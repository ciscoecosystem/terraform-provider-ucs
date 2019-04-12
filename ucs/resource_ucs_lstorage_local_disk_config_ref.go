package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsLstorageLocalDiskConfigRef() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsLstorageLocalDiskConfigRefCreate,
		Update: resourceUcsLstorageLocalDiskConfigRefUpdate,
		Read:   resourceUcsLstorageLocalDiskConfigRefRead,
		Delete: resourceUcsLstorageLocalDiskConfigRefDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsLstorageLocalDiskConfigRefImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"lstorage_disk_group_config_policy_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"slot_num": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"child_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"sacl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"span_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		}),
	}
}

func getRemoteLstorageLocalDiskConfigRef(client *client.Client, dn string) (*models.LstorageLocalDiskConfigRef, error) {
	lstorageLocalDiskConfigRefDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	lstorageLocalDiskConfigRef := models.LstorageLocalDiskConfigRefFromDoc(lstorageLocalDiskConfigRefDoc, "configResolveDn")

	if lstorageLocalDiskConfigRef.DistinguishedName == "" {
		return nil, fmt.Errorf("LstorageLocalDiskConfigRef %s not found", dn)
	}

	return lstorageLocalDiskConfigRef, nil
}

func setLstorageLocalDiskConfigRefAttributes(lstorageLocalDiskConfigRef *models.LstorageLocalDiskConfigRef, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(lstorageLocalDiskConfigRef.DistinguishedName)
	d.Set("description", lstorageLocalDiskConfigRef.Description)
	d.Set("lstorage_disk_group_config_policy_dn", GetParentDn(lstorageLocalDiskConfigRef.DistinguishedName))
	lstorageLocalDiskConfigRefMap, _ := lstorageLocalDiskConfigRef.ToMap()

	d.Set("child_action", lstorageLocalDiskConfigRefMap["childAction"])

	d.Set("role", lstorageLocalDiskConfigRefMap["role"])

	d.Set("sacl", lstorageLocalDiskConfigRefMap["sacl"])

	d.Set("span_id", lstorageLocalDiskConfigRefMap["spanId"])
	return d
}

func resourceUcsLstorageLocalDiskConfigRefImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	lstorageLocalDiskConfigRef, err := getRemoteLstorageLocalDiskConfigRef(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setLstorageLocalDiskConfigRefAttributes(lstorageLocalDiskConfigRef, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsLstorageLocalDiskConfigRefCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Slot_num := d.Get("slot_num").(string)

	LstorageDiskGroupConfigPolicy := d.Get("lstorage_disk_group_config_policy_dn").(string)

	lstorageLocalDiskConfigRefAttr := models.LstorageLocalDiskConfigRefAttributes{}

	if Child_action, ok := d.GetOk("child_action"); ok {
		lstorageLocalDiskConfigRefAttr.Child_action = Child_action.(string)
	}

	if Role, ok := d.GetOk("role"); ok {
		lstorageLocalDiskConfigRefAttr.Role = Role.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		lstorageLocalDiskConfigRefAttr.Sacl = Sacl.(string)
	}

	if Span_id, ok := d.GetOk("span_id"); ok {
		lstorageLocalDiskConfigRefAttr.Span_id = Span_id.(string)
	}

	lstorageLocalDiskConfigRef := models.NewLstorageLocalDiskConfigRef(fmt.Sprintf("slot-%s", Slot_num), LstorageDiskGroupConfigPolicy, desc, lstorageLocalDiskConfigRefAttr)

	err := ucsClient.Save(lstorageLocalDiskConfigRef)
	if err != nil {
		return err
	}

	d.SetId(lstorageLocalDiskConfigRef.DistinguishedName)
	return resourceUcsLstorageLocalDiskConfigRefRead(d, m)
}

func resourceUcsLstorageLocalDiskConfigRefRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	lstorageLocalDiskConfigRef, err := getRemoteLstorageLocalDiskConfigRef(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setLstorageLocalDiskConfigRefAttributes(lstorageLocalDiskConfigRef, d)

	return nil
}

func resourceUcsLstorageLocalDiskConfigRefDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "lstorageLocalDiskConfigRef")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsLstorageLocalDiskConfigRefUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Slot_num := d.Get("slot_num").(string)

	LstorageDiskGroupConfigPolicy := d.Get("lstorage_disk_group_config_policy_dn").(string)

	lstorageLocalDiskConfigRefAttr := models.LstorageLocalDiskConfigRefAttributes{}
	if Child_action, ok := d.GetOk("child_action"); ok {
		lstorageLocalDiskConfigRefAttr.Child_action = Child_action.(string)
	}
	if Role, ok := d.GetOk("role"); ok {
		lstorageLocalDiskConfigRefAttr.Role = Role.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		lstorageLocalDiskConfigRefAttr.Sacl = Sacl.(string)
	}
	if Span_id, ok := d.GetOk("span_id"); ok {
		lstorageLocalDiskConfigRefAttr.Span_id = Span_id.(string)
	}

	lstorageLocalDiskConfigRef := models.NewLstorageLocalDiskConfigRef(fmt.Sprintf("slot-%s", Slot_num), LstorageDiskGroupConfigPolicy, desc, lstorageLocalDiskConfigRefAttr)
	lstorageLocalDiskConfigRef.Status = "modified"
	err := ucsClient.Save(lstorageLocalDiskConfigRef)
	if err != nil {
		return err
	}

	d.SetId(lstorageLocalDiskConfigRef.DistinguishedName)
	return resourceUcsLstorageLocalDiskConfigRefRead(d, m)
}
