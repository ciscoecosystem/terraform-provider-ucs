package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsMacpoolPool() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsMacpoolPoolCreate,
		Update: resourceUcsMacpoolPoolUpdate,
		Read:   resourceUcsMacpoolPoolRead,
		Delete: resourceUcsMacpoolPoolDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsMacpoolPoolImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"org_org_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"assigned": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"assignment_order": &schema.Schema{
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

			"size": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		}),
	}
}

func getRemoteMacpoolPool(client *client.Client, dn string) (*models.MacpoolPool, error) {
	macpoolPoolDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	macpoolPool := models.MacpoolPoolFromDoc(macpoolPoolDoc, "configResolveDn")

	if macpoolPool.DistinguishedName == "" {
		return nil, fmt.Errorf("MacpoolPool %s not found", dn)
	}

	return macpoolPool, nil
}

func setMacpoolPoolAttributes(macpoolPool *models.MacpoolPool, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(macpoolPool.DistinguishedName)
	d.Set("description", macpoolPool.Description)
	d.Set("org_org_dn", GetParentDn(macpoolPool.DistinguishedName))
	macpoolPoolMap, _ := macpoolPool.ToMap()

	d.Set("assigned", macpoolPoolMap["assigned"])

	d.Set("assignment_order", macpoolPoolMap["assignmentOrder"])

	d.Set("child_action", macpoolPoolMap["childAction"])

	d.Set("int_id", macpoolPoolMap["intId"])

	d.Set("name", macpoolPoolMap["name"])

	d.Set("policy_level", macpoolPoolMap["policyLevel"])

	d.Set("policy_owner", macpoolPoolMap["policyOwner"])

	d.Set("sacl", macpoolPoolMap["sacl"])

	d.Set("size", macpoolPoolMap["size"])
	return d
}

func resourceUcsMacpoolPoolImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	macpoolPool, err := getRemoteMacpoolPool(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setMacpoolPoolAttributes(macpoolPool, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsMacpoolPoolCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)
	Name := d.Get("name").(string)

	OrgOrg := d.Get("org_org_dn").(string)

	macpoolPoolAttr := models.MacpoolPoolAttributes{}

	if Assigned, ok := d.GetOk("assigned"); ok {
		macpoolPoolAttr.Name = Assigned.(string)
	}

	if Assignment_order, ok := d.GetOk("assignment_order"); ok {
		macpoolPoolAttr.Name = Assignment_order.(string)
	}

	if Child_action, ok := d.GetOk("child_action"); ok {
		macpoolPoolAttr.Name = Child_action.(string)
	}

	if Int_id, ok := d.GetOk("int_id"); ok {
		macpoolPoolAttr.Name = Int_id.(string)
	}

	if Policy_level, ok := d.GetOk("policy_level"); ok {
		macpoolPoolAttr.Name = Policy_level.(string)
	}

	if Policy_owner, ok := d.GetOk("policy_owner"); ok {
		macpoolPoolAttr.Name = Policy_owner.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		macpoolPoolAttr.Name = Sacl.(string)
	}

	if Size, ok := d.GetOk("size"); ok {
		macpoolPoolAttr.Name = Size.(string)
	}

	macpoolPool := models.NewMacpoolPool(fmt.Sprintf("mac-pool-%s", Name), OrgOrg, desc, macpoolPoolAttr)

	err := ucsClient.Save(macpoolPool)
	if err != nil {
		return err
	}

	d.SetId(macpoolPool.DistinguishedName)
	return resourceUcsMacpoolPoolRead(d, m)
}

func resourceUcsMacpoolPoolRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	macpoolPool, err := getRemoteMacpoolPool(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setMacpoolPoolAttributes(macpoolPool, d)

	return nil
}

func resourceUcsMacpoolPoolDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "macpoolPool")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsMacpoolPoolUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)
	Name := d.Get("name").(string)

	OrgOrg := d.Get("org_org_dn").(string)

	macpoolPoolAttr := models.MacpoolPoolAttributes{}
	if Assigned, ok := d.GetOk("assigned"); ok {
		macpoolPoolAttr.Name = Assigned.(string)
	}
	if Assignment_order, ok := d.GetOk("assignment_order"); ok {
		macpoolPoolAttr.Name = Assignment_order.(string)
	}
	if Child_action, ok := d.GetOk("child_action"); ok {
		macpoolPoolAttr.Name = Child_action.(string)
	}
	if Int_id, ok := d.GetOk("int_id"); ok {
		macpoolPoolAttr.Name = Int_id.(string)
	}
	if Policy_level, ok := d.GetOk("policy_level"); ok {
		macpoolPoolAttr.Name = Policy_level.(string)
	}
	if Policy_owner, ok := d.GetOk("policy_owner"); ok {
		macpoolPoolAttr.Name = Policy_owner.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		macpoolPoolAttr.Name = Sacl.(string)
	}
	if Size, ok := d.GetOk("size"); ok {
		macpoolPoolAttr.Name = Size.(string)
	}

	macpoolPool := models.NewMacpoolPool(fmt.Sprintf("mac-pool-%s", Name), OrgOrg, desc, macpoolPoolAttr)
	macpoolPool.Status = "modified"
	err := ucsClient.Save(macpoolPool)
	if err != nil {
		return err
	}

	d.SetId(macpoolPool.DistinguishedName)
	return resourceUcsMacpoolPoolRead(d, m)
}
