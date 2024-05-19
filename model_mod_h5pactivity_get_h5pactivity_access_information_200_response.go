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

// checks if the ModH5pactivityGetH5pactivityAccessInformation200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModH5pactivityGetH5pactivityAccessInformation200Response{}

// ModH5pactivityGetH5pactivityAccessInformation200Response struct for ModH5pactivityGetH5pactivityAccessInformation200Response
type ModH5pactivityGetH5pactivityAccessInformation200Response struct {
	// Whether the user has the capability mod/h5pactivity:addinstance allowed.
	Canaddinstance *bool `json:"canaddinstance,omitempty"`
	// Whether the user has the capability mod/h5pactivity:reviewattempts allowed.
	Canreviewattempts *bool `json:"canreviewattempts,omitempty"`
	// Whether the user has the capability mod/h5pactivity:submit allowed.
	Cansubmit *bool `json:"cansubmit,omitempty"`
	// Whether the user has the capability mod/h5pactivity:view allowed.
	Canview *bool `json:"canview,omitempty"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

// NewModH5pactivityGetH5pactivityAccessInformation200Response instantiates a new ModH5pactivityGetH5pactivityAccessInformation200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModH5pactivityGetH5pactivityAccessInformation200Response() *ModH5pactivityGetH5pactivityAccessInformation200Response {
	this := ModH5pactivityGetH5pactivityAccessInformation200Response{}
	var canaddinstance bool = null
	this.Canaddinstance = &canaddinstance
	var canreviewattempts bool = null
	this.Canreviewattempts = &canreviewattempts
	var cansubmit bool = null
	this.Cansubmit = &cansubmit
	var canview bool = null
	this.Canview = &canview
	return &this
}

// NewModH5pactivityGetH5pactivityAccessInformation200ResponseWithDefaults instantiates a new ModH5pactivityGetH5pactivityAccessInformation200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModH5pactivityGetH5pactivityAccessInformation200ResponseWithDefaults() *ModH5pactivityGetH5pactivityAccessInformation200Response {
	this := ModH5pactivityGetH5pactivityAccessInformation200Response{}
	var canaddinstance bool = null
	this.Canaddinstance = &canaddinstance
	var canreviewattempts bool = null
	this.Canreviewattempts = &canreviewattempts
	var cansubmit bool = null
	this.Cansubmit = &cansubmit
	var canview bool = null
	this.Canview = &canview
	return &this
}

// GetCanaddinstance returns the Canaddinstance field value if set, zero value otherwise.
func (o *ModH5pactivityGetH5pactivityAccessInformation200Response) GetCanaddinstance() bool {
	if o == nil || IsNil(o.Canaddinstance) {
		var ret bool
		return ret
	}
	return *o.Canaddinstance
}

// GetCanaddinstanceOk returns a tuple with the Canaddinstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModH5pactivityGetH5pactivityAccessInformation200Response) GetCanaddinstanceOk() (*bool, bool) {
	if o == nil || IsNil(o.Canaddinstance) {
		return nil, false
	}
	return o.Canaddinstance, true
}

// HasCanaddinstance returns a boolean if a field has been set.
func (o *ModH5pactivityGetH5pactivityAccessInformation200Response) HasCanaddinstance() bool {
	if o != nil && !IsNil(o.Canaddinstance) {
		return true
	}

	return false
}

// SetCanaddinstance gets a reference to the given bool and assigns it to the Canaddinstance field.
func (o *ModH5pactivityGetH5pactivityAccessInformation200Response) SetCanaddinstance(v bool) {
	o.Canaddinstance = &v
}

// GetCanreviewattempts returns the Canreviewattempts field value if set, zero value otherwise.
func (o *ModH5pactivityGetH5pactivityAccessInformation200Response) GetCanreviewattempts() bool {
	if o == nil || IsNil(o.Canreviewattempts) {
		var ret bool
		return ret
	}
	return *o.Canreviewattempts
}

// GetCanreviewattemptsOk returns a tuple with the Canreviewattempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModH5pactivityGetH5pactivityAccessInformation200Response) GetCanreviewattemptsOk() (*bool, bool) {
	if o == nil || IsNil(o.Canreviewattempts) {
		return nil, false
	}
	return o.Canreviewattempts, true
}

// HasCanreviewattempts returns a boolean if a field has been set.
func (o *ModH5pactivityGetH5pactivityAccessInformation200Response) HasCanreviewattempts() bool {
	if o != nil && !IsNil(o.Canreviewattempts) {
		return true
	}

	return false
}

// SetCanreviewattempts gets a reference to the given bool and assigns it to the Canreviewattempts field.
func (o *ModH5pactivityGetH5pactivityAccessInformation200Response) SetCanreviewattempts(v bool) {
	o.Canreviewattempts = &v
}

// GetCansubmit returns the Cansubmit field value if set, zero value otherwise.
func (o *ModH5pactivityGetH5pactivityAccessInformation200Response) GetCansubmit() bool {
	if o == nil || IsNil(o.Cansubmit) {
		var ret bool
		return ret
	}
	return *o.Cansubmit
}

// GetCansubmitOk returns a tuple with the Cansubmit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModH5pactivityGetH5pactivityAccessInformation200Response) GetCansubmitOk() (*bool, bool) {
	if o == nil || IsNil(o.Cansubmit) {
		return nil, false
	}
	return o.Cansubmit, true
}

// HasCansubmit returns a boolean if a field has been set.
func (o *ModH5pactivityGetH5pactivityAccessInformation200Response) HasCansubmit() bool {
	if o != nil && !IsNil(o.Cansubmit) {
		return true
	}

	return false
}

// SetCansubmit gets a reference to the given bool and assigns it to the Cansubmit field.
func (o *ModH5pactivityGetH5pactivityAccessInformation200Response) SetCansubmit(v bool) {
	o.Cansubmit = &v
}

// GetCanview returns the Canview field value if set, zero value otherwise.
func (o *ModH5pactivityGetH5pactivityAccessInformation200Response) GetCanview() bool {
	if o == nil || IsNil(o.Canview) {
		var ret bool
		return ret
	}
	return *o.Canview
}

// GetCanviewOk returns a tuple with the Canview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModH5pactivityGetH5pactivityAccessInformation200Response) GetCanviewOk() (*bool, bool) {
	if o == nil || IsNil(o.Canview) {
		return nil, false
	}
	return o.Canview, true
}

// HasCanview returns a boolean if a field has been set.
func (o *ModH5pactivityGetH5pactivityAccessInformation200Response) HasCanview() bool {
	if o != nil && !IsNil(o.Canview) {
		return true
	}

	return false
}

// SetCanview gets a reference to the given bool and assigns it to the Canview field.
func (o *ModH5pactivityGetH5pactivityAccessInformation200Response) SetCanview(v bool) {
	o.Canview = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ModH5pactivityGetH5pactivityAccessInformation200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModH5pactivityGetH5pactivityAccessInformation200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ModH5pactivityGetH5pactivityAccessInformation200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *ModH5pactivityGetH5pactivityAccessInformation200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o ModH5pactivityGetH5pactivityAccessInformation200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModH5pactivityGetH5pactivityAccessInformation200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Canaddinstance) {
		toSerialize["canaddinstance"] = o.Canaddinstance
	}
	if !IsNil(o.Canreviewattempts) {
		toSerialize["canreviewattempts"] = o.Canreviewattempts
	}
	if !IsNil(o.Cansubmit) {
		toSerialize["cansubmit"] = o.Cansubmit
	}
	if !IsNil(o.Canview) {
		toSerialize["canview"] = o.Canview
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableModH5pactivityGetH5pactivityAccessInformation200Response struct {
	value *ModH5pactivityGetH5pactivityAccessInformation200Response
	isSet bool
}

func (v NullableModH5pactivityGetH5pactivityAccessInformation200Response) Get() *ModH5pactivityGetH5pactivityAccessInformation200Response {
	return v.value
}

func (v *NullableModH5pactivityGetH5pactivityAccessInformation200Response) Set(val *ModH5pactivityGetH5pactivityAccessInformation200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModH5pactivityGetH5pactivityAccessInformation200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModH5pactivityGetH5pactivityAccessInformation200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModH5pactivityGetH5pactivityAccessInformation200Response(val *ModH5pactivityGetH5pactivityAccessInformation200Response) *NullableModH5pactivityGetH5pactivityAccessInformation200Response {
	return &NullableModH5pactivityGetH5pactivityAccessInformation200Response{value: val, isSet: true}
}

func (v NullableModH5pactivityGetH5pactivityAccessInformation200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModH5pactivityGetH5pactivityAccessInformation200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

