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

// checks if the CoreGroupDeleteGroupsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreGroupDeleteGroupsRequest{}

// CoreGroupDeleteGroupsRequest struct for CoreGroupDeleteGroupsRequest
type CoreGroupDeleteGroupsRequest struct {
	Groupids []map[string]interface{} `json:"groupids"`
}

type _CoreGroupDeleteGroupsRequest CoreGroupDeleteGroupsRequest

// NewCoreGroupDeleteGroupsRequest instantiates a new CoreGroupDeleteGroupsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreGroupDeleteGroupsRequest(groupids []map[string]interface{}) *CoreGroupDeleteGroupsRequest {
	this := CoreGroupDeleteGroupsRequest{}
	this.Groupids = groupids
	return &this
}

// NewCoreGroupDeleteGroupsRequestWithDefaults instantiates a new CoreGroupDeleteGroupsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreGroupDeleteGroupsRequestWithDefaults() *CoreGroupDeleteGroupsRequest {
	this := CoreGroupDeleteGroupsRequest{}
	return &this
}

// GetGroupids returns the Groupids field value
func (o *CoreGroupDeleteGroupsRequest) GetGroupids() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Groupids
}

// GetGroupidsOk returns a tuple with the Groupids field value
// and a boolean to check if the value has been set.
func (o *CoreGroupDeleteGroupsRequest) GetGroupidsOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Groupids, true
}

// SetGroupids sets field value
func (o *CoreGroupDeleteGroupsRequest) SetGroupids(v []map[string]interface{}) {
	o.Groupids = v
}

func (o CoreGroupDeleteGroupsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreGroupDeleteGroupsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["groupids"] = o.Groupids
	return toSerialize, nil
}

func (o *CoreGroupDeleteGroupsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"groupids",
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

	varCoreGroupDeleteGroupsRequest := _CoreGroupDeleteGroupsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreGroupDeleteGroupsRequest)

	if err != nil {
		return err
	}

	*o = CoreGroupDeleteGroupsRequest(varCoreGroupDeleteGroupsRequest)

	return err
}

type NullableCoreGroupDeleteGroupsRequest struct {
	value *CoreGroupDeleteGroupsRequest
	isSet bool
}

func (v NullableCoreGroupDeleteGroupsRequest) Get() *CoreGroupDeleteGroupsRequest {
	return v.value
}

func (v *NullableCoreGroupDeleteGroupsRequest) Set(val *CoreGroupDeleteGroupsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreGroupDeleteGroupsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreGroupDeleteGroupsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreGroupDeleteGroupsRequest(val *CoreGroupDeleteGroupsRequest) *NullableCoreGroupDeleteGroupsRequest {
	return &NullableCoreGroupDeleteGroupsRequest{value: val, isSet: true}
}

func (v NullableCoreGroupDeleteGroupsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreGroupDeleteGroupsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


