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

// checks if the ModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner{}

// ModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner struct for ModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner
type ModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner struct {
	// Scale name.
	Name *string `json:"name,omitempty"`
	// Scale value/option id.
	Value *int32 `json:"value,omitempty"`
}

// NewModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner instantiates a new ModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner() *ModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner {
	this := ModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner{}
	var name string = "null"
	this.Name = &name
	var value int32 = null
	this.Value = &value
	return &this
}

// NewModDataGetEntry200ResponseRatinginfoScalesInnerItemsInnerWithDefaults instantiates a new ModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModDataGetEntry200ResponseRatinginfoScalesInnerItemsInnerWithDefaults() *ModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner {
	this := ModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner{}
	var name string = "null"
	this.Name = &name
	var value int32 = null
	this.Value = &value
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner) GetValue() int32 {
	if o == nil || IsNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner) GetValueOk() (*int32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *ModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner) SetValue(v int32) {
	o.Value = &v
}

func (o ModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner struct {
	value *ModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner
	isSet bool
}

func (v NullableModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner) Get() *ModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner {
	return v.value
}

func (v *NullableModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner) Set(val *ModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner(val *ModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner) *NullableModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner {
	return &NullableModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner{value: val, isSet: true}
}

func (v NullableModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModDataGetEntry200ResponseRatinginfoScalesInnerItemsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


