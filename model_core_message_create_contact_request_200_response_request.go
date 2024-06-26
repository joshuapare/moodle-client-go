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

// checks if the CoreMessageCreateContactRequest200ResponseRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreMessageCreateContactRequest200ResponseRequest{}

// CoreMessageCreateContactRequest200ResponseRequest struct for CoreMessageCreateContactRequest200ResponseRequest
type CoreMessageCreateContactRequest200ResponseRequest struct {
	// Message id
	Id int32 `json:"id"`
	// User to id
	Requesteduserid int32 `json:"requesteduserid"`
	// Time created
	Timecreated int32 `json:"timecreated"`
	// User from id
	Userid int32 `json:"userid"`
}

type _CoreMessageCreateContactRequest200ResponseRequest CoreMessageCreateContactRequest200ResponseRequest

// NewCoreMessageCreateContactRequest200ResponseRequest instantiates a new CoreMessageCreateContactRequest200ResponseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreMessageCreateContactRequest200ResponseRequest(id int32, requesteduserid int32, timecreated int32, userid int32) *CoreMessageCreateContactRequest200ResponseRequest {
	this := CoreMessageCreateContactRequest200ResponseRequest{}
	this.Id = id
	this.Requesteduserid = requesteduserid
	this.Timecreated = timecreated
	this.Userid = userid
	return &this
}

// NewCoreMessageCreateContactRequest200ResponseRequestWithDefaults instantiates a new CoreMessageCreateContactRequest200ResponseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreMessageCreateContactRequest200ResponseRequestWithDefaults() *CoreMessageCreateContactRequest200ResponseRequest {
	this := CoreMessageCreateContactRequest200ResponseRequest{}
	var id int32 = null
	this.Id = id
	var requesteduserid int32 = null
	this.Requesteduserid = requesteduserid
	var userid int32 = null
	this.Userid = userid
	return &this
}

// GetId returns the Id field value
func (o *CoreMessageCreateContactRequest200ResponseRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CoreMessageCreateContactRequest200ResponseRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CoreMessageCreateContactRequest200ResponseRequest) SetId(v int32) {
	o.Id = v
}

// GetRequesteduserid returns the Requesteduserid field value
func (o *CoreMessageCreateContactRequest200ResponseRequest) GetRequesteduserid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Requesteduserid
}

// GetRequesteduseridOk returns a tuple with the Requesteduserid field value
// and a boolean to check if the value has been set.
func (o *CoreMessageCreateContactRequest200ResponseRequest) GetRequesteduseridOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requesteduserid, true
}

// SetRequesteduserid sets field value
func (o *CoreMessageCreateContactRequest200ResponseRequest) SetRequesteduserid(v int32) {
	o.Requesteduserid = v
}

// GetTimecreated returns the Timecreated field value
func (o *CoreMessageCreateContactRequest200ResponseRequest) GetTimecreated() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timecreated
}

// GetTimecreatedOk returns a tuple with the Timecreated field value
// and a boolean to check if the value has been set.
func (o *CoreMessageCreateContactRequest200ResponseRequest) GetTimecreatedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timecreated, true
}

// SetTimecreated sets field value
func (o *CoreMessageCreateContactRequest200ResponseRequest) SetTimecreated(v int32) {
	o.Timecreated = v
}

// GetUserid returns the Userid field value
func (o *CoreMessageCreateContactRequest200ResponseRequest) GetUserid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Userid
}

// GetUseridOk returns a tuple with the Userid field value
// and a boolean to check if the value has been set.
func (o *CoreMessageCreateContactRequest200ResponseRequest) GetUseridOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Userid, true
}

// SetUserid sets field value
func (o *CoreMessageCreateContactRequest200ResponseRequest) SetUserid(v int32) {
	o.Userid = v
}

func (o CoreMessageCreateContactRequest200ResponseRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreMessageCreateContactRequest200ResponseRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["requesteduserid"] = o.Requesteduserid
	toSerialize["timecreated"] = o.Timecreated
	toSerialize["userid"] = o.Userid
	return toSerialize, nil
}

func (o *CoreMessageCreateContactRequest200ResponseRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"requesteduserid",
		"timecreated",
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

	varCoreMessageCreateContactRequest200ResponseRequest := _CoreMessageCreateContactRequest200ResponseRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreMessageCreateContactRequest200ResponseRequest)

	if err != nil {
		return err
	}

	*o = CoreMessageCreateContactRequest200ResponseRequest(varCoreMessageCreateContactRequest200ResponseRequest)

	return err
}

type NullableCoreMessageCreateContactRequest200ResponseRequest struct {
	value *CoreMessageCreateContactRequest200ResponseRequest
	isSet bool
}

func (v NullableCoreMessageCreateContactRequest200ResponseRequest) Get() *CoreMessageCreateContactRequest200ResponseRequest {
	return v.value
}

func (v *NullableCoreMessageCreateContactRequest200ResponseRequest) Set(val *CoreMessageCreateContactRequest200ResponseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreMessageCreateContactRequest200ResponseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreMessageCreateContactRequest200ResponseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreMessageCreateContactRequest200ResponseRequest(val *CoreMessageCreateContactRequest200ResponseRequest) *NullableCoreMessageCreateContactRequest200ResponseRequest {
	return &NullableCoreMessageCreateContactRequest200ResponseRequest{value: val, isSet: true}
}

func (v NullableCoreMessageCreateContactRequest200ResponseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreMessageCreateContactRequest200ResponseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


