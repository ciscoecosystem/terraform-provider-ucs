package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsVnicIpV4MgmtPooledAddr() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsVnicIpV4MgmtPooledAddrCreate,
		Update: resourceUcsVnicIpV4MgmtPooledAddrUpdate,
		Read:   resourceUcsVnicIpV4MgmtPooledAddrRead,
		Delete: resourceUcsVnicIpV4MgmtPooledAddrDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsVnicIpV4MgmtPooledAddrImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"mgmt_interface_dn": &schema.Schema{
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

			"def_gw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_name": &schema.Schema{
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

func getRemoteVnicIpV4MgmtPooledAddr(client *client.Client, dn string) (*models.VnicIpV4MgmtPooledAddr, error) {
	vnicIpV4MgmtPooledAddrDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	vnicIpV4MgmtPooledAddr := models.VnicIpV4MgmtPooledAddrFromDoc(vnicIpV4MgmtPooledAddrDoc, "configResolveDn")

	if vnicIpV4MgmtPooledAddr.DistinguishedName == "" {
		return nil, fmt.Errorf("VnicIpV4MgmtPooledAddr %s not found", dn)
	}

	return vnicIpV4MgmtPooledAddr, nil
}

func setVnicIpV4MgmtPooledAddrAttributes(vnicIpV4MgmtPooledAddr *models.VnicIpV4MgmtPooledAddr, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(vnicIpV4MgmtPooledAddr.DistinguishedName)
	d.Set("description", vnicIpV4MgmtPooledAddr.Description)
	d.Set("mgmt_interface_dn", GetParentDn(vnicIpV4MgmtPooledAddr.DistinguishedName))
	vnicIpV4MgmtPooledAddrMap, _ := vnicIpV4MgmtPooledAddr.ToMap()

	d.Set("addr", vnicIpV4MgmtPooledAddrMap["addr"])

	d.Set("child_action", vnicIpV4MgmtPooledAddrMap["childAction"])

	d.Set("def_gw", vnicIpV4MgmtPooledAddrMap["defGw"])

	d.Set("name", vnicIpV4MgmtPooledAddrMap["name"])

	d.Set("oper_name", vnicIpV4MgmtPooledAddrMap["operName"])

	d.Set("prim_dns", vnicIpV4MgmtPooledAddrMap["primDns"])

	d.Set("sacl", vnicIpV4MgmtPooledAddrMap["sacl"])

	d.Set("sec_dns", vnicIpV4MgmtPooledAddrMap["secDns"])

	d.Set("subnet", vnicIpV4MgmtPooledAddrMap["subnet"])
	return d
}

func resourceUcsVnicIpV4MgmtPooledAddrImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	vnicIpV4MgmtPooledAddr, err := getRemoteVnicIpV4MgmtPooledAddr(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setVnicIpV4MgmtPooledAddrAttributes(vnicIpV4MgmtPooledAddr, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsVnicIpV4MgmtPooledAddrCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	MgmtInterface := d.Get("mgmt_interface_dn").(string)

	vnicIpV4MgmtPooledAddrAttr := models.VnicIpV4MgmtPooledAddrAttributes{}

	if Addr, ok := d.GetOk("addr"); ok {
		vnicIpV4MgmtPooledAddrAttr.Addr = Addr.(string)
	}

	if Child_action, ok := d.GetOk("child_action"); ok {
		vnicIpV4MgmtPooledAddrAttr.Child_action = Child_action.(string)
	}

	if Def_gw, ok := d.GetOk("def_gw"); ok {
		vnicIpV4MgmtPooledAddrAttr.Def_gw = Def_gw.(string)
	}

	if Oper_name, ok := d.GetOk("oper_name"); ok {
		vnicIpV4MgmtPooledAddrAttr.Oper_name = Oper_name.(string)
	}

	if Prim_dns, ok := d.GetOk("prim_dns"); ok {
		vnicIpV4MgmtPooledAddrAttr.Prim_dns = Prim_dns.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		vnicIpV4MgmtPooledAddrAttr.Sacl = Sacl.(string)
	}

	if Sec_dns, ok := d.GetOk("sec_dns"); ok {
		vnicIpV4MgmtPooledAddrAttr.Sec_dns = Sec_dns.(string)
	}

	if Subnet, ok := d.GetOk("subnet"); ok {
		vnicIpV4MgmtPooledAddrAttr.Subnet = Subnet.(string)
	}

	vnicIpV4MgmtPooledAddr := models.NewVnicIpV4MgmtPooledAddr(fmt.Sprintf("ipv4-pooled-addr"), MgmtInterface, desc, vnicIpV4MgmtPooledAddrAttr)

	err := ucsClient.Save(vnicIpV4MgmtPooledAddr)
	if err != nil {
		return err
	}

	d.SetId(vnicIpV4MgmtPooledAddr.DistinguishedName)
	return resourceUcsVnicIpV4MgmtPooledAddrRead(d, m)
}

func resourceUcsVnicIpV4MgmtPooledAddrRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	vnicIpV4MgmtPooledAddr, err := getRemoteVnicIpV4MgmtPooledAddr(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setVnicIpV4MgmtPooledAddrAttributes(vnicIpV4MgmtPooledAddr, d)

	return nil
}

func resourceUcsVnicIpV4MgmtPooledAddrDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "vnicIpV4MgmtPooledAddr")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsVnicIpV4MgmtPooledAddrUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	MgmtInterface := d.Get("mgmt_interface_dn").(string)

	vnicIpV4MgmtPooledAddrAttr := models.VnicIpV4MgmtPooledAddrAttributes{}
	if Addr, ok := d.GetOk("addr"); ok {
		vnicIpV4MgmtPooledAddrAttr.Addr = Addr.(string)
	}
	if Child_action, ok := d.GetOk("child_action"); ok {
		vnicIpV4MgmtPooledAddrAttr.Child_action = Child_action.(string)
	}
	if Def_gw, ok := d.GetOk("def_gw"); ok {
		vnicIpV4MgmtPooledAddrAttr.Def_gw = Def_gw.(string)
	}
	if Oper_name, ok := d.GetOk("oper_name"); ok {
		vnicIpV4MgmtPooledAddrAttr.Oper_name = Oper_name.(string)
	}
	if Prim_dns, ok := d.GetOk("prim_dns"); ok {
		vnicIpV4MgmtPooledAddrAttr.Prim_dns = Prim_dns.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		vnicIpV4MgmtPooledAddrAttr.Sacl = Sacl.(string)
	}
	if Sec_dns, ok := d.GetOk("sec_dns"); ok {
		vnicIpV4MgmtPooledAddrAttr.Sec_dns = Sec_dns.(string)
	}
	if Subnet, ok := d.GetOk("subnet"); ok {
		vnicIpV4MgmtPooledAddrAttr.Subnet = Subnet.(string)
	}

	vnicIpV4MgmtPooledAddr := models.NewVnicIpV4MgmtPooledAddr(fmt.Sprintf("ipv4-pooled-addr"), MgmtInterface, desc, vnicIpV4MgmtPooledAddrAttr)
	vnicIpV4MgmtPooledAddr.Status = "modified"
	err := ucsClient.Save(vnicIpV4MgmtPooledAddr)
	if err != nil {
		return err
	}

	d.SetId(vnicIpV4MgmtPooledAddr.DistinguishedName)
	return resourceUcsVnicIpV4MgmtPooledAddrRead(d, m)
}
