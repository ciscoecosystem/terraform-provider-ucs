package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsFabricVsan() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsFabricVsanCreate,
		Update: resourceUcsFabricVsanUpdate,
		Read:   resourceUcsFabricVsanRead,
		Delete: resourceUcsFabricVsanDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsFabricVsanImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"child_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"config_overlap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"default_zoning": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"ep_dn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"fc_zone_sharing_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"fcoe_vlan": &schema.Schema{
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

			"fabric_vsan_id": &schema.Schema{
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

			"oper_state": &schema.Schema{
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

			"sacl": &schema.Schema{
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

			"zoning_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		}),
	}
}

func getRemoteFabricVsan(client *client.Client, dn string) (*models.FabricVsan, error) {
	fabricVsanDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	fabricVsan := models.FabricVsanFromDoc(fabricVsanDoc, "configResolveDn")

	if fabricVsan.DistinguishedName == "" {
		return nil, fmt.Errorf("FabricVsan %s not found", dn)
	}

	return fabricVsan, nil
}

func setFabricVsanAttributes(fabricVsan *models.FabricVsan, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(fabricVsan.DistinguishedName)
	d.Set("description", fabricVsan.Description)
	fabricVsanMap, _ := fabricVsan.ToMap()

	d.Set("child_action", fabricVsanMap["childAction"])

	d.Set("config_overlap", fabricVsanMap["configOverlap"])

	d.Set("default_zoning", fabricVsanMap["defaultZoning"])

	d.Set("ep_dn", fabricVsanMap["epDn"])

	d.Set("fc_zone_sharing_mode", fabricVsanMap["fcZoneSharingMode"])

	d.Set("fcoe_vlan", fabricVsanMap["fcoeVlan"])

	d.Set("flt_aggr", fabricVsanMap["fltAggr"])

	d.Set("r_global", fabricVsanMap["global"])

	d.Set("fabric_vsan_id", fabricVsanMap["id"])

	d.Set("if_role", fabricVsanMap["ifRole"])

	d.Set("if_type", fabricVsanMap["ifType"])

	d.Set("local", fabricVsanMap["local"])

	d.Set("locale", fabricVsanMap["locale"])

	d.Set("oper_state", fabricVsanMap["operState"])

	d.Set("peer_dn", fabricVsanMap["peerDn"])

	d.Set("policy_owner", fabricVsanMap["policyOwner"])

	d.Set("sacl", fabricVsanMap["sacl"])

	d.Set("switch_id", fabricVsanMap["switchId"])

	d.Set("transport", fabricVsanMap["transport"])

	d.Set("type", fabricVsanMap["type"])

	d.Set("zoning_state", fabricVsanMap["zoningState"])
	return d
}

func resourceUcsFabricVsanImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	fabricVsan, err := getRemoteFabricVsan(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setFabricVsanAttributes(fabricVsan, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsFabricVsanCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	fabricVsanAttr := models.FabricVsanAttributes{}

	if Child_action, ok := d.GetOk("child_action"); ok {
		fabricVsanAttr.Child_action = Child_action.(string)
	}

	if Config_overlap, ok := d.GetOk("config_overlap"); ok {
		fabricVsanAttr.Config_overlap = Config_overlap.(string)
	}

	if Default_zoning, ok := d.GetOk("default_zoning"); ok {
		fabricVsanAttr.Default_zoning = Default_zoning.(string)
	}

	if Ep_dn, ok := d.GetOk("ep_dn"); ok {
		fabricVsanAttr.Ep_dn = Ep_dn.(string)
	}

	if Fc_zone_sharing_mode, ok := d.GetOk("fc_zone_sharing_mode"); ok {
		fabricVsanAttr.Fc_zone_sharing_mode = Fc_zone_sharing_mode.(string)
	}

	if Fcoe_vlan, ok := d.GetOk("fcoe_vlan"); ok {
		fabricVsanAttr.Fcoe_vlan = Fcoe_vlan.(string)
	}

	if Flt_aggr, ok := d.GetOk("flt_aggr"); ok {
		fabricVsanAttr.Flt_aggr = Flt_aggr.(string)
	}

	if R_global, ok := d.GetOk("r_global"); ok {
		fabricVsanAttr.R_global = R_global.(string)
	}

	if Fabric_vsan_id, ok := d.GetOk("fabric_vsan_id"); ok {
		fabricVsanAttr.Fabric_vsan_id = Fabric_vsan_id.(string)
	}

	if If_role, ok := d.GetOk("if_role"); ok {
		fabricVsanAttr.If_role = If_role.(string)
	}

	if If_type, ok := d.GetOk("if_type"); ok {
		fabricVsanAttr.If_type = If_type.(string)
	}

	if Local, ok := d.GetOk("local"); ok {
		fabricVsanAttr.Local = Local.(string)
	}

	if Locale, ok := d.GetOk("locale"); ok {
		fabricVsanAttr.Locale = Locale.(string)
	}

	if Oper_state, ok := d.GetOk("oper_state"); ok {
		fabricVsanAttr.Oper_state = Oper_state.(string)
	}

	if Peer_dn, ok := d.GetOk("peer_dn"); ok {
		fabricVsanAttr.Peer_dn = Peer_dn.(string)
	}

	if Policy_owner, ok := d.GetOk("policy_owner"); ok {
		fabricVsanAttr.Policy_owner = Policy_owner.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		fabricVsanAttr.Sacl = Sacl.(string)
	}

	if Switch_id, ok := d.GetOk("switch_id"); ok {
		fabricVsanAttr.Switch_id = Switch_id.(string)
	}

	if Transport, ok := d.GetOk("transport"); ok {
		fabricVsanAttr.Transport = Transport.(string)
	}

	if Type, ok := d.GetOk("type"); ok {
		fabricVsanAttr.Type = Type.(string)
	}

	if Zoning_state, ok := d.GetOk("zoning_state"); ok {
		fabricVsanAttr.Zoning_state = Zoning_state.(string)
	}

	fabricVsan := models.NewFabricVsan(fmt.Sprintf("fabric/san/net-%s", Name), desc, fabricVsanAttr)

	err := ucsClient.Save(fabricVsan)
	if err != nil {
		return err
	}

	d.SetId(fabricVsan.DistinguishedName)
	return resourceUcsFabricVsanRead(d, m)
}

func resourceUcsFabricVsanRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	fabricVsan, err := getRemoteFabricVsan(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setFabricVsanAttributes(fabricVsan, d)

	return nil
}

func resourceUcsFabricVsanDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "fabricVsan")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsFabricVsanUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	fabricVsanAttr := models.FabricVsanAttributes{}
	if Child_action, ok := d.GetOk("child_action"); ok {
		fabricVsanAttr.Child_action = Child_action.(string)
	}
	if Config_overlap, ok := d.GetOk("config_overlap"); ok {
		fabricVsanAttr.Config_overlap = Config_overlap.(string)
	}
	if Default_zoning, ok := d.GetOk("default_zoning"); ok {
		fabricVsanAttr.Default_zoning = Default_zoning.(string)
	}
	if Ep_dn, ok := d.GetOk("ep_dn"); ok {
		fabricVsanAttr.Ep_dn = Ep_dn.(string)
	}
	if Fc_zone_sharing_mode, ok := d.GetOk("fc_zone_sharing_mode"); ok {
		fabricVsanAttr.Fc_zone_sharing_mode = Fc_zone_sharing_mode.(string)
	}
	if Fcoe_vlan, ok := d.GetOk("fcoe_vlan"); ok {
		fabricVsanAttr.Fcoe_vlan = Fcoe_vlan.(string)
	}
	if Flt_aggr, ok := d.GetOk("flt_aggr"); ok {
		fabricVsanAttr.Flt_aggr = Flt_aggr.(string)
	}
	if R_global, ok := d.GetOk("r_global"); ok {
		fabricVsanAttr.R_global = R_global.(string)
	}
	if Fabric_vsan_id, ok := d.GetOk("fabric_vsan_id"); ok {
		fabricVsanAttr.Fabric_vsan_id = Fabric_vsan_id.(string)
	}
	if If_role, ok := d.GetOk("if_role"); ok {
		fabricVsanAttr.If_role = If_role.(string)
	}
	if If_type, ok := d.GetOk("if_type"); ok {
		fabricVsanAttr.If_type = If_type.(string)
	}
	if Local, ok := d.GetOk("local"); ok {
		fabricVsanAttr.Local = Local.(string)
	}
	if Locale, ok := d.GetOk("locale"); ok {
		fabricVsanAttr.Locale = Locale.(string)
	}
	if Oper_state, ok := d.GetOk("oper_state"); ok {
		fabricVsanAttr.Oper_state = Oper_state.(string)
	}
	if Peer_dn, ok := d.GetOk("peer_dn"); ok {
		fabricVsanAttr.Peer_dn = Peer_dn.(string)
	}
	if Policy_owner, ok := d.GetOk("policy_owner"); ok {
		fabricVsanAttr.Policy_owner = Policy_owner.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		fabricVsanAttr.Sacl = Sacl.(string)
	}
	if Switch_id, ok := d.GetOk("switch_id"); ok {
		fabricVsanAttr.Switch_id = Switch_id.(string)
	}
	if Transport, ok := d.GetOk("transport"); ok {
		fabricVsanAttr.Transport = Transport.(string)
	}
	if Type, ok := d.GetOk("type"); ok {
		fabricVsanAttr.Type = Type.(string)
	}
	if Zoning_state, ok := d.GetOk("zoning_state"); ok {
		fabricVsanAttr.Zoning_state = Zoning_state.(string)
	}

	fabricVsan := models.NewFabricVsan(fmt.Sprintf("fabric/san/net-%s", Name), desc, fabricVsanAttr)
	fabricVsan.Status = "modified"
	err := ucsClient.Save(fabricVsan)
	if err != nil {
		return err
	}

	d.SetId(fabricVsan.DistinguishedName)
	return resourceUcsFabricVsanRead(d, m)
}
