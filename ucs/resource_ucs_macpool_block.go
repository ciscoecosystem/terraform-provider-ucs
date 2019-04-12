package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsMacpoolBlock() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsMacpoolBlockCreate,
		Update: resourceUcsMacpoolBlockUpdate,
		Read:   resourceUcsMacpoolBlockRead,
		Delete: resourceUcsMacpoolBlockDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsMacpoolBlockImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"macpool_pool_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"r_from": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"to": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"child_action": &schema.Schema{
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

func getRemoteMacpoolBlock(client *client.Client, dn string) (*models.MacpoolBlock, error) {
	macpoolBlockDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	macpoolBlock := models.MacpoolBlockFromDoc(macpoolBlockDoc, "configResolveDn")

	if macpoolBlock.DistinguishedName == "" {
		return nil, fmt.Errorf("MacpoolBlock %s not found", dn)
	}

	return macpoolBlock, nil
}

func setMacpoolBlockAttributes(macpoolBlock *models.MacpoolBlock, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(macpoolBlock.DistinguishedName)
	d.Set("description", macpoolBlock.Description)
	d.Set("macpool_pool_dn", GetParentDn(macpoolBlock.DistinguishedName))
	macpoolBlockMap, _ := macpoolBlock.ToMap()

	d.Set("child_action", macpoolBlockMap["childAction"])

	d.Set("sacl", macpoolBlockMap["sacl"])
	return d
}

func resourceUcsMacpoolBlockImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	macpoolBlock, err := getRemoteMacpoolBlock(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setMacpoolBlockAttributes(macpoolBlock, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsMacpoolBlockCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	R_from := d.Get("r_from").(string)
	To := d.Get("to").(string)

	MacpoolPool := d.Get("macpool_pool_dn").(string)

	macpoolBlockAttr := models.MacpoolBlockAttributes{}

	if Child_action, ok := d.GetOk("child_action"); ok {
		macpoolBlockAttr.Child_action = Child_action.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		macpoolBlockAttr.Sacl = Sacl.(string)
	}

	macpoolBlock := models.NewMacpoolBlock(fmt.Sprintf("block-%s-%s", R_from, To), MacpoolPool, desc, macpoolBlockAttr)

	err := ucsClient.Save(macpoolBlock)
	if err != nil {
		return err
	}

	d.SetId(macpoolBlock.DistinguishedName)
	return resourceUcsMacpoolBlockRead(d, m)
}

func resourceUcsMacpoolBlockRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	macpoolBlock, err := getRemoteMacpoolBlock(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setMacpoolBlockAttributes(macpoolBlock, d)

	return nil
}

func resourceUcsMacpoolBlockDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "macpoolBlock")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsMacpoolBlockUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	R_from := d.Get("r_from").(string)
	To := d.Get("to").(string)

	MacpoolPool := d.Get("macpool_pool_dn").(string)

	macpoolBlockAttr := models.MacpoolBlockAttributes{}
	if Child_action, ok := d.GetOk("child_action"); ok {
		macpoolBlockAttr.Child_action = Child_action.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		macpoolBlockAttr.Sacl = Sacl.(string)
	}

	macpoolBlock := models.NewMacpoolBlock(fmt.Sprintf("block-%s-%s", R_from, To), MacpoolPool, desc, macpoolBlockAttr)
	macpoolBlock.Status = "modified"
	err := ucsClient.Save(macpoolBlock)
	if err != nil {
		return err
	}

	d.SetId(macpoolBlock.DistinguishedName)
	return resourceUcsMacpoolBlockRead(d, m)
}
