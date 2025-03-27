variable "public_ip" {
  type        = string
  description = "Public IP that will host the API server hostname."
}

variable "public_domain" {
  type        = string
  description = "Public domain where the 'hostname' will be created."
}

variable "hostname" {
  type        = string
  description = "Hostname, within 'public_domain' that is associated with the 'public_ip'."
}

#
# namecheap authentication information
#
variable "namecheap_user_name" {
  type        = string
  description = "User name associated with the account hosting the DNS records for 'public_hostname'."
  sensitive   = true
}

variable "namecheap_api_user_name" {
  type        = string
  default     = null
  description = "API User name associated with the account hosting the DNS records for 'public_hostname'.  Defaults to 'namecheap_user_name' if not provided."
  sensitive   = true
  nullable    = true
}

variable "namecheap_api_key" {
  type        = string
  description = "API key used for updating DNS records associated with the account hosting the DNS records for 'public_hostname'."
  sensitive   = true
}
