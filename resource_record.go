package main

import (
    "github.com/hashicorp/terraform/helper/schema"
    "./src/loopia"
)

func resourceRecord() *schema.Resource {
    return &schema.Resource{
        Create: resourceRecordCreate,
        Read:   resourceRecordRead,
        Update: resourceRecordUpdate,
        Delete: resourceRecordDelete,

        Schema: map[string]*schema.Schema{
            "domain": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "value": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "type": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "ttl": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Default:  3600,
            },
            "priority": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Default:  0,
            },
        },
    }
}

func resourceRecordCreate(d *schema.ResourceData, m interface{}) error {
    return resourceRecordRead(d, m)
}

func resourceRecordRead(d *schema.ResourceData, m interface{}) error {
    records := loopia.GetZoneRecords()
    return nil
}

func resourceRecordUpdate(d *schema.ResourceData, m interface{}) error {
    return resourceRecordRead(d, m)
}

func resourceRecordDelete(d *schema.ResourceData, m interface{}) error {
    return nil
}
