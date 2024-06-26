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

// checks if the ToolLpDataForUserEvidencePage200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ToolLpDataForUserEvidencePage200Response{}

// ToolLpDataForUserEvidencePage200Response struct for ToolLpDataForUserEvidencePage200Response
type ToolLpDataForUserEvidencePage200Response struct {
	// Url to the tool_lp plugin folder on this Moodle site
	Pluginbaseurl string `json:"pluginbaseurl"`
	Userevidence ToolLpDataForUserEvidencePage200ResponseUserevidence `json:"userevidence"`
}

type _ToolLpDataForUserEvidencePage200Response ToolLpDataForUserEvidencePage200Response

// NewToolLpDataForUserEvidencePage200Response instantiates a new ToolLpDataForUserEvidencePage200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToolLpDataForUserEvidencePage200Response(pluginbaseurl string, userevidence ToolLpDataForUserEvidencePage200ResponseUserevidence) *ToolLpDataForUserEvidencePage200Response {
	this := ToolLpDataForUserEvidencePage200Response{}
	this.Pluginbaseurl = pluginbaseurl
	this.Userevidence = userevidence
	return &this
}

// NewToolLpDataForUserEvidencePage200ResponseWithDefaults instantiates a new ToolLpDataForUserEvidencePage200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToolLpDataForUserEvidencePage200ResponseWithDefaults() *ToolLpDataForUserEvidencePage200Response {
	this := ToolLpDataForUserEvidencePage200Response{}
	return &this
}

// GetPluginbaseurl returns the Pluginbaseurl field value
func (o *ToolLpDataForUserEvidencePage200Response) GetPluginbaseurl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pluginbaseurl
}

// GetPluginbaseurlOk returns a tuple with the Pluginbaseurl field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForUserEvidencePage200Response) GetPluginbaseurlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pluginbaseurl, true
}

// SetPluginbaseurl sets field value
func (o *ToolLpDataForUserEvidencePage200Response) SetPluginbaseurl(v string) {
	o.Pluginbaseurl = v
}

// GetUserevidence returns the Userevidence field value
func (o *ToolLpDataForUserEvidencePage200Response) GetUserevidence() ToolLpDataForUserEvidencePage200ResponseUserevidence {
	if o == nil {
		var ret ToolLpDataForUserEvidencePage200ResponseUserevidence
		return ret
	}

	return o.Userevidence
}

// GetUserevidenceOk returns a tuple with the Userevidence field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForUserEvidencePage200Response) GetUserevidenceOk() (*ToolLpDataForUserEvidencePage200ResponseUserevidence, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Userevidence, true
}

// SetUserevidence sets field value
func (o *ToolLpDataForUserEvidencePage200Response) SetUserevidence(v ToolLpDataForUserEvidencePage200ResponseUserevidence) {
	o.Userevidence = v
}

func (o ToolLpDataForUserEvidencePage200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToolLpDataForUserEvidencePage200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pluginbaseurl"] = o.Pluginbaseurl
	toSerialize["userevidence"] = o.Userevidence
	return toSerialize, nil
}

func (o *ToolLpDataForUserEvidencePage200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pluginbaseurl",
		"userevidence",
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

	varToolLpDataForUserEvidencePage200Response := _ToolLpDataForUserEvidencePage200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varToolLpDataForUserEvidencePage200Response)

	if err != nil {
		return err
	}

	*o = ToolLpDataForUserEvidencePage200Response(varToolLpDataForUserEvidencePage200Response)

	return err
}

type NullableToolLpDataForUserEvidencePage200Response struct {
	value *ToolLpDataForUserEvidencePage200Response
	isSet bool
}

func (v NullableToolLpDataForUserEvidencePage200Response) Get() *ToolLpDataForUserEvidencePage200Response {
	return v.value
}

func (v *NullableToolLpDataForUserEvidencePage200Response) Set(val *ToolLpDataForUserEvidencePage200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableToolLpDataForUserEvidencePage200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableToolLpDataForUserEvidencePage200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToolLpDataForUserEvidencePage200Response(val *ToolLpDataForUserEvidencePage200Response) *NullableToolLpDataForUserEvidencePage200Response {
	return &NullableToolLpDataForUserEvidencePage200Response{value: val, isSet: true}
}

func (v NullableToolLpDataForUserEvidencePage200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToolLpDataForUserEvidencePage200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


