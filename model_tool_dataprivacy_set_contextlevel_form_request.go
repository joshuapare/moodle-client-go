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

// checks if the ToolDataprivacySetContextlevelFormRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ToolDataprivacySetContextlevelFormRequest{}

// ToolDataprivacySetContextlevelFormRequest struct for ToolDataprivacySetContextlevelFormRequest
type ToolDataprivacySetContextlevelFormRequest struct {
	// The context level data, encoded as a json array
	Jsonformdata string `json:"jsonformdata"`
}

type _ToolDataprivacySetContextlevelFormRequest ToolDataprivacySetContextlevelFormRequest

// NewToolDataprivacySetContextlevelFormRequest instantiates a new ToolDataprivacySetContextlevelFormRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToolDataprivacySetContextlevelFormRequest(jsonformdata string) *ToolDataprivacySetContextlevelFormRequest {
	this := ToolDataprivacySetContextlevelFormRequest{}
	this.Jsonformdata = jsonformdata
	return &this
}

// NewToolDataprivacySetContextlevelFormRequestWithDefaults instantiates a new ToolDataprivacySetContextlevelFormRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToolDataprivacySetContextlevelFormRequestWithDefaults() *ToolDataprivacySetContextlevelFormRequest {
	this := ToolDataprivacySetContextlevelFormRequest{}
	return &this
}

// GetJsonformdata returns the Jsonformdata field value
func (o *ToolDataprivacySetContextlevelFormRequest) GetJsonformdata() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Jsonformdata
}

// GetJsonformdataOk returns a tuple with the Jsonformdata field value
// and a boolean to check if the value has been set.
func (o *ToolDataprivacySetContextlevelFormRequest) GetJsonformdataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Jsonformdata, true
}

// SetJsonformdata sets field value
func (o *ToolDataprivacySetContextlevelFormRequest) SetJsonformdata(v string) {
	o.Jsonformdata = v
}

func (o ToolDataprivacySetContextlevelFormRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToolDataprivacySetContextlevelFormRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["jsonformdata"] = o.Jsonformdata
	return toSerialize, nil
}

func (o *ToolDataprivacySetContextlevelFormRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"jsonformdata",
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

	varToolDataprivacySetContextlevelFormRequest := _ToolDataprivacySetContextlevelFormRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varToolDataprivacySetContextlevelFormRequest)

	if err != nil {
		return err
	}

	*o = ToolDataprivacySetContextlevelFormRequest(varToolDataprivacySetContextlevelFormRequest)

	return err
}

type NullableToolDataprivacySetContextlevelFormRequest struct {
	value *ToolDataprivacySetContextlevelFormRequest
	isSet bool
}

func (v NullableToolDataprivacySetContextlevelFormRequest) Get() *ToolDataprivacySetContextlevelFormRequest {
	return v.value
}

func (v *NullableToolDataprivacySetContextlevelFormRequest) Set(val *ToolDataprivacySetContextlevelFormRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableToolDataprivacySetContextlevelFormRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableToolDataprivacySetContextlevelFormRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToolDataprivacySetContextlevelFormRequest(val *ToolDataprivacySetContextlevelFormRequest) *NullableToolDataprivacySetContextlevelFormRequest {
	return &NullableToolDataprivacySetContextlevelFormRequest{value: val, isSet: true}
}

func (v NullableToolDataprivacySetContextlevelFormRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToolDataprivacySetContextlevelFormRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

