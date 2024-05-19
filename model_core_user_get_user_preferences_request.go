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

// checks if the CoreUserGetUserPreferencesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreUserGetUserPreferencesRequest{}

// CoreUserGetUserPreferencesRequest struct for CoreUserGetUserPreferencesRequest
type CoreUserGetUserPreferencesRequest struct {
	// preference name, empty for all
	Name *string `json:"name,omitempty"`
	// id of the user, default to current user
	Userid *int32 `json:"userid,omitempty"`
}

// NewCoreUserGetUserPreferencesRequest instantiates a new CoreUserGetUserPreferencesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreUserGetUserPreferencesRequest() *CoreUserGetUserPreferencesRequest {
	this := CoreUserGetUserPreferencesRequest{}
	var name string = ""
	this.Name = &name
	var userid int32 = 0
	this.Userid = &userid
	return &this
}

// NewCoreUserGetUserPreferencesRequestWithDefaults instantiates a new CoreUserGetUserPreferencesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreUserGetUserPreferencesRequestWithDefaults() *CoreUserGetUserPreferencesRequest {
	this := CoreUserGetUserPreferencesRequest{}
	var name string = ""
	this.Name = &name
	var userid int32 = 0
	this.Userid = &userid
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CoreUserGetUserPreferencesRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreUserGetUserPreferencesRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CoreUserGetUserPreferencesRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CoreUserGetUserPreferencesRequest) SetName(v string) {
	o.Name = &v
}

// GetUserid returns the Userid field value if set, zero value otherwise.
func (o *CoreUserGetUserPreferencesRequest) GetUserid() int32 {
	if o == nil || IsNil(o.Userid) {
		var ret int32
		return ret
	}
	return *o.Userid
}

// GetUseridOk returns a tuple with the Userid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreUserGetUserPreferencesRequest) GetUseridOk() (*int32, bool) {
	if o == nil || IsNil(o.Userid) {
		return nil, false
	}
	return o.Userid, true
}

// HasUserid returns a boolean if a field has been set.
func (o *CoreUserGetUserPreferencesRequest) HasUserid() bool {
	if o != nil && !IsNil(o.Userid) {
		return true
	}

	return false
}

// SetUserid gets a reference to the given int32 and assigns it to the Userid field.
func (o *CoreUserGetUserPreferencesRequest) SetUserid(v int32) {
	o.Userid = &v
}

func (o CoreUserGetUserPreferencesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreUserGetUserPreferencesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Userid) {
		toSerialize["userid"] = o.Userid
	}
	return toSerialize, nil
}

type NullableCoreUserGetUserPreferencesRequest struct {
	value *CoreUserGetUserPreferencesRequest
	isSet bool
}

func (v NullableCoreUserGetUserPreferencesRequest) Get() *CoreUserGetUserPreferencesRequest {
	return v.value
}

func (v *NullableCoreUserGetUserPreferencesRequest) Set(val *CoreUserGetUserPreferencesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreUserGetUserPreferencesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreUserGetUserPreferencesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreUserGetUserPreferencesRequest(val *CoreUserGetUserPreferencesRequest) *NullableCoreUserGetUserPreferencesRequest {
	return &NullableCoreUserGetUserPreferencesRequest{value: val, isSet: true}
}

func (v NullableCoreUserGetUserPreferencesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreUserGetUserPreferencesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

