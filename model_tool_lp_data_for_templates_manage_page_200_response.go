/*
Moodle Webservice API

Auto-generated OpenAPI spec for Moodle's Webservice API. Requires installation of the Catalyst RESTful API plugin: https://github.com/catalyst/moodle-webservice_restful.

API version: 4.3.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package moodleclient

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ToolLpDataForTemplatesManagePage200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ToolLpDataForTemplatesManagePage200Response{}

// ToolLpDataForTemplatesManagePage200Response struct for ToolLpDataForTemplatesManagePage200Response
type ToolLpDataForTemplatesManagePage200Response struct {
	// Whether the user manage the templates
	Canmanage bool `json:"canmanage"`
	Navigation []map[string]interface{} `json:"navigation"`
	// The page context id
	Pagecontextid int32 `json:"pagecontextid"`
	// Url to the tool_lp plugin folder on this Moodle site
	Pluginbaseurl string `json:"pluginbaseurl"`
	Templates []ToolLpDataForTemplatesManagePage200ResponseTemplatesInner `json:"templates"`
}

type _ToolLpDataForTemplatesManagePage200Response ToolLpDataForTemplatesManagePage200Response

// NewToolLpDataForTemplatesManagePage200Response instantiates a new ToolLpDataForTemplatesManagePage200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToolLpDataForTemplatesManagePage200Response(canmanage bool, navigation []map[string]interface{}, pagecontextid int32, pluginbaseurl string, templates []ToolLpDataForTemplatesManagePage200ResponseTemplatesInner) *ToolLpDataForTemplatesManagePage200Response {
	this := ToolLpDataForTemplatesManagePage200Response{}
	this.Canmanage = canmanage
	this.Navigation = navigation
	this.Pagecontextid = pagecontextid
	this.Pluginbaseurl = pluginbaseurl
	this.Templates = templates
	return &this
}

// NewToolLpDataForTemplatesManagePage200ResponseWithDefaults instantiates a new ToolLpDataForTemplatesManagePage200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToolLpDataForTemplatesManagePage200ResponseWithDefaults() *ToolLpDataForTemplatesManagePage200Response {
	this := ToolLpDataForTemplatesManagePage200Response{}
	var canmanage bool = null
	this.Canmanage = canmanage
	return &this
}

// GetCanmanage returns the Canmanage field value
func (o *ToolLpDataForTemplatesManagePage200Response) GetCanmanage() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Canmanage
}

// GetCanmanageOk returns a tuple with the Canmanage field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForTemplatesManagePage200Response) GetCanmanageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Canmanage, true
}

// SetCanmanage sets field value
func (o *ToolLpDataForTemplatesManagePage200Response) SetCanmanage(v bool) {
	o.Canmanage = v
}

// GetNavigation returns the Navigation field value
func (o *ToolLpDataForTemplatesManagePage200Response) GetNavigation() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Navigation
}

// GetNavigationOk returns a tuple with the Navigation field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForTemplatesManagePage200Response) GetNavigationOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Navigation, true
}

// SetNavigation sets field value
func (o *ToolLpDataForTemplatesManagePage200Response) SetNavigation(v []map[string]interface{}) {
	o.Navigation = v
}

// GetPagecontextid returns the Pagecontextid field value
func (o *ToolLpDataForTemplatesManagePage200Response) GetPagecontextid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pagecontextid
}

// GetPagecontextidOk returns a tuple with the Pagecontextid field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForTemplatesManagePage200Response) GetPagecontextidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagecontextid, true
}

// SetPagecontextid sets field value
func (o *ToolLpDataForTemplatesManagePage200Response) SetPagecontextid(v int32) {
	o.Pagecontextid = v
}

// GetPluginbaseurl returns the Pluginbaseurl field value
func (o *ToolLpDataForTemplatesManagePage200Response) GetPluginbaseurl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pluginbaseurl
}

// GetPluginbaseurlOk returns a tuple with the Pluginbaseurl field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForTemplatesManagePage200Response) GetPluginbaseurlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pluginbaseurl, true
}

// SetPluginbaseurl sets field value
func (o *ToolLpDataForTemplatesManagePage200Response) SetPluginbaseurl(v string) {
	o.Pluginbaseurl = v
}

// GetTemplates returns the Templates field value
func (o *ToolLpDataForTemplatesManagePage200Response) GetTemplates() []ToolLpDataForTemplatesManagePage200ResponseTemplatesInner {
	if o == nil {
		var ret []ToolLpDataForTemplatesManagePage200ResponseTemplatesInner
		return ret
	}

	return o.Templates
}

// GetTemplatesOk returns a tuple with the Templates field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForTemplatesManagePage200Response) GetTemplatesOk() ([]ToolLpDataForTemplatesManagePage200ResponseTemplatesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Templates, true
}

// SetTemplates sets field value
func (o *ToolLpDataForTemplatesManagePage200Response) SetTemplates(v []ToolLpDataForTemplatesManagePage200ResponseTemplatesInner) {
	o.Templates = v
}

func (o ToolLpDataForTemplatesManagePage200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToolLpDataForTemplatesManagePage200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["canmanage"] = o.Canmanage
	toSerialize["navigation"] = o.Navigation
	toSerialize["pagecontextid"] = o.Pagecontextid
	toSerialize["pluginbaseurl"] = o.Pluginbaseurl
	toSerialize["templates"] = o.Templates
	return toSerialize, nil
}

func (o *ToolLpDataForTemplatesManagePage200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"canmanage",
		"navigation",
		"pagecontextid",
		"pluginbaseurl",
		"templates",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varToolLpDataForTemplatesManagePage200Response := _ToolLpDataForTemplatesManagePage200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varToolLpDataForTemplatesManagePage200Response)

	if err != nil {
		return err
	}

	*o = ToolLpDataForTemplatesManagePage200Response(varToolLpDataForTemplatesManagePage200Response)

	return err
}

type NullableToolLpDataForTemplatesManagePage200Response struct {
	value *ToolLpDataForTemplatesManagePage200Response
	isSet bool
}

func (v NullableToolLpDataForTemplatesManagePage200Response) Get() *ToolLpDataForTemplatesManagePage200Response {
	return v.value
}

func (v *NullableToolLpDataForTemplatesManagePage200Response) Set(val *ToolLpDataForTemplatesManagePage200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableToolLpDataForTemplatesManagePage200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableToolLpDataForTemplatesManagePage200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToolLpDataForTemplatesManagePage200Response(val *ToolLpDataForTemplatesManagePage200Response) *NullableToolLpDataForTemplatesManagePage200Response {
	return &NullableToolLpDataForTemplatesManagePage200Response{value: val, isSet: true}
}

func (v NullableToolLpDataForTemplatesManagePage200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToolLpDataForTemplatesManagePage200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


