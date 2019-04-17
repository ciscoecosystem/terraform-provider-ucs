package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsVnicSanConnPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsVnicSanConnPolicyCreate,
		Update: resourceUcsVnicSanConnPolicyUpdate,
		Read:   resourceUcsVnicSanConnPolicyRead,
		Delete: resourceUcsVnicSanConnPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsVnicSanConnPolicyImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"org_org_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"name": &schema.Schema{
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

			"int_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"policy_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"policy_owner": &schema.Schema{
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

func getRemoteVnicSanConnPolicy(client *client.Client, dn string) (*models.VnicSanConnPolicy, error) {
	vnicSanConnPolicyDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	vnicSanConnPolicy := models.VnicSanConnPolicyFromDoc(vnicSanConnPolicyDoc, "configResolveDn")

	if vnicSanConnPolicy.DistinguishedName == "" {
		return nil, fmt.Errorf("VnicSanConnPolicy %s not found", dn)
	}

	return vnicSanConnPolicy, nil
}

func setVnicSanConnPolicyAttributes(vnicSanConnPolicy *models.VnicSanConnPolicy, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(vnicSanConnPolicy.DistinguishedName)
	d.Set("description", vnicSanConnPolicy.Description)
	d.Set("org_org_dn", GetParentDn(vnicSanConnPolicy.DistinguishedName))
	vnicSanConnPolicyMap, _ := vnicSanConnPolicy.ToMap()

	d.Set("child_action", vnicSanConnPolicyMap["childAction"])

	d.Set("flt_aggr", vnicSanConnPolicyMap["fltAggr"])

	d.Set("int_id", vnicSanConnPolicyMap["intId"])

	d.Set("policy_level", vnicSanConnPolicyMap["policyLevel"])

	d.Set("policy_owner", vnicSanConnPolicyMap["policyOwner"])

	d.Set("sacl", vnicSanConnPolicyMap["sacl"])
	return d
}

func resourceUcsVnicSanConnPolicyImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	vnicSanConnPolicy, err := getRemoteVnicSanConnPolicy(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setVnicSanConnPolicyAttributes(vnicSanConnPolicy, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsVnicSanConnPolicyCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	OrgOrg := d.Get("org_org_dn").(string)

	vnicSanConnPolicyAttr := models.VnicSanConnPolicyAttributes{}

	if Child_action, ok := d.GetOk("child_action"); ok {
		vnicSanConnPolicyAttr.Child_action = Child_action.(string)
	}

	if Flt_aggr, ok := d.GetOk("flt_aggr"); ok {
		vnicSanConnPolicyAttr.Flt_aggr = Flt_aggr.(string)
	}

	if Int_id, ok := d.GetOk("int_id"); ok {
		vnicSanConnPolicyAttr.Int_id = Int_id.(string)
	}

	if Policy_level, ok := d.GetOk("policy_level"); ok {
		vnicSanConnPolicyAttr.Policy_level = Policy_level.(string)
	}

	if Policy_owner, ok := d.GetOk("policy_owner"); ok {
		vnicSanConnPolicyAttr.Policy_owner = Policy_owner.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		vnicSanConnPolicyAttr.Sacl = Sacl.(string)
	}

	vnicSanConnPolicy := models.NewVnicSanConnPolicy(fmt.Sprintf("san-conn-pol-%s", Name), OrgOrg, desc, vnicSanConnPolicyAttr)

	err := ucsClient.Save(vnicSanConnPolicy)
	if err != nil {
		return err
	}

	d.SetId(vnicSanConnPolicy.DistinguishedName)
	return resourceUcsVnicSanConnPolicyRead(d, m)
}

func resourceUcsVnicSanConnPolicyRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	vnicSanConnPolicy, err := getRemoteVnicSanConnPolicy(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setVnicSanConnPolicyAttributes(vnicSanConnPolicy, d)

	return nil
}

func resourceUcsVnicSanConnPolicyDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "vnicSanConnPolicy")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsVnicSanConnPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	OrgOrg := d.Get("org_org_dn").(string)

	vnicSanConnPolicyAttr := models.VnicSanConnPolicyAttributes{}
	if Child_action, ok := d.GetOk("child_action"); ok {
		vnicSanConnPolicyAttr.Child_action = Child_action.(string)
	}
	if Flt_aggr, ok := d.GetOk("flt_aggr"); ok {
		vnicSanConnPolicyAttr.Flt_aggr = Flt_aggr.(string)
	}
	if Int_id, ok := d.GetOk("int_id"); ok {
		vnicSanConnPolicyAttr.Int_id = Int_id.(string)
	}
	if Policy_level, ok := d.GetOk("policy_level"); ok {
		vnicSanConnPolicyAttr.Policy_level = Policy_level.(string)
	}
	if Policy_owner, ok := d.GetOk("policy_owner"); ok {
		vnicSanConnPolicyAttr.Policy_owner = Policy_owner.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		vnicSanConnPolicyAttr.Sacl = Sacl.(string)
	}

	vnicSanConnPolicy := models.NewVnicSanConnPolicy(fmt.Sprintf("san-conn-pol-%s", Name), OrgOrg, desc, vnicSanConnPolicyAttr)
	vnicSanConnPolicy.Status = "modified"
	err := ucsClient.Save(vnicSanConnPolicy)
	if err != nil {
		return err
	}

	d.SetId(vnicSanConnPolicy.DistinguishedName)
	return resourceUcsVnicSanConnPolicyRead(d, m)
}
