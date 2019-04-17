package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsVnicLanConnPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsVnicLanConnPolicyCreate,
		Update: resourceUcsVnicLanConnPolicyUpdate,
		Read:   resourceUcsVnicLanConnPolicyRead,
		Delete: resourceUcsVnicLanConnPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsVnicLanConnPolicyImport,
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

func getRemoteVnicLanConnPolicy(client *client.Client, dn string) (*models.VnicLanConnPolicy, error) {
	vnicLanConnPolicyDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	vnicLanConnPolicy := models.VnicLanConnPolicyFromDoc(vnicLanConnPolicyDoc, "configResolveDn")

	if vnicLanConnPolicy.DistinguishedName == "" {
		return nil, fmt.Errorf("VnicLanConnPolicy %s not found", dn)
	}

	return vnicLanConnPolicy, nil
}

func setVnicLanConnPolicyAttributes(vnicLanConnPolicy *models.VnicLanConnPolicy, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(vnicLanConnPolicy.DistinguishedName)
	d.Set("description", vnicLanConnPolicy.Description)
	d.Set("org_org_dn", GetParentDn(vnicLanConnPolicy.DistinguishedName))
	vnicLanConnPolicyMap, _ := vnicLanConnPolicy.ToMap()

	d.Set("child_action", vnicLanConnPolicyMap["childAction"])

	d.Set("flt_aggr", vnicLanConnPolicyMap["fltAggr"])

	d.Set("int_id", vnicLanConnPolicyMap["intId"])

	d.Set("policy_level", vnicLanConnPolicyMap["policyLevel"])

	d.Set("policy_owner", vnicLanConnPolicyMap["policyOwner"])

	d.Set("sacl", vnicLanConnPolicyMap["sacl"])
	return d
}

func resourceUcsVnicLanConnPolicyImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	vnicLanConnPolicy, err := getRemoteVnicLanConnPolicy(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setVnicLanConnPolicyAttributes(vnicLanConnPolicy, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsVnicLanConnPolicyCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	OrgOrg := d.Get("org_org_dn").(string)

	vnicLanConnPolicyAttr := models.VnicLanConnPolicyAttributes{}

	if Child_action, ok := d.GetOk("child_action"); ok {
		vnicLanConnPolicyAttr.Child_action = Child_action.(string)
	}

	if Flt_aggr, ok := d.GetOk("flt_aggr"); ok {
		vnicLanConnPolicyAttr.Flt_aggr = Flt_aggr.(string)
	}

	if Int_id, ok := d.GetOk("int_id"); ok {
		vnicLanConnPolicyAttr.Int_id = Int_id.(string)
	}

	if Policy_level, ok := d.GetOk("policy_level"); ok {
		vnicLanConnPolicyAttr.Policy_level = Policy_level.(string)
	}

	if Policy_owner, ok := d.GetOk("policy_owner"); ok {
		vnicLanConnPolicyAttr.Policy_owner = Policy_owner.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		vnicLanConnPolicyAttr.Sacl = Sacl.(string)
	}

	vnicLanConnPolicy := models.NewVnicLanConnPolicy(fmt.Sprintf("lan-conn-pol-%s", Name), OrgOrg, desc, vnicLanConnPolicyAttr)

	err := ucsClient.Save(vnicLanConnPolicy)
	if err != nil {
		return err
	}

	d.SetId(vnicLanConnPolicy.DistinguishedName)
	return resourceUcsVnicLanConnPolicyRead(d, m)
}

func resourceUcsVnicLanConnPolicyRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	vnicLanConnPolicy, err := getRemoteVnicLanConnPolicy(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setVnicLanConnPolicyAttributes(vnicLanConnPolicy, d)

	return nil
}

func resourceUcsVnicLanConnPolicyDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "vnicLanConnPolicy")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsVnicLanConnPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	OrgOrg := d.Get("org_org_dn").(string)

	vnicLanConnPolicyAttr := models.VnicLanConnPolicyAttributes{}
	if Child_action, ok := d.GetOk("child_action"); ok {
		vnicLanConnPolicyAttr.Child_action = Child_action.(string)
	}
	if Flt_aggr, ok := d.GetOk("flt_aggr"); ok {
		vnicLanConnPolicyAttr.Flt_aggr = Flt_aggr.(string)
	}
	if Int_id, ok := d.GetOk("int_id"); ok {
		vnicLanConnPolicyAttr.Int_id = Int_id.(string)
	}
	if Policy_level, ok := d.GetOk("policy_level"); ok {
		vnicLanConnPolicyAttr.Policy_level = Policy_level.(string)
	}
	if Policy_owner, ok := d.GetOk("policy_owner"); ok {
		vnicLanConnPolicyAttr.Policy_owner = Policy_owner.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		vnicLanConnPolicyAttr.Sacl = Sacl.(string)
	}

	vnicLanConnPolicy := models.NewVnicLanConnPolicy(fmt.Sprintf("lan-conn-pol-%s", Name), OrgOrg, desc, vnicLanConnPolicyAttr)
	vnicLanConnPolicy.Status = "modified"
	err := ucsClient.Save(vnicLanConnPolicy)
	if err != nil {
		return err
	}

	d.SetId(vnicLanConnPolicy.DistinguishedName)
	return resourceUcsVnicLanConnPolicyRead(d, m)
}
