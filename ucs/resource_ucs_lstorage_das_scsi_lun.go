package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsLstorageDasScsiLun() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsLstorageDasScsiLunCreate,
		Update: resourceUcsLstorageDasScsiLunUpdate,
		Read:   resourceUcsLstorageDasScsiLunRead,
		Delete: resourceUcsLstorageDasScsiLunDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsLstorageDasScsiLunImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"lstorage_profile_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"admin_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"auto_deploy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"boot_dev": &schema.Schema{
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

			"deferred_naming": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"expand_to_avail": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"fractional_size": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"local_disk_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"lun_dn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"lun_map_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_local_disk_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"order": &schema.Schema{
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

			"storage_class": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		}),
	}
}

func getRemoteLstorageDasScsiLun(client *client.Client, dn string) (*models.LstorageDasScsiLun, error) {
	lstorageDasScsiLunDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	lstorageDasScsiLun := models.LstorageDasScsiLunFromDoc(lstorageDasScsiLunDoc, "configResolveDn")

	if lstorageDasScsiLun.DistinguishedName == "" {
		return nil, fmt.Errorf("LstorageDasScsiLun %s not found", dn)
	}

	return lstorageDasScsiLun, nil
}

func setLstorageDasScsiLunAttributes(lstorageDasScsiLun *models.LstorageDasScsiLun, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(lstorageDasScsiLun.DistinguishedName)
	d.Set("description", lstorageDasScsiLun.Description)
	d.Set("lstorage_profile_dn", GetParentDn(lstorageDasScsiLun.DistinguishedName))
	lstorageDasScsiLunMap, _ := lstorageDasScsiLun.ToMap()

	d.Set("admin_state", lstorageDasScsiLunMap["adminState"])

	d.Set("auto_deploy", lstorageDasScsiLunMap["autoDeploy"])

	d.Set("boot_dev", lstorageDasScsiLunMap["bootDev"])

	d.Set("child_action", lstorageDasScsiLunMap["childAction"])

	d.Set("config_qualifier", lstorageDasScsiLunMap["configQualifier"])

	d.Set("config_state", lstorageDasScsiLunMap["configState"])

	d.Set("deferred_naming", lstorageDasScsiLunMap["deferredNaming"])

	d.Set("expand_to_avail", lstorageDasScsiLunMap["expandToAvail"])

	d.Set("fractional_size", lstorageDasScsiLunMap["fractionalSize"])

	d.Set("local_disk_policy_name", lstorageDasScsiLunMap["localDiskPolicyName"])

	d.Set("lun_dn", lstorageDasScsiLunMap["lunDn"])

	d.Set("lun_map_type", lstorageDasScsiLunMap["lunMapType"])

	d.Set("oper_local_disk_policy_name", lstorageDasScsiLunMap["operLocalDiskPolicyName"])

	d.Set("oper_state", lstorageDasScsiLunMap["operState"])

	d.Set("order", lstorageDasScsiLunMap["order"])

	d.Set("sacl", lstorageDasScsiLunMap["sacl"])

	d.Set("size", lstorageDasScsiLunMap["size"])

	d.Set("storage_class", lstorageDasScsiLunMap["storageClass"])
	return d
}

func resourceUcsLstorageDasScsiLunImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	lstorageDasScsiLun, err := getRemoteLstorageDasScsiLun(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setLstorageDasScsiLunAttributes(lstorageDasScsiLun, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsLstorageDasScsiLunCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	LstorageProfile := d.Get("lstorage_profile_dn").(string)

	lstorageDasScsiLunAttr := models.LstorageDasScsiLunAttributes{}

	if Admin_state, ok := d.GetOk("admin_state"); ok {
		lstorageDasScsiLunAttr.Admin_state = Admin_state.(string)
	}

	if Auto_deploy, ok := d.GetOk("auto_deploy"); ok {
		lstorageDasScsiLunAttr.Auto_deploy = Auto_deploy.(string)
	}

	if Boot_dev, ok := d.GetOk("boot_dev"); ok {
		lstorageDasScsiLunAttr.Boot_dev = Boot_dev.(string)
	}

	if Child_action, ok := d.GetOk("child_action"); ok {
		lstorageDasScsiLunAttr.Child_action = Child_action.(string)
	}

	if Config_qualifier, ok := d.GetOk("config_qualifier"); ok {
		lstorageDasScsiLunAttr.Config_qualifier = Config_qualifier.(string)
	}

	if Config_state, ok := d.GetOk("config_state"); ok {
		lstorageDasScsiLunAttr.Config_state = Config_state.(string)
	}

	if Deferred_naming, ok := d.GetOk("deferred_naming"); ok {
		lstorageDasScsiLunAttr.Deferred_naming = Deferred_naming.(string)
	}

	if Expand_to_avail, ok := d.GetOk("expand_to_avail"); ok {
		lstorageDasScsiLunAttr.Expand_to_avail = Expand_to_avail.(string)
	}

	if Fractional_size, ok := d.GetOk("fractional_size"); ok {
		lstorageDasScsiLunAttr.Fractional_size = Fractional_size.(string)
	}

	if Local_disk_policy_name, ok := d.GetOk("local_disk_policy_name"); ok {
		lstorageDasScsiLunAttr.Local_disk_policy_name = Local_disk_policy_name.(string)
	}

	if Lun_dn, ok := d.GetOk("lun_dn"); ok {
		lstorageDasScsiLunAttr.Lun_dn = Lun_dn.(string)
	}

	if Lun_map_type, ok := d.GetOk("lun_map_type"); ok {
		lstorageDasScsiLunAttr.Lun_map_type = Lun_map_type.(string)
	}

	if Oper_local_disk_policy_name, ok := d.GetOk("oper_local_disk_policy_name"); ok {
		lstorageDasScsiLunAttr.Oper_local_disk_policy_name = Oper_local_disk_policy_name.(string)
	}

	if Oper_state, ok := d.GetOk("oper_state"); ok {
		lstorageDasScsiLunAttr.Oper_state = Oper_state.(string)
	}

	if Order, ok := d.GetOk("order"); ok {
		lstorageDasScsiLunAttr.Order = Order.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		lstorageDasScsiLunAttr.Sacl = Sacl.(string)
	}

	if Size, ok := d.GetOk("size"); ok {
		lstorageDasScsiLunAttr.Size = Size.(string)
	}

	if Storage_class, ok := d.GetOk("storage_class"); ok {
		lstorageDasScsiLunAttr.Storage_class = Storage_class.(string)
	}

	lstorageDasScsiLun := models.NewLstorageDasScsiLun(fmt.Sprintf("das-scsi-lun-%s", Name), LstorageProfile, desc, lstorageDasScsiLunAttr)

	err := ucsClient.Save(lstorageDasScsiLun)
	if err != nil {
		return err
	}

	d.SetId(lstorageDasScsiLun.DistinguishedName)
	return resourceUcsLstorageDasScsiLunRead(d, m)
}

func resourceUcsLstorageDasScsiLunRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	lstorageDasScsiLun, err := getRemoteLstorageDasScsiLun(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setLstorageDasScsiLunAttributes(lstorageDasScsiLun, d)

	return nil
}

func resourceUcsLstorageDasScsiLunDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "lstorageDasScsiLun")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsLstorageDasScsiLunUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	LstorageProfile := d.Get("lstorage_profile_dn").(string)

	lstorageDasScsiLunAttr := models.LstorageDasScsiLunAttributes{}
	if Admin_state, ok := d.GetOk("admin_state"); ok {
		lstorageDasScsiLunAttr.Admin_state = Admin_state.(string)
	}
	if Auto_deploy, ok := d.GetOk("auto_deploy"); ok {
		lstorageDasScsiLunAttr.Auto_deploy = Auto_deploy.(string)
	}
	if Boot_dev, ok := d.GetOk("boot_dev"); ok {
		lstorageDasScsiLunAttr.Boot_dev = Boot_dev.(string)
	}
	if Child_action, ok := d.GetOk("child_action"); ok {
		lstorageDasScsiLunAttr.Child_action = Child_action.(string)
	}
	if Config_qualifier, ok := d.GetOk("config_qualifier"); ok {
		lstorageDasScsiLunAttr.Config_qualifier = Config_qualifier.(string)
	}
	if Config_state, ok := d.GetOk("config_state"); ok {
		lstorageDasScsiLunAttr.Config_state = Config_state.(string)
	}
	if Deferred_naming, ok := d.GetOk("deferred_naming"); ok {
		lstorageDasScsiLunAttr.Deferred_naming = Deferred_naming.(string)
	}
	if Expand_to_avail, ok := d.GetOk("expand_to_avail"); ok {
		lstorageDasScsiLunAttr.Expand_to_avail = Expand_to_avail.(string)
	}
	if Fractional_size, ok := d.GetOk("fractional_size"); ok {
		lstorageDasScsiLunAttr.Fractional_size = Fractional_size.(string)
	}
	if Local_disk_policy_name, ok := d.GetOk("local_disk_policy_name"); ok {
		lstorageDasScsiLunAttr.Local_disk_policy_name = Local_disk_policy_name.(string)
	}
	if Lun_dn, ok := d.GetOk("lun_dn"); ok {
		lstorageDasScsiLunAttr.Lun_dn = Lun_dn.(string)
	}
	if Lun_map_type, ok := d.GetOk("lun_map_type"); ok {
		lstorageDasScsiLunAttr.Lun_map_type = Lun_map_type.(string)
	}
	if Oper_local_disk_policy_name, ok := d.GetOk("oper_local_disk_policy_name"); ok {
		lstorageDasScsiLunAttr.Oper_local_disk_policy_name = Oper_local_disk_policy_name.(string)
	}
	if Oper_state, ok := d.GetOk("oper_state"); ok {
		lstorageDasScsiLunAttr.Oper_state = Oper_state.(string)
	}
	if Order, ok := d.GetOk("order"); ok {
		lstorageDasScsiLunAttr.Order = Order.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		lstorageDasScsiLunAttr.Sacl = Sacl.(string)
	}
	if Size, ok := d.GetOk("size"); ok {
		lstorageDasScsiLunAttr.Size = Size.(string)
	}
	if Storage_class, ok := d.GetOk("storage_class"); ok {
		lstorageDasScsiLunAttr.Storage_class = Storage_class.(string)
	}

	lstorageDasScsiLun := models.NewLstorageDasScsiLun(fmt.Sprintf("das-scsi-lun-%s", Name), LstorageProfile, desc, lstorageDasScsiLunAttr)
	lstorageDasScsiLun.Status = "modified"
	err := ucsClient.Save(lstorageDasScsiLun)
	if err != nil {
		return err
	}

	d.SetId(lstorageDasScsiLun.DistinguishedName)
	return resourceUcsLstorageDasScsiLunRead(d, m)
}
