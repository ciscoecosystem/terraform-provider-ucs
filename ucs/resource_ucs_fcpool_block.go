package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsFcpoolBlock() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsFcpoolBlockCreate,
		Update: resourceUcsFcpoolBlockUpdate,
		Read:   resourceUcsFcpoolBlockRead,
		Delete: resourceUcsFcpoolBlockDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsFcpoolBlockImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"fcpool_initiators_dn": &schema.Schema{
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

func getRemoteFcpoolBlock(client *client.Client, dn string) (*models.FcpoolBlock, error) {
	fcpoolBlockDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	fcpoolBlock := models.FcpoolBlockFromDoc(fcpoolBlockDoc, "configResolveDn")

	if fcpoolBlock.DistinguishedName == "" {
		return nil, fmt.Errorf("FcpoolBlock %s not found", dn)
	}

	return fcpoolBlock, nil
}

func setFcpoolBlockAttributes(fcpoolBlock *models.FcpoolBlock, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(fcpoolBlock.DistinguishedName)
	d.Set("description", fcpoolBlock.Description)
	d.Set("fcpool_initiators_dn", GetParentDn(fcpoolBlock.DistinguishedName))
	fcpoolBlockMap, _ := fcpoolBlock.ToMap()

	d.Set("child_action", fcpoolBlockMap["childAction"])

	d.Set("sacl", fcpoolBlockMap["sacl"])
	return d
}

func resourceUcsFcpoolBlockImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	fcpoolBlock, err := getRemoteFcpoolBlock(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setFcpoolBlockAttributes(fcpoolBlock, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsFcpoolBlockCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	R_from := d.Get("r_from").(string)
	To := d.Get("to").(string)

	FcpoolInitiators := d.Get("fcpool_initiators_dn").(string)

	fcpoolBlockAttr := models.FcpoolBlockAttributes{}

	if Child_action, ok := d.GetOk("child_action"); ok {
		fcpoolBlockAttr.Child_action = Child_action.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		fcpoolBlockAttr.Sacl = Sacl.(string)
	}

	fcpoolBlock := models.NewFcpoolBlock(fmt.Sprintf("block-%s-%s", R_from, To), FcpoolInitiators, desc, fcpoolBlockAttr)

	err := ucsClient.Save(fcpoolBlock)
	if err != nil {
		return err
	}

	d.SetId(fcpoolBlock.DistinguishedName)
	return resourceUcsFcpoolBlockRead(d, m)
}

func resourceUcsFcpoolBlockRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	fcpoolBlock, err := getRemoteFcpoolBlock(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setFcpoolBlockAttributes(fcpoolBlock, d)

	return nil
}

func resourceUcsFcpoolBlockDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "fcpoolBlock")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsFcpoolBlockUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	R_from := d.Get("r_from").(string)
	To := d.Get("to").(string)

	FcpoolInitiators := d.Get("fcpool_initiators_dn").(string)

	fcpoolBlockAttr := models.FcpoolBlockAttributes{}
	if Child_action, ok := d.GetOk("child_action"); ok {
		fcpoolBlockAttr.Child_action = Child_action.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		fcpoolBlockAttr.Sacl = Sacl.(string)
	}

	fcpoolBlock := models.NewFcpoolBlock(fmt.Sprintf("block-%s-%s", R_from, To), FcpoolInitiators, desc, fcpoolBlockAttr)
	fcpoolBlock.Status = "modified"
	err := ucsClient.Save(fcpoolBlock)
	if err != nil {
		return err
	}

	d.SetId(fcpoolBlock.DistinguishedName)
	return resourceUcsFcpoolBlockRead(d, m)
}
