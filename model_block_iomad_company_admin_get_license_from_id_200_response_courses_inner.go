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

// checks if the BlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInner{}

// BlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInner struct for BlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInner
type BlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInner struct {
	// Course full name
	Fullname *string `json:"fullname,omitempty"`
	// Course ID
	Id *int32 `json:"id,omitempty"`
}

// NewBlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInner instantiates a new BlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInner() *BlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInner {
	this := BlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInner{}
	return &this
}

// NewBlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInnerWithDefaults instantiates a new BlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInnerWithDefaults() *BlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInner {
	this := BlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInner{}
	return &this
}

// GetFullname returns the Fullname field value if set, zero value otherwise.
func (o *BlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInner) GetFullname() string {
	if o == nil || IsNil(o.Fullname) {
		var ret string
		return ret
	}
	return *o.Fullname
}

// GetFullnameOk returns a tuple with the Fullname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInner) GetFullnameOk() (*string, bool) {
	if o == nil || IsNil(o.Fullname) {
		return nil, false
	}
	return o.Fullname, true
}

// HasFullname returns a boolean if a field has been set.
func (o *BlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInner) HasFullname() bool {
	if o != nil && !IsNil(o.Fullname) {
		return true
	}

	return false
}

// SetFullname gets a reference to the given string and assigns it to the Fullname field.
func (o *BlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInner) SetFullname(v string) {
	o.Fullname = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *BlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInner) SetId(v int32) {
	o.Id = &v
}

func (o BlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fullname) {
		toSerialize["fullname"] = o.Fullname
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableBlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInner struct {
	value *BlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInner
	isSet bool
}

func (v NullableBlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInner) Get() *BlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInner {
	return v.value
}

func (v *NullableBlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInner) Set(val *BlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInner(val *BlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInner) *NullableBlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInner {
	return &NullableBlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInner{value: val, isSet: true}
}

func (v NullableBlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockIomadCompanyAdminGetLicenseFromId200ResponseCoursesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


