package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsLsServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsLsServerCreate,
		Update: resourceUcsLsServerUpdate,
		Read:   resourceUcsLsServerRead,
		Delete: resourceUcsLsServerDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsLsServerImport,
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

			"agent_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"assign_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"assoc_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"bios_profile_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"boot_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"child_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"config_qualifier": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"config_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"dynamic_con_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"ext_ip_pool_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"ext_ip_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"flt_aggr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"fsm_descr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"fsm_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"fsm_prev": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"fsm_progr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"fsm_rmt_inv_err_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"fsm_rmt_inv_err_descr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"fsm_rmt_inv_rslt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"fsm_stage_descr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"fsm_stamp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"fsm_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"fsm_try": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"graphics_card_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"host_fw_policy_name": &schema.Schema{
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

			"kvm_mgmt_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"local_disk_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"maint_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"mgmt_access_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"mgmt_fw_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_bios_profile_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_boot_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_dynamic_con_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_ext_ip_pool_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_graphics_card_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_host_fw_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_ident_pool_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_kvm_mgmt_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_local_disk_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_maint_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_mgmt_access_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_mgmt_fw_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_power_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_power_sync_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_scrub_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_sol_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_src_templ_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_stats_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_vcon_profile_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_vmedia_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"owner": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"pn_dn": &schema.Schema{
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

			"power_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"power_sync_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"prop_acl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"resolve_remote": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"sacl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"scrub_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"sol_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"src_templ_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"stats_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"svnic_config": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"usr_lbl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"uuid_suffix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"vcon_profile_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"vmedia_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		}),
	}
}

func getRemoteLsServer(client *client.Client, dn string) (*models.LsServer, error) {
	lsServerDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	lsServer := models.LsServerFromDoc(lsServerDoc, "configResolveDn")

	if lsServer.DistinguishedName == "" {
		return nil, fmt.Errorf("LsServer %s not found", dn)
	}

	return lsServer, nil
}

