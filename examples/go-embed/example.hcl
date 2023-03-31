server "localhost" {
  listen_addr = "0.0.0.0:8081"

  tls_key_bytes = "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDP3nKZ7pJjqkYl\n2MgzJz/TDy5XJByrp7uKQ5LUGgF+Lw3q50TYh3i/SGEEW3GorsNORl5SLYw3tcYX\nJiDRIxhpMng/43AZLsr5GJIVYk+hMHwt6z40+rZqIqixTT6Y3scJTQHtLrJz/svd\nrQf8UTDfKk7+OVq9Jjca1KIE7YDeHI3biI+qAqwyv8wh6gEJgQYJ6DP3liFX3Mws\nKyXCOZxkAJXxOgmco5ih5vxtlWMypJSs9J9pt6IjJbJsu3pdfHTRXt83MCUAH7AL\nUfnX0X81rMc9OeiiJVdni42f1fIT+S/Y6i4VPDJrsavQzpxcldLkfW1Nyu3Phep7\ncpLNrcx9AgMBAAECggEAfus0j/ZR3ZamtA8T1d/eIXFAeyZqdwi4AmjV6rliNfhn\nCAljSM4WlLyNuApZcIQcbdubVZPH7HAJXjMCkI9cJUcSkve8hzPB7Kvq7jGTqzie\naa8b0V9PJ6i31WHzmYVUg8JILdh4T0jAkz5GBPJR48DCcIj2dU0zEifIMTgOmg49\najzVmbM8VrzhfEbd3hHS/seFz9vun8vlhS259lVxfypIKg2sPVuc+ZaDauy6IjE+\ng5quc/wGaPfJu/yPjZdv1zfZ0OngaYMqXpopXPqMYZpcCpCQ/QLwwWUJdr8MnVNi\n69gQRYaivOf37/3qh6nrekireIJIKmAh2W7MUqqXZQKBgQDsar7LLm0D2r+Xjw8v\npBgb1gZUqiZ4z9ImVGDj+rIcMrXCi/oHD91hj2FnPDz2a+DkD3bUwFwoXnx5joUd\nQ2sadrTx7gahabYNVmSTjY89wkALpBgNS61+IPYV5fxdxmXD7+iyx2VLX8KcPXOg\nL2bDJe6AtwsC7ZeJBunzp8N66wKBgQDhFlZX0Oq2pH7Z0PeA/dJkRXw16ssOyTeO\nlK16yb2CqHsjaiBFA65X8sYW+mObh9XPTjLA+kv1fdAOtzXwH8ZGOsEoB7KOEHgB\nXfyod0kjHm8sLV+ih58S7B+EGa7FSXzf8smuWMRkgAbNp6zVJ2DaZl3OEjfzjbEq\nwIoYPSosNwKBgQC/3LlOJzUthZ8rHlySqzctj/m/R/7E2NORa7TeD5vwKm87nWGh\nzYw6GqRPpaFL0qzGKI6lhh7EXDnkZHc/EnGfHmsmU6dp42JtXIlu/dkbo8NaLyed\n0qbPz2wmxWiTqej4pIBSZmOcUfCb33swqodZJDV0nBgBnlP2Tnga1zFbVQKBgAER\ncRFznVmGiE+STUpgafot1jwPRC1qyQe1LumRQYP4NrVhUIvjS5WBQm1jfMuFJn+Y\nWGdQCMI+eZlIR0o/bFpy9u2ws6k27Mrc3lESV9eB0nc8M+L6j8ewNSoUUR8AjUYm\nycw5AZ2UefPJ4ukUCcKfj7xlm7xdJMUWKtGYoLWbAoGBAIvs0/xgieN/sLnWMN2H\npBcdkD92FiGE3QnHNLA/jAsAay4j5ysmJeP+49jlkzuXIEis6NLkTHSndimCuNQn\ncfCfMoItyfVJkKJ103hBgaPy0+JDQN0LyyN5zjuaHLeSKy0qCorIiPMNH97VbYS9\nD4O1KagNJBoOh1YVQxC76oU0\n-----END PRIVATE KEY-----\n"
  tls_cert_bytes = "-----BEGIN CERTIFICATE-----\nMIIDeDCCAmACCQDn2yaGHqfn8jANBgkqhkiG9w0BAQsFADB+MQswCQYDVQQGEwJV\nUzETMBEGA1UECAwKQ2FsaWZvcm5pYTEWMBQGA1UEBwwNTW91bnRhaW4gVmlldzEa\nMBgGA1UECgwRWW91ciBPcmdhbml6YXRpb24xEjAQBgNVBAsMCVlvdXIgVW5pdDES\nMBAGA1UEAwwJbG9jYWxob3N0MB4XDTIyMTIzMDA2NTQ0MFoXDTMyMTIyNzA2NTQ0\nMFowfjELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExFjAUBgNVBAcM\nDU1vdW50YWluIFZpZXcxGjAYBgNVBAoMEVlvdXIgT3JnYW5pemF0aW9uMRIwEAYD\nVQQLDAlZb3VyIFVuaXQxEjAQBgNVBAMMCWxvY2FsaG9zdDCCASIwDQYJKoZIhvcN\nAQEBBQADggEPADCCAQoCggEBAM/ecpnukmOqRiXYyDMnP9MPLlckHKunu4pDktQa\nAX4vDernRNiHeL9IYQRbcaiuw05GXlItjDe1xhcmINEjGGkyeD/jcBkuyvkYkhVi\nT6EwfC3rPjT6tmoiqLFNPpjexwlNAe0usnP+y92tB/xRMN8qTv45Wr0mNxrUogTt\ngN4cjduIj6oCrDK/zCHqAQmBBgnoM/eWIVfczCwrJcI5nGQAlfE6CZyjmKHm/G2V\nYzKklKz0n2m3oiMlsmy7el18dNFe3zcwJQAfsAtR+dfRfzWsxz056KIlV2eLjZ/V\n8hP5L9jqLhU8Mmuxq9DOnFyV0uR9bU3K7c+F6ntyks2tzH0CAwEAATANBgkqhkiG\n9w0BAQsFAAOCAQEAW68AbIgWOW3wAB+HboNOSVsfGc9zMpA97XaACbZq9zkGVDTe\n0k1YmYlcxF/9PNqw5FQXNlSuNaIi3Be5mb0m30Vrp+vPg78r9oxaeYnt2ZLs/KcS\nLiNU8X4gMCT2jxot74EJAG7MWP/1i0EIYmFNnesSnRv9hJfbczG+MIoBNw2Oqlfs\nukXUH1ash5E9PVzKZINEcLFJXbJNrzwYSrYAnEPDV/DxlTM671pCOiSeovt3VJET\niPqMO7ldgQhxw6A74XidRrtzHw7/vlc+XEuP3L6PYi/BAWaP+o3fdXz2x1XIc72E\nQyDpMWpQRZhDiCUp7T/zM/OXSrAkdUaDQEE2ig==\n-----END CERTIFICATE-----"

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

        headers = {
          test = "test"
        }
      }
    }
  }
}
