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

// checks if the ToolMobileGetTokensForQrLoginRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ToolMobileGetTokensForQrLoginRequest{}

// ToolMobileGetTokensForQrLoginRequest struct for ToolMobileGetTokensForQrLoginRequest
type ToolMobileGetTokensForQrLoginRequest struct {
	// The user key for validating the request.
	Qrloginkey string `json:"qrloginkey"`
	// The user the key belongs to.
	Userid int32 `json:"userid"`
}

type _ToolMobileGetTokensForQrLoginRequest ToolMobileGetTokensForQrLoginRequest

// NewToolMobileGetTokensForQrLoginRequest instantiates a new ToolMobileGetTokensForQrLoginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToolMobileGetTokensForQrLoginRequest(qrloginkey string, userid int32) *ToolMobileGetTokensForQrLoginRequest {
	this := ToolMobileGetTokensForQrLoginRequest{}
	this.Qrloginkey = qrloginkey
	this.Userid = userid
	return &this
}

// NewToolMobileGetTokensForQrLoginRequestWithDefaults instantiates a new ToolMobileGetTokensForQrLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToolMobileGetTokensForQrLoginRequestWithDefaults() *ToolMobileGetTokensForQrLoginRequest {
	this := ToolMobileGetTokensForQrLoginRequest{}
	var qrloginkey string = "null"
	this.Qrloginkey = qrloginkey
	var userid int32 = null
	this.Userid = userid
	return &this
}

// GetQrloginkey returns the Qrloginkey field value
func (o *ToolMobileGetTokensForQrLoginRequest) GetQrloginkey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Qrloginkey
}

// GetQrloginkeyOk returns a tuple with the Qrloginkey field value
// and a boolean to check if the value has been set.
func (o *ToolMobileGetTokensForQrLoginRequest) GetQrloginkeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Qrloginkey, true
}

// SetQrloginkey sets field value
func (o *ToolMobileGetTokensForQrLoginRequest) SetQrloginkey(v string) {
	o.Qrloginkey = v
}

// GetUserid returns the Userid field value
func (o *ToolMobileGetTokensForQrLoginRequest) GetUserid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Userid
}

// GetUseridOk returns a tuple with the Userid field value
// and a boolean to check if the value has been set.
func (o *ToolMobileGetTokensForQrLoginRequest) GetUseridOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Userid, true
}

// SetUserid sets field value
func (o *ToolMobileGetTokensForQrLoginRequest) SetUserid(v int32) {
	o.Userid = v
}

func (o ToolMobileGetTokensForQrLoginRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToolMobileGetTokensForQrLoginRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["qrloginkey"] = o.Qrloginkey
	toSerialize["userid"] = o.Userid
	return toSerialize, nil
}

func (o *ToolMobileGetTokensForQrLoginRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"qrloginkey",
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

	varToolMobileGetTokensForQrLoginRequest := _ToolMobileGetTokensForQrLoginRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varToolMobileGetTokensForQrLoginRequest)

	if err != nil {
		return err
	}

	*o = ToolMobileGetTokensForQrLoginRequest(varToolMobileGetTokensForQrLoginRequest)

	return err
}

type NullableToolMobileGetTokensForQrLoginRequest struct {
	value *ToolMobileGetTokensForQrLoginRequest
	isSet bool
}

func (v NullableToolMobileGetTokensForQrLoginRequest) Get() *ToolMobileGetTokensForQrLoginRequest {
	return v.value
}

func (v *NullableToolMobileGetTokensForQrLoginRequest) Set(val *ToolMobileGetTokensForQrLoginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableToolMobileGetTokensForQrLoginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableToolMobileGetTokensForQrLoginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToolMobileGetTokensForQrLoginRequest(val *ToolMobileGetTokensForQrLoginRequest) *NullableToolMobileGetTokensForQrLoginRequest {
	return &NullableToolMobileGetTokensForQrLoginRequest{value: val, isSet: true}
}

func (v NullableToolMobileGetTokensForQrLoginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToolMobileGetTokensForQrLoginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