func setLsServerAttributes(lsServer *models.LsServer, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(lsServer.DistinguishedName)
	d.Set("description", lsServer.Description)
	d.Set("org_org_dn", GetParentDn(lsServer.DistinguishedName))
	lsServerMap, _ := lsServer.ToMap()

	d.Set("agent_policy_name", lsServerMap["agentPolicyName"])

	d.Set("assign_state", lsServerMap["assignState"])

	d.Set("assoc_state", lsServerMap["assocState"])

	d.Set("bios_profile_name", lsServerMap["biosProfileName"])

	d.Set("boot_policy_name", lsServerMap["bootPolicyName"])

	d.Set("child_action", lsServerMap["childAction"])

	d.Set("config_qualifier", lsServerMap["configQualifier"])

	d.Set("config_state", lsServerMap["configState"])

	d.Set("dynamic_con_policy_name", lsServerMap["dynamicConPolicyName"])

	d.Set("ext_ip_pool_name", lsServerMap["extIPPoolName"])

	d.Set("ext_ip_state", lsServerMap["extIPState"])

	d.Set("flt_aggr", lsServerMap["fltAggr"])

	d.Set("fsm_descr", lsServerMap["fsmDescr"])

	d.Set("fsm_flags", lsServerMap["fsmFlags"])

	d.Set("fsm_prev", lsServerMap["fsmPrev"])

	d.Set("fsm_progr", lsServerMap["fsmProgr"])

	d.Set("fsm_rmt_inv_err_code", lsServerMap["fsmRmtInvErrCode"])

	d.Set("fsm_rmt_inv_err_descr", lsServerMap["fsmRmtInvErrDescr"])

	d.Set("fsm_rmt_inv_rslt", lsServerMap["fsmRmtInvRslt"])

	d.Set("fsm_stage_descr", lsServerMap["fsmStageDescr"])

	d.Set("fsm_stamp", lsServerMap["fsmStamp"])

	d.Set("fsm_status", lsServerMap["fsmStatus"])

	d.Set("fsm_try", lsServerMap["fsmTry"])

	d.Set("graphics_card_policy_name", lsServerMap["graphicsCardPolicyName"])

	d.Set("host_fw_policy_name", lsServerMap["hostFwPolicyName"])

	d.Set("ident_pool_name", lsServerMap["identPoolName"])

	d.Set("int_id", lsServerMap["intId"])

	d.Set("kvm_mgmt_policy_name", lsServerMap["kvmMgmtPolicyName"])

	d.Set("local_disk_policy_name", lsServerMap["localDiskPolicyName"])

	d.Set("maint_policy_name", lsServerMap["maintPolicyName"])

	d.Set("mgmt_access_policy_name", lsServerMap["mgmtAccessPolicyName"])

	d.Set("mgmt_fw_policy_name", lsServerMap["mgmtFwPolicyName"])

	d.Set("oper_bios_profile_name", lsServerMap["operBiosProfileName"])

	d.Set("oper_boot_policy_name", lsServerMap["operBootPolicyName"])

	d.Set("oper_dynamic_con_policy_name", lsServerMap["operDynamicConPolicyName"])

	d.Set("oper_ext_ip_pool_name", lsServerMap["operExtIPPoolName"])

	d.Set("oper_graphics_card_policy_name", lsServerMap["operGraphicsCardPolicyName"])

	d.Set("oper_host_fw_policy_name", lsServerMap["operHostFwPolicyName"])

	d.Set("oper_ident_pool_name", lsServerMap["operIdentPoolName"])

	d.Set("oper_kvm_mgmt_policy_name", lsServerMap["operKvmMgmtPolicyName"])

	d.Set("oper_local_disk_policy_name", lsServerMap["operLocalDiskPolicyName"])

	d.Set("oper_maint_policy_name", lsServerMap["operMaintPolicyName"])

	d.Set("oper_mgmt_access_policy_name", lsServerMap["operMgmtAccessPolicyName"])

	d.Set("oper_mgmt_fw_policy_name", lsServerMap["operMgmtFwPolicyName"])

	d.Set("oper_power_policy_name", lsServerMap["operPowerPolicyName"])

	d.Set("oper_power_sync_policy_name", lsServerMap["operPowerSyncPolicyName"])

	d.Set("oper_scrub_policy_name", lsServerMap["operScrubPolicyName"])

	d.Set("oper_sol_policy_name", lsServerMap["operSolPolicyName"])

	d.Set("oper_src_templ_name", lsServerMap["operSrcTemplName"])

	d.Set("oper_state", lsServerMap["operState"])

	d.Set("oper_stats_policy_name", lsServerMap["operStatsPolicyName"])

	d.Set("oper_vcon_profile_name", lsServerMap["operVconProfileName"])

	d.Set("oper_vmedia_policy_name", lsServerMap["operVmediaPolicyName"])

	d.Set("owner", lsServerMap["owner"])

	d.Set("pn_dn", lsServerMap["pnDn"])

	d.Set("policy_level", lsServerMap["policyLevel"])

	d.Set("policy_owner", lsServerMap["policyOwner"])

	d.Set("power_policy_name", lsServerMap["powerPolicyName"])

	d.Set("power_sync_policy_name", lsServerMap["powerSyncPolicyName"])

	d.Set("prop_acl", lsServerMap["propAcl"])

	d.Set("resolve_remote", lsServerMap["resolveRemote"])

	d.Set("sacl", lsServerMap["sacl"])

	d.Set("scrub_policy_name", lsServerMap["scrubPolicyName"])

	d.Set("sol_policy_name", lsServerMap["solPolicyName"])

	d.Set("src_templ_name", lsServerMap["srcTemplName"])

	d.Set("stats_policy_name", lsServerMap["statsPolicyName"])

	d.Set("svnic_config", lsServerMap["svnicConfig"])

	d.Set("type", lsServerMap["type"])

	d.Set("usr_lbl", lsServerMap["usrLbl"])

	d.Set("uuid", lsServerMap["uuid"])

	d.Set("uuid_suffix", lsServerMap["uuidSuffix"])

	d.Set("vcon_profile_name", lsServerMap["vconProfileName"])

	d.Set("vmedia_policy_name", lsServerMap["vmediaPolicyName"])
	return d
}

func resourceUcsLsServerImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	lsServer, err := getRemoteLsServer(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setLsServerAttributes(lsServer, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsLsServerCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	OrgOrg := d.Get("org_org_dn").(string)

	lsServerAttr := models.LsServerAttributes{}

	if Agent_policy_name, ok := d.GetOk("agent_policy_name"); ok {
		lsServerAttr.Agent_policy_name = Agent_policy_name.(string)
	}

	if Assign_state, ok := d.GetOk("assign_state"); ok {
		lsServerAttr.Assign_state = Assign_state.(string)
	}

	if Assoc_state, ok := d.GetOk("assoc_state"); ok {
		lsServerAttr.Assoc_state = Assoc_state.(string)
	}

	if Bios_profile_name, ok := d.GetOk("bios_profile_name"); ok {
		lsServerAttr.Bios_profile_name = Bios_profile_name.(string)
	}

	if Boot_policy_name, ok := d.GetOk("boot_policy_name"); ok {
		lsServerAttr.Boot_policy_name = Boot_policy_name.(string)
	}

	if Child_action, ok := d.GetOk("child_action"); ok {
		lsServerAttr.Child_action = Child_action.(string)
	}

	if Config_qualifier, ok := d.GetOk("config_qualifier"); ok {
		lsServerAttr.Config_qualifier = Config_qualifier.(string)
	}

	if Config_state, ok := d.GetOk("config_state"); ok {
		lsServerAttr.Config_state = Config_state.(string)
	}

	if Dynamic_con_policy_name, ok := d.GetOk("dynamic_con_policy_name"); ok {
		lsServerAttr.Dynamic_con_policy_name = Dynamic_con_policy_name.(string)
	}

	if Ext_ip_pool_name, ok := d.GetOk("ext_ip_pool_name"); ok {
		lsServerAttr.Ext_ip_pool_name = Ext_ip_pool_name.(string)
	}

	if Ext_ip_state, ok := d.GetOk("ext_ip_state"); ok {
		lsServerAttr.Ext_ip_state = Ext_ip_state.(string)
	}

	if Flt_aggr, ok := d.GetOk("flt_aggr"); ok {
		lsServerAttr.Flt_aggr = Flt_aggr.(string)
	}

	if Fsm_descr, ok := d.GetOk("fsm_descr"); ok {
		lsServerAttr.Fsm_descr = Fsm_descr.(string)
	}

	if Fsm_flags, ok := d.GetOk("fsm_flags"); ok {
		lsServerAttr.Fsm_flags = Fsm_flags.(string)
	}

	if Fsm_prev, ok := d.GetOk("fsm_prev"); ok {
		lsServerAttr.Fsm_prev = Fsm_prev.(string)
	}

	if Fsm_progr, ok := d.GetOk("fsm_progr"); ok {
		lsServerAttr.Fsm_progr = Fsm_progr.(string)
	}

	if Fsm_rmt_inv_err_code, ok := d.GetOk("fsm_rmt_inv_err_code"); ok {
		lsServerAttr.Fsm_rmt_inv_err_code = Fsm_rmt_inv_err_code.(string)
	}

	if Fsm_rmt_inv_err_descr, ok := d.GetOk("fsm_rmt_inv_err_descr"); ok {
		lsServerAttr.Fsm_rmt_inv_err_descr = Fsm_rmt_inv_err_descr.(string)
	}

	if Fsm_rmt_inv_rslt, ok := d.GetOk("fsm_rmt_inv_rslt"); ok {
		lsServerAttr.Fsm_rmt_inv_rslt = Fsm_rmt_inv_rslt.(string)
	}

	if Fsm_stage_descr, ok := d.GetOk("fsm_stage_descr"); ok {
		lsServerAttr.Fsm_stage_descr = Fsm_stage_descr.(string)
	}

	if Fsm_stamp, ok := d.GetOk("fsm_stamp"); ok {
		lsServerAttr.Fsm_stamp = Fsm_stamp.(string)
	}

	if Fsm_status, ok := d.GetOk("fsm_status"); ok {
		lsServerAttr.Fsm_status = Fsm_status.(string)
	}

	if Fsm_try, ok := d.GetOk("fsm_try"); ok {
		lsServerAttr.Fsm_try = Fsm_try.(string)
	}

	if Graphics_card_policy_name, ok := d.GetOk("graphics_card_policy_name"); ok {
		lsServerAttr.Graphics_card_policy_name = Graphics_card_policy_name.(string)
	}

	if Host_fw_policy_name, ok := d.GetOk("host_fw_policy_name"); ok {
		lsServerAttr.Host_fw_policy_name = Host_fw_policy_name.(string)
	}

	if Ident_pool_name, ok := d.GetOk("ident_pool_name"); ok {
		lsServerAttr.Ident_pool_name = Ident_pool_name.(string)
	}

	if Int_id, ok := d.GetOk("int_id"); ok {
		lsServerAttr.Int_id = Int_id.(string)
	}

	if Kvm_mgmt_policy_name, ok := d.GetOk("kvm_mgmt_policy_name"); ok {
		lsServerAttr.Kvm_mgmt_policy_name = Kvm_mgmt_policy_name.(string)
	}

	if Local_disk_policy_name, ok := d.GetOk("local_disk_policy_name"); ok {
		lsServerAttr.Local_disk_policy_name = Local_disk_policy_name.(string)
	}

	if Maint_policy_name, ok := d.GetOk("maint_policy_name"); ok {
		lsServerAttr.Maint_policy_name = Maint_policy_name.(string)
	}

	if Mgmt_access_policy_name, ok := d.GetOk("mgmt_access_policy_name"); ok {
		lsServerAttr.Mgmt_access_policy_name = Mgmt_access_policy_name.(string)
	}

	if Mgmt_fw_policy_name, ok := d.GetOk("mgmt_fw_policy_name"); ok {
		lsServerAttr.Mgmt_fw_policy_name = Mgmt_fw_policy_name.(string)
	}

	if Oper_bios_profile_name, ok := d.GetOk("oper_bios_profile_name"); ok {
		lsServerAttr.Oper_bios_profile_name = Oper_bios_profile_name.(string)
	}

	if Oper_boot_policy_name, ok := d.GetOk("oper_boot_policy_name"); ok {
		lsServerAttr.Oper_boot_policy_name = Oper_boot_policy_name.(string)
	}

	if Oper_dynamic_con_policy_name, ok := d.GetOk("oper_dynamic_con_policy_name"); ok {
		lsServerAttr.Oper_dynamic_con_policy_name = Oper_dynamic_con_policy_name.(string)
	}

	if Oper_ext_ip_pool_name, ok := d.GetOk("oper_ext_ip_pool_name"); ok {
		lsServerAttr.Oper_ext_ip_pool_name = Oper_ext_ip_pool_name.(string)
	}

	if Oper_graphics_card_policy_name, ok := d.GetOk("oper_graphics_card_policy_name"); ok {
		lsServerAttr.Oper_graphics_card_policy_name = Oper_graphics_card_policy_name.(string)
	}

	if Oper_host_fw_policy_name, ok := d.GetOk("oper_host_fw_policy_name"); ok {
		lsServerAttr.Oper_host_fw_policy_name = Oper_host_fw_policy_name.(string)
	}

	if Oper_ident_pool_name, ok := d.GetOk("oper_ident_pool_name"); ok {
		lsServerAttr.Oper_ident_pool_name = Oper_ident_pool_name.(string)
	}

	if Oper_kvm_mgmt_policy_name, ok := d.GetOk("oper_kvm_mgmt_policy_name"); ok {
		lsServerAttr.Oper_kvm_mgmt_policy_name = Oper_kvm_mgmt_policy_name.(string)
	}

	if Oper_local_disk_policy_name, ok := d.GetOk("oper_local_disk_policy_name"); ok {
		lsServerAttr.Oper_local_disk_policy_name = Oper_local_disk_policy_name.(string)
	}

	if Oper_maint_policy_name, ok := d.GetOk("oper_maint_policy_name"); ok {
		lsServerAttr.Oper_maint_policy_name = Oper_maint_policy_name.(string)
	}

	if Oper_mgmt_access_policy_name, ok := d.GetOk("oper_mgmt_access_policy_name"); ok {
		lsServerAttr.Oper_mgmt_access_policy_name = Oper_mgmt_access_policy_name.(string)
	}

	if Oper_mgmt_fw_policy_name, ok := d.GetOk("oper_mgmt_fw_policy_name"); ok {
		lsServerAttr.Oper_mgmt_fw_policy_name = Oper_mgmt_fw_policy_name.(string)
	}

	if Oper_power_policy_name, ok := d.GetOk("oper_power_policy_name"); ok {
		lsServerAttr.Oper_power_policy_name = Oper_power_policy_name.(string)
	}

	if Oper_power_sync_policy_name, ok := d.GetOk("oper_power_sync_policy_name"); ok {
		lsServerAttr.Oper_power_sync_policy_name = Oper_power_sync_policy_name.(string)
	}

	if Oper_scrub_policy_name, ok := d.GetOk("oper_scrub_policy_name"); ok {
		lsServerAttr.Oper_scrub_policy_name = Oper_scrub_policy_name.(string)
	}

	if Oper_sol_policy_name, ok := d.GetOk("oper_sol_policy_name"); ok {
		lsServerAttr.Oper_sol_policy_name = Oper_sol_policy_name.(string)
	}

	if Oper_src_templ_name, ok := d.GetOk("oper_src_templ_name"); ok {
		lsServerAttr.Oper_src_templ_name = Oper_src_templ_name.(string)
	}

	if Oper_state, ok := d.GetOk("oper_state"); ok {
		lsServerAttr.Oper_state = Oper_state.(string)
	}

	if Oper_stats_policy_name, ok := d.GetOk("oper_stats_policy_name"); ok {
		lsServerAttr.Oper_stats_policy_name = Oper_stats_policy_name.(string)
	}

	if Oper_vcon_profile_name, ok := d.GetOk("oper_vcon_profile_name"); ok {
		lsServerAttr.Oper_vcon_profile_name = Oper_vcon_profile_name.(string)
	}

	if Oper_vmedia_policy_name, ok := d.GetOk("oper_vmedia_policy_name"); ok {
		lsServerAttr.Oper_vmedia_policy_name = Oper_vmedia_policy_name.(string)
	}

	if Owner, ok := d.GetOk("owner"); ok {
		lsServerAttr.Owner = Owner.(string)
	}

	if Pn_dn, ok := d.GetOk("pn_dn"); ok {
		lsServerAttr.Pn_dn = Pn_dn.(string)
	}

	if Policy_level, ok := d.GetOk("policy_level"); ok {
		lsServerAttr.Policy_level = Policy_level.(string)
	}

	if Policy_owner, ok := d.GetOk("policy_owner"); ok {
		lsServerAttr.Policy_owner = Policy_owner.(string)
	}

	if Power_policy_name, ok := d.GetOk("power_policy_name"); ok {
		lsServerAttr.Power_policy_name = Power_policy_name.(string)
	}

	if Power_sync_policy_name, ok := d.GetOk("power_sync_policy_name"); ok {
		lsServerAttr.Power_sync_policy_name = Power_sync_policy_name.(string)
	}

	if Prop_acl, ok := d.GetOk("prop_acl"); ok {
		lsServerAttr.Prop_acl = Prop_acl.(string)
	}

	if Resolve_remote, ok := d.GetOk("resolve_remote"); ok {
		lsServerAttr.Resolve_remote = Resolve_remote.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		lsServerAttr.Sacl = Sacl.(string)
	}

	if Scrub_policy_name, ok := d.GetOk("scrub_policy_name"); ok {
		lsServerAttr.Scrub_policy_name = Scrub_policy_name.(string)
	}

	if Sol_policy_name, ok := d.GetOk("sol_policy_name"); ok {
		lsServerAttr.Sol_policy_name = Sol_policy_name.(string)
	}

	if Src_templ_name, ok := d.GetOk("src_templ_name"); ok {
		lsServerAttr.Src_templ_name = Src_templ_name.(string)
	}

	if Stats_policy_name, ok := d.GetOk("stats_policy_name"); ok {
		lsServerAttr.Stats_policy_name = Stats_policy_name.(string)
	}

	if Svnic_config, ok := d.GetOk("svnic_config"); ok {
		lsServerAttr.Svnic_config = Svnic_config.(string)
	}

	if Type, ok := d.GetOk("type"); ok {
		lsServerAttr.Type = Type.(string)
	}

	if Usr_lbl, ok := d.GetOk("usr_lbl"); ok {
		lsServerAttr.Usr_lbl = Usr_lbl.(string)
	}

	if Uuid, ok := d.GetOk("uuid"); ok {
		lsServerAttr.Uuid = Uuid.(string)
	}

	if Uuid_suffix, ok := d.GetOk("uuid_suffix"); ok {
		lsServerAttr.Uuid_suffix = Uuid_suffix.(string)
	}

	if Vcon_profile_name, ok := d.GetOk("vcon_profile_name"); ok {
		lsServerAttr.Vcon_profile_name = Vcon_profile_name.(string)
	}

	if Vmedia_policy_name, ok := d.GetOk("vmedia_policy_name"); ok {
		lsServerAttr.Vmedia_policy_name = Vmedia_policy_name.(string)
	}

	lsServer := models.NewLsServer(fmt.Sprintf("ls-%s", Name), OrgOrg, desc, lsServerAttr)

	err := ucsClient.Save(lsServer)
	if err != nil {
		return err
	}

	d.SetId(lsServer.DistinguishedName)
	return resourceUcsLsServerRead(d, m)
}

func resourceUcsLsServerRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	lsServer, err := getRemoteLsServer(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setLsServerAttributes(lsServer, d)

	return nil
}

func resourceUcsLsServerDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "lsServer")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsLsServerUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	OrgOrg := d.Get("org_org_dn").(string)

	lsServerAttr := models.LsServerAttributes{}
	if Agent_policy_name, ok := d.GetOk("agent_policy_name"); ok {
		lsServerAttr.Agent_policy_name = Agent_policy_name.(string)
	}
	if Assign_state, ok := d.GetOk("assign_state"); ok {
		lsServerAttr.Assign_state = Assign_state.(string)
	}
	if Assoc_state, ok := d.GetOk("assoc_state"); ok {
		lsServerAttr.Assoc_state = Assoc_state.(string)
	}
	if Bios_profile_name, ok := d.GetOk("bios_profile_name"); ok {
		lsServerAttr.Bios_profile_name = Bios_profile_name.(string)
	}
	if Boot_policy_name, ok := d.GetOk("boot_policy_name"); ok {
		lsServerAttr.Boot_policy_name = Boot_policy_name.(string)
	}
	if Child_action, ok := d.GetOk("child_action"); ok {
		lsServerAttr.Child_action = Child_action.(string)
	}
	if Config_qualifier, ok := d.GetOk("config_qualifier"); ok {
		lsServerAttr.Config_qualifier = Config_qualifier.(string)
	}
	if Config_state, ok := d.GetOk("config_state"); ok {
		lsServerAttr.Config_state = Config_state.(string)
	}
	if Dynamic_con_policy_name, ok := d.GetOk("dynamic_con_policy_name"); ok {
		lsServerAttr.Dynamic_con_policy_name = Dynamic_con_policy_name.(string)
	}
	if Ext_ip_pool_name, ok := d.GetOk("ext_ip_pool_name"); ok {
		lsServerAttr.Ext_ip_pool_name = Ext_ip_pool_name.(string)
	}
	if Ext_ip_state, ok := d.GetOk("ext_ip_state"); ok {
		lsServerAttr.Ext_ip_state = Ext_ip_state.(string)
	}
	if Flt_aggr, ok := d.GetOk("flt_aggr"); ok {
		lsServerAttr.Flt_aggr = Flt_aggr.(string)
	}
	if Fsm_descr, ok := d.GetOk("fsm_descr"); ok {
		lsServerAttr.Fsm_descr = Fsm_descr.(string)
	}
	if Fsm_flags, ok := d.GetOk("fsm_flags"); ok {
		lsServerAttr.Fsm_flags = Fsm_flags.(string)
	}
	if Fsm_prev, ok := d.GetOk("fsm_prev"); ok {
		lsServerAttr.Fsm_prev = Fsm_prev.(string)
	}
	if Fsm_progr, ok := d.GetOk("fsm_progr"); ok {
		lsServerAttr.Fsm_progr = Fsm_progr.(string)
	}
	if Fsm_rmt_inv_err_code, ok := d.GetOk("fsm_rmt_inv_err_code"); ok {
		lsServerAttr.Fsm_rmt_inv_err_code = Fsm_rmt_inv_err_code.(string)
	}
	if Fsm_rmt_inv_err_descr, ok := d.GetOk("fsm_rmt_inv_err_descr"); ok {
		lsServerAttr.Fsm_rmt_inv_err_descr = Fsm_rmt_inv_err_descr.(string)
	}
	if Fsm_rmt_inv_rslt, ok := d.GetOk("fsm_rmt_inv_rslt"); ok {
		lsServerAttr.Fsm_rmt_inv_rslt = Fsm_rmt_inv_rslt.(string)
	}
	if Fsm_stage_descr, ok := d.GetOk("fsm_stage_descr"); ok {
		lsServerAttr.Fsm_stage_descr = Fsm_stage_descr.(string)
	}
	if Fsm_stamp, ok := d.GetOk("fsm_stamp"); ok {
		lsServerAttr.Fsm_stamp = Fsm_stamp.(string)
	}
	if Fsm_status, ok := d.GetOk("fsm_status"); ok {
		lsServerAttr.Fsm_status = Fsm_status.(string)
	}
	if Fsm_try, ok := d.GetOk("fsm_try"); ok {
		lsServerAttr.Fsm_try = Fsm_try.(string)
	}
	if Graphics_card_policy_name, ok := d.GetOk("graphics_card_policy_name"); ok {
		lsServerAttr.Graphics_card_policy_name = Graphics_card_policy_name.(string)
	}
	if Host_fw_policy_name, ok := d.GetOk("host_fw_policy_name"); ok {
		lsServerAttr.Host_fw_policy_name = Host_fw_policy_name.(string)
	}
	if Ident_pool_name, ok := d.GetOk("ident_pool_name"); ok {
		lsServerAttr.Ident_pool_name = Ident_pool_name.(string)
	}
	if Int_id, ok := d.GetOk("int_id"); ok {
		lsServerAttr.Int_id = Int_id.(string)
	}
	if Kvm_mgmt_policy_name, ok := d.GetOk("kvm_mgmt_policy_name"); ok {
		lsServerAttr.Kvm_mgmt_policy_name = Kvm_mgmt_policy_name.(string)
	}
	if Local_disk_policy_name, ok := d.GetOk("local_disk_policy_name"); ok {
		lsServerAttr.Local_disk_policy_name = Local_disk_policy_name.(string)
	}
	if Maint_policy_name, ok := d.GetOk("maint_policy_name"); ok {
		lsServerAttr.Maint_policy_name = Maint_policy_name.(string)
	}
	if Mgmt_access_policy_name, ok := d.GetOk("mgmt_access_policy_name"); ok {
		lsServerAttr.Mgmt_access_policy_name = Mgmt_access_policy_name.(string)
	}
	if Mgmt_fw_policy_name, ok := d.GetOk("mgmt_fw_policy_name"); ok {
		lsServerAttr.Mgmt_fw_policy_name = Mgmt_fw_policy_name.(string)
	}
	if Oper_bios_profile_name, ok := d.GetOk("oper_bios_profile_name"); ok {
		lsServerAttr.Oper_bios_profile_name = Oper_bios_profile_name.(string)
	}
	if Oper_boot_policy_name, ok := d.GetOk("oper_boot_policy_name"); ok {
		lsServerAttr.Oper_boot_policy_name = Oper_boot_policy_name.(string)
	}
	if Oper_dynamic_con_policy_name, ok := d.GetOk("oper_dynamic_con_policy_name"); ok {
		lsServerAttr.Oper_dynamic_con_policy_name = Oper_dynamic_con_policy_name.(string)
	}
	if Oper_ext_ip_pool_name, ok := d.GetOk("oper_ext_ip_pool_name"); ok {
		lsServerAttr.Oper_ext_ip_pool_name = Oper_ext_ip_pool_name.(string)
	}
	if Oper_graphics_card_policy_name, ok := d.GetOk("oper_graphics_card_policy_name"); ok {
		lsServerAttr.Oper_graphics_card_policy_name = Oper_graphics_card_policy_name.(string)
	}
	if Oper_host_fw_policy_name, ok := d.GetOk("oper_host_fw_policy_name"); ok {
		lsServerAttr.Oper_host_fw_policy_name = Oper_host_fw_policy_name.(string)
	}
	if Oper_ident_pool_name, ok := d.GetOk("oper_ident_pool_name"); ok {
		lsServerAttr.Oper_ident_pool_name = Oper_ident_pool_name.(string)
	}
	if Oper_kvm_mgmt_policy_name, ok := d.GetOk("oper_kvm_mgmt_policy_name"); ok {
		lsServerAttr.Oper_kvm_mgmt_policy_name = Oper_kvm_mgmt_policy_name.(string)
	}
	if Oper_local_disk_policy_name, ok := d.GetOk("oper_local_disk_policy_name"); ok {
		lsServerAttr.Oper_local_disk_policy_name = Oper_local_disk_policy_name.(string)
	}
	if Oper_maint_policy_name, ok := d.GetOk("oper_maint_policy_name"); ok {
		lsServerAttr.Oper_maint_policy_name = Oper_maint_policy_name.(string)
	}
	if Oper_mgmt_access_policy_name, ok := d.GetOk("oper_mgmt_access_policy_name"); ok {
		lsServerAttr.Oper_mgmt_access_policy_name = Oper_mgmt_access_policy_name.(string)
	}
	if Oper_mgmt_fw_policy_name, ok := d.GetOk("oper_mgmt_fw_policy_name"); ok {
		lsServerAttr.Oper_mgmt_fw_policy_name = Oper_mgmt_fw_policy_name.(string)
	}
	if Oper_power_policy_name, ok := d.GetOk("oper_power_policy_name"); ok {
		lsServerAttr.Oper_power_policy_name = Oper_power_policy_name.(string)
	}
	if Oper_power_sync_policy_name, ok := d.GetOk("oper_power_sync_policy_name"); ok {
		lsServerAttr.Oper_power_sync_policy_name = Oper_power_sync_policy_name.(string)
	}
	if Oper_scrub_policy_name, ok := d.GetOk("oper_scrub_policy_name"); ok {
		lsServerAttr.Oper_scrub_policy_name = Oper_scrub_policy_name.(string)
	}
	if Oper_sol_policy_name, ok := d.GetOk("oper_sol_policy_name"); ok {
		lsServerAttr.Oper_sol_policy_name = Oper_sol_policy_name.(string)
	}
	if Oper_src_templ_name, ok := d.GetOk("oper_src_templ_name"); ok {
		lsServerAttr.Oper_src_templ_name = Oper_src_templ_name.(string)
	}
	if Oper_state, ok := d.GetOk("oper_state"); ok {
		lsServerAttr.Oper_state = Oper_state.(string)
	}
	if Oper_stats_policy_name, ok := d.GetOk("oper_stats_policy_name"); ok {
		lsServerAttr.Oper_stats_policy_name = Oper_stats_policy_name.(string)
	}
	if Oper_vcon_profile_name, ok := d.GetOk("oper_vcon_profile_name"); ok {
		lsServerAttr.Oper_vcon_profile_name = Oper_vcon_profile_name.(string)
	}
	if Oper_vmedia_policy_name, ok := d.GetOk("oper_vmedia_policy_name"); ok {
		lsServerAttr.Oper_vmedia_policy_name = Oper_vmedia_policy_name.(string)
	}
	if Owner, ok := d.GetOk("owner"); ok {
		lsServerAttr.Owner = Owner.(string)
	}
	if Pn_dn, ok := d.GetOk("pn_dn"); ok {
		lsServerAttr.Pn_dn = Pn_dn.(string)
	}
	if Policy_level, ok := d.GetOk("policy_level"); ok {
		lsServerAttr.Policy_level = Policy_level.(string)
	}
	if Policy_owner, ok := d.GetOk("policy_owner"); ok {
		lsServerAttr.Policy_owner = Policy_owner.(string)
	}
	if Power_policy_name, ok := d.GetOk("power_policy_name"); ok {
		lsServerAttr.Power_policy_name = Power_policy_name.(string)
	}
	if Power_sync_policy_name, ok := d.GetOk("power_sync_policy_name"); ok {
		lsServerAttr.Power_sync_policy_name = Power_sync_policy_name.(string)
	}
	if Prop_acl, ok := d.GetOk("prop_acl"); ok {
		lsServerAttr.Prop_acl = Prop_acl.(string)
	}
	if Resolve_remote, ok := d.GetOk("resolve_remote"); ok {
		lsServerAttr.Resolve_remote = Resolve_remote.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		lsServerAttr.Sacl = Sacl.(string)
	}
	if Scrub_policy_name, ok := d.GetOk("scrub_policy_name"); ok {
		lsServerAttr.Scrub_policy_name = Scrub_policy_name.(string)
	}
	if Sol_policy_name, ok := d.GetOk("sol_policy_name"); ok {
		lsServerAttr.Sol_policy_name = Sol_policy_name.(string)
	}
	if Src_templ_name, ok := d.GetOk("src_templ_name"); ok {
		lsServerAttr.Src_templ_name = Src_templ_name.(string)
	}
	if Stats_policy_name, ok := d.GetOk("stats_policy_name"); ok {
		lsServerAttr.Stats_policy_name = Stats_policy_name.(string)
	}
	if Svnic_config, ok := d.GetOk("svnic_config"); ok {
		lsServerAttr.Svnic_config = Svnic_config.(string)
	}
	if Type, ok := d.GetOk("type"); ok {
		lsServerAttr.Type = Type.(string)
	}
	if Usr_lbl, ok := d.GetOk("usr_lbl"); ok {
		lsServerAttr.Usr_lbl = Usr_lbl.(string)
	}
	if Uuid, ok := d.GetOk("uuid"); ok {
		lsServerAttr.Uuid = Uuid.(string)
	}
	if Uuid_suffix, ok := d.GetOk("uuid_suffix"); ok {
		lsServerAttr.Uuid_suffix = Uuid_suffix.(string)
	}
	if Vcon_profile_name, ok := d.GetOk("vcon_profile_name"); ok {
		lsServerAttr.Vcon_profile_name = Vcon_profile_name.(string)
	}
	if Vmedia_policy_name, ok := d.GetOk("vmedia_policy_name"); ok {
		lsServerAttr.Vmedia_policy_name = Vmedia_policy_name.(string)
	}

	lsServer := models.NewLsServer(fmt.Sprintf("ls-%s", Name), OrgOrg, desc, lsServerAttr)
	lsServer.Status = "modified"
	err := ucsClient.Save(lsServer)
	if err != nil {
		return err
	}

	d.SetId(lsServer.DistinguishedName)
	return resourceUcsLsServerRead(d, m)
}
