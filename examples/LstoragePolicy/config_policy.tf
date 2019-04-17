resource "ucs_disk_group_config_policy" "test_disk_policy" {
  org_org_dn = "org-root"
  name       = "test-policy"
  raid_level = "simple"
}

resource "ucs_disk_group_qualifier" "test_disk_qualifier" {
    lstorage_disk_group_config_policy_dn = "${ucs_disk_group_config_policy.test_disk_policy.id}"
    name = "test-qualifier"
  
}

# resource "ucs_disk_config_reference" "test-ref" {
#     lstorage_disk_group_config_policy_dn = "${ucs_disk_group_config_policy.test_disk_policy.id}"
#     slot_num = "1"
#     name = "test-ref"
  
# }


