package loopia

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jonlil/loopia-go"
	"strconv"
	"strings"
)

func resourceLoopiaRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceLoopiaRecordCreate,
		Read:   resourceLoopiaRecordRead,
		Update: resourceLoopiaRecordUpdate,
		Delete: resourceLoopiaRecordDelete,
		Importer: &schema.ResourceImporter{
			State: resourceLoopiaeRecordImport,
		},

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

func resourceLoopiaRecordCreate(d *schema.ResourceData, m interface{}) error {
	return resourceLoopiaRecordRead(d, m)
}

func resourceLoopiaRecordRead(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*loopia.API)
	return nil
	// record, err := client.GetZoneRecord()
	// if err != nil {
	//   return nil, fmt.Errorf("Unable to find record with ID %q: %q", d.Id())
	// }
	// return nil
}

func resourceLoopiaRecordUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceLoopiaRecordRead(d, m)
}

func resourceLoopiaRecordDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceLoopiaeRecordImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*loopia.API)
	idAttr := strings.SplitN(d.Id(), "/", 3)
	var domain string
	var subdomain string

	id, _ := strconv.ParseInt(idAttr[2], 10, 0)
	domain = idAttr[1]
	subdomain = idAttr[0]

	record, err := client.GetZoneRecord(domain, subdomain, id)

	if err != nil {
		return nil, fmt.Errorf("Unable to find record with ID %q: %q", id, err)
	}

	d.Set("name", subdomain)
	d.Set("domain", domain)
	d.Set("priority", record.Priority)
	d.Set("value", record.Value)
	d.Set("type", record.Type)
	d.Set("ttl", record.TTL)
	d.SetId(idAttr[2])

	return []*schema.ResourceData{d}, nil
}
