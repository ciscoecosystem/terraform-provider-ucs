package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsLsRequirement() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsLsRequirementCreate,
		Update: resourceUcsLsRequirementUpdate,
		Read:   resourceUcsLsRequirementRead,
		Delete: resourceUcsLsRequirementDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsLsRequirementImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"ls_server_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"admin_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"admin_action_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"assigned_to_dn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"child_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"compute_ep_dn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"issues": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"pn_dn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"pn_pool_dn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"qualifier": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"restrict_migration": &schema.Schema{
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

func getRemoteLsRequirement(client *client.Client, dn string) (*models.LsRequirement, error) {
	lsRequirementDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	lsRequirement := models.LsRequirementFromDoc(lsRequirementDoc, "configResolveDn")

	if lsRequirement.DistinguishedName == "" {
		return nil, fmt.Errorf("LsRequirement %s not found", dn)
	}

	return lsRequirement, nil
}

func setLsRequirementAttributes(lsRequirement *models.LsRequirement, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(lsRequirement.DistinguishedName)
	d.Set("description", lsRequirement.Description)
	d.Set("ls_server_dn", GetParentDn(lsRequirement.DistinguishedName))
	lsRequirementMap, _ := lsRequirement.ToMap()

	d.Set("admin_action", lsRequirementMap["adminAction"])

	d.Set("admin_action_trigger", lsRequirementMap["adminActionTrigger"])

	d.Set("assigned_to_dn", lsRequirementMap["assignedToDn"])

	d.Set("child_action", lsRequirementMap["childAction"])

	d.Set("compute_ep_dn", lsRequirementMap["computeEpDn"])

	d.Set("issues", lsRequirementMap["issues"])

	d.Set("name", lsRequirementMap["name"])

	d.Set("oper_name", lsRequirementMap["operName"])

	d.Set("oper_state", lsRequirementMap["operState"])

	d.Set("pn_dn", lsRequirementMap["pnDn"])

	d.Set("pn_pool_dn", lsRequirementMap["pnPoolDn"])

	d.Set("qualifier", lsRequirementMap["qualifier"])

	d.Set("restrict_migration", lsRequirementMap["restrictMigration"])

	d.Set("sacl", lsRequirementMap["sacl"])
	return d
}

func resourceUcsLsRequirementImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	lsRequirement, err := getRemoteLsRequirement(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setLsRequirementAttributes(lsRequirement, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsLsRequirementCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	LsServer := d.Get("ls_server_dn").(string)

	lsRequirementAttr := models.LsRequirementAttributes{}

	if Admin_action, ok := d.GetOk("admin_action"); ok {
		lsRequirementAttr.Admin_action = Admin_action.(string)
	}

	if Admin_action_trigger, ok := d.GetOk("admin_action_trigger"); ok {
		lsRequirementAttr.Admin_action_trigger = Admin_action_trigger.(string)
	}

	if Assigned_to_dn, ok := d.GetOk("assigned_to_dn"); ok {
		lsRequirementAttr.Assigned_to_dn = Assigned_to_dn.(string)
	}

	if Child_action, ok := d.GetOk("child_action"); ok {
		lsRequirementAttr.Child_action = Child_action.(string)
	}

	if Compute_ep_dn, ok := d.GetOk("compute_ep_dn"); ok {
		lsRequirementAttr.Compute_ep_dn = Compute_ep_dn.(string)
	}

	if Issues, ok := d.GetOk("issues"); ok {
		lsRequirementAttr.Issues = Issues.(string)
	}

	if Oper_name, ok := d.GetOk("oper_name"); ok {
		lsRequirementAttr.Oper_name = Oper_name.(string)
	}

	if Oper_state, ok := d.GetOk("oper_state"); ok {
		lsRequirementAttr.Oper_state = Oper_state.(string)
	}

	if Pn_dn, ok := d.GetOk("pn_dn"); ok {
		lsRequirementAttr.Pn_dn = Pn_dn.(string)
	}

	if Pn_pool_dn, ok := d.GetOk("pn_pool_dn"); ok {
		lsRequirementAttr.Pn_pool_dn = Pn_pool_dn.(string)
	}

	if Qualifier, ok := d.GetOk("qualifier"); ok {
		lsRequirementAttr.Qualifier = Qualifier.(string)
	}

	if Restrict_migration, ok := d.GetOk("restrict_migration"); ok {
		lsRequirementAttr.Restrict_migration = Restrict_migration.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		lsRequirementAttr.Sacl = Sacl.(string)
	}

	lsRequirement := models.NewLsRequirement(fmt.Sprintf("pn-req"), LsServer, desc, lsRequirementAttr)

	err := ucsClient.Save(lsRequirement)
	if err != nil {
		return err
	}

	d.SetId(lsRequirement.DistinguishedName)
	return resourceUcsLsRequirementRead(d, m)
}

func resourceUcsLsRequirementRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	lsRequirement, err := getRemoteLsRequirement(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setLsRequirementAttributes(lsRequirement, d)

	return nil
}

func resourceUcsLsRequirementDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "lsRequirement")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsLsRequirementUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	LsServer := d.Get("ls_server_dn").(string)

	lsRequirementAttr := models.LsRequirementAttributes{}
	if Admin_action, ok := d.GetOk("admin_action"); ok {
		lsRequirementAttr.Admin_action = Admin_action.(string)
	}
	if Admin_action_trigger, ok := d.GetOk("admin_action_trigger"); ok {
		lsRequirementAttr.Admin_action_trigger = Admin_action_trigger.(string)
	}
	if Assigned_to_dn, ok := d.GetOk("assigned_to_dn"); ok {
		lsRequirementAttr.Assigned_to_dn = Assigned_to_dn.(string)
	}
	if Child_action, ok := d.GetOk("child_action"); ok {
		lsRequirementAttr.Child_action = Child_action.(string)
	}
	if Compute_ep_dn, ok := d.GetOk("compute_ep_dn"); ok {
		lsRequirementAttr.Compute_ep_dn = Compute_ep_dn.(string)
	}
	if Issues, ok := d.GetOk("issues"); ok {
		lsRequirementAttr.Issues = Issues.(string)
	}
	if Oper_name, ok := d.GetOk("oper_name"); ok {
		lsRequirementAttr.Oper_name = Oper_name.(string)
	}
	if Oper_state, ok := d.GetOk("oper_state"); ok {
		lsRequirementAttr.Oper_state = Oper_state.(string)
	}
	if Pn_dn, ok := d.GetOk("pn_dn"); ok {
		lsRequirementAttr.Pn_dn = Pn_dn.(string)
	}
	if Pn_pool_dn, ok := d.GetOk("pn_pool_dn"); ok {
		lsRequirementAttr.Pn_pool_dn = Pn_pool_dn.(string)
	}
	if Qualifier, ok := d.GetOk("qualifier"); ok {
		lsRequirementAttr.Qualifier = Qualifier.(string)
	}
	if Restrict_migration, ok := d.GetOk("restrict_migration"); ok {
		lsRequirementAttr.Restrict_migration = Restrict_migration.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		lsRequirementAttr.Sacl = Sacl.(string)
	}

	lsRequirement := models.NewLsRequirement(fmt.Sprintf("pn-req"), LsServer, desc, lsRequirementAttr)
	lsRequirement.Status = "modified"
	err := ucsClient.Save(lsRequirement)
	if err != nil {
		return err
	}

	d.SetId(lsRequirement.DistinguishedName)
	return resourceUcsLsRequirementRead(d, m)
}
