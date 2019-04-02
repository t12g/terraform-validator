// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"google.golang.org/api/compute/v1"
)

func resourceComputeHttpsHealthCheck() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeHttpsHealthCheckCreate,
		Read:   resourceComputeHttpsHealthCheckRead,
		Update: resourceComputeHttpsHealthCheckUpdate,
		Delete: resourceComputeHttpsHealthCheckDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeHttpsHealthCheckImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(240 * time.Second),
			Update: schema.DefaultTimeout(240 * time.Second),
			Delete: schema.DefaultTimeout(240 * time.Second),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"check_interval_sec": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"healthy_threshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2,
			},
			"host": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  443,
			},
			"request_path": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/",
			},
			"timeout_sec": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"unhealthy_threshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2,
			},
			"creation_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceComputeHttpsHealthCheckCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	checkIntervalSecProp, err := expandComputeHttpsHealthCheckCheckIntervalSec(d.Get("check_interval_sec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("check_interval_sec"); !isEmptyValue(reflect.ValueOf(checkIntervalSecProp)) && (ok || !reflect.DeepEqual(v, checkIntervalSecProp)) {
		obj["checkIntervalSec"] = checkIntervalSecProp
	}
	descriptionProp, err := expandComputeHttpsHealthCheckDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	healthyThresholdProp, err := expandComputeHttpsHealthCheckHealthyThreshold(d.Get("healthy_threshold"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("healthy_threshold"); !isEmptyValue(reflect.ValueOf(healthyThresholdProp)) && (ok || !reflect.DeepEqual(v, healthyThresholdProp)) {
		obj["healthyThreshold"] = healthyThresholdProp
	}
	hostProp, err := expandComputeHttpsHealthCheckHost(d.Get("host"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("host"); !isEmptyValue(reflect.ValueOf(hostProp)) && (ok || !reflect.DeepEqual(v, hostProp)) {
		obj["host"] = hostProp
	}
	nameProp, err := expandComputeHttpsHealthCheckName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	portProp, err := expandComputeHttpsHealthCheckPort(d.Get("port"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("port"); !isEmptyValue(reflect.ValueOf(portProp)) && (ok || !reflect.DeepEqual(v, portProp)) {
		obj["port"] = portProp
	}
	requestPathProp, err := expandComputeHttpsHealthCheckRequestPath(d.Get("request_path"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("request_path"); !isEmptyValue(reflect.ValueOf(requestPathProp)) && (ok || !reflect.DeepEqual(v, requestPathProp)) {
		obj["requestPath"] = requestPathProp
	}
	timeoutSecProp, err := expandComputeHttpsHealthCheckTimeoutSec(d.Get("timeout_sec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("timeout_sec"); !isEmptyValue(reflect.ValueOf(timeoutSecProp)) && (ok || !reflect.DeepEqual(v, timeoutSecProp)) {
		obj["timeoutSec"] = timeoutSecProp
	}
	unhealthyThresholdProp, err := expandComputeHttpsHealthCheckUnhealthyThreshold(d.Get("unhealthy_threshold"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("unhealthy_threshold"); !isEmptyValue(reflect.ValueOf(unhealthyThresholdProp)) && (ok || !reflect.DeepEqual(v, unhealthyThresholdProp)) {
		obj["unhealthyThreshold"] = unhealthyThresholdProp
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/httpsHealthChecks")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new HttpsHealthCheck: %#v", obj)
	res, err := sendRequestWithTimeout(config, "POST", url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating HttpsHealthCheck: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	waitErr := computeOperationWaitTime(
		config.clientCompute, op, project, "Creating HttpsHealthCheck",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create HttpsHealthCheck: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating HttpsHealthCheck %q: %#v", d.Id(), res)

	return resourceComputeHttpsHealthCheckRead(d, meta)
}

func resourceComputeHttpsHealthCheckRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/httpsHealthChecks/{{name}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeHttpsHealthCheck %q", d.Id()))
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading HttpsHealthCheck: %s", err)
	}

	if err := d.Set("check_interval_sec", flattenComputeHttpsHealthCheckCheckIntervalSec(res["checkIntervalSec"], d)); err != nil {
		return fmt.Errorf("Error reading HttpsHealthCheck: %s", err)
	}
	if err := d.Set("creation_timestamp", flattenComputeHttpsHealthCheckCreationTimestamp(res["creationTimestamp"], d)); err != nil {
		return fmt.Errorf("Error reading HttpsHealthCheck: %s", err)
	}
	if err := d.Set("description", flattenComputeHttpsHealthCheckDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading HttpsHealthCheck: %s", err)
	}
	if err := d.Set("healthy_threshold", flattenComputeHttpsHealthCheckHealthyThreshold(res["healthyThreshold"], d)); err != nil {
		return fmt.Errorf("Error reading HttpsHealthCheck: %s", err)
	}
	if err := d.Set("host", flattenComputeHttpsHealthCheckHost(res["host"], d)); err != nil {
		return fmt.Errorf("Error reading HttpsHealthCheck: %s", err)
	}
	if err := d.Set("name", flattenComputeHttpsHealthCheckName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading HttpsHealthCheck: %s", err)
	}
	if err := d.Set("port", flattenComputeHttpsHealthCheckPort(res["port"], d)); err != nil {
		return fmt.Errorf("Error reading HttpsHealthCheck: %s", err)
	}
	if err := d.Set("request_path", flattenComputeHttpsHealthCheckRequestPath(res["requestPath"], d)); err != nil {
		return fmt.Errorf("Error reading HttpsHealthCheck: %s", err)
	}
	if err := d.Set("timeout_sec", flattenComputeHttpsHealthCheckTimeoutSec(res["timeoutSec"], d)); err != nil {
		return fmt.Errorf("Error reading HttpsHealthCheck: %s", err)
	}
	if err := d.Set("unhealthy_threshold", flattenComputeHttpsHealthCheckUnhealthyThreshold(res["unhealthyThreshold"], d)); err != nil {
		return fmt.Errorf("Error reading HttpsHealthCheck: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading HttpsHealthCheck: %s", err)
	}

	return nil
}

func resourceComputeHttpsHealthCheckUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	checkIntervalSecProp, err := expandComputeHttpsHealthCheckCheckIntervalSec(d.Get("check_interval_sec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("check_interval_sec"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, checkIntervalSecProp)) {
		obj["checkIntervalSec"] = checkIntervalSecProp
	}
	descriptionProp, err := expandComputeHttpsHealthCheckDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	healthyThresholdProp, err := expandComputeHttpsHealthCheckHealthyThreshold(d.Get("healthy_threshold"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("healthy_threshold"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, healthyThresholdProp)) {
		obj["healthyThreshold"] = healthyThresholdProp
	}
	hostProp, err := expandComputeHttpsHealthCheckHost(d.Get("host"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("host"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, hostProp)) {
		obj["host"] = hostProp
	}
	nameProp, err := expandComputeHttpsHealthCheckName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	portProp, err := expandComputeHttpsHealthCheckPort(d.Get("port"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("port"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, portProp)) {
		obj["port"] = portProp
	}
	requestPathProp, err := expandComputeHttpsHealthCheckRequestPath(d.Get("request_path"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("request_path"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, requestPathProp)) {
		obj["requestPath"] = requestPathProp
	}
	timeoutSecProp, err := expandComputeHttpsHealthCheckTimeoutSec(d.Get("timeout_sec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("timeout_sec"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, timeoutSecProp)) {
		obj["timeoutSec"] = timeoutSecProp
	}
	unhealthyThresholdProp, err := expandComputeHttpsHealthCheckUnhealthyThreshold(d.Get("unhealthy_threshold"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("unhealthy_threshold"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, unhealthyThresholdProp)) {
		obj["unhealthyThreshold"] = unhealthyThresholdProp
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/httpsHealthChecks/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating HttpsHealthCheck %q: %#v", d.Id(), obj)
	res, err := sendRequestWithTimeout(config, "PUT", url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating HttpsHealthCheck %q: %s", d.Id(), err)
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWaitTime(
		config.clientCompute, op, project, "Updating HttpsHealthCheck",
		int(d.Timeout(schema.TimeoutUpdate).Minutes()))

	if err != nil {
		return err
	}

	return resourceComputeHttpsHealthCheckRead(d, meta)
}

func resourceComputeHttpsHealthCheckDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/httpsHealthChecks/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting HttpsHealthCheck %q", d.Id())
	res, err := sendRequestWithTimeout(config, "DELETE", url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "HttpsHealthCheck")
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWaitTime(
		config.clientCompute, op, project, "Deleting HttpsHealthCheck",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting HttpsHealthCheck %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeHttpsHealthCheckImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{"projects/(?P<project>[^/]+)/global/httpsHealthChecks/(?P<name>[^/]+)", "(?P<project>[^/]+)/(?P<name>[^/]+)", "(?P<name>[^/]+)"}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeHttpsHealthCheckCheckIntervalSec(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeHttpsHealthCheckCreationTimestamp(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeHttpsHealthCheckDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeHttpsHealthCheckHealthyThreshold(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeHttpsHealthCheckHost(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeHttpsHealthCheckName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeHttpsHealthCheckPort(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeHttpsHealthCheckRequestPath(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeHttpsHealthCheckTimeoutSec(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeHttpsHealthCheckUnhealthyThreshold(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func expandComputeHttpsHealthCheckCheckIntervalSec(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHttpsHealthCheckDescription(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHttpsHealthCheckHealthyThreshold(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHttpsHealthCheckHost(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHttpsHealthCheckName(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHttpsHealthCheckPort(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHttpsHealthCheckRequestPath(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHttpsHealthCheckTimeoutSec(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHttpsHealthCheckUnhealthyThreshold(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}