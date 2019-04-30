resource "ucs_fabric_vsan" "test_vsan" {
    name = "tf_vsan"
    fabric_vsan_id = "2"
    fcoe_vlan = "2"
  
}
