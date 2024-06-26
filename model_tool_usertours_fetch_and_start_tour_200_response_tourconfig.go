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

// checks if the ToolUsertoursFetchAndStartTour200ResponseTourconfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ToolUsertoursFetchAndStartTour200ResponseTourconfig{}

// ToolUsertoursFetchAndStartTour200ResponseTourconfig struct for ToolUsertoursFetchAndStartTour200ResponseTourconfig
type ToolUsertoursFetchAndStartTour200ResponseTourconfig struct {
	// display step number
	Displaystepnumbers bool `json:"displaystepnumbers"`
	// Label of the end tour button
	Endtourlabel string `json:"endtourlabel"`
	// Tour Name
	Name string `json:"name"`
	Steps []ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner `json:"steps"`
}

type _ToolUsertoursFetchAndStartTour200ResponseTourconfig ToolUsertoursFetchAndStartTour200ResponseTourconfig

// NewToolUsertoursFetchAndStartTour200ResponseTourconfig instantiates a new ToolUsertoursFetchAndStartTour200ResponseTourconfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToolUsertoursFetchAndStartTour200ResponseTourconfig(displaystepnumbers bool, endtourlabel string, name string, steps []ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) *ToolUsertoursFetchAndStartTour200ResponseTourconfig {
	this := ToolUsertoursFetchAndStartTour200ResponseTourconfig{}
	this.Displaystepnumbers = displaystepnumbers
	this.Endtourlabel = endtourlabel
	this.Name = name
	this.Steps = steps
	return &this
}

// NewToolUsertoursFetchAndStartTour200ResponseTourconfigWithDefaults instantiates a new ToolUsertoursFetchAndStartTour200ResponseTourconfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToolUsertoursFetchAndStartTour200ResponseTourconfigWithDefaults() *ToolUsertoursFetchAndStartTour200ResponseTourconfig {
	this := ToolUsertoursFetchAndStartTour200ResponseTourconfig{}
	var displaystepnumbers bool = null
	this.Displaystepnumbers = displaystepnumbers
	var endtourlabel string = "null"
	this.Endtourlabel = endtourlabel
	var name string = "null"
	this.Name = name
	return &this
}

// GetDisplaystepnumbers returns the Displaystepnumbers field value
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfig) GetDisplaystepnumbers() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Displaystepnumbers
}

// GetDisplaystepnumbersOk returns a tuple with the Displaystepnumbers field value
// and a boolean to check if the value has been set.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfig) GetDisplaystepnumbersOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Displaystepnumbers, true
}

// SetDisplaystepnumbers sets field value
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfig) SetDisplaystepnumbers(v bool) {
	o.Displaystepnumbers = v
}

// GetEndtourlabel returns the Endtourlabel field value
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfig) GetEndtourlabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Endtourlabel
}

// GetEndtourlabelOk returns a tuple with the Endtourlabel field value
// and a boolean to check if the value has been set.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfig) GetEndtourlabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endtourlabel, true
}

// SetEndtourlabel sets field value
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfig) SetEndtourlabel(v string) {
	o.Endtourlabel = v
}

// GetName returns the Name field value
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfig) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfig) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfig) SetName(v string) {
	o.Name = v
}

// GetSteps returns the Steps field value
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfig) GetSteps() []ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner {
	if o == nil {
		var ret []ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner
		return ret
	}

	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value
// and a boolean to check if the value has been set.
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfig) GetStepsOk() ([]ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Steps, true
}

// SetSteps sets field value
func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfig) SetSteps(v []ToolUsertoursFetchAndStartTour200ResponseTourconfigStepsInner) {
	o.Steps = v
}

func (o ToolUsertoursFetchAndStartTour200ResponseTourconfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToolUsertoursFetchAndStartTour200ResponseTourconfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["displaystepnumbers"] = o.Displaystepnumbers
	toSerialize["endtourlabel"] = o.Endtourlabel
	toSerialize["name"] = o.Name
	toSerialize["steps"] = o.Steps
	return toSerialize, nil
}

func (o *ToolUsertoursFetchAndStartTour200ResponseTourconfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"displaystepnumbers",
		"endtourlabel",
		"name",
		"steps",
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

	varToolUsertoursFetchAndStartTour200ResponseTourconfig := _ToolUsertoursFetchAndStartTour200ResponseTourconfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varToolUsertoursFetchAndStartTour200ResponseTourconfig)

	if err != nil {
		return err
	}

	*o = ToolUsertoursFetchAndStartTour200ResponseTourconfig(varToolUsertoursFetchAndStartTour200ResponseTourconfig)

	return err
}

type NullableToolUsertoursFetchAndStartTour200ResponseTourconfig struct {
	value *ToolUsertoursFetchAndStartTour200ResponseTourconfig
	isSet bool
}

func (v NullableToolUsertoursFetchAndStartTour200ResponseTourconfig) Get() *ToolUsertoursFetchAndStartTour200ResponseTourconfig {
	return v.value
}

func (v *NullableToolUsertoursFetchAndStartTour200ResponseTourconfig) Set(val *ToolUsertoursFetchAndStartTour200ResponseTourconfig) {
	v.value = val
	v.isSet = true
}

func (v NullableToolUsertoursFetchAndStartTour200ResponseTourconfig) IsSet() bool {
	return v.isSet
}

func (v *NullableToolUsertoursFetchAndStartTour200ResponseTourconfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToolUsertoursFetchAndStartTour200ResponseTourconfig(val *ToolUsertoursFetchAndStartTour200ResponseTourconfig) *NullableToolUsertoursFetchAndStartTour200ResponseTourconfig {
	return &NullableToolUsertoursFetchAndStartTour200ResponseTourconfig{value: val, isSet: true}
}

func (v NullableToolUsertoursFetchAndStartTour200ResponseTourconfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToolUsertoursFetchAndStartTour200ResponseTourconfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


