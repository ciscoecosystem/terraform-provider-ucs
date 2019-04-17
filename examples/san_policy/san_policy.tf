resource "ucs_vnic_san_policy" "tf_san_policy" {
    org_org_dn = "org-root"
    name = "tf_test_san"
  
}

resource "ucs_vnic_fc" "tf_vnic_fc" {
    vnic_san_conn_policy_dn = "${ucs_vnic_san_policy.tf_san_policy.id}"
    name = "test-fc"
  
}

resource "ucs_vnic_fc_node" "tf_vnic_fc_node" {
    vnic_san_conn_policy_dn = "${ucs_vnic_san_policy.tf_san_policy.id}"
    name = "test_node"
  
}

resource "ucs_vnic_fc_if" "tf_vnic_fcif" {
    vnic_fc_dn = "${ucs_vnic_fc.tf_vnic_fc.id}"
    name = "test_fcif"
  
}

