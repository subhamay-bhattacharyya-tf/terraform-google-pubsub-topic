# ============================================================================
# Google Pub/Sub Topic Module - Locals
# ============================================================================

locals {
  topic_name = "${var.project_code}-${var.pubsub_config.base_name}-${var.pubsub_config.location}-${var.environment}"

  labels = {
    environment  = var.environment
    project_code = var.project_code
  }
}
