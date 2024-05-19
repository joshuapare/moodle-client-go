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

// checks if the CoreCourseCheckUpdates200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCourseCheckUpdates200Response{}

// CoreCourseCheckUpdates200Response struct for CoreCourseCheckUpdates200Response
type CoreCourseCheckUpdates200Response struct {
	Instances []CoreCourseCheckUpdates200ResponseInstancesInner `json:"instances"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _CoreCourseCheckUpdates200Response CoreCourseCheckUpdates200Response

// NewCoreCourseCheckUpdates200Response instantiates a new CoreCourseCheckUpdates200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCourseCheckUpdates200Response(instances []CoreCourseCheckUpdates200ResponseInstancesInner) *CoreCourseCheckUpdates200Response {
	this := CoreCourseCheckUpdates200Response{}
	this.Instances = instances
	return &this
}

// NewCoreCourseCheckUpdates200ResponseWithDefaults instantiates a new CoreCourseCheckUpdates200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCourseCheckUpdates200ResponseWithDefaults() *CoreCourseCheckUpdates200Response {
	this := CoreCourseCheckUpdates200Response{}
	return &this
}

// GetInstances returns the Instances field value
func (o *CoreCourseCheckUpdates200Response) GetInstances() []CoreCourseCheckUpdates200ResponseInstancesInner {
	if o == nil {
		var ret []CoreCourseCheckUpdates200ResponseInstancesInner
		return ret
	}

	return o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value
// and a boolean to check if the value has been set.
func (o *CoreCourseCheckUpdates200Response) GetInstancesOk() ([]CoreCourseCheckUpdates200ResponseInstancesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Instances, true
}

// SetInstances sets field value
func (o *CoreCourseCheckUpdates200Response) SetInstances(v []CoreCourseCheckUpdates200ResponseInstancesInner) {
	o.Instances = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *CoreCourseCheckUpdates200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseCheckUpdates200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *CoreCourseCheckUpdates200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *CoreCourseCheckUpdates200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o CoreCourseCheckUpdates200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCourseCheckUpdates200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["instances"] = o.Instances
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *CoreCourseCheckUpdates200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"instances",
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

	varCoreCourseCheckUpdates200Response := _CoreCourseCheckUpdates200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreCourseCheckUpdates200Response)

	if err != nil {
		return err
	}

	*o = CoreCourseCheckUpdates200Response(varCoreCourseCheckUpdates200Response)

	return err
}

type NullableCoreCourseCheckUpdates200Response struct {
	value *CoreCourseCheckUpdates200Response
	isSet bool
}

func (v NullableCoreCourseCheckUpdates200Response) Get() *CoreCourseCheckUpdates200Response {
	return v.value
}

func (v *NullableCoreCourseCheckUpdates200Response) Set(val *CoreCourseCheckUpdates200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCourseCheckUpdates200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCourseCheckUpdates200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCourseCheckUpdates200Response(val *CoreCourseCheckUpdates200Response) *NullableCoreCourseCheckUpdates200Response {
	return &NullableCoreCourseCheckUpdates200Response{value: val, isSet: true}
}

func (v NullableCoreCourseCheckUpdates200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCourseCheckUpdates200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


