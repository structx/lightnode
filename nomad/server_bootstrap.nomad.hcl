
job "level0" {
    datacenters = ["dc1"]
    type = "service"

    group "structx" {
        count = 1

        network {
            mode = "bridge"

            port "metrics" {
                to = 2112
            }

            port "rpc" {}

            port "http" {}
        }

        template {
            data =<< EOH
            server {
                bootstrap = true
                listen_addr = ":{{ env 'NOMAD_PORT_http' }}"
            }
            EOH
            destination = "/etc/server/config.hcl"
        }

        task "server" {
            driver = "docker"

            config {
                image = ""
            }

            env {
                __CONFIG = "/etc/server/config.hcl"
            }

            resources {}
        }
    }
}