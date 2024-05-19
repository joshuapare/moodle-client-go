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

// checks if the CoreXapiPostStateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreXapiPostStateRequest{}

// CoreXapiPostStateRequest struct for CoreXapiPostStateRequest
type CoreXapiPostStateRequest struct {
	// xAPI activity ID IRI
	ActivityId string `json:"activityId"`
	// The xAPI agent json
	Agent string `json:"agent"`
	// Component name
	Component string `json:"component"`
	// The xAPI registration UUID
	Registration *string `json:"registration,omitempty"`
	// JSON object with the state data
	StateData string `json:"stateData"`
	// The xAPI state ID
	StateId string `json:"stateId"`
}

type _CoreXapiPostStateRequest CoreXapiPostStateRequest

// NewCoreXapiPostStateRequest instantiates a new CoreXapiPostStateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreXapiPostStateRequest(activityId string, agent string, component string, stateData string, stateId string) *CoreXapiPostStateRequest {
	this := CoreXapiPostStateRequest{}
	this.ActivityId = activityId
	this.Agent = agent
	this.Component = component
	this.StateData = stateData
	this.StateId = stateId
	return &this
}

// NewCoreXapiPostStateRequestWithDefaults instantiates a new CoreXapiPostStateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreXapiPostStateRequestWithDefaults() *CoreXapiPostStateRequest {
	this := CoreXapiPostStateRequest{}
	var stateData string = "null"
	this.StateData = stateData
	return &this
}

// GetActivityId returns the ActivityId field value
func (o *CoreXapiPostStateRequest) GetActivityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActivityId
}

// GetActivityIdOk returns a tuple with the ActivityId field value
// and a boolean to check if the value has been set.
func (o *CoreXapiPostStateRequest) GetActivityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActivityId, true
}

// SetActivityId sets field value
func (o *CoreXapiPostStateRequest) SetActivityId(v string) {
	o.ActivityId = v
}

// GetAgent returns the Agent field value
func (o *CoreXapiPostStateRequest) GetAgent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Agent
}

// GetAgentOk returns a tuple with the Agent field value
// and a boolean to check if the value has been set.
func (o *CoreXapiPostStateRequest) GetAgentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Agent, true
}

// SetAgent sets field value
func (o *CoreXapiPostStateRequest) SetAgent(v string) {
	o.Agent = v
}

// GetComponent returns the Component field value
func (o *CoreXapiPostStateRequest) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *CoreXapiPostStateRequest) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *CoreXapiPostStateRequest) SetComponent(v string) {
	o.Component = v
}

// GetRegistration returns the Registration field value if set, zero value otherwise.
func (o *CoreXapiPostStateRequest) GetRegistration() string {
	if o == nil || IsNil(o.Registration) {
		var ret string
		return ret
	}
	return *o.Registration
}

// GetRegistrationOk returns a tuple with the Registration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreXapiPostStateRequest) GetRegistrationOk() (*string, bool) {
	if o == nil || IsNil(o.Registration) {
		return nil, false
	}
	return o.Registration, true
}

// HasRegistration returns a boolean if a field has been set.
func (o *CoreXapiPostStateRequest) HasRegistration() bool {
	if o != nil && !IsNil(o.Registration) {
		return true
	}

	return false
}

// SetRegistration gets a reference to the given string and assigns it to the Registration field.
func (o *CoreXapiPostStateRequest) SetRegistration(v string) {
	o.Registration = &v
}

// GetStateData returns the StateData field value
func (o *CoreXapiPostStateRequest) GetStateData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StateData
}

// GetStateDataOk returns a tuple with the StateData field value
// and a boolean to check if the value has been set.
func (o *CoreXapiPostStateRequest) GetStateDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StateData, true
}

// SetStateData sets field value
func (o *CoreXapiPostStateRequest) SetStateData(v string) {
	o.StateData = v
}

// GetStateId returns the StateId field value
func (o *CoreXapiPostStateRequest) GetStateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StateId
}

// GetStateIdOk returns a tuple with the StateId field value
// and a boolean to check if the value has been set.
func (o *CoreXapiPostStateRequest) GetStateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StateId, true
}

// SetStateId sets field value
func (o *CoreXapiPostStateRequest) SetStateId(v string) {
	o.StateId = v
}

func (o CoreXapiPostStateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreXapiPostStateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["activityId"] = o.ActivityId
	toSerialize["agent"] = o.Agent
	toSerialize["component"] = o.Component
	if !IsNil(o.Registration) {
		toSerialize["registration"] = o.Registration
	}
	toSerialize["stateData"] = o.StateData
	toSerialize["stateId"] = o.StateId
	return toSerialize, nil
}

func (o *CoreXapiPostStateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"activityId",
		"agent",
		"component",
		"stateData",
		"stateId",
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

	varCoreXapiPostStateRequest := _CoreXapiPostStateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreXapiPostStateRequest)

	if err != nil {
		return err
	}

	*o = CoreXapiPostStateRequest(varCoreXapiPostStateRequest)

	return err
}

type NullableCoreXapiPostStateRequest struct {
	value *CoreXapiPostStateRequest
	isSet bool
}

func (v NullableCoreXapiPostStateRequest) Get() *CoreXapiPostStateRequest {
	return v.value
}

func (v *NullableCoreXapiPostStateRequest) Set(val *CoreXapiPostStateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreXapiPostStateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreXapiPostStateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreXapiPostStateRequest(val *CoreXapiPostStateRequest) *NullableCoreXapiPostStateRequest {
	return &NullableCoreXapiPostStateRequest{value: val, isSet: true}
}

func (v NullableCoreXapiPostStateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreXapiPostStateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


