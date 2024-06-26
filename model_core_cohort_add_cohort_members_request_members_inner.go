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

// checks if the CoreCohortAddCohortMembersRequestMembersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCohortAddCohortMembersRequestMembersInner{}

// CoreCohortAddCohortMembersRequestMembersInner struct for CoreCohortAddCohortMembersRequestMembersInner
type CoreCohortAddCohortMembersRequestMembersInner struct {
	Cohorttype *CoreCohortAddCohortMembersRequestMembersInnerCohorttype `json:"cohorttype,omitempty"`
	Usertype *CoreCohortAddCohortMembersRequestMembersInnerUsertype `json:"usertype,omitempty"`
}

// NewCoreCohortAddCohortMembersRequestMembersInner instantiates a new CoreCohortAddCohortMembersRequestMembersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCohortAddCohortMembersRequestMembersInner() *CoreCohortAddCohortMembersRequestMembersInner {
	this := CoreCohortAddCohortMembersRequestMembersInner{}
	return &this
}

// NewCoreCohortAddCohortMembersRequestMembersInnerWithDefaults instantiates a new CoreCohortAddCohortMembersRequestMembersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCohortAddCohortMembersRequestMembersInnerWithDefaults() *CoreCohortAddCohortMembersRequestMembersInner {
	this := CoreCohortAddCohortMembersRequestMembersInner{}
	return &this
}

// GetCohorttype returns the Cohorttype field value if set, zero value otherwise.
func (o *CoreCohortAddCohortMembersRequestMembersInner) GetCohorttype() CoreCohortAddCohortMembersRequestMembersInnerCohorttype {
	if o == nil || IsNil(o.Cohorttype) {
		var ret CoreCohortAddCohortMembersRequestMembersInnerCohorttype
		return ret
	}
	return *o.Cohorttype
}

// GetCohorttypeOk returns a tuple with the Cohorttype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCohortAddCohortMembersRequestMembersInner) GetCohorttypeOk() (*CoreCohortAddCohortMembersRequestMembersInnerCohorttype, bool) {
	if o == nil || IsNil(o.Cohorttype) {
		return nil, false
	}
	return o.Cohorttype, true
}

// HasCohorttype returns a boolean if a field has been set.
func (o *CoreCohortAddCohortMembersRequestMembersInner) HasCohorttype() bool {
	if o != nil && !IsNil(o.Cohorttype) {
		return true
	}

	return false
}

// SetCohorttype gets a reference to the given CoreCohortAddCohortMembersRequestMembersInnerCohorttype and assigns it to the Cohorttype field.
func (o *CoreCohortAddCohortMembersRequestMembersInner) SetCohorttype(v CoreCohortAddCohortMembersRequestMembersInnerCohorttype) {
	o.Cohorttype = &v
}

// GetUsertype returns the Usertype field value if set, zero value otherwise.
func (o *CoreCohortAddCohortMembersRequestMembersInner) GetUsertype() CoreCohortAddCohortMembersRequestMembersInnerUsertype {
	if o == nil || IsNil(o.Usertype) {
		var ret CoreCohortAddCohortMembersRequestMembersInnerUsertype
		return ret
	}
	return *o.Usertype
}

// GetUsertypeOk returns a tuple with the Usertype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCohortAddCohortMembersRequestMembersInner) GetUsertypeOk() (*CoreCohortAddCohortMembersRequestMembersInnerUsertype, bool) {
	if o == nil || IsNil(o.Usertype) {
		return nil, false
	}
	return o.Usertype, true
}

// HasUsertype returns a boolean if a field has been set.
func (o *CoreCohortAddCohortMembersRequestMembersInner) HasUsertype() bool {
	if o != nil && !IsNil(o.Usertype) {
		return true
	}

	return false
}

// SetUsertype gets a reference to the given CoreCohortAddCohortMembersRequestMembersInnerUsertype and assigns it to the Usertype field.
func (o *CoreCohortAddCohortMembersRequestMembersInner) SetUsertype(v CoreCohortAddCohortMembersRequestMembersInnerUsertype) {
	o.Usertype = &v
}

func (o CoreCohortAddCohortMembersRequestMembersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCohortAddCohortMembersRequestMembersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cohorttype) {
		toSerialize["cohorttype"] = o.Cohorttype
	}
	if !IsNil(o.Usertype) {
		toSerialize["usertype"] = o.Usertype
	}
	return toSerialize, nil
}

type NullableCoreCohortAddCohortMembersRequestMembersInner struct {
	value *CoreCohortAddCohortMembersRequestMembersInner
	isSet bool
}

func (v NullableCoreCohortAddCohortMembersRequestMembersInner) Get() *CoreCohortAddCohortMembersRequestMembersInner {
	return v.value
}

func (v *NullableCoreCohortAddCohortMembersRequestMembersInner) Set(val *CoreCohortAddCohortMembersRequestMembersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCohortAddCohortMembersRequestMembersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCohortAddCohortMembersRequestMembersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCohortAddCohortMembersRequestMembersInner(val *CoreCohortAddCohortMembersRequestMembersInner) *NullableCoreCohortAddCohortMembersRequestMembersInner {
	return &NullableCoreCohortAddCohortMembersRequestMembersInner{value: val, isSet: true}
}

func (v NullableCoreCohortAddCohortMembersRequestMembersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCohortAddCohortMembersRequestMembersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


