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

// checks if the ModScormGetScormAttemptCountRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModScormGetScormAttemptCountRequest{}

// ModScormGetScormAttemptCountRequest struct for ModScormGetScormAttemptCountRequest
type ModScormGetScormAttemptCountRequest struct {
	// Ignores attempts that haven't reported a grade/completion
	Ignoremissingcompletion *bool `json:"ignoremissingcompletion,omitempty"`
	// SCORM instance id
	Scormid int32 `json:"scormid"`
	// User id
	Userid int32 `json:"userid"`
}

type _ModScormGetScormAttemptCountRequest ModScormGetScormAttemptCountRequest

// NewModScormGetScormAttemptCountRequest instantiates a new ModScormGetScormAttemptCountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModScormGetScormAttemptCountRequest(scormid int32, userid int32) *ModScormGetScormAttemptCountRequest {
	this := ModScormGetScormAttemptCountRequest{}
	var ignoremissingcompletion bool = false
	this.Ignoremissingcompletion = &ignoremissingcompletion
	this.Scormid = scormid
	this.Userid = userid
	return &this
}

// NewModScormGetScormAttemptCountRequestWithDefaults instantiates a new ModScormGetScormAttemptCountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModScormGetScormAttemptCountRequestWithDefaults() *ModScormGetScormAttemptCountRequest {
	this := ModScormGetScormAttemptCountRequest{}
	var ignoremissingcompletion bool = false
	this.Ignoremissingcompletion = &ignoremissingcompletion
	var scormid int32 = null
	this.Scormid = scormid
	return &this
}

// GetIgnoremissingcompletion returns the Ignoremissingcompletion field value if set, zero value otherwise.
func (o *ModScormGetScormAttemptCountRequest) GetIgnoremissingcompletion() bool {
	if o == nil || IsNil(o.Ignoremissingcompletion) {
		var ret bool
		return ret
	}
	return *o.Ignoremissingcompletion
}

// GetIgnoremissingcompletionOk returns a tuple with the Ignoremissingcompletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModScormGetScormAttemptCountRequest) GetIgnoremissingcompletionOk() (*bool, bool) {
	if o == nil || IsNil(o.Ignoremissingcompletion) {
		return nil, false
	}
	return o.Ignoremissingcompletion, true
}

// HasIgnoremissingcompletion returns a boolean if a field has been set.
func (o *ModScormGetScormAttemptCountRequest) HasIgnoremissingcompletion() bool {
	if o != nil && !IsNil(o.Ignoremissingcompletion) {
		return true
	}

	return false
}

// SetIgnoremissingcompletion gets a reference to the given bool and assigns it to the Ignoremissingcompletion field.
func (o *ModScormGetScormAttemptCountRequest) SetIgnoremissingcompletion(v bool) {
	o.Ignoremissingcompletion = &v
}

// GetScormid returns the Scormid field value
func (o *ModScormGetScormAttemptCountRequest) GetScormid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Scormid
}

// GetScormidOk returns a tuple with the Scormid field value
// and a boolean to check if the value has been set.
func (o *ModScormGetScormAttemptCountRequest) GetScormidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scormid, true
}

// SetScormid sets field value
func (o *ModScormGetScormAttemptCountRequest) SetScormid(v int32) {
	o.Scormid = v
}

// GetUserid returns the Userid field value
func (o *ModScormGetScormAttemptCountRequest) GetUserid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Userid
}

// GetUseridOk returns a tuple with the Userid field value
// and a boolean to check if the value has been set.
func (o *ModScormGetScormAttemptCountRequest) GetUseridOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Userid, true
}

// SetUserid sets field value
func (o *ModScormGetScormAttemptCountRequest) SetUserid(v int32) {
	o.Userid = v
}

func (o ModScormGetScormAttemptCountRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModScormGetScormAttemptCountRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ignoremissingcompletion) {
		toSerialize["ignoremissingcompletion"] = o.Ignoremissingcompletion
	}
	toSerialize["scormid"] = o.Scormid
	toSerialize["userid"] = o.Userid
	return toSerialize, nil
}

func (o *ModScormGetScormAttemptCountRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"scormid",
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

	varModScormGetScormAttemptCountRequest := _ModScormGetScormAttemptCountRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModScormGetScormAttemptCountRequest)

	if err != nil {
		return err
	}

	*o = ModScormGetScormAttemptCountRequest(varModScormGetScormAttemptCountRequest)

	return err
}

type NullableModScormGetScormAttemptCountRequest struct {
	value *ModScormGetScormAttemptCountRequest
	isSet bool
}

func (v NullableModScormGetScormAttemptCountRequest) Get() *ModScormGetScormAttemptCountRequest {
	return v.value
}

func (v *NullableModScormGetScormAttemptCountRequest) Set(val *ModScormGetScormAttemptCountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModScormGetScormAttemptCountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModScormGetScormAttemptCountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModScormGetScormAttemptCountRequest(val *ModScormGetScormAttemptCountRequest) *NullableModScormGetScormAttemptCountRequest {
	return &NullableModScormGetScormAttemptCountRequest{value: val, isSet: true}
}

func (v NullableModScormGetScormAttemptCountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModScormGetScormAttemptCountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


