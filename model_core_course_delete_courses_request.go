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

// checks if the CoreCourseDeleteCoursesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCourseDeleteCoursesRequest{}

// CoreCourseDeleteCoursesRequest struct for CoreCourseDeleteCoursesRequest
type CoreCourseDeleteCoursesRequest struct {
	Courseids []map[string]interface{} `json:"courseids"`
}

type _CoreCourseDeleteCoursesRequest CoreCourseDeleteCoursesRequest

// NewCoreCourseDeleteCoursesRequest instantiates a new CoreCourseDeleteCoursesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCourseDeleteCoursesRequest(courseids []map[string]interface{}) *CoreCourseDeleteCoursesRequest {
	this := CoreCourseDeleteCoursesRequest{}
	this.Courseids = courseids
	return &this
}

// NewCoreCourseDeleteCoursesRequestWithDefaults instantiates a new CoreCourseDeleteCoursesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCourseDeleteCoursesRequestWithDefaults() *CoreCourseDeleteCoursesRequest {
	this := CoreCourseDeleteCoursesRequest{}
	return &this
}

// GetCourseids returns the Courseids field value
func (o *CoreCourseDeleteCoursesRequest) GetCourseids() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Courseids
}

// GetCourseidsOk returns a tuple with the Courseids field value
// and a boolean to check if the value has been set.
func (o *CoreCourseDeleteCoursesRequest) GetCourseidsOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Courseids, true
}

// SetCourseids sets field value
func (o *CoreCourseDeleteCoursesRequest) SetCourseids(v []map[string]interface{}) {
	o.Courseids = v
}

func (o CoreCourseDeleteCoursesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCourseDeleteCoursesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["courseids"] = o.Courseids
	return toSerialize, nil
}

func (o *CoreCourseDeleteCoursesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"courseids",
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

	varCoreCourseDeleteCoursesRequest := _CoreCourseDeleteCoursesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreCourseDeleteCoursesRequest)

	if err != nil {
		return err
	}

	*o = CoreCourseDeleteCoursesRequest(varCoreCourseDeleteCoursesRequest)

	return err
}

type NullableCoreCourseDeleteCoursesRequest struct {
	value *CoreCourseDeleteCoursesRequest
	isSet bool
}

func (v NullableCoreCourseDeleteCoursesRequest) Get() *CoreCourseDeleteCoursesRequest {
	return v.value
}

func (v *NullableCoreCourseDeleteCoursesRequest) Set(val *CoreCourseDeleteCoursesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCourseDeleteCoursesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCourseDeleteCoursesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCourseDeleteCoursesRequest(val *CoreCourseDeleteCoursesRequest) *NullableCoreCourseDeleteCoursesRequest {
	return &NullableCoreCourseDeleteCoursesRequest{value: val, isSet: true}
}

func (v NullableCoreCourseDeleteCoursesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCourseDeleteCoursesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


