resource "ucs_lstorage_profile" "test_storage_profile" {
    org_org_dn = "org-root"
    name = "tf_profile"
  
}

resource "ucs_lstorage_das_scsi_lun" "test_das_scsi" {
  lstorage_profile_dn = "${ucs_lstorage_profile.test_storage_profile.id}"
  name = "tf_test"
}
