resource "namecheap_domain_records" "service_info" {
  domain = var.public_domain

  record {
    hostname = var.hostname
    type     = "A"
    address  = var.public_ip
  }
}
