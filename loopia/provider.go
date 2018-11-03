package loopia

import (
    "github.com/hashicorp/terraform/helper/schema"
)

func Provider() terraform.ResourceProvider {
    return &schema.Provider{
        Schema: map[string]*schema.Schema{
            "username": &schema.Schema{
              Type:        schema.TypeString,
              Required:    true,
              DefaultFunc: schema.EnvDefaultFunc("LOOPIA_USERNAME", nil),
            },

            "password": &schema.Schema{
              Type:        schema.TypeString,
              Required:    true,
              DefaultFunc: schema.EnvDefaultFunc("LOOPIA_PASSWORD", nil),
            },
        },
        ResourcesMap: map[string]*schema.Resource{
            "loopia_record": resourceLoopiaRecord(),
        },
    }
}
