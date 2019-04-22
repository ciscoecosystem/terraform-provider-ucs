package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsVnicSanConnTempl() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsVnicSanConnTemplCreate,
		Update: resourceUcsVnicSanConnTemplUpdate,
		Read:   resourceUcsVnicSanConnTemplRead,
		Delete: resourceUcsVnicSanConnTemplDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsVnicSanConnTemplImport,
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

			"ident_pool_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"int_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"max_data_field_size": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"nw_ctrl_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_ident_pool_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_peer_redundancy_templ_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_qos_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_stats_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"peer_redundancy_templ_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"pin_to_group_name": &schema.Schema{
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

			"qos_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"redundancy_pair_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"sacl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"stats_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"switch_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"target": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"templ_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		}),
	}
}

func getRemoteVnicSanConnTempl(client *client.Client, dn string) (*models.VnicSanConnTempl, error) {
	vnicSanConnTemplDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	vnicSanConnTempl := models.VnicSanConnTemplFromDoc(vnicSanConnTemplDoc, "configResolveDn")

	if vnicSanConnTempl.DistinguishedName == "" {
		return nil, fmt.Errorf("VnicSanConnTempl %s not found", dn)
	}

	return vnicSanConnTempl, nil
}

func setVnicSanConnTemplAttributes(vnicSanConnTempl *models.VnicSanConnTempl, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(vnicSanConnTempl.DistinguishedName)
	d.Set("description", vnicSanConnTempl.Description)
	d.Set("org_org_dn", GetParentDn(vnicSanConnTempl.DistinguishedName))
	vnicSanConnTemplMap, _ := vnicSanConnTempl.ToMap()

	d.Set("child_action", vnicSanConnTemplMap["childAction"])

	d.Set("ident_pool_name", vnicSanConnTemplMap["identPoolName"])

	d.Set("int_id", vnicSanConnTemplMap["intId"])

	d.Set("max_data_field_size", vnicSanConnTemplMap["maxDataFieldSize"])

	d.Set("nw_ctrl_policy_name", vnicSanConnTemplMap["nwCtrlPolicyName"])

	d.Set("oper_ident_pool_name", vnicSanConnTemplMap["operIdentPoolName"])

	d.Set("oper_peer_redundancy_templ_name", vnicSanConnTemplMap["operPeerRedundancyTemplName"])

	d.Set("oper_qos_policy_name", vnicSanConnTemplMap["operQosPolicyName"])

	d.Set("oper_stats_policy_name", vnicSanConnTemplMap["operStatsPolicyName"])

	d.Set("peer_redundancy_templ_name", vnicSanConnTemplMap["peerRedundancyTemplName"])

	d.Set("pin_to_group_name", vnicSanConnTemplMap["pinToGroupName"])

	d.Set("policy_level", vnicSanConnTemplMap["policyLevel"])

	d.Set("policy_owner", vnicSanConnTemplMap["policyOwner"])

	d.Set("qos_policy_name", vnicSanConnTemplMap["qosPolicyName"])

	d.Set("redundancy_pair_type", vnicSanConnTemplMap["redundancyPairType"])

	d.Set("sacl", vnicSanConnTemplMap["sacl"])

	d.Set("stats_policy_name", vnicSanConnTemplMap["statsPolicyName"])

	d.Set("switch_id", vnicSanConnTemplMap["switchId"])

	d.Set("target", vnicSanConnTemplMap["target"])

	d.Set("templ_type", vnicSanConnTemplMap["templType"])
	return d
}

func resourceUcsVnicSanConnTemplImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	vnicSanConnTempl, err := getRemoteVnicSanConnTempl(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setVnicSanConnTemplAttributes(vnicSanConnTempl, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsVnicSanConnTemplCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	OrgOrg := d.Get("org_org_dn").(string)

	vnicSanConnTemplAttr := models.VnicSanConnTemplAttributes{}

	if Child_action, ok := d.GetOk("child_action"); ok {
		vnicSanConnTemplAttr.Child_action = Child_action.(string)
	}

	if Ident_pool_name, ok := d.GetOk("ident_pool_name"); ok {
		vnicSanConnTemplAttr.Ident_pool_name = Ident_pool_name.(string)
	}

	if Int_id, ok := d.GetOk("int_id"); ok {
		vnicSanConnTemplAttr.Int_id = Int_id.(string)
	}

	if Max_data_field_size, ok := d.GetOk("max_data_field_size"); ok {
		vnicSanConnTemplAttr.Max_data_field_size = Max_data_field_size.(string)
	}

	if Nw_ctrl_policy_name, ok := d.GetOk("nw_ctrl_policy_name"); ok {
		vnicSanConnTemplAttr.Nw_ctrl_policy_name = Nw_ctrl_policy_name.(string)
	}

	if Oper_ident_pool_name, ok := d.GetOk("oper_ident_pool_name"); ok {
		vnicSanConnTemplAttr.Oper_ident_pool_name = Oper_ident_pool_name.(string)
	}

	if Oper_peer_redundancy_templ_name, ok := d.GetOk("oper_peer_redundancy_templ_name"); ok {
		vnicSanConnTemplAttr.Oper_peer_redundancy_templ_name = Oper_peer_redundancy_templ_name.(string)
	}

	if Oper_qos_policy_name, ok := d.GetOk("oper_qos_policy_name"); ok {
		vnicSanConnTemplAttr.Oper_qos_policy_name = Oper_qos_policy_name.(string)
	}

	if Oper_stats_policy_name, ok := d.GetOk("oper_stats_policy_name"); ok {
		vnicSanConnTemplAttr.Oper_stats_policy_name = Oper_stats_policy_name.(string)
	}

	if Peer_redundancy_templ_name, ok := d.GetOk("peer_redundancy_templ_name"); ok {
		vnicSanConnTemplAttr.Peer_redundancy_templ_name = Peer_redundancy_templ_name.(string)
	}

	if Pin_to_group_name, ok := d.GetOk("pin_to_group_name"); ok {
		vnicSanConnTemplAttr.Pin_to_group_name = Pin_to_group_name.(string)
	}

	if Policy_level, ok := d.GetOk("policy_level"); ok {
		vnicSanConnTemplAttr.Policy_level = Policy_level.(string)
	}

	if Policy_owner, ok := d.GetOk("policy_owner"); ok {
		vnicSanConnTemplAttr.Policy_owner = Policy_owner.(string)
	}

	if Qos_policy_name, ok := d.GetOk("qos_policy_name"); ok {
		vnicSanConnTemplAttr.Qos_policy_name = Qos_policy_name.(string)
	}

	if Redundancy_pair_type, ok := d.GetOk("redundancy_pair_type"); ok {
		vnicSanConnTemplAttr.Redundancy_pair_type = Redundancy_pair_type.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		vnicSanConnTemplAttr.Sacl = Sacl.(string)
	}

	if Stats_policy_name, ok := d.GetOk("stats_policy_name"); ok {
		vnicSanConnTemplAttr.Stats_policy_name = Stats_policy_name.(string)
	}

	if Switch_id, ok := d.GetOk("switch_id"); ok {
		vnicSanConnTemplAttr.Switch_id = Switch_id.(string)
	}

	if Target, ok := d.GetOk("target"); ok {
		vnicSanConnTemplAttr.Target = Target.(string)
	}

	if Templ_type, ok := d.GetOk("templ_type"); ok {
		vnicSanConnTemplAttr.Templ_type = Templ_type.(string)
	}

	vnicSanConnTempl := models.NewVnicSanConnTempl(fmt.Sprintf("san-conn-templ-%s", Name), OrgOrg, desc, vnicSanConnTemplAttr)

	err := ucsClient.Save(vnicSanConnTempl)
	if err != nil {
		return err
	}

	d.SetId(vnicSanConnTempl.DistinguishedName)
	return resourceUcsVnicSanConnTemplRead(d, m)
}

func resourceUcsVnicSanConnTemplRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	vnicSanConnTempl, err := getRemoteVnicSanConnTempl(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setVnicSanConnTemplAttributes(vnicSanConnTempl, d)

	return nil
}

func resourceUcsVnicSanConnTemplDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "vnicSanConnTempl")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsVnicSanConnTemplUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	OrgOrg := d.Get("org_org_dn").(string)

	vnicSanConnTemplAttr := models.VnicSanConnTemplAttributes{}
	if Child_action, ok := d.GetOk("child_action"); ok {
		vnicSanConnTemplAttr.Child_action = Child_action.(string)
	}
	if Ident_pool_name, ok := d.GetOk("ident_pool_name"); ok {
		vnicSanConnTemplAttr.Ident_pool_name = Ident_pool_name.(string)
	}
	if Int_id, ok := d.GetOk("int_id"); ok {
		vnicSanConnTemplAttr.Int_id = Int_id.(string)
	}
	if Max_data_field_size, ok := d.GetOk("max_data_field_size"); ok {
		vnicSanConnTemplAttr.Max_data_field_size = Max_data_field_size.(string)
	}
	if Nw_ctrl_policy_name, ok := d.GetOk("nw_ctrl_policy_name"); ok {
		vnicSanConnTemplAttr.Nw_ctrl_policy_name = Nw_ctrl_policy_name.(string)
	}
	if Oper_ident_pool_name, ok := d.GetOk("oper_ident_pool_name"); ok {
		vnicSanConnTemplAttr.Oper_ident_pool_name = Oper_ident_pool_name.(string)
	}
	if Oper_peer_redundancy_templ_name, ok := d.GetOk("oper_peer_redundancy_templ_name"); ok {
		vnicSanConnTemplAttr.Oper_peer_redundancy_templ_name = Oper_peer_redundancy_templ_name.(string)
	}
	if Oper_qos_policy_name, ok := d.GetOk("oper_qos_policy_name"); ok {
		vnicSanConnTemplAttr.Oper_qos_policy_name = Oper_qos_policy_name.(string)
	}
	if Oper_stats_policy_name, ok := d.GetOk("oper_stats_policy_name"); ok {
		vnicSanConnTemplAttr.Oper_stats_policy_name = Oper_stats_policy_name.(string)
	}
	if Peer_redundancy_templ_name, ok := d.GetOk("peer_redundancy_templ_name"); ok {
		vnicSanConnTemplAttr.Peer_redundancy_templ_name = Peer_redundancy_templ_name.(string)
	}
	if Pin_to_group_name, ok := d.GetOk("pin_to_group_name"); ok {
		vnicSanConnTemplAttr.Pin_to_group_name = Pin_to_group_name.(string)
	}
	if Policy_level, ok := d.GetOk("policy_level"); ok {
		vnicSanConnTemplAttr.Policy_level = Policy_level.(string)
	}
	if Policy_owner, ok := d.GetOk("policy_owner"); ok {
		vnicSanConnTemplAttr.Policy_owner = Policy_owner.(string)
	}
	if Qos_policy_name, ok := d.GetOk("qos_policy_name"); ok {
		vnicSanConnTemplAttr.Qos_policy_name = Qos_policy_name.(string)
	}
	if Redundancy_pair_type, ok := d.GetOk("redundancy_pair_type"); ok {
		vnicSanConnTemplAttr.Redundancy_pair_type = Redundancy_pair_type.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		vnicSanConnTemplAttr.Sacl = Sacl.(string)
	}
	if Stats_policy_name, ok := d.GetOk("stats_policy_name"); ok {
		vnicSanConnTemplAttr.Stats_policy_name = Stats_policy_name.(string)
	}
	if Switch_id, ok := d.GetOk("switch_id"); ok {
		vnicSanConnTemplAttr.Switch_id = Switch_id.(string)
	}
	if Target, ok := d.GetOk("target"); ok {
		vnicSanConnTemplAttr.Target = Target.(string)
	}
	if Templ_type, ok := d.GetOk("templ_type"); ok {
		vnicSanConnTemplAttr.Templ_type = Templ_type.(string)
	}

	vnicSanConnTempl := models.NewVnicSanConnTempl(fmt.Sprintf("san-conn-templ-%s", Name), OrgOrg, desc, vnicSanConnTemplAttr)
	vnicSanConnTempl.Status = "modified"
	err := ucsClient.Save(vnicSanConnTempl)
	if err != nil {
		return err
	}

	d.SetId(vnicSanConnTempl.DistinguishedName)
	return resourceUcsVnicSanConnTemplRead(d, m)
}
