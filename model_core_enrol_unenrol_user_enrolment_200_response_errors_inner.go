/*
Moodle Webservice API

Auto-generated OpenAPI spec for Moodle's Webservice API. Requires installation of the Catalyst RESTful API plugin: https://github.com/catalyst/moodle-webservice_restful.

API version: 4.3.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package moodleclient

import (
	"encoding/json"
)

// checks if the CoreEnrolUnenrolUserEnrolment200ResponseErrorsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreEnrolUnenrolUserEnrolment200ResponseErrorsInner{}

// CoreEnrolUnenrolUserEnrolment200ResponseErrorsInner struct for CoreEnrolUnenrolUserEnrolment200ResponseErrorsInner
type CoreEnrolUnenrolUserEnrolment200ResponseErrorsInner struct {
	// The data that failed the validation
	Key *string `json:"key,omitempty"`
	// The error message
	Message *string `json:"message,omitempty"`
}

// NewCoreEnrolUnenrolUserEnrolment200ResponseErrorsInner instantiates a new CoreEnrolUnenrolUserEnrolment200ResponseErrorsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreEnrolUnenrolUserEnrolment200ResponseErrorsInner() *CoreEnrolUnenrolUserEnrolment200ResponseErrorsInner {
	this := CoreEnrolUnenrolUserEnrolment200ResponseErrorsInner{}
	var key string = "null"
	this.Key = &key
	var message string = "null"
	this.Message = &message
	return &this
}

// NewCoreEnrolUnenrolUserEnrolment200ResponseErrorsInnerWithDefaults instantiates a new CoreEnrolUnenrolUserEnrolment200ResponseErrorsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreEnrolUnenrolUserEnrolment200ResponseErrorsInnerWithDefaults() *CoreEnrolUnenrolUserEnrolment200ResponseErrorsInner {
	this := CoreEnrolUnenrolUserEnrolment200ResponseErrorsInner{}
	var key string = "null"
	this.Key = &key
	var message string = "null"
	this.Message = &message
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *CoreEnrolUnenrolUserEnrolment200ResponseErrorsInner) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreEnrolUnenrolUserEnrolment200ResponseErrorsInner) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *CoreEnrolUnenrolUserEnrolment200ResponseErrorsInner) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *CoreEnrolUnenrolUserEnrolment200ResponseErrorsInner) SetKey(v string) {
	o.Key = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CoreEnrolUnenrolUserEnrolment200ResponseErrorsInner) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreEnrolUnenrolUserEnrolment200ResponseErrorsInner) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CoreEnrolUnenrolUserEnrolment200ResponseErrorsInner) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CoreEnrolUnenrolUserEnrolment200ResponseErrorsInner) SetMessage(v string) {
	o.Message = &v
}

func (o CoreEnrolUnenrolUserEnrolment200ResponseErrorsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreEnrolUnenrolUserEnrolment200ResponseErrorsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableCoreEnrolUnenrolUserEnrolment200ResponseErrorsInner struct {
	value *CoreEnrolUnenrolUserEnrolment200ResponseErrorsInner
	isSet bool
}

func (v NullableCoreEnrolUnenrolUserEnrolment200ResponseErrorsInner) Get() *CoreEnrolUnenrolUserEnrolment200ResponseErrorsInner {
	return v.value
}

func (v *NullableCoreEnrolUnenrolUserEnrolment200ResponseErrorsInner) Set(val *CoreEnrolUnenrolUserEnrolment200ResponseErrorsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreEnrolUnenrolUserEnrolment200ResponseErrorsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreEnrolUnenrolUserEnrolment200ResponseErrorsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreEnrolUnenrolUserEnrolment200ResponseErrorsInner(val *CoreEnrolUnenrolUserEnrolment200ResponseErrorsInner) *NullableCoreEnrolUnenrolUserEnrolment200ResponseErrorsInner {
	return &NullableCoreEnrolUnenrolUserEnrolment200ResponseErrorsInner{value: val, isSet: true}
}

func (v NullableCoreEnrolUnenrolUserEnrolment200ResponseErrorsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreEnrolUnenrolUserEnrolment200ResponseErrorsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


