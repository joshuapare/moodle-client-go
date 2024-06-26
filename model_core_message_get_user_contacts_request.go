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

// checks if the CoreMessageGetUserContactsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreMessageGetUserContactsRequest{}

// CoreMessageGetUserContactsRequest struct for CoreMessageGetUserContactsRequest
type CoreMessageGetUserContactsRequest struct {
	// Limit from
	Limitfrom *int32 `json:"limitfrom,omitempty"`
	// Limit number
	Limitnum *int32 `json:"limitnum,omitempty"`
	// The id of the user who we retrieving the contacts for
	Userid int32 `json:"userid"`
}

type _CoreMessageGetUserContactsRequest CoreMessageGetUserContactsRequest

// NewCoreMessageGetUserContactsRequest instantiates a new CoreMessageGetUserContactsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreMessageGetUserContactsRequest(userid int32) *CoreMessageGetUserContactsRequest {
	this := CoreMessageGetUserContactsRequest{}
	var limitfrom int32 = 0
	this.Limitfrom = &limitfrom
	var limitnum int32 = 0
	this.Limitnum = &limitnum
	this.Userid = userid
	return &this
}

// NewCoreMessageGetUserContactsRequestWithDefaults instantiates a new CoreMessageGetUserContactsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreMessageGetUserContactsRequestWithDefaults() *CoreMessageGetUserContactsRequest {
	this := CoreMessageGetUserContactsRequest{}
	var limitfrom int32 = 0
	this.Limitfrom = &limitfrom
	var limitnum int32 = 0
	this.Limitnum = &limitnum
	var userid int32 = null
	this.Userid = userid
	return &this
}

// GetLimitfrom returns the Limitfrom field value if set, zero value otherwise.
func (o *CoreMessageGetUserContactsRequest) GetLimitfrom() int32 {
	if o == nil || IsNil(o.Limitfrom) {
		var ret int32
		return ret
	}
	return *o.Limitfrom
}

// GetLimitfromOk returns a tuple with the Limitfrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetUserContactsRequest) GetLimitfromOk() (*int32, bool) {
	if o == nil || IsNil(o.Limitfrom) {
		return nil, false
	}
	return o.Limitfrom, true
}

// HasLimitfrom returns a boolean if a field has been set.
func (o *CoreMessageGetUserContactsRequest) HasLimitfrom() bool {
	if o != nil && !IsNil(o.Limitfrom) {
		return true
	}

	return false
}

// SetLimitfrom gets a reference to the given int32 and assigns it to the Limitfrom field.
func (o *CoreMessageGetUserContactsRequest) SetLimitfrom(v int32) {
	o.Limitfrom = &v
}

// GetLimitnum returns the Limitnum field value if set, zero value otherwise.
func (o *CoreMessageGetUserContactsRequest) GetLimitnum() int32 {
	if o == nil || IsNil(o.Limitnum) {
		var ret int32
		return ret
	}
	return *o.Limitnum
}

// GetLimitnumOk returns a tuple with the Limitnum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetUserContactsRequest) GetLimitnumOk() (*int32, bool) {
	if o == nil || IsNil(o.Limitnum) {
		return nil, false
	}
	return o.Limitnum, true
}

// HasLimitnum returns a boolean if a field has been set.
func (o *CoreMessageGetUserContactsRequest) HasLimitnum() bool {
	if o != nil && !IsNil(o.Limitnum) {
		return true
	}

	return false
}

// SetLimitnum gets a reference to the given int32 and assigns it to the Limitnum field.
func (o *CoreMessageGetUserContactsRequest) SetLimitnum(v int32) {
	o.Limitnum = &v
}

// GetUserid returns the Userid field value
func (o *CoreMessageGetUserContactsRequest) GetUserid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Userid
}

// GetUseridOk returns a tuple with the Userid field value
// and a boolean to check if the value has been set.
func (o *CoreMessageGetUserContactsRequest) GetUseridOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Userid, true
}

// SetUserid sets field value
func (o *CoreMessageGetUserContactsRequest) SetUserid(v int32) {
	o.Userid = v
}

func (o CoreMessageGetUserContactsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreMessageGetUserContactsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Limitfrom) {
		toSerialize["limitfrom"] = o.Limitfrom
	}
	if !IsNil(o.Limitnum) {
		toSerialize["limitnum"] = o.Limitnum
	}
	toSerialize["userid"] = o.Userid
	return toSerialize, nil
}

func (o *CoreMessageGetUserContactsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varCoreMessageGetUserContactsRequest := _CoreMessageGetUserContactsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreMessageGetUserContactsRequest)

	if err != nil {
		return err
	}

	*o = CoreMessageGetUserContactsRequest(varCoreMessageGetUserContactsRequest)

	return err
}

type NullableCoreMessageGetUserContactsRequest struct {
	value *CoreMessageGetUserContactsRequest
	isSet bool
}

func (v NullableCoreMessageGetUserContactsRequest) Get() *CoreMessageGetUserContactsRequest {
	return v.value
}

func (v *NullableCoreMessageGetUserContactsRequest) Set(val *CoreMessageGetUserContactsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreMessageGetUserContactsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreMessageGetUserContactsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreMessageGetUserContactsRequest(val *CoreMessageGetUserContactsRequest) *NullableCoreMessageGetUserContactsRequest {
	return &NullableCoreMessageGetUserContactsRequest{value: val, isSet: true}
}

func (v NullableCoreMessageGetUserContactsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreMessageGetUserContactsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


