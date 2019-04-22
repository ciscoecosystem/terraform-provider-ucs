package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsMgmtVnet() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsMgmtVnetCreate,
		Update: resourceUcsMgmtVnetUpdate,
		Read:   resourceUcsMgmtVnetRead,
		Delete: resourceUcsMgmtVnetDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsMgmtVnetImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"mgmt_interface_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"child_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"config_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"mgmt_vnet_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"sacl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		}),
	}
}

func getRemoteMgmtVnet(client *client.Client, dn string) (*models.MgmtVnet, error) {
	mgmtVnetDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	mgmtVnet := models.MgmtVnetFromDoc(mgmtVnetDoc, "configResolveDn")

	if mgmtVnet.DistinguishedName == "" {
		return nil, fmt.Errorf("MgmtVnet %s not found", dn)
	}

	return mgmtVnet, nil
}

func setMgmtVnetAttributes(mgmtVnet *models.MgmtVnet, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(mgmtVnet.DistinguishedName)
	d.Set("description", mgmtVnet.Description)
	d.Set("mgmt_interface_dn", GetParentDn(mgmtVnet.DistinguishedName))
	mgmtVnetMap, _ := mgmtVnet.ToMap()

	d.Set("child_action", mgmtVnetMap["childAction"])

	d.Set("config_state", mgmtVnetMap["configState"])

	d.Set("mgmt_vnet_id", mgmtVnetMap["id"])

	d.Set("name", mgmtVnetMap["name"])

	d.Set("sacl", mgmtVnetMap["sacl"])
	return d
}

func resourceUcsMgmtVnetImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	mgmtVnet, err := getRemoteMgmtVnet(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setMgmtVnetAttributes(mgmtVnet, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsMgmtVnetCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	MgmtInterface := d.Get("mgmt_interface_dn").(string)

	mgmtVnetAttr := models.MgmtVnetAttributes{}

	if Child_action, ok := d.GetOk("child_action"); ok {
		mgmtVnetAttr.Child_action = Child_action.(string)
	}

	if Config_state, ok := d.GetOk("config_state"); ok {
		mgmtVnetAttr.Config_state = Config_state.(string)
	}

	if Mgmt_vnet_id, ok := d.GetOk("mgmt_vnet_id"); ok {
		mgmtVnetAttr.Mgmt_vnet_id = Mgmt_vnet_id.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		mgmtVnetAttr.Sacl = Sacl.(string)
	}

	mgmtVnet := models.NewMgmtVnet(fmt.Sprintf("network"), MgmtInterface, desc, mgmtVnetAttr)

	err := ucsClient.Save(mgmtVnet)
	if err != nil {
		return err
	}

	d.SetId(mgmtVnet.DistinguishedName)
	return resourceUcsMgmtVnetRead(d, m)
}

func resourceUcsMgmtVnetRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	mgmtVnet, err := getRemoteMgmtVnet(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setMgmtVnetAttributes(mgmtVnet, d)

	return nil
}

func resourceUcsMgmtVnetDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "mgmtVnet")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsMgmtVnetUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	MgmtInterface := d.Get("mgmt_interface_dn").(string)

	mgmtVnetAttr := models.MgmtVnetAttributes{}
	if Child_action, ok := d.GetOk("child_action"); ok {
		mgmtVnetAttr.Child_action = Child_action.(string)
	}
	if Config_state, ok := d.GetOk("config_state"); ok {
		mgmtVnetAttr.Config_state = Config_state.(string)
	}
	if Mgmt_vnet_id, ok := d.GetOk("mgmt_vnet_id"); ok {
		mgmtVnetAttr.Mgmt_vnet_id = Mgmt_vnet_id.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		mgmtVnetAttr.Sacl = Sacl.(string)
	}

	mgmtVnet := models.NewMgmtVnet(fmt.Sprintf("network"), MgmtInterface, desc, mgmtVnetAttr)
	mgmtVnet.Status = "modified"
	err := ucsClient.Save(mgmtVnet)
	if err != nil {
		return err
	}

	d.SetId(mgmtVnet.DistinguishedName)
	return resourceUcsMgmtVnetRead(d, m)
}
