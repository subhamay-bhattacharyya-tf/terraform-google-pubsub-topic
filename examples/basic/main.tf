module "pubsub_topic" {
  source = "../../"

  environment  = var.environment
  project_code = var.project_code
  region       = var.region

  pubsub_config = {
    base_name = var.base_name
    location  = var.location
  }
}
