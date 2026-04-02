# ============================================================================
# Google Pub/Sub Topic Module - Main
# Creates and manages a Google Pub/Sub Topic.
# ============================================================================

resource "google_pubsub_topic" "this" {
  name   = local.topic_name
  labels = local.labels

  message_storage_policy {
    allowed_persistence_regions = [var.pubsub_config.location]
  }
}
