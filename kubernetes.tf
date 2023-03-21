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
  #client_certificate     = google_container_cluster.primary.master_auth.0.client_certificate
  #client_key             = google_container_cluster.primary.master_auth.0.client_key
  #cluster_ca_certificate = google_container_cluster.primary.master_auth.0.cluster_ca_certificate
}

resource "kubernetes_deployment" "consent" {
  metadata {
    name = "consent-service"
    labels = {
      App = "ConsentService"
    }
  }

  spec {
    replicas = 2
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
            container_port = 8180
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
      target_port = 8180
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
    replicas = 2
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
            container_port = 8380
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
      target_port = 8280
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
    replicas = 2
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
            container_port = 8380
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
      target_port = 8380
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
    replicas = 2
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
            container_port = 8480
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
      target_port = 8480
    }

    type = "LoadBalancer"
  }
}
