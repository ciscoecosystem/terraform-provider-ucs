resource "ucs_vnic_lan_conn_policy" "vnic_lan_policy" {
  org_org_dn = "org-root"
  name = "test_vlan_policy"
}

resource "ucs_vnic_ether" "vnic_ether" {
  vnic_lan_conn_policy_dn = "${ucs_vnic_lan_conn_policy.vnic_lan_policy.id}"
  name = "test_vnic_ether"
}

resource "ucs_vnic_iscsi_lcp" "vnic_iscsi_lcp" {
  vnic_lan_conn_policy_dn = "${ucs_vnic_lan_conn_policy.vnic_lan_policy.id}"
  name = "test_vnic_lcp"
}

resource "ucs_vnic_vlan" "vnic_vlan" {
  vnic_i_scsi_lcp_dn = "${ucs_vnic_iscsi_lcp.vnic_iscsi_lcp.id}"
  name = "test_vlan"
}

