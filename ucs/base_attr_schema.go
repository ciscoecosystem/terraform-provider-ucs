package ucs

import "github.com/hashicorp/terraform/helper/schema"

func GetBaseAttrSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// AppendBaseAttrSchema adds the BaseAttr to any schema
func AppendBaseAttrSchema(attrs map[string]*schema.Schema) map[string]*schema.Schema {
	for key, value := range GetBaseAttrSchema() {
		attrs[key] = value
	}

	return attrs
}
