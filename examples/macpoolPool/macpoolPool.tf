resource "ucs_mac_pool_pool" "test_mac_pool" {
  name        = "tf_test_macpool"
  org_org_dn  = "org-root"
  description = "mac poo created with terraform"
}
