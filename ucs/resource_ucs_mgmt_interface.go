package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/ciscoecosystem/ucs-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUcsMgmtInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceUcsMgmtInterfaceCreate,
		Update: resourceUcsMgmtInterfaceUpdate,
		Read:   resourceUcsMgmtInterfaceRead,
		Delete: resourceUcsMgmtInterfaceDelete,

		Importer: &schema.ResourceImporter{
			State: resourceUcsMgmtInterfaceImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"ls_server_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"child_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"config_message": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"config_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"ip_v4_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"ip_v6_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"is_default_derived": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"monitor_interval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"oper_state": &schema.Schema{
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

func getRemoteMgmtInterface(client *client.Client, dn string) (*models.MgmtInterface, error) {
	mgmtInterfaceDoc, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	mgmtInterface := models.MgmtInterfaceFromDoc(mgmtInterfaceDoc, "configResolveDn")

	if mgmtInterface.DistinguishedName == "" {
		return nil, fmt.Errorf("MgmtInterface %s not found", dn)
	}

	return mgmtInterface, nil
}

func setMgmtInterfaceAttributes(mgmtInterface *models.MgmtInterface, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(mgmtInterface.DistinguishedName)
	d.Set("description", mgmtInterface.Description)
	d.Set("ls_server_dn", GetParentDn(mgmtInterface.DistinguishedName))
	mgmtInterfaceMap, _ := mgmtInterface.ToMap()

	d.Set("child_action", mgmtInterfaceMap["childAction"])

	d.Set("config_message", mgmtInterfaceMap["configMessage"])

	d.Set("config_state", mgmtInterfaceMap["configState"])

	d.Set("ip_v4_state", mgmtInterfaceMap["ipV4State"])

	d.Set("ip_v6_state", mgmtInterfaceMap["ipV6State"])

	d.Set("is_default_derived", mgmtInterfaceMap["isDefaultDerived"])

	d.Set("monitor_interval", mgmtInterfaceMap["monitorInterval"])

	d.Set("oper_state", mgmtInterfaceMap["operState"])

	d.Set("sacl", mgmtInterfaceMap["sacl"])
	return d
}

func resourceUcsMgmtInterfaceImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	ucsClient := m.(*client.Client)

	dn := d.Id()

	mgmtInterface, err := getRemoteMgmtInterface(ucsClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setMgmtInterfaceAttributes(mgmtInterface, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceUcsMgmtInterfaceCreate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Mode := d.Get("mode").(string)

	LsServer := d.Get("ls_server_dn").(string)

	mgmtInterfaceAttr := models.MgmtInterfaceAttributes{}

	if Child_action, ok := d.GetOk("child_action"); ok {
		mgmtInterfaceAttr.Child_action = Child_action.(string)
	}

	if Config_message, ok := d.GetOk("config_message"); ok {
		mgmtInterfaceAttr.Config_message = Config_message.(string)
	}

	if Config_state, ok := d.GetOk("config_state"); ok {
		mgmtInterfaceAttr.Config_state = Config_state.(string)
	}

	if Ip_v4_state, ok := d.GetOk("ip_v4_state"); ok {
		mgmtInterfaceAttr.Ip_v4_state = Ip_v4_state.(string)
	}

	if Ip_v6_state, ok := d.GetOk("ip_v6_state"); ok {
		mgmtInterfaceAttr.Ip_v6_state = Ip_v6_state.(string)
	}

	if Is_default_derived, ok := d.GetOk("is_default_derived"); ok {
		mgmtInterfaceAttr.Is_default_derived = Is_default_derived.(string)
	}

	if Monitor_interval, ok := d.GetOk("monitor_interval"); ok {
		mgmtInterfaceAttr.Monitor_interval = Monitor_interval.(string)
	}

	if Oper_state, ok := d.GetOk("oper_state"); ok {
		mgmtInterfaceAttr.Oper_state = Oper_state.(string)
	}

	if Sacl, ok := d.GetOk("sacl"); ok {
		mgmtInterfaceAttr.Sacl = Sacl.(string)
	}

	mgmtInterface := models.NewMgmtInterface(fmt.Sprintf("iface-%s", Mode), LsServer, desc, mgmtInterfaceAttr)

	err := ucsClient.Save(mgmtInterface)
	if err != nil {
		return err
	}

	d.SetId(mgmtInterface.DistinguishedName)
	return resourceUcsMgmtInterfaceRead(d, m)
}

func resourceUcsMgmtInterfaceRead(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)

	dn := d.Id()
	mgmtInterface, err := getRemoteMgmtInterface(ucsClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setMgmtInterfaceAttributes(mgmtInterface, d)

	return nil
}

func resourceUcsMgmtInterfaceDelete(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	dn := d.Id()
	err := ucsClient.DeleteByDn(dn, "mgmtInterface")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}

func resourceUcsMgmtInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
	ucsClient := m.(*client.Client)
	desc := d.Get("description").(string)

	Mode := d.Get("mode").(string)

	LsServer := d.Get("ls_server_dn").(string)

	mgmtInterfaceAttr := models.MgmtInterfaceAttributes{}
	if Child_action, ok := d.GetOk("child_action"); ok {
		mgmtInterfaceAttr.Child_action = Child_action.(string)
	}
	if Config_message, ok := d.GetOk("config_message"); ok {
		mgmtInterfaceAttr.Config_message = Config_message.(string)
	}
	if Config_state, ok := d.GetOk("config_state"); ok {
		mgmtInterfaceAttr.Config_state = Config_state.(string)
	}
	if Ip_v4_state, ok := d.GetOk("ip_v4_state"); ok {
		mgmtInterfaceAttr.Ip_v4_state = Ip_v4_state.(string)
	}
	if Ip_v6_state, ok := d.GetOk("ip_v6_state"); ok {
		mgmtInterfaceAttr.Ip_v6_state = Ip_v6_state.(string)
	}
	if Is_default_derived, ok := d.GetOk("is_default_derived"); ok {
		mgmtInterfaceAttr.Is_default_derived = Is_default_derived.(string)
	}
	if Monitor_interval, ok := d.GetOk("monitor_interval"); ok {
		mgmtInterfaceAttr.Monitor_interval = Monitor_interval.(string)
	}
	if Oper_state, ok := d.GetOk("oper_state"); ok {
		mgmtInterfaceAttr.Oper_state = Oper_state.(string)
	}
	if Sacl, ok := d.GetOk("sacl"); ok {
		mgmtInterfaceAttr.Sacl = Sacl.(string)
	}

	mgmtInterface := models.NewMgmtInterface(fmt.Sprintf("iface-%s", Mode), LsServer, desc, mgmtInterfaceAttr)
	mgmtInterface.Status = "modified"
	err := ucsClient.Save(mgmtInterface)
	if err != nil {
		return err
	}

	d.SetId(mgmtInterface.DistinguishedName)
	return resourceUcsMgmtInterfaceRead(d, m)
}
