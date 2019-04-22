package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsLstorageProfileBinding() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsLstorageProfileBindingCreate,
		Update: resourceUcsLstorageProfileBindingUpdate,
		Read:   resourceUcsLstorageProfileBindingRead,
		Delete: resourceUcsLstorageProfileBindingDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsLstorageProfileBindingImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"ls_server_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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

			"issues": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_storage_profile_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"sacl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"storage_profile_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		}),
	}
}

func getRemoteLstorageProfileBinding(client *client.Client, dn string) (*models.LstorageProfileBinding, error) {
	lstorageProfileBindingDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	lstorageProfileBinding := models.LstorageProfileBindingFromDoc(lstorageProfileBindingDoc, "configResolveDn")

	if lstorageProfileBinding.DistinguishedName == "" {
		return nil, fmt.Errorf("LstorageProfileBinding %s not found", dn)
	}

	return lstorageProfileBinding, nil
}

func setLstorageProfileBindingAttributes(lstorageProfileBinding *models.LstorageProfileBinding, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(lstorageProfileBinding.DistinguishedName)
	d.Set("description", lstorageProfileBinding.Description)
	d.Set("ls_server_dn", GetParentDn(lstorageProfileBinding.DistinguishedName))
	lstorageProfileBindingMap, _ := lstorageProfileBinding.ToMap()

	d.Set("assigned_to_dn", lstorageProfileBindingMap["assignedToDn"])

	d.Set("child_action", lstorageProfileBindingMap["childAction"])

	d.Set("issues", lstorageProfileBindingMap["issues"])

	d.Set("name", lstorageProfileBindingMap["name"])

	d.Set("oper_storage_profile_name", lstorageProfileBindingMap["operStorageProfileName"])

	d.Set("sacl", lstorageProfileBindingMap["sacl"])

	d.Set("storage_profile_name", lstorageProfileBindingMap["storageProfileName"])
	return d
}

func resourceUcsLstorageProfileBindingImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	lstorageProfileBinding, err := getRemoteLstorageProfileBinding(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setLstorageProfileBindingAttributes(lstorageProfileBinding, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsLstorageProfileBindingCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	LsServer := d.Get("ls_server_dn").(string)

	lstorageProfileBindingAttr := models.LstorageProfileBindingAttributes{}

	if Assigned_to_dn, ok := d.GetOk("assigned_to_dn"); ok {
		lstorageProfileBindingAttr.Assigned_to_dn = Assigned_to_dn.(string)
	}

	if Child_action, ok := d.GetOk("child_action"); ok {
		lstorageProfileBindingAttr.Child_action = Child_action.(string)
	}

	if Issues, ok := d.GetOk("issues"); ok {
		lstorageProfileBindingAttr.Issues = Issues.(string)
	}

	if Oper_storage_profile_name, ok := d.GetOk("oper_storage_profile_name"); ok {
		lstorageProfileBindingAttr.Oper_storage_profile_name = Oper_storage_profile_name.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		lstorageProfileBindingAttr.Sacl = Sacl.(string)
	}

	if Storage_profile_name, ok := d.GetOk("storage_profile_name"); ok {
		lstorageProfileBindingAttr.Storage_profile_name = Storage_profile_name.(string)
	}

	lstorageProfileBinding := models.NewLstorageProfileBinding(fmt.Sprintf("profile-binding"), LsServer, desc, lstorageProfileBindingAttr)

	err := ucsClient.Save(lstorageProfileBinding)
	if err != nil {
		return err
	}

	d.SetId(lstorageProfileBinding.DistinguishedName)
	return resourceUcsLstorageProfileBindingRead(d, m)
}

func resourceUcsLstorageProfileBindingRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	lstorageProfileBinding, err := getRemoteLstorageProfileBinding(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setLstorageProfileBindingAttributes(lstorageProfileBinding, d)

	return nil
}

func resourceUcsLstorageProfileBindingDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "lstorageProfileBinding")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsLstorageProfileBindingUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	LsServer := d.Get("ls_server_dn").(string)

	lstorageProfileBindingAttr := models.LstorageProfileBindingAttributes{}
	if Assigned_to_dn, ok := d.GetOk("assigned_to_dn"); ok {
		lstorageProfileBindingAttr.Assigned_to_dn = Assigned_to_dn.(string)
	}
	if Child_action, ok := d.GetOk("child_action"); ok {
		lstorageProfileBindingAttr.Child_action = Child_action.(string)
	}
	if Issues, ok := d.GetOk("issues"); ok {
		lstorageProfileBindingAttr.Issues = Issues.(string)
	}
	if Oper_storage_profile_name, ok := d.GetOk("oper_storage_profile_name"); ok {
		lstorageProfileBindingAttr.Oper_storage_profile_name = Oper_storage_profile_name.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		lstorageProfileBindingAttr.Sacl = Sacl.(string)
	}
	if Storage_profile_name, ok := d.GetOk("storage_profile_name"); ok {
		lstorageProfileBindingAttr.Storage_profile_name = Storage_profile_name.(string)
	}

	lstorageProfileBinding := models.NewLstorageProfileBinding(fmt.Sprintf("profile-binding"), LsServer, desc, lstorageProfileBindingAttr)
	lstorageProfileBinding.Status = "modified"
	err := ucsClient.Save(lstorageProfileBinding)
	if err != nil {
		return err
	}

	d.SetId(lstorageProfileBinding.DistinguishedName)
	return resourceUcsLstorageProfileBindingRead(d, m)
}
