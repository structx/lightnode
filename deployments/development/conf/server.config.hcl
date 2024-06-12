
server {
  bind_addr = "0.0.0.0"
  default_timeout = 15

  ports {
    http = 8080
    grpc = 50051
  }
}

chain {
  base_dir = "/opt/lightnode/data"
}

raft {
  bootstrap = true
  local_id = "1"
  base_dir = "/opt/lightnode/raft"
}

logger {
  log_path = "/var/log/lightnode/lightnode.log"
  log_level = "DEBUG"
  raft_log_path = "/var/log"
}