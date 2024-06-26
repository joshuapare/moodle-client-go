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

// checks if the CoreGradesGetGradeitems200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreGradesGetGradeitems200Response{}

// CoreGradesGetGradeitems200Response struct for CoreGradesGetGradeitems200Response
type CoreGradesGetGradeitems200Response struct {
	GradeItems []CoreGradesGetGradeitems200ResponseGradeItemsInner `json:"gradeItems"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _CoreGradesGetGradeitems200Response CoreGradesGetGradeitems200Response

// NewCoreGradesGetGradeitems200Response instantiates a new CoreGradesGetGradeitems200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreGradesGetGradeitems200Response(gradeItems []CoreGradesGetGradeitems200ResponseGradeItemsInner) *CoreGradesGetGradeitems200Response {
	this := CoreGradesGetGradeitems200Response{}
	this.GradeItems = gradeItems
	return &this
}

// NewCoreGradesGetGradeitems200ResponseWithDefaults instantiates a new CoreGradesGetGradeitems200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreGradesGetGradeitems200ResponseWithDefaults() *CoreGradesGetGradeitems200Response {
	this := CoreGradesGetGradeitems200Response{}
	return &this
}

// GetGradeItems returns the GradeItems field value
func (o *CoreGradesGetGradeitems200Response) GetGradeItems() []CoreGradesGetGradeitems200ResponseGradeItemsInner {
	if o == nil {
		var ret []CoreGradesGetGradeitems200ResponseGradeItemsInner
		return ret
	}

	return o.GradeItems
}

// GetGradeItemsOk returns a tuple with the GradeItems field value
// and a boolean to check if the value has been set.
func (o *CoreGradesGetGradeitems200Response) GetGradeItemsOk() ([]CoreGradesGetGradeitems200ResponseGradeItemsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.GradeItems, true
}

// SetGradeItems sets field value
func (o *CoreGradesGetGradeitems200Response) SetGradeItems(v []CoreGradesGetGradeitems200ResponseGradeItemsInner) {
	o.GradeItems = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *CoreGradesGetGradeitems200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreGradesGetGradeitems200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *CoreGradesGetGradeitems200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *CoreGradesGetGradeitems200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o CoreGradesGetGradeitems200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreGradesGetGradeitems200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["gradeItems"] = o.GradeItems
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *CoreGradesGetGradeitems200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"gradeItems",
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

	varCoreGradesGetGradeitems200Response := _CoreGradesGetGradeitems200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreGradesGetGradeitems200Response)

	if err != nil {
		return err
	}

	*o = CoreGradesGetGradeitems200Response(varCoreGradesGetGradeitems200Response)

	return err
}

type NullableCoreGradesGetGradeitems200Response struct {
	value *CoreGradesGetGradeitems200Response
	isSet bool
}

func (v NullableCoreGradesGetGradeitems200Response) Get() *CoreGradesGetGradeitems200Response {
	return v.value
}

func (v *NullableCoreGradesGetGradeitems200Response) Set(val *CoreGradesGetGradeitems200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreGradesGetGradeitems200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreGradesGetGradeitems200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreGradesGetGradeitems200Response(val *CoreGradesGetGradeitems200Response) *NullableCoreGradesGetGradeitems200Response {
	return &NullableCoreGradesGetGradeitems200Response{value: val, isSet: true}
}

func (v NullableCoreGradesGetGradeitems200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreGradesGetGradeitems200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


