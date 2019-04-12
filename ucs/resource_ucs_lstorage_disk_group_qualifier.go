package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsLstorageDiskGroupQualifier() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsLstorageDiskGroupQualifierCreate,
		Update: resourceUcsLstorageDiskGroupQualifierUpdate,
		Read:   resourceUcsLstorageDiskGroupQualifierRead,
		Delete: resourceUcsLstorageDiskGroupQualifierDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsLstorageDiskGroupQualifierImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"lstorage_disk_group_config_policy_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"child_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"drive_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"min_drive_size": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"num_ded_hot_spares": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"num_drives": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"num_glob_hot_spares": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"sacl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"use_jbod_disks": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"use_remaining_disks": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		}),
	}
}

func getRemoteLstorageDiskGroupQualifier(client *client.Client, dn string) (*models.LstorageDiskGroupQualifier, error) {
	lstorageDiskGroupQualifierDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	lstorageDiskGroupQualifier := models.LstorageDiskGroupQualifierFromDoc(lstorageDiskGroupQualifierDoc, "configResolveDn")

	if lstorageDiskGroupQualifier.DistinguishedName == "" {
		return nil, fmt.Errorf("LstorageDiskGroupQualifier %s not found", dn)
	}

	return lstorageDiskGroupQualifier, nil
}

func setLstorageDiskGroupQualifierAttributes(lstorageDiskGroupQualifier *models.LstorageDiskGroupQualifier, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(lstorageDiskGroupQualifier.DistinguishedName)
	d.Set("description", lstorageDiskGroupQualifier.Description)
	d.Set("lstorage_disk_group_config_policy_dn", GetParentDn(lstorageDiskGroupQualifier.DistinguishedName))
	lstorageDiskGroupQualifierMap, _ := lstorageDiskGroupQualifier.ToMap()

	d.Set("child_action", lstorageDiskGroupQualifierMap["childAction"])

	d.Set("drive_type", lstorageDiskGroupQualifierMap["driveType"])

	d.Set("min_drive_size", lstorageDiskGroupQualifierMap["minDriveSize"])

	d.Set("num_ded_hot_spares", lstorageDiskGroupQualifierMap["numDedHotSpares"])

	d.Set("num_drives", lstorageDiskGroupQualifierMap["numDrives"])

	d.Set("num_glob_hot_spares", lstorageDiskGroupQualifierMap["numGlobHotSpares"])

	d.Set("sacl", lstorageDiskGroupQualifierMap["sacl"])

	d.Set("use_jbod_disks", lstorageDiskGroupQualifierMap["useJbodDisks"])

	d.Set("use_remaining_disks", lstorageDiskGroupQualifierMap["useRemainingDisks"])
	return d
}

