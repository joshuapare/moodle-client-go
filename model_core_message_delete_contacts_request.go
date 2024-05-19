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

// checks if the CoreMessageDeleteContactsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreMessageDeleteContactsRequest{}

// CoreMessageDeleteContactsRequest struct for CoreMessageDeleteContactsRequest
type CoreMessageDeleteContactsRequest struct {
	// The id of the user we are deleting the contacts for, 0 for the                     current user
	Userid *int32 `json:"userid,omitempty"`
	Userids []map[string]interface{} `json:"userids"`
}

type _CoreMessageDeleteContactsRequest CoreMessageDeleteContactsRequest

// NewCoreMessageDeleteContactsRequest instantiates a new CoreMessageDeleteContactsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreMessageDeleteContactsRequest(userids []map[string]interface{}) *CoreMessageDeleteContactsRequest {
	this := CoreMessageDeleteContactsRequest{}
	var userid int32 = 0
	this.Userid = &userid
	this.Userids = userids
	return &this
}

// NewCoreMessageDeleteContactsRequestWithDefaults instantiates a new CoreMessageDeleteContactsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreMessageDeleteContactsRequestWithDefaults() *CoreMessageDeleteContactsRequest {
	this := CoreMessageDeleteContactsRequest{}
	var userid int32 = 0
	this.Userid = &userid
	return &this
}

// GetUserid returns the Userid field value if set, zero value otherwise.
func (o *CoreMessageDeleteContactsRequest) GetUserid() int32 {
	if o == nil || IsNil(o.Userid) {
		var ret int32
		return ret
	}
	return *o.Userid
}

// GetUseridOk returns a tuple with the Userid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageDeleteContactsRequest) GetUseridOk() (*int32, bool) {
	if o == nil || IsNil(o.Userid) {
		return nil, false
	}
	return o.Userid, true
}

// HasUserid returns a boolean if a field has been set.
func (o *CoreMessageDeleteContactsRequest) HasUserid() bool {
	if o != nil && !IsNil(o.Userid) {
		return true
	}

	return false
}

// SetUserid gets a reference to the given int32 and assigns it to the Userid field.
func (o *CoreMessageDeleteContactsRequest) SetUserid(v int32) {
	o.Userid = &v
}

// GetUserids returns the Userids field value
func (o *CoreMessageDeleteContactsRequest) GetUserids() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Userids
}

// GetUseridsOk returns a tuple with the Userids field value
// and a boolean to check if the value has been set.
func (o *CoreMessageDeleteContactsRequest) GetUseridsOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Userids, true
}

// SetUserids sets field value
func (o *CoreMessageDeleteContactsRequest) SetUserids(v []map[string]interface{}) {
	o.Userids = v
}

func (o CoreMessageDeleteContactsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreMessageDeleteContactsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Userid) {
		toSerialize["userid"] = o.Userid
	}
	toSerialize["userids"] = o.Userids
	return toSerialize, nil
}

func (o *CoreMessageDeleteContactsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"userids",
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

	varCoreMessageDeleteContactsRequest := _CoreMessageDeleteContactsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreMessageDeleteContactsRequest)

	if err != nil {
		return err
	}

	*o = CoreMessageDeleteContactsRequest(varCoreMessageDeleteContactsRequest)

	return err
}

type NullableCoreMessageDeleteContactsRequest struct {
	value *CoreMessageDeleteContactsRequest
	isSet bool
}

func (v NullableCoreMessageDeleteContactsRequest) Get() *CoreMessageDeleteContactsRequest {
	return v.value
}

func (v *NullableCoreMessageDeleteContactsRequest) Set(val *CoreMessageDeleteContactsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreMessageDeleteContactsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreMessageDeleteContactsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreMessageDeleteContactsRequest(val *CoreMessageDeleteContactsRequest) *NullableCoreMessageDeleteContactsRequest {
	return &NullableCoreMessageDeleteContactsRequest{value: val, isSet: true}
}

func (v NullableCoreMessageDeleteContactsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreMessageDeleteContactsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


