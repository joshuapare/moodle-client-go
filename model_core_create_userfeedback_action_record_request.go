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

// checks if the CoreCreateUserfeedbackActionRecordRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCreateUserfeedbackActionRecordRequest{}

// CoreCreateUserfeedbackActionRecordRequest struct for CoreCreateUserfeedbackActionRecordRequest
type CoreCreateUserfeedbackActionRecordRequest struct {
	// The action taken by user
	Action string `json:"action"`
	// The context id of the page the user is in
	Contextid int32 `json:"contextid"`
}

type _CoreCreateUserfeedbackActionRecordRequest CoreCreateUserfeedbackActionRecordRequest

// NewCoreCreateUserfeedbackActionRecordRequest instantiates a new CoreCreateUserfeedbackActionRecordRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCreateUserfeedbackActionRecordRequest(action string, contextid int32) *CoreCreateUserfeedbackActionRecordRequest {
	this := CoreCreateUserfeedbackActionRecordRequest{}
	this.Action = action
	this.Contextid = contextid
	return &this
}

// NewCoreCreateUserfeedbackActionRecordRequestWithDefaults instantiates a new CoreCreateUserfeedbackActionRecordRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCreateUserfeedbackActionRecordRequestWithDefaults() *CoreCreateUserfeedbackActionRecordRequest {
	this := CoreCreateUserfeedbackActionRecordRequest{}
	var action string = "null"
	this.Action = action
	var contextid int32 = null
	this.Contextid = contextid
	return &this
}

// GetAction returns the Action field value
func (o *CoreCreateUserfeedbackActionRecordRequest) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *CoreCreateUserfeedbackActionRecordRequest) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *CoreCreateUserfeedbackActionRecordRequest) SetAction(v string) {
	o.Action = v
}

// GetContextid returns the Contextid field value
func (o *CoreCreateUserfeedbackActionRecordRequest) GetContextid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Contextid
}

// GetContextidOk returns a tuple with the Contextid field value
// and a boolean to check if the value has been set.
func (o *CoreCreateUserfeedbackActionRecordRequest) GetContextidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contextid, true
}

// SetContextid sets field value
func (o *CoreCreateUserfeedbackActionRecordRequest) SetContextid(v int32) {
	o.Contextid = v
}

func (o CoreCreateUserfeedbackActionRecordRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCreateUserfeedbackActionRecordRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	toSerialize["contextid"] = o.Contextid
	return toSerialize, nil
}

func (o *CoreCreateUserfeedbackActionRecordRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
		"contextid",
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

	varCoreCreateUserfeedbackActionRecordRequest := _CoreCreateUserfeedbackActionRecordRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreCreateUserfeedbackActionRecordRequest)

	if err != nil {
		return err
	}

	*o = CoreCreateUserfeedbackActionRecordRequest(varCoreCreateUserfeedbackActionRecordRequest)

	return err
}

type NullableCoreCreateUserfeedbackActionRecordRequest struct {
	value *CoreCreateUserfeedbackActionRecordRequest
	isSet bool
}

func (v NullableCoreCreateUserfeedbackActionRecordRequest) Get() *CoreCreateUserfeedbackActionRecordRequest {
	return v.value
}

func (v *NullableCoreCreateUserfeedbackActionRecordRequest) Set(val *CoreCreateUserfeedbackActionRecordRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCreateUserfeedbackActionRecordRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCreateUserfeedbackActionRecordRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCreateUserfeedbackActionRecordRequest(val *CoreCreateUserfeedbackActionRecordRequest) *NullableCoreCreateUserfeedbackActionRecordRequest {
	return &NullableCoreCreateUserfeedbackActionRecordRequest{value: val, isSet: true}
}

func (v NullableCoreCreateUserfeedbackActionRecordRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCreateUserfeedbackActionRecordRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


