
terraform {
  required_providers {
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = ">= 2.0"
    }
  }
}

provider "kubernetes" {
    config_path = "~/.kube/config"
}

resource "kubernetes_config_map" "lightnode" {
    metadata {
        name = "lightnode-config"
        namespace = "testnet"
    }

    data = {
        "server.config.hcl" = "${file("${path.module}/config/server.config.hcl")}"
    }
}

resource "kubernetes_persistent_volume_claim" "lightnode-data" {
    metadata {
        name = "lightnode-data-pvc"
        namespace = "testnet"
    }
    spec {
        access_modes = ["ReadWriteMany"]
        resources {
            requests = {
                storage = "5Gi"
            }
        }
        volume_name = "${kubernetes_persistent_volume.lightnode-data.metadata.0.name}"
    }
}

resource "kubernetes_persistent_volume_claim" "lightnode-log" {
    metadata {
        name = "lightnode-log-pvc"
        namespace = "testnet"
    }
    spec {
        access_modes = ["ReadWriteMany"]
        resources {
            requests = {
                storage = "5Gi"
            }
        }
        volume_name = "${kubernetes_persistent_volume.lightnode-log.metadata.0.name}"
    }
}

resource "kubernetes_persistent_volume_claim" "lightnode-raft" {
    metadata {
        name = "lightnode-raft-pvc"
        namespace = "testnet"
    }
    spec {
        access_modes = ["ReadWriteMany"]
        resources {
            requests = {
                storage = "5Gi"
            }
        }
        volume_name = "${kubernetes_persistent_volume.lightnode-raft.metadata.0.name}"
    }
}

resource "kubernetes_persistent_volume" "lightnode-data" {
    metadata {
        name = "lightnode-data-pv"
    }
    spec {
        capacity = {
            storage = "5Gi"
        }
        access_modes = ["ReadWriteMany"]
        persistent_volume_reclaim_policy = "Retain"
        storage_class_name = "standard"
        persistent_volume_source {
          host_path {
            path = "/vagrant/lightnode/data"
          }
        }
    }
}

resource "kubernetes_persistent_volume" "lightnode-log" {
    metadata {
        name = "lightnode-log-pv"
    }
    spec {
        capacity = {
            storage = "5Gi"
        }
        access_modes = ["ReadWriteMany"]
        persistent_volume_reclaim_policy = "Retain"
        storage_class_name = "standard"
        persistent_volume_source {
          host_path {
            path = "/vagrant/lightnode/log"
          }
        }
    }
}

resource "kubernetes_persistent_volume" "lightnode-raft" {
    metadata {
        name = "lightnode-raft-pv"
    }
    spec {
        capacity = {
            storage = "5Gi"
        }
        access_modes = ["ReadWriteMany"]
        persistent_volume_reclaim_policy = "Retain"
        storage_class_name = "standard"
        persistent_volume_source {
          host_path {
            path = "/vagrant/lightnode/raft"
          }
        }
    }
}

resource "kubernetes_deployment" "lightnode" {
    metadata {
        name = "lightnode"
        labels = {
            app = "LightnodeApp"
        }
    }
    spec {
        selector {
            match_labels = {
                app = "LightnodeApp"
            }
        }
        template {
            metadata {
                labels = {
                    app = "LightnodeApp"
                }
            }
            spec {
                volume {
                    name = "lightnode-config"
                    config_map {
                      name = kubernetes_config_map.lightnode.metadata.0.name
                      optional = false
                    }
                }
                volume {
                  name = kubernetes_persistent_volume.lightnode-data.metadata.0.name
                  persistent_volume_claim {
                    claim_name = kubernetes_persistent_volume.lightnode-data.metadata.0.name
                  }
                }
                volume {
                  name = kubernetes_persistent_volume.lightnode-log.metadata.0.name
                  persistent_volume_claim {
                    claim_name = kubernetes_persistent_volume.lightnode-log.metadata.0.name
                  }
                }
                volume {
                  name = kubernetes_persistent_volume.lightnode-raft.metadata.0.name
                  persistent_volume_claim {
                    claim_name = kubernetes_persistent_volume.lightnode-raft.metadata.0.name
                  }
                }

                container {
                    image = "decentralized/structx:latest"
                    name = "lightnode"

                    resources {
                        limits = {
                            cpu = "0.25"
                            memory = "512Mi"
                        }
                    }

                    env {
                        name = "DSERVICE_CONFIG"
                        value = "/local/lightnode/server.config.hcl"
                    }

                    volume_mount {
                      name = "lightnode-config"
                      mount_path = "/local/lightnode"
                    }

                    volume_mount {
                        name = kubernetes_persistent_volume.lightnode-data.metadata.0.name
                        mount_path = "/opt/lightnode/data"
                    }
                    volume_mount {
                        name = kubernetes_persistent_volume.lightnode-log.metadata.0.name
                        mount_path = "/var/log/lightnode"
                    }
                    volume_mount {
                        name = kubernetes_persistent_volume.lightnode-data.metadata.0.name
                        mount_path = "/opt/lightnode/raft"
                    }

                    port {
                        container_port = 8080
                        name = "http"
                    }
                    port {
                        container_port = 50051
                        name = "messenger-tcp"
                    }
                    port {
                        container_port = 50053
                        name = "raft-tcp"
                    }

                    liveness_probe {
                        http_get {
                            path = "/health"
                            port = "http"
                        }
                    }
                }
            }
        }
    }
}