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

// checks if the ModFeedbackViewFeedbackRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModFeedbackViewFeedbackRequest{}

// ModFeedbackViewFeedbackRequest struct for ModFeedbackViewFeedbackRequest
type ModFeedbackViewFeedbackRequest struct {
	// Course where user completes the feedback (for site feedbacks only).
	Courseid *int32 `json:"courseid,omitempty"`
	// Feedback instance id
	Feedbackid int32 `json:"feedbackid"`
	// If we need to mark the module as viewed for completion
	Moduleviewed *bool `json:"moduleviewed,omitempty"`
}

type _ModFeedbackViewFeedbackRequest ModFeedbackViewFeedbackRequest

// NewModFeedbackViewFeedbackRequest instantiates a new ModFeedbackViewFeedbackRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModFeedbackViewFeedbackRequest(feedbackid int32) *ModFeedbackViewFeedbackRequest {
	this := ModFeedbackViewFeedbackRequest{}
	var courseid int32 = 0
	this.Courseid = &courseid
	this.Feedbackid = feedbackid
	var moduleviewed bool = false
	this.Moduleviewed = &moduleviewed
	return &this
}

// NewModFeedbackViewFeedbackRequestWithDefaults instantiates a new ModFeedbackViewFeedbackRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModFeedbackViewFeedbackRequestWithDefaults() *ModFeedbackViewFeedbackRequest {
	this := ModFeedbackViewFeedbackRequest{}
	var courseid int32 = 0
	this.Courseid = &courseid
	var moduleviewed bool = false
	this.Moduleviewed = &moduleviewed
	return &this
}

// GetCourseid returns the Courseid field value if set, zero value otherwise.
func (o *ModFeedbackViewFeedbackRequest) GetCourseid() int32 {
	if o == nil || IsNil(o.Courseid) {
		var ret int32
		return ret
	}
	return *o.Courseid
}

// GetCourseidOk returns a tuple with the Courseid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFeedbackViewFeedbackRequest) GetCourseidOk() (*int32, bool) {
	if o == nil || IsNil(o.Courseid) {
		return nil, false
	}
	return o.Courseid, true
}

// HasCourseid returns a boolean if a field has been set.
func (o *ModFeedbackViewFeedbackRequest) HasCourseid() bool {
	if o != nil && !IsNil(o.Courseid) {
		return true
	}

	return false
}

// SetCourseid gets a reference to the given int32 and assigns it to the Courseid field.
func (o *ModFeedbackViewFeedbackRequest) SetCourseid(v int32) {
	o.Courseid = &v
}

// GetFeedbackid returns the Feedbackid field value
func (o *ModFeedbackViewFeedbackRequest) GetFeedbackid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Feedbackid
}

// GetFeedbackidOk returns a tuple with the Feedbackid field value
// and a boolean to check if the value has been set.
func (o *ModFeedbackViewFeedbackRequest) GetFeedbackidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Feedbackid, true
}

// SetFeedbackid sets field value
func (o *ModFeedbackViewFeedbackRequest) SetFeedbackid(v int32) {
	o.Feedbackid = v
}

// GetModuleviewed returns the Moduleviewed field value if set, zero value otherwise.
func (o *ModFeedbackViewFeedbackRequest) GetModuleviewed() bool {
	if o == nil || IsNil(o.Moduleviewed) {
		var ret bool
		return ret
	}
	return *o.Moduleviewed
}

// GetModuleviewedOk returns a tuple with the Moduleviewed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFeedbackViewFeedbackRequest) GetModuleviewedOk() (*bool, bool) {
	if o == nil || IsNil(o.Moduleviewed) {
		return nil, false
	}
	return o.Moduleviewed, true
}

// HasModuleviewed returns a boolean if a field has been set.
func (o *ModFeedbackViewFeedbackRequest) HasModuleviewed() bool {
	if o != nil && !IsNil(o.Moduleviewed) {
		return true
	}

	return false
}

// SetModuleviewed gets a reference to the given bool and assigns it to the Moduleviewed field.
func (o *ModFeedbackViewFeedbackRequest) SetModuleviewed(v bool) {
	o.Moduleviewed = &v
}

func (o ModFeedbackViewFeedbackRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModFeedbackViewFeedbackRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Courseid) {
		toSerialize["courseid"] = o.Courseid
	}
	toSerialize["feedbackid"] = o.Feedbackid
	if !IsNil(o.Moduleviewed) {
		toSerialize["moduleviewed"] = o.Moduleviewed
	}
	return toSerialize, nil
}

func (o *ModFeedbackViewFeedbackRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"feedbackid",
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

	varModFeedbackViewFeedbackRequest := _ModFeedbackViewFeedbackRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModFeedbackViewFeedbackRequest)

	if err != nil {
		return err
	}

	*o = ModFeedbackViewFeedbackRequest(varModFeedbackViewFeedbackRequest)

	return err
}

type NullableModFeedbackViewFeedbackRequest struct {
	value *ModFeedbackViewFeedbackRequest
	isSet bool
}

func (v NullableModFeedbackViewFeedbackRequest) Get() *ModFeedbackViewFeedbackRequest {
	return v.value
}

func (v *NullableModFeedbackViewFeedbackRequest) Set(val *ModFeedbackViewFeedbackRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModFeedbackViewFeedbackRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModFeedbackViewFeedbackRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModFeedbackViewFeedbackRequest(val *ModFeedbackViewFeedbackRequest) *NullableModFeedbackViewFeedbackRequest {
	return &NullableModFeedbackViewFeedbackRequest{value: val, isSet: true}
}

func (v NullableModFeedbackViewFeedbackRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModFeedbackViewFeedbackRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


