package exo

import (
	"crypto/tls"
	"errors"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"regexp"
	"strings"
	"sync"

	stdslog "golang.org/x/exp/slog"
)

var (
	jsonHandler = stdslog.HandlerOptions{Level: stdslog.LevelInfo}.NewJSONHandler(os.Stdout)
	slog        = stdslog.New(jsonHandler)
)

func StartServers(config *Config) error {
	servers := StructureHandlers(config)

	wg := new(sync.WaitGroup)
	for listener, serverConfig := range *servers {
		wg.Add(1)
		server, err := createServer(listener, serverConfig)
		if err != nil {
			return err
		}

		go func(wg *sync.WaitGroup, serverConfig *Server) {
			if len(serverConfig.TLSCerts) > 0 {
				if err = server.ListenAndServeTLS("", ""); err != nil {
					log.Fatal(err)
				}
			} else {

				if err = server.ListenAndServe(); err != nil {
					log.Fatal(err)
				}
			}
			defer wg.Done()
		}(wg, serverConfig)
	}

	wg.Wait()
	return nil
}

type Servers map[string]*Server
type Server struct {
	Hosts    map[string][]RouteConfig
	TLSCerts []tls.Certificate
	Handler  http.HandlerFunc
}

// A listener may need to listen for multiple domains
func StructureHandlers(config *Config) *Servers {
	var servers = Servers{}

	for _, host := range config.Servers {
		hostname := strings.ToLower(host.Name)
		if _, ok := servers[host.ListenAddr]; !ok {
			servers[host.ListenAddr] = &Server{
				Hosts: map[string][]RouteConfig{},
			}

			if host.TLSCertificateFile != nil && host.TLSPrivateLKeyFile != nil {
				certificate, err := tls.LoadX509KeyPair(*host.TLSCertificateFile, *host.TLSPrivateLKeyFile)
				if err != nil {
					log.Fatal(err)
				}
				servers[host.ListenAddr].TLSCerts = append(servers[host.ListenAddr].TLSCerts, certificate)
			}

			if host.TLSCertificate != nil && host.TLSPrivateLKey != nil {
				certificate, err := tls.X509KeyPair([]byte(*host.TLSCertificate), []byte(*host.TLSPrivateLKey))
				if err != nil {
					log.Fatal(err)
				}
				servers[host.ListenAddr].TLSCerts = append(servers[host.ListenAddr].TLSCerts, certificate)
			}

		}

		for _, routes := range host.Routes.Routes {
			if _, ok := servers[host.ListenAddr].Hosts[hostname]; !ok {
				servers[host.ListenAddr].Hosts[hostname] = []RouteConfig{}
			}

			servers[host.ListenAddr].Hosts[hostname] = append(servers[host.ListenAddr].Hosts[hostname], routes)
		}

		// servers[host.ListenAddr].Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 	http.NotFound(w, r)
		// })
	}

	return &servers
}

func createServer(listener string, config *Server) (*http.Server, error) {
	slog.Debug("Creating server for listener " + listener)
	server := &http.Server{
		Addr: listener,
		Handler: Logger(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

				hostname := r.Host

				slog.WithGroup("debugCtx").LogAttrs(
					r.Context(),
					stdslog.LevelDebug,
					"request pre-processing",
					[]stdslog.Attr{
						stdslog.Any("request-url", r.URL.String()),
						stdslog.Any("request-host", r.Host),
						stdslog.Any("route-host", hostname),
						stdslog.Any("request-path", r.RequestURI),
						// stdslog.Any("wanted-hosts", MapKeys(config.Hosts)),
					}...,
				)

				if _, ok := config.Hosts[strings.ToLower(hostname)]; !ok {
					http.Error(w, errors.New("host not found").Error(), http.StatusNotFound)
					return
				}
				for _, routeCfg := range config.Hosts[hostname] {
					if routeCfg.Headers != nil {
						for headerK, heaverV := range *routeCfg.Headers {
							r.Header.Add(headerK, heaverV)
						}
					}

					re := regexp.MustCompile(routeCfg.Match)
					if re.MatchString(r.URL.Path) {

						// Run file server
						if routeCfg.FileServer != nil {
							fs := http.FileServer(http.Dir(routeCfg.FileServer.Root))
							handler := http.StripPrefix(routeCfg.Match, fs)
							handler.ServeHTTP(w, r)
							return
						}

						if routeCfg.ReverseProxy != nil {
							// Use upstream default config
							upstream, err := url.Parse(routeCfg.ReverseProxy.Upstreams[0])
							if err != nil {
								log.Fatal(err)
							}

							proxy := httputil.NewSingleHostReverseProxy(upstream)
							proxy.ServeHTTP(w, r)
							return
						}

						// error out
						http.Error(w, errors.New("resource not found").Error(), http.StatusNotFound)
						return
					}

				}

				// error out
				http.Error(w, errors.New("Not found").Error(), http.StatusNotFound)
			}),
		),
	}

	if len(config.TLSCerts) > 0 {
		server.TLSConfig = &tls.Config{Certificates: config.TLSCerts}
	}

	return server, nil
}

func MapKeys(arr map[string][]RouteConfig) []any {
	var out []any
	for k := range arr {
		out = append(out, k)
	}

	return out
}
