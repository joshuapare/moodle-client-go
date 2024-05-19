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

// checks if the CoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInner{}

// CoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInner struct for CoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInner
type CoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInner struct {
	// The name of the preferences
	Name *string `json:"name,omitempty"`
	// The value of the preference
	Value *string `json:"value,omitempty"`
}

// NewCoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInner instantiates a new CoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInner() *CoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInner {
	this := CoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInner{}
	var name string = "null"
	this.Name = &name
	var value string = "null"
	this.Value = &value
	return &this
}

// NewCoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInnerWithDefaults instantiates a new CoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInnerWithDefaults() *CoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInner {
	this := CoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInner{}
	var name string = "null"
	this.Name = &name
	var value string = "null"
	this.Value = &value
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInner) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInner) SetValue(v string) {
	o.Value = &v
}

func (o CoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableCoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInner struct {
	value *CoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInner
	isSet bool
}

func (v NullableCoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInner) Get() *CoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInner {
	return v.value
}

func (v *NullableCoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInner) Set(val *CoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInner(val *CoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInner) *NullableCoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInner {
	return &NullableCoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInner{value: val, isSet: true}
}

func (v NullableCoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreGradesGetEnrolledUsersForSelector200ResponseUsersInnerPreferencesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

