package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsIppoolPool() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsIppoolPoolCreate,
		Update: resourceUcsIppoolPoolUpdate,
		Read:   resourceUcsIppoolPoolRead,
		Delete: resourceUcsIppoolPoolDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsIppoolPoolImport,
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

			"ext_managed": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"guid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"int_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"is_net_bios_enabled": &schema.Schema{
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

			"size": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"supports_dhcp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"v4_assigned": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"v4_size": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"v6_assigned": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"v6_size": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		}),
	}
}

func getRemoteIppoolPool(client *client.Client, dn string) (*models.IppoolPool, error) {
	ippoolPoolDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	ippoolPool := models.IppoolPoolFromDoc(ippoolPoolDoc, "configResolveDn")

	if ippoolPool.DistinguishedName == "" {
		return nil, fmt.Errorf("IppoolPool %s not found", dn)
	}

	return ippoolPool, nil
}

func setIppoolPoolAttributes(ippoolPool *models.IppoolPool, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(ippoolPool.DistinguishedName)
	d.Set("description", ippoolPool.Description)
	d.Set("org_org_dn", GetParentDn(ippoolPool.DistinguishedName))
	ippoolPoolMap, _ := ippoolPool.ToMap()

	d.Set("assigned", ippoolPoolMap["assigned"])

	d.Set("assignment_order", ippoolPoolMap["assignmentOrder"])

	d.Set("child_action", ippoolPoolMap["childAction"])

	d.Set("ext_managed", ippoolPoolMap["extManaged"])

	d.Set("guid", ippoolPoolMap["guid"])

	d.Set("int_id", ippoolPoolMap["intId"])

	d.Set("is_net_bios_enabled", ippoolPoolMap["isNetBIOSEnabled"])

	d.Set("policy_level", ippoolPoolMap["policyLevel"])

	d.Set("policy_owner", ippoolPoolMap["policyOwner"])

	d.Set("prop_acl", ippoolPoolMap["propAcl"])

	d.Set("sacl", ippoolPoolMap["sacl"])

	d.Set("size", ippoolPoolMap["size"])

	d.Set("supports_dhcp", ippoolPoolMap["supportsDHCP"])

	d.Set("v4_assigned", ippoolPoolMap["v4Assigned"])

	d.Set("v4_size", ippoolPoolMap["v4Size"])

	d.Set("v6_assigned", ippoolPoolMap["v6Assigned"])

	d.Set("v6_size", ippoolPoolMap["v6Size"])
	return d
}

func resourceUcsIppoolPoolImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	ippoolPool, err := getRemoteIppoolPool(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setIppoolPoolAttributes(ippoolPool, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsIppoolPoolCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	OrgOrg := d.Get("org_org_dn").(string)

	ippoolPoolAttr := models.IppoolPoolAttributes{}

	if Assigned, ok := d.GetOk("assigned"); ok {
		ippoolPoolAttr.Assigned = Assigned.(string)
	}

	if Assignment_order, ok := d.GetOk("assignment_order"); ok {
		ippoolPoolAttr.Assignment_order = Assignment_order.(string)
	}

	if Child_action, ok := d.GetOk("child_action"); ok {
		ippoolPoolAttr.Child_action = Child_action.(string)
	}

	if Ext_managed, ok := d.GetOk("ext_managed"); ok {
		ippoolPoolAttr.Ext_managed = Ext_managed.(string)
	}

	if Guid, ok := d.GetOk("guid"); ok {
		ippoolPoolAttr.Guid = Guid.(string)
	}

	if Int_id, ok := d.GetOk("int_id"); ok {
		ippoolPoolAttr.Int_id = Int_id.(string)
	}

	if Is_net_bios_enabled, ok := d.GetOk("is_net_bios_enabled"); ok {
		ippoolPoolAttr.Is_net_bios_enabled = Is_net_bios_enabled.(string)
	}

	if Policy_level, ok := d.GetOk("policy_level"); ok {
		ippoolPoolAttr.Policy_level = Policy_level.(string)
	}

	if Policy_owner, ok := d.GetOk("policy_owner"); ok {
		ippoolPoolAttr.Policy_owner = Policy_owner.(string)
	}

	if Prop_acl, ok := d.GetOk("prop_acl"); ok {
		ippoolPoolAttr.Prop_acl = Prop_acl.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		ippoolPoolAttr.Sacl = Sacl.(string)
	}

	if Size, ok := d.GetOk("size"); ok {
		ippoolPoolAttr.Size = Size.(string)
	}

	if Supports_dhcp, ok := d.GetOk("supports_dhcp"); ok {
		ippoolPoolAttr.Supports_dhcp = Supports_dhcp.(string)
	}

	if V4_assigned, ok := d.GetOk("v4_assigned"); ok {
		ippoolPoolAttr.V4_assigned = V4_assigned.(string)
	}

	if V4_size, ok := d.GetOk("v4_size"); ok {
		ippoolPoolAttr.V4_size = V4_size.(string)
	}

	if V6_assigned, ok := d.GetOk("v6_assigned"); ok {
		ippoolPoolAttr.V6_assigned = V6_assigned.(string)
	}

	if V6_size, ok := d.GetOk("v6_size"); ok {
		ippoolPoolAttr.V6_size = V6_size.(string)
	}

	ippoolPool := models.NewIppoolPool(fmt.Sprintf("ip-pool-%s", Name), OrgOrg, desc, ippoolPoolAttr)

	err := ucsClient.Save(ippoolPool)
	if err != nil {
		return err
	}

	d.SetId(ippoolPool.DistinguishedName)
	return resourceUcsIppoolPoolRead(d, m)
}

func resourceUcsIppoolPoolRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	ippoolPool, err := getRemoteIppoolPool(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setIppoolPoolAttributes(ippoolPool, d)

	return nil
}

func resourceUcsIppoolPoolDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "ippoolPool")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsIppoolPoolUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	OrgOrg := d.Get("org_org_dn").(string)

	ippoolPoolAttr := models.IppoolPoolAttributes{}
	if Assigned, ok := d.GetOk("assigned"); ok {
		ippoolPoolAttr.Assigned = Assigned.(string)
	}
	if Assignment_order, ok := d.GetOk("assignment_order"); ok {
		ippoolPoolAttr.Assignment_order = Assignment_order.(string)
	}
	if Child_action, ok := d.GetOk("child_action"); ok {
		ippoolPoolAttr.Child_action = Child_action.(string)
	}
	if Ext_managed, ok := d.GetOk("ext_managed"); ok {
		ippoolPoolAttr.Ext_managed = Ext_managed.(string)
	}
	if Guid, ok := d.GetOk("guid"); ok {
		ippoolPoolAttr.Guid = Guid.(string)
	}
	if Int_id, ok := d.GetOk("int_id"); ok {
		ippoolPoolAttr.Int_id = Int_id.(string)
	}
	if Is_net_bios_enabled, ok := d.GetOk("is_net_bios_enabled"); ok {
		ippoolPoolAttr.Is_net_bios_enabled = Is_net_bios_enabled.(string)
	}
	if Policy_level, ok := d.GetOk("policy_level"); ok {
		ippoolPoolAttr.Policy_level = Policy_level.(string)
	}
	if Policy_owner, ok := d.GetOk("policy_owner"); ok {
		ippoolPoolAttr.Policy_owner = Policy_owner.(string)
	}
	if Prop_acl, ok := d.GetOk("prop_acl"); ok {
		ippoolPoolAttr.Prop_acl = Prop_acl.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		ippoolPoolAttr.Sacl = Sacl.(string)
	}
	if Size, ok := d.GetOk("size"); ok {
		ippoolPoolAttr.Size = Size.(string)
	}
	if Supports_dhcp, ok := d.GetOk("supports_dhcp"); ok {
		ippoolPoolAttr.Supports_dhcp = Supports_dhcp.(string)
	}
	if V4_assigned, ok := d.GetOk("v4_assigned"); ok {
		ippoolPoolAttr.V4_assigned = V4_assigned.(string)
	}
	if V4_size, ok := d.GetOk("v4_size"); ok {
		ippoolPoolAttr.V4_size = V4_size.(string)
	}
	if V6_assigned, ok := d.GetOk("v6_assigned"); ok {
		ippoolPoolAttr.V6_assigned = V6_assigned.(string)
	}
	if V6_size, ok := d.GetOk("v6_size"); ok {
		ippoolPoolAttr.V6_size = V6_size.(string)
	}

	ippoolPool := models.NewIppoolPool(fmt.Sprintf("ip-pool-%s", Name), OrgOrg, desc, ippoolPoolAttr)
	ippoolPool.Status = "modified"
	err := ucsClient.Save(ippoolPool)
	if err != nil {
		return err
	}

	d.SetId(ippoolPool.DistinguishedName)
	return resourceUcsIppoolPoolRead(d, m)
}
