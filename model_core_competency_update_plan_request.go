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

// checks if the CoreCompetencyUpdatePlanRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCompetencyUpdatePlanRequest{}

// CoreCompetencyUpdatePlanRequest struct for CoreCompetencyUpdatePlanRequest
type CoreCompetencyUpdatePlanRequest struct {
	Plan CoreCompetencyUpdatePlanRequestPlan `json:"plan"`
}

type _CoreCompetencyUpdatePlanRequest CoreCompetencyUpdatePlanRequest

// NewCoreCompetencyUpdatePlanRequest instantiates a new CoreCompetencyUpdatePlanRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCompetencyUpdatePlanRequest(plan CoreCompetencyUpdatePlanRequestPlan) *CoreCompetencyUpdatePlanRequest {
	this := CoreCompetencyUpdatePlanRequest{}
	this.Plan = plan
	return &this
}

// NewCoreCompetencyUpdatePlanRequestWithDefaults instantiates a new CoreCompetencyUpdatePlanRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCompetencyUpdatePlanRequestWithDefaults() *CoreCompetencyUpdatePlanRequest {
	this := CoreCompetencyUpdatePlanRequest{}
	return &this
}

// GetPlan returns the Plan field value
func (o *CoreCompetencyUpdatePlanRequest) GetPlan() CoreCompetencyUpdatePlanRequestPlan {
	if o == nil {
		var ret CoreCompetencyUpdatePlanRequestPlan
		return ret
	}

	return o.Plan
}

// GetPlanOk returns a tuple with the Plan field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyUpdatePlanRequest) GetPlanOk() (*CoreCompetencyUpdatePlanRequestPlan, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Plan, true
}

// SetPlan sets field value
func (o *CoreCompetencyUpdatePlanRequest) SetPlan(v CoreCompetencyUpdatePlanRequestPlan) {
	o.Plan = v
}

func (o CoreCompetencyUpdatePlanRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCompetencyUpdatePlanRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["plan"] = o.Plan
	return toSerialize, nil
}

func (o *CoreCompetencyUpdatePlanRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"plan",
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

	varCoreCompetencyUpdatePlanRequest := _CoreCompetencyUpdatePlanRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreCompetencyUpdatePlanRequest)

	if err != nil {
		return err
	}

	*o = CoreCompetencyUpdatePlanRequest(varCoreCompetencyUpdatePlanRequest)

	return err
}

type NullableCoreCompetencyUpdatePlanRequest struct {
	value *CoreCompetencyUpdatePlanRequest
	isSet bool
}

func (v NullableCoreCompetencyUpdatePlanRequest) Get() *CoreCompetencyUpdatePlanRequest {
	return v.value
}

func (v *NullableCoreCompetencyUpdatePlanRequest) Set(val *CoreCompetencyUpdatePlanRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCompetencyUpdatePlanRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCompetencyUpdatePlanRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCompetencyUpdatePlanRequest(val *CoreCompetencyUpdatePlanRequest) *NullableCoreCompetencyUpdatePlanRequest {
	return &NullableCoreCompetencyUpdatePlanRequest{value: val, isSet: true}
}

func (v NullableCoreCompetencyUpdatePlanRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCompetencyUpdatePlanRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


