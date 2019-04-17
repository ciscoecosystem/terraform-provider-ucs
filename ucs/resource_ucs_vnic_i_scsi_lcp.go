package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsVnicIScsiLCP() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsVnicIScsiLCPCreate,
		Update: resourceUcsVnicIScsiLCPUpdate,
		Read:   resourceUcsVnicIScsiLCPRead,
		Delete: resourceUcsVnicIScsiLCPDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsVnicIScsiLCPImport,
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

			"oper_order": &schema.Schema{
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

			"vnic_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		}),
	}
}

func getRemoteVnicIScsiLCP(client *client.Client, dn string) (*models.VnicIScsiLCP, error) {
	vnicIScsiLCPDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	vnicIScsiLCP := models.VnicIScsiLCPFromDoc(vnicIScsiLCPDoc, "configResolveDn")

	if vnicIScsiLCP.DistinguishedName == "" {
		return nil, fmt.Errorf("VnicIScsiLCP %s not found", dn)
	}

	return vnicIScsiLCP, nil
}

func setVnicIScsiLCPAttributes(vnicIScsiLCP *models.VnicIScsiLCP, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(vnicIScsiLCP.DistinguishedName)
	d.Set("description", vnicIScsiLCP.Description)
	d.Set("vnic_lan_conn_policy_dn", GetParentDn(vnicIScsiLCP.DistinguishedName))
	vnicIScsiLCPMap, _ := vnicIScsiLCP.ToMap()

	d.Set("adaptor_profile_name", vnicIScsiLCPMap["adaptorProfileName"])

	d.Set("addr", vnicIScsiLCPMap["addr"])

	d.Set("admin_cdn_name", vnicIScsiLCPMap["adminCdnName"])

	d.Set("admin_host_port", vnicIScsiLCPMap["adminHostPort"])

	d.Set("admin_vcon", vnicIScsiLCPMap["adminVcon"])

	d.Set("boot_dev", vnicIScsiLCPMap["bootDev"])

	d.Set("cdn_prop_in_sync", vnicIScsiLCPMap["cdnPropInSync"])

	d.Set("cdn_source", vnicIScsiLCPMap["cdnSource"])

	d.Set("child_action", vnicIScsiLCPMap["childAction"])

	d.Set("config_qualifier", vnicIScsiLCPMap["configQualifier"])

	d.Set("config_state", vnicIScsiLCPMap["configState"])

	d.Set("equipment_dn", vnicIScsiLCPMap["equipmentDn"])

	d.Set("flt_aggr", vnicIScsiLCPMap["fltAggr"])

	d.Set("ident_pool_name", vnicIScsiLCPMap["identPoolName"])

	d.Set("inst_type", vnicIScsiLCPMap["instType"])

	d.Set("nw_templ_name", vnicIScsiLCPMap["nwTemplName"])

	d.Set("oper_adaptor_profile_name", vnicIScsiLCPMap["operAdaptorProfileName"])

	d.Set("oper_cdn_name", vnicIScsiLCPMap["operCdnName"])

	d.Set("oper_host_port", vnicIScsiLCPMap["operHostPort"])

	d.Set("oper_ident_pool_name", vnicIScsiLCPMap["operIdentPoolName"])

	d.Set("oper_order", vnicIScsiLCPMap["operOrder"])

	d.Set("oper_speed", vnicIScsiLCPMap["operSpeed"])

	d.Set("oper_stats_policy_name", vnicIScsiLCPMap["operStatsPolicyName"])

	d.Set("oper_vcon", vnicIScsiLCPMap["operVcon"])

	d.Set("order", vnicIScsiLCPMap["order"])

	d.Set("owner", vnicIScsiLCPMap["owner"])

	d.Set("pin_to_group_name", vnicIScsiLCPMap["pinToGroupName"])

	d.Set("qos_policy_name", vnicIScsiLCPMap["qosPolicyName"])

	d.Set("redundancy_pair_type", vnicIScsiLCPMap["redundancyPairType"])

	d.Set("redundancy_peer", vnicIScsiLCPMap["redundancyPeer"])

	d.Set("sacl", vnicIScsiLCPMap["sacl"])

	d.Set("stats_policy_name", vnicIScsiLCPMap["statsPolicyName"])

	d.Set("switch_id", vnicIScsiLCPMap["switchId"])

	d.Set("type", vnicIScsiLCPMap["type"])

	d.Set("vnic_name", vnicIScsiLCPMap["vnicName"])
	return d
}

func resourceUcsVnicIScsiLCPImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	vnicIScsiLCP, err := getRemoteVnicIScsiLCP(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setVnicIScsiLCPAttributes(vnicIScsiLCP, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsVnicIScsiLCPCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	VnicLanConnPolicy := d.Get("vnic_lan_conn_policy_dn").(string)

	vnicIScsiLCPAttr := models.VnicIScsiLCPAttributes{}

	if Adaptor_profile_name, ok := d.GetOk("adaptor_profile_name"); ok {
		vnicIScsiLCPAttr.Adaptor_profile_name = Adaptor_profile_name.(string)
	}

	if Addr, ok := d.GetOk("addr"); ok {
		vnicIScsiLCPAttr.Addr = Addr.(string)
	}

	if Admin_cdn_name, ok := d.GetOk("admin_cdn_name"); ok {
		vnicIScsiLCPAttr.Admin_cdn_name = Admin_cdn_name.(string)
	}

	if Admin_host_port, ok := d.GetOk("admin_host_port"); ok {
		vnicIScsiLCPAttr.Admin_host_port = Admin_host_port.(string)
	}

	if Admin_vcon, ok := d.GetOk("admin_vcon"); ok {
		vnicIScsiLCPAttr.Admin_vcon = Admin_vcon.(string)
	}

	if Boot_dev, ok := d.GetOk("boot_dev"); ok {
		vnicIScsiLCPAttr.Boot_dev = Boot_dev.(string)
	}

	if Cdn_prop_in_sync, ok := d.GetOk("cdn_prop_in_sync"); ok {
		vnicIScsiLCPAttr.Cdn_prop_in_sync = Cdn_prop_in_sync.(string)
	}

	if Cdn_source, ok := d.GetOk("cdn_source"); ok {
		vnicIScsiLCPAttr.Cdn_source = Cdn_source.(string)
	}

	if Child_action, ok := d.GetOk("child_action"); ok {
		vnicIScsiLCPAttr.Child_action = Child_action.(string)
	}

	if Config_qualifier, ok := d.GetOk("config_qualifier"); ok {
		vnicIScsiLCPAttr.Config_qualifier = Config_qualifier.(string)
	}

	if Config_state, ok := d.GetOk("config_state"); ok {
		vnicIScsiLCPAttr.Config_state = Config_state.(string)
	}

	if Equipment_dn, ok := d.GetOk("equipment_dn"); ok {
		vnicIScsiLCPAttr.Equipment_dn = Equipment_dn.(string)
	}

	if Flt_aggr, ok := d.GetOk("flt_aggr"); ok {
		vnicIScsiLCPAttr.Flt_aggr = Flt_aggr.(string)
	}

	if Ident_pool_name, ok := d.GetOk("ident_pool_name"); ok {
		vnicIScsiLCPAttr.Ident_pool_name = Ident_pool_name.(string)
	}

	if Inst_type, ok := d.GetOk("inst_type"); ok {
		vnicIScsiLCPAttr.Inst_type = Inst_type.(string)
	}

	if Nw_templ_name, ok := d.GetOk("nw_templ_name"); ok {
		vnicIScsiLCPAttr.Nw_templ_name = Nw_templ_name.(string)
	}

	if Oper_adaptor_profile_name, ok := d.GetOk("oper_adaptor_profile_name"); ok {
		vnicIScsiLCPAttr.Oper_adaptor_profile_name = Oper_adaptor_profile_name.(string)
	}

	if Oper_cdn_name, ok := d.GetOk("oper_cdn_name"); ok {
		vnicIScsiLCPAttr.Oper_cdn_name = Oper_cdn_name.(string)
	}

	if Oper_host_port, ok := d.GetOk("oper_host_port"); ok {
		vnicIScsiLCPAttr.Oper_host_port = Oper_host_port.(string)
	}

	if Oper_ident_pool_name, ok := d.GetOk("oper_ident_pool_name"); ok {
		vnicIScsiLCPAttr.Oper_ident_pool_name = Oper_ident_pool_name.(string)
	}

	if Oper_order, ok := d.GetOk("oper_order"); ok {
		vnicIScsiLCPAttr.Oper_order = Oper_order.(string)
	}

	if Oper_speed, ok := d.GetOk("oper_speed"); ok {
		vnicIScsiLCPAttr.Oper_speed = Oper_speed.(string)
	}

	if Oper_stats_policy_name, ok := d.GetOk("oper_stats_policy_name"); ok {
		vnicIScsiLCPAttr.Oper_stats_policy_name = Oper_stats_policy_name.(string)
	}

	if Oper_vcon, ok := d.GetOk("oper_vcon"); ok {
		vnicIScsiLCPAttr.Oper_vcon = Oper_vcon.(string)
	}

	if Order, ok := d.GetOk("order"); ok {
		vnicIScsiLCPAttr.Order = Order.(string)
	}

	if Owner, ok := d.GetOk("owner"); ok {
		vnicIScsiLCPAttr.Owner = Owner.(string)
	}

	if Pin_to_group_name, ok := d.GetOk("pin_to_group_name"); ok {
		vnicIScsiLCPAttr.Pin_to_group_name = Pin_to_group_name.(string)
	}

	if Qos_policy_name, ok := d.GetOk("qos_policy_name"); ok {
		vnicIScsiLCPAttr.Qos_policy_name = Qos_policy_name.(string)
	}

	if Redundancy_pair_type, ok := d.GetOk("redundancy_pair_type"); ok {
		vnicIScsiLCPAttr.Redundancy_pair_type = Redundancy_pair_type.(string)
	}

	if Redundancy_peer, ok := d.GetOk("redundancy_peer"); ok {
		vnicIScsiLCPAttr.Redundancy_peer = Redundancy_peer.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		vnicIScsiLCPAttr.Sacl = Sacl.(string)
	}

	if Stats_policy_name, ok := d.GetOk("stats_policy_name"); ok {
		vnicIScsiLCPAttr.Stats_policy_name = Stats_policy_name.(string)
	}

	if Switch_id, ok := d.GetOk("switch_id"); ok {
		vnicIScsiLCPAttr.Switch_id = Switch_id.(string)
	}

	if Type, ok := d.GetOk("type"); ok {
		vnicIScsiLCPAttr.Type = Type.(string)
	}

	if Vnic_name, ok := d.GetOk("vnic_name"); ok {
		vnicIScsiLCPAttr.Vnic_name = Vnic_name.(string)
	}

	vnicIScsiLCP := models.NewVnicIScsiLCP(fmt.Sprintf("iscsi-%s", Name), VnicLanConnPolicy, desc, vnicIScsiLCPAttr)

	err := ucsClient.Save(vnicIScsiLCP)
	if err != nil {
		return err
	}

	d.SetId(vnicIScsiLCP.DistinguishedName)
	return resourceUcsVnicIScsiLCPRead(d, m)
}

func resourceUcsVnicIScsiLCPRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	vnicIScsiLCP, err := getRemoteVnicIScsiLCP(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setVnicIScsiLCPAttributes(vnicIScsiLCP, d)

	return nil
}

func resourceUcsVnicIScsiLCPDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "vnicIScsiLCP")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsVnicIScsiLCPUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	VnicLanConnPolicy := d.Get("vnic_lan_conn_policy_dn").(string)

	vnicIScsiLCPAttr := models.VnicIScsiLCPAttributes{}
	if Adaptor_profile_name, ok := d.GetOk("adaptor_profile_name"); ok {
		vnicIScsiLCPAttr.Adaptor_profile_name = Adaptor_profile_name.(string)
	}
	if Addr, ok := d.GetOk("addr"); ok {
		vnicIScsiLCPAttr.Addr = Addr.(string)
	}
	if Admin_cdn_name, ok := d.GetOk("admin_cdn_name"); ok {
		vnicIScsiLCPAttr.Admin_cdn_name = Admin_cdn_name.(string)
	}
	if Admin_host_port, ok := d.GetOk("admin_host_port"); ok {
		vnicIScsiLCPAttr.Admin_host_port = Admin_host_port.(string)
	}
	if Admin_vcon, ok := d.GetOk("admin_vcon"); ok {
		vnicIScsiLCPAttr.Admin_vcon = Admin_vcon.(string)
	}
	if Boot_dev, ok := d.GetOk("boot_dev"); ok {
		vnicIScsiLCPAttr.Boot_dev = Boot_dev.(string)
	}
	if Cdn_prop_in_sync, ok := d.GetOk("cdn_prop_in_sync"); ok {
		vnicIScsiLCPAttr.Cdn_prop_in_sync = Cdn_prop_in_sync.(string)
	}
	if Cdn_source, ok := d.GetOk("cdn_source"); ok {
		vnicIScsiLCPAttr.Cdn_source = Cdn_source.(string)
	}
	if Child_action, ok := d.GetOk("child_action"); ok {
		vnicIScsiLCPAttr.Child_action = Child_action.(string)
	}
	if Config_qualifier, ok := d.GetOk("config_qualifier"); ok {
		vnicIScsiLCPAttr.Config_qualifier = Config_qualifier.(string)
	}
	if Config_state, ok := d.GetOk("config_state"); ok {
		vnicIScsiLCPAttr.Config_state = Config_state.(string)
	}
	if Equipment_dn, ok := d.GetOk("equipment_dn"); ok {
		vnicIScsiLCPAttr.Equipment_dn = Equipment_dn.(string)
	}
	if Flt_aggr, ok := d.GetOk("flt_aggr"); ok {
		vnicIScsiLCPAttr.Flt_aggr = Flt_aggr.(string)
	}
	if Ident_pool_name, ok := d.GetOk("ident_pool_name"); ok {
		vnicIScsiLCPAttr.Ident_pool_name = Ident_pool_name.(string)
	}
	if Inst_type, ok := d.GetOk("inst_type"); ok {
		vnicIScsiLCPAttr.Inst_type = Inst_type.(string)
	}
	if Nw_templ_name, ok := d.GetOk("nw_templ_name"); ok {
		vnicIScsiLCPAttr.Nw_templ_name = Nw_templ_name.(string)
	}
	if Oper_adaptor_profile_name, ok := d.GetOk("oper_adaptor_profile_name"); ok {
		vnicIScsiLCPAttr.Oper_adaptor_profile_name = Oper_adaptor_profile_name.(string)
	}
	if Oper_cdn_name, ok := d.GetOk("oper_cdn_name"); ok {
		vnicIScsiLCPAttr.Oper_cdn_name = Oper_cdn_name.(string)
	}
	if Oper_host_port, ok := d.GetOk("oper_host_port"); ok {
		vnicIScsiLCPAttr.Oper_host_port = Oper_host_port.(string)
	}
	if Oper_ident_pool_name, ok := d.GetOk("oper_ident_pool_name"); ok {
		vnicIScsiLCPAttr.Oper_ident_pool_name = Oper_ident_pool_name.(string)
	}
	if Oper_order, ok := d.GetOk("oper_order"); ok {
		vnicIScsiLCPAttr.Oper_order = Oper_order.(string)
	}
	if Oper_speed, ok := d.GetOk("oper_speed"); ok {
		vnicIScsiLCPAttr.Oper_speed = Oper_speed.(string)
	}
	if Oper_stats_policy_name, ok := d.GetOk("oper_stats_policy_name"); ok {
		vnicIScsiLCPAttr.Oper_stats_policy_name = Oper_stats_policy_name.(string)
	}
	if Oper_vcon, ok := d.GetOk("oper_vcon"); ok {
		vnicIScsiLCPAttr.Oper_vcon = Oper_vcon.(string)
	}
	if Order, ok := d.GetOk("order"); ok {
		vnicIScsiLCPAttr.Order = Order.(string)
	}
	if Owner, ok := d.GetOk("owner"); ok {
		vnicIScsiLCPAttr.Owner = Owner.(string)
	}
	if Pin_to_group_name, ok := d.GetOk("pin_to_group_name"); ok {
		vnicIScsiLCPAttr.Pin_to_group_name = Pin_to_group_name.(string)
	}
	if Qos_policy_name, ok := d.GetOk("qos_policy_name"); ok {
		vnicIScsiLCPAttr.Qos_policy_name = Qos_policy_name.(string)
	}
	if Redundancy_pair_type, ok := d.GetOk("redundancy_pair_type"); ok {
		vnicIScsiLCPAttr.Redundancy_pair_type = Redundancy_pair_type.(string)
	}
	if Redundancy_peer, ok := d.GetOk("redundancy_peer"); ok {
		vnicIScsiLCPAttr.Redundancy_peer = Redundancy_peer.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		vnicIScsiLCPAttr.Sacl = Sacl.(string)
	}
	if Stats_policy_name, ok := d.GetOk("stats_policy_name"); ok {
		vnicIScsiLCPAttr.Stats_policy_name = Stats_policy_name.(string)
	}
	if Switch_id, ok := d.GetOk("switch_id"); ok {
		vnicIScsiLCPAttr.Switch_id = Switch_id.(string)
	}
	if Type, ok := d.GetOk("type"); ok {
		vnicIScsiLCPAttr.Type = Type.(string)
	}
	if Vnic_name, ok := d.GetOk("vnic_name"); ok {
		vnicIScsiLCPAttr.Vnic_name = Vnic_name.(string)
	}

	vnicIScsiLCP := models.NewVnicIScsiLCP(fmt.Sprintf("iscsi-%s", Name), VnicLanConnPolicy, desc, vnicIScsiLCPAttr)
	vnicIScsiLCP.Status = "modified"
	err := ucsClient.Save(vnicIScsiLCP)
	if err != nil {
		return err
	}

	d.SetId(vnicIScsiLCP.DistinguishedName)
	return resourceUcsVnicIScsiLCPRead(d, m)
}
