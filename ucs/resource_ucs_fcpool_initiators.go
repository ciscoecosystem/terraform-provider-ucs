package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsFcpoolInitiators() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsFcpoolInitiatorsCreate,
		Update: resourceUcsFcpoolInitiatorsUpdate,
		Read:   resourceUcsFcpoolInitiatorsRead,
		Delete: resourceUcsFcpoolInitiatorsDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsFcpoolInitiatorsImport,
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

			"max_ports_per_node": &schema.Schema{
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

			"purpose": &schema.Schema{
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

func getRemoteFcpoolInitiators(client *client.Client, dn string) (*models.FcpoolInitiators, error) {
	fcpoolInitiatorsDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	fcpoolInitiators := models.FcpoolInitiatorsFromDoc(fcpoolInitiatorsDoc, "configResolveDn")

	if fcpoolInitiators.DistinguishedName == "" {
		return nil, fmt.Errorf("FcpoolInitiators %s not found", dn)
	}

	return fcpoolInitiators, nil
}

func setFcpoolInitiatorsAttributes(fcpoolInitiators *models.FcpoolInitiators, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(fcpoolInitiators.DistinguishedName)
	d.Set("description", fcpoolInitiators.Description)
	d.Set("org_org_dn", GetParentDn(fcpoolInitiators.DistinguishedName))
	fcpoolInitiatorsMap, _ := fcpoolInitiators.ToMap()

	d.Set("assigned", fcpoolInitiatorsMap["assigned"])

	d.Set("assignment_order", fcpoolInitiatorsMap["assignmentOrder"])

	d.Set("child_action", fcpoolInitiatorsMap["childAction"])

	d.Set("int_id", fcpoolInitiatorsMap["intId"])

	d.Set("max_ports_per_node", fcpoolInitiatorsMap["maxPortsPerNode"])

	d.Set("policy_level", fcpoolInitiatorsMap["policyLevel"])

	d.Set("policy_owner", fcpoolInitiatorsMap["policyOwner"])

	d.Set("purpose", fcpoolInitiatorsMap["purpose"])

	d.Set("sacl", fcpoolInitiatorsMap["sacl"])

	d.Set("size", fcpoolInitiatorsMap["size"])
	return d
}

func resourceUcsFcpoolInitiatorsImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	fcpoolInitiators, err := getRemoteFcpoolInitiators(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setFcpoolInitiatorsAttributes(fcpoolInitiators, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsFcpoolInitiatorsCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	OrgOrg := d.Get("org_org_dn").(string)

	fcpoolInitiatorsAttr := models.FcpoolInitiatorsAttributes{}

	if Assigned, ok := d.GetOk("assigned"); ok {
		fcpoolInitiatorsAttr.Assigned = Assigned.(string)
	}

	if Assignment_order, ok := d.GetOk("assignment_order"); ok {
		fcpoolInitiatorsAttr.Assignment_order = Assignment_order.(string)
	}

	if Child_action, ok := d.GetOk("child_action"); ok {
		fcpoolInitiatorsAttr.Child_action = Child_action.(string)
	}

	if Int_id, ok := d.GetOk("int_id"); ok {
		fcpoolInitiatorsAttr.Int_id = Int_id.(string)
	}

	if Max_ports_per_node, ok := d.GetOk("max_ports_per_node"); ok {
		fcpoolInitiatorsAttr.Max_ports_per_node = Max_ports_per_node.(string)
	}

	if Policy_level, ok := d.GetOk("policy_level"); ok {
		fcpoolInitiatorsAttr.Policy_level = Policy_level.(string)
	}

	if Policy_owner, ok := d.GetOk("policy_owner"); ok {
		fcpoolInitiatorsAttr.Policy_owner = Policy_owner.(string)
	}

	if Purpose, ok := d.GetOk("purpose"); ok {
		fcpoolInitiatorsAttr.Purpose = Purpose.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		fcpoolInitiatorsAttr.Sacl = Sacl.(string)
	}

	if Size, ok := d.GetOk("size"); ok {
		fcpoolInitiatorsAttr.Size = Size.(string)
	}

	fcpoolInitiators := models.NewFcpoolInitiators(fmt.Sprintf("wwn-pool-%s", Name), OrgOrg, desc, fcpoolInitiatorsAttr)

	err := ucsClient.Save(fcpoolInitiators)
	if err != nil {
		return err
	}

	d.SetId(fcpoolInitiators.DistinguishedName)
	return resourceUcsFcpoolInitiatorsRead(d, m)
}

func resourceUcsFcpoolInitiatorsRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	fcpoolInitiators, err := getRemoteFcpoolInitiators(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setFcpoolInitiatorsAttributes(fcpoolInitiators, d)

	return nil
}

func resourceUcsFcpoolInitiatorsDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "fcpoolInitiators")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsFcpoolInitiatorsUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	OrgOrg := d.Get("org_org_dn").(string)

	fcpoolInitiatorsAttr := models.FcpoolInitiatorsAttributes{}
	if Assigned, ok := d.GetOk("assigned"); ok {
		fcpoolInitiatorsAttr.Assigned = Assigned.(string)
	}
	if Assignment_order, ok := d.GetOk("assignment_order"); ok {
		fcpoolInitiatorsAttr.Assignment_order = Assignment_order.(string)
	}
	if Child_action, ok := d.GetOk("child_action"); ok {
		fcpoolInitiatorsAttr.Child_action = Child_action.(string)
	}
	if Int_id, ok := d.GetOk("int_id"); ok {
		fcpoolInitiatorsAttr.Int_id = Int_id.(string)
	}
	if Max_ports_per_node, ok := d.GetOk("max_ports_per_node"); ok {
		fcpoolInitiatorsAttr.Max_ports_per_node = Max_ports_per_node.(string)
	}
	if Policy_level, ok := d.GetOk("policy_level"); ok {
		fcpoolInitiatorsAttr.Policy_level = Policy_level.(string)
	}
	if Policy_owner, ok := d.GetOk("policy_owner"); ok {
		fcpoolInitiatorsAttr.Policy_owner = Policy_owner.(string)
	}
	if Purpose, ok := d.GetOk("purpose"); ok {
		fcpoolInitiatorsAttr.Purpose = Purpose.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		fcpoolInitiatorsAttr.Sacl = Sacl.(string)
	}
	if Size, ok := d.GetOk("size"); ok {
		fcpoolInitiatorsAttr.Size = Size.(string)
	}

	fcpoolInitiators := models.NewFcpoolInitiators(fmt.Sprintf("wwn-pool-%s", Name), OrgOrg, desc, fcpoolInitiatorsAttr)
	fcpoolInitiators.Status = "modified"
	err := ucsClient.Save(fcpoolInitiators)
	if err != nil {
		return err
	}

	d.SetId(fcpoolInitiators.DistinguishedName)
	return resourceUcsFcpoolInitiatorsRead(d, m)
}
