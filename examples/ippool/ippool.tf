resource "ucs_ip_pool_pool" "test_ip_pool" {
  name        = "tf_test_ip"
  org_org_dn  = "org-root"
  description = "ip pool created with terraform"
}

resource "ucs_ip_pool_block" "test_ip_pool_block" {
  name           = "tf_test_ip_pool_block"
  ippool_pool_dn = "${ucs_ip_pool_pool.test_ip_pool.id}"
  r_from         = "192.168.10.102"
  to             = "192.168.10.104"
}

resource "ucs_ip_pool_ipv6_block" "name" {
  ippool_pool_dn = "${ucs_ip_pool_pool.test_ip_pool.id}"
  name           = "tf_test_ipv6_block"
  r_from         = "2002::"
  to             = "2002::"
}
