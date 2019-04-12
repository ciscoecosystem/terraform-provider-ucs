package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsLstorageDiskGroupConfigPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsLstorageDiskGroupConfigPolicyCreate,
		Update: resourceUcsLstorageDiskGroupConfigPolicyUpdate,
		Read:   resourceUcsLstorageDiskGroupConfigPolicyRead,
		Delete: resourceUcsLstorageDiskGroupConfigPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsLstorageDiskGroupConfigPolicyImport,
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

			"raid_level": &schema.Schema{
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

func getRemoteLstorageDiskGroupConfigPolicy(client *client.Client, dn string) (*models.LstorageDiskGroupConfigPolicy, error) {
	lstorageDiskGroupConfigPolicyDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	lstorageDiskGroupConfigPolicy := models.LstorageDiskGroupConfigPolicyFromDoc(lstorageDiskGroupConfigPolicyDoc, "configResolveDn")

	if lstorageDiskGroupConfigPolicy.DistinguishedName == "" {
		return nil, fmt.Errorf("LstorageDiskGroupConfigPolicy %s not found", dn)
	}

	return lstorageDiskGroupConfigPolicy, nil
}

func setLstorageDiskGroupConfigPolicyAttributes(lstorageDiskGroupConfigPolicy *models.LstorageDiskGroupConfigPolicy, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(lstorageDiskGroupConfigPolicy.DistinguishedName)
	d.Set("description", lstorageDiskGroupConfigPolicy.Description)
	d.Set("org_org_dn", GetParentDn(lstorageDiskGroupConfigPolicy.DistinguishedName))
	lstorageDiskGroupConfigPolicyMap, _ := lstorageDiskGroupConfigPolicy.ToMap()

	d.Set("child_action", lstorageDiskGroupConfigPolicyMap["childAction"])

	d.Set("int_id", lstorageDiskGroupConfigPolicyMap["intId"])

	d.Set("policy_level", lstorageDiskGroupConfigPolicyMap["policyLevel"])

	d.Set("policy_owner", lstorageDiskGroupConfigPolicyMap["policyOwner"])

	d.Set("raid_level", lstorageDiskGroupConfigPolicyMap["raidLevel"])

	d.Set("sacl", lstorageDiskGroupConfigPolicyMap["sacl"])
	return d
}

func resourceUcsLstorageDiskGroupConfigPolicyImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	lstorageDiskGroupConfigPolicy, err := getRemoteLstorageDiskGroupConfigPolicy(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setLstorageDiskGroupConfigPolicyAttributes(lstorageDiskGroupConfigPolicy, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsLstorageDiskGroupConfigPolicyCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	OrgOrg := d.Get("org_org_dn").(string)

	lstorageDiskGroupConfigPolicyAttr := models.LstorageDiskGroupConfigPolicyAttributes{}

	if Child_action, ok := d.GetOk("child_action"); ok {
		lstorageDiskGroupConfigPolicyAttr.Child_action = Child_action.(string)
	}

	if Int_id, ok := d.GetOk("int_id"); ok {
		lstorageDiskGroupConfigPolicyAttr.Int_id = Int_id.(string)
	}

	if Policy_level, ok := d.GetOk("policy_level"); ok {
		lstorageDiskGroupConfigPolicyAttr.Policy_level = Policy_level.(string)
	}

	if Policy_owner, ok := d.GetOk("policy_owner"); ok {
		lstorageDiskGroupConfigPolicyAttr.Policy_owner = Policy_owner.(string)
	}

	if Raid_level, ok := d.GetOk("raid_level"); ok {
		lstorageDiskGroupConfigPolicyAttr.Raid_level = Raid_level.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		lstorageDiskGroupConfigPolicyAttr.Sacl = Sacl.(string)
	}

	lstorageDiskGroupConfigPolicy := models.NewLstorageDiskGroupConfigPolicy(fmt.Sprintf("disk-group-config-%s", Name), OrgOrg, desc, lstorageDiskGroupConfigPolicyAttr)

	err := ucsClient.Save(lstorageDiskGroupConfigPolicy)
	if err != nil {
		return err
	}

	d.SetId(lstorageDiskGroupConfigPolicy.DistinguishedName)
	return resourceUcsLstorageDiskGroupConfigPolicyRead(d, m)
}

func resourceUcsLstorageDiskGroupConfigPolicyRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	lstorageDiskGroupConfigPolicy, err := getRemoteLstorageDiskGroupConfigPolicy(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setLstorageDiskGroupConfigPolicyAttributes(lstorageDiskGroupConfigPolicy, d)

	return nil
}

func resourceUcsLstorageDiskGroupConfigPolicyDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "lstorageDiskGroupConfigPolicy")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsLstorageDiskGroupConfigPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	OrgOrg := d.Get("org_org_dn").(string)

	lstorageDiskGroupConfigPolicyAttr := models.LstorageDiskGroupConfigPolicyAttributes{}
	if Child_action, ok := d.GetOk("child_action"); ok {
		lstorageDiskGroupConfigPolicyAttr.Child_action = Child_action.(string)
	}
	if Int_id, ok := d.GetOk("int_id"); ok {
		lstorageDiskGroupConfigPolicyAttr.Int_id = Int_id.(string)
	}
	if Policy_level, ok := d.GetOk("policy_level"); ok {
		lstorageDiskGroupConfigPolicyAttr.Policy_level = Policy_level.(string)
	}
	if Policy_owner, ok := d.GetOk("policy_owner"); ok {
		lstorageDiskGroupConfigPolicyAttr.Policy_owner = Policy_owner.(string)
	}
	if Raid_level, ok := d.GetOk("raid_level"); ok {
		lstorageDiskGroupConfigPolicyAttr.Raid_level = Raid_level.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		lstorageDiskGroupConfigPolicyAttr.Sacl = Sacl.(string)
	}

	lstorageDiskGroupConfigPolicy := models.NewLstorageDiskGroupConfigPolicy(fmt.Sprintf("disk-group-config-%s", Name), OrgOrg, desc, lstorageDiskGroupConfigPolicyAttr)
	lstorageDiskGroupConfigPolicy.Status = "modified"
	err := ucsClient.Save(lstorageDiskGroupConfigPolicy)
	if err != nil {
		return err
	}

	d.SetId(lstorageDiskGroupConfigPolicy.DistinguishedName)
	return resourceUcsLstorageDiskGroupConfigPolicyRead(d, m)
}
