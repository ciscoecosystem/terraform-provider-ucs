package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsLstorageProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsLstorageProfileCreate,
		Update: resourceUcsLstorageProfileUpdate,
		Read:   resourceUcsLstorageProfileRead,
		Delete: resourceUcsLstorageProfileDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsLstorageProfileImport,
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

			"assigned_to_dn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"availability": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

			"sacl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		}),
	}
}

func getRemoteLstorageProfile(client *client.Client, dn string) (*models.LstorageProfile, error) {
	lstorageProfileDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	lstorageProfile := models.LstorageProfileFromDoc(lstorageProfileDoc, "configResolveDn")

	if lstorageProfile.DistinguishedName == "" {
		return nil, fmt.Errorf("LstorageProfile %s not found", dn)
	}

	return lstorageProfile, nil
}

func setLstorageProfileAttributes(lstorageProfile *models.LstorageProfile, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(lstorageProfile.DistinguishedName)
	d.Set("description", lstorageProfile.Description)
	d.Set("org_org_dn", GetParentDn(lstorageProfile.DistinguishedName))
	lstorageProfileMap, _ := lstorageProfile.ToMap()

	d.Set("assigned_to_dn", lstorageProfileMap["assignedToDn"])

	d.Set("availability", lstorageProfileMap["availability"])

	d.Set("child_action", lstorageProfileMap["childAction"])

	d.Set("int_id", lstorageProfileMap["intId"])

	d.Set("policy_level", lstorageProfileMap["policyLevel"])

	d.Set("policy_owner", lstorageProfileMap["policyOwner"])

	d.Set("sacl", lstorageProfileMap["sacl"])
	return d
}

func resourceUcsLstorageProfileImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	lstorageProfile, err := getRemoteLstorageProfile(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setLstorageProfileAttributes(lstorageProfile, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsLstorageProfileCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	OrgOrg := d.Get("org_org_dn").(string)

	lstorageProfileAttr := models.LstorageProfileAttributes{}

	if Assigned_to_dn, ok := d.GetOk("assigned_to_dn"); ok {
		lstorageProfileAttr.Assigned_to_dn = Assigned_to_dn.(string)
	}

	if Availability, ok := d.GetOk("availability"); ok {
		lstorageProfileAttr.Availability = Availability.(string)
	}

	if Child_action, ok := d.GetOk("child_action"); ok {
		lstorageProfileAttr.Child_action = Child_action.(string)
	}

	if Int_id, ok := d.GetOk("int_id"); ok {
		lstorageProfileAttr.Int_id = Int_id.(string)
	}

	if Policy_level, ok := d.GetOk("policy_level"); ok {
		lstorageProfileAttr.Policy_level = Policy_level.(string)
	}

	if Policy_owner, ok := d.GetOk("policy_owner"); ok {
		lstorageProfileAttr.Policy_owner = Policy_owner.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		lstorageProfileAttr.Sacl = Sacl.(string)
	}

	lstorageProfile := models.NewLstorageProfile(fmt.Sprintf("profile-%s", Name), OrgOrg, desc, lstorageProfileAttr)

	err := ucsClient.Save(lstorageProfile)
	if err != nil {
		return err
	}

	d.SetId(lstorageProfile.DistinguishedName)
	return resourceUcsLstorageProfileRead(d, m)
}

func resourceUcsLstorageProfileRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	lstorageProfile, err := getRemoteLstorageProfile(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setLstorageProfileAttributes(lstorageProfile, d)

	return nil
}

func resourceUcsLstorageProfileDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "lstorageProfile")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsLstorageProfileUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	OrgOrg := d.Get("org_org_dn").(string)

	lstorageProfileAttr := models.LstorageProfileAttributes{}
	if Assigned_to_dn, ok := d.GetOk("assigned_to_dn"); ok {
		lstorageProfileAttr.Assigned_to_dn = Assigned_to_dn.(string)
	}
	if Availability, ok := d.GetOk("availability"); ok {
		lstorageProfileAttr.Availability = Availability.(string)
	}
	if Child_action, ok := d.GetOk("child_action"); ok {
		lstorageProfileAttr.Child_action = Child_action.(string)
	}
	if Int_id, ok := d.GetOk("int_id"); ok {
		lstorageProfileAttr.Int_id = Int_id.(string)
	}
	if Policy_level, ok := d.GetOk("policy_level"); ok {
		lstorageProfileAttr.Policy_level = Policy_level.(string)
	}
	if Policy_owner, ok := d.GetOk("policy_owner"); ok {
		lstorageProfileAttr.Policy_owner = Policy_owner.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		lstorageProfileAttr.Sacl = Sacl.(string)
	}

	lstorageProfile := models.NewLstorageProfile(fmt.Sprintf("profile-%s", Name), OrgOrg, desc, lstorageProfileAttr)
	lstorageProfile.Status = "modified"
	err := ucsClient.Save(lstorageProfile)
	if err != nil {
		return err
	}

	d.SetId(lstorageProfile.DistinguishedName)
	return resourceUcsLstorageProfileRead(d, m)
}
