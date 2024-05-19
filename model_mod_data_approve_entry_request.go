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

// checks if the ModDataApproveEntryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModDataApproveEntryRequest{}

// ModDataApproveEntryRequest struct for ModDataApproveEntryRequest
type ModDataApproveEntryRequest struct {
	// Whether to approve (true) or unapprove the entry.
	Approve *bool `json:"approve,omitempty"`
	// Record entry id.
	Entryid int32 `json:"entryid"`
}

type _ModDataApproveEntryRequest ModDataApproveEntryRequest

// NewModDataApproveEntryRequest instantiates a new ModDataApproveEntryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModDataApproveEntryRequest(entryid int32) *ModDataApproveEntryRequest {
	this := ModDataApproveEntryRequest{}
	var approve bool = true
	this.Approve = &approve
	this.Entryid = entryid
	return &this
}

// NewModDataApproveEntryRequestWithDefaults instantiates a new ModDataApproveEntryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModDataApproveEntryRequestWithDefaults() *ModDataApproveEntryRequest {
	this := ModDataApproveEntryRequest{}
	var approve bool = true
	this.Approve = &approve
	var entryid int32 = null
	this.Entryid = entryid
	return &this
}

// GetApprove returns the Approve field value if set, zero value otherwise.
func (o *ModDataApproveEntryRequest) GetApprove() bool {
	if o == nil || IsNil(o.Approve) {
		var ret bool
		return ret
	}
	return *o.Approve
}

// GetApproveOk returns a tuple with the Approve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModDataApproveEntryRequest) GetApproveOk() (*bool, bool) {
	if o == nil || IsNil(o.Approve) {
		return nil, false
	}
	return o.Approve, true
}

// HasApprove returns a boolean if a field has been set.
func (o *ModDataApproveEntryRequest) HasApprove() bool {
	if o != nil && !IsNil(o.Approve) {
		return true
	}

	return false
}

// SetApprove gets a reference to the given bool and assigns it to the Approve field.
func (o *ModDataApproveEntryRequest) SetApprove(v bool) {
	o.Approve = &v
}

// GetEntryid returns the Entryid field value
func (o *ModDataApproveEntryRequest) GetEntryid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Entryid
}

// GetEntryidOk returns a tuple with the Entryid field value
// and a boolean to check if the value has been set.
func (o *ModDataApproveEntryRequest) GetEntryidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entryid, true
}

// SetEntryid sets field value
func (o *ModDataApproveEntryRequest) SetEntryid(v int32) {
	o.Entryid = v
}

func (o ModDataApproveEntryRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModDataApproveEntryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Approve) {
		toSerialize["approve"] = o.Approve
	}
	toSerialize["entryid"] = o.Entryid
	return toSerialize, nil
}

func (o *ModDataApproveEntryRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entryid",
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

	varModDataApproveEntryRequest := _ModDataApproveEntryRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModDataApproveEntryRequest)

	if err != nil {
		return err
	}

	*o = ModDataApproveEntryRequest(varModDataApproveEntryRequest)

	return err
}

type NullableModDataApproveEntryRequest struct {
	value *ModDataApproveEntryRequest
	isSet bool
}

func (v NullableModDataApproveEntryRequest) Get() *ModDataApproveEntryRequest {
	return v.value
}

func (v *NullableModDataApproveEntryRequest) Set(val *ModDataApproveEntryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModDataApproveEntryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModDataApproveEntryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModDataApproveEntryRequest(val *ModDataApproveEntryRequest) *NullableModDataApproveEntryRequest {
	return &NullableModDataApproveEntryRequest{value: val, isSet: true}
}

func (v NullableModDataApproveEntryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModDataApproveEntryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

