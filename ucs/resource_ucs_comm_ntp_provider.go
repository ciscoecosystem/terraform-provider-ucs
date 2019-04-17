package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsCommNtpProvider() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsCommNtpProviderCreate,
		Update: resourceUcsCommNtpProviderUpdate,
		Read:   resourceUcsCommNtpProviderRead,
		Delete: resourceUcsCommNtpProviderDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsCommNtpProviderImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"admin_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"child_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"hostname": &schema.Schema{
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

func getRemoteCommNtpProvider(client *client.Client, dn string) (*models.CommNtpProvider, error) {
	commNtpProviderDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	commNtpProvider := models.CommNtpProviderFromDoc(commNtpProviderDoc, "configResolveDn")

	if commNtpProvider.DistinguishedName == "" {
		return nil, fmt.Errorf("CommNtpProvider %s not found", dn)
	}

	return commNtpProvider, nil
}

func setCommNtpProviderAttributes(commNtpProvider *models.CommNtpProvider, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(commNtpProvider.DistinguishedName)
	d.Set("description", commNtpProvider.Description)
	commNtpProviderMap, _ := commNtpProvider.ToMap()

	d.Set("admin_state", commNtpProviderMap["adminState"])

	d.Set("child_action", commNtpProviderMap["childAction"])

	d.Set("hostname", commNtpProviderMap["hostname"])

	d.Set("sacl", commNtpProviderMap["sacl"])
	return d
}

func resourceUcsCommNtpProviderImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	commNtpProvider, err := getRemoteCommNtpProvider(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setCommNtpProviderAttributes(commNtpProvider, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsCommNtpProviderCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	commNtpProviderAttr := models.CommNtpProviderAttributes{}

	if Admin_state, ok := d.GetOk("admin_state"); ok {
		commNtpProviderAttr.Admin_state = Admin_state.(string)
	}

	if Child_action, ok := d.GetOk("child_action"); ok {
		commNtpProviderAttr.Child_action = Child_action.(string)
	}

	if Hostname, ok := d.GetOk("hostname"); ok {
		commNtpProviderAttr.Hostname = Hostname.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		commNtpProviderAttr.Sacl = Sacl.(string)
	}

	commNtpProvider := models.NewCommNtpProvider(fmt.Sprintf("sys/svc-ext/datetime-svc/ntp-%s", Name), desc, commNtpProviderAttr)

	err := ucsClient.Save(commNtpProvider)
	if err != nil {
		return err
	}

	d.SetId(commNtpProvider.DistinguishedName)
	return resourceUcsCommNtpProviderRead(d, m)
}

func resourceUcsCommNtpProviderRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	commNtpProvider, err := getRemoteCommNtpProvider(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setCommNtpProviderAttributes(commNtpProvider, d)

	return nil
}

func resourceUcsCommNtpProviderDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "commNtpProvider")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsCommNtpProviderUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	commNtpProviderAttr := models.CommNtpProviderAttributes{}
	if Admin_state, ok := d.GetOk("admin_state"); ok {
		commNtpProviderAttr.Admin_state = Admin_state.(string)
	}
	if Child_action, ok := d.GetOk("child_action"); ok {
		commNtpProviderAttr.Child_action = Child_action.(string)
	}
	if Hostname, ok := d.GetOk("hostname"); ok {
		commNtpProviderAttr.Hostname = Hostname.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		commNtpProviderAttr.Sacl = Sacl.(string)
	}

	commNtpProvider := models.NewCommNtpProvider(fmt.Sprintf("sys/svc-ext/datetime-svc/ntp-%s", Name), desc, commNtpProviderAttr)
	commNtpProvider.Status = "modified"
	err := ucsClient.Save(commNtpProvider)
	if err != nil {
		return err
	}

	d.SetId(commNtpProvider.DistinguishedName)
	return resourceUcsCommNtpProviderRead(d, m)
}
