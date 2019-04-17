package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsIppoolIpV6Block() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsIppoolIpV6BlockCreate,
		Update: resourceUcsIppoolIpV6BlockUpdate,
		Read:   resourceUcsIppoolIpV6BlockRead,
		Delete: resourceUcsIppoolIpV6BlockDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsIppoolIpV6BlockImport,
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

			"prefix": &schema.Schema{
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
		}),
	}
}

func getRemoteIppoolIpV6Block(client *client.Client, dn string) (*models.IppoolIpV6Block, error) {
	ippoolIpV6BlockDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	ippoolIpV6Block := models.IppoolIpV6BlockFromDoc(ippoolIpV6BlockDoc, "configResolveDn")

	if ippoolIpV6Block.DistinguishedName == "" {
		return nil, fmt.Errorf("IppoolIpV6Block %s not found", dn)
	}

	return ippoolIpV6Block, nil
}

func setIppoolIpV6BlockAttributes(ippoolIpV6Block *models.IppoolIpV6Block, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(ippoolIpV6Block.DistinguishedName)
	d.Set("description", ippoolIpV6Block.Description)
	d.Set("ippool_pool_dn", GetParentDn(ippoolIpV6Block.DistinguishedName))
	ippoolIpV6BlockMap, _ := ippoolIpV6Block.ToMap()

	d.Set("child_action", ippoolIpV6BlockMap["childAction"])

	d.Set("def_gw", ippoolIpV6BlockMap["defGw"])

	d.Set("prefix", ippoolIpV6BlockMap["prefix"])

	d.Set("prim_dns", ippoolIpV6BlockMap["primDns"])

	d.Set("sacl", ippoolIpV6BlockMap["sacl"])

	d.Set("sec_dns", ippoolIpV6BlockMap["secDns"])
	return d
}

func resourceUcsIppoolIpV6BlockImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	ippoolIpV6Block, err := getRemoteIppoolIpV6Block(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setIppoolIpV6BlockAttributes(ippoolIpV6Block, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsIppoolIpV6BlockCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	R_from := d.Get("r_from").(string)
	To := d.Get("to").(string)

	IppoolPool := d.Get("ippool_pool_dn").(string)

	ippoolIpV6BlockAttr := models.IppoolIpV6BlockAttributes{}

	if Child_action, ok := d.GetOk("child_action"); ok {
		ippoolIpV6BlockAttr.Child_action = Child_action.(string)
	}

	if Def_gw, ok := d.GetOk("def_gw"); ok {
		ippoolIpV6BlockAttr.Def_gw = Def_gw.(string)
	}

	if Prefix, ok := d.GetOk("prefix"); ok {
		ippoolIpV6BlockAttr.Prefix = Prefix.(string)
	}

	if Prim_dns, ok := d.GetOk("prim_dns"); ok {
		ippoolIpV6BlockAttr.Prim_dns = Prim_dns.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		ippoolIpV6BlockAttr.Sacl = Sacl.(string)
	}

	if Sec_dns, ok := d.GetOk("sec_dns"); ok {
		ippoolIpV6BlockAttr.Sec_dns = Sec_dns.(string)
	}

	ippoolIpV6Block := models.NewIppoolIpV6Block(fmt.Sprintf("v6block-%s-%s", R_from, To), IppoolPool, desc, ippoolIpV6BlockAttr)

	err := ucsClient.Save(ippoolIpV6Block)
	if err != nil {
		return err
	}

	d.SetId(ippoolIpV6Block.DistinguishedName)
	return resourceUcsIppoolIpV6BlockRead(d, m)
}

func resourceUcsIppoolIpV6BlockRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	ippoolIpV6Block, err := getRemoteIppoolIpV6Block(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setIppoolIpV6BlockAttributes(ippoolIpV6Block, d)

	return nil
}

func resourceUcsIppoolIpV6BlockDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "ippoolIpV6Block")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsIppoolIpV6BlockUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	R_from := d.Get("r_from").(string)
	To := d.Get("to").(string)

	IppoolPool := d.Get("ippool_pool_dn").(string)

	ippoolIpV6BlockAttr := models.IppoolIpV6BlockAttributes{}
	if Child_action, ok := d.GetOk("child_action"); ok {
		ippoolIpV6BlockAttr.Child_action = Child_action.(string)
	}
	if Def_gw, ok := d.GetOk("def_gw"); ok {
		ippoolIpV6BlockAttr.Def_gw = Def_gw.(string)
	}
	if Prefix, ok := d.GetOk("prefix"); ok {
		ippoolIpV6BlockAttr.Prefix = Prefix.(string)
	}
	if Prim_dns, ok := d.GetOk("prim_dns"); ok {
		ippoolIpV6BlockAttr.Prim_dns = Prim_dns.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		ippoolIpV6BlockAttr.Sacl = Sacl.(string)
	}
	if Sec_dns, ok := d.GetOk("sec_dns"); ok {
		ippoolIpV6BlockAttr.Sec_dns = Sec_dns.(string)
	}

	ippoolIpV6Block := models.NewIppoolIpV6Block(fmt.Sprintf("v6block-%s-%s", R_from, To), IppoolPool, desc, ippoolIpV6BlockAttr)
	ippoolIpV6Block.Status = "modified"
	err := ucsClient.Save(ippoolIpV6Block)
	if err != nil {
		return err
	}

	d.SetId(ippoolIpV6Block.DistinguishedName)
	return resourceUcsIppoolIpV6BlockRead(d, m)
}
