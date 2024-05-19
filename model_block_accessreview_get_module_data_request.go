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

// checks if the BlockAccessreviewGetModuleDataRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockAccessreviewGetModuleDataRequest{}

// BlockAccessreviewGetModuleDataRequest struct for BlockAccessreviewGetModuleDataRequest
type BlockAccessreviewGetModuleDataRequest struct {
	// The course id to obtain results for.
	Courseid int32 `json:"courseid"`
}

type _BlockAccessreviewGetModuleDataRequest BlockAccessreviewGetModuleDataRequest

// NewBlockAccessreviewGetModuleDataRequest instantiates a new BlockAccessreviewGetModuleDataRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockAccessreviewGetModuleDataRequest(courseid int32) *BlockAccessreviewGetModuleDataRequest {
	this := BlockAccessreviewGetModuleDataRequest{}
	this.Courseid = courseid
	return &this
}

// NewBlockAccessreviewGetModuleDataRequestWithDefaults instantiates a new BlockAccessreviewGetModuleDataRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockAccessreviewGetModuleDataRequestWithDefaults() *BlockAccessreviewGetModuleDataRequest {
	this := BlockAccessreviewGetModuleDataRequest{}
	var courseid int32 = null
	this.Courseid = courseid
	return &this
}

// GetCourseid returns the Courseid field value
func (o *BlockAccessreviewGetModuleDataRequest) GetCourseid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Courseid
}

// GetCourseidOk returns a tuple with the Courseid field value
// and a boolean to check if the value has been set.
func (o *BlockAccessreviewGetModuleDataRequest) GetCourseidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Courseid, true
}

// SetCourseid sets field value
func (o *BlockAccessreviewGetModuleDataRequest) SetCourseid(v int32) {
	o.Courseid = v
}

func (o BlockAccessreviewGetModuleDataRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockAccessreviewGetModuleDataRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["courseid"] = o.Courseid
	return toSerialize, nil
}

func (o *BlockAccessreviewGetModuleDataRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"courseid",
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

	varBlockAccessreviewGetModuleDataRequest := _BlockAccessreviewGetModuleDataRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBlockAccessreviewGetModuleDataRequest)

	if err != nil {
		return err
	}

	*o = BlockAccessreviewGetModuleDataRequest(varBlockAccessreviewGetModuleDataRequest)

	return err
}

type NullableBlockAccessreviewGetModuleDataRequest struct {
	value *BlockAccessreviewGetModuleDataRequest
	isSet bool
}

func (v NullableBlockAccessreviewGetModuleDataRequest) Get() *BlockAccessreviewGetModuleDataRequest {
	return v.value
}

func (v *NullableBlockAccessreviewGetModuleDataRequest) Set(val *BlockAccessreviewGetModuleDataRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockAccessreviewGetModuleDataRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockAccessreviewGetModuleDataRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockAccessreviewGetModuleDataRequest(val *BlockAccessreviewGetModuleDataRequest) *NullableBlockAccessreviewGetModuleDataRequest {
	return &NullableBlockAccessreviewGetModuleDataRequest{value: val, isSet: true}
}

func (v NullableBlockAccessreviewGetModuleDataRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockAccessreviewGetModuleDataRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

