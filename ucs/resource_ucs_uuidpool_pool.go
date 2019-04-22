package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsUuidpoolPool() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsUuidpoolPoolCreate,
		Update: resourceUcsUuidpoolPoolUpdate,
		Read:   resourceUcsUuidpoolPoolRead,
		Delete: resourceUcsUuidpoolPoolDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsUuidpoolPoolImport,
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

			"prefix": &schema.Schema{
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

func getRemoteUuidpoolPool(client *client.Client, dn string) (*models.UuidpoolPool, error) {
	uuidpoolPoolDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	uuidpoolPool := models.UuidpoolPoolFromDoc(uuidpoolPoolDoc, "configResolveDn")

	if uuidpoolPool.DistinguishedName == "" {
		return nil, fmt.Errorf("UuidpoolPool %s not found", dn)
	}

	return uuidpoolPool, nil
}

func setUuidpoolPoolAttributes(uuidpoolPool *models.UuidpoolPool, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(uuidpoolPool.DistinguishedName)
	d.Set("description", uuidpoolPool.Description)
	d.Set("org_org_dn", GetParentDn(uuidpoolPool.DistinguishedName))
	uuidpoolPoolMap, _ := uuidpoolPool.ToMap()

	d.Set("assigned", uuidpoolPoolMap["assigned"])

	d.Set("assignment_order", uuidpoolPoolMap["assignmentOrder"])

	d.Set("child_action", uuidpoolPoolMap["childAction"])

	d.Set("int_id", uuidpoolPoolMap["intId"])

	d.Set("policy_level", uuidpoolPoolMap["policyLevel"])

	d.Set("policy_owner", uuidpoolPoolMap["policyOwner"])

	d.Set("prefix", uuidpoolPoolMap["prefix"])

	d.Set("sacl", uuidpoolPoolMap["sacl"])

	d.Set("size", uuidpoolPoolMap["size"])
	return d
}

func resourceUcsUuidpoolPoolImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	uuidpoolPool, err := getRemoteUuidpoolPool(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setUuidpoolPoolAttributes(uuidpoolPool, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsUuidpoolPoolCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	OrgOrg := d.Get("org_org_dn").(string)

	uuidpoolPoolAttr := models.UuidpoolPoolAttributes{}

	if Assigned, ok := d.GetOk("assigned"); ok {
		uuidpoolPoolAttr.Assigned = Assigned.(string)
	}

	if Assignment_order, ok := d.GetOk("assignment_order"); ok {
		uuidpoolPoolAttr.Assignment_order = Assignment_order.(string)
	}

	if Child_action, ok := d.GetOk("child_action"); ok {
		uuidpoolPoolAttr.Child_action = Child_action.(string)
	}

	if Int_id, ok := d.GetOk("int_id"); ok {
		uuidpoolPoolAttr.Int_id = Int_id.(string)
	}

	if Policy_level, ok := d.GetOk("policy_level"); ok {
		uuidpoolPoolAttr.Policy_level = Policy_level.(string)
	}

	if Policy_owner, ok := d.GetOk("policy_owner"); ok {
		uuidpoolPoolAttr.Policy_owner = Policy_owner.(string)
	}

	if Prefix, ok := d.GetOk("prefix"); ok {
		uuidpoolPoolAttr.Prefix = Prefix.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		uuidpoolPoolAttr.Sacl = Sacl.(string)
	}

	if Size, ok := d.GetOk("size"); ok {
		uuidpoolPoolAttr.Size = Size.(string)
	}

	uuidpoolPool := models.NewUuidpoolPool(fmt.Sprintf("uuid-pool-%s", Name), OrgOrg, desc, uuidpoolPoolAttr)

	err := ucsClient.Save(uuidpoolPool)
	if err != nil {
		return err
	}

	d.SetId(uuidpoolPool.DistinguishedName)
	return resourceUcsUuidpoolPoolRead(d, m)
}

func resourceUcsUuidpoolPoolRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	uuidpoolPool, err := getRemoteUuidpoolPool(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setUuidpoolPoolAttributes(uuidpoolPool, d)

	return nil
}

func resourceUcsUuidpoolPoolDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "uuidpoolPool")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsUuidpoolPoolUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	OrgOrg := d.Get("org_org_dn").(string)

	uuidpoolPoolAttr := models.UuidpoolPoolAttributes{}
	if Assigned, ok := d.GetOk("assigned"); ok {
		uuidpoolPoolAttr.Assigned = Assigned.(string)
	}
	if Assignment_order, ok := d.GetOk("assignment_order"); ok {
		uuidpoolPoolAttr.Assignment_order = Assignment_order.(string)
	}
	if Child_action, ok := d.GetOk("child_action"); ok {
		uuidpoolPoolAttr.Child_action = Child_action.(string)
	}
	if Int_id, ok := d.GetOk("int_id"); ok {
		uuidpoolPoolAttr.Int_id = Int_id.(string)
	}
	if Policy_level, ok := d.GetOk("policy_level"); ok {
		uuidpoolPoolAttr.Policy_level = Policy_level.(string)
	}
	if Policy_owner, ok := d.GetOk("policy_owner"); ok {
		uuidpoolPoolAttr.Policy_owner = Policy_owner.(string)
	}
	if Prefix, ok := d.GetOk("prefix"); ok {
		uuidpoolPoolAttr.Prefix = Prefix.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		uuidpoolPoolAttr.Sacl = Sacl.(string)
	}
	if Size, ok := d.GetOk("size"); ok {
		uuidpoolPoolAttr.Size = Size.(string)
	}

	uuidpoolPool := models.NewUuidpoolPool(fmt.Sprintf("uuid-pool-%s", Name), OrgOrg, desc, uuidpoolPoolAttr)
	uuidpoolPool.Status = "modified"
	err := ucsClient.Save(uuidpoolPool)
	if err != nil {
		return err
	}

	d.SetId(uuidpoolPool.DistinguishedName)
	return resourceUcsUuidpoolPoolRead(d, m)
}
