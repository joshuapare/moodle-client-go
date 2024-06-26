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

// checks if the CoreMessageGetMessageProcessorRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreMessageGetMessageProcessorRequest{}

// CoreMessageGetMessageProcessorRequest struct for CoreMessageGetMessageProcessorRequest
type CoreMessageGetMessageProcessorRequest struct {
	// The name of the message processor
	Name string `json:"name"`
	// id of the user, 0 for current user
	Userid int32 `json:"userid"`
}

type _CoreMessageGetMessageProcessorRequest CoreMessageGetMessageProcessorRequest

// NewCoreMessageGetMessageProcessorRequest instantiates a new CoreMessageGetMessageProcessorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreMessageGetMessageProcessorRequest(name string, userid int32) *CoreMessageGetMessageProcessorRequest {
	this := CoreMessageGetMessageProcessorRequest{}
	this.Name = name
	this.Userid = userid
	return &this
}

// NewCoreMessageGetMessageProcessorRequestWithDefaults instantiates a new CoreMessageGetMessageProcessorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreMessageGetMessageProcessorRequestWithDefaults() *CoreMessageGetMessageProcessorRequest {
	this := CoreMessageGetMessageProcessorRequest{}
	var name string = "null"
	this.Name = name
	var userid int32 = null
	this.Userid = userid
	return &this
}

// GetName returns the Name field value
func (o *CoreMessageGetMessageProcessorRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CoreMessageGetMessageProcessorRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CoreMessageGetMessageProcessorRequest) SetName(v string) {
	o.Name = v
}

// GetUserid returns the Userid field value
func (o *CoreMessageGetMessageProcessorRequest) GetUserid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Userid
}

// GetUseridOk returns a tuple with the Userid field value
// and a boolean to check if the value has been set.
func (o *CoreMessageGetMessageProcessorRequest) GetUseridOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Userid, true
}

// SetUserid sets field value
func (o *CoreMessageGetMessageProcessorRequest) SetUserid(v int32) {
	o.Userid = v
}

func (o CoreMessageGetMessageProcessorRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreMessageGetMessageProcessorRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["userid"] = o.Userid
	return toSerialize, nil
}

func (o *CoreMessageGetMessageProcessorRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varCoreMessageGetMessageProcessorRequest := _CoreMessageGetMessageProcessorRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreMessageGetMessageProcessorRequest)

	if err != nil {
		return err
	}

	*o = CoreMessageGetMessageProcessorRequest(varCoreMessageGetMessageProcessorRequest)

	return err
}

type NullableCoreMessageGetMessageProcessorRequest struct {
	value *CoreMessageGetMessageProcessorRequest
	isSet bool
}

func (v NullableCoreMessageGetMessageProcessorRequest) Get() *CoreMessageGetMessageProcessorRequest {
	return v.value
}

func (v *NullableCoreMessageGetMessageProcessorRequest) Set(val *CoreMessageGetMessageProcessorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreMessageGetMessageProcessorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreMessageGetMessageProcessorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreMessageGetMessageProcessorRequest(val *CoreMessageGetMessageProcessorRequest) *NullableCoreMessageGetMessageProcessorRequest {
	return &NullableCoreMessageGetMessageProcessorRequest{value: val, isSet: true}
}

func (v NullableCoreMessageGetMessageProcessorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreMessageGetMessageProcessorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


