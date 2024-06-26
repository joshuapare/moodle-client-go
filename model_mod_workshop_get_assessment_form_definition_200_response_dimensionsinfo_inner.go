/*
Moodle Webservice API

Auto-generated OpenAPI spec for Moodle's Webservice API. Requires installation of the Catalyst RESTful API plugin: https://github.com/catalyst/moodle-webservice_restful.

API version: 4.3.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package moodleclient

import (
	"encoding/json"
)

// checks if the ModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner{}

// ModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner struct for ModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner
type ModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner struct {
	// Dimension id.
	Id *int32 `json:"id,omitempty"`
	// Maximum grade for the dimension.
	Max *float32 `json:"max,omitempty"`
	// Minimum grade for the dimension.
	Min *float32 `json:"min,omitempty"`
	// Scale items (if used).
	Scale *string `json:"scale,omitempty"`
	// The weight of the dimension.
	Weight *string `json:"weight,omitempty"`
}

// NewModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner instantiates a new ModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner() *ModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner {
	this := ModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner{}
	var id int32 = null
	this.Id = &id
	var max float32 = null
	this.Max = &max
	var min float32 = null
	this.Min = &min
	var scale string = "null"
	this.Scale = &scale
	var weight string = "null"
	this.Weight = &weight
	return &this
}

// NewModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInnerWithDefaults instantiates a new ModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInnerWithDefaults() *ModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner {
	this := ModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner{}
	var id int32 = null
	this.Id = &id
	var max float32 = null
	this.Max = &max
	var min float32 = null
	this.Min = &min
	var scale string = "null"
	this.Scale = &scale
	var weight string = "null"
	this.Weight = &weight
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner) SetId(v int32) {
	o.Id = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *ModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner) GetMax() float32 {
	if o == nil || IsNil(o.Max) {
		var ret float32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner) GetMaxOk() (*float32, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *ModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given float32 and assigns it to the Max field.
func (o *ModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner) SetMax(v float32) {
	o.Max = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *ModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner) GetMin() float32 {
	if o == nil || IsNil(o.Min) {
		var ret float32
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner) GetMinOk() (*float32, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *ModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given float32 and assigns it to the Min field.
func (o *ModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner) SetMin(v float32) {
	o.Min = &v
}

// GetScale returns the Scale field value if set, zero value otherwise.
func (o *ModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner) GetScale() string {
	if o == nil || IsNil(o.Scale) {
		var ret string
		return ret
	}
	return *o.Scale
}

// GetScaleOk returns a tuple with the Scale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner) GetScaleOk() (*string, bool) {
	if o == nil || IsNil(o.Scale) {
		return nil, false
	}
	return o.Scale, true
}

// HasScale returns a boolean if a field has been set.
func (o *ModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner) HasScale() bool {
	if o != nil && !IsNil(o.Scale) {
		return true
	}

	return false
}

// SetScale gets a reference to the given string and assigns it to the Scale field.
func (o *ModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner) SetScale(v string) {
	o.Scale = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *ModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner) GetWeight() string {
	if o == nil || IsNil(o.Weight) {
		var ret string
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner) GetWeightOk() (*string, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *ModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given string and assigns it to the Weight field.
func (o *ModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner) SetWeight(v string) {
	o.Weight = &v
}

func (o ModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	if !IsNil(o.Scale) {
		toSerialize["scale"] = o.Scale
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	return toSerialize, nil
}

type NullableModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner struct {
	value *ModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner
	isSet bool
}

func (v NullableModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner) Get() *ModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner {
	return v.value
}

func (v *NullableModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner) Set(val *ModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner) {
	v.value = val
	v.isSet = true
}

func (v NullableModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner) IsSet() bool {
	return v.isSet
}

func (v *NullableModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner(val *ModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner) *NullableModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner {
	return &NullableModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner{value: val, isSet: true}
}

func (v NullableModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModWorkshopGetAssessmentFormDefinition200ResponseDimensionsinfoInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


