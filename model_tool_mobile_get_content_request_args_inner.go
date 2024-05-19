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

// checks if the ToolMobileGetContentRequestArgsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ToolMobileGetContentRequestArgsInner{}

// ToolMobileGetContentRequestArgsInner struct for ToolMobileGetContentRequestArgsInner
type ToolMobileGetContentRequestArgsInner struct {
	// Param name.
	Name *string `json:"name,omitempty"`
	// Param value.
	Value *string `json:"value,omitempty"`
}

// NewToolMobileGetContentRequestArgsInner instantiates a new ToolMobileGetContentRequestArgsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToolMobileGetContentRequestArgsInner() *ToolMobileGetContentRequestArgsInner {
	this := ToolMobileGetContentRequestArgsInner{}
	var name string = "null"
	this.Name = &name
	var value string = "null"
	this.Value = &value
	return &this
}

// NewToolMobileGetContentRequestArgsInnerWithDefaults instantiates a new ToolMobileGetContentRequestArgsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToolMobileGetContentRequestArgsInnerWithDefaults() *ToolMobileGetContentRequestArgsInner {
	this := ToolMobileGetContentRequestArgsInner{}
	var name string = "null"
	this.Name = &name
	var value string = "null"
	this.Value = &value
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ToolMobileGetContentRequestArgsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolMobileGetContentRequestArgsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ToolMobileGetContentRequestArgsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ToolMobileGetContentRequestArgsInner) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ToolMobileGetContentRequestArgsInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolMobileGetContentRequestArgsInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ToolMobileGetContentRequestArgsInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ToolMobileGetContentRequestArgsInner) SetValue(v string) {
	o.Value = &v
}

func (o ToolMobileGetContentRequestArgsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToolMobileGetContentRequestArgsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableToolMobileGetContentRequestArgsInner struct {
	value *ToolMobileGetContentRequestArgsInner
	isSet bool
}

func (v NullableToolMobileGetContentRequestArgsInner) Get() *ToolMobileGetContentRequestArgsInner {
	return v.value
}

func (v *NullableToolMobileGetContentRequestArgsInner) Set(val *ToolMobileGetContentRequestArgsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableToolMobileGetContentRequestArgsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableToolMobileGetContentRequestArgsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToolMobileGetContentRequestArgsInner(val *ToolMobileGetContentRequestArgsInner) *NullableToolMobileGetContentRequestArgsInner {
	return &NullableToolMobileGetContentRequestArgsInner{value: val, isSet: true}
}

func (v NullableToolMobileGetContentRequestArgsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToolMobileGetContentRequestArgsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


