package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsFabricVlan() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsFabricVlanCreate,
		Update: resourceUcsFabricVlanUpdate,
		Read:   resourceUcsFabricVlanRead,
		Delete: resourceUcsFabricVlanDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsFabricVlanImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"assoc_primary_vlan_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"assoc_primary_vlan_switch_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"child_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"cloud": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"compression_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"config_issues": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"config_overlap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"default_net": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"ep_dn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"flt_aggr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"r_global": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"fabric_vlan_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"if_role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"if_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"local": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"locale": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"mcast_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_mcast_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"overlap_state_for_a": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"overlap_state_for_b": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"peer_dn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"policy_owner": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"pub_nw_dn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"pub_nw_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"pub_nw_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"sacl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"sharing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"switch_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"transport": &schema.Schema{
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

func getRemoteFabricVlan(client *client.Client, dn string) (*models.FabricVlan, error) {
	fabricVlanDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	fabricVlan := models.FabricVlanFromDoc(fabricVlanDoc, "configResolveDn")

	if fabricVlan.DistinguishedName == "" {
		return nil, fmt.Errorf("FabricVlan %s not found", dn)
	}

	return fabricVlan, nil
}

func setFabricVlanAttributes(fabricVlan *models.FabricVlan, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(fabricVlan.DistinguishedName)
	d.Set("description", fabricVlan.Description)
	fabricVlanMap, _ := fabricVlan.ToMap()

	d.Set("assoc_primary_vlan_state", fabricVlanMap["assocPrimaryVlanState"])

	d.Set("assoc_primary_vlan_switch_id", fabricVlanMap["assocPrimaryVlanSwitchId"])

	d.Set("child_action", fabricVlanMap["childAction"])

	d.Set("cloud", fabricVlanMap["cloud"])

	d.Set("compression_type", fabricVlanMap["compressionType"])

	d.Set("config_issues", fabricVlanMap["configIssues"])

	d.Set("config_overlap", fabricVlanMap["configOverlap"])

	d.Set("default_net", fabricVlanMap["defaultNet"])

	d.Set("ep_dn", fabricVlanMap["epDn"])

	d.Set("flt_aggr", fabricVlanMap["fltAggr"])

	d.Set("r_global", fabricVlanMap["global"])

	d.Set("fabric_vlan_id", fabricVlanMap["id"])

	d.Set("if_role", fabricVlanMap["ifRole"])

	d.Set("if_type", fabricVlanMap["ifType"])

	d.Set("local", fabricVlanMap["local"])

	d.Set("locale", fabricVlanMap["locale"])

	d.Set("mcast_policy_name", fabricVlanMap["mcastPolicyName"])

	d.Set("oper_mcast_policy_name", fabricVlanMap["operMcastPolicyName"])

	d.Set("oper_state", fabricVlanMap["operState"])

	d.Set("overlap_state_for_a", fabricVlanMap["overlapStateForA"])

	d.Set("overlap_state_for_b", fabricVlanMap["overlapStateForB"])

	d.Set("peer_dn", fabricVlanMap["peerDn"])

	d.Set("policy_owner", fabricVlanMap["policyOwner"])

	d.Set("pub_nw_dn", fabricVlanMap["pubNwDn"])

	d.Set("pub_nw_id", fabricVlanMap["pubNwId"])

	d.Set("pub_nw_name", fabricVlanMap["pubNwName"])

	d.Set("sacl", fabricVlanMap["sacl"])

	d.Set("sharing", fabricVlanMap["sharing"])

	d.Set("switch_id", fabricVlanMap["switchId"])

	d.Set("transport", fabricVlanMap["transport"])

	d.Set("type", fabricVlanMap["type"])
	return d
}

func resourceUcsFabricVlanImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	fabricVlan, err := getRemoteFabricVlan(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setFabricVlanAttributes(fabricVlan, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsFabricVlanCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	fabricVlanAttr := models.FabricVlanAttributes{}

	if Assoc_primary_vlan_state, ok := d.GetOk("assoc_primary_vlan_state"); ok {
		fabricVlanAttr.Assoc_primary_vlan_state = Assoc_primary_vlan_state.(string)
	}

	if Assoc_primary_vlan_switch_id, ok := d.GetOk("assoc_primary_vlan_switch_id"); ok {
		fabricVlanAttr.Assoc_primary_vlan_switch_id = Assoc_primary_vlan_switch_id.(string)
	}

	if Child_action, ok := d.GetOk("child_action"); ok {
		fabricVlanAttr.Child_action = Child_action.(string)
	}

	if Cloud, ok := d.GetOk("cloud"); ok {
		fabricVlanAttr.Cloud = Cloud.(string)
	}

	if Compression_type, ok := d.GetOk("compression_type"); ok {
		fabricVlanAttr.Compression_type = Compression_type.(string)
	}

	if Config_issues, ok := d.GetOk("config_issues"); ok {
		fabricVlanAttr.Config_issues = Config_issues.(string)
	}

	if Config_overlap, ok := d.GetOk("config_overlap"); ok {
		fabricVlanAttr.Config_overlap = Config_overlap.(string)
	}

	if Default_net, ok := d.GetOk("default_net"); ok {
		fabricVlanAttr.Default_net = Default_net.(string)
	}

	if Ep_dn, ok := d.GetOk("ep_dn"); ok {
		fabricVlanAttr.Ep_dn = Ep_dn.(string)
	}

	if Flt_aggr, ok := d.GetOk("flt_aggr"); ok {
		fabricVlanAttr.Flt_aggr = Flt_aggr.(string)
	}

	if R_global, ok := d.GetOk("r_global"); ok {
		fabricVlanAttr.R_global = R_global.(string)
	}

	if Fabric_vlan_id, ok := d.GetOk("fabric_vlan_id"); ok {
		fabricVlanAttr.Fabric_vlan_id = Fabric_vlan_id.(string)
	}

	if If_role, ok := d.GetOk("if_role"); ok {
		fabricVlanAttr.If_role = If_role.(string)
	}

	if If_type, ok := d.GetOk("if_type"); ok {
		fabricVlanAttr.If_type = If_type.(string)
	}

	if Local, ok := d.GetOk("local"); ok {
		fabricVlanAttr.Local = Local.(string)
	}

	if Locale, ok := d.GetOk("locale"); ok {
		fabricVlanAttr.Locale = Locale.(string)
	}

	if Mcast_policy_name, ok := d.GetOk("mcast_policy_name"); ok {
		fabricVlanAttr.Mcast_policy_name = Mcast_policy_name.(string)
	}

	if Oper_mcast_policy_name, ok := d.GetOk("oper_mcast_policy_name"); ok {
		fabricVlanAttr.Oper_mcast_policy_name = Oper_mcast_policy_name.(string)
	}

	if Oper_state, ok := d.GetOk("oper_state"); ok {
		fabricVlanAttr.Oper_state = Oper_state.(string)
	}

	if Overlap_state_for_a, ok := d.GetOk("overlap_state_for_a"); ok {
		fabricVlanAttr.Overlap_state_for_a = Overlap_state_for_a.(string)
	}

	if Overlap_state_for_b, ok := d.GetOk("overlap_state_for_b"); ok {
		fabricVlanAttr.Overlap_state_for_b = Overlap_state_for_b.(string)
	}

	if Peer_dn, ok := d.GetOk("peer_dn"); ok {
		fabricVlanAttr.Peer_dn = Peer_dn.(string)
	}

	if Policy_owner, ok := d.GetOk("policy_owner"); ok {
		fabricVlanAttr.Policy_owner = Policy_owner.(string)
	}

	if Pub_nw_dn, ok := d.GetOk("pub_nw_dn"); ok {
		fabricVlanAttr.Pub_nw_dn = Pub_nw_dn.(string)
	}

	if Pub_nw_id, ok := d.GetOk("pub_nw_id"); ok {
		fabricVlanAttr.Pub_nw_id = Pub_nw_id.(string)
	}

	if Pub_nw_name, ok := d.GetOk("pub_nw_name"); ok {
		fabricVlanAttr.Pub_nw_name = Pub_nw_name.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		fabricVlanAttr.Sacl = Sacl.(string)
	}

	if Sharing, ok := d.GetOk("sharing"); ok {
		fabricVlanAttr.Sharing = Sharing.(string)
	}

	if Switch_id, ok := d.GetOk("switch_id"); ok {
		fabricVlanAttr.Switch_id = Switch_id.(string)
	}

	if Transport, ok := d.GetOk("transport"); ok {
		fabricVlanAttr.Transport = Transport.(string)
	}

	if Type, ok := d.GetOk("type"); ok {
		fabricVlanAttr.Type = Type.(string)
	}

	fabricVlan := models.NewFabricVlan(fmt.Sprintf("fabric/lan/net-%s", Name), desc, fabricVlanAttr)

	err := ucsClient.Save(fabricVlan)
	if err != nil {
		return err
	}

	d.SetId(fabricVlan.DistinguishedName)
	return resourceUcsFabricVlanRead(d, m)
}

func resourceUcsFabricVlanRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	fabricVlan, err := getRemoteFabricVlan(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setFabricVlanAttributes(fabricVlan, d)

	return nil
}

func resourceUcsFabricVlanDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "fabricVlan")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsFabricVlanUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	fabricVlanAttr := models.FabricVlanAttributes{}
	if Assoc_primary_vlan_state, ok := d.GetOk("assoc_primary_vlan_state"); ok {
		fabricVlanAttr.Assoc_primary_vlan_state = Assoc_primary_vlan_state.(string)
	}
	if Assoc_primary_vlan_switch_id, ok := d.GetOk("assoc_primary_vlan_switch_id"); ok {
		fabricVlanAttr.Assoc_primary_vlan_switch_id = Assoc_primary_vlan_switch_id.(string)
	}
	if Child_action, ok := d.GetOk("child_action"); ok {
		fabricVlanAttr.Child_action = Child_action.(string)
	}
	if Cloud, ok := d.GetOk("cloud"); ok {
		fabricVlanAttr.Cloud = Cloud.(string)
	}
	if Compression_type, ok := d.GetOk("compression_type"); ok {
		fabricVlanAttr.Compression_type = Compression_type.(string)
	}
	if Config_issues, ok := d.GetOk("config_issues"); ok {
		fabricVlanAttr.Config_issues = Config_issues.(string)
	}
	if Config_overlap, ok := d.GetOk("config_overlap"); ok {
		fabricVlanAttr.Config_overlap = Config_overlap.(string)
	}
	if Default_net, ok := d.GetOk("default_net"); ok {
		fabricVlanAttr.Default_net = Default_net.(string)
	}
	if Ep_dn, ok := d.GetOk("ep_dn"); ok {
		fabricVlanAttr.Ep_dn = Ep_dn.(string)
	}
	if Flt_aggr, ok := d.GetOk("flt_aggr"); ok {
		fabricVlanAttr.Flt_aggr = Flt_aggr.(string)
	}
	if R_global, ok := d.GetOk("r_global"); ok {
		fabricVlanAttr.R_global = R_global.(string)
	}
	if Fabric_vlan_id, ok := d.GetOk("fabric_vlan_id"); ok {
		fabricVlanAttr.Fabric_vlan_id = Fabric_vlan_id.(string)
	}
	if If_role, ok := d.GetOk("if_role"); ok {
		fabricVlanAttr.If_role = If_role.(string)
	}
	if If_type, ok := d.GetOk("if_type"); ok {
		fabricVlanAttr.If_type = If_type.(string)
	}
	if Local, ok := d.GetOk("local"); ok {
		fabricVlanAttr.Local = Local.(string)
	}
	if Locale, ok := d.GetOk("locale"); ok {
		fabricVlanAttr.Locale = Locale.(string)
	}
	if Mcast_policy_name, ok := d.GetOk("mcast_policy_name"); ok {
		fabricVlanAttr.Mcast_policy_name = Mcast_policy_name.(string)
	}
	if Oper_mcast_policy_name, ok := d.GetOk("oper_mcast_policy_name"); ok {
		fabricVlanAttr.Oper_mcast_policy_name = Oper_mcast_policy_name.(string)
	}
	if Oper_state, ok := d.GetOk("oper_state"); ok {
		fabricVlanAttr.Oper_state = Oper_state.(string)
	}
	if Overlap_state_for_a, ok := d.GetOk("overlap_state_for_a"); ok {
		fabricVlanAttr.Overlap_state_for_a = Overlap_state_for_a.(string)
	}
	if Overlap_state_for_b, ok := d.GetOk("overlap_state_for_b"); ok {
		fabricVlanAttr.Overlap_state_for_b = Overlap_state_for_b.(string)
	}
	if Peer_dn, ok := d.GetOk("peer_dn"); ok {
		fabricVlanAttr.Peer_dn = Peer_dn.(string)
	}
	if Policy_owner, ok := d.GetOk("policy_owner"); ok {
		fabricVlanAttr.Policy_owner = Policy_owner.(string)
	}
	if Pub_nw_dn, ok := d.GetOk("pub_nw_dn"); ok {
		fabricVlanAttr.Pub_nw_dn = Pub_nw_dn.(string)
	}
	if Pub_nw_id, ok := d.GetOk("pub_nw_id"); ok {
		fabricVlanAttr.Pub_nw_id = Pub_nw_id.(string)
	}
	if Pub_nw_name, ok := d.GetOk("pub_nw_name"); ok {
		fabricVlanAttr.Pub_nw_name = Pub_nw_name.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		fabricVlanAttr.Sacl = Sacl.(string)
	}
	if Sharing, ok := d.GetOk("sharing"); ok {
		fabricVlanAttr.Sharing = Sharing.(string)
	}
	if Switch_id, ok := d.GetOk("switch_id"); ok {
		fabricVlanAttr.Switch_id = Switch_id.(string)
	}
	if Transport, ok := d.GetOk("transport"); ok {
		fabricVlanAttr.Transport = Transport.(string)
	}
	if Type, ok := d.GetOk("type"); ok {
		fabricVlanAttr.Type = Type.(string)
	}

	fabricVlan := models.NewFabricVlan(fmt.Sprintf("fabric/lan/net-%s", Name), desc, fabricVlanAttr)
	fabricVlan.Status = "modified"
	err := ucsClient.Save(fabricVlan)
	if err != nil {
		return err
	}

	d.SetId(fabricVlan.DistinguishedName)
	return resourceUcsFabricVlanRead(d, m)
}
