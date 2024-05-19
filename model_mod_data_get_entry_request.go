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

// checks if the ModDataGetEntryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModDataGetEntryRequest{}

// ModDataGetEntryRequest struct for ModDataGetEntryRequest
type ModDataGetEntryRequest struct {
	// record entry id
	Entryid int32 `json:"entryid"`
	// Whether to return contents or not.
	Returncontents *bool `json:"returncontents,omitempty"`
}

type _ModDataGetEntryRequest ModDataGetEntryRequest

// NewModDataGetEntryRequest instantiates a new ModDataGetEntryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModDataGetEntryRequest(entryid int32) *ModDataGetEntryRequest {
	this := ModDataGetEntryRequest{}
	this.Entryid = entryid
	var returncontents bool = false
	this.Returncontents = &returncontents
	return &this
}

// NewModDataGetEntryRequestWithDefaults instantiates a new ModDataGetEntryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModDataGetEntryRequestWithDefaults() *ModDataGetEntryRequest {
	this := ModDataGetEntryRequest{}
	var entryid int32 = null
	this.Entryid = entryid
	var returncontents bool = false
	this.Returncontents = &returncontents
	return &this
}

// GetEntryid returns the Entryid field value
func (o *ModDataGetEntryRequest) GetEntryid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Entryid
}

// GetEntryidOk returns a tuple with the Entryid field value
// and a boolean to check if the value has been set.
func (o *ModDataGetEntryRequest) GetEntryidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entryid, true
}

// SetEntryid sets field value
func (o *ModDataGetEntryRequest) SetEntryid(v int32) {
	o.Entryid = v
}

// GetReturncontents returns the Returncontents field value if set, zero value otherwise.
func (o *ModDataGetEntryRequest) GetReturncontents() bool {
	if o == nil || IsNil(o.Returncontents) {
		var ret bool
		return ret
	}
	return *o.Returncontents
}

// GetReturncontentsOk returns a tuple with the Returncontents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModDataGetEntryRequest) GetReturncontentsOk() (*bool, bool) {
	if o == nil || IsNil(o.Returncontents) {
		return nil, false
	}
	return o.Returncontents, true
}

// HasReturncontents returns a boolean if a field has been set.
func (o *ModDataGetEntryRequest) HasReturncontents() bool {
	if o != nil && !IsNil(o.Returncontents) {
		return true
	}

	return false
}

// SetReturncontents gets a reference to the given bool and assigns it to the Returncontents field.
func (o *ModDataGetEntryRequest) SetReturncontents(v bool) {
	o.Returncontents = &v
}

func (o ModDataGetEntryRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModDataGetEntryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entryid"] = o.Entryid
	if !IsNil(o.Returncontents) {
		toSerialize["returncontents"] = o.Returncontents
	}
	return toSerialize, nil
}

func (o *ModDataGetEntryRequest) UnmarshalJSON(data []byte) (err error) {
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

	varModDataGetEntryRequest := _ModDataGetEntryRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModDataGetEntryRequest)

	if err != nil {
		return err
	}

	*o = ModDataGetEntryRequest(varModDataGetEntryRequest)

	return err
}

type NullableModDataGetEntryRequest struct {
	value *ModDataGetEntryRequest
	isSet bool
}

func (v NullableModDataGetEntryRequest) Get() *ModDataGetEntryRequest {
	return v.value
}

func (v *NullableModDataGetEntryRequest) Set(val *ModDataGetEntryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModDataGetEntryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModDataGetEntryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModDataGetEntryRequest(val *ModDataGetEntryRequest) *NullableModDataGetEntryRequest {
	return &NullableModDataGetEntryRequest{value: val, isSet: true}
}

func (v NullableModDataGetEntryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModDataGetEntryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

