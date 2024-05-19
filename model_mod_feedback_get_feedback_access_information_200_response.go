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

// checks if the ModFeedbackGetFeedbackAccessInformation200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModFeedbackGetFeedbackAccessInformation200Response{}

// ModFeedbackGetFeedbackAccessInformation200Response struct for ModFeedbackGetFeedbackAccessInformation200Response
type ModFeedbackGetFeedbackAccessInformation200Response struct {
	// Whether the user can complete the feedback or not.
	Cancomplete bool `json:"cancomplete"`
	// Whether the user can delete submissions or not.
	Candeletesubmissions bool `json:"candeletesubmissions"`
	// Whether the user can edit feedback items or not.
	Canedititems bool `json:"canedititems"`
	// Whether the user can submit the feedback or not.
	Cansubmit bool `json:"cansubmit"`
	// Whether the user can view the analysis or not.
	Canviewanalysis bool `json:"canviewanalysis"`
	// Whether the user can view the feedback reports or not.
	Canviewreports bool `json:"canviewreports"`
	// Whether the feedback is already submitted or not.
	Isalreadysubmitted bool `json:"isalreadysubmitted"`
	// Whether the feedback is anonymous or not.
	Isanonymous bool `json:"isanonymous"`
	// Whether the feedback has questions or not.
	Isempty bool `json:"isempty"`
	// Whether the feedback has active access time restrictions or not.
	Isopen bool `json:"isopen"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _ModFeedbackGetFeedbackAccessInformation200Response ModFeedbackGetFeedbackAccessInformation200Response

// NewModFeedbackGetFeedbackAccessInformation200Response instantiates a new ModFeedbackGetFeedbackAccessInformation200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModFeedbackGetFeedbackAccessInformation200Response(cancomplete bool, candeletesubmissions bool, canedititems bool, cansubmit bool, canviewanalysis bool, canviewreports bool, isalreadysubmitted bool, isanonymous bool, isempty bool, isopen bool) *ModFeedbackGetFeedbackAccessInformation200Response {
	this := ModFeedbackGetFeedbackAccessInformation200Response{}
	this.Cancomplete = cancomplete
	this.Candeletesubmissions = candeletesubmissions
	this.Canedititems = canedititems
	this.Cansubmit = cansubmit
	this.Canviewanalysis = canviewanalysis
	this.Canviewreports = canviewreports
	this.Isalreadysubmitted = isalreadysubmitted
	this.Isanonymous = isanonymous
	this.Isempty = isempty
	this.Isopen = isopen
	return &this
}

// NewModFeedbackGetFeedbackAccessInformation200ResponseWithDefaults instantiates a new ModFeedbackGetFeedbackAccessInformation200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModFeedbackGetFeedbackAccessInformation200ResponseWithDefaults() *ModFeedbackGetFeedbackAccessInformation200Response {
	this := ModFeedbackGetFeedbackAccessInformation200Response{}
	var cancomplete bool = null
	this.Cancomplete = cancomplete
	var candeletesubmissions bool = null
	this.Candeletesubmissions = candeletesubmissions
	var canedititems bool = null
	this.Canedititems = canedititems
	var cansubmit bool = null
	this.Cansubmit = cansubmit
	var canviewanalysis bool = null
	this.Canviewanalysis = canviewanalysis
	var canviewreports bool = null
	this.Canviewreports = canviewreports
	var isalreadysubmitted bool = null
	this.Isalreadysubmitted = isalreadysubmitted
	var isanonymous bool = null
	this.Isanonymous = isanonymous
	var isempty bool = null
	this.Isempty = isempty
	var isopen bool = null
	this.Isopen = isopen
	return &this
}

// GetCancomplete returns the Cancomplete field value
func (o *ModFeedbackGetFeedbackAccessInformation200Response) GetCancomplete() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Cancomplete
}

// GetCancompleteOk returns a tuple with the Cancomplete field value
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetFeedbackAccessInformation200Response) GetCancompleteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cancomplete, true
}

// SetCancomplete sets field value
func (o *ModFeedbackGetFeedbackAccessInformation200Response) SetCancomplete(v bool) {
	o.Cancomplete = v
}

// GetCandeletesubmissions returns the Candeletesubmissions field value
func (o *ModFeedbackGetFeedbackAccessInformation200Response) GetCandeletesubmissions() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Candeletesubmissions
}

// GetCandeletesubmissionsOk returns a tuple with the Candeletesubmissions field value
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetFeedbackAccessInformation200Response) GetCandeletesubmissionsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Candeletesubmissions, true
}

// SetCandeletesubmissions sets field value
func (o *ModFeedbackGetFeedbackAccessInformation200Response) SetCandeletesubmissions(v bool) {
	o.Candeletesubmissions = v
}

// GetCanedititems returns the Canedititems field value
func (o *ModFeedbackGetFeedbackAccessInformation200Response) GetCanedititems() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Canedititems
}

// GetCanedititemsOk returns a tuple with the Canedititems field value
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetFeedbackAccessInformation200Response) GetCanedititemsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Canedititems, true
}

// SetCanedititems sets field value
func (o *ModFeedbackGetFeedbackAccessInformation200Response) SetCanedititems(v bool) {
	o.Canedititems = v
}

// GetCansubmit returns the Cansubmit field value
func (o *ModFeedbackGetFeedbackAccessInformation200Response) GetCansubmit() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Cansubmit
}

// GetCansubmitOk returns a tuple with the Cansubmit field value
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetFeedbackAccessInformation200Response) GetCansubmitOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cansubmit, true
}

// SetCansubmit sets field value
func (o *ModFeedbackGetFeedbackAccessInformation200Response) SetCansubmit(v bool) {
	o.Cansubmit = v
}

// GetCanviewanalysis returns the Canviewanalysis field value
func (o *ModFeedbackGetFeedbackAccessInformation200Response) GetCanviewanalysis() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Canviewanalysis
}

// GetCanviewanalysisOk returns a tuple with the Canviewanalysis field value
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetFeedbackAccessInformation200Response) GetCanviewanalysisOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Canviewanalysis, true
}

// SetCanviewanalysis sets field value
func (o *ModFeedbackGetFeedbackAccessInformation200Response) SetCanviewanalysis(v bool) {
	o.Canviewanalysis = v
}

// GetCanviewreports returns the Canviewreports field value
func (o *ModFeedbackGetFeedbackAccessInformation200Response) GetCanviewreports() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Canviewreports
}

// GetCanviewreportsOk returns a tuple with the Canviewreports field value
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetFeedbackAccessInformation200Response) GetCanviewreportsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Canviewreports, true
}

// SetCanviewreports sets field value
func (o *ModFeedbackGetFeedbackAccessInformation200Response) SetCanviewreports(v bool) {
	o.Canviewreports = v
}

// GetIsalreadysubmitted returns the Isalreadysubmitted field value
func (o *ModFeedbackGetFeedbackAccessInformation200Response) GetIsalreadysubmitted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Isalreadysubmitted
}

// GetIsalreadysubmittedOk returns a tuple with the Isalreadysubmitted field value
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetFeedbackAccessInformation200Response) GetIsalreadysubmittedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Isalreadysubmitted, true
}

// SetIsalreadysubmitted sets field value
func (o *ModFeedbackGetFeedbackAccessInformation200Response) SetIsalreadysubmitted(v bool) {
	o.Isalreadysubmitted = v
}

// GetIsanonymous returns the Isanonymous field value
func (o *ModFeedbackGetFeedbackAccessInformation200Response) GetIsanonymous() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Isanonymous
}

// GetIsanonymousOk returns a tuple with the Isanonymous field value
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetFeedbackAccessInformation200Response) GetIsanonymousOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Isanonymous, true
}

// SetIsanonymous sets field value
func (o *ModFeedbackGetFeedbackAccessInformation200Response) SetIsanonymous(v bool) {
	o.Isanonymous = v
}

// GetIsempty returns the Isempty field value
func (o *ModFeedbackGetFeedbackAccessInformation200Response) GetIsempty() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Isempty
}

// GetIsemptyOk returns a tuple with the Isempty field value
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetFeedbackAccessInformation200Response) GetIsemptyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Isempty, true
}

// SetIsempty sets field value
func (o *ModFeedbackGetFeedbackAccessInformation200Response) SetIsempty(v bool) {
	o.Isempty = v
}

// GetIsopen returns the Isopen field value
func (o *ModFeedbackGetFeedbackAccessInformation200Response) GetIsopen() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Isopen
}

// GetIsopenOk returns a tuple with the Isopen field value
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetFeedbackAccessInformation200Response) GetIsopenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Isopen, true
}

// SetIsopen sets field value
func (o *ModFeedbackGetFeedbackAccessInformation200Response) SetIsopen(v bool) {
	o.Isopen = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ModFeedbackGetFeedbackAccessInformation200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFeedbackGetFeedbackAccessInformation200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ModFeedbackGetFeedbackAccessInformation200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *ModFeedbackGetFeedbackAccessInformation200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o ModFeedbackGetFeedbackAccessInformation200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModFeedbackGetFeedbackAccessInformation200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cancomplete"] = o.Cancomplete
	toSerialize["candeletesubmissions"] = o.Candeletesubmissions
	toSerialize["canedititems"] = o.Canedititems
	toSerialize["cansubmit"] = o.Cansubmit
	toSerialize["canviewanalysis"] = o.Canviewanalysis
	toSerialize["canviewreports"] = o.Canviewreports
	toSerialize["isalreadysubmitted"] = o.Isalreadysubmitted
	toSerialize["isanonymous"] = o.Isanonymous
	toSerialize["isempty"] = o.Isempty
	toSerialize["isopen"] = o.Isopen
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *ModFeedbackGetFeedbackAccessInformation200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cancomplete",
		"candeletesubmissions",
		"canedititems",
		"cansubmit",
		"canviewanalysis",
		"canviewreports",
		"isalreadysubmitted",
		"isanonymous",
		"isempty",
		"isopen",
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

	varModFeedbackGetFeedbackAccessInformation200Response := _ModFeedbackGetFeedbackAccessInformation200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModFeedbackGetFeedbackAccessInformation200Response)

	if err != nil {
		return err
	}

	*o = ModFeedbackGetFeedbackAccessInformation200Response(varModFeedbackGetFeedbackAccessInformation200Response)

	return err
}

type NullableModFeedbackGetFeedbackAccessInformation200Response struct {
	value *ModFeedbackGetFeedbackAccessInformation200Response
	isSet bool
}

func (v NullableModFeedbackGetFeedbackAccessInformation200Response) Get() *ModFeedbackGetFeedbackAccessInformation200Response {
	return v.value
}

func (v *NullableModFeedbackGetFeedbackAccessInformation200Response) Set(val *ModFeedbackGetFeedbackAccessInformation200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModFeedbackGetFeedbackAccessInformation200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModFeedbackGetFeedbackAccessInformation200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModFeedbackGetFeedbackAccessInformation200Response(val *ModFeedbackGetFeedbackAccessInformation200Response) *NullableModFeedbackGetFeedbackAccessInformation200Response {
	return &NullableModFeedbackGetFeedbackAccessInformation200Response{value: val, isSet: true}
}

func (v NullableModFeedbackGetFeedbackAccessInformation200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModFeedbackGetFeedbackAccessInformation200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

