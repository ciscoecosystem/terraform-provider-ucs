package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsVnicVlan() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsVnicVlanCreate,
		Update: resourceUcsVnicVlanUpdate,
		Read:   resourceUcsVnicVlanRead,
		Delete: resourceUcsVnicVlanDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsVnicVlanImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"vnic_i_scsi_lcp_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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

			"vlan_name": &schema.Schema{
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

func getRemoteVnicVlan(client *client.Client, dn string) (*models.VnicVlan, error) {
	vnicVlanDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	vnicVlan := models.VnicVlanFromDoc(vnicVlanDoc, "configResolveDn")

	if vnicVlan.DistinguishedName == "" {
		return nil, fmt.Errorf("VnicVlan %s not found", dn)
	}

	return vnicVlan, nil
}

func setVnicVlanAttributes(vnicVlan *models.VnicVlan, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(vnicVlan.DistinguishedName)
	d.Set("description", vnicVlan.Description)
	d.Set("vnic_i_scsi_lcp_dn", GetParentDn(vnicVlan.DistinguishedName))
	vnicVlanMap, _ := vnicVlan.ToMap()

	d.Set("child_action", vnicVlanMap["childAction"])

	d.Set("config_qualifier", vnicVlanMap["configQualifier"])

	d.Set("flt_aggr", vnicVlanMap["fltAggr"])

	d.Set("name", vnicVlanMap["name"])

	d.Set("oper_primary_vnet_dn", vnicVlanMap["operPrimaryVnetDn"])

	d.Set("oper_primary_vnet_name", vnicVlanMap["operPrimaryVnetName"])

	d.Set("oper_state", vnicVlanMap["operState"])

	d.Set("oper_vnet_dn", vnicVlanMap["operVnetDn"])

	d.Set("oper_vnet_name", vnicVlanMap["operVnetName"])

	d.Set("owner", vnicVlanMap["owner"])

	d.Set("pub_nw_id", vnicVlanMap["pubNwId"])

	d.Set("sacl", vnicVlanMap["sacl"])

	d.Set("sharing", vnicVlanMap["sharing"])

	d.Set("switch_id", vnicVlanMap["switchId"])

	d.Set("type", vnicVlanMap["type"])

	d.Set("vlan_name", vnicVlanMap["vlanName"])

	d.Set("vnet", vnicVlanMap["vnet"])
	return d
}

func resourceUcsVnicVlanImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	vnicVlan, err := getRemoteVnicVlan(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setVnicVlanAttributes(vnicVlan, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsVnicVlanCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	VnicIScsiLCP := d.Get("vnic_i_scsi_lcp_dn").(string)

	vnicVlanAttr := models.VnicVlanAttributes{}

	if Child_action, ok := d.GetOk("child_action"); ok {
		vnicVlanAttr.Child_action = Child_action.(string)
	}

	if Config_qualifier, ok := d.GetOk("config_qualifier"); ok {
		vnicVlanAttr.Config_qualifier = Config_qualifier.(string)
	}

	if Flt_aggr, ok := d.GetOk("flt_aggr"); ok {
		vnicVlanAttr.Flt_aggr = Flt_aggr.(string)
	}

	if Oper_primary_vnet_dn, ok := d.GetOk("oper_primary_vnet_dn"); ok {
		vnicVlanAttr.Oper_primary_vnet_dn = Oper_primary_vnet_dn.(string)
	}

	if Oper_primary_vnet_name, ok := d.GetOk("oper_primary_vnet_name"); ok {
		vnicVlanAttr.Oper_primary_vnet_name = Oper_primary_vnet_name.(string)
	}

	if Oper_state, ok := d.GetOk("oper_state"); ok {
		vnicVlanAttr.Oper_state = Oper_state.(string)
	}

	if Oper_vnet_dn, ok := d.GetOk("oper_vnet_dn"); ok {
		vnicVlanAttr.Oper_vnet_dn = Oper_vnet_dn.(string)
	}

	if Oper_vnet_name, ok := d.GetOk("oper_vnet_name"); ok {
		vnicVlanAttr.Oper_vnet_name = Oper_vnet_name.(string)
	}

	if Owner, ok := d.GetOk("owner"); ok {
		vnicVlanAttr.Owner = Owner.(string)
	}

	if Pub_nw_id, ok := d.GetOk("pub_nw_id"); ok {
		vnicVlanAttr.Pub_nw_id = Pub_nw_id.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		vnicVlanAttr.Sacl = Sacl.(string)
	}

	if Sharing, ok := d.GetOk("sharing"); ok {
		vnicVlanAttr.Sharing = Sharing.(string)
	}

	if Switch_id, ok := d.GetOk("switch_id"); ok {
		vnicVlanAttr.Switch_id = Switch_id.(string)
	}

	if Type, ok := d.GetOk("type"); ok {
		vnicVlanAttr.Type = Type.(string)
	}

	if Vlan_name, ok := d.GetOk("vlan_name"); ok {
		vnicVlanAttr.Vlan_name = Vlan_name.(string)
	}

	if Vnet, ok := d.GetOk("vnet"); ok {
		vnicVlanAttr.Vnet = Vnet.(string)
	}

	vnicVlan := models.NewVnicVlan(fmt.Sprintf("vlan"), VnicIScsiLCP, desc, vnicVlanAttr)

	err := ucsClient.Save(vnicVlan)
	if err != nil {
		return err
	}

	d.SetId(vnicVlan.DistinguishedName)
	return resourceUcsVnicVlanRead(d, m)
}

func resourceUcsVnicVlanRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	vnicVlan, err := getRemoteVnicVlan(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setVnicVlanAttributes(vnicVlan, d)

	return nil
}

func resourceUcsVnicVlanDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "vnicVlan")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsVnicVlanUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	VnicIScsiLCP := d.Get("vnic_i_scsi_lcp_dn").(string)

	vnicVlanAttr := models.VnicVlanAttributes{}
	if Child_action, ok := d.GetOk("child_action"); ok {
		vnicVlanAttr.Child_action = Child_action.(string)
	}
	if Config_qualifier, ok := d.GetOk("config_qualifier"); ok {
		vnicVlanAttr.Config_qualifier = Config_qualifier.(string)
	}
	if Flt_aggr, ok := d.GetOk("flt_aggr"); ok {
		vnicVlanAttr.Flt_aggr = Flt_aggr.(string)
	}
	if Oper_primary_vnet_dn, ok := d.GetOk("oper_primary_vnet_dn"); ok {
		vnicVlanAttr.Oper_primary_vnet_dn = Oper_primary_vnet_dn.(string)
	}
	if Oper_primary_vnet_name, ok := d.GetOk("oper_primary_vnet_name"); ok {
		vnicVlanAttr.Oper_primary_vnet_name = Oper_primary_vnet_name.(string)
	}
	if Oper_state, ok := d.GetOk("oper_state"); ok {
		vnicVlanAttr.Oper_state = Oper_state.(string)
	}
	if Oper_vnet_dn, ok := d.GetOk("oper_vnet_dn"); ok {
		vnicVlanAttr.Oper_vnet_dn = Oper_vnet_dn.(string)
	}
	if Oper_vnet_name, ok := d.GetOk("oper_vnet_name"); ok {
		vnicVlanAttr.Oper_vnet_name = Oper_vnet_name.(string)
	}
	if Owner, ok := d.GetOk("owner"); ok {
		vnicVlanAttr.Owner = Owner.(string)
	}
	if Pub_nw_id, ok := d.GetOk("pub_nw_id"); ok {
		vnicVlanAttr.Pub_nw_id = Pub_nw_id.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		vnicVlanAttr.Sacl = Sacl.(string)
	}
	if Sharing, ok := d.GetOk("sharing"); ok {
		vnicVlanAttr.Sharing = Sharing.(string)
	}
	if Switch_id, ok := d.GetOk("switch_id"); ok {
		vnicVlanAttr.Switch_id = Switch_id.(string)
	}
	if Type, ok := d.GetOk("type"); ok {
		vnicVlanAttr.Type = Type.(string)
	}
	if Vlan_name, ok := d.GetOk("vlan_name"); ok {
		vnicVlanAttr.Vlan_name = Vlan_name.(string)
	}
	if Vnet, ok := d.GetOk("vnet"); ok {
		vnicVlanAttr.Vnet = Vnet.(string)
	}

	vnicVlan := models.NewVnicVlan(fmt.Sprintf("vlan"), VnicIScsiLCP, desc, vnicVlanAttr)
	vnicVlan.Status = "modified"
	err := ucsClient.Save(vnicVlan)
	if err != nil {
		return err
	}

	d.SetId(vnicVlan.DistinguishedName)
	return resourceUcsVnicVlanRead(d, m)
}
