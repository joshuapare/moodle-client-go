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

// checks if the CoreNotesGetNotesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreNotesGetNotesRequest{}

// CoreNotesGetNotesRequest struct for CoreNotesGetNotesRequest
type CoreNotesGetNotesRequest struct {
	Notes []map[string]interface{} `json:"notes"`
}

type _CoreNotesGetNotesRequest CoreNotesGetNotesRequest

// NewCoreNotesGetNotesRequest instantiates a new CoreNotesGetNotesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreNotesGetNotesRequest(notes []map[string]interface{}) *CoreNotesGetNotesRequest {
	this := CoreNotesGetNotesRequest{}
	this.Notes = notes
	return &this
}

// NewCoreNotesGetNotesRequestWithDefaults instantiates a new CoreNotesGetNotesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreNotesGetNotesRequestWithDefaults() *CoreNotesGetNotesRequest {
	this := CoreNotesGetNotesRequest{}
	return &this
}

// GetNotes returns the Notes field value
func (o *CoreNotesGetNotesRequest) GetNotes() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value
// and a boolean to check if the value has been set.
func (o *CoreNotesGetNotesRequest) GetNotesOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notes, true
}

// SetNotes sets field value
func (o *CoreNotesGetNotesRequest) SetNotes(v []map[string]interface{}) {
	o.Notes = v
}

func (o CoreNotesGetNotesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreNotesGetNotesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["notes"] = o.Notes
	return toSerialize, nil
}

func (o *CoreNotesGetNotesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"notes",
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

	varCoreNotesGetNotesRequest := _CoreNotesGetNotesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreNotesGetNotesRequest)

	if err != nil {
		return err
	}

	*o = CoreNotesGetNotesRequest(varCoreNotesGetNotesRequest)

	return err
}

type NullableCoreNotesGetNotesRequest struct {
	value *CoreNotesGetNotesRequest
	isSet bool
}

func (v NullableCoreNotesGetNotesRequest) Get() *CoreNotesGetNotesRequest {
	return v.value
}

func (v *NullableCoreNotesGetNotesRequest) Set(val *CoreNotesGetNotesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreNotesGetNotesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreNotesGetNotesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreNotesGetNotesRequest(val *CoreNotesGetNotesRequest) *NullableCoreNotesGetNotesRequest {
	return &NullableCoreNotesGetNotesRequest{value: val, isSet: true}
}

func (v NullableCoreNotesGetNotesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreNotesGetNotesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

