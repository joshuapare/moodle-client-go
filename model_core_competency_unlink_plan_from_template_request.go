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

// checks if the CoreCompetencyUnlinkPlanFromTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCompetencyUnlinkPlanFromTemplateRequest{}

// CoreCompetencyUnlinkPlanFromTemplateRequest struct for CoreCompetencyUnlinkPlanFromTemplateRequest
type CoreCompetencyUnlinkPlanFromTemplateRequest struct {
	// Data base record id for the plan
	Planid int32 `json:"planid"`
}

type _CoreCompetencyUnlinkPlanFromTemplateRequest CoreCompetencyUnlinkPlanFromTemplateRequest

// NewCoreCompetencyUnlinkPlanFromTemplateRequest instantiates a new CoreCompetencyUnlinkPlanFromTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCompetencyUnlinkPlanFromTemplateRequest(planid int32) *CoreCompetencyUnlinkPlanFromTemplateRequest {
	this := CoreCompetencyUnlinkPlanFromTemplateRequest{}
	this.Planid = planid
	return &this
}

// NewCoreCompetencyUnlinkPlanFromTemplateRequestWithDefaults instantiates a new CoreCompetencyUnlinkPlanFromTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCompetencyUnlinkPlanFromTemplateRequestWithDefaults() *CoreCompetencyUnlinkPlanFromTemplateRequest {
	this := CoreCompetencyUnlinkPlanFromTemplateRequest{}
	return &this
}

// GetPlanid returns the Planid field value
func (o *CoreCompetencyUnlinkPlanFromTemplateRequest) GetPlanid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Planid
}

// GetPlanidOk returns a tuple with the Planid field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyUnlinkPlanFromTemplateRequest) GetPlanidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Planid, true
}

// SetPlanid sets field value
func (o *CoreCompetencyUnlinkPlanFromTemplateRequest) SetPlanid(v int32) {
	o.Planid = v
}

func (o CoreCompetencyUnlinkPlanFromTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCompetencyUnlinkPlanFromTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["planid"] = o.Planid
	return toSerialize, nil
}

func (o *CoreCompetencyUnlinkPlanFromTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"planid",
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

	varCoreCompetencyUnlinkPlanFromTemplateRequest := _CoreCompetencyUnlinkPlanFromTemplateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreCompetencyUnlinkPlanFromTemplateRequest)

	if err != nil {
		return err
	}

	*o = CoreCompetencyUnlinkPlanFromTemplateRequest(varCoreCompetencyUnlinkPlanFromTemplateRequest)

	return err
}

type NullableCoreCompetencyUnlinkPlanFromTemplateRequest struct {
	value *CoreCompetencyUnlinkPlanFromTemplateRequest
	isSet bool
}

func (v NullableCoreCompetencyUnlinkPlanFromTemplateRequest) Get() *CoreCompetencyUnlinkPlanFromTemplateRequest {
	return v.value
}

func (v *NullableCoreCompetencyUnlinkPlanFromTemplateRequest) Set(val *CoreCompetencyUnlinkPlanFromTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCompetencyUnlinkPlanFromTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCompetencyUnlinkPlanFromTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCompetencyUnlinkPlanFromTemplateRequest(val *CoreCompetencyUnlinkPlanFromTemplateRequest) *NullableCoreCompetencyUnlinkPlanFromTemplateRequest {
	return &NullableCoreCompetencyUnlinkPlanFromTemplateRequest{value: val, isSet: true}
}

func (v NullableCoreCompetencyUnlinkPlanFromTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCompetencyUnlinkPlanFromTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


