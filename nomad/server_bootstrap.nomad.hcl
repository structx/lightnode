
job "olivia" {
    datacenters = ["dc1"]
    type = "service"

    group "blockchain" {
        count = 1

        network {
            mode = "bridge"

            port "rpc" {}

            port "http" {}
        }
        service {
            name = "olivia-grpc"
            port = "rpc"

            provider = "consul"

            connect {
                sidecar_service {
                    proxy {
                        upstreams {
                            destination_name = "mora"
                            local_bind_port = 8081
                        }
                    }
                }
            }
        }

        service {
            name = "olivia-api"
            port = "http"

            provider = "consul"

            connect {
                sidecar_service {}
            }
            
        }

        task "server" {
            driver = "docker"

            config {
                image = "trevatk/olivia:v0.0.1"
                ports = ["http", "rpc"]
            }

            env {
                DSERVICE_CONFIG = "${NOMAD_TASK_DIR}/node.config.hcl"
            }

            template {
                data = <<EOH
server {
    bind_addr = "0.0.0.0"

    ports {
        http = {{ env "NOMAD_PORT_http"}}
        grpc = {{ env "NOMAD_PORT_rpc" }}
    }
}

raft {
    bootstrap = true
    local_id = "12345566777434"
    base_dir = "/opt/olivia/raft"
}

logger {
    log_path = "/var/log/olivia/node.log"
    log_level = "DEBUG"
    raft_log_path = "/var/log/olivia"
}

chain {
    base_dir = "/opt/olivia/data"
}

message_broker {
    {{- range service "mora-rpc"}}
    server_addr = "{{ .Address}}:{{ .Port}};{{- end}}"
}
                EOH
                
                destination = "local/node.config.hcl"
                change_mode = "restart"
                change_signal = "SIGTERM"
            }


        }
    }
}