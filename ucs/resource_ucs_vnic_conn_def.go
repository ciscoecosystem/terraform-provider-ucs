package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsVnicConnDef() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsVnicConnDefCreate,
		Update: resourceUcsVnicConnDefUpdate,
		Read:   resourceUcsVnicConnDefRead,
		Delete: resourceUcsVnicConnDefDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsVnicConnDefImport,
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

			"lan_conn_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_lan_conn_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_san_conn_policy_name": &schema.Schema{
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

			"san_conn_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		}),
	}
}

func getRemoteVnicConnDef(client *client.Client, dn string) (*models.VnicConnDef, error) {
	vnicConnDefDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	vnicConnDef := models.VnicConnDefFromDoc(vnicConnDefDoc, "configResolveDn")

	if vnicConnDef.DistinguishedName == "" {
		return nil, fmt.Errorf("VnicConnDef %s not found", dn)
	}

	return vnicConnDef, nil
}

func setVnicConnDefAttributes(vnicConnDef *models.VnicConnDef, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(vnicConnDef.DistinguishedName)
	d.Set("description", vnicConnDef.Description)
	d.Set("ls_server_dn", GetParentDn(vnicConnDef.DistinguishedName))
	vnicConnDefMap, _ := vnicConnDef.ToMap()

	d.Set("child_action", vnicConnDefMap["childAction"])

	d.Set("flt_aggr", vnicConnDefMap["fltAggr"])

	d.Set("lan_conn_policy_name", vnicConnDefMap["lanConnPolicyName"])

	d.Set("oper_lan_conn_policy_name", vnicConnDefMap["operLanConnPolicyName"])

	d.Set("oper_san_conn_policy_name", vnicConnDefMap["operSanConnPolicyName"])

	d.Set("prop_acl", vnicConnDefMap["propAcl"])

	d.Set("sacl", vnicConnDefMap["sacl"])

	d.Set("san_conn_policy_name", vnicConnDefMap["sanConnPolicyName"])
	return d
}

func resourceUcsVnicConnDefImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	vnicConnDef, err := getRemoteVnicConnDef(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setVnicConnDefAttributes(vnicConnDef, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsVnicConnDefCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	LsServer := d.Get("ls_server_dn").(string)

	vnicConnDefAttr := models.VnicConnDefAttributes{}

	if Child_action, ok := d.GetOk("child_action"); ok {
		vnicConnDefAttr.Child_action = Child_action.(string)
	}

	if Flt_aggr, ok := d.GetOk("flt_aggr"); ok {
		vnicConnDefAttr.Flt_aggr = Flt_aggr.(string)
	}

	if Lan_conn_policy_name, ok := d.GetOk("lan_conn_policy_name"); ok {
		vnicConnDefAttr.Lan_conn_policy_name = Lan_conn_policy_name.(string)
	}

	if Oper_lan_conn_policy_name, ok := d.GetOk("oper_lan_conn_policy_name"); ok {
		vnicConnDefAttr.Oper_lan_conn_policy_name = Oper_lan_conn_policy_name.(string)
	}

	if Oper_san_conn_policy_name, ok := d.GetOk("oper_san_conn_policy_name"); ok {
		vnicConnDefAttr.Oper_san_conn_policy_name = Oper_san_conn_policy_name.(string)
	}

	if Prop_acl, ok := d.GetOk("prop_acl"); ok {
		vnicConnDefAttr.Prop_acl = Prop_acl.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		vnicConnDefAttr.Sacl = Sacl.(string)
	}

	if San_conn_policy_name, ok := d.GetOk("san_conn_policy_name"); ok {
		vnicConnDefAttr.San_conn_policy_name = San_conn_policy_name.(string)
	}

	vnicConnDef := models.NewVnicConnDef(fmt.Sprintf("conn-def"), LsServer, desc, vnicConnDefAttr)

	err := ucsClient.Save(vnicConnDef)
	if err != nil {
		return err
	}

	d.SetId(vnicConnDef.DistinguishedName)
	return resourceUcsVnicConnDefRead(d, m)
}

func resourceUcsVnicConnDefRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	vnicConnDef, err := getRemoteVnicConnDef(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setVnicConnDefAttributes(vnicConnDef, d)

	return nil
}

func resourceUcsVnicConnDefDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "vnicConnDef")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsVnicConnDefUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	LsServer := d.Get("ls_server_dn").(string)

	vnicConnDefAttr := models.VnicConnDefAttributes{}
	if Child_action, ok := d.GetOk("child_action"); ok {
		vnicConnDefAttr.Child_action = Child_action.(string)
	}
	if Flt_aggr, ok := d.GetOk("flt_aggr"); ok {
		vnicConnDefAttr.Flt_aggr = Flt_aggr.(string)
	}
	if Lan_conn_policy_name, ok := d.GetOk("lan_conn_policy_name"); ok {
		vnicConnDefAttr.Lan_conn_policy_name = Lan_conn_policy_name.(string)
	}
	if Oper_lan_conn_policy_name, ok := d.GetOk("oper_lan_conn_policy_name"); ok {
		vnicConnDefAttr.Oper_lan_conn_policy_name = Oper_lan_conn_policy_name.(string)
	}
	if Oper_san_conn_policy_name, ok := d.GetOk("oper_san_conn_policy_name"); ok {
		vnicConnDefAttr.Oper_san_conn_policy_name = Oper_san_conn_policy_name.(string)
	}
	if Prop_acl, ok := d.GetOk("prop_acl"); ok {
		vnicConnDefAttr.Prop_acl = Prop_acl.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		vnicConnDefAttr.Sacl = Sacl.(string)
	}
	if San_conn_policy_name, ok := d.GetOk("san_conn_policy_name"); ok {
		vnicConnDefAttr.San_conn_policy_name = San_conn_policy_name.(string)
	}

	vnicConnDef := models.NewVnicConnDef(fmt.Sprintf("conn-def"), LsServer, desc, vnicConnDefAttr)
	vnicConnDef.Status = "modified"
	err := ucsClient.Save(vnicConnDef)
	if err != nil {
		return err
	}

	d.SetId(vnicConnDef.DistinguishedName)
	return resourceUcsVnicConnDefRead(d, m)
}
