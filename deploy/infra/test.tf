#
# variables
#
variable "namecheap_user_name" {
  type        = string
  description = "User name associated with the account hosting the DNS records for 'public_hostname'."
  sensitive   = true
}

variable "namecheap_api_key" {
  type        = string
  description = "API key used for updating DNS records associated with the account hosting the DNS records for 'public_hostname'."
  sensitive   = true
}

locals {
  base_domain = "klearwave.io"
}

#
# test environments
# 
module "test_sandbox" {
  source = "./sandbox"

  public_ip     = "10.10.10.10"
  public_domain = local.base_domain
  hostname      = "test.api.sandbox"

  namecheap_user_name = var.namecheap_user_name
  namecheap_api_key   = var.namecheap_api_key
}
