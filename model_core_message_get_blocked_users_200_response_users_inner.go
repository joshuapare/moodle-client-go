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

// checks if the CoreMessageGetBlockedUsers200ResponseUsersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreMessageGetBlockedUsers200ResponseUsersInner{}

// CoreMessageGetBlockedUsers200ResponseUsersInner struct for CoreMessageGetBlockedUsers200ResponseUsersInner
type CoreMessageGetBlockedUsers200ResponseUsersInner struct {
	// User full name
	Fullname *string `json:"fullname,omitempty"`
	// User ID
	Id *int32 `json:"id,omitempty"`
	// User picture URL
	Profileimageurl *string `json:"profileimageurl,omitempty"`
}

// NewCoreMessageGetBlockedUsers200ResponseUsersInner instantiates a new CoreMessageGetBlockedUsers200ResponseUsersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreMessageGetBlockedUsers200ResponseUsersInner() *CoreMessageGetBlockedUsers200ResponseUsersInner {
	this := CoreMessageGetBlockedUsers200ResponseUsersInner{}
	var fullname string = "null"
	this.Fullname = &fullname
	return &this
}

// NewCoreMessageGetBlockedUsers200ResponseUsersInnerWithDefaults instantiates a new CoreMessageGetBlockedUsers200ResponseUsersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreMessageGetBlockedUsers200ResponseUsersInnerWithDefaults() *CoreMessageGetBlockedUsers200ResponseUsersInner {
	this := CoreMessageGetBlockedUsers200ResponseUsersInner{}
	var fullname string = "null"
	this.Fullname = &fullname
	return &this
}

// GetFullname returns the Fullname field value if set, zero value otherwise.
func (o *CoreMessageGetBlockedUsers200ResponseUsersInner) GetFullname() string {
	if o == nil || IsNil(o.Fullname) {
		var ret string
		return ret
	}
	return *o.Fullname
}

// GetFullnameOk returns a tuple with the Fullname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetBlockedUsers200ResponseUsersInner) GetFullnameOk() (*string, bool) {
	if o == nil || IsNil(o.Fullname) {
		return nil, false
	}
	return o.Fullname, true
}

// HasFullname returns a boolean if a field has been set.
func (o *CoreMessageGetBlockedUsers200ResponseUsersInner) HasFullname() bool {
	if o != nil && !IsNil(o.Fullname) {
		return true
	}

	return false
}

// SetFullname gets a reference to the given string and assigns it to the Fullname field.
func (o *CoreMessageGetBlockedUsers200ResponseUsersInner) SetFullname(v string) {
	o.Fullname = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CoreMessageGetBlockedUsers200ResponseUsersInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetBlockedUsers200ResponseUsersInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CoreMessageGetBlockedUsers200ResponseUsersInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CoreMessageGetBlockedUsers200ResponseUsersInner) SetId(v int32) {
	o.Id = &v
}

// GetProfileimageurl returns the Profileimageurl field value if set, zero value otherwise.
func (o *CoreMessageGetBlockedUsers200ResponseUsersInner) GetProfileimageurl() string {
	if o == nil || IsNil(o.Profileimageurl) {
		var ret string
		return ret
	}
	return *o.Profileimageurl
}

// GetProfileimageurlOk returns a tuple with the Profileimageurl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetBlockedUsers200ResponseUsersInner) GetProfileimageurlOk() (*string, bool) {
	if o == nil || IsNil(o.Profileimageurl) {
		return nil, false
	}
	return o.Profileimageurl, true
}

// HasProfileimageurl returns a boolean if a field has been set.
func (o *CoreMessageGetBlockedUsers200ResponseUsersInner) HasProfileimageurl() bool {
	if o != nil && !IsNil(o.Profileimageurl) {
		return true
	}

	return false
}

// SetProfileimageurl gets a reference to the given string and assigns it to the Profileimageurl field.
func (o *CoreMessageGetBlockedUsers200ResponseUsersInner) SetProfileimageurl(v string) {
	o.Profileimageurl = &v
}

func (o CoreMessageGetBlockedUsers200ResponseUsersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreMessageGetBlockedUsers200ResponseUsersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fullname) {
		toSerialize["fullname"] = o.Fullname
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Profileimageurl) {
		toSerialize["profileimageurl"] = o.Profileimageurl
	}
	return toSerialize, nil
}

type NullableCoreMessageGetBlockedUsers200ResponseUsersInner struct {
	value *CoreMessageGetBlockedUsers200ResponseUsersInner
	isSet bool
}

func (v NullableCoreMessageGetBlockedUsers200ResponseUsersInner) Get() *CoreMessageGetBlockedUsers200ResponseUsersInner {
	return v.value
}

func (v *NullableCoreMessageGetBlockedUsers200ResponseUsersInner) Set(val *CoreMessageGetBlockedUsers200ResponseUsersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreMessageGetBlockedUsers200ResponseUsersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreMessageGetBlockedUsers200ResponseUsersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreMessageGetBlockedUsers200ResponseUsersInner(val *CoreMessageGetBlockedUsers200ResponseUsersInner) *NullableCoreMessageGetBlockedUsers200ResponseUsersInner {
	return &NullableCoreMessageGetBlockedUsers200ResponseUsersInner{value: val, isSet: true}
}

func (v NullableCoreMessageGetBlockedUsers200ResponseUsersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreMessageGetBlockedUsers200ResponseUsersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

