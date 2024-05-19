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

// checks if the CoreGetUserDatesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreGetUserDatesRequest{}

// CoreGetUserDatesRequest struct for CoreGetUserDatesRequest
type CoreGetUserDatesRequest struct {
	// Context ID. Either use this value, or level and instanceid.
	Contextid *int32 `json:"contextid,omitempty"`
	// Context level. To be used with instanceid.
	Contextlevel *string `json:"contextlevel,omitempty"`
	// Context instance ID. To be used with level
	Instanceid *int32 `json:"instanceid,omitempty"`
	Timestamps []CoreGetUserDatesRequestTimestampsInner `json:"timestamps"`
}

type _CoreGetUserDatesRequest CoreGetUserDatesRequest

// NewCoreGetUserDatesRequest instantiates a new CoreGetUserDatesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreGetUserDatesRequest(timestamps []CoreGetUserDatesRequestTimestampsInner) *CoreGetUserDatesRequest {
	this := CoreGetUserDatesRequest{}
	var contextid int32 = 0
	this.Contextid = &contextid
	var contextlevel string = ""
	this.Contextlevel = &contextlevel
	var instanceid int32 = 0
	this.Instanceid = &instanceid
	this.Timestamps = timestamps
	return &this
}

// NewCoreGetUserDatesRequestWithDefaults instantiates a new CoreGetUserDatesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreGetUserDatesRequestWithDefaults() *CoreGetUserDatesRequest {
	this := CoreGetUserDatesRequest{}
	var contextid int32 = 0
	this.Contextid = &contextid
	var contextlevel string = ""
	this.Contextlevel = &contextlevel
	var instanceid int32 = 0
	this.Instanceid = &instanceid
	return &this
}

// GetContextid returns the Contextid field value if set, zero value otherwise.
func (o *CoreGetUserDatesRequest) GetContextid() int32 {
	if o == nil || IsNil(o.Contextid) {
		var ret int32
		return ret
	}
	return *o.Contextid
}

// GetContextidOk returns a tuple with the Contextid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGetUserDatesRequest) GetContextidOk() (*int32, bool) {
	if o == nil || IsNil(o.Contextid) {
		return nil, false
	}
	return o.Contextid, true
}

// HasContextid returns a boolean if a field has been set.
func (o *CoreGetUserDatesRequest) HasContextid() bool {
	if o != nil && !IsNil(o.Contextid) {
		return true
	}

	return false
}

// SetContextid gets a reference to the given int32 and assigns it to the Contextid field.
func (o *CoreGetUserDatesRequest) SetContextid(v int32) {
	o.Contextid = &v
}

// GetContextlevel returns the Contextlevel field value if set, zero value otherwise.
func (o *CoreGetUserDatesRequest) GetContextlevel() string {
	if o == nil || IsNil(o.Contextlevel) {
		var ret string
		return ret
	}
	return *o.Contextlevel
}

// GetContextlevelOk returns a tuple with the Contextlevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGetUserDatesRequest) GetContextlevelOk() (*string, bool) {
	if o == nil || IsNil(o.Contextlevel) {
		return nil, false
	}
	return o.Contextlevel, true
}

// HasContextlevel returns a boolean if a field has been set.
func (o *CoreGetUserDatesRequest) HasContextlevel() bool {
	if o != nil && !IsNil(o.Contextlevel) {
		return true
	}

	return false
}

// SetContextlevel gets a reference to the given string and assigns it to the Contextlevel field.
func (o *CoreGetUserDatesRequest) SetContextlevel(v string) {
	o.Contextlevel = &v
}

// GetInstanceid returns the Instanceid field value if set, zero value otherwise.
func (o *CoreGetUserDatesRequest) GetInstanceid() int32 {
	if o == nil || IsNil(o.Instanceid) {
		var ret int32
		return ret
	}
	return *o.Instanceid
}

// GetInstanceidOk returns a tuple with the Instanceid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGetUserDatesRequest) GetInstanceidOk() (*int32, bool) {
	if o == nil || IsNil(o.Instanceid) {
		return nil, false
	}
	return o.Instanceid, true
}

// HasInstanceid returns a boolean if a field has been set.
func (o *CoreGetUserDatesRequest) HasInstanceid() bool {
	if o != nil && !IsNil(o.Instanceid) {
		return true
	}

	return false
}

// SetInstanceid gets a reference to the given int32 and assigns it to the Instanceid field.
func (o *CoreGetUserDatesRequest) SetInstanceid(v int32) {
	o.Instanceid = &v
}

// GetTimestamps returns the Timestamps field value
func (o *CoreGetUserDatesRequest) GetTimestamps() []CoreGetUserDatesRequestTimestampsInner {
	if o == nil {
		var ret []CoreGetUserDatesRequestTimestampsInner
		return ret
	}

	return o.Timestamps
}

// GetTimestampsOk returns a tuple with the Timestamps field value
// and a boolean to check if the value has been set.
func (o *CoreGetUserDatesRequest) GetTimestampsOk() ([]CoreGetUserDatesRequestTimestampsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timestamps, true
}

// SetTimestamps sets field value
func (o *CoreGetUserDatesRequest) SetTimestamps(v []CoreGetUserDatesRequestTimestampsInner) {
	o.Timestamps = v
}

func (o CoreGetUserDatesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreGetUserDatesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Contextid) {
		toSerialize["contextid"] = o.Contextid
	}
	if !IsNil(o.Contextlevel) {
		toSerialize["contextlevel"] = o.Contextlevel
	}
	if !IsNil(o.Instanceid) {
		toSerialize["instanceid"] = o.Instanceid
	}
	toSerialize["timestamps"] = o.Timestamps
	return toSerialize, nil
}

func (o *CoreGetUserDatesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"timestamps",
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

	varCoreGetUserDatesRequest := _CoreGetUserDatesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreGetUserDatesRequest)

	if err != nil {
		return err
	}

	*o = CoreGetUserDatesRequest(varCoreGetUserDatesRequest)

	return err
}

type NullableCoreGetUserDatesRequest struct {
	value *CoreGetUserDatesRequest
	isSet bool
}

func (v NullableCoreGetUserDatesRequest) Get() *CoreGetUserDatesRequest {
	return v.value
}

func (v *NullableCoreGetUserDatesRequest) Set(val *CoreGetUserDatesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreGetUserDatesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreGetUserDatesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreGetUserDatesRequest(val *CoreGetUserDatesRequest) *NullableCoreGetUserDatesRequest {
	return &NullableCoreGetUserDatesRequest{value: val, isSet: true}
}

func (v NullableCoreGetUserDatesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreGetUserDatesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

