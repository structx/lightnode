
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
    "lightnode_config" = "${file("${path.module}/config/server.config.hcl")}"
  }
}

resource "kubernetes_persistent_volume" "lightnode-data" {
  metadata {
    name = "lightnode-data"
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
    name = "lightnode-log"
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
    name = "lightnode-raft"
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
    namespace = "testnet"
    labels = {
      app = var.app_name
    }
  }
  spec {
    selector {
      match_labels = {
        app = var.app_name
      }
    }
    template {
      metadata {
        labels = {
          app = var.app_name
        }
      }
      spec {
        security_context {
          run_as_user = "1000"
          run_as_group = "1000"
        }
        volume {
          name = "lightnode-config"
          config_map {
            name = kubernetes_config_map.lightnode.metadata.0.name
            optional = false
          }
        }
        volume {
          name = kubernetes_persistent_volume.lightnode-data.metadata.0.name
          host_path {
            path = "/vagrant/lightnode/data"
            type = "DirectoryOrCreate"
          }
        }
        volume {
          name = kubernetes_persistent_volume.lightnode-log.metadata.0.name
          host_path {
            path = "/vagrant/lightnode/log"
            type = "DirectoryOrCreate"
          }
        }
        volume {
          name = kubernetes_persistent_volume.lightnode-raft.metadata.0.name
          host_path {
            path = "/vagrant/lightnode/raft"
            type = "DirectoryOrCreate"
          }
        }

        container {
          image = "registry.structx.local/structx/lightnode:v0.0.1"
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
            mount_path = "/local/lightnode/server.config.hcl"
            sub_path = "lightnode_config"
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

resource "kubernetes_service" "lightnode" {
  metadata {
    name = kubernetes_deployment.lightnode.metadata.0.name
    namespace = "testnet"
  }
  spec {
    selector = {
      app = var.app_name
    }
    port {
      name = "http"
      port = 80
      target_port = 8080
    }
    type = "ClusterIP"
  }
}

resource "kubernetes_ingress_v1" "lightnode" {
  metadata {
    name = kubernetes_deployment.lightnode.metadata.0.name
    namespace = "testnet"
    annotations = {
      "traefik.ingress.kubernetes.io/router.entrypoints" = "web"
    }
  }
  spec {
    default_backend {
      service {
        name = kubernetes_service.lightnode.metadata.0.name
        port {
          name = "http"
        }
      }
    }
    rule {
      host = "lightnode.structx.localnet"
      http {
        path {
          backend {
            service {
              name = kubernetes_service.lightnode.metadata.0.name
              port {
                name = "http"
              }
            }
          }
        }
      }
    }
  }
}