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

// checks if the ToolIomadpolicyGetIomadpolicyVersionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ToolIomadpolicyGetIomadpolicyVersionRequest{}

// ToolIomadpolicyGetIomadpolicyVersionRequest struct for ToolIomadpolicyGetIomadpolicyVersionRequest
type ToolIomadpolicyGetIomadpolicyVersionRequest struct {
	// The id of user on whose behalf the user is viewing the iomadpolicy
	Behalfid *int32 `json:"behalfid,omitempty"`
	// The iomadpolicy version ID
	Versionid int32 `json:"versionid"`
}

type _ToolIomadpolicyGetIomadpolicyVersionRequest ToolIomadpolicyGetIomadpolicyVersionRequest

// NewToolIomadpolicyGetIomadpolicyVersionRequest instantiates a new ToolIomadpolicyGetIomadpolicyVersionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToolIomadpolicyGetIomadpolicyVersionRequest(versionid int32) *ToolIomadpolicyGetIomadpolicyVersionRequest {
	this := ToolIomadpolicyGetIomadpolicyVersionRequest{}
	var behalfid int32 = 0
	this.Behalfid = &behalfid
	this.Versionid = versionid
	return &this
}

// NewToolIomadpolicyGetIomadpolicyVersionRequestWithDefaults instantiates a new ToolIomadpolicyGetIomadpolicyVersionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToolIomadpolicyGetIomadpolicyVersionRequestWithDefaults() *ToolIomadpolicyGetIomadpolicyVersionRequest {
	this := ToolIomadpolicyGetIomadpolicyVersionRequest{}
	var behalfid int32 = 0
	this.Behalfid = &behalfid
	var versionid int32 = null
	this.Versionid = versionid
	return &this
}

// GetBehalfid returns the Behalfid field value if set, zero value otherwise.
func (o *ToolIomadpolicyGetIomadpolicyVersionRequest) GetBehalfid() int32 {
	if o == nil || IsNil(o.Behalfid) {
		var ret int32
		return ret
	}
	return *o.Behalfid
}

// GetBehalfidOk returns a tuple with the Behalfid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolIomadpolicyGetIomadpolicyVersionRequest) GetBehalfidOk() (*int32, bool) {
	if o == nil || IsNil(o.Behalfid) {
		return nil, false
	}
	return o.Behalfid, true
}

// HasBehalfid returns a boolean if a field has been set.
func (o *ToolIomadpolicyGetIomadpolicyVersionRequest) HasBehalfid() bool {
	if o != nil && !IsNil(o.Behalfid) {
		return true
	}

	return false
}

// SetBehalfid gets a reference to the given int32 and assigns it to the Behalfid field.
func (o *ToolIomadpolicyGetIomadpolicyVersionRequest) SetBehalfid(v int32) {
	o.Behalfid = &v
}

// GetVersionid returns the Versionid field value
func (o *ToolIomadpolicyGetIomadpolicyVersionRequest) GetVersionid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Versionid
}

// GetVersionidOk returns a tuple with the Versionid field value
// and a boolean to check if the value has been set.
func (o *ToolIomadpolicyGetIomadpolicyVersionRequest) GetVersionidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Versionid, true
}

// SetVersionid sets field value
func (o *ToolIomadpolicyGetIomadpolicyVersionRequest) SetVersionid(v int32) {
	o.Versionid = v
}

func (o ToolIomadpolicyGetIomadpolicyVersionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToolIomadpolicyGetIomadpolicyVersionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Behalfid) {
		toSerialize["behalfid"] = o.Behalfid
	}
	toSerialize["versionid"] = o.Versionid
	return toSerialize, nil
}

func (o *ToolIomadpolicyGetIomadpolicyVersionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"versionid",
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

	varToolIomadpolicyGetIomadpolicyVersionRequest := _ToolIomadpolicyGetIomadpolicyVersionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varToolIomadpolicyGetIomadpolicyVersionRequest)

	if err != nil {
		return err
	}

	*o = ToolIomadpolicyGetIomadpolicyVersionRequest(varToolIomadpolicyGetIomadpolicyVersionRequest)

	return err
}

type NullableToolIomadpolicyGetIomadpolicyVersionRequest struct {
	value *ToolIomadpolicyGetIomadpolicyVersionRequest
	isSet bool
}

func (v NullableToolIomadpolicyGetIomadpolicyVersionRequest) Get() *ToolIomadpolicyGetIomadpolicyVersionRequest {
	return v.value
}

func (v *NullableToolIomadpolicyGetIomadpolicyVersionRequest) Set(val *ToolIomadpolicyGetIomadpolicyVersionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableToolIomadpolicyGetIomadpolicyVersionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableToolIomadpolicyGetIomadpolicyVersionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToolIomadpolicyGetIomadpolicyVersionRequest(val *ToolIomadpolicyGetIomadpolicyVersionRequest) *NullableToolIomadpolicyGetIomadpolicyVersionRequest {
	return &NullableToolIomadpolicyGetIomadpolicyVersionRequest{value: val, isSet: true}
}

func (v NullableToolIomadpolicyGetIomadpolicyVersionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToolIomadpolicyGetIomadpolicyVersionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


