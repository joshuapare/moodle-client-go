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

// checks if the CoreUserGetPrivateFilesInfoRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreUserGetPrivateFilesInfoRequest{}

// CoreUserGetPrivateFilesInfoRequest struct for CoreUserGetPrivateFilesInfoRequest
type CoreUserGetPrivateFilesInfoRequest struct {
	// Id of the user, default to current user.
	Userid *int32 `json:"userid,omitempty"`
}

// NewCoreUserGetPrivateFilesInfoRequest instantiates a new CoreUserGetPrivateFilesInfoRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreUserGetPrivateFilesInfoRequest() *CoreUserGetPrivateFilesInfoRequest {
	this := CoreUserGetPrivateFilesInfoRequest{}
	var userid int32 = 0
	this.Userid = &userid
	return &this
}

// NewCoreUserGetPrivateFilesInfoRequestWithDefaults instantiates a new CoreUserGetPrivateFilesInfoRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreUserGetPrivateFilesInfoRequestWithDefaults() *CoreUserGetPrivateFilesInfoRequest {
	this := CoreUserGetPrivateFilesInfoRequest{}
	var userid int32 = 0
	this.Userid = &userid
	return &this
}

// GetUserid returns the Userid field value if set, zero value otherwise.
func (o *CoreUserGetPrivateFilesInfoRequest) GetUserid() int32 {
	if o == nil || IsNil(o.Userid) {
		var ret int32
		return ret
	}
	return *o.Userid
}

// GetUseridOk returns a tuple with the Userid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreUserGetPrivateFilesInfoRequest) GetUseridOk() (*int32, bool) {
	if o == nil || IsNil(o.Userid) {
		return nil, false
	}
	return o.Userid, true
}

// HasUserid returns a boolean if a field has been set.
func (o *CoreUserGetPrivateFilesInfoRequest) HasUserid() bool {
	if o != nil && !IsNil(o.Userid) {
		return true
	}

	return false
}

// SetUserid gets a reference to the given int32 and assigns it to the Userid field.
func (o *CoreUserGetPrivateFilesInfoRequest) SetUserid(v int32) {
	o.Userid = &v
}

func (o CoreUserGetPrivateFilesInfoRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreUserGetPrivateFilesInfoRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Userid) {
		toSerialize["userid"] = o.Userid
	}
	return toSerialize, nil
}

type NullableCoreUserGetPrivateFilesInfoRequest struct {
	value *CoreUserGetPrivateFilesInfoRequest
	isSet bool
}

func (v NullableCoreUserGetPrivateFilesInfoRequest) Get() *CoreUserGetPrivateFilesInfoRequest {
	return v.value
}

func (v *NullableCoreUserGetPrivateFilesInfoRequest) Set(val *CoreUserGetPrivateFilesInfoRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreUserGetPrivateFilesInfoRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreUserGetPrivateFilesInfoRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreUserGetPrivateFilesInfoRequest(val *CoreUserGetPrivateFilesInfoRequest) *NullableCoreUserGetPrivateFilesInfoRequest {
	return &NullableCoreUserGetPrivateFilesInfoRequest{value: val, isSet: true}
}

func (v NullableCoreUserGetPrivateFilesInfoRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreUserGetPrivateFilesInfoRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


