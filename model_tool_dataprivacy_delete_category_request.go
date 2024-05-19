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

// checks if the ToolDataprivacyDeleteCategoryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ToolDataprivacyDeleteCategoryRequest{}

// ToolDataprivacyDeleteCategoryRequest struct for ToolDataprivacyDeleteCategoryRequest
type ToolDataprivacyDeleteCategoryRequest struct {
	// The category ID
	Id int32 `json:"id"`
}

type _ToolDataprivacyDeleteCategoryRequest ToolDataprivacyDeleteCategoryRequest

// NewToolDataprivacyDeleteCategoryRequest instantiates a new ToolDataprivacyDeleteCategoryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToolDataprivacyDeleteCategoryRequest(id int32) *ToolDataprivacyDeleteCategoryRequest {
	this := ToolDataprivacyDeleteCategoryRequest{}
	this.Id = id
	return &this
}

// NewToolDataprivacyDeleteCategoryRequestWithDefaults instantiates a new ToolDataprivacyDeleteCategoryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToolDataprivacyDeleteCategoryRequestWithDefaults() *ToolDataprivacyDeleteCategoryRequest {
	this := ToolDataprivacyDeleteCategoryRequest{}
	return &this
}

// GetId returns the Id field value
func (o *ToolDataprivacyDeleteCategoryRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ToolDataprivacyDeleteCategoryRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ToolDataprivacyDeleteCategoryRequest) SetId(v int32) {
	o.Id = v
}

func (o ToolDataprivacyDeleteCategoryRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToolDataprivacyDeleteCategoryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *ToolDataprivacyDeleteCategoryRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varToolDataprivacyDeleteCategoryRequest := _ToolDataprivacyDeleteCategoryRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varToolDataprivacyDeleteCategoryRequest)

	if err != nil {
		return err
	}

	*o = ToolDataprivacyDeleteCategoryRequest(varToolDataprivacyDeleteCategoryRequest)

	return err
}

type NullableToolDataprivacyDeleteCategoryRequest struct {
	value *ToolDataprivacyDeleteCategoryRequest
	isSet bool
}

func (v NullableToolDataprivacyDeleteCategoryRequest) Get() *ToolDataprivacyDeleteCategoryRequest {
	return v.value
}

func (v *NullableToolDataprivacyDeleteCategoryRequest) Set(val *ToolDataprivacyDeleteCategoryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableToolDataprivacyDeleteCategoryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableToolDataprivacyDeleteCategoryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToolDataprivacyDeleteCategoryRequest(val *ToolDataprivacyDeleteCategoryRequest) *NullableToolDataprivacyDeleteCategoryRequest {
	return &NullableToolDataprivacyDeleteCategoryRequest{value: val, isSet: true}
}

func (v NullableToolDataprivacyDeleteCategoryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToolDataprivacyDeleteCategoryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


