/*
Moodle Webservice API

Auto-generated OpenAPI spec for Moodle's Webservice API. Requires installation of the Catalyst RESTful API plugin: https://github.com/catalyst/moodle-webservice_restful.

API version: 4.3.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package moodleclient

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CoreMessageGetBlockedUsersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreMessageGetBlockedUsersRequest{}

// CoreMessageGetBlockedUsersRequest struct for CoreMessageGetBlockedUsersRequest
type CoreMessageGetBlockedUsersRequest struct {
	// the user whose blocked users we want to retrieve
	Userid int32 `json:"userid"`
}

type _CoreMessageGetBlockedUsersRequest CoreMessageGetBlockedUsersRequest

// NewCoreMessageGetBlockedUsersRequest instantiates a new CoreMessageGetBlockedUsersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreMessageGetBlockedUsersRequest(userid int32) *CoreMessageGetBlockedUsersRequest {
	this := CoreMessageGetBlockedUsersRequest{}
	this.Userid = userid
	return &this
}

// NewCoreMessageGetBlockedUsersRequestWithDefaults instantiates a new CoreMessageGetBlockedUsersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreMessageGetBlockedUsersRequestWithDefaults() *CoreMessageGetBlockedUsersRequest {
	this := CoreMessageGetBlockedUsersRequest{}
	var userid int32 = null
	this.Userid = userid
	return &this
}

// GetUserid returns the Userid field value
func (o *CoreMessageGetBlockedUsersRequest) GetUserid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Userid
}

// GetUseridOk returns a tuple with the Userid field value
// and a boolean to check if the value has been set.
func (o *CoreMessageGetBlockedUsersRequest) GetUseridOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Userid, true
}

// SetUserid sets field value
func (o *CoreMessageGetBlockedUsersRequest) SetUserid(v int32) {
	o.Userid = v
}

func (o CoreMessageGetBlockedUsersRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreMessageGetBlockedUsersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userid"] = o.Userid
	return toSerialize, nil
}

func (o *CoreMessageGetBlockedUsersRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"userid",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCoreMessageGetBlockedUsersRequest := _CoreMessageGetBlockedUsersRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreMessageGetBlockedUsersRequest)

	if err != nil {
		return err
	}

	*o = CoreMessageGetBlockedUsersRequest(varCoreMessageGetBlockedUsersRequest)

	return err
}

type NullableCoreMessageGetBlockedUsersRequest struct {
	value *CoreMessageGetBlockedUsersRequest
	isSet bool
}

func (v NullableCoreMessageGetBlockedUsersRequest) Get() *CoreMessageGetBlockedUsersRequest {
	return v.value
}

func (v *NullableCoreMessageGetBlockedUsersRequest) Set(val *CoreMessageGetBlockedUsersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreMessageGetBlockedUsersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreMessageGetBlockedUsersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreMessageGetBlockedUsersRequest(val *CoreMessageGetBlockedUsersRequest) *NullableCoreMessageGetBlockedUsersRequest {
	return &NullableCoreMessageGetBlockedUsersRequest{value: val, isSet: true}
}

func (v NullableCoreMessageGetBlockedUsersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreMessageGetBlockedUsersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


