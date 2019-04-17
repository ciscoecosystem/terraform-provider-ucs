package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsIppoolBlock() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsIppoolBlockCreate,
		Update: resourceUcsIppoolBlockUpdate,
		Read:   resourceUcsIppoolBlockRead,
		Delete: resourceUcsIppoolBlockDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsIppoolBlockImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"ippool_pool_dn": &schema.Schema{
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

			"def_gw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"prim_dns": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"sacl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"sec_dns": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"subnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		}),
	}
}

func getRemoteIppoolBlock(client *client.Client, dn string) (*models.IppoolBlock, error) {
	ippoolBlockDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	ippoolBlock := models.IppoolBlockFromDoc(ippoolBlockDoc, "configResolveDn")

	if ippoolBlock.DistinguishedName == "" {
		return nil, fmt.Errorf("IppoolBlock %s not found", dn)
	}

	return ippoolBlock, nil
}

func setIppoolBlockAttributes(ippoolBlock *models.IppoolBlock, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(ippoolBlock.DistinguishedName)
	d.Set("description", ippoolBlock.Description)
	d.Set("ippool_pool_dn", GetParentDn(ippoolBlock.DistinguishedName))
	ippoolBlockMap, _ := ippoolBlock.ToMap()

	d.Set("child_action", ippoolBlockMap["childAction"])

	d.Set("def_gw", ippoolBlockMap["defGw"])

	d.Set("prim_dns", ippoolBlockMap["primDns"])

	d.Set("sacl", ippoolBlockMap["sacl"])

	d.Set("sec_dns", ippoolBlockMap["secDns"])

	d.Set("subnet", ippoolBlockMap["subnet"])
	return d
}

func resourceUcsIppoolBlockImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	ippoolBlock, err := getRemoteIppoolBlock(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setIppoolBlockAttributes(ippoolBlock, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsIppoolBlockCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	R_from := d.Get("r_from").(string)
	To := d.Get("to").(string)

	IppoolPool := d.Get("ippool_pool_dn").(string)

	ippoolBlockAttr := models.IppoolBlockAttributes{}

	if Child_action, ok := d.GetOk("child_action"); ok {
		ippoolBlockAttr.Child_action = Child_action.(string)
	}

	if Def_gw, ok := d.GetOk("def_gw"); ok {
		ippoolBlockAttr.Def_gw = Def_gw.(string)
	}

	if Prim_dns, ok := d.GetOk("prim_dns"); ok {
		ippoolBlockAttr.Prim_dns = Prim_dns.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		ippoolBlockAttr.Sacl = Sacl.(string)
	}

	if Sec_dns, ok := d.GetOk("sec_dns"); ok {
		ippoolBlockAttr.Sec_dns = Sec_dns.(string)
	}

	if Subnet, ok := d.GetOk("subnet"); ok {
		ippoolBlockAttr.Subnet = Subnet.(string)
	}

	ippoolBlock := models.NewIppoolBlock(fmt.Sprintf("block-%s-%s", R_from, To), IppoolPool, desc, ippoolBlockAttr)

	err := ucsClient.Save(ippoolBlock)
	if err != nil {
		return err
	}

	d.SetId(ippoolBlock.DistinguishedName)
	return resourceUcsIppoolBlockRead(d, m)
}

func resourceUcsIppoolBlockRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	ippoolBlock, err := getRemoteIppoolBlock(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setIppoolBlockAttributes(ippoolBlock, d)

	return nil
}

func resourceUcsIppoolBlockDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "ippoolBlock")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsIppoolBlockUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	R_from := d.Get("r_from").(string)
	To := d.Get("to").(string)

	IppoolPool := d.Get("ippool_pool_dn").(string)

	ippoolBlockAttr := models.IppoolBlockAttributes{}
	if Child_action, ok := d.GetOk("child_action"); ok {
		ippoolBlockAttr.Child_action = Child_action.(string)
	}
	if Def_gw, ok := d.GetOk("def_gw"); ok {
		ippoolBlockAttr.Def_gw = Def_gw.(string)
	}
	if Prim_dns, ok := d.GetOk("prim_dns"); ok {
		ippoolBlockAttr.Prim_dns = Prim_dns.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		ippoolBlockAttr.Sacl = Sacl.(string)
	}
	if Sec_dns, ok := d.GetOk("sec_dns"); ok {
		ippoolBlockAttr.Sec_dns = Sec_dns.(string)
	}
	if Subnet, ok := d.GetOk("subnet"); ok {
		ippoolBlockAttr.Subnet = Subnet.(string)
	}

	ippoolBlock := models.NewIppoolBlock(fmt.Sprintf("block-%s-%s", R_from, To), IppoolPool, desc, ippoolBlockAttr)
	ippoolBlock.Status = "modified"
	err := ucsClient.Save(ippoolBlock)
	if err != nil {
		return err
	}

	d.SetId(ippoolBlock.DistinguishedName)
	return resourceUcsIppoolBlockRead(d, m)
}
