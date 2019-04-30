package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsVnicLanConnTempl() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsVnicLanConnTemplCreate,
		Update: resourceUcsVnicLanConnTemplUpdate,
		Read:   resourceUcsVnicLanConnTemplRead,
		Delete: resourceUcsVnicLanConnTemplDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsVnicLanConnTemplImport,
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

			"admin_cdn_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"cdn_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

			"mtu": &schema.Schema{
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

			"oper_nw_ctrl_policy_name": &schema.Schema{
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

func getRemoteVnicLanConnTempl(client *client.Client, dn string) (*models.VnicLanConnTempl, error) {
	vnicLanConnTemplDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	vnicLanConnTempl := models.VnicLanConnTemplFromDoc(vnicLanConnTemplDoc, "configResolveDn")

	if vnicLanConnTempl.DistinguishedName == "" {
		return nil, fmt.Errorf("VnicLanConnTempl %s not found", dn)
	}

	return vnicLanConnTempl, nil
}

func setVnicLanConnTemplAttributes(vnicLanConnTempl *models.VnicLanConnTempl, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(vnicLanConnTempl.DistinguishedName)
	d.Set("description", vnicLanConnTempl.Description)
	d.Set("org_org_dn", GetParentDn(vnicLanConnTempl.DistinguishedName))
	vnicLanConnTemplMap, _ := vnicLanConnTempl.ToMap()

	d.Set("admin_cdn_name", vnicLanConnTemplMap["adminCdnName"])

	d.Set("cdn_source", vnicLanConnTemplMap["cdnSource"])

	d.Set("child_action", vnicLanConnTemplMap["childAction"])

	d.Set("ident_pool_name", vnicLanConnTemplMap["identPoolName"])

	d.Set("int_id", vnicLanConnTemplMap["intId"])

	d.Set("mtu", vnicLanConnTemplMap["mtu"])

	d.Set("nw_ctrl_policy_name", vnicLanConnTemplMap["nwCtrlPolicyName"])

	d.Set("oper_ident_pool_name", vnicLanConnTemplMap["operIdentPoolName"])

	d.Set("oper_nw_ctrl_policy_name", vnicLanConnTemplMap["operNwCtrlPolicyName"])

	d.Set("oper_peer_redundancy_templ_name", vnicLanConnTemplMap["operPeerRedundancyTemplName"])

	d.Set("oper_qos_policy_name", vnicLanConnTemplMap["operQosPolicyName"])

	d.Set("oper_stats_policy_name", vnicLanConnTemplMap["operStatsPolicyName"])

	d.Set("peer_redundancy_templ_name", vnicLanConnTemplMap["peerRedundancyTemplName"])

	d.Set("pin_to_group_name", vnicLanConnTemplMap["pinToGroupName"])

	d.Set("policy_level", vnicLanConnTemplMap["policyLevel"])

	d.Set("policy_owner", vnicLanConnTemplMap["policyOwner"])

	d.Set("qos_policy_name", vnicLanConnTemplMap["qosPolicyName"])

	d.Set("redundancy_pair_type", vnicLanConnTemplMap["redundancyPairType"])

	d.Set("sacl", vnicLanConnTemplMap["sacl"])

	d.Set("stats_policy_name", vnicLanConnTemplMap["statsPolicyName"])

	d.Set("switch_id", vnicLanConnTemplMap["switchId"])

	d.Set("target", vnicLanConnTemplMap["target"])

	d.Set("templ_type", vnicLanConnTemplMap["templType"])
	return d
}

func resourceUcsVnicLanConnTemplImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	vnicLanConnTempl, err := getRemoteVnicLanConnTempl(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setVnicLanConnTemplAttributes(vnicLanConnTempl, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsVnicLanConnTemplCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	OrgOrg := d.Get("org_org_dn").(string)

	vnicLanConnTemplAttr := models.VnicLanConnTemplAttributes{}

	if Admin_cdn_name, ok := d.GetOk("admin_cdn_name"); ok {
		vnicLanConnTemplAttr.Admin_cdn_name = Admin_cdn_name.(string)
	}

	if Cdn_source, ok := d.GetOk("cdn_source"); ok {
		vnicLanConnTemplAttr.Cdn_source = Cdn_source.(string)
	}

	if Child_action, ok := d.GetOk("child_action"); ok {
		vnicLanConnTemplAttr.Child_action = Child_action.(string)
	}

	if Ident_pool_name, ok := d.GetOk("ident_pool_name"); ok {
		vnicLanConnTemplAttr.Ident_pool_name = Ident_pool_name.(string)
	}

	if Int_id, ok := d.GetOk("int_id"); ok {
		vnicLanConnTemplAttr.Int_id = Int_id.(string)
	}

	if Mtu, ok := d.GetOk("mtu"); ok {
		vnicLanConnTemplAttr.Mtu = Mtu.(string)
	}

	if Nw_ctrl_policy_name, ok := d.GetOk("nw_ctrl_policy_name"); ok {
		vnicLanConnTemplAttr.Nw_ctrl_policy_name = Nw_ctrl_policy_name.(string)
	}

	if Oper_ident_pool_name, ok := d.GetOk("oper_ident_pool_name"); ok {
		vnicLanConnTemplAttr.Oper_ident_pool_name = Oper_ident_pool_name.(string)
	}

	if Oper_nw_ctrl_policy_name, ok := d.GetOk("oper_nw_ctrl_policy_name"); ok {
		vnicLanConnTemplAttr.Oper_nw_ctrl_policy_name = Oper_nw_ctrl_policy_name.(string)
	}

	if Oper_peer_redundancy_templ_name, ok := d.GetOk("oper_peer_redundancy_templ_name"); ok {
		vnicLanConnTemplAttr.Oper_peer_redundancy_templ_name = Oper_peer_redundancy_templ_name.(string)
	}

	if Oper_qos_policy_name, ok := d.GetOk("oper_qos_policy_name"); ok {
		vnicLanConnTemplAttr.Oper_qos_policy_name = Oper_qos_policy_name.(string)
	}

	if Oper_stats_policy_name, ok := d.GetOk("oper_stats_policy_name"); ok {
		vnicLanConnTemplAttr.Oper_stats_policy_name = Oper_stats_policy_name.(string)
	}

	if Peer_redundancy_templ_name, ok := d.GetOk("peer_redundancy_templ_name"); ok {
		vnicLanConnTemplAttr.Peer_redundancy_templ_name = Peer_redundancy_templ_name.(string)
	}

	if Pin_to_group_name, ok := d.GetOk("pin_to_group_name"); ok {
		vnicLanConnTemplAttr.Pin_to_group_name = Pin_to_group_name.(string)
	}

	if Policy_level, ok := d.GetOk("policy_level"); ok {
		vnicLanConnTemplAttr.Policy_level = Policy_level.(string)
	}

	if Policy_owner, ok := d.GetOk("policy_owner"); ok {
		vnicLanConnTemplAttr.Policy_owner = Policy_owner.(string)
	}

	if Qos_policy_name, ok := d.GetOk("qos_policy_name"); ok {
		vnicLanConnTemplAttr.Qos_policy_name = Qos_policy_name.(string)
	}

	if Redundancy_pair_type, ok := d.GetOk("redundancy_pair_type"); ok {
		vnicLanConnTemplAttr.Redundancy_pair_type = Redundancy_pair_type.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		vnicLanConnTemplAttr.Sacl = Sacl.(string)
	}

	if Stats_policy_name, ok := d.GetOk("stats_policy_name"); ok {
		vnicLanConnTemplAttr.Stats_policy_name = Stats_policy_name.(string)
	}

	if Switch_id, ok := d.GetOk("switch_id"); ok {
		vnicLanConnTemplAttr.Switch_id = Switch_id.(string)
	}

	if Target, ok := d.GetOk("target"); ok {
		vnicLanConnTemplAttr.Target = Target.(string)
	}

	if Templ_type, ok := d.GetOk("templ_type"); ok {
		vnicLanConnTemplAttr.Templ_type = Templ_type.(string)
	}

	vnicLanConnTempl := models.NewVnicLanConnTempl(fmt.Sprintf("lan-conn-templ-%s", Name), OrgOrg, desc, vnicLanConnTemplAttr)

	err := ucsClient.Save(vnicLanConnTempl)
	if err != nil {
		return err
	}

	d.SetId(vnicLanConnTempl.DistinguishedName)
	return resourceUcsVnicLanConnTemplRead(d, m)
}

func resourceUcsVnicLanConnTemplRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	vnicLanConnTempl, err := getRemoteVnicLanConnTempl(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setVnicLanConnTemplAttributes(vnicLanConnTempl, d)

	return nil
}

func resourceUcsVnicLanConnTemplDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "vnicLanConnTempl")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsVnicLanConnTemplUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	OrgOrg := d.Get("org_org_dn").(string)

	vnicLanConnTemplAttr := models.VnicLanConnTemplAttributes{}
	if Admin_cdn_name, ok := d.GetOk("admin_cdn_name"); ok {
		vnicLanConnTemplAttr.Admin_cdn_name = Admin_cdn_name.(string)
	}
	if Cdn_source, ok := d.GetOk("cdn_source"); ok {
		vnicLanConnTemplAttr.Cdn_source = Cdn_source.(string)
	}
	if Child_action, ok := d.GetOk("child_action"); ok {
		vnicLanConnTemplAttr.Child_action = Child_action.(string)
	}
	if Ident_pool_name, ok := d.GetOk("ident_pool_name"); ok {
		vnicLanConnTemplAttr.Ident_pool_name = Ident_pool_name.(string)
	}
	if Int_id, ok := d.GetOk("int_id"); ok {
		vnicLanConnTemplAttr.Int_id = Int_id.(string)
	}
	if Mtu, ok := d.GetOk("mtu"); ok {
		vnicLanConnTemplAttr.Mtu = Mtu.(string)
	}
	if Nw_ctrl_policy_name, ok := d.GetOk("nw_ctrl_policy_name"); ok {
		vnicLanConnTemplAttr.Nw_ctrl_policy_name = Nw_ctrl_policy_name.(string)
	}
	if Oper_ident_pool_name, ok := d.GetOk("oper_ident_pool_name"); ok {
		vnicLanConnTemplAttr.Oper_ident_pool_name = Oper_ident_pool_name.(string)
	}
	if Oper_nw_ctrl_policy_name, ok := d.GetOk("oper_nw_ctrl_policy_name"); ok {
		vnicLanConnTemplAttr.Oper_nw_ctrl_policy_name = Oper_nw_ctrl_policy_name.(string)
	}
	if Oper_peer_redundancy_templ_name, ok := d.GetOk("oper_peer_redundancy_templ_name"); ok {
		vnicLanConnTemplAttr.Oper_peer_redundancy_templ_name = Oper_peer_redundancy_templ_name.(string)
	}
	if Oper_qos_policy_name, ok := d.GetOk("oper_qos_policy_name"); ok {
		vnicLanConnTemplAttr.Oper_qos_policy_name = Oper_qos_policy_name.(string)
	}
	if Oper_stats_policy_name, ok := d.GetOk("oper_stats_policy_name"); ok {
		vnicLanConnTemplAttr.Oper_stats_policy_name = Oper_stats_policy_name.(string)
	}
	if Peer_redundancy_templ_name, ok := d.GetOk("peer_redundancy_templ_name"); ok {
		vnicLanConnTemplAttr.Peer_redundancy_templ_name = Peer_redundancy_templ_name.(string)
	}
	if Pin_to_group_name, ok := d.GetOk("pin_to_group_name"); ok {
		vnicLanConnTemplAttr.Pin_to_group_name = Pin_to_group_name.(string)
	}
	if Policy_level, ok := d.GetOk("policy_level"); ok {
		vnicLanConnTemplAttr.Policy_level = Policy_level.(string)
	}
	if Policy_owner, ok := d.GetOk("policy_owner"); ok {
		vnicLanConnTemplAttr.Policy_owner = Policy_owner.(string)
	}
	if Qos_policy_name, ok := d.GetOk("qos_policy_name"); ok {
		vnicLanConnTemplAttr.Qos_policy_name = Qos_policy_name.(string)
	}
	if Redundancy_pair_type, ok := d.GetOk("redundancy_pair_type"); ok {
		vnicLanConnTemplAttr.Redundancy_pair_type = Redundancy_pair_type.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		vnicLanConnTemplAttr.Sacl = Sacl.(string)
	}
	if Stats_policy_name, ok := d.GetOk("stats_policy_name"); ok {
		vnicLanConnTemplAttr.Stats_policy_name = Stats_policy_name.(string)
	}
	if Switch_id, ok := d.GetOk("switch_id"); ok {
		vnicLanConnTemplAttr.Switch_id = Switch_id.(string)
	}
	if Target, ok := d.GetOk("target"); ok {
		vnicLanConnTemplAttr.Target = Target.(string)
	}
	if Templ_type, ok := d.GetOk("templ_type"); ok {
		vnicLanConnTemplAttr.Templ_type = Templ_type.(string)
	}

	vnicLanConnTempl := models.NewVnicLanConnTempl(fmt.Sprintf("lan-conn-templ-%s", Name), OrgOrg, desc, vnicLanConnTemplAttr)
	vnicLanConnTempl.Status = "modified"
	err := ucsClient.Save(vnicLanConnTempl)
	if err != nil {
		return err
	}

	d.SetId(vnicLanConnTempl.DistinguishedName)
	return resourceUcsVnicLanConnTemplRead(d, m)
}
