resource "ucs_ls_server" "test_ls_server" {
    org_org_dn = "org-root"
    name = "tf_ls_server"
  
}


resource "ucs_ls_requirement" "test_req" {
    ls_server_dn = "${ucs_ls_server.test_ls_server.id}"
  
}

resource "ucs_mgmt_interface" "test_mgmt_interface" {
    ls_server_dn = "${ucs_ls_server.test_ls_server.id}"
    mode = "in-band"
  
}

resource "ucs_mgmt_vnet" "test_mgmt_vnet" {
    mgmt_interface_dn = "${ucs_mgmt_interface.test_mgmt_interface.id}"
  
}


