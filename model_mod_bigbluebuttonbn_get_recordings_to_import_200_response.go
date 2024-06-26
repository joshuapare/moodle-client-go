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

// checks if the ModBigbluebuttonbnGetRecordingsToImport200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModBigbluebuttonbnGetRecordingsToImport200Response{}

// ModBigbluebuttonbnGetRecordingsToImport200Response struct for ModBigbluebuttonbnGetRecordingsToImport200Response
type ModBigbluebuttonbnGetRecordingsToImport200Response struct {
	// Whether the fetch was successful
	Status bool `json:"status"`
	Tabledata *ModBigbluebuttonbnGetRecordingsToImport200ResponseTabledata `json:"tabledata,omitempty"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _ModBigbluebuttonbnGetRecordingsToImport200Response ModBigbluebuttonbnGetRecordingsToImport200Response

// NewModBigbluebuttonbnGetRecordingsToImport200Response instantiates a new ModBigbluebuttonbnGetRecordingsToImport200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModBigbluebuttonbnGetRecordingsToImport200Response(status bool) *ModBigbluebuttonbnGetRecordingsToImport200Response {
	this := ModBigbluebuttonbnGetRecordingsToImport200Response{}
	this.Status = status
	return &this
}

// NewModBigbluebuttonbnGetRecordingsToImport200ResponseWithDefaults instantiates a new ModBigbluebuttonbnGetRecordingsToImport200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModBigbluebuttonbnGetRecordingsToImport200ResponseWithDefaults() *ModBigbluebuttonbnGetRecordingsToImport200Response {
	this := ModBigbluebuttonbnGetRecordingsToImport200Response{}
	return &this
}

// GetStatus returns the Status field value
func (o *ModBigbluebuttonbnGetRecordingsToImport200Response) GetStatus() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ModBigbluebuttonbnGetRecordingsToImport200Response) GetStatusOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ModBigbluebuttonbnGetRecordingsToImport200Response) SetStatus(v bool) {
	o.Status = v
}

// GetTabledata returns the Tabledata field value if set, zero value otherwise.
func (o *ModBigbluebuttonbnGetRecordingsToImport200Response) GetTabledata() ModBigbluebuttonbnGetRecordingsToImport200ResponseTabledata {
	if o == nil || IsNil(o.Tabledata) {
		var ret ModBigbluebuttonbnGetRecordingsToImport200ResponseTabledata
		return ret
	}
	return *o.Tabledata
}

// GetTabledataOk returns a tuple with the Tabledata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModBigbluebuttonbnGetRecordingsToImport200Response) GetTabledataOk() (*ModBigbluebuttonbnGetRecordingsToImport200ResponseTabledata, bool) {
	if o == nil || IsNil(o.Tabledata) {
		return nil, false
	}
	return o.Tabledata, true
}

// HasTabledata returns a boolean if a field has been set.
func (o *ModBigbluebuttonbnGetRecordingsToImport200Response) HasTabledata() bool {
	if o != nil && !IsNil(o.Tabledata) {
		return true
	}

	return false
}

// SetTabledata gets a reference to the given ModBigbluebuttonbnGetRecordingsToImport200ResponseTabledata and assigns it to the Tabledata field.
func (o *ModBigbluebuttonbnGetRecordingsToImport200Response) SetTabledata(v ModBigbluebuttonbnGetRecordingsToImport200ResponseTabledata) {
	o.Tabledata = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ModBigbluebuttonbnGetRecordingsToImport200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModBigbluebuttonbnGetRecordingsToImport200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ModBigbluebuttonbnGetRecordingsToImport200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *ModBigbluebuttonbnGetRecordingsToImport200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o ModBigbluebuttonbnGetRecordingsToImport200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModBigbluebuttonbnGetRecordingsToImport200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	if !IsNil(o.Tabledata) {
		toSerialize["tabledata"] = o.Tabledata
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *ModBigbluebuttonbnGetRecordingsToImport200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
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

	varModBigbluebuttonbnGetRecordingsToImport200Response := _ModBigbluebuttonbnGetRecordingsToImport200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModBigbluebuttonbnGetRecordingsToImport200Response)

	if err != nil {
		return err
	}

	*o = ModBigbluebuttonbnGetRecordingsToImport200Response(varModBigbluebuttonbnGetRecordingsToImport200Response)

	return err
}

type NullableModBigbluebuttonbnGetRecordingsToImport200Response struct {
	value *ModBigbluebuttonbnGetRecordingsToImport200Response
	isSet bool
}

func (v NullableModBigbluebuttonbnGetRecordingsToImport200Response) Get() *ModBigbluebuttonbnGetRecordingsToImport200Response {
	return v.value
}

func (v *NullableModBigbluebuttonbnGetRecordingsToImport200Response) Set(val *ModBigbluebuttonbnGetRecordingsToImport200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModBigbluebuttonbnGetRecordingsToImport200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModBigbluebuttonbnGetRecordingsToImport200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModBigbluebuttonbnGetRecordingsToImport200Response(val *ModBigbluebuttonbnGetRecordingsToImport200Response) *NullableModBigbluebuttonbnGetRecordingsToImport200Response {
	return &NullableModBigbluebuttonbnGetRecordingsToImport200Response{value: val, isSet: true}
}

func (v NullableModBigbluebuttonbnGetRecordingsToImport200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModBigbluebuttonbnGetRecordingsToImport200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


