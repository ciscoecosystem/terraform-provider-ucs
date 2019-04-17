package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsCommDnsProvider() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsCommDnsProviderCreate,
		Update: resourceUcsCommDnsProviderUpdate,
		Read:   resourceUcsCommDnsProviderRead,
		Delete: resourceUcsCommDnsProviderDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsCommDnsProviderImport,
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

func getRemoteCommDnsProvider(client *client.Client, dn string) (*models.CommDnsProvider, error) {
	commDnsProviderDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	commDnsProvider := models.CommDnsProviderFromDoc(commDnsProviderDoc, "configResolveDn")

	if commDnsProvider.DistinguishedName == "" {
		return nil, fmt.Errorf("CommDnsProvider %s not found", dn)
	}

	return commDnsProvider, nil
}

func setCommDnsProviderAttributes(commDnsProvider *models.CommDnsProvider, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(commDnsProvider.DistinguishedName)
	d.Set("description", commDnsProvider.Description)
	commDnsProviderMap, _ := commDnsProvider.ToMap()

	d.Set("admin_state", commDnsProviderMap["adminState"])

	d.Set("child_action", commDnsProviderMap["childAction"])

	d.Set("hostname", commDnsProviderMap["hostname"])

	d.Set("sacl", commDnsProviderMap["sacl"])
	return d
}

func resourceUcsCommDnsProviderImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	commDnsProvider, err := getRemoteCommDnsProvider(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setCommDnsProviderAttributes(commDnsProvider, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsCommDnsProviderCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	commDnsProviderAttr := models.CommDnsProviderAttributes{}

	if Admin_state, ok := d.GetOk("admin_state"); ok {
		commDnsProviderAttr.Admin_state = Admin_state.(string)
	}

	if Child_action, ok := d.GetOk("child_action"); ok {
		commDnsProviderAttr.Child_action = Child_action.(string)
	}

	if Hostname, ok := d.GetOk("hostname"); ok {
		commDnsProviderAttr.Hostname = Hostname.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		commDnsProviderAttr.Sacl = Sacl.(string)
	}

	commDnsProvider := models.NewCommDnsProvider(fmt.Sprintf("sys/svc-ext/dns-svc/dns-%s", Name), desc, commDnsProviderAttr)

	err := ucsClient.Save(commDnsProvider)
	if err != nil {
		return err
	}

	d.SetId(commDnsProvider.DistinguishedName)
	return resourceUcsCommDnsProviderRead(d, m)
}

func resourceUcsCommDnsProviderRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	commDnsProvider, err := getRemoteCommDnsProvider(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setCommDnsProviderAttributes(commDnsProvider, d)

	return nil
}

func resourceUcsCommDnsProviderDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "commDnsProvider")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsCommDnsProviderUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Name := d.Get("name").(string)

	commDnsProviderAttr := models.CommDnsProviderAttributes{}
	if Admin_state, ok := d.GetOk("admin_state"); ok {
		commDnsProviderAttr.Admin_state = Admin_state.(string)
	}
	if Child_action, ok := d.GetOk("child_action"); ok {
		commDnsProviderAttr.Child_action = Child_action.(string)
	}
	if Hostname, ok := d.GetOk("hostname"); ok {
		commDnsProviderAttr.Hostname = Hostname.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		commDnsProviderAttr.Sacl = Sacl.(string)
	}

	commDnsProvider := models.NewCommDnsProvider(fmt.Sprintf("sys/svc-ext/dns-svc/dns-%s", Name), desc, commDnsProviderAttr)
	commDnsProvider.Status = "modified"
	err := ucsClient.Save(commDnsProvider)
	if err != nil {
		return err
	}

	d.SetId(commDnsProvider.DistinguishedName)
	return resourceUcsCommDnsProviderRead(d, m)
}