func resourceUcsLstorageDiskGroupQualifierImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	lstorageDiskGroupQualifier, err := getRemoteLstorageDiskGroupQualifier(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setLstorageDiskGroupQualifierAttributes(lstorageDiskGroupQualifier, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsLstorageDiskGroupQualifierCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	LstorageDiskGroupConfigPolicy := d.Get("lstorage_disk_group_config_policy_dn").(string)

	lstorageDiskGroupQualifierAttr := models.LstorageDiskGroupQualifierAttributes{}

	if Child_action, ok := d.GetOk("child_action"); ok {
		lstorageDiskGroupQualifierAttr.Child_action = Child_action.(string)
	}

	if Drive_type, ok := d.GetOk("drive_type"); ok {
		lstorageDiskGroupQualifierAttr.Drive_type = Drive_type.(string)
	}

	if Min_drive_size, ok := d.GetOk("min_drive_size"); ok {
		lstorageDiskGroupQualifierAttr.Min_drive_size = Min_drive_size.(string)
	}

	if Num_ded_hot_spares, ok := d.GetOk("num_ded_hot_spares"); ok {
		lstorageDiskGroupQualifierAttr.Num_ded_hot_spares = Num_ded_hot_spares.(string)
	}

	if Num_drives, ok := d.GetOk("num_drives"); ok {
		lstorageDiskGroupQualifierAttr.Num_drives = Num_drives.(string)
	}

	if Num_glob_hot_spares, ok := d.GetOk("num_glob_hot_spares"); ok {
		lstorageDiskGroupQualifierAttr.Num_glob_hot_spares = Num_glob_hot_spares.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		lstorageDiskGroupQualifierAttr.Sacl = Sacl.(string)
	}

	if Use_jbod_disks, ok := d.GetOk("use_jbod_disks"); ok {
		lstorageDiskGroupQualifierAttr.Use_jbod_disks = Use_jbod_disks.(string)
	}

	if Use_remaining_disks, ok := d.GetOk("use_remaining_disks"); ok {
		lstorageDiskGroupQualifierAttr.Use_remaining_disks = Use_remaining_disks.(string)
	}

	lstorageDiskGroupQualifier := models.NewLstorageDiskGroupQualifier(fmt.Sprintf("disk-group-qual"), LstorageDiskGroupConfigPolicy, desc, lstorageDiskGroupQualifierAttr)

	err := ucsClient.Save(lstorageDiskGroupQualifier)
	if err != nil {
		return err
	}

	d.SetId(lstorageDiskGroupQualifier.DistinguishedName)
	return resourceUcsLstorageDiskGroupQualifierRead(d, m)
}

func resourceUcsLstorageDiskGroupQualifierRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	lstorageDiskGroupQualifier, err := getRemoteLstorageDiskGroupQualifier(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setLstorageDiskGroupQualifierAttributes(lstorageDiskGroupQualifier, d)

	return nil
}

func resourceUcsLstorageDiskGroupQualifierDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "lstorageDiskGroupQualifier")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsLstorageDiskGroupQualifierUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	LstorageDiskGroupConfigPolicy := d.Get("lstorage_disk_group_config_policy_dn").(string)

	lstorageDiskGroupQualifierAttr := models.LstorageDiskGroupQualifierAttributes{}
	if Child_action, ok := d.GetOk("child_action"); ok {
		lstorageDiskGroupQualifierAttr.Child_action = Child_action.(string)
	}
	if Drive_type, ok := d.GetOk("drive_type"); ok {
		lstorageDiskGroupQualifierAttr.Drive_type = Drive_type.(string)
	}
	if Min_drive_size, ok := d.GetOk("min_drive_size"); ok {
		lstorageDiskGroupQualifierAttr.Min_drive_size = Min_drive_size.(string)
	}
	if Num_ded_hot_spares, ok := d.GetOk("num_ded_hot_spares"); ok {
		lstorageDiskGroupQualifierAttr.Num_ded_hot_spares = Num_ded_hot_spares.(string)
	}
	if Num_drives, ok := d.GetOk("num_drives"); ok {
		lstorageDiskGroupQualifierAttr.Num_drives = Num_drives.(string)
	}
	if Num_glob_hot_spares, ok := d.GetOk("num_glob_hot_spares"); ok {
		lstorageDiskGroupQualifierAttr.Num_glob_hot_spares = Num_glob_hot_spares.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		lstorageDiskGroupQualifierAttr.Sacl = Sacl.(string)
	}
	if Use_jbod_disks, ok := d.GetOk("use_jbod_disks"); ok {
		lstorageDiskGroupQualifierAttr.Use_jbod_disks = Use_jbod_disks.(string)
	}
	if Use_remaining_disks, ok := d.GetOk("use_remaining_disks"); ok {
		lstorageDiskGroupQualifierAttr.Use_remaining_disks = Use_remaining_disks.(string)
	}

	lstorageDiskGroupQualifier := models.NewLstorageDiskGroupQualifier(fmt.Sprintf("disk-group-qual"), LstorageDiskGroupConfigPolicy, desc, lstorageDiskGroupQualifierAttr)
	lstorageDiskGroupQualifier.Status = "modified"
	err := ucsClient.Save(lstorageDiskGroupQualifier)
	if err != nil {
		return err
	}

	d.SetId(lstorageDiskGroupQualifier.DistinguishedName)
	return resourceUcsLstorageDiskGroupQualifierRead(d, m)
}
