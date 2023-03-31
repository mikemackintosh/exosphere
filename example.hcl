server "accounts.zush.dev" {
  listen_addr = "0.0.0.0:8081"

  tls_key = "archive/accounts.zush.dev/privkey1.pem"
  tls_cert = "archive/accounts.zush.dev/fullchain1.pem"

  routes {
    path "/static" {
      file_server {
        root = "/var/www/html"
      }
    }
    path "/*" {
      reverse_proxy {
        upstreams = [
          "https://google.com",
        ]

        load_balancing = "round_robin"

        headers = {
          test = "test"
        }
      }
    }
  }
}

server "auth.zush.dev" {
  listen_addr = "0.0.0.0:8081"

  tls_key = "archive/auth.zush.dev/privkey1.pem"
  tls_cert = "archive/auth.zush.dev/fullchain1.pem"

  routes {
    path "/static" {
      file_server {
        root = "./www"
      }
    }
    path "/(test)*" {
      reverse_proxy {
        upstreams = [
          "https://google.com",
        ]

        load_balancing = "round_robin"

        headers = {
          test = "test"
        }
      }
    }
  }
}

server "artifacts.zush.dev" {
  listen_addr = "0.0.0.0:8082"

  tls_key = "archive/artifacts.zush.dev/privkey1.pem"
  tls_cert = "archive/artifacts.zush.dev/fullchain1.pem"

  routes {
    path "/static" {
      file_server {
        bundle = true
        root = "/var/www/html"
      }
    }
    path "*" {
      reverse_proxy {
        upstreams = [
          "http://localhost:8000",
          "http://localhost:8001",
        ]

        load_balancing = "round_robin"

        headers = {
          test = "test"
        }
      }
    }
  }
}
