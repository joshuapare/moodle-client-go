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

// checks if the ModDataDeleteSavedPresetRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModDataDeleteSavedPresetRequest{}

// ModDataDeleteSavedPresetRequest struct for ModDataDeleteSavedPresetRequest
type ModDataDeleteSavedPresetRequest struct {
	// Id of the data activity
	Dataid int32 `json:"dataid"`
	Presetnames []map[string]interface{} `json:"presetnames"`
}

type _ModDataDeleteSavedPresetRequest ModDataDeleteSavedPresetRequest

// NewModDataDeleteSavedPresetRequest instantiates a new ModDataDeleteSavedPresetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModDataDeleteSavedPresetRequest(dataid int32, presetnames []map[string]interface{}) *ModDataDeleteSavedPresetRequest {
	this := ModDataDeleteSavedPresetRequest{}
	this.Dataid = dataid
	this.Presetnames = presetnames
	return &this
}

// NewModDataDeleteSavedPresetRequestWithDefaults instantiates a new ModDataDeleteSavedPresetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModDataDeleteSavedPresetRequestWithDefaults() *ModDataDeleteSavedPresetRequest {
	this := ModDataDeleteSavedPresetRequest{}
	var dataid int32 = null
	this.Dataid = dataid
	return &this
}

// GetDataid returns the Dataid field value
func (o *ModDataDeleteSavedPresetRequest) GetDataid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Dataid
}

// GetDataidOk returns a tuple with the Dataid field value
// and a boolean to check if the value has been set.
func (o *ModDataDeleteSavedPresetRequest) GetDataidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dataid, true
}

// SetDataid sets field value
func (o *ModDataDeleteSavedPresetRequest) SetDataid(v int32) {
	o.Dataid = v
}

// GetPresetnames returns the Presetnames field value
func (o *ModDataDeleteSavedPresetRequest) GetPresetnames() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Presetnames
}

// GetPresetnamesOk returns a tuple with the Presetnames field value
// and a boolean to check if the value has been set.
func (o *ModDataDeleteSavedPresetRequest) GetPresetnamesOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Presetnames, true
}

// SetPresetnames sets field value
func (o *ModDataDeleteSavedPresetRequest) SetPresetnames(v []map[string]interface{}) {
	o.Presetnames = v
}

func (o ModDataDeleteSavedPresetRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModDataDeleteSavedPresetRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dataid"] = o.Dataid
	toSerialize["presetnames"] = o.Presetnames
	return toSerialize, nil
}

func (o *ModDataDeleteSavedPresetRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dataid",
		"presetnames",
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

	varModDataDeleteSavedPresetRequest := _ModDataDeleteSavedPresetRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModDataDeleteSavedPresetRequest)

	if err != nil {
		return err
	}

	*o = ModDataDeleteSavedPresetRequest(varModDataDeleteSavedPresetRequest)

	return err
}

type NullableModDataDeleteSavedPresetRequest struct {
	value *ModDataDeleteSavedPresetRequest
	isSet bool
}

func (v NullableModDataDeleteSavedPresetRequest) Get() *ModDataDeleteSavedPresetRequest {
	return v.value
}

func (v *NullableModDataDeleteSavedPresetRequest) Set(val *ModDataDeleteSavedPresetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModDataDeleteSavedPresetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModDataDeleteSavedPresetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModDataDeleteSavedPresetRequest(val *ModDataDeleteSavedPresetRequest) *NullableModDataDeleteSavedPresetRequest {
	return &NullableModDataDeleteSavedPresetRequest{value: val, isSet: true}
}

func (v NullableModDataDeleteSavedPresetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModDataDeleteSavedPresetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


