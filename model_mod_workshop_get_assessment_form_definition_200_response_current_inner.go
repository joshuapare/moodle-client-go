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

// checks if the ModWorkshopGetAssessmentFormDefinition200ResponseCurrentInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModWorkshopGetAssessmentFormDefinition200ResponseCurrentInner{}

// ModWorkshopGetAssessmentFormDefinition200ResponseCurrentInner struct for ModWorkshopGetAssessmentFormDefinition200ResponseCurrentInner
type ModWorkshopGetAssessmentFormDefinition200ResponseCurrentInner struct {
	// Field name.
	Name *string `json:"name,omitempty"`
	// Current field value.
	Value *string `json:"value,omitempty"`
}

// NewModWorkshopGetAssessmentFormDefinition200ResponseCurrentInner instantiates a new ModWorkshopGetAssessmentFormDefinition200ResponseCurrentInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModWorkshopGetAssessmentFormDefinition200ResponseCurrentInner() *ModWorkshopGetAssessmentFormDefinition200ResponseCurrentInner {
	this := ModWorkshopGetAssessmentFormDefinition200ResponseCurrentInner{}
	var name string = "null"
	this.Name = &name
	var value string = "null"
	this.Value = &value
	return &this
}

// NewModWorkshopGetAssessmentFormDefinition200ResponseCurrentInnerWithDefaults instantiates a new ModWorkshopGetAssessmentFormDefinition200ResponseCurrentInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModWorkshopGetAssessmentFormDefinition200ResponseCurrentInnerWithDefaults() *ModWorkshopGetAssessmentFormDefinition200ResponseCurrentInner {
	this := ModWorkshopGetAssessmentFormDefinition200ResponseCurrentInner{}
	var name string = "null"
	this.Name = &name
	var value string = "null"
	this.Value = &value
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModWorkshopGetAssessmentFormDefinition200ResponseCurrentInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModWorkshopGetAssessmentFormDefinition200ResponseCurrentInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModWorkshopGetAssessmentFormDefinition200ResponseCurrentInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModWorkshopGetAssessmentFormDefinition200ResponseCurrentInner) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ModWorkshopGetAssessmentFormDefinition200ResponseCurrentInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModWorkshopGetAssessmentFormDefinition200ResponseCurrentInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ModWorkshopGetAssessmentFormDefinition200ResponseCurrentInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ModWorkshopGetAssessmentFormDefinition200ResponseCurrentInner) SetValue(v string) {
	o.Value = &v
}

func (o ModWorkshopGetAssessmentFormDefinition200ResponseCurrentInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModWorkshopGetAssessmentFormDefinition200ResponseCurrentInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableModWorkshopGetAssessmentFormDefinition200ResponseCurrentInner struct {
	value *ModWorkshopGetAssessmentFormDefinition200ResponseCurrentInner
	isSet bool
}

func (v NullableModWorkshopGetAssessmentFormDefinition200ResponseCurrentInner) Get() *ModWorkshopGetAssessmentFormDefinition200ResponseCurrentInner {
	return v.value
}

func (v *NullableModWorkshopGetAssessmentFormDefinition200ResponseCurrentInner) Set(val *ModWorkshopGetAssessmentFormDefinition200ResponseCurrentInner) {
	v.value = val
	v.isSet = true
}

func (v NullableModWorkshopGetAssessmentFormDefinition200ResponseCurrentInner) IsSet() bool {
	return v.isSet
}

func (v *NullableModWorkshopGetAssessmentFormDefinition200ResponseCurrentInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModWorkshopGetAssessmentFormDefinition200ResponseCurrentInner(val *ModWorkshopGetAssessmentFormDefinition200ResponseCurrentInner) *NullableModWorkshopGetAssessmentFormDefinition200ResponseCurrentInner {
	return &NullableModWorkshopGetAssessmentFormDefinition200ResponseCurrentInner{value: val, isSet: true}
}

func (v NullableModWorkshopGetAssessmentFormDefinition200ResponseCurrentInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModWorkshopGetAssessmentFormDefinition200ResponseCurrentInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

