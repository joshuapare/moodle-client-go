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

// checks if the CoreCompetencyDuplicateCompetencyFrameworkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCompetencyDuplicateCompetencyFrameworkRequest{}

// CoreCompetencyDuplicateCompetencyFrameworkRequest struct for CoreCompetencyDuplicateCompetencyFrameworkRequest
type CoreCompetencyDuplicateCompetencyFrameworkRequest struct {
	// Data base record id for the framework
	Id int32 `json:"id"`
}

type _CoreCompetencyDuplicateCompetencyFrameworkRequest CoreCompetencyDuplicateCompetencyFrameworkRequest

// NewCoreCompetencyDuplicateCompetencyFrameworkRequest instantiates a new CoreCompetencyDuplicateCompetencyFrameworkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCompetencyDuplicateCompetencyFrameworkRequest(id int32) *CoreCompetencyDuplicateCompetencyFrameworkRequest {
	this := CoreCompetencyDuplicateCompetencyFrameworkRequest{}
	this.Id = id
	return &this
}

// NewCoreCompetencyDuplicateCompetencyFrameworkRequestWithDefaults instantiates a new CoreCompetencyDuplicateCompetencyFrameworkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCompetencyDuplicateCompetencyFrameworkRequestWithDefaults() *CoreCompetencyDuplicateCompetencyFrameworkRequest {
	this := CoreCompetencyDuplicateCompetencyFrameworkRequest{}
	return &this
}

// GetId returns the Id field value
func (o *CoreCompetencyDuplicateCompetencyFrameworkRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyDuplicateCompetencyFrameworkRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CoreCompetencyDuplicateCompetencyFrameworkRequest) SetId(v int32) {
	o.Id = v
}

func (o CoreCompetencyDuplicateCompetencyFrameworkRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCompetencyDuplicateCompetencyFrameworkRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *CoreCompetencyDuplicateCompetencyFrameworkRequest) UnmarshalJSON(data []byte) (err error) {
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

	varCoreCompetencyDuplicateCompetencyFrameworkRequest := _CoreCompetencyDuplicateCompetencyFrameworkRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreCompetencyDuplicateCompetencyFrameworkRequest)

	if err != nil {
		return err
	}

	*o = CoreCompetencyDuplicateCompetencyFrameworkRequest(varCoreCompetencyDuplicateCompetencyFrameworkRequest)

	return err
}

type NullableCoreCompetencyDuplicateCompetencyFrameworkRequest struct {
	value *CoreCompetencyDuplicateCompetencyFrameworkRequest
	isSet bool
}

func (v NullableCoreCompetencyDuplicateCompetencyFrameworkRequest) Get() *CoreCompetencyDuplicateCompetencyFrameworkRequest {
	return v.value
}

func (v *NullableCoreCompetencyDuplicateCompetencyFrameworkRequest) Set(val *CoreCompetencyDuplicateCompetencyFrameworkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCompetencyDuplicateCompetencyFrameworkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCompetencyDuplicateCompetencyFrameworkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCompetencyDuplicateCompetencyFrameworkRequest(val *CoreCompetencyDuplicateCompetencyFrameworkRequest) *NullableCoreCompetencyDuplicateCompetencyFrameworkRequest {
	return &NullableCoreCompetencyDuplicateCompetencyFrameworkRequest{value: val, isSet: true}
}

func (v NullableCoreCompetencyDuplicateCompetencyFrameworkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCompetencyDuplicateCompetencyFrameworkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


