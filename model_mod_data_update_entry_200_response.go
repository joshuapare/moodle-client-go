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

// checks if the ModDataUpdateEntry200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModDataUpdateEntry200Response{}

// ModDataUpdateEntry200Response struct for ModDataUpdateEntry200Response
type ModDataUpdateEntry200Response struct {
	Fieldnotifications []ModDataUpdateEntry200ResponseFieldnotificationsInner `json:"fieldnotifications"`
	Generalnotifications []map[string]interface{} `json:"generalnotifications"`
	// True if the entry was successfully updated, false other wise.
	Updated bool `json:"updated"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _ModDataUpdateEntry200Response ModDataUpdateEntry200Response

// NewModDataUpdateEntry200Response instantiates a new ModDataUpdateEntry200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModDataUpdateEntry200Response(fieldnotifications []ModDataUpdateEntry200ResponseFieldnotificationsInner, generalnotifications []map[string]interface{}, updated bool) *ModDataUpdateEntry200Response {
	this := ModDataUpdateEntry200Response{}
	this.Fieldnotifications = fieldnotifications
	this.Generalnotifications = generalnotifications
	this.Updated = updated
	return &this
}

// NewModDataUpdateEntry200ResponseWithDefaults instantiates a new ModDataUpdateEntry200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModDataUpdateEntry200ResponseWithDefaults() *ModDataUpdateEntry200Response {
	this := ModDataUpdateEntry200Response{}
	var updated bool = null
	this.Updated = updated
	return &this
}

// GetFieldnotifications returns the Fieldnotifications field value
func (o *ModDataUpdateEntry200Response) GetFieldnotifications() []ModDataUpdateEntry200ResponseFieldnotificationsInner {
	if o == nil {
		var ret []ModDataUpdateEntry200ResponseFieldnotificationsInner
		return ret
	}

	return o.Fieldnotifications
}

// GetFieldnotificationsOk returns a tuple with the Fieldnotifications field value
// and a boolean to check if the value has been set.
func (o *ModDataUpdateEntry200Response) GetFieldnotificationsOk() ([]ModDataUpdateEntry200ResponseFieldnotificationsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fieldnotifications, true
}

// SetFieldnotifications sets field value
func (o *ModDataUpdateEntry200Response) SetFieldnotifications(v []ModDataUpdateEntry200ResponseFieldnotificationsInner) {
	o.Fieldnotifications = v
}

// GetGeneralnotifications returns the Generalnotifications field value
func (o *ModDataUpdateEntry200Response) GetGeneralnotifications() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Generalnotifications
}

// GetGeneralnotificationsOk returns a tuple with the Generalnotifications field value
// and a boolean to check if the value has been set.
func (o *ModDataUpdateEntry200Response) GetGeneralnotificationsOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Generalnotifications, true
}

// SetGeneralnotifications sets field value
func (o *ModDataUpdateEntry200Response) SetGeneralnotifications(v []map[string]interface{}) {
	o.Generalnotifications = v
}

// GetUpdated returns the Updated field value
func (o *ModDataUpdateEntry200Response) GetUpdated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
func (o *ModDataUpdateEntry200Response) GetUpdatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updated, true
}

// SetUpdated sets field value
func (o *ModDataUpdateEntry200Response) SetUpdated(v bool) {
	o.Updated = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ModDataUpdateEntry200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModDataUpdateEntry200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ModDataUpdateEntry200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *ModDataUpdateEntry200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o ModDataUpdateEntry200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModDataUpdateEntry200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fieldnotifications"] = o.Fieldnotifications
	toSerialize["generalnotifications"] = o.Generalnotifications
	toSerialize["updated"] = o.Updated
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *ModDataUpdateEntry200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fieldnotifications",
		"generalnotifications",
		"updated",
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

	varModDataUpdateEntry200Response := _ModDataUpdateEntry200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModDataUpdateEntry200Response)

	if err != nil {
		return err
	}

	*o = ModDataUpdateEntry200Response(varModDataUpdateEntry200Response)

	return err
}

type NullableModDataUpdateEntry200Response struct {
	value *ModDataUpdateEntry200Response
	isSet bool
}

func (v NullableModDataUpdateEntry200Response) Get() *ModDataUpdateEntry200Response {
	return v.value
}

func (v *NullableModDataUpdateEntry200Response) Set(val *ModDataUpdateEntry200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModDataUpdateEntry200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModDataUpdateEntry200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModDataUpdateEntry200Response(val *ModDataUpdateEntry200Response) *NullableModDataUpdateEntry200Response {
	return &NullableModDataUpdateEntry200Response{value: val, isSet: true}
}

func (v NullableModDataUpdateEntry200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModDataUpdateEntry200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

