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

// checks if the CoreCalendarUpdateEventStartDayRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCalendarUpdateEventStartDayRequest{}

// CoreCalendarUpdateEventStartDayRequest struct for CoreCalendarUpdateEventStartDayRequest
type CoreCalendarUpdateEventStartDayRequest struct {
	// Timestamp for the new start day
	Daytimestamp int32 `json:"daytimestamp"`
	// Id of event to be updated
	Eventid int32 `json:"eventid"`
}

type _CoreCalendarUpdateEventStartDayRequest CoreCalendarUpdateEventStartDayRequest

// NewCoreCalendarUpdateEventStartDayRequest instantiates a new CoreCalendarUpdateEventStartDayRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCalendarUpdateEventStartDayRequest(daytimestamp int32, eventid int32) *CoreCalendarUpdateEventStartDayRequest {
	this := CoreCalendarUpdateEventStartDayRequest{}
	this.Daytimestamp = daytimestamp
	this.Eventid = eventid
	return &this
}

// NewCoreCalendarUpdateEventStartDayRequestWithDefaults instantiates a new CoreCalendarUpdateEventStartDayRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCalendarUpdateEventStartDayRequestWithDefaults() *CoreCalendarUpdateEventStartDayRequest {
	this := CoreCalendarUpdateEventStartDayRequest{}
	var daytimestamp int32 = null
	this.Daytimestamp = daytimestamp
	var eventid int32 = null
	this.Eventid = eventid
	return &this
}

// GetDaytimestamp returns the Daytimestamp field value
func (o *CoreCalendarUpdateEventStartDayRequest) GetDaytimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Daytimestamp
}

// GetDaytimestampOk returns a tuple with the Daytimestamp field value
// and a boolean to check if the value has been set.
func (o *CoreCalendarUpdateEventStartDayRequest) GetDaytimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Daytimestamp, true
}

// SetDaytimestamp sets field value
func (o *CoreCalendarUpdateEventStartDayRequest) SetDaytimestamp(v int32) {
	o.Daytimestamp = v
}

// GetEventid returns the Eventid field value
func (o *CoreCalendarUpdateEventStartDayRequest) GetEventid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Eventid
}

// GetEventidOk returns a tuple with the Eventid field value
// and a boolean to check if the value has been set.
func (o *CoreCalendarUpdateEventStartDayRequest) GetEventidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Eventid, true
}

// SetEventid sets field value
func (o *CoreCalendarUpdateEventStartDayRequest) SetEventid(v int32) {
	o.Eventid = v
}

func (o CoreCalendarUpdateEventStartDayRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCalendarUpdateEventStartDayRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["daytimestamp"] = o.Daytimestamp
	toSerialize["eventid"] = o.Eventid
	return toSerialize, nil
}

func (o *CoreCalendarUpdateEventStartDayRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"daytimestamp",
		"eventid",
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

	varCoreCalendarUpdateEventStartDayRequest := _CoreCalendarUpdateEventStartDayRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreCalendarUpdateEventStartDayRequest)

	if err != nil {
		return err
	}

	*o = CoreCalendarUpdateEventStartDayRequest(varCoreCalendarUpdateEventStartDayRequest)

	return err
}

type NullableCoreCalendarUpdateEventStartDayRequest struct {
	value *CoreCalendarUpdateEventStartDayRequest
	isSet bool
}

func (v NullableCoreCalendarUpdateEventStartDayRequest) Get() *CoreCalendarUpdateEventStartDayRequest {
	return v.value
}

func (v *NullableCoreCalendarUpdateEventStartDayRequest) Set(val *CoreCalendarUpdateEventStartDayRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCalendarUpdateEventStartDayRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCalendarUpdateEventStartDayRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCalendarUpdateEventStartDayRequest(val *CoreCalendarUpdateEventStartDayRequest) *NullableCoreCalendarUpdateEventStartDayRequest {
	return &NullableCoreCalendarUpdateEventStartDayRequest{value: val, isSet: true}
}

func (v NullableCoreCalendarUpdateEventStartDayRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCalendarUpdateEventStartDayRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


