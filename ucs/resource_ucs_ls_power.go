package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsLsPower() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsLsPowerCreate,
		Update: resourceUcsLsPowerUpdate,
		Read:   resourceUcsLsPowerRead,
		Delete: resourceUcsLsPowerDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsLsPowerImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"ls_server_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"child_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"prop_acl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"sacl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"soft_shutdown_timer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		}),
	}
}

func getRemoteLsPower(client *client.Client, dn string) (*models.LsPower, error) {
	lsPowerDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	lsPower := models.LsPowerFromDoc(lsPowerDoc, "configResolveDn")

	if lsPower.DistinguishedName == "" {
		return nil, fmt.Errorf("LsPower %s not found", dn)
	}

	return lsPower, nil
}

func setLsPowerAttributes(lsPower *models.LsPower, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(lsPower.DistinguishedName)
	d.Set("description", lsPower.Description)
	d.Set("ls_server_dn", GetParentDn(lsPower.DistinguishedName))
	lsPowerMap, _ := lsPower.ToMap()

	d.Set("child_action", lsPowerMap["childAction"])

	d.Set("prop_acl", lsPowerMap["propAcl"])

	d.Set("sacl", lsPowerMap["sacl"])

	d.Set("soft_shutdown_timer", lsPowerMap["softShutdownTimer"])

	d.Set("state", lsPowerMap["state"])
	return d
}

func resourceUcsLsPowerImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	lsPower, err := getRemoteLsPower(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setLsPowerAttributes(lsPower, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsLsPowerCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	LsServer := d.Get("ls_server_dn").(string)

	lsPowerAttr := models.LsPowerAttributes{}

	if Child_action, ok := d.GetOk("child_action"); ok {
		lsPowerAttr.Child_action = Child_action.(string)
	}

	if Prop_acl, ok := d.GetOk("prop_acl"); ok {
		lsPowerAttr.Prop_acl = Prop_acl.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		lsPowerAttr.Sacl = Sacl.(string)
	}

	if Soft_shutdown_timer, ok := d.GetOk("soft_shutdown_timer"); ok {
		lsPowerAttr.Soft_shutdown_timer = Soft_shutdown_timer.(string)
	}

	if State, ok := d.GetOk("state"); ok {
		lsPowerAttr.State = State.(string)
	}

	lsPower := models.NewLsPower(fmt.Sprintf("power"), LsServer, desc, lsPowerAttr)

	err := ucsClient.Save(lsPower)
	if err != nil {
		return err
	}

	d.SetId(lsPower.DistinguishedName)
	return resourceUcsLsPowerRead(d, m)
}

func resourceUcsLsPowerRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	lsPower, err := getRemoteLsPower(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setLsPowerAttributes(lsPower, d)

	return nil
}

func resourceUcsLsPowerDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "lsPower")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsLsPowerUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	LsServer := d.Get("ls_server_dn").(string)

	lsPowerAttr := models.LsPowerAttributes{}
	if Child_action, ok := d.GetOk("child_action"); ok {
		lsPowerAttr.Child_action = Child_action.(string)
	}
	if Prop_acl, ok := d.GetOk("prop_acl"); ok {
		lsPowerAttr.Prop_acl = Prop_acl.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		lsPowerAttr.Sacl = Sacl.(string)
	}
	if Soft_shutdown_timer, ok := d.GetOk("soft_shutdown_timer"); ok {
		lsPowerAttr.Soft_shutdown_timer = Soft_shutdown_timer.(string)
	}
	if State, ok := d.GetOk("state"); ok {
		lsPowerAttr.State = State.(string)
	}

	lsPower := models.NewLsPower(fmt.Sprintf("power"), LsServer, desc, lsPowerAttr)
	lsPower.Status = "modified"
	err := ucsClient.Save(lsPower)
	if err != nil {
		return err
	}

	d.SetId(lsPower.DistinguishedName)
	return resourceUcsLsPowerRead(d, m)
}
