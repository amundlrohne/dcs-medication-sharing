resource "google_healthcare_dataset" "dataset" {
  name     = "${var.project_id}-dataset"
  location = var.region
}

resource "google_healthcare_fhir_store" "fhir-store" {
  name    = "${var.project_id}-fhir-store"
  dataset = google_healthcare_dataset.dataset.id
  version = "R4"

  #enable_update_create          = false
  #disable_referential_integrity = false
  disable_resource_versioning = false
  enable_history_import       = false
}

