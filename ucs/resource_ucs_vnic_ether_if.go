package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsVnicEtherIf() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsVnicEtherIfCreate,
		Update: resourceUcsVnicEtherIfUpdate,
		Read:   resourceUcsVnicEtherIfRead,
		Delete: resourceUcsVnicEtherIfDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsVnicEtherIfImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"vnic_lan_conn_templ_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"addr": &schema.Schema{
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

			"default_net": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"flt_aggr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_primary_vnet_dn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_primary_vnet_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_vnet_dn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_vnet_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"owner": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"prop_acl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"pub_nw_id": &schema.Schema{
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

			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"vnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		}),
	}
}

func getRemoteVnicEtherIf(client *client.Client, dn string) (*models.VnicEtherIf, error) {
	vnicEtherIfDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	vnicEtherIf := models.VnicEtherIfFromDoc(vnicEtherIfDoc, "configResolveDn")

	if vnicEtherIf.DistinguishedName == "" {
		return nil, fmt.Errorf("VnicEtherIf %s not found", dn)
	}

	return vnicEtherIf, nil
}

func setVnicEtherIfAttributes(vnicEtherIf *models.VnicEtherIf, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(vnicEtherIf.DistinguishedName)
	d.Set("description", vnicEtherIf.Description)
	d.Set("vnic_lan_conn_templ_dn", GetParentDn(vnicEtherIf.DistinguishedName))
	vnicEtherIfMap, _ := vnicEtherIf.ToMap()

	d.Set("addr", vnicEtherIfMap["addr"])

	d.Set("child_action", vnicEtherIfMap["childAction"])

	d.Set("config_qualifier", vnicEtherIfMap["configQualifier"])

	d.Set("default_net", vnicEtherIfMap["defaultNet"])

	d.Set("flt_aggr", vnicEtherIfMap["fltAggr"])

	d.Set("oper_primary_vnet_dn", vnicEtherIfMap["operPrimaryVnetDn"])

	d.Set("oper_primary_vnet_name", vnicEtherIfMap["operPrimaryVnetName"])

	d.Set("oper_state", vnicEtherIfMap["operState"])

	d.Set("oper_vnet_dn", vnicEtherIfMap["operVnetDn"])

	d.Set("oper_vnet_name", vnicEtherIfMap["operVnetName"])

	d.Set("owner", vnicEtherIfMap["owner"])

	d.Set("prop_acl", vnicEtherIfMap["propAcl"])

	d.Set("pub_nw_id", vnicEtherIfMap["pubNwId"])

	d.Set("sacl", vnicEtherIfMap["sacl"])

	d.Set("sharing", vnicEtherIfMap["sharing"])

	d.Set("switch_id", vnicEtherIfMap["switchId"])

	d.Set("type", vnicEtherIfMap["type"])

	d.Set("vnet", vnicEtherIfMap["vnet"])
	return d
}

func resourceUcsVnicEtherIfImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	vnicEtherIf, err := getRemoteVnicEtherIf(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setVnicEtherIfAttributes(vnicEtherIf, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsVnicEtherIfCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	VnicLanConnTempl := d.Get("vnic_lan_conn_templ_dn").(string)

	vnicEtherIfAttr := models.VnicEtherIfAttributes{}

	if Addr, ok := d.GetOk("addr"); ok {
		vnicEtherIfAttr.Addr = Addr.(string)
	}

	if Child_action, ok := d.GetOk("child_action"); ok {
		vnicEtherIfAttr.Child_action = Child_action.(string)
	}

	if Config_qualifier, ok := d.GetOk("config_qualifier"); ok {
		vnicEtherIfAttr.Config_qualifier = Config_qualifier.(string)
	}

	if Default_net, ok := d.GetOk("default_net"); ok {
		vnicEtherIfAttr.Default_net = Default_net.(string)
	}

	if Flt_aggr, ok := d.GetOk("flt_aggr"); ok {
		vnicEtherIfAttr.Flt_aggr = Flt_aggr.(string)
	}

	if Oper_primary_vnet_dn, ok := d.GetOk("oper_primary_vnet_dn"); ok {
		vnicEtherIfAttr.Oper_primary_vnet_dn = Oper_primary_vnet_dn.(string)
	}

	if Oper_primary_vnet_name, ok := d.GetOk("oper_primary_vnet_name"); ok {
		vnicEtherIfAttr.Oper_primary_vnet_name = Oper_primary_vnet_name.(string)
	}

	if Oper_state, ok := d.GetOk("oper_state"); ok {
		vnicEtherIfAttr.Oper_state = Oper_state.(string)
	}

	if Oper_vnet_dn, ok := d.GetOk("oper_vnet_dn"); ok {
		vnicEtherIfAttr.Oper_vnet_dn = Oper_vnet_dn.(string)
	}

	if Oper_vnet_name, ok := d.GetOk("oper_vnet_name"); ok {
		vnicEtherIfAttr.Oper_vnet_name = Oper_vnet_name.(string)
	}

	if Owner, ok := d.GetOk("owner"); ok {
		vnicEtherIfAttr.Owner = Owner.(string)
	}

	if Prop_acl, ok := d.GetOk("prop_acl"); ok {
		vnicEtherIfAttr.Prop_acl = Prop_acl.(string)
	}

	if Pub_nw_id, ok := d.GetOk("pub_nw_id"); ok {
		vnicEtherIfAttr.Pub_nw_id = Pub_nw_id.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		vnicEtherIfAttr.Sacl = Sacl.(string)
	}

	if Sharing, ok := d.GetOk("sharing"); ok {
		vnicEtherIfAttr.Sharing = Sharing.(string)
	}

	if Switch_id, ok := d.GetOk("switch_id"); ok {
		vnicEtherIfAttr.Switch_id = Switch_id.(string)
	}

	if Type, ok := d.GetOk("type"); ok {
		vnicEtherIfAttr.Type = Type.(string)
	}

	if Vnet, ok := d.GetOk("vnet"); ok {
		vnicEtherIfAttr.Vnet = Vnet.(string)
	}

	vnicEtherIf := models.NewVnicEtherIf(fmt.Sprintf("if-%s", Name), VnicLanConnTempl, desc, vnicEtherIfAttr)

	err := ucsClient.Save(vnicEtherIf)
	if err != nil {
		return err
	}

	d.SetId(vnicEtherIf.DistinguishedName)
	return resourceUcsVnicEtherIfRead(d, m)
}

func resourceUcsVnicEtherIfRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	vnicEtherIf, err := getRemoteVnicEtherIf(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setVnicEtherIfAttributes(vnicEtherIf, d)

	return nil
}

func resourceUcsVnicEtherIfDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "vnicEtherIf")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsVnicEtherIfUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	VnicLanConnTempl := d.Get("vnic_lan_conn_templ_dn").(string)

	vnicEtherIfAttr := models.VnicEtherIfAttributes{}
	if Addr, ok := d.GetOk("addr"); ok {
		vnicEtherIfAttr.Addr = Addr.(string)
	}
	if Child_action, ok := d.GetOk("child_action"); ok {
		vnicEtherIfAttr.Child_action = Child_action.(string)
	}
	if Config_qualifier, ok := d.GetOk("config_qualifier"); ok {
		vnicEtherIfAttr.Config_qualifier = Config_qualifier.(string)
	}
	if Default_net, ok := d.GetOk("default_net"); ok {
		vnicEtherIfAttr.Default_net = Default_net.(string)
	}
	if Flt_aggr, ok := d.GetOk("flt_aggr"); ok {
		vnicEtherIfAttr.Flt_aggr = Flt_aggr.(string)
	}
	if Oper_primary_vnet_dn, ok := d.GetOk("oper_primary_vnet_dn"); ok {
		vnicEtherIfAttr.Oper_primary_vnet_dn = Oper_primary_vnet_dn.(string)
	}
	if Oper_primary_vnet_name, ok := d.GetOk("oper_primary_vnet_name"); ok {
		vnicEtherIfAttr.Oper_primary_vnet_name = Oper_primary_vnet_name.(string)
	}
	if Oper_state, ok := d.GetOk("oper_state"); ok {
		vnicEtherIfAttr.Oper_state = Oper_state.(string)
	}
	if Oper_vnet_dn, ok := d.GetOk("oper_vnet_dn"); ok {
		vnicEtherIfAttr.Oper_vnet_dn = Oper_vnet_dn.(string)
	}
	if Oper_vnet_name, ok := d.GetOk("oper_vnet_name"); ok {
		vnicEtherIfAttr.Oper_vnet_name = Oper_vnet_name.(string)
	}
	if Owner, ok := d.GetOk("owner"); ok {
		vnicEtherIfAttr.Owner = Owner.(string)
	}
	if Prop_acl, ok := d.GetOk("prop_acl"); ok {
		vnicEtherIfAttr.Prop_acl = Prop_acl.(string)
	}
	if Pub_nw_id, ok := d.GetOk("pub_nw_id"); ok {
		vnicEtherIfAttr.Pub_nw_id = Pub_nw_id.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		vnicEtherIfAttr.Sacl = Sacl.(string)
	}
	if Sharing, ok := d.GetOk("sharing"); ok {
		vnicEtherIfAttr.Sharing = Sharing.(string)
	}
	if Switch_id, ok := d.GetOk("switch_id"); ok {
		vnicEtherIfAttr.Switch_id = Switch_id.(string)
	}
	if Type, ok := d.GetOk("type"); ok {
		vnicEtherIfAttr.Type = Type.(string)
	}
	if Vnet, ok := d.GetOk("vnet"); ok {
		vnicEtherIfAttr.Vnet = Vnet.(string)
	}

	vnicEtherIf := models.NewVnicEtherIf(fmt.Sprintf("if-%s", Name), VnicLanConnTempl, desc, vnicEtherIfAttr)
	vnicEtherIf.Status = "modified"
	err := ucsClient.Save(vnicEtherIf)
	if err != nil {
		return err
	}

	d.SetId(vnicEtherIf.DistinguishedName)
	return resourceUcsVnicEtherIfRead(d, m)
}
