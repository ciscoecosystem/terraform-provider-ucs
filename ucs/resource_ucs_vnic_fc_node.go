package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsVnicFcNode() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsVnicFcNodeCreate,
		Update: resourceUcsVnicFcNodeUpdate,
		Read:   resourceUcsVnicFcNodeRead,
		Delete: resourceUcsVnicFcNodeDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsVnicFcNodeImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"vnic_san_conn_policy_dn": &schema.Schema{
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

			"flt_aggr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"ident_pool_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"max_derivable_wwpn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_ident_pool_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"owner": &schema.Schema{
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

func getRemoteVnicFcNode(client *client.Client, dn string) (*models.VnicFcNode, error) {
	vnicFcNodeDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	vnicFcNode := models.VnicFcNodeFromDoc(vnicFcNodeDoc, "configResolveDn")

	if vnicFcNode.DistinguishedName == "" {
		return nil, fmt.Errorf("VnicFcNode %s not found", dn)
	}

	return vnicFcNode, nil
}

func setVnicFcNodeAttributes(vnicFcNode *models.VnicFcNode, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(vnicFcNode.DistinguishedName)
	d.Set("description", vnicFcNode.Description)
	d.Set("vnic_san_conn_policy_dn", GetParentDn(vnicFcNode.DistinguishedName))
	vnicFcNodeMap, _ := vnicFcNode.ToMap()

	d.Set("addr", vnicFcNodeMap["addr"])

	d.Set("child_action", vnicFcNodeMap["childAction"])

	d.Set("flt_aggr", vnicFcNodeMap["fltAggr"])

	d.Set("ident_pool_name", vnicFcNodeMap["identPoolName"])

	d.Set("max_derivable_wwpn", vnicFcNodeMap["maxDerivableWWPN"])

	d.Set("oper_ident_pool_name", vnicFcNodeMap["operIdentPoolName"])

	d.Set("owner", vnicFcNodeMap["owner"])

	d.Set("sacl", vnicFcNodeMap["sacl"])
	return d
}

func resourceUcsVnicFcNodeImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	vnicFcNode, err := getRemoteVnicFcNode(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setVnicFcNodeAttributes(vnicFcNode, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsVnicFcNodeCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	VnicSanConnPolicy := d.Get("vnic_san_conn_policy_dn").(string)

	vnicFcNodeAttr := models.VnicFcNodeAttributes{}

	if Addr, ok := d.GetOk("addr"); ok {
		vnicFcNodeAttr.Addr = Addr.(string)
	}

	if Child_action, ok := d.GetOk("child_action"); ok {
		vnicFcNodeAttr.Child_action = Child_action.(string)
	}

	if Flt_aggr, ok := d.GetOk("flt_aggr"); ok {
		vnicFcNodeAttr.Flt_aggr = Flt_aggr.(string)
	}

	if Ident_pool_name, ok := d.GetOk("ident_pool_name"); ok {
		vnicFcNodeAttr.Ident_pool_name = Ident_pool_name.(string)
	}

	if Max_derivable_wwpn, ok := d.GetOk("max_derivable_wwpn"); ok {
		vnicFcNodeAttr.Max_derivable_wwpn = Max_derivable_wwpn.(string)
	}

	if Oper_ident_pool_name, ok := d.GetOk("oper_ident_pool_name"); ok {
		vnicFcNodeAttr.Oper_ident_pool_name = Oper_ident_pool_name.(string)
	}

	if Owner, ok := d.GetOk("owner"); ok {
		vnicFcNodeAttr.Owner = Owner.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		vnicFcNodeAttr.Sacl = Sacl.(string)
	}

	vnicFcNode := models.NewVnicFcNode(fmt.Sprintf("fc-node"), VnicSanConnPolicy, desc, vnicFcNodeAttr)

	err := ucsClient.Save(vnicFcNode)
	if err != nil {
		return err
	}

	d.SetId(vnicFcNode.DistinguishedName)
	return resourceUcsVnicFcNodeRead(d, m)
}

func resourceUcsVnicFcNodeRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	vnicFcNode, err := getRemoteVnicFcNode(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setVnicFcNodeAttributes(vnicFcNode, d)

	return nil
}

func resourceUcsVnicFcNodeDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "vnicFcNode")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsVnicFcNodeUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	VnicSanConnPolicy := d.Get("vnic_san_conn_policy_dn").(string)

	vnicFcNodeAttr := models.VnicFcNodeAttributes{}
	if Addr, ok := d.GetOk("addr"); ok {
		vnicFcNodeAttr.Addr = Addr.(string)
	}
	if Child_action, ok := d.GetOk("child_action"); ok {
		vnicFcNodeAttr.Child_action = Child_action.(string)
	}
	if Flt_aggr, ok := d.GetOk("flt_aggr"); ok {
		vnicFcNodeAttr.Flt_aggr = Flt_aggr.(string)
	}
	if Ident_pool_name, ok := d.GetOk("ident_pool_name"); ok {
		vnicFcNodeAttr.Ident_pool_name = Ident_pool_name.(string)
	}
	if Max_derivable_wwpn, ok := d.GetOk("max_derivable_wwpn"); ok {
		vnicFcNodeAttr.Max_derivable_wwpn = Max_derivable_wwpn.(string)
	}
	if Oper_ident_pool_name, ok := d.GetOk("oper_ident_pool_name"); ok {
		vnicFcNodeAttr.Oper_ident_pool_name = Oper_ident_pool_name.(string)
	}
	if Owner, ok := d.GetOk("owner"); ok {
		vnicFcNodeAttr.Owner = Owner.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		vnicFcNodeAttr.Sacl = Sacl.(string)
	}

	vnicFcNode := models.NewVnicFcNode(fmt.Sprintf("fc-node"), VnicSanConnPolicy, desc, vnicFcNodeAttr)
	vnicFcNode.Status = "modified"
	err := ucsClient.Save(vnicFcNode)
	if err != nil {
		return err
	}

	d.SetId(vnicFcNode.DistinguishedName)
	return resourceUcsVnicFcNodeRead(d, m)
}
