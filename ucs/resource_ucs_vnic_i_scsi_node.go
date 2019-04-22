package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsVnicIScsiNode() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsVnicIScsiNodeCreate,
		Update: resourceUcsVnicIScsiNodeUpdate,
		Read:   resourceUcsVnicIScsiNodeRead,
		Delete: resourceUcsVnicIScsiNodeDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsVnicIScsiNodeImport,
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

			"flt_aggr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"init_name_suffix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"initiator_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"initiator_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"iqn_ident_pool_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_initiator_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_iqn_ident_pool_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"owner": &schema.Schema{
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
		}),
	}
}

func getRemoteVnicIScsiNode(client *client.Client, dn string) (*models.VnicIScsiNode, error) {
	vnicIScsiNodeDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	vnicIScsiNode := models.VnicIScsiNodeFromDoc(vnicIScsiNodeDoc, "configResolveDn")

	if vnicIScsiNode.DistinguishedName == "" {
		return nil, fmt.Errorf("VnicIScsiNode %s not found", dn)
	}

	return vnicIScsiNode, nil
}

func setVnicIScsiNodeAttributes(vnicIScsiNode *models.VnicIScsiNode, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(vnicIScsiNode.DistinguishedName)
	d.Set("description", vnicIScsiNode.Description)
	d.Set("ls_server_dn", GetParentDn(vnicIScsiNode.DistinguishedName))
	vnicIScsiNodeMap, _ := vnicIScsiNode.ToMap()

	d.Set("child_action", vnicIScsiNodeMap["childAction"])

	d.Set("flt_aggr", vnicIScsiNodeMap["fltAggr"])

	d.Set("init_name_suffix", vnicIScsiNodeMap["initNameSuffix"])

	d.Set("initiator_name", vnicIScsiNodeMap["initiatorName"])

	d.Set("initiator_policy_name", vnicIScsiNodeMap["initiatorPolicyName"])

	d.Set("iqn_ident_pool_name", vnicIScsiNodeMap["iqnIdentPoolName"])

	d.Set("oper_initiator_policy_name", vnicIScsiNodeMap["operInitiatorPolicyName"])

	d.Set("oper_iqn_ident_pool_name", vnicIScsiNodeMap["operIqnIdentPoolName"])

	d.Set("owner", vnicIScsiNodeMap["owner"])

	d.Set("prop_acl", vnicIScsiNodeMap["propAcl"])

	d.Set("sacl", vnicIScsiNodeMap["sacl"])
	return d
}

func resourceUcsVnicIScsiNodeImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	vnicIScsiNode, err := getRemoteVnicIScsiNode(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setVnicIScsiNodeAttributes(vnicIScsiNode, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsVnicIScsiNodeCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	LsServer := d.Get("ls_server_dn").(string)

	vnicIScsiNodeAttr := models.VnicIScsiNodeAttributes{}

	if Child_action, ok := d.GetOk("child_action"); ok {
		vnicIScsiNodeAttr.Child_action = Child_action.(string)
	}

	if Flt_aggr, ok := d.GetOk("flt_aggr"); ok {
		vnicIScsiNodeAttr.Flt_aggr = Flt_aggr.(string)
	}

	if Init_name_suffix, ok := d.GetOk("init_name_suffix"); ok {
		vnicIScsiNodeAttr.Init_name_suffix = Init_name_suffix.(string)
	}

	if Initiator_name, ok := d.GetOk("initiator_name"); ok {
		vnicIScsiNodeAttr.Initiator_name = Initiator_name.(string)
	}

	if Initiator_policy_name, ok := d.GetOk("initiator_policy_name"); ok {
		vnicIScsiNodeAttr.Initiator_policy_name = Initiator_policy_name.(string)
	}

	if Iqn_ident_pool_name, ok := d.GetOk("iqn_ident_pool_name"); ok {
		vnicIScsiNodeAttr.Iqn_ident_pool_name = Iqn_ident_pool_name.(string)
	}

	if Oper_initiator_policy_name, ok := d.GetOk("oper_initiator_policy_name"); ok {
		vnicIScsiNodeAttr.Oper_initiator_policy_name = Oper_initiator_policy_name.(string)
	}

	if Oper_iqn_ident_pool_name, ok := d.GetOk("oper_iqn_ident_pool_name"); ok {
		vnicIScsiNodeAttr.Oper_iqn_ident_pool_name = Oper_iqn_ident_pool_name.(string)
	}

	if Owner, ok := d.GetOk("owner"); ok {
		vnicIScsiNodeAttr.Owner = Owner.(string)
	}

	if Prop_acl, ok := d.GetOk("prop_acl"); ok {
		vnicIScsiNodeAttr.Prop_acl = Prop_acl.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		vnicIScsiNodeAttr.Sacl = Sacl.(string)
	}

	vnicIScsiNode := models.NewVnicIScsiNode(fmt.Sprintf("iscsi-node"), LsServer, desc, vnicIScsiNodeAttr)

	err := ucsClient.Save(vnicIScsiNode)
	if err != nil {
		return err
	}

	d.SetId(vnicIScsiNode.DistinguishedName)
	return resourceUcsVnicIScsiNodeRead(d, m)
}

func resourceUcsVnicIScsiNodeRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	vnicIScsiNode, err := getRemoteVnicIScsiNode(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setVnicIScsiNodeAttributes(vnicIScsiNode, d)

	return nil
}

func resourceUcsVnicIScsiNodeDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "vnicIScsiNode")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsVnicIScsiNodeUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	LsServer := d.Get("ls_server_dn").(string)

	vnicIScsiNodeAttr := models.VnicIScsiNodeAttributes{}
	if Child_action, ok := d.GetOk("child_action"); ok {
		vnicIScsiNodeAttr.Child_action = Child_action.(string)
	}
	if Flt_aggr, ok := d.GetOk("flt_aggr"); ok {
		vnicIScsiNodeAttr.Flt_aggr = Flt_aggr.(string)
	}
	if Init_name_suffix, ok := d.GetOk("init_name_suffix"); ok {
		vnicIScsiNodeAttr.Init_name_suffix = Init_name_suffix.(string)
	}
	if Initiator_name, ok := d.GetOk("initiator_name"); ok {
		vnicIScsiNodeAttr.Initiator_name = Initiator_name.(string)
	}
	if Initiator_policy_name, ok := d.GetOk("initiator_policy_name"); ok {
		vnicIScsiNodeAttr.Initiator_policy_name = Initiator_policy_name.(string)
	}
	if Iqn_ident_pool_name, ok := d.GetOk("iqn_ident_pool_name"); ok {
		vnicIScsiNodeAttr.Iqn_ident_pool_name = Iqn_ident_pool_name.(string)
	}
	if Oper_initiator_policy_name, ok := d.GetOk("oper_initiator_policy_name"); ok {
		vnicIScsiNodeAttr.Oper_initiator_policy_name = Oper_initiator_policy_name.(string)
	}
	if Oper_iqn_ident_pool_name, ok := d.GetOk("oper_iqn_ident_pool_name"); ok {
		vnicIScsiNodeAttr.Oper_iqn_ident_pool_name = Oper_iqn_ident_pool_name.(string)
	}
	if Owner, ok := d.GetOk("owner"); ok {
		vnicIScsiNodeAttr.Owner = Owner.(string)
	}
	if Prop_acl, ok := d.GetOk("prop_acl"); ok {
		vnicIScsiNodeAttr.Prop_acl = Prop_acl.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		vnicIScsiNodeAttr.Sacl = Sacl.(string)
	}

	vnicIScsiNode := models.NewVnicIScsiNode(fmt.Sprintf("iscsi-node"), LsServer, desc, vnicIScsiNodeAttr)
	vnicIScsiNode.Status = "modified"
	err := ucsClient.Save(vnicIScsiNode)
	if err != nil {
		return err
	}

	d.SetId(vnicIScsiNode.DistinguishedName)
	return resourceUcsVnicIScsiNodeRead(d, m)
}
