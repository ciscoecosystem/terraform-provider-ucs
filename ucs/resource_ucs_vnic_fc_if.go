package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsVnicFcIf() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsVnicFcIfCreate,
		Update: resourceUcsVnicFcIfUpdate,
		Read:   resourceUcsVnicFcIfRead,
		Delete: resourceUcsVnicFcIfDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsVnicFcIfImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"vnic_fc_dn": &schema.Schema{
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

			"initiator": &schema.Schema{
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

			"vnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		}),
	}
}

func getRemoteVnicFcIf(client *client.Client, dn string) (*models.VnicFcIf, error) {
	vnicFcIfDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	vnicFcIf := models.VnicFcIfFromDoc(vnicFcIfDoc, "configResolveDn")

	if vnicFcIf.DistinguishedName == "" {
		return nil, fmt.Errorf("VnicFcIf %s not found", dn)
	}

	return vnicFcIf, nil
}

func setVnicFcIfAttributes(vnicFcIf *models.VnicFcIf, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(vnicFcIf.DistinguishedName)
	d.Set("description", vnicFcIf.Description)
	d.Set("vnic_fc_dn", GetParentDn(vnicFcIf.DistinguishedName))
	vnicFcIfMap, _ := vnicFcIf.ToMap()

	d.Set("child_action", vnicFcIfMap["childAction"])

	d.Set("config_qualifier", vnicFcIfMap["configQualifier"])

	d.Set("initiator", vnicFcIfMap["initiator"])

	d.Set("name", vnicFcIfMap["name"])

	d.Set("oper_primary_vnet_dn", vnicFcIfMap["operPrimaryVnetDn"])

	d.Set("oper_primary_vnet_name", vnicFcIfMap["operPrimaryVnetName"])

	d.Set("oper_state", vnicFcIfMap["operState"])

	d.Set("oper_vnet_dn", vnicFcIfMap["operVnetDn"])

	d.Set("oper_vnet_name", vnicFcIfMap["operVnetName"])

	d.Set("owner", vnicFcIfMap["owner"])

	d.Set("pub_nw_id", vnicFcIfMap["pubNwId"])

	d.Set("sacl", vnicFcIfMap["sacl"])

	d.Set("sharing", vnicFcIfMap["sharing"])

	d.Set("switch_id", vnicFcIfMap["switchId"])

	d.Set("type", vnicFcIfMap["type"])

	d.Set("vnet", vnicFcIfMap["vnet"])
	return d
}

func resourceUcsVnicFcIfImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	vnicFcIf, err := getRemoteVnicFcIf(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setVnicFcIfAttributes(vnicFcIf, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsVnicFcIfCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	VnicFc := d.Get("vnic_fc_dn").(string)

	vnicFcIfAttr := models.VnicFcIfAttributes{}

	if Child_action, ok := d.GetOk("child_action"); ok {
		vnicFcIfAttr.Child_action = Child_action.(string)
	}

	if Config_qualifier, ok := d.GetOk("config_qualifier"); ok {
		vnicFcIfAttr.Config_qualifier = Config_qualifier.(string)
	}

	if Initiator, ok := d.GetOk("initiator"); ok {
		vnicFcIfAttr.Initiator = Initiator.(string)
	}

	if Oper_primary_vnet_dn, ok := d.GetOk("oper_primary_vnet_dn"); ok {
		vnicFcIfAttr.Oper_primary_vnet_dn = Oper_primary_vnet_dn.(string)
	}

	if Oper_primary_vnet_name, ok := d.GetOk("oper_primary_vnet_name"); ok {
		vnicFcIfAttr.Oper_primary_vnet_name = Oper_primary_vnet_name.(string)
	}

	if Oper_state, ok := d.GetOk("oper_state"); ok {
		vnicFcIfAttr.Oper_state = Oper_state.(string)
	}

	if Oper_vnet_dn, ok := d.GetOk("oper_vnet_dn"); ok {
		vnicFcIfAttr.Oper_vnet_dn = Oper_vnet_dn.(string)
	}

	if Oper_vnet_name, ok := d.GetOk("oper_vnet_name"); ok {
		vnicFcIfAttr.Oper_vnet_name = Oper_vnet_name.(string)
	}

	if Owner, ok := d.GetOk("owner"); ok {
		vnicFcIfAttr.Owner = Owner.(string)
	}

	if Pub_nw_id, ok := d.GetOk("pub_nw_id"); ok {
		vnicFcIfAttr.Pub_nw_id = Pub_nw_id.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		vnicFcIfAttr.Sacl = Sacl.(string)
	}

	if Sharing, ok := d.GetOk("sharing"); ok {
		vnicFcIfAttr.Sharing = Sharing.(string)
	}

	if Switch_id, ok := d.GetOk("switch_id"); ok {
		vnicFcIfAttr.Switch_id = Switch_id.(string)
	}

	if Type, ok := d.GetOk("type"); ok {
		vnicFcIfAttr.Type = Type.(string)
	}

	if Vnet, ok := d.GetOk("vnet"); ok {
		vnicFcIfAttr.Vnet = Vnet.(string)
	}

	vnicFcIf := models.NewVnicFcIf(fmt.Sprintf("if-default"), VnicFc, desc, vnicFcIfAttr)

	err := ucsClient.Save(vnicFcIf)
	if err != nil {
		return err
	}

	d.SetId(vnicFcIf.DistinguishedName)
	return resourceUcsVnicFcIfRead(d, m)
}

func resourceUcsVnicFcIfRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	vnicFcIf, err := getRemoteVnicFcIf(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setVnicFcIfAttributes(vnicFcIf, d)

	return nil
}

func resourceUcsVnicFcIfDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "vnicFcIf")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsVnicFcIfUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	VnicFc := d.Get("vnic_fc_dn").(string)

	vnicFcIfAttr := models.VnicFcIfAttributes{}
	if Child_action, ok := d.GetOk("child_action"); ok {
		vnicFcIfAttr.Child_action = Child_action.(string)
	}
	if Config_qualifier, ok := d.GetOk("config_qualifier"); ok {
		vnicFcIfAttr.Config_qualifier = Config_qualifier.(string)
	}
	if Initiator, ok := d.GetOk("initiator"); ok {
		vnicFcIfAttr.Initiator = Initiator.(string)
	}
	if Oper_primary_vnet_dn, ok := d.GetOk("oper_primary_vnet_dn"); ok {
		vnicFcIfAttr.Oper_primary_vnet_dn = Oper_primary_vnet_dn.(string)
	}
	if Oper_primary_vnet_name, ok := d.GetOk("oper_primary_vnet_name"); ok {
		vnicFcIfAttr.Oper_primary_vnet_name = Oper_primary_vnet_name.(string)
	}
	if Oper_state, ok := d.GetOk("oper_state"); ok {
		vnicFcIfAttr.Oper_state = Oper_state.(string)
	}
	if Oper_vnet_dn, ok := d.GetOk("oper_vnet_dn"); ok {
		vnicFcIfAttr.Oper_vnet_dn = Oper_vnet_dn.(string)
	}
	if Oper_vnet_name, ok := d.GetOk("oper_vnet_name"); ok {
		vnicFcIfAttr.Oper_vnet_name = Oper_vnet_name.(string)
	}
	if Owner, ok := d.GetOk("owner"); ok {
		vnicFcIfAttr.Owner = Owner.(string)
	}
	if Pub_nw_id, ok := d.GetOk("pub_nw_id"); ok {
		vnicFcIfAttr.Pub_nw_id = Pub_nw_id.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		vnicFcIfAttr.Sacl = Sacl.(string)
	}
	if Sharing, ok := d.GetOk("sharing"); ok {
		vnicFcIfAttr.Sharing = Sharing.(string)
	}
	if Switch_id, ok := d.GetOk("switch_id"); ok {
		vnicFcIfAttr.Switch_id = Switch_id.(string)
	}
	if Type, ok := d.GetOk("type"); ok {
		vnicFcIfAttr.Type = Type.(string)
	}
	if Vnet, ok := d.GetOk("vnet"); ok {
		vnicFcIfAttr.Vnet = Vnet.(string)
	}

	vnicFcIf := models.NewVnicFcIf(fmt.Sprintf("if-default"), VnicFc, desc, vnicFcIfAttr)
	vnicFcIf.Status = "modified"
	err := ucsClient.Save(vnicFcIf)
	if err != nil {
		return err
	}

	d.SetId(vnicFcIf.DistinguishedName)
	return resourceUcsVnicFcIfRead(d, m)
}
