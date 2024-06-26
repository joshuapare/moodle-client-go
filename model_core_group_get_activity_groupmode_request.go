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

// checks if the CoreGroupGetActivityGroupmodeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreGroupGetActivityGroupmodeRequest{}

// CoreGroupGetActivityGroupmodeRequest struct for CoreGroupGetActivityGroupmodeRequest
type CoreGroupGetActivityGroupmodeRequest struct {
	// course module id
	Cmid int32 `json:"cmid"`
}

type _CoreGroupGetActivityGroupmodeRequest CoreGroupGetActivityGroupmodeRequest

// NewCoreGroupGetActivityGroupmodeRequest instantiates a new CoreGroupGetActivityGroupmodeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreGroupGetActivityGroupmodeRequest(cmid int32) *CoreGroupGetActivityGroupmodeRequest {
	this := CoreGroupGetActivityGroupmodeRequest{}
	this.Cmid = cmid
	return &this
}

// NewCoreGroupGetActivityGroupmodeRequestWithDefaults instantiates a new CoreGroupGetActivityGroupmodeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreGroupGetActivityGroupmodeRequestWithDefaults() *CoreGroupGetActivityGroupmodeRequest {
	this := CoreGroupGetActivityGroupmodeRequest{}
	return &this
}

// GetCmid returns the Cmid field value
func (o *CoreGroupGetActivityGroupmodeRequest) GetCmid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Cmid
}

// GetCmidOk returns a tuple with the Cmid field value
// and a boolean to check if the value has been set.
func (o *CoreGroupGetActivityGroupmodeRequest) GetCmidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cmid, true
}

// SetCmid sets field value
func (o *CoreGroupGetActivityGroupmodeRequest) SetCmid(v int32) {
	o.Cmid = v
}

func (o CoreGroupGetActivityGroupmodeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreGroupGetActivityGroupmodeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cmid"] = o.Cmid
	return toSerialize, nil
}

func (o *CoreGroupGetActivityGroupmodeRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cmid",
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

	varCoreGroupGetActivityGroupmodeRequest := _CoreGroupGetActivityGroupmodeRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreGroupGetActivityGroupmodeRequest)

	if err != nil {
		return err
	}

	*o = CoreGroupGetActivityGroupmodeRequest(varCoreGroupGetActivityGroupmodeRequest)

	return err
}

type NullableCoreGroupGetActivityGroupmodeRequest struct {
	value *CoreGroupGetActivityGroupmodeRequest
	isSet bool
}

func (v NullableCoreGroupGetActivityGroupmodeRequest) Get() *CoreGroupGetActivityGroupmodeRequest {
	return v.value
}

func (v *NullableCoreGroupGetActivityGroupmodeRequest) Set(val *CoreGroupGetActivityGroupmodeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreGroupGetActivityGroupmodeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreGroupGetActivityGroupmodeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreGroupGetActivityGroupmodeRequest(val *CoreGroupGetActivityGroupmodeRequest) *NullableCoreGroupGetActivityGroupmodeRequest {
	return &NullableCoreGroupGetActivityGroupmodeRequest{value: val, isSet: true}
}

func (v NullableCoreGroupGetActivityGroupmodeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreGroupGetActivityGroupmodeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


