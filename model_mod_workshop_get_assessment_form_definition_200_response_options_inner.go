/*
Moodle Webservice API

Auto-generated OpenAPI spec for Moodle's Webservice API. Requires installation of the Catalyst RESTful API plugin: https://github.com/catalyst/moodle-webservice_restful.

API version: 4.3.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package moodleclient

import (
	"encoding/json"
)

// checks if the ModWorkshopGetAssessmentFormDefinition200ResponseOptionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModWorkshopGetAssessmentFormDefinition200ResponseOptionsInner{}

// ModWorkshopGetAssessmentFormDefinition200ResponseOptionsInner struct for ModWorkshopGetAssessmentFormDefinition200ResponseOptionsInner
type ModWorkshopGetAssessmentFormDefinition200ResponseOptionsInner struct {
	// Option name.
	Name *string `json:"name,omitempty"`
	// Option value.
	Value *string `json:"value,omitempty"`
}

// NewModWorkshopGetAssessmentFormDefinition200ResponseOptionsInner instantiates a new ModWorkshopGetAssessmentFormDefinition200ResponseOptionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModWorkshopGetAssessmentFormDefinition200ResponseOptionsInner() *ModWorkshopGetAssessmentFormDefinition200ResponseOptionsInner {
	this := ModWorkshopGetAssessmentFormDefinition200ResponseOptionsInner{}
	var name string = "null"
	this.Name = &name
	var value string = "null"
	this.Value = &value
	return &this
}

// NewModWorkshopGetAssessmentFormDefinition200ResponseOptionsInnerWithDefaults instantiates a new ModWorkshopGetAssessmentFormDefinition200ResponseOptionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModWorkshopGetAssessmentFormDefinition200ResponseOptionsInnerWithDefaults() *ModWorkshopGetAssessmentFormDefinition200ResponseOptionsInner {
	this := ModWorkshopGetAssessmentFormDefinition200ResponseOptionsInner{}
	var name string = "null"
	this.Name = &name
	var value string = "null"
	this.Value = &value
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModWorkshopGetAssessmentFormDefinition200ResponseOptionsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModWorkshopGetAssessmentFormDefinition200ResponseOptionsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModWorkshopGetAssessmentFormDefinition200ResponseOptionsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModWorkshopGetAssessmentFormDefinition200ResponseOptionsInner) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ModWorkshopGetAssessmentFormDefinition200ResponseOptionsInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModWorkshopGetAssessmentFormDefinition200ResponseOptionsInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ModWorkshopGetAssessmentFormDefinition200ResponseOptionsInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ModWorkshopGetAssessmentFormDefinition200ResponseOptionsInner) SetValue(v string) {
	o.Value = &v
}

func (o ModWorkshopGetAssessmentFormDefinition200ResponseOptionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModWorkshopGetAssessmentFormDefinition200ResponseOptionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableModWorkshopGetAssessmentFormDefinition200ResponseOptionsInner struct {
	value *ModWorkshopGetAssessmentFormDefinition200ResponseOptionsInner
	isSet bool
}

func (v NullableModWorkshopGetAssessmentFormDefinition200ResponseOptionsInner) Get() *ModWorkshopGetAssessmentFormDefinition200ResponseOptionsInner {
	return v.value
}

func (v *NullableModWorkshopGetAssessmentFormDefinition200ResponseOptionsInner) Set(val *ModWorkshopGetAssessmentFormDefinition200ResponseOptionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableModWorkshopGetAssessmentFormDefinition200ResponseOptionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableModWorkshopGetAssessmentFormDefinition200ResponseOptionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModWorkshopGetAssessmentFormDefinition200ResponseOptionsInner(val *ModWorkshopGetAssessmentFormDefinition200ResponseOptionsInner) *NullableModWorkshopGetAssessmentFormDefinition200ResponseOptionsInner {
	return &NullableModWorkshopGetAssessmentFormDefinition200ResponseOptionsInner{value: val, isSet: true}
}

func (v NullableModWorkshopGetAssessmentFormDefinition200ResponseOptionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModWorkshopGetAssessmentFormDefinition200ResponseOptionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

