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

// checks if the ModChatGetChatUsersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModChatGetChatUsersRequest{}

// ModChatGetChatUsersRequest struct for ModChatGetChatUsersRequest
type ModChatGetChatUsersRequest struct {
	// chat session id (obtained via mod_chat_login_user)
	Chatsid string `json:"chatsid"`
}

type _ModChatGetChatUsersRequest ModChatGetChatUsersRequest

// NewModChatGetChatUsersRequest instantiates a new ModChatGetChatUsersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModChatGetChatUsersRequest(chatsid string) *ModChatGetChatUsersRequest {
	this := ModChatGetChatUsersRequest{}
	this.Chatsid = chatsid
	return &this
}

// NewModChatGetChatUsersRequestWithDefaults instantiates a new ModChatGetChatUsersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModChatGetChatUsersRequestWithDefaults() *ModChatGetChatUsersRequest {
	this := ModChatGetChatUsersRequest{}
	return &this
}

// GetChatsid returns the Chatsid field value
func (o *ModChatGetChatUsersRequest) GetChatsid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Chatsid
}

// GetChatsidOk returns a tuple with the Chatsid field value
// and a boolean to check if the value has been set.
func (o *ModChatGetChatUsersRequest) GetChatsidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chatsid, true
}

// SetChatsid sets field value
func (o *ModChatGetChatUsersRequest) SetChatsid(v string) {
	o.Chatsid = v
}

func (o ModChatGetChatUsersRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModChatGetChatUsersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chatsid"] = o.Chatsid
	return toSerialize, nil
}

func (o *ModChatGetChatUsersRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"chatsid",
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

	varModChatGetChatUsersRequest := _ModChatGetChatUsersRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModChatGetChatUsersRequest)

	if err != nil {
		return err
	}

	*o = ModChatGetChatUsersRequest(varModChatGetChatUsersRequest)

	return err
}

type NullableModChatGetChatUsersRequest struct {
	value *ModChatGetChatUsersRequest
	isSet bool
}

func (v NullableModChatGetChatUsersRequest) Get() *ModChatGetChatUsersRequest {
	return v.value
}

func (v *NullableModChatGetChatUsersRequest) Set(val *ModChatGetChatUsersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModChatGetChatUsersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModChatGetChatUsersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModChatGetChatUsersRequest(val *ModChatGetChatUsersRequest) *NullableModChatGetChatUsersRequest {
	return &NullableModChatGetChatUsersRequest{value: val, isSet: true}
}

func (v NullableModChatGetChatUsersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModChatGetChatUsersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


