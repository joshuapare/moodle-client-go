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

// checks if the CoreGroupCreateGroupsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreGroupCreateGroupsRequest{}

// CoreGroupCreateGroupsRequest struct for CoreGroupCreateGroupsRequest
type CoreGroupCreateGroupsRequest struct {
	Groups []CoreGroupCreateGroupsRequestGroupsInner `json:"groups"`
}

type _CoreGroupCreateGroupsRequest CoreGroupCreateGroupsRequest

// NewCoreGroupCreateGroupsRequest instantiates a new CoreGroupCreateGroupsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreGroupCreateGroupsRequest(groups []CoreGroupCreateGroupsRequestGroupsInner) *CoreGroupCreateGroupsRequest {
	this := CoreGroupCreateGroupsRequest{}
	this.Groups = groups
	return &this
}

// NewCoreGroupCreateGroupsRequestWithDefaults instantiates a new CoreGroupCreateGroupsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreGroupCreateGroupsRequestWithDefaults() *CoreGroupCreateGroupsRequest {
	this := CoreGroupCreateGroupsRequest{}
	return &this
}

// GetGroups returns the Groups field value
func (o *CoreGroupCreateGroupsRequest) GetGroups() []CoreGroupCreateGroupsRequestGroupsInner {
	if o == nil {
		var ret []CoreGroupCreateGroupsRequestGroupsInner
		return ret
	}

	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value
// and a boolean to check if the value has been set.
func (o *CoreGroupCreateGroupsRequest) GetGroupsOk() ([]CoreGroupCreateGroupsRequestGroupsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Groups, true
}

// SetGroups sets field value
func (o *CoreGroupCreateGroupsRequest) SetGroups(v []CoreGroupCreateGroupsRequestGroupsInner) {
	o.Groups = v
}

func (o CoreGroupCreateGroupsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreGroupCreateGroupsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["groups"] = o.Groups
	return toSerialize, nil
}

func (o *CoreGroupCreateGroupsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"groups",
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

	varCoreGroupCreateGroupsRequest := _CoreGroupCreateGroupsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreGroupCreateGroupsRequest)

	if err != nil {
		return err
	}

	*o = CoreGroupCreateGroupsRequest(varCoreGroupCreateGroupsRequest)

	return err
}

type NullableCoreGroupCreateGroupsRequest struct {
	value *CoreGroupCreateGroupsRequest
	isSet bool
}

func (v NullableCoreGroupCreateGroupsRequest) Get() *CoreGroupCreateGroupsRequest {
	return v.value
}

func (v *NullableCoreGroupCreateGroupsRequest) Set(val *CoreGroupCreateGroupsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreGroupCreateGroupsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreGroupCreateGroupsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreGroupCreateGroupsRequest(val *CoreGroupCreateGroupsRequest) *NullableCoreGroupCreateGroupsRequest {
	return &NullableCoreGroupCreateGroupsRequest{value: val, isSet: true}
}

func (v NullableCoreGroupCreateGroupsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreGroupCreateGroupsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


