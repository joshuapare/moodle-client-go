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

// checks if the CoreCompetencyListCompetenciesRequestFiltersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCompetencyListCompetenciesRequestFiltersInner{}

// CoreCompetencyListCompetenciesRequestFiltersInner struct for CoreCompetencyListCompetenciesRequestFiltersInner
type CoreCompetencyListCompetenciesRequestFiltersInner struct {
	// Column name to filter by
	Column *string `json:"column,omitempty"`
	// Value to filter by. Must be exact match
	Value *string `json:"value,omitempty"`
}

// NewCoreCompetencyListCompetenciesRequestFiltersInner instantiates a new CoreCompetencyListCompetenciesRequestFiltersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCompetencyListCompetenciesRequestFiltersInner() *CoreCompetencyListCompetenciesRequestFiltersInner {
	this := CoreCompetencyListCompetenciesRequestFiltersInner{}
	return &this
}

// NewCoreCompetencyListCompetenciesRequestFiltersInnerWithDefaults instantiates a new CoreCompetencyListCompetenciesRequestFiltersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCompetencyListCompetenciesRequestFiltersInnerWithDefaults() *CoreCompetencyListCompetenciesRequestFiltersInner {
	this := CoreCompetencyListCompetenciesRequestFiltersInner{}
	return &this
}

// GetColumn returns the Column field value if set, zero value otherwise.
func (o *CoreCompetencyListCompetenciesRequestFiltersInner) GetColumn() string {
	if o == nil || IsNil(o.Column) {
		var ret string
		return ret
	}
	return *o.Column
}

// GetColumnOk returns a tuple with the Column field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyListCompetenciesRequestFiltersInner) GetColumnOk() (*string, bool) {
	if o == nil || IsNil(o.Column) {
		return nil, false
	}
	return o.Column, true
}

// HasColumn returns a boolean if a field has been set.
func (o *CoreCompetencyListCompetenciesRequestFiltersInner) HasColumn() bool {
	if o != nil && !IsNil(o.Column) {
		return true
	}

	return false
}

// SetColumn gets a reference to the given string and assigns it to the Column field.
func (o *CoreCompetencyListCompetenciesRequestFiltersInner) SetColumn(v string) {
	o.Column = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CoreCompetencyListCompetenciesRequestFiltersInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyListCompetenciesRequestFiltersInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CoreCompetencyListCompetenciesRequestFiltersInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CoreCompetencyListCompetenciesRequestFiltersInner) SetValue(v string) {
	o.Value = &v
}

func (o CoreCompetencyListCompetenciesRequestFiltersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCompetencyListCompetenciesRequestFiltersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Column) {
		toSerialize["column"] = o.Column
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableCoreCompetencyListCompetenciesRequestFiltersInner struct {
	value *CoreCompetencyListCompetenciesRequestFiltersInner
	isSet bool
}

func (v NullableCoreCompetencyListCompetenciesRequestFiltersInner) Get() *CoreCompetencyListCompetenciesRequestFiltersInner {
	return v.value
}

func (v *NullableCoreCompetencyListCompetenciesRequestFiltersInner) Set(val *CoreCompetencyListCompetenciesRequestFiltersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCompetencyListCompetenciesRequestFiltersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCompetencyListCompetenciesRequestFiltersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCompetencyListCompetenciesRequestFiltersInner(val *CoreCompetencyListCompetenciesRequestFiltersInner) *NullableCoreCompetencyListCompetenciesRequestFiltersInner {
	return &NullableCoreCompetencyListCompetenciesRequestFiltersInner{value: val, isSet: true}
}

func (v NullableCoreCompetencyListCompetenciesRequestFiltersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCompetencyListCompetenciesRequestFiltersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


