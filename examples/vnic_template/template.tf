resource "ucs_vnic_lan_conn_template" "test_temp" {
    org_org_dn = "org-root"
    name = "tf_test_template"
}

resource "ucs_vnic_ether_if" "test_ether_if" {
    vnic_lan_conn_templ_dn = "${ucs_vnic_lan_conn_template.test_temp.id}"
    name = "tf_ether"
  
}
