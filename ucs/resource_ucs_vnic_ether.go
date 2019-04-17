package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsVnicEther() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsVnicEtherCreate,
		Update: resourceUcsVnicEtherUpdate,
		Read:   resourceUcsVnicEtherRead,
		Delete: resourceUcsVnicEtherDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsVnicEtherImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"vnic_lan_conn_policy_dn": &schema.Schema{
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

			"dynamic_id": &schema.Schema{
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

			"mtu": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"nw_ctrl_policy_name": &schema.Schema{
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

			"oper_nw_ctrl_policy_name": &schema.Schema{
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

			"pf_dn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"pin_to_group_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"prop_acl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"purpose": &schema.Schema{
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

			"virtualization_preference": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		}),
	}
}

func getRemoteVnicEther(client *client.Client, dn string) (*models.VnicEther, error) {
	vnicEtherDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	vnicEther := models.VnicEtherFromDoc(vnicEtherDoc, "configResolveDn")

	if vnicEther.DistinguishedName == "" {
		return nil, fmt.Errorf("VnicEther %s not found", dn)
	}

	return vnicEther, nil
}

func setVnicEtherAttributes(vnicEther *models.VnicEther, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(vnicEther.DistinguishedName)
	d.Set("description", vnicEther.Description)
	d.Set("vnic_lan_conn_policy_dn", GetParentDn(vnicEther.DistinguishedName))
	vnicEtherMap, _ := vnicEther.ToMap()

	d.Set("adaptor_profile_name", vnicEtherMap["adaptorProfileName"])

	d.Set("addr", vnicEtherMap["addr"])

	d.Set("admin_cdn_name", vnicEtherMap["adminCdnName"])

	d.Set("admin_host_port", vnicEtherMap["adminHostPort"])

	d.Set("admin_vcon", vnicEtherMap["adminVcon"])

	d.Set("boot_dev", vnicEtherMap["bootDev"])

	d.Set("cdn_prop_in_sync", vnicEtherMap["cdnPropInSync"])

	d.Set("cdn_source", vnicEtherMap["cdnSource"])

	d.Set("child_action", vnicEtherMap["childAction"])

	d.Set("config_qualifier", vnicEtherMap["configQualifier"])

	d.Set("config_state", vnicEtherMap["configState"])

	d.Set("dynamic_id", vnicEtherMap["dynamicId"])

	d.Set("equipment_dn", vnicEtherMap["equipmentDn"])

	d.Set("flt_aggr", vnicEtherMap["fltAggr"])

	d.Set("ident_pool_name", vnicEtherMap["identPoolName"])

	d.Set("inst_type", vnicEtherMap["instType"])

	d.Set("mtu", vnicEtherMap["mtu"])

	d.Set("nw_ctrl_policy_name", vnicEtherMap["nwCtrlPolicyName"])

	d.Set("nw_templ_name", vnicEtherMap["nwTemplName"])

	d.Set("oper_adaptor_profile_name", vnicEtherMap["operAdaptorProfileName"])

	d.Set("oper_cdn_name", vnicEtherMap["operCdnName"])

	d.Set("oper_host_port", vnicEtherMap["operHostPort"])

	d.Set("oper_ident_pool_name", vnicEtherMap["operIdentPoolName"])

	d.Set("oper_nw_ctrl_policy_name", vnicEtherMap["operNwCtrlPolicyName"])

	d.Set("oper_nw_templ_name", vnicEtherMap["operNwTemplName"])

	d.Set("oper_order", vnicEtherMap["operOrder"])

	d.Set("oper_pin_to_group_name", vnicEtherMap["operPinToGroupName"])

	d.Set("oper_qos_policy_name", vnicEtherMap["operQosPolicyName"])

	d.Set("oper_speed", vnicEtherMap["operSpeed"])

	d.Set("oper_stats_policy_name", vnicEtherMap["operStatsPolicyName"])

	d.Set("oper_vcon", vnicEtherMap["operVcon"])

	d.Set("order", vnicEtherMap["order"])

	d.Set("owner", vnicEtherMap["owner"])

	d.Set("pf_dn", vnicEtherMap["pfDn"])

	d.Set("pin_to_group_name", vnicEtherMap["pinToGroupName"])

	d.Set("prop_acl", vnicEtherMap["propAcl"])

	d.Set("purpose", vnicEtherMap["purpose"])

	d.Set("qos_policy_name", vnicEtherMap["qosPolicyName"])

	d.Set("redundancy_pair_type", vnicEtherMap["redundancyPairType"])

	d.Set("redundancy_peer", vnicEtherMap["redundancyPeer"])

	d.Set("sacl", vnicEtherMap["sacl"])

	d.Set("stats_policy_name", vnicEtherMap["statsPolicyName"])

	d.Set("switch_id", vnicEtherMap["switchId"])

	d.Set("type", vnicEtherMap["type"])

	d.Set("virtualization_preference", vnicEtherMap["virtualizationPreference"])
	return d
}

func resourceUcsVnicEtherImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	vnicEther, err := getRemoteVnicEther(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setVnicEtherAttributes(vnicEther, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsVnicEtherCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	VnicLanConnPolicy := d.Get("vnic_lan_conn_policy_dn").(string)

	vnicEtherAttr := models.VnicEtherAttributes{}

	if Adaptor_profile_name, ok := d.GetOk("adaptor_profile_name"); ok {
		vnicEtherAttr.Adaptor_profile_name = Adaptor_profile_name.(string)
	}

	if Addr, ok := d.GetOk("addr"); ok {
		vnicEtherAttr.Addr = Addr.(string)
	}

	if Admin_cdn_name, ok := d.GetOk("admin_cdn_name"); ok {
		vnicEtherAttr.Admin_cdn_name = Admin_cdn_name.(string)
	}

	if Admin_host_port, ok := d.GetOk("admin_host_port"); ok {
		vnicEtherAttr.Admin_host_port = Admin_host_port.(string)
	}

	if Admin_vcon, ok := d.GetOk("admin_vcon"); ok {
		vnicEtherAttr.Admin_vcon = Admin_vcon.(string)
	}

	if Boot_dev, ok := d.GetOk("boot_dev"); ok {
		vnicEtherAttr.Boot_dev = Boot_dev.(string)
	}

	if Cdn_prop_in_sync, ok := d.GetOk("cdn_prop_in_sync"); ok {
		vnicEtherAttr.Cdn_prop_in_sync = Cdn_prop_in_sync.(string)
	}

	if Cdn_source, ok := d.GetOk("cdn_source"); ok {
		vnicEtherAttr.Cdn_source = Cdn_source.(string)
	}

	if Child_action, ok := d.GetOk("child_action"); ok {
		vnicEtherAttr.Child_action = Child_action.(string)
	}

	if Config_qualifier, ok := d.GetOk("config_qualifier"); ok {
		vnicEtherAttr.Config_qualifier = Config_qualifier.(string)
	}

	if Config_state, ok := d.GetOk("config_state"); ok {
		vnicEtherAttr.Config_state = Config_state.(string)
	}

	if Dynamic_id, ok := d.GetOk("dynamic_id"); ok {
		vnicEtherAttr.Dynamic_id = Dynamic_id.(string)
	}

	if Equipment_dn, ok := d.GetOk("equipment_dn"); ok {
		vnicEtherAttr.Equipment_dn = Equipment_dn.(string)
	}

	if Flt_aggr, ok := d.GetOk("flt_aggr"); ok {
		vnicEtherAttr.Flt_aggr = Flt_aggr.(string)
	}

	if Ident_pool_name, ok := d.GetOk("ident_pool_name"); ok {
		vnicEtherAttr.Ident_pool_name = Ident_pool_name.(string)
	}

	if Inst_type, ok := d.GetOk("inst_type"); ok {
		vnicEtherAttr.Inst_type = Inst_type.(string)
	}

	if Mtu, ok := d.GetOk("mtu"); ok {
		vnicEtherAttr.Mtu = Mtu.(string)
	}

	if Nw_ctrl_policy_name, ok := d.GetOk("nw_ctrl_policy_name"); ok {
		vnicEtherAttr.Nw_ctrl_policy_name = Nw_ctrl_policy_name.(string)
	}

	if Nw_templ_name, ok := d.GetOk("nw_templ_name"); ok {
		vnicEtherAttr.Nw_templ_name = Nw_templ_name.(string)
	}

	if Oper_adaptor_profile_name, ok := d.GetOk("oper_adaptor_profile_name"); ok {
		vnicEtherAttr.Oper_adaptor_profile_name = Oper_adaptor_profile_name.(string)
	}

	if Oper_cdn_name, ok := d.GetOk("oper_cdn_name"); ok {
		vnicEtherAttr.Oper_cdn_name = Oper_cdn_name.(string)
	}

	if Oper_host_port, ok := d.GetOk("oper_host_port"); ok {
		vnicEtherAttr.Oper_host_port = Oper_host_port.(string)
	}

	if Oper_ident_pool_name, ok := d.GetOk("oper_ident_pool_name"); ok {
		vnicEtherAttr.Oper_ident_pool_name = Oper_ident_pool_name.(string)
	}

	if Oper_nw_ctrl_policy_name, ok := d.GetOk("oper_nw_ctrl_policy_name"); ok {
		vnicEtherAttr.Oper_nw_ctrl_policy_name = Oper_nw_ctrl_policy_name.(string)
	}

	if Oper_nw_templ_name, ok := d.GetOk("oper_nw_templ_name"); ok {
		vnicEtherAttr.Oper_nw_templ_name = Oper_nw_templ_name.(string)
	}

	if Oper_order, ok := d.GetOk("oper_order"); ok {
		vnicEtherAttr.Oper_order = Oper_order.(string)
	}

	if Oper_pin_to_group_name, ok := d.GetOk("oper_pin_to_group_name"); ok {
		vnicEtherAttr.Oper_pin_to_group_name = Oper_pin_to_group_name.(string)
	}

	if Oper_qos_policy_name, ok := d.GetOk("oper_qos_policy_name"); ok {
		vnicEtherAttr.Oper_qos_policy_name = Oper_qos_policy_name.(string)
	}

	if Oper_speed, ok := d.GetOk("oper_speed"); ok {
		vnicEtherAttr.Oper_speed = Oper_speed.(string)
	}

	if Oper_stats_policy_name, ok := d.GetOk("oper_stats_policy_name"); ok {
		vnicEtherAttr.Oper_stats_policy_name = Oper_stats_policy_name.(string)
	}

	if Oper_vcon, ok := d.GetOk("oper_vcon"); ok {
		vnicEtherAttr.Oper_vcon = Oper_vcon.(string)
	}

	if Order, ok := d.GetOk("order"); ok {
		vnicEtherAttr.Order = Order.(string)
	}

	if Owner, ok := d.GetOk("owner"); ok {
		vnicEtherAttr.Owner = Owner.(string)
	}

	if Pf_dn, ok := d.GetOk("pf_dn"); ok {
		vnicEtherAttr.Pf_dn = Pf_dn.(string)
	}

	if Pin_to_group_name, ok := d.GetOk("pin_to_group_name"); ok {
		vnicEtherAttr.Pin_to_group_name = Pin_to_group_name.(string)
	}

	if Prop_acl, ok := d.GetOk("prop_acl"); ok {
		vnicEtherAttr.Prop_acl = Prop_acl.(string)
	}

	if Purpose, ok := d.GetOk("purpose"); ok {
		vnicEtherAttr.Purpose = Purpose.(string)
	}

	if Qos_policy_name, ok := d.GetOk("qos_policy_name"); ok {
		vnicEtherAttr.Qos_policy_name = Qos_policy_name.(string)
	}

	if Redundancy_pair_type, ok := d.GetOk("redundancy_pair_type"); ok {
		vnicEtherAttr.Redundancy_pair_type = Redundancy_pair_type.(string)
	}

	if Redundancy_peer, ok := d.GetOk("redundancy_peer"); ok {
		vnicEtherAttr.Redundancy_peer = Redundancy_peer.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		vnicEtherAttr.Sacl = Sacl.(string)
	}

	if Stats_policy_name, ok := d.GetOk("stats_policy_name"); ok {
		vnicEtherAttr.Stats_policy_name = Stats_policy_name.(string)
	}

	if Switch_id, ok := d.GetOk("switch_id"); ok {
		vnicEtherAttr.Switch_id = Switch_id.(string)
	}

	if Type, ok := d.GetOk("type"); ok {
		vnicEtherAttr.Type = Type.(string)
	}

	if Virtualization_preference, ok := d.GetOk("virtualization_preference"); ok {
		vnicEtherAttr.Virtualization_preference = Virtualization_preference.(string)
	}

	vnicEther := models.NewVnicEther(fmt.Sprintf("ether-%s", Name), VnicLanConnPolicy, desc, vnicEtherAttr)

	err := ucsClient.Save(vnicEther)
	if err != nil {
		return err
	}

	d.SetId(vnicEther.DistinguishedName)
	return resourceUcsVnicEtherRead(d, m)
}

func resourceUcsVnicEtherRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	vnicEther, err := getRemoteVnicEther(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setVnicEtherAttributes(vnicEther, d)

	return nil
}

func resourceUcsVnicEtherDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "vnicEther")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsVnicEtherUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	VnicLanConnPolicy := d.Get("vnic_lan_conn_policy_dn").(string)

	vnicEtherAttr := models.VnicEtherAttributes{}
	if Adaptor_profile_name, ok := d.GetOk("adaptor_profile_name"); ok {
		vnicEtherAttr.Adaptor_profile_name = Adaptor_profile_name.(string)
	}
	if Addr, ok := d.GetOk("addr"); ok {
		vnicEtherAttr.Addr = Addr.(string)
	}
	if Admin_cdn_name, ok := d.GetOk("admin_cdn_name"); ok {
		vnicEtherAttr.Admin_cdn_name = Admin_cdn_name.(string)
	}
	if Admin_host_port, ok := d.GetOk("admin_host_port"); ok {
		vnicEtherAttr.Admin_host_port = Admin_host_port.(string)
	}
	if Admin_vcon, ok := d.GetOk("admin_vcon"); ok {
		vnicEtherAttr.Admin_vcon = Admin_vcon.(string)
	}
	if Boot_dev, ok := d.GetOk("boot_dev"); ok {
		vnicEtherAttr.Boot_dev = Boot_dev.(string)
	}
	if Cdn_prop_in_sync, ok := d.GetOk("cdn_prop_in_sync"); ok {
		vnicEtherAttr.Cdn_prop_in_sync = Cdn_prop_in_sync.(string)
	}
	if Cdn_source, ok := d.GetOk("cdn_source"); ok {
		vnicEtherAttr.Cdn_source = Cdn_source.(string)
	}
	if Child_action, ok := d.GetOk("child_action"); ok {
		vnicEtherAttr.Child_action = Child_action.(string)
	}
	if Config_qualifier, ok := d.GetOk("config_qualifier"); ok {
		vnicEtherAttr.Config_qualifier = Config_qualifier.(string)
	}
	if Config_state, ok := d.GetOk("config_state"); ok {
		vnicEtherAttr.Config_state = Config_state.(string)
	}
	if Dynamic_id, ok := d.GetOk("dynamic_id"); ok {
		vnicEtherAttr.Dynamic_id = Dynamic_id.(string)
	}
	if Equipment_dn, ok := d.GetOk("equipment_dn"); ok {
		vnicEtherAttr.Equipment_dn = Equipment_dn.(string)
	}
	if Flt_aggr, ok := d.GetOk("flt_aggr"); ok {
		vnicEtherAttr.Flt_aggr = Flt_aggr.(string)
	}
	if Ident_pool_name, ok := d.GetOk("ident_pool_name"); ok {
		vnicEtherAttr.Ident_pool_name = Ident_pool_name.(string)
	}
	if Inst_type, ok := d.GetOk("inst_type"); ok {
		vnicEtherAttr.Inst_type = Inst_type.(string)
	}
	if Mtu, ok := d.GetOk("mtu"); ok {
		vnicEtherAttr.Mtu = Mtu.(string)
	}
	if Nw_ctrl_policy_name, ok := d.GetOk("nw_ctrl_policy_name"); ok {
		vnicEtherAttr.Nw_ctrl_policy_name = Nw_ctrl_policy_name.(string)
	}
	if Nw_templ_name, ok := d.GetOk("nw_templ_name"); ok {
		vnicEtherAttr.Nw_templ_name = Nw_templ_name.(string)
	}
	if Oper_adaptor_profile_name, ok := d.GetOk("oper_adaptor_profile_name"); ok {
		vnicEtherAttr.Oper_adaptor_profile_name = Oper_adaptor_profile_name.(string)
	}
	if Oper_cdn_name, ok := d.GetOk("oper_cdn_name"); ok {
		vnicEtherAttr.Oper_cdn_name = Oper_cdn_name.(string)
	}
	if Oper_host_port, ok := d.GetOk("oper_host_port"); ok {
		vnicEtherAttr.Oper_host_port = Oper_host_port.(string)
	}
	if Oper_ident_pool_name, ok := d.GetOk("oper_ident_pool_name"); ok {
		vnicEtherAttr.Oper_ident_pool_name = Oper_ident_pool_name.(string)
	}
	if Oper_nw_ctrl_policy_name, ok := d.GetOk("oper_nw_ctrl_policy_name"); ok {
		vnicEtherAttr.Oper_nw_ctrl_policy_name = Oper_nw_ctrl_policy_name.(string)
	}
	if Oper_nw_templ_name, ok := d.GetOk("oper_nw_templ_name"); ok {
		vnicEtherAttr.Oper_nw_templ_name = Oper_nw_templ_name.(string)
	}
	if Oper_order, ok := d.GetOk("oper_order"); ok {
		vnicEtherAttr.Oper_order = Oper_order.(string)
	}
	if Oper_pin_to_group_name, ok := d.GetOk("oper_pin_to_group_name"); ok {
		vnicEtherAttr.Oper_pin_to_group_name = Oper_pin_to_group_name.(string)
	}
	if Oper_qos_policy_name, ok := d.GetOk("oper_qos_policy_name"); ok {
		vnicEtherAttr.Oper_qos_policy_name = Oper_qos_policy_name.(string)
	}
	if Oper_speed, ok := d.GetOk("oper_speed"); ok {
		vnicEtherAttr.Oper_speed = Oper_speed.(string)
	}
	if Oper_stats_policy_name, ok := d.GetOk("oper_stats_policy_name"); ok {
		vnicEtherAttr.Oper_stats_policy_name = Oper_stats_policy_name.(string)
	}
	if Oper_vcon, ok := d.GetOk("oper_vcon"); ok {
		vnicEtherAttr.Oper_vcon = Oper_vcon.(string)
	}
	if Order, ok := d.GetOk("order"); ok {
		vnicEtherAttr.Order = Order.(string)
	}
	if Owner, ok := d.GetOk("owner"); ok {
		vnicEtherAttr.Owner = Owner.(string)
	}
	if Pf_dn, ok := d.GetOk("pf_dn"); ok {
		vnicEtherAttr.Pf_dn = Pf_dn.(string)
	}
	if Pin_to_group_name, ok := d.GetOk("pin_to_group_name"); ok {
		vnicEtherAttr.Pin_to_group_name = Pin_to_group_name.(string)
	}
	if Prop_acl, ok := d.GetOk("prop_acl"); ok {
		vnicEtherAttr.Prop_acl = Prop_acl.(string)
	}
	if Purpose, ok := d.GetOk("purpose"); ok {
		vnicEtherAttr.Purpose = Purpose.(string)
	}
	if Qos_policy_name, ok := d.GetOk("qos_policy_name"); ok {
		vnicEtherAttr.Qos_policy_name = Qos_policy_name.(string)
	}
	if Redundancy_pair_type, ok := d.GetOk("redundancy_pair_type"); ok {
		vnicEtherAttr.Redundancy_pair_type = Redundancy_pair_type.(string)
	}
	if Redundancy_peer, ok := d.GetOk("redundancy_peer"); ok {
		vnicEtherAttr.Redundancy_peer = Redundancy_peer.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		vnicEtherAttr.Sacl = Sacl.(string)
	}
	if Stats_policy_name, ok := d.GetOk("stats_policy_name"); ok {
		vnicEtherAttr.Stats_policy_name = Stats_policy_name.(string)
	}
	if Switch_id, ok := d.GetOk("switch_id"); ok {
		vnicEtherAttr.Switch_id = Switch_id.(string)
	}
	if Type, ok := d.GetOk("type"); ok {
		vnicEtherAttr.Type = Type.(string)
	}
	if Virtualization_preference, ok := d.GetOk("virtualization_preference"); ok {
		vnicEtherAttr.Virtualization_preference = Virtualization_preference.(string)
	}

	vnicEther := models.NewVnicEther(fmt.Sprintf("ether-%s", Name), VnicLanConnPolicy, desc, vnicEtherAttr)
	vnicEther.Status = "modified"
	err := ucsClient.Save(vnicEther)
	if err != nil {
		return err
	}

	d.SetId(vnicEther.DistinguishedName)
	return resourceUcsVnicEtherRead(d, m)
}
