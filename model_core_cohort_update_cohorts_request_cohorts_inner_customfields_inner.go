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

// checks if the CoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInner{}

// CoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInner struct for CoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInner
type CoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInner struct {
	// The shortname of the custom field
	Shortname *string `json:"shortname,omitempty"`
	// The value of the custom field
	Value *string `json:"value,omitempty"`
}

// NewCoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInner instantiates a new CoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInner() *CoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInner {
	this := CoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInner{}
	return &this
}

// NewCoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInnerWithDefaults instantiates a new CoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInnerWithDefaults() *CoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInner {
	this := CoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInner{}
	return &this
}

// GetShortname returns the Shortname field value if set, zero value otherwise.
func (o *CoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInner) GetShortname() string {
	if o == nil || IsNil(o.Shortname) {
		var ret string
		return ret
	}
	return *o.Shortname
}

// GetShortnameOk returns a tuple with the Shortname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInner) GetShortnameOk() (*string, bool) {
	if o == nil || IsNil(o.Shortname) {
		return nil, false
	}
	return o.Shortname, true
}

// HasShortname returns a boolean if a field has been set.
func (o *CoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInner) HasShortname() bool {
	if o != nil && !IsNil(o.Shortname) {
		return true
	}

	return false
}

// SetShortname gets a reference to the given string and assigns it to the Shortname field.
func (o *CoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInner) SetShortname(v string) {
	o.Shortname = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInner) SetValue(v string) {
	o.Value = &v
}

func (o CoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Shortname) {
		toSerialize["shortname"] = o.Shortname
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableCoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInner struct {
	value *CoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInner
	isSet bool
}

func (v NullableCoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInner) Get() *CoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInner {
	return v.value
}

func (v *NullableCoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInner) Set(val *CoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInner(val *CoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInner) *NullableCoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInner {
	return &NullableCoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInner{value: val, isSet: true}
}

func (v NullableCoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCohortUpdateCohortsRequestCohortsInnerCustomfieldsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


