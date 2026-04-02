# ============================================================================
# Google Pub/Sub Topic Module - Variables
# ============================================================================

variable "environment" {
  description = "Deployment environment. One of: devl, test, prod."
  type        = string

  validation {
    condition     = contains(["devl", "test", "prod"], var.environment)
    error_message = "environment must be one of: devl, test, prod."
  }
}

variable "project_code" {
  description = "Short identifier used in resource naming for governance and cost attribution."
  type        = string

  validation {
    condition     = length(var.project_code) > 0
    error_message = "project_code must not be empty."
  }
}

variable "region" {
  description = "GCP region where the Pub/Sub topic will be created."
  type        = string
  default     = "us-central1"
}

variable "pubsub_config" {
  description = "Configuration object for the Google Pub/Sub topic."
  type = object({
    base_name = string
  })

  validation {
    condition     = can(regex("^[a-zA-Z0-9][a-zA-Z0-9-]{0,29}$", var.pubsub_config.base_name))
    error_message = "base_name must be alphanumeric or dashes, start with a letter or digit, and be at most 30 characters."
  }
}
