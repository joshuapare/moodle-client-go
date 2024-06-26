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

// checks if the CoreNotesGetNotes200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreNotesGetNotes200Response{}

// CoreNotesGetNotes200Response struct for CoreNotesGetNotes200Response
type CoreNotesGetNotes200Response struct {
	Notes []CoreNotesGetNotes200ResponseNotesInner `json:"notes"`
	Warnings []CoreNotesGetNotes200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _CoreNotesGetNotes200Response CoreNotesGetNotes200Response

// NewCoreNotesGetNotes200Response instantiates a new CoreNotesGetNotes200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreNotesGetNotes200Response(notes []CoreNotesGetNotes200ResponseNotesInner) *CoreNotesGetNotes200Response {
	this := CoreNotesGetNotes200Response{}
	this.Notes = notes
	return &this
}

// NewCoreNotesGetNotes200ResponseWithDefaults instantiates a new CoreNotesGetNotes200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreNotesGetNotes200ResponseWithDefaults() *CoreNotesGetNotes200Response {
	this := CoreNotesGetNotes200Response{}
	return &this
}

// GetNotes returns the Notes field value
func (o *CoreNotesGetNotes200Response) GetNotes() []CoreNotesGetNotes200ResponseNotesInner {
	if o == nil {
		var ret []CoreNotesGetNotes200ResponseNotesInner
		return ret
	}

	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value
// and a boolean to check if the value has been set.
func (o *CoreNotesGetNotes200Response) GetNotesOk() ([]CoreNotesGetNotes200ResponseNotesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notes, true
}

// SetNotes sets field value
func (o *CoreNotesGetNotes200Response) SetNotes(v []CoreNotesGetNotes200ResponseNotesInner) {
	o.Notes = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *CoreNotesGetNotes200Response) GetWarnings() []CoreNotesGetNotes200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []CoreNotesGetNotes200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreNotesGetNotes200Response) GetWarningsOk() ([]CoreNotesGetNotes200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *CoreNotesGetNotes200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []CoreNotesGetNotes200ResponseWarningsInner and assigns it to the Warnings field.
func (o *CoreNotesGetNotes200Response) SetWarnings(v []CoreNotesGetNotes200ResponseWarningsInner) {
	o.Warnings = v
}

func (o CoreNotesGetNotes200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreNotesGetNotes200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["notes"] = o.Notes
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *CoreNotesGetNotes200Response) UnmarshalJSON(data []byte) (err error) {
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

	varCoreNotesGetNotes200Response := _CoreNotesGetNotes200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreNotesGetNotes200Response)

	if err != nil {
		return err
	}

	*o = CoreNotesGetNotes200Response(varCoreNotesGetNotes200Response)

	return err
}

type NullableCoreNotesGetNotes200Response struct {
	value *CoreNotesGetNotes200Response
	isSet bool
}

func (v NullableCoreNotesGetNotes200Response) Get() *CoreNotesGetNotes200Response {
	return v.value
}

func (v *NullableCoreNotesGetNotes200Response) Set(val *CoreNotesGetNotes200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreNotesGetNotes200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreNotesGetNotes200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreNotesGetNotes200Response(val *CoreNotesGetNotes200Response) *NullableCoreNotesGetNotes200Response {
	return &NullableCoreNotesGetNotes200Response{value: val, isSet: true}
}

func (v NullableCoreNotesGetNotes200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreNotesGetNotes200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


