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

// checks if the QuizaccessSebValidateQuizKeysRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuizaccessSebValidateQuizKeysRequest{}

// QuizaccessSebValidateQuizKeysRequest struct for QuizaccessSebValidateQuizKeysRequest
type QuizaccessSebValidateQuizKeysRequest struct {
	// SEB browser exam key
	Browserexamkey *string `json:"browserexamkey,omitempty"`
	// Course module ID
	Cmid int32 `json:"cmid"`
	// SEB config key
	Configkey *string `json:"configkey,omitempty"`
	// Page URL to check
	Url string `json:"url"`
}

type _QuizaccessSebValidateQuizKeysRequest QuizaccessSebValidateQuizKeysRequest

// NewQuizaccessSebValidateQuizKeysRequest instantiates a new QuizaccessSebValidateQuizKeysRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuizaccessSebValidateQuizKeysRequest(cmid int32, url string) *QuizaccessSebValidateQuizKeysRequest {
	this := QuizaccessSebValidateQuizKeysRequest{}
	var browserexamkey string = "null"
	this.Browserexamkey = &browserexamkey
	this.Cmid = cmid
	var configkey string = "null"
	this.Configkey = &configkey
	this.Url = url
	return &this
}

// NewQuizaccessSebValidateQuizKeysRequestWithDefaults instantiates a new QuizaccessSebValidateQuizKeysRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuizaccessSebValidateQuizKeysRequestWithDefaults() *QuizaccessSebValidateQuizKeysRequest {
	this := QuizaccessSebValidateQuizKeysRequest{}
	var browserexamkey string = "null"
	this.Browserexamkey = &browserexamkey
	var cmid int32 = null
	this.Cmid = cmid
	var configkey string = "null"
	this.Configkey = &configkey
	var url string = "null"
	this.Url = url
	return &this
}

// GetBrowserexamkey returns the Browserexamkey field value if set, zero value otherwise.
func (o *QuizaccessSebValidateQuizKeysRequest) GetBrowserexamkey() string {
	if o == nil || IsNil(o.Browserexamkey) {
		var ret string
		return ret
	}
	return *o.Browserexamkey
}

// GetBrowserexamkeyOk returns a tuple with the Browserexamkey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuizaccessSebValidateQuizKeysRequest) GetBrowserexamkeyOk() (*string, bool) {
	if o == nil || IsNil(o.Browserexamkey) {
		return nil, false
	}
	return o.Browserexamkey, true
}

// HasBrowserexamkey returns a boolean if a field has been set.
func (o *QuizaccessSebValidateQuizKeysRequest) HasBrowserexamkey() bool {
	if o != nil && !IsNil(o.Browserexamkey) {
		return true
	}

	return false
}

// SetBrowserexamkey gets a reference to the given string and assigns it to the Browserexamkey field.
func (o *QuizaccessSebValidateQuizKeysRequest) SetBrowserexamkey(v string) {
	o.Browserexamkey = &v
}

// GetCmid returns the Cmid field value
func (o *QuizaccessSebValidateQuizKeysRequest) GetCmid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Cmid
}

// GetCmidOk returns a tuple with the Cmid field value
// and a boolean to check if the value has been set.
func (o *QuizaccessSebValidateQuizKeysRequest) GetCmidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cmid, true
}

// SetCmid sets field value
func (o *QuizaccessSebValidateQuizKeysRequest) SetCmid(v int32) {
	o.Cmid = v
}

// GetConfigkey returns the Configkey field value if set, zero value otherwise.
func (o *QuizaccessSebValidateQuizKeysRequest) GetConfigkey() string {
	if o == nil || IsNil(o.Configkey) {
		var ret string
		return ret
	}
	return *o.Configkey
}

// GetConfigkeyOk returns a tuple with the Configkey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuizaccessSebValidateQuizKeysRequest) GetConfigkeyOk() (*string, bool) {
	if o == nil || IsNil(o.Configkey) {
		return nil, false
	}
	return o.Configkey, true
}

// HasConfigkey returns a boolean if a field has been set.
func (o *QuizaccessSebValidateQuizKeysRequest) HasConfigkey() bool {
	if o != nil && !IsNil(o.Configkey) {
		return true
	}

	return false
}

// SetConfigkey gets a reference to the given string and assigns it to the Configkey field.
func (o *QuizaccessSebValidateQuizKeysRequest) SetConfigkey(v string) {
	o.Configkey = &v
}

// GetUrl returns the Url field value
func (o *QuizaccessSebValidateQuizKeysRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *QuizaccessSebValidateQuizKeysRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *QuizaccessSebValidateQuizKeysRequest) SetUrl(v string) {
	o.Url = v
}

func (o QuizaccessSebValidateQuizKeysRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuizaccessSebValidateQuizKeysRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Browserexamkey) {
		toSerialize["browserexamkey"] = o.Browserexamkey
	}
	toSerialize["cmid"] = o.Cmid
	if !IsNil(o.Configkey) {
		toSerialize["configkey"] = o.Configkey
	}
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *QuizaccessSebValidateQuizKeysRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cmid",
		"url",
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

	varQuizaccessSebValidateQuizKeysRequest := _QuizaccessSebValidateQuizKeysRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQuizaccessSebValidateQuizKeysRequest)

	if err != nil {
		return err
	}

	*o = QuizaccessSebValidateQuizKeysRequest(varQuizaccessSebValidateQuizKeysRequest)

	return err
}

type NullableQuizaccessSebValidateQuizKeysRequest struct {
	value *QuizaccessSebValidateQuizKeysRequest
	isSet bool
}

func (v NullableQuizaccessSebValidateQuizKeysRequest) Get() *QuizaccessSebValidateQuizKeysRequest {
	return v.value
}

func (v *NullableQuizaccessSebValidateQuizKeysRequest) Set(val *QuizaccessSebValidateQuizKeysRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableQuizaccessSebValidateQuizKeysRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableQuizaccessSebValidateQuizKeysRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuizaccessSebValidateQuizKeysRequest(val *QuizaccessSebValidateQuizKeysRequest) *NullableQuizaccessSebValidateQuizKeysRequest {
	return &NullableQuizaccessSebValidateQuizKeysRequest{value: val, isSet: true}
}

func (v NullableQuizaccessSebValidateQuizKeysRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuizaccessSebValidateQuizKeysRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


