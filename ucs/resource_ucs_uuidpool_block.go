package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsUuidpoolBlock() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsUuidpoolBlockCreate,
		Update: resourceUcsUuidpoolBlockUpdate,
		Read:   resourceUcsUuidpoolBlockRead,
		Delete: resourceUcsUuidpoolBlockDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsUuidpoolBlockImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"uuidpool_pool_dn": &schema.Schema{
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

func getRemoteUuidpoolBlock(client *client.Client, dn string) (*models.UuidpoolBlock, error) {
	uuidpoolBlockDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	uuidpoolBlock := models.UuidpoolBlockFromDoc(uuidpoolBlockDoc, "configResolveDn")

	if uuidpoolBlock.DistinguishedName == "" {
		return nil, fmt.Errorf("UuidpoolBlock %s not found", dn)
	}

	return uuidpoolBlock, nil
}

func setUuidpoolBlockAttributes(uuidpoolBlock *models.UuidpoolBlock, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(uuidpoolBlock.DistinguishedName)
	d.Set("description", uuidpoolBlock.Description)
	d.Set("uuidpool_pool_dn", GetParentDn(uuidpoolBlock.DistinguishedName))
	uuidpoolBlockMap, _ := uuidpoolBlock.ToMap()

	d.Set("child_action", uuidpoolBlockMap["childAction"])

	d.Set("sacl", uuidpoolBlockMap["sacl"])
	return d
}

func resourceUcsUuidpoolBlockImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	uuidpoolBlock, err := getRemoteUuidpoolBlock(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setUuidpoolBlockAttributes(uuidpoolBlock, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsUuidpoolBlockCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	R_from := d.Get("r_from").(string)
	To := d.Get("to").(string)

	UuidpoolPool := d.Get("uuidpool_pool_dn").(string)

	uuidpoolBlockAttr := models.UuidpoolBlockAttributes{}

	if Child_action, ok := d.GetOk("child_action"); ok {
		uuidpoolBlockAttr.Child_action = Child_action.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		uuidpoolBlockAttr.Sacl = Sacl.(string)
	}

	uuidpoolBlock := models.NewUuidpoolBlock(fmt.Sprintf("block-from-%s-to-%s", R_from, To), UuidpoolPool, desc, uuidpoolBlockAttr)

	err := ucsClient.Save(uuidpoolBlock)
	if err != nil {
		return err
	}

	d.SetId(uuidpoolBlock.DistinguishedName)
	return resourceUcsUuidpoolBlockRead(d, m)
}

func resourceUcsUuidpoolBlockRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	uuidpoolBlock, err := getRemoteUuidpoolBlock(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setUuidpoolBlockAttributes(uuidpoolBlock, d)

	return nil
}

func resourceUcsUuidpoolBlockDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "uuidpoolBlock")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsUuidpoolBlockUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	R_from := d.Get("r_from").(string)
	To := d.Get("to").(string)

	UuidpoolPool := d.Get("uuidpool_pool_dn").(string)

	uuidpoolBlockAttr := models.UuidpoolBlockAttributes{}
	if Child_action, ok := d.GetOk("child_action"); ok {
		uuidpoolBlockAttr.Child_action = Child_action.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		uuidpoolBlockAttr.Sacl = Sacl.(string)
	}

	uuidpoolBlock := models.NewUuidpoolBlock(fmt.Sprintf("block-from-%s-to-%s", R_from, To), UuidpoolPool, desc, uuidpoolBlockAttr)
	uuidpoolBlock.Status = "modified"
	err := ucsClient.Save(uuidpoolBlock)
	if err != nil {
		return err
	}

	d.SetId(uuidpoolBlock.DistinguishedName)
	return resourceUcsUuidpoolBlockRead(d, m)
}
