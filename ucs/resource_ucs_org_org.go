package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsOrgOrg() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsOrgOrgCreate,
		Update: resourceUcsOrgOrgUpdate,
		Read:   resourceUcsOrgOrgRead,
		Delete: resourceUcsOrgOrgDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsOrgOrgImport,
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

			"level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"perm_access": &schema.Schema{
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

func getRemoteOrgOrg(client *client.Client, dn string) (*models.OrgOrg, error) {
	orgOrgDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	orgOrg := models.OrgOrgFromDoc(orgOrgDoc, "configResolveDn")

	if orgOrg.DistinguishedName == "" {
		return nil, fmt.Errorf("OrgOrg %s not found", dn)
	}

	return orgOrg, nil
}

func setOrgOrgAttributes(orgOrg *models.OrgOrg, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(orgOrg.DistinguishedName)
	d.Set("description", orgOrg.Description)
	d.Set("org_org_dn", GetParentDn(orgOrg.DistinguishedName))
	orgOrgMap, _ := orgOrg.ToMap()

	d.Set("child_action", orgOrgMap["childAction"])

	d.Set("flt_aggr", orgOrgMap["fltAggr"])

	d.Set("level", orgOrgMap["level"])

	d.Set("perm_access", orgOrgMap["permAccess"])

	d.Set("sacl", orgOrgMap["sacl"])
	return d
}

func resourceUcsOrgOrgImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	orgOrg, err := getRemoteOrgOrg(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setOrgOrgAttributes(orgOrg, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsOrgOrgCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	OrgOrg := d.Get("org_org_dn").(string)

	orgOrgAttr := models.OrgOrgAttributes{}

	if Child_action, ok := d.GetOk("child_action"); ok {
		orgOrgAttr.Child_action = Child_action.(string)
	}

	if Flt_aggr, ok := d.GetOk("flt_aggr"); ok {
		orgOrgAttr.Flt_aggr = Flt_aggr.(string)
	}

	if Level, ok := d.GetOk("level"); ok {
		orgOrgAttr.Level = Level.(string)
	}

	if Perm_access, ok := d.GetOk("perm_access"); ok {
		orgOrgAttr.Perm_access = Perm_access.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		orgOrgAttr.Sacl = Sacl.(string)
	}

	orgOrg := models.NewOrgOrg(fmt.Sprintf("org-%s", Name), OrgOrg, desc, orgOrgAttr)

	err := ucsClient.Save(orgOrg)
	if err != nil {
		return err
	}

	d.SetId(orgOrg.DistinguishedName)
	return resourceUcsOrgOrgRead(d, m)
}

func resourceUcsOrgOrgRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	orgOrg, err := getRemoteOrgOrg(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setOrgOrgAttributes(orgOrg, d)

	return nil
}

func resourceUcsOrgOrgDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "orgOrg")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsOrgOrgUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	OrgOrg := d.Get("org_org_dn").(string)

	orgOrgAttr := models.OrgOrgAttributes{}
	if Child_action, ok := d.GetOk("child_action"); ok {
		orgOrgAttr.Child_action = Child_action.(string)
	}
	if Flt_aggr, ok := d.GetOk("flt_aggr"); ok {
		orgOrgAttr.Flt_aggr = Flt_aggr.(string)
	}
	if Level, ok := d.GetOk("level"); ok {
		orgOrgAttr.Level = Level.(string)
	}
	if Perm_access, ok := d.GetOk("perm_access"); ok {
		orgOrgAttr.Perm_access = Perm_access.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		orgOrgAttr.Sacl = Sacl.(string)
	}

	orgOrg := models.NewOrgOrg(fmt.Sprintf("org-%s", Name), OrgOrg, desc, orgOrgAttr)
	orgOrg.Status = "modified"
	err := ucsClient.Save(orgOrg)
	if err != nil {
		return err
	}

	d.SetId(orgOrg.DistinguishedName)
	return resourceUcsOrgOrgRead(d, m)
}
