package exo

import (
	"errors"

	"github.com/hashicorp/hcl2/gohcl"
	"github.com/hashicorp/hcl2/hclparse"
)

type Config struct {
	Servers []*ServerConfig `hcl:"server,block"`
}

type ServerConfig struct {
	Name       string `hcl:"name,label"`
	ListenAddr string `hcl:"listen_addr"`
	Routes     Routes `hcl:"routes,block"`

	// TLSOn              bool    `hcl:"tls_on,optional"`
	TLSCertificateFile *string `hcl:"tls_cert,optional"`
	TLSPrivateLKeyFile *string `hcl:"tls_key,optional"`
	TLSCertificate     *string `hcl:"tls_cert_bytes,optional"`
	TLSPrivateLKey     *string `hcl:"tls_key_bytes,optional"`
}

type Routes struct {
	Routes []RouteConfig `hcl:"path,block"`
}

type RouteConfig struct {
	Match        string              `hcl:"match,label"`
	FileServer   *FileServerConfig   `hcl:"file_server,block"`
	ReverseProxy *ReverseProxyConfig `hcl:"reverse_proxy,block"`
	Headers      *Headers            `hcl:"headers,optional"`
}

type FileServerConfig struct {
	Root    string   `hcl:"root"`
	Bundle  bool     `hcl:"bundle,optional"`
	Headers *Headers `hcl:"headers,optional"`
}

type ReverseProxyConfig struct {
	Upstreams     []string `hcl:"upstreams"`
	LoadBalancing string   `hcl:"load_balancing,optional"`
	Headers       *Headers `hcl:"headers,optional"`
}

type Headers map[string]string

func LoadConfig(config string, cfg *Config) error {
	parser := hclparse.NewParser()
	file, diags := parser.ParseHCLFile(config)
	if diags.HasErrors() {
		return errors.New(diags.Error())
	}

	diags = gohcl.DecodeBody(file.Body, nil, cfg)
	if diags.HasErrors() {
		return errors.New(diags.Error())
	}

	return nil
}

func LoadConfigFromBytes(config []byte, cfg *Config) error {
	parser := hclparse.NewParser()
	file, diags := parser.ParseHCL(config, "byte.hcl")
	if diags.HasErrors() {
		return errors.New(diags.Error())
	}

	diags = gohcl.DecodeBody(file.Body, nil, cfg)
	if diags.HasErrors() {
		return errors.New(diags.Error())
	}

	return nil
}
