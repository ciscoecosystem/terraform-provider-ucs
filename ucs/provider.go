package ucs

import (
	"fmt"

	"github.com/ciscoecosystem/ucs-go-client/client"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ACI_USERNAME", nil),
				Description: "Username for the APIC Account",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ACI_PASSWORD", nil),
				Description: "Password for the APIC Account",
			},
			"url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ACI_URL", nil),
				Description: "URL of the Cisco ACI web interface",
			},
			"insecure": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Allow insecure HTTPS client",
			},
			"proxy_url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ACI_PROXY_URL", nil),
				Description: "Proxy Server URL with port number",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"ucs_mac_pool_pool":            resourceUcsMacpoolPool(),
			"ucs_mac_pool_pool_block":      resourceUcsMacpoolBlock(),
			"ucs_disk_group_config_policy": resourceUcsLstorageDiskGroupConfigPolicy(),
			"ucs_disk_group_qualifier":     resourceUcsLstorageDiskGroupQualifier(),
			"ucs_disk_config_reference":    resourceUcsLstorageLocalDiskConfigRef(),
			"ucs_dns_provider":             resourceUcsCommDnsProvider(),
			"ucs_ip_pool_pool":             resourceUcsIppoolPool(),
			"ucs_ip_pool_block":            resourceUcsIppoolBlock(),
			"ucs_ip_pool_ipv6_block":       resourceUcsIppoolIpV6Block(),
			"ucs_vnic_lan_conn_policy":     resourceUcsVnicLanConnPolicy(),
			"ucs_vnic_ether":               resourceUcsVnicEther(),
			"ucs_vnic_iscsi_lcp":           resourceUcsVnicIScsiLCP(),
			"ucs_vnic_vlan":                resourceUcsVnicVlan(),
			"ucs_ntp_provider":             resourceUcsCommNtpProvider(),
			"ucs_org":                      resourceUcsOrgOrg(),
			"ucs_vnic_san_policy":          resourceUcsVnicSanConnPolicy(),
			"ucs_vnic_fc_if":               resourceUcsVnicFcIf(),
			"ucs_vnic_fc_node":             resourceUcsVnicFcNode(),
			"ucs_vnic_fc":                  resourceUcsVnicFc(),
		},

		// DataSourcesMap: map[string]*schema.Resource{
		// 	"aci_tenant":                            dataSourceAciTenant(),
		// 	"aci_application_profile":               dataSourceAciApplicationProfile(),
		// 	"aci_bridge_domain":                     dataSourceAciBridgeDomain(),
		// 	"aci_contract":                          dataSourceAciContract(),
		// 	"aci_application_epg":                   dataSourceAciApplicationEPG(),
		// 	"aci_contract_subject":                  dataSourceAciContractSubject(),
		// 	"aci_subnet":                            dataSourceAciSubnet(),
		// 	"aci_filter":                            dataSourceAciFilter(),
		// 	"aci_filter_entry":                      dataSourceAciFilterEntry(),
		// 	"aci_vmm_domain":                        dataSourceAciVMMDomain(),
		// 	"aci_vrf":                               dataSourceAciVRF(),
		// 	"aci_external_network_instance_profile": dataSourceAciExternalNetworkInstanceProfile(),
		// 	"aci_l3_outside":                        dataSourceAciL3Outside(),
		// },

		ConfigureFunc: configureClient,
	}
}

func configureClient(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Username:   d.Get("username").(string),
		Password:   d.Get("password").(string),
		URL:        d.Get("url").(string),
		IsInsecure: d.Get("insecure").(bool),
		ProxyUrl:   d.Get("proxy_url").(string),
	}

	if err := config.Valid(); err != nil {
		return nil, err
	}

	return config.getClient(), nil
}

func (c Config) Valid() error {

	if c.Username == "" {
		return fmt.Errorf("Username must be provided for the ACI provider")
	}

	if c.Password == "" {
		return fmt.Errorf("Password must be provided")
	}

	if c.URL == "" {
		return fmt.Errorf("The URL must be provided for the ACI provider")
	}

	return nil
}

func (c Config) getClient() interface{} {

	return client.GetClient(c.URL, c.Username, client.Password(c.Password), client.Insecure(c.IsInsecure), client.ProxyUrl(c.ProxyUrl))

}

// Config
type Config struct {
	Username   string
	Password   string
	URL        string
	IsInsecure bool
	ProxyUrl   string
}
