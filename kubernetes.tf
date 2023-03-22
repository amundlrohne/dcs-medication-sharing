# Configure kubernetes provider with Oauth2 access token.
# https://registry.terraform.io/providers/hashicorp/google/latest/docs/data-sources/client_config
# This fetches a new token, which will expire in 1 hour.
data "google_client_config" "default" {}

data "google_container_cluster" "primary" {
  name     = google_container_cluster.primary.name
  location = var.region
}

provider "kubernetes" {
  host = "https://${google_container_cluster.primary.endpoint}"

  token                  = data.google_client_config.default.access_token
  cluster_ca_certificate = base64decode(data.google_container_cluster.primary.master_auth[0].cluster_ca_certificate)
}

resource "kubernetes_deployment" "consent" {
  metadata {
    name = "consent-service"
    labels = {
      App = "ConsentService"
    }
  }

  spec {
    replicas = 1
    selector {
      match_labels = {
        App = "ConsentService"
      }
    }
    template {
      metadata {
        labels = {
          App = "ConsentService"
        }
      }
      spec {
        container {
          image = "ghcr.io/amundlrohne/dcs-medication-sharing/consent:latest"
          name  = "consent"

          port {
            container_port = 8080
          }

          env {
            name  = "MONGO_URL"
            value = mongodbatlas_advanced_cluster.dcs-medication-sharing.connection_strings[0].standard_srv
          }

          env {
            name  = "MONGO_USERNAME"
            value = var.mongodb_atlas_username
          }

          env {
            name  = "MONGO_PASSWORD"
            value = var.mongodb_atlas_password
          }

          env {
            name  = "MONGO_DB_NAME"
            value = "consent"
          }

          resources {
            limits = {
              cpu    = "0.5"
              memory = "512Mi"
            }
            requests = {
              cpu    = "250m"
              memory = "50Mi"
            }
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "consent" {
  metadata {
    name = "consent-service"
  }
  spec {
    selector = {
      App = kubernetes_deployment.consent.spec.0.template.0.metadata[0].labels.App
    }
    port {
      port        = 8180
      target_port = 8080
    }

    type = "LoadBalancer"
  }
}

resource "kubernetes_deployment" "healthcare-provider" {
  metadata {
    name = "healthcare-provider-service"
    labels = {
      App = "HealthcareProviderService"
    }
  }

  spec {
    replicas = 1
    selector {
      match_labels = {
        App = "HealthcareProviderService"
      }
    }
    template {
      metadata {
        labels = {
          App = "HealthcareProviderService"
        }
      }
      spec {
        container {
          image = "ghcr.io/amundlrohne/dcs-medication-sharing/healthcare-provider:latest"
          name  = "healthcare-provider"

          port {
            container_port = 8080
          }

          env {
            name  = "MONGO_URL"
            value = mongodbatlas_advanced_cluster.dcs-medication-sharing.connection_strings[0].standard_srv
          }

          env {
            name  = "MONGO_USERNAME"
            value = var.mongodb_atlas_username
          }

          env {
            name  = "MONGO_PASSWORD"
            value = var.mongodb_atlas_password
          }

          env {
            name  = "MONGO_DB_NAME"
            value = "healthcare-provider"
          }
          resources {
            limits = {
              cpu    = "0.5"
              memory = "512Mi"
            }
            requests = {
              cpu    = "250m"
              memory = "50Mi"
            }
          }
        }
      }
    }
  }
}


resource "kubernetes_service" "healthcare-provider" {
  metadata {
    name = "healthcare-provider-service"
  }
  spec {
    selector = {
      App = kubernetes_deployment.healthcare-provider.spec.0.template.0.metadata[0].labels.App
    }
    port {
      port        = 8280
      target_port = 8080
    }

    type = "LoadBalancer"
  }
}

resource "kubernetes_deployment" "medication-record" {
  metadata {
    name = "medication-record-service"
    labels = {
      App = "MedicationRecordService"
    }
  }

  spec {
    replicas = 1
    selector {
      match_labels = {
        App = "MedicationRecordService"
      }
    }
    template {
      metadata {
        labels = {
          App = "MedicationRecordService"
        }
      }
      spec {
        container {
          image = "ghcr.io/amundlrohne/dcs-medication-sharing/medication-record:latest"
          name  = "medication-record"

          port {
            container_port = 8080
          }

          env {
            name  = "GCP_SERVICE_ACCOUNT_PRIVATE_KEY"
            value = base64decode(google_service_account_key.medication-record-service-key.private_key)
          }

          env {
            name  = "PRODUCTION"
            value = true
          }

          resources {
            limits = {
              cpu    = "0.5"
              memory = "512Mi"
            }
            requests = {
              cpu    = "250m"
              memory = "50Mi"
            }
          }
        }
      }
    }
  }
}


resource "kubernetes_service" "medication-record" {
  metadata {
    name = "medication-record-service"
  }
  spec {
    selector = {
      App = kubernetes_deployment.medication-record.spec.0.template.0.metadata[0].labels.App
    }
    port {
      port        = 8380
      target_port = 8080
    }

    type = "LoadBalancer"
  }
}

resource "kubernetes_deployment" "standardization" {
  metadata {
    name = "standardization-service"
    labels = {
      App = "StandardizationService"
    }
  }

  spec {
    replicas = 1
    selector {
      match_labels = {
        App = "StandardizationService"
      }
    }
    template {
      metadata {
        labels = {
          App = "StandardizationService"
        }
      }
      spec {
        container {
          image = "ghcr.io/amundlrohne/dcs-medication-sharing/standardization:latest"
          name  = "standardization"

          port {
            container_port = 8080
          }

          resources {
            limits = {
              cpu    = "0.5"
              memory = "512Mi"
            }
            requests = {
              cpu    = "250m"
              memory = "50Mi"
            }
          }
        }
      }
    }
  }
}


resource "kubernetes_service" "standardization" {
  metadata {
    name = "standardization-service"
  }
  spec {
    selector = {
      App = kubernetes_deployment.standardization.spec.0.template.0.metadata[0].labels.App
    }
    port {
      port        = 8480
      target_port = 8080
    }

    type = "LoadBalancer"
  }
}

resource "kubernetes_deployment" "react-frontend" {
  metadata {
    name = "react-frontend-service"
    labels = {
      App = "ReactFrontendService"
    }
  }

  spec {
    replicas = 1
    selector {
      match_labels = {
        App = "ReactFrontendService"
      }
    }
    template {
      metadata {
        labels = {
          App = "ReactFrontendService"
        }
      }
      spec {
        container {
          image = "ghcr.io/amundlrohne/dcs-medication-sharing/react-frontend:latest"
          name  = "react-frontend"

          port {
            container_port = 3000
          }
          env {
            name  = "REACT_APP_PRODUCTION"
            value = true
          }

          resources {
            limits = {
              cpu    = "0.5"
              memory = "512Mi"
            }
            requests = {
              cpu    = "250m"
              memory = "50Mi"
            }
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "react-frontend" {
  metadata {
    name = "react-frontend-service"
  }
  spec {
    selector = {
      App = kubernetes_deployment.react-frontend.spec.0.template.0.metadata[0].labels.App
    }
    port {
      port        = 8080
      target_port = 3000
    }

    type = "LoadBalancer"
  }
}

resource "kubernetes_ingress_v1" "dcs-medication-sharing-ingress" {
  metadata {
    name = "dcs-medication-sharing-ingress"
  }

  spec {
    default_backend {
      service {
        name = kubernetes_deployment.react-frontend.metadata[0].name
        port {
          number = 8080
        }
      }
    }


    rule {
      http {
        path {
          backend {
            service {
              name = kubernetes_deployment.consent.metadata[0].name
              port {
                number = 8180
              }
            }
          }

          path = "/consent/*"
        }

        path {
          backend {
            service {
              name = kubernetes_deployment.healthcare-provider.metadata[0].name
              port {
                number = 8280
              }
            }
          }

          path = "/healthcare-provider/*"
        }
        path {
          backend {
            service {
              name = kubernetes_deployment.medication-record.metadata[0].name
              port {
                number = 8380
              }
            }
          }

          path = "/medication-record/*"
        }
        path {
          backend {
            service {
              name = kubernetes_deployment.standardization.metadata[0].name
              port {
                number = 8480
              }
            }
          }

          path = "/standardization/*"
        }
      }
    }
  }
}
