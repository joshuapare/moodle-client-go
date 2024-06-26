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

// checks if the ToolLpDataForPlansPageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ToolLpDataForPlansPageRequest{}

// ToolLpDataForPlansPageRequest struct for ToolLpDataForPlansPageRequest
type ToolLpDataForPlansPageRequest struct {
	// The user id
	Userid int32 `json:"userid"`
}

type _ToolLpDataForPlansPageRequest ToolLpDataForPlansPageRequest

// NewToolLpDataForPlansPageRequest instantiates a new ToolLpDataForPlansPageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToolLpDataForPlansPageRequest(userid int32) *ToolLpDataForPlansPageRequest {
	this := ToolLpDataForPlansPageRequest{}
	this.Userid = userid
	return &this
}

// NewToolLpDataForPlansPageRequestWithDefaults instantiates a new ToolLpDataForPlansPageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToolLpDataForPlansPageRequestWithDefaults() *ToolLpDataForPlansPageRequest {
	this := ToolLpDataForPlansPageRequest{}
	return &this
}

// GetUserid returns the Userid field value
func (o *ToolLpDataForPlansPageRequest) GetUserid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Userid
}

// GetUseridOk returns a tuple with the Userid field value
// and a boolean to check if the value has been set.
func (o *ToolLpDataForPlansPageRequest) GetUseridOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Userid, true
}

// SetUserid sets field value
func (o *ToolLpDataForPlansPageRequest) SetUserid(v int32) {
	o.Userid = v
}

func (o ToolLpDataForPlansPageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToolLpDataForPlansPageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userid"] = o.Userid
	return toSerialize, nil
}

func (o *ToolLpDataForPlansPageRequest) UnmarshalJSON(data []byte) (err error) {
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

	varToolLpDataForPlansPageRequest := _ToolLpDataForPlansPageRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varToolLpDataForPlansPageRequest)

	if err != nil {
		return err
	}

	*o = ToolLpDataForPlansPageRequest(varToolLpDataForPlansPageRequest)

	return err
}

type NullableToolLpDataForPlansPageRequest struct {
	value *ToolLpDataForPlansPageRequest
	isSet bool
}

func (v NullableToolLpDataForPlansPageRequest) Get() *ToolLpDataForPlansPageRequest {
	return v.value
}

func (v *NullableToolLpDataForPlansPageRequest) Set(val *ToolLpDataForPlansPageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableToolLpDataForPlansPageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableToolLpDataForPlansPageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToolLpDataForPlansPageRequest(val *ToolLpDataForPlansPageRequest) *NullableToolLpDataForPlansPageRequest {
	return &NullableToolLpDataForPlansPageRequest{value: val, isSet: true}
}

func (v NullableToolLpDataForPlansPageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToolLpDataForPlansPageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


