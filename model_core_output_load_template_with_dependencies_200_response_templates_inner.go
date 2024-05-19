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

// checks if the CoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner{}

// CoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner struct for CoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner
type CoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner struct {
	// component containing the resource
	Component *string `json:"component,omitempty"`
	// name of the resource
	Name *string `json:"name,omitempty"`
	// resource value
	Value *string `json:"value,omitempty"`
}

// NewCoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner instantiates a new CoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner() *CoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner {
	this := CoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner{}
	return &this
}

// NewCoreOutputLoadTemplateWithDependencies200ResponseTemplatesInnerWithDefaults instantiates a new CoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreOutputLoadTemplateWithDependencies200ResponseTemplatesInnerWithDefaults() *CoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner {
	this := CoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner{}
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *CoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner) GetComponent() string {
	if o == nil || IsNil(o.Component) {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner) GetComponentOk() (*string, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *CoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *CoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner) SetComponent(v string) {
	o.Component = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner) SetValue(v string) {
	o.Value = &v
}

func (o CoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Component) {
		toSerialize["component"] = o.Component
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableCoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner struct {
	value *CoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner
	isSet bool
}

func (v NullableCoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner) Get() *CoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner {
	return v.value
}

func (v *NullableCoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner) Set(val *CoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner(val *CoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner) *NullableCoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner {
	return &NullableCoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner{value: val, isSet: true}
}

func (v NullableCoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

