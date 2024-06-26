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

// checks if the CoreReportbuilderListReports200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreReportbuilderListReports200Response{}

// CoreReportbuilderListReports200Response struct for CoreReportbuilderListReports200Response
type CoreReportbuilderListReports200Response struct {
	Reports []CoreReportbuilderListReports200ResponseReportsInner `json:"reports"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _CoreReportbuilderListReports200Response CoreReportbuilderListReports200Response

// NewCoreReportbuilderListReports200Response instantiates a new CoreReportbuilderListReports200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreReportbuilderListReports200Response(reports []CoreReportbuilderListReports200ResponseReportsInner) *CoreReportbuilderListReports200Response {
	this := CoreReportbuilderListReports200Response{}
	this.Reports = reports
	return &this
}

// NewCoreReportbuilderListReports200ResponseWithDefaults instantiates a new CoreReportbuilderListReports200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreReportbuilderListReports200ResponseWithDefaults() *CoreReportbuilderListReports200Response {
	this := CoreReportbuilderListReports200Response{}
	return &this
}

// GetReports returns the Reports field value
func (o *CoreReportbuilderListReports200Response) GetReports() []CoreReportbuilderListReports200ResponseReportsInner {
	if o == nil {
		var ret []CoreReportbuilderListReports200ResponseReportsInner
		return ret
	}

	return o.Reports
}

// GetReportsOk returns a tuple with the Reports field value
// and a boolean to check if the value has been set.
func (o *CoreReportbuilderListReports200Response) GetReportsOk() ([]CoreReportbuilderListReports200ResponseReportsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reports, true
}

// SetReports sets field value
func (o *CoreReportbuilderListReports200Response) SetReports(v []CoreReportbuilderListReports200ResponseReportsInner) {
	o.Reports = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *CoreReportbuilderListReports200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreReportbuilderListReports200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *CoreReportbuilderListReports200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *CoreReportbuilderListReports200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o CoreReportbuilderListReports200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreReportbuilderListReports200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reports"] = o.Reports
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *CoreReportbuilderListReports200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"reports",
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

	varCoreReportbuilderListReports200Response := _CoreReportbuilderListReports200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreReportbuilderListReports200Response)

	if err != nil {
		return err
	}

	*o = CoreReportbuilderListReports200Response(varCoreReportbuilderListReports200Response)

	return err
}

type NullableCoreReportbuilderListReports200Response struct {
	value *CoreReportbuilderListReports200Response
	isSet bool
}

func (v NullableCoreReportbuilderListReports200Response) Get() *CoreReportbuilderListReports200Response {
	return v.value
}

func (v *NullableCoreReportbuilderListReports200Response) Set(val *CoreReportbuilderListReports200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreReportbuilderListReports200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreReportbuilderListReports200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreReportbuilderListReports200Response(val *CoreReportbuilderListReports200Response) *NullableCoreReportbuilderListReports200Response {
	return &NullableCoreReportbuilderListReports200Response{value: val, isSet: true}
}

func (v NullableCoreReportbuilderListReports200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreReportbuilderListReports200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


