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

// checks if the CoreMessageGetReceivedContactRequestsCountRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreMessageGetReceivedContactRequestsCountRequest{}

// CoreMessageGetReceivedContactRequestsCountRequest struct for CoreMessageGetReceivedContactRequestsCountRequest
type CoreMessageGetReceivedContactRequestsCountRequest struct {
	// The id of the user we want to return the number of received contact requests for
	Userid int32 `json:"userid"`
}

type _CoreMessageGetReceivedContactRequestsCountRequest CoreMessageGetReceivedContactRequestsCountRequest

// NewCoreMessageGetReceivedContactRequestsCountRequest instantiates a new CoreMessageGetReceivedContactRequestsCountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreMessageGetReceivedContactRequestsCountRequest(userid int32) *CoreMessageGetReceivedContactRequestsCountRequest {
	this := CoreMessageGetReceivedContactRequestsCountRequest{}
	this.Userid = userid
	return &this
}

// NewCoreMessageGetReceivedContactRequestsCountRequestWithDefaults instantiates a new CoreMessageGetReceivedContactRequestsCountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreMessageGetReceivedContactRequestsCountRequestWithDefaults() *CoreMessageGetReceivedContactRequestsCountRequest {
	this := CoreMessageGetReceivedContactRequestsCountRequest{}
	var userid int32 = null
	this.Userid = userid
	return &this
}

// GetUserid returns the Userid field value
func (o *CoreMessageGetReceivedContactRequestsCountRequest) GetUserid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Userid
}

// GetUseridOk returns a tuple with the Userid field value
// and a boolean to check if the value has been set.
func (o *CoreMessageGetReceivedContactRequestsCountRequest) GetUseridOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Userid, true
}

// SetUserid sets field value
func (o *CoreMessageGetReceivedContactRequestsCountRequest) SetUserid(v int32) {
	o.Userid = v
}

func (o CoreMessageGetReceivedContactRequestsCountRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreMessageGetReceivedContactRequestsCountRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userid"] = o.Userid
	return toSerialize, nil
}

func (o *CoreMessageGetReceivedContactRequestsCountRequest) UnmarshalJSON(data []byte) (err error) {
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

	varCoreMessageGetReceivedContactRequestsCountRequest := _CoreMessageGetReceivedContactRequestsCountRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreMessageGetReceivedContactRequestsCountRequest)

	if err != nil {
		return err
	}

	*o = CoreMessageGetReceivedContactRequestsCountRequest(varCoreMessageGetReceivedContactRequestsCountRequest)

	return err
}

type NullableCoreMessageGetReceivedContactRequestsCountRequest struct {
	value *CoreMessageGetReceivedContactRequestsCountRequest
	isSet bool
}

func (v NullableCoreMessageGetReceivedContactRequestsCountRequest) Get() *CoreMessageGetReceivedContactRequestsCountRequest {
	return v.value
}

func (v *NullableCoreMessageGetReceivedContactRequestsCountRequest) Set(val *CoreMessageGetReceivedContactRequestsCountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreMessageGetReceivedContactRequestsCountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreMessageGetReceivedContactRequestsCountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreMessageGetReceivedContactRequestsCountRequest(val *CoreMessageGetReceivedContactRequestsCountRequest) *NullableCoreMessageGetReceivedContactRequestsCountRequest {
	return &NullableCoreMessageGetReceivedContactRequestsCountRequest{value: val, isSet: true}
}

func (v NullableCoreMessageGetReceivedContactRequestsCountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreMessageGetReceivedContactRequestsCountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


