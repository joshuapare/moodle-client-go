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

// checks if the CoreMessageUnmuteConversationsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreMessageUnmuteConversationsRequest{}

// CoreMessageUnmuteConversationsRequest struct for CoreMessageUnmuteConversationsRequest
type CoreMessageUnmuteConversationsRequest struct {
	Conversationids []map[string]interface{} `json:"conversationids"`
	// The id of the user who is unblocking
	Userid int32 `json:"userid"`
}

type _CoreMessageUnmuteConversationsRequest CoreMessageUnmuteConversationsRequest

// NewCoreMessageUnmuteConversationsRequest instantiates a new CoreMessageUnmuteConversationsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreMessageUnmuteConversationsRequest(conversationids []map[string]interface{}, userid int32) *CoreMessageUnmuteConversationsRequest {
	this := CoreMessageUnmuteConversationsRequest{}
	this.Conversationids = conversationids
	this.Userid = userid
	return &this
}

// NewCoreMessageUnmuteConversationsRequestWithDefaults instantiates a new CoreMessageUnmuteConversationsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreMessageUnmuteConversationsRequestWithDefaults() *CoreMessageUnmuteConversationsRequest {
	this := CoreMessageUnmuteConversationsRequest{}
	return &this
}

// GetConversationids returns the Conversationids field value
func (o *CoreMessageUnmuteConversationsRequest) GetConversationids() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Conversationids
}

// GetConversationidsOk returns a tuple with the Conversationids field value
// and a boolean to check if the value has been set.
func (o *CoreMessageUnmuteConversationsRequest) GetConversationidsOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Conversationids, true
}

// SetConversationids sets field value
func (o *CoreMessageUnmuteConversationsRequest) SetConversationids(v []map[string]interface{}) {
	o.Conversationids = v
}

// GetUserid returns the Userid field value
func (o *CoreMessageUnmuteConversationsRequest) GetUserid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Userid
}

// GetUseridOk returns a tuple with the Userid field value
// and a boolean to check if the value has been set.
func (o *CoreMessageUnmuteConversationsRequest) GetUseridOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Userid, true
}

// SetUserid sets field value
func (o *CoreMessageUnmuteConversationsRequest) SetUserid(v int32) {
	o.Userid = v
}

func (o CoreMessageUnmuteConversationsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreMessageUnmuteConversationsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["conversationids"] = o.Conversationids
	toSerialize["userid"] = o.Userid
	return toSerialize, nil
}

func (o *CoreMessageUnmuteConversationsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"conversationids",
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

	varCoreMessageUnmuteConversationsRequest := _CoreMessageUnmuteConversationsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreMessageUnmuteConversationsRequest)

	if err != nil {
		return err
	}

	*o = CoreMessageUnmuteConversationsRequest(varCoreMessageUnmuteConversationsRequest)

	return err
}

type NullableCoreMessageUnmuteConversationsRequest struct {
	value *CoreMessageUnmuteConversationsRequest
	isSet bool
}

func (v NullableCoreMessageUnmuteConversationsRequest) Get() *CoreMessageUnmuteConversationsRequest {
	return v.value
}

func (v *NullableCoreMessageUnmuteConversationsRequest) Set(val *CoreMessageUnmuteConversationsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreMessageUnmuteConversationsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreMessageUnmuteConversationsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreMessageUnmuteConversationsRequest(val *CoreMessageUnmuteConversationsRequest) *NullableCoreMessageUnmuteConversationsRequest {
	return &NullableCoreMessageUnmuteConversationsRequest{value: val, isSet: true}
}

func (v NullableCoreMessageUnmuteConversationsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreMessageUnmuteConversationsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


