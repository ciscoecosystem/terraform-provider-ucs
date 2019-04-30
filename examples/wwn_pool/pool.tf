resource "ucs_fcpool_initiators" "pool_init" {
    org_org_dn = "org-root"
    name = "tf_init"
  
}

resource "ucs_fcpool_block" "pool_block" {
    fcpool_initiators_dn = "${ucs_fcpool_initiators.pool_init.id}"
    r_from = "20:00:00:25:B5:00:00:00"
    to = "20:00:00:25:B5:00:00:02"
}

