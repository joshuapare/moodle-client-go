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

// checks if the ToolDataprivacyTreeExtraBranches200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ToolDataprivacyTreeExtraBranches200Response{}

// ToolDataprivacyTreeExtraBranches200Response struct for ToolDataprivacyTreeExtraBranches200Response
type ToolDataprivacyTreeExtraBranches200Response struct {
	Branches []ToolDataprivacyTreeExtraBranches200ResponseBranchesInner `json:"branches"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _ToolDataprivacyTreeExtraBranches200Response ToolDataprivacyTreeExtraBranches200Response

// NewToolDataprivacyTreeExtraBranches200Response instantiates a new ToolDataprivacyTreeExtraBranches200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToolDataprivacyTreeExtraBranches200Response(branches []ToolDataprivacyTreeExtraBranches200ResponseBranchesInner) *ToolDataprivacyTreeExtraBranches200Response {
	this := ToolDataprivacyTreeExtraBranches200Response{}
	this.Branches = branches
	return &this
}

// NewToolDataprivacyTreeExtraBranches200ResponseWithDefaults instantiates a new ToolDataprivacyTreeExtraBranches200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToolDataprivacyTreeExtraBranches200ResponseWithDefaults() *ToolDataprivacyTreeExtraBranches200Response {
	this := ToolDataprivacyTreeExtraBranches200Response{}
	return &this
}

// GetBranches returns the Branches field value
func (o *ToolDataprivacyTreeExtraBranches200Response) GetBranches() []ToolDataprivacyTreeExtraBranches200ResponseBranchesInner {
	if o == nil {
		var ret []ToolDataprivacyTreeExtraBranches200ResponseBranchesInner
		return ret
	}

	return o.Branches
}

// GetBranchesOk returns a tuple with the Branches field value
// and a boolean to check if the value has been set.
func (o *ToolDataprivacyTreeExtraBranches200Response) GetBranchesOk() ([]ToolDataprivacyTreeExtraBranches200ResponseBranchesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Branches, true
}

// SetBranches sets field value
func (o *ToolDataprivacyTreeExtraBranches200Response) SetBranches(v []ToolDataprivacyTreeExtraBranches200ResponseBranchesInner) {
	o.Branches = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ToolDataprivacyTreeExtraBranches200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolDataprivacyTreeExtraBranches200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ToolDataprivacyTreeExtraBranches200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *ToolDataprivacyTreeExtraBranches200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o ToolDataprivacyTreeExtraBranches200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToolDataprivacyTreeExtraBranches200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["branches"] = o.Branches
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *ToolDataprivacyTreeExtraBranches200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"branches",
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

	varToolDataprivacyTreeExtraBranches200Response := _ToolDataprivacyTreeExtraBranches200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varToolDataprivacyTreeExtraBranches200Response)

	if err != nil {
		return err
	}

	*o = ToolDataprivacyTreeExtraBranches200Response(varToolDataprivacyTreeExtraBranches200Response)

	return err
}

type NullableToolDataprivacyTreeExtraBranches200Response struct {
	value *ToolDataprivacyTreeExtraBranches200Response
	isSet bool
}

func (v NullableToolDataprivacyTreeExtraBranches200Response) Get() *ToolDataprivacyTreeExtraBranches200Response {
	return v.value
}

func (v *NullableToolDataprivacyTreeExtraBranches200Response) Set(val *ToolDataprivacyTreeExtraBranches200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableToolDataprivacyTreeExtraBranches200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableToolDataprivacyTreeExtraBranches200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToolDataprivacyTreeExtraBranches200Response(val *ToolDataprivacyTreeExtraBranches200Response) *NullableToolDataprivacyTreeExtraBranches200Response {
	return &NullableToolDataprivacyTreeExtraBranches200Response{value: val, isSet: true}
}

func (v NullableToolDataprivacyTreeExtraBranches200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToolDataprivacyTreeExtraBranches200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


