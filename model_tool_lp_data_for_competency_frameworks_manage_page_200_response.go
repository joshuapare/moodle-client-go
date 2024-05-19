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

// checks if the ToolLpDataForCompetencyFrameworksManagePage200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ToolLpDataForCompetencyFrameworksManagePage200Response{}

// ToolLpDataForCompetencyFrameworksManagePage200Response struct for ToolLpDataForCompetencyFrameworksManagePage200Response
type ToolLpDataForCompetencyFrameworksManagePage200Response struct {
	Competencyframeworks []ToolLpDataForCompetencyFrameworksManagePage200ResponseCompetencyframeworksInner `json:"competencyframeworks"`
	Navigation []map[string]interface{} `json:"navigation"`
	// The page context id
	Pagecontextid int32 `json:"pagecontextid"`
	// Url to the tool_lp plugin folder on this Moodle site
	Pluginbaseurl string `json:"pluginbaseurl"`
}

type _ToolLpDataForCompetencyFrameworksManagePage200Response ToolLpDataForCompetencyFrameworksManagePage200Response

// NewToolLpDataForCompetencyFrameworksManagePage200Response instantiates a new ToolLpDataForCompetencyFrameworksManagePage200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToolLpDataForCompetencyFrameworksManagePage200Response(competencyframeworks []ToolLpDataForCompetencyFrameworksManagePage200ResponseCompetencyframeworksInner, navigation []map[string]interface{}, pagecontextid int32, pluginbaseurl string) *ToolLpDataForCompetencyFrameworksManagePage200Response {
	this := ToolLpDataForCompetencyFrameworksManagePage200Response{}
	this.Competencyframeworks = competencyframeworks
	this.Navigation = navigation
	this.Pagecontextid = pagecontextid
	this.Pluginbaseurl = pluginbaseurl
	return &this
}

// NewToolLpDataForCompetencyFrameworksManagePage200ResponseWithDefaults instantiates a new ToolLpDataForCompetencyFrameworksManagePage200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToolLpDataForCompetencyFrameworksManagePage200ResponseWithDefaults() *ToolLpDataForCompetencyFrameworksManagePage200Response {
	this := ToolLpDataForCompetencyFrameworksManagePage200Response{}
	var pagecontextid int32 = null
	this.Pagecontextid = pagecontextid
	var pluginbaseurl string = "null"
	this.Pluginbaseurl = pluginbaseurl
	return &this
}

// GetCompetencyframeworks returns the Competencyframeworks field value
func (o *ToolLpDataForCompetencyFrameworksManagePage200Response) GetCompetencyframeworks() []ToolLpDataForCompetencyFrameworksManagePage200ResponseCompetencyframeworksInner {
	if o == nil {
		var ret []ToolLpDataForCompetencyFrameworksManagePage200ResponseCompetencyframeworksInner
		return ret
	}

	return o.Competencyframeworks
}

// GetCompetencyframeworksOk returns a tuple with the Competencyframeworks field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForCompetencyFrameworksManagePage200Response) GetCompetencyframeworksOk() ([]ToolLpDataForCompetencyFrameworksManagePage200ResponseCompetencyframeworksInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Competencyframeworks, true
}

// SetCompetencyframeworks sets field value
func (o *ToolLpDataForCompetencyFrameworksManagePage200Response) SetCompetencyframeworks(v []ToolLpDataForCompetencyFrameworksManagePage200ResponseCompetencyframeworksInner) {
	o.Competencyframeworks = v
}

// GetNavigation returns the Navigation field value
func (o *ToolLpDataForCompetencyFrameworksManagePage200Response) GetNavigation() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Navigation
}

// GetNavigationOk returns a tuple with the Navigation field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForCompetencyFrameworksManagePage200Response) GetNavigationOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Navigation, true
}

// SetNavigation sets field value
func (o *ToolLpDataForCompetencyFrameworksManagePage200Response) SetNavigation(v []map[string]interface{}) {
	o.Navigation = v
}

// GetPagecontextid returns the Pagecontextid field value
func (o *ToolLpDataForCompetencyFrameworksManagePage200Response) GetPagecontextid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pagecontextid
}

// GetPagecontextidOk returns a tuple with the Pagecontextid field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForCompetencyFrameworksManagePage200Response) GetPagecontextidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagecontextid, true
}

// SetPagecontextid sets field value
func (o *ToolLpDataForCompetencyFrameworksManagePage200Response) SetPagecontextid(v int32) {
	o.Pagecontextid = v
}

// GetPluginbaseurl returns the Pluginbaseurl field value
func (o *ToolLpDataForCompetencyFrameworksManagePage200Response) GetPluginbaseurl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pluginbaseurl
}

// GetPluginbaseurlOk returns a tuple with the Pluginbaseurl field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForCompetencyFrameworksManagePage200Response) GetPluginbaseurlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pluginbaseurl, true
}

// SetPluginbaseurl sets field value
func (o *ToolLpDataForCompetencyFrameworksManagePage200Response) SetPluginbaseurl(v string) {
	o.Pluginbaseurl = v
}

func (o ToolLpDataForCompetencyFrameworksManagePage200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToolLpDataForCompetencyFrameworksManagePage200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["competencyframeworks"] = o.Competencyframeworks
	toSerialize["navigation"] = o.Navigation
	toSerialize["pagecontextid"] = o.Pagecontextid
	toSerialize["pluginbaseurl"] = o.Pluginbaseurl
	return toSerialize, nil
}

func (o *ToolLpDataForCompetencyFrameworksManagePage200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"competencyframeworks",
		"navigation",
		"pagecontextid",
		"pluginbaseurl",
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

	varToolLpDataForCompetencyFrameworksManagePage200Response := _ToolLpDataForCompetencyFrameworksManagePage200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varToolLpDataForCompetencyFrameworksManagePage200Response)

	if err != nil {
		return err
	}

	*o = ToolLpDataForCompetencyFrameworksManagePage200Response(varToolLpDataForCompetencyFrameworksManagePage200Response)

	return err
}

type NullableToolLpDataForCompetencyFrameworksManagePage200Response struct {
	value *ToolLpDataForCompetencyFrameworksManagePage200Response
	isSet bool
}

func (v NullableToolLpDataForCompetencyFrameworksManagePage200Response) Get() *ToolLpDataForCompetencyFrameworksManagePage200Response {
	return v.value
}

func (v *NullableToolLpDataForCompetencyFrameworksManagePage200Response) Set(val *ToolLpDataForCompetencyFrameworksManagePage200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableToolLpDataForCompetencyFrameworksManagePage200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableToolLpDataForCompetencyFrameworksManagePage200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToolLpDataForCompetencyFrameworksManagePage200Response(val *ToolLpDataForCompetencyFrameworksManagePage200Response) *NullableToolLpDataForCompetencyFrameworksManagePage200Response {
	return &NullableToolLpDataForCompetencyFrameworksManagePage200Response{value: val, isSet: true}
}

func (v NullableToolLpDataForCompetencyFrameworksManagePage200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToolLpDataForCompetencyFrameworksManagePage200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

