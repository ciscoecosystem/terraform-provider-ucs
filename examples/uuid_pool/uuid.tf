resource "ucs_uuid_pool" "test_uuid_pool" {
  org_org_dn = "org-root"
  name = "tf_test_uuid"
}

resource "ucs_uuid_pool_block" "test_uuid_pool_block" {
    uuidpool_pool_dn = "${ucs_uuid_pool.test_uuid_pool.id}"
    r_from = "0000-000000000001"
    to = "0000-000000000003"

}

