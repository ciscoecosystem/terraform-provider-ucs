resource "ucs_mac_pool_pool_block" "test_macpool_block" {
  macpool_pool_dn = "${ucs_mac_pool_pool.test_mac_pool.id}"
  r_from          = "00:25:B5:00:00:00"
  to              = "00:25:B5:00:00:01"
  name            = "test_mac_pool_block"
}
