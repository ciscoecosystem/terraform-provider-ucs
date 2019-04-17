package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsVnicFc() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsVnicFcCreate,
		Update: resourceUcsVnicFcUpdate,
		Read:   resourceUcsVnicFcRead,
		Delete: resourceUcsVnicFcDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsVnicFcImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"vnic_san_conn_policy_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"adaptor_profile_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"admin_cdn_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"admin_host_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"admin_vcon": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"boot_dev": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"cdn_prop_in_sync": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"cdn_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"child_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"config_qualifier": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"config_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"equipment_dn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"flt_aggr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"ident_pool_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"inst_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"is_supported": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"max_data_field_size": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"node_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"nw_templ_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_adaptor_profile_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_cdn_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_host_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_ident_pool_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_nw_templ_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_order": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_pin_to_group_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_qos_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_speed": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_stats_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_vcon": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"order": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"owner": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"pers_bind": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"pers_bind_clear": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"pin_to_group_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"qos_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"redundancy_pair_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"redundancy_peer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"sacl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"stats_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"switch_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		}),
	}
}

func getRemoteVnicFc(client *client.Client, dn string) (*models.VnicFc, error) {
	vnicFcDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	vnicFc := models.VnicFcFromDoc(vnicFcDoc, "configResolveDn")

	if vnicFc.DistinguishedName == "" {
		return nil, fmt.Errorf("VnicFc %s not found", dn)
	}

	return vnicFc, nil
}

func setVnicFcAttributes(vnicFc *models.VnicFc, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(vnicFc.DistinguishedName)
	d.Set("description", vnicFc.Description)
	d.Set("vnic_san_conn_policy_dn", GetParentDn(vnicFc.DistinguishedName))
	vnicFcMap, _ := vnicFc.ToMap()

	d.Set("adaptor_profile_name", vnicFcMap["adaptorProfileName"])

	d.Set("addr", vnicFcMap["addr"])

	d.Set("admin_cdn_name", vnicFcMap["adminCdnName"])

	d.Set("admin_host_port", vnicFcMap["adminHostPort"])

	d.Set("admin_vcon", vnicFcMap["adminVcon"])

	d.Set("boot_dev", vnicFcMap["bootDev"])

	d.Set("cdn_prop_in_sync", vnicFcMap["cdnPropInSync"])

	d.Set("cdn_source", vnicFcMap["cdnSource"])

	d.Set("child_action", vnicFcMap["childAction"])

	d.Set("config_qualifier", vnicFcMap["configQualifier"])

	d.Set("config_state", vnicFcMap["configState"])

	d.Set("equipment_dn", vnicFcMap["equipmentDn"])

	d.Set("flt_aggr", vnicFcMap["fltAggr"])

	d.Set("ident_pool_name", vnicFcMap["identPoolName"])

	d.Set("inst_type", vnicFcMap["instType"])

	d.Set("is_supported", vnicFcMap["isSupported"])

	d.Set("max_data_field_size", vnicFcMap["maxDataFieldSize"])

	d.Set("node_addr", vnicFcMap["nodeAddr"])

	d.Set("nw_templ_name", vnicFcMap["nwTemplName"])

	d.Set("oper_adaptor_profile_name", vnicFcMap["operAdaptorProfileName"])

	d.Set("oper_cdn_name", vnicFcMap["operCdnName"])

	d.Set("oper_host_port", vnicFcMap["operHostPort"])

	d.Set("oper_ident_pool_name", vnicFcMap["operIdentPoolName"])

	d.Set("oper_nw_templ_name", vnicFcMap["operNwTemplName"])

	d.Set("oper_order", vnicFcMap["operOrder"])

	d.Set("oper_pin_to_group_name", vnicFcMap["operPinToGroupName"])

	d.Set("oper_qos_policy_name", vnicFcMap["operQosPolicyName"])

	d.Set("oper_speed", vnicFcMap["operSpeed"])

	d.Set("oper_stats_policy_name", vnicFcMap["operStatsPolicyName"])

	d.Set("oper_vcon", vnicFcMap["operVcon"])

	d.Set("order", vnicFcMap["order"])

	d.Set("owner", vnicFcMap["owner"])

	d.Set("pers_bind", vnicFcMap["persBind"])

	d.Set("pers_bind_clear", vnicFcMap["persBindClear"])

	d.Set("pin_to_group_name", vnicFcMap["pinToGroupName"])

	d.Set("qos_policy_name", vnicFcMap["qosPolicyName"])

	d.Set("redundancy_pair_type", vnicFcMap["redundancyPairType"])

	d.Set("redundancy_peer", vnicFcMap["redundancyPeer"])

	d.Set("sacl", vnicFcMap["sacl"])

	d.Set("stats_policy_name", vnicFcMap["statsPolicyName"])

	d.Set("switch_id", vnicFcMap["switchId"])

	d.Set("type", vnicFcMap["type"])
	return d
}

func resourceUcsVnicFcImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	vnicFc, err := getRemoteVnicFc(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setVnicFcAttributes(vnicFc, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsVnicFcCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	VnicSanConnPolicy := d.Get("vnic_san_conn_policy_dn").(string)

	vnicFcAttr := models.VnicFcAttributes{}

	if Adaptor_profile_name, ok := d.GetOk("adaptor_profile_name"); ok {
		vnicFcAttr.Adaptor_profile_name = Adaptor_profile_name.(string)
	}

	if Addr, ok := d.GetOk("addr"); ok {
		vnicFcAttr.Addr = Addr.(string)
	}

	if Admin_cdn_name, ok := d.GetOk("admin_cdn_name"); ok {
		vnicFcAttr.Admin_cdn_name = Admin_cdn_name.(string)
	}

	if Admin_host_port, ok := d.GetOk("admin_host_port"); ok {
		vnicFcAttr.Admin_host_port = Admin_host_port.(string)
	}

	if Admin_vcon, ok := d.GetOk("admin_vcon"); ok {
		vnicFcAttr.Admin_vcon = Admin_vcon.(string)
	}

	if Boot_dev, ok := d.GetOk("boot_dev"); ok {
		vnicFcAttr.Boot_dev = Boot_dev.(string)
	}

	if Cdn_prop_in_sync, ok := d.GetOk("cdn_prop_in_sync"); ok {
		vnicFcAttr.Cdn_prop_in_sync = Cdn_prop_in_sync.(string)
	}

	if Cdn_source, ok := d.GetOk("cdn_source"); ok {
		vnicFcAttr.Cdn_source = Cdn_source.(string)
	}

	if Child_action, ok := d.GetOk("child_action"); ok {
		vnicFcAttr.Child_action = Child_action.(string)
	}

	if Config_qualifier, ok := d.GetOk("config_qualifier"); ok {
		vnicFcAttr.Config_qualifier = Config_qualifier.(string)
	}

	if Config_state, ok := d.GetOk("config_state"); ok {
		vnicFcAttr.Config_state = Config_state.(string)
	}

	if Equipment_dn, ok := d.GetOk("equipment_dn"); ok {
		vnicFcAttr.Equipment_dn = Equipment_dn.(string)
	}

	if Flt_aggr, ok := d.GetOk("flt_aggr"); ok {
		vnicFcAttr.Flt_aggr = Flt_aggr.(string)
	}

	if Ident_pool_name, ok := d.GetOk("ident_pool_name"); ok {
		vnicFcAttr.Ident_pool_name = Ident_pool_name.(string)
	}

	if Inst_type, ok := d.GetOk("inst_type"); ok {
		vnicFcAttr.Inst_type = Inst_type.(string)
	}

	if Is_supported, ok := d.GetOk("is_supported"); ok {
		vnicFcAttr.Is_supported = Is_supported.(string)
	}

	if Max_data_field_size, ok := d.GetOk("max_data_field_size"); ok {
		vnicFcAttr.Max_data_field_size = Max_data_field_size.(string)
	}

	if Node_addr, ok := d.GetOk("node_addr"); ok {
		vnicFcAttr.Node_addr = Node_addr.(string)
	}

	if Nw_templ_name, ok := d.GetOk("nw_templ_name"); ok {
		vnicFcAttr.Nw_templ_name = Nw_templ_name.(string)
	}

	if Oper_adaptor_profile_name, ok := d.GetOk("oper_adaptor_profile_name"); ok {
		vnicFcAttr.Oper_adaptor_profile_name = Oper_adaptor_profile_name.(string)
	}

	if Oper_cdn_name, ok := d.GetOk("oper_cdn_name"); ok {
		vnicFcAttr.Oper_cdn_name = Oper_cdn_name.(string)
	}

	if Oper_host_port, ok := d.GetOk("oper_host_port"); ok {
		vnicFcAttr.Oper_host_port = Oper_host_port.(string)
	}

	if Oper_ident_pool_name, ok := d.GetOk("oper_ident_pool_name"); ok {
		vnicFcAttr.Oper_ident_pool_name = Oper_ident_pool_name.(string)
	}

	if Oper_nw_templ_name, ok := d.GetOk("oper_nw_templ_name"); ok {
		vnicFcAttr.Oper_nw_templ_name = Oper_nw_templ_name.(string)
	}

	if Oper_order, ok := d.GetOk("oper_order"); ok {
		vnicFcAttr.Oper_order = Oper_order.(string)
	}

	if Oper_pin_to_group_name, ok := d.GetOk("oper_pin_to_group_name"); ok {
		vnicFcAttr.Oper_pin_to_group_name = Oper_pin_to_group_name.(string)
	}

	if Oper_qos_policy_name, ok := d.GetOk("oper_qos_policy_name"); ok {
		vnicFcAttr.Oper_qos_policy_name = Oper_qos_policy_name.(string)
	}

	if Oper_speed, ok := d.GetOk("oper_speed"); ok {
		vnicFcAttr.Oper_speed = Oper_speed.(string)
	}

	if Oper_stats_policy_name, ok := d.GetOk("oper_stats_policy_name"); ok {
		vnicFcAttr.Oper_stats_policy_name = Oper_stats_policy_name.(string)
	}

	if Oper_vcon, ok := d.GetOk("oper_vcon"); ok {
		vnicFcAttr.Oper_vcon = Oper_vcon.(string)
	}

	if Order, ok := d.GetOk("order"); ok {
		vnicFcAttr.Order = Order.(string)
	}

	if Owner, ok := d.GetOk("owner"); ok {
		vnicFcAttr.Owner = Owner.(string)
	}

	if Pers_bind, ok := d.GetOk("pers_bind"); ok {
		vnicFcAttr.Pers_bind = Pers_bind.(string)
	}

	if Pers_bind_clear, ok := d.GetOk("pers_bind_clear"); ok {
		vnicFcAttr.Pers_bind_clear = Pers_bind_clear.(string)
	}

	if Pin_to_group_name, ok := d.GetOk("pin_to_group_name"); ok {
		vnicFcAttr.Pin_to_group_name = Pin_to_group_name.(string)
	}

	if Qos_policy_name, ok := d.GetOk("qos_policy_name"); ok {
		vnicFcAttr.Qos_policy_name = Qos_policy_name.(string)
	}

	if Redundancy_pair_type, ok := d.GetOk("redundancy_pair_type"); ok {
		vnicFcAttr.Redundancy_pair_type = Redundancy_pair_type.(string)
	}

	if Redundancy_peer, ok := d.GetOk("redundancy_peer"); ok {
		vnicFcAttr.Redundancy_peer = Redundancy_peer.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		vnicFcAttr.Sacl = Sacl.(string)
	}

	if Stats_policy_name, ok := d.GetOk("stats_policy_name"); ok {
		vnicFcAttr.Stats_policy_name = Stats_policy_name.(string)
	}

	if Switch_id, ok := d.GetOk("switch_id"); ok {
		vnicFcAttr.Switch_id = Switch_id.(string)
	}

	if Type, ok := d.GetOk("type"); ok {
		vnicFcAttr.Type = Type.(string)
	}

	vnicFc := models.NewVnicFc(fmt.Sprintf("fc-%s", Name), VnicSanConnPolicy, desc, vnicFcAttr)

	err := ucsClient.Save(vnicFc)
	if err != nil {
		return err
	}

	d.SetId(vnicFc.DistinguishedName)
	return resourceUcsVnicFcRead(d, m)
}

func resourceUcsVnicFcRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	vnicFc, err := getRemoteVnicFc(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setVnicFcAttributes(vnicFc, d)

	return nil
}

func resourceUcsVnicFcDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "vnicFc")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsVnicFcUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	VnicSanConnPolicy := d.Get("vnic_san_conn_policy_dn").(string)

	vnicFcAttr := models.VnicFcAttributes{}
	if Adaptor_profile_name, ok := d.GetOk("adaptor_profile_name"); ok {
		vnicFcAttr.Adaptor_profile_name = Adaptor_profile_name.(string)
	}
	if Addr, ok := d.GetOk("addr"); ok {
		vnicFcAttr.Addr = Addr.(string)
	}
	if Admin_cdn_name, ok := d.GetOk("admin_cdn_name"); ok {
		vnicFcAttr.Admin_cdn_name = Admin_cdn_name.(string)
	}
	if Admin_host_port, ok := d.GetOk("admin_host_port"); ok {
		vnicFcAttr.Admin_host_port = Admin_host_port.(string)
	}
	if Admin_vcon, ok := d.GetOk("admin_vcon"); ok {
		vnicFcAttr.Admin_vcon = Admin_vcon.(string)
	}
	if Boot_dev, ok := d.GetOk("boot_dev"); ok {
		vnicFcAttr.Boot_dev = Boot_dev.(string)
	}
	if Cdn_prop_in_sync, ok := d.GetOk("cdn_prop_in_sync"); ok {
		vnicFcAttr.Cdn_prop_in_sync = Cdn_prop_in_sync.(string)
	}
	if Cdn_source, ok := d.GetOk("cdn_source"); ok {
		vnicFcAttr.Cdn_source = Cdn_source.(string)
	}
	if Child_action, ok := d.GetOk("child_action"); ok {
		vnicFcAttr.Child_action = Child_action.(string)
	}
	if Config_qualifier, ok := d.GetOk("config_qualifier"); ok {
		vnicFcAttr.Config_qualifier = Config_qualifier.(string)
	}
	if Config_state, ok := d.GetOk("config_state"); ok {
		vnicFcAttr.Config_state = Config_state.(string)
	}
	if Equipment_dn, ok := d.GetOk("equipment_dn"); ok {
		vnicFcAttr.Equipment_dn = Equipment_dn.(string)
	}
	if Flt_aggr, ok := d.GetOk("flt_aggr"); ok {
		vnicFcAttr.Flt_aggr = Flt_aggr.(string)
	}
	if Ident_pool_name, ok := d.GetOk("ident_pool_name"); ok {
		vnicFcAttr.Ident_pool_name = Ident_pool_name.(string)
	}
	if Inst_type, ok := d.GetOk("inst_type"); ok {
		vnicFcAttr.Inst_type = Inst_type.(string)
	}
	if Is_supported, ok := d.GetOk("is_supported"); ok {
		vnicFcAttr.Is_supported = Is_supported.(string)
	}
	if Max_data_field_size, ok := d.GetOk("max_data_field_size"); ok {
		vnicFcAttr.Max_data_field_size = Max_data_field_size.(string)
	}
	if Node_addr, ok := d.GetOk("node_addr"); ok {
		vnicFcAttr.Node_addr = Node_addr.(string)
	}
	if Nw_templ_name, ok := d.GetOk("nw_templ_name"); ok {
		vnicFcAttr.Nw_templ_name = Nw_templ_name.(string)
	}
	if Oper_adaptor_profile_name, ok := d.GetOk("oper_adaptor_profile_name"); ok {
		vnicFcAttr.Oper_adaptor_profile_name = Oper_adaptor_profile_name.(string)
	}
	if Oper_cdn_name, ok := d.GetOk("oper_cdn_name"); ok {
		vnicFcAttr.Oper_cdn_name = Oper_cdn_name.(string)
	}
	if Oper_host_port, ok := d.GetOk("oper_host_port"); ok {
		vnicFcAttr.Oper_host_port = Oper_host_port.(string)
	}
	if Oper_ident_pool_name, ok := d.GetOk("oper_ident_pool_name"); ok {
		vnicFcAttr.Oper_ident_pool_name = Oper_ident_pool_name.(string)
	}
	if Oper_nw_templ_name, ok := d.GetOk("oper_nw_templ_name"); ok {
		vnicFcAttr.Oper_nw_templ_name = Oper_nw_templ_name.(string)
	}
	if Oper_order, ok := d.GetOk("oper_order"); ok {
		vnicFcAttr.Oper_order = Oper_order.(string)
	}
	if Oper_pin_to_group_name, ok := d.GetOk("oper_pin_to_group_name"); ok {
		vnicFcAttr.Oper_pin_to_group_name = Oper_pin_to_group_name.(string)
	}
	if Oper_qos_policy_name, ok := d.GetOk("oper_qos_policy_name"); ok {
		vnicFcAttr.Oper_qos_policy_name = Oper_qos_policy_name.(string)
	}
	if Oper_speed, ok := d.GetOk("oper_speed"); ok {
		vnicFcAttr.Oper_speed = Oper_speed.(string)
	}
	if Oper_stats_policy_name, ok := d.GetOk("oper_stats_policy_name"); ok {
		vnicFcAttr.Oper_stats_policy_name = Oper_stats_policy_name.(string)
	}
	if Oper_vcon, ok := d.GetOk("oper_vcon"); ok {
		vnicFcAttr.Oper_vcon = Oper_vcon.(string)
	}
	if Order, ok := d.GetOk("order"); ok {
		vnicFcAttr.Order = Order.(string)
	}
	if Owner, ok := d.GetOk("owner"); ok {
		vnicFcAttr.Owner = Owner.(string)
	}
	if Pers_bind, ok := d.GetOk("pers_bind"); ok {
		vnicFcAttr.Pers_bind = Pers_bind.(string)
	}
	if Pers_bind_clear, ok := d.GetOk("pers_bind_clear"); ok {
		vnicFcAttr.Pers_bind_clear = Pers_bind_clear.(string)
	}
	if Pin_to_group_name, ok := d.GetOk("pin_to_group_name"); ok {
		vnicFcAttr.Pin_to_group_name = Pin_to_group_name.(string)
	}
	if Qos_policy_name, ok := d.GetOk("qos_policy_name"); ok {
		vnicFcAttr.Qos_policy_name = Qos_policy_name.(string)
	}
	if Redundancy_pair_type, ok := d.GetOk("redundancy_pair_type"); ok {
		vnicFcAttr.Redundancy_pair_type = Redundancy_pair_type.(string)
	}
	if Redundancy_peer, ok := d.GetOk("redundancy_peer"); ok {
		vnicFcAttr.Redundancy_peer = Redundancy_peer.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		vnicFcAttr.Sacl = Sacl.(string)
	}
	if Stats_policy_name, ok := d.GetOk("stats_policy_name"); ok {
		vnicFcAttr.Stats_policy_name = Stats_policy_name.(string)
	}
	if Switch_id, ok := d.GetOk("switch_id"); ok {
		vnicFcAttr.Switch_id = Switch_id.(string)
	}
	if Type, ok := d.GetOk("type"); ok {
		vnicFcAttr.Type = Type.(string)
	}

	vnicFc := models.NewVnicFc(fmt.Sprintf("fc-%s", Name), VnicSanConnPolicy, desc, vnicFcAttr)
	vnicFc.Status = "modified"
	err := ucsClient.Save(vnicFc)
	if err != nil {
		return err
	}

	d.SetId(vnicFc.DistinguishedName)
	return resourceUcsVnicFcRead(d, m)
}
