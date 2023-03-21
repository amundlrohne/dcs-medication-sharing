variable "mongodb_atlas_username" {
  description = "mongodb atlas username"
}

variable "mongodb_atlas_password" {
  description = "mongodb atlas password"
}

variable "mongodbatlas_public_key" {
  description = "mongodb atlas public key"
}

variable "mongodbatlas_private_key" {
  description = "mongodb atlas private key"
}

provider "mongodbatlas" {
  public_key  = var.mongodbatlas_public_key
  private_key = var.mongodbatlas_private_key
}

resource "mongodbatlas_advanced_cluster" "dcs-medication-sharing" {
  project_id   = "64199ea462a42d47918b8293"
  name         = "dcs-medication-sharing"
  cluster_type = "REPLICASET"
  replication_specs {
    region_configs {
      electable_specs {
        instance_size = "M0"
        node_count    = 1
      }
      analytics_specs {
        instance_size = "M0"
        node_count    = 1
      }
      provider_name         = "TENANT"
      backing_provider_name = "GCP"
      priority              = 1
      region_name           = "WESTERN_EUROPE"
    }
  }
}

output "mongodbatlas_connection_url" {
  value = mongodbatlas_advanced_cluster.dcs-medication-sharing.connection_strings[0].standard_srv
}
