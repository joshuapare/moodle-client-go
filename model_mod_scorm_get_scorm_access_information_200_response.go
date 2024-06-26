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

// checks if the ModScormGetScormAccessInformation200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModScormGetScormAccessInformation200Response{}

// ModScormGetScormAccessInformation200Response struct for ModScormGetScormAccessInformation200Response
type ModScormGetScormAccessInformation200Response struct {
	// Whether the user has the capability mod/scorm:addinstance allowed.
	Canaddinstance *bool `json:"canaddinstance,omitempty"`
	// Whether the user has the capability mod/scorm:deleteownresponses allowed.
	Candeleteownresponses *bool `json:"candeleteownresponses,omitempty"`
	// Whether the user has the capability mod/scorm:deleteresponses allowed.
	Candeleteresponses *bool `json:"candeleteresponses,omitempty"`
	// Whether the user has the capability mod/scorm:savetrack allowed.
	Cansavetrack *bool `json:"cansavetrack,omitempty"`
	// Whether the user has the capability mod/scorm:skipview allowed.
	Canskipview *bool `json:"canskipview,omitempty"`
	// Whether the user has the capability mod/scorm:viewreport allowed.
	Canviewreport *bool `json:"canviewreport,omitempty"`
	// Whether the user has the capability mod/scorm:viewscores allowed.
	Canviewscores *bool `json:"canviewscores,omitempty"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

// NewModScormGetScormAccessInformation200Response instantiates a new ModScormGetScormAccessInformation200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModScormGetScormAccessInformation200Response() *ModScormGetScormAccessInformation200Response {
	this := ModScormGetScormAccessInformation200Response{}
	var canaddinstance bool = null
	this.Canaddinstance = &canaddinstance
	var candeleteownresponses bool = null
	this.Candeleteownresponses = &candeleteownresponses
	var candeleteresponses bool = null
	this.Candeleteresponses = &candeleteresponses
	var cansavetrack bool = null
	this.Cansavetrack = &cansavetrack
	var canskipview bool = null
	this.Canskipview = &canskipview
	var canviewreport bool = null
	this.Canviewreport = &canviewreport
	var canviewscores bool = null
	this.Canviewscores = &canviewscores
	return &this
}

// NewModScormGetScormAccessInformation200ResponseWithDefaults instantiates a new ModScormGetScormAccessInformation200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModScormGetScormAccessInformation200ResponseWithDefaults() *ModScormGetScormAccessInformation200Response {
	this := ModScormGetScormAccessInformation200Response{}
	var canaddinstance bool = null
	this.Canaddinstance = &canaddinstance
	var candeleteownresponses bool = null
	this.Candeleteownresponses = &candeleteownresponses
	var candeleteresponses bool = null
	this.Candeleteresponses = &candeleteresponses
	var cansavetrack bool = null
	this.Cansavetrack = &cansavetrack
	var canskipview bool = null
	this.Canskipview = &canskipview
	var canviewreport bool = null
	this.Canviewreport = &canviewreport
	var canviewscores bool = null
	this.Canviewscores = &canviewscores
	return &this
}

// GetCanaddinstance returns the Canaddinstance field value if set, zero value otherwise.
func (o *ModScormGetScormAccessInformation200Response) GetCanaddinstance() bool {
	if o == nil || IsNil(o.Canaddinstance) {
		var ret bool
		return ret
	}
	return *o.Canaddinstance
}

// GetCanaddinstanceOk returns a tuple with the Canaddinstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModScormGetScormAccessInformation200Response) GetCanaddinstanceOk() (*bool, bool) {
	if o == nil || IsNil(o.Canaddinstance) {
		return nil, false
	}
	return o.Canaddinstance, true
}

// HasCanaddinstance returns a boolean if a field has been set.
func (o *ModScormGetScormAccessInformation200Response) HasCanaddinstance() bool {
	if o != nil && !IsNil(o.Canaddinstance) {
		return true
	}

	return false
}

// SetCanaddinstance gets a reference to the given bool and assigns it to the Canaddinstance field.
func (o *ModScormGetScormAccessInformation200Response) SetCanaddinstance(v bool) {
	o.Canaddinstance = &v
}

// GetCandeleteownresponses returns the Candeleteownresponses field value if set, zero value otherwise.
func (o *ModScormGetScormAccessInformation200Response) GetCandeleteownresponses() bool {
	if o == nil || IsNil(o.Candeleteownresponses) {
		var ret bool
		return ret
	}
	return *o.Candeleteownresponses
}

// GetCandeleteownresponsesOk returns a tuple with the Candeleteownresponses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModScormGetScormAccessInformation200Response) GetCandeleteownresponsesOk() (*bool, bool) {
	if o == nil || IsNil(o.Candeleteownresponses) {
		return nil, false
	}
	return o.Candeleteownresponses, true
}

// HasCandeleteownresponses returns a boolean if a field has been set.
func (o *ModScormGetScormAccessInformation200Response) HasCandeleteownresponses() bool {
	if o != nil && !IsNil(o.Candeleteownresponses) {
		return true
	}

	return false
}

// SetCandeleteownresponses gets a reference to the given bool and assigns it to the Candeleteownresponses field.
func (o *ModScormGetScormAccessInformation200Response) SetCandeleteownresponses(v bool) {
	o.Candeleteownresponses = &v
}

// GetCandeleteresponses returns the Candeleteresponses field value if set, zero value otherwise.
func (o *ModScormGetScormAccessInformation200Response) GetCandeleteresponses() bool {
	if o == nil || IsNil(o.Candeleteresponses) {
		var ret bool
		return ret
	}
	return *o.Candeleteresponses
}

// GetCandeleteresponsesOk returns a tuple with the Candeleteresponses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModScormGetScormAccessInformation200Response) GetCandeleteresponsesOk() (*bool, bool) {
	if o == nil || IsNil(o.Candeleteresponses) {
		return nil, false
	}
	return o.Candeleteresponses, true
}

// HasCandeleteresponses returns a boolean if a field has been set.
func (o *ModScormGetScormAccessInformation200Response) HasCandeleteresponses() bool {
	if o != nil && !IsNil(o.Candeleteresponses) {
		return true
	}

	return false
}

// SetCandeleteresponses gets a reference to the given bool and assigns it to the Candeleteresponses field.
func (o *ModScormGetScormAccessInformation200Response) SetCandeleteresponses(v bool) {
	o.Candeleteresponses = &v
}

// GetCansavetrack returns the Cansavetrack field value if set, zero value otherwise.
func (o *ModScormGetScormAccessInformation200Response) GetCansavetrack() bool {
	if o == nil || IsNil(o.Cansavetrack) {
		var ret bool
		return ret
	}
	return *o.Cansavetrack
}

// GetCansavetrackOk returns a tuple with the Cansavetrack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModScormGetScormAccessInformation200Response) GetCansavetrackOk() (*bool, bool) {
	if o == nil || IsNil(o.Cansavetrack) {
		return nil, false
	}
	return o.Cansavetrack, true
}

// HasCansavetrack returns a boolean if a field has been set.
func (o *ModScormGetScormAccessInformation200Response) HasCansavetrack() bool {
	if o != nil && !IsNil(o.Cansavetrack) {
		return true
	}

	return false
}

// SetCansavetrack gets a reference to the given bool and assigns it to the Cansavetrack field.
func (o *ModScormGetScormAccessInformation200Response) SetCansavetrack(v bool) {
	o.Cansavetrack = &v
}

// GetCanskipview returns the Canskipview field value if set, zero value otherwise.
func (o *ModScormGetScormAccessInformation200Response) GetCanskipview() bool {
	if o == nil || IsNil(o.Canskipview) {
		var ret bool
		return ret
	}
	return *o.Canskipview
}

// GetCanskipviewOk returns a tuple with the Canskipview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModScormGetScormAccessInformation200Response) GetCanskipviewOk() (*bool, bool) {
	if o == nil || IsNil(o.Canskipview) {
		return nil, false
	}
	return o.Canskipview, true
}

// HasCanskipview returns a boolean if a field has been set.
func (o *ModScormGetScormAccessInformation200Response) HasCanskipview() bool {
	if o != nil && !IsNil(o.Canskipview) {
		return true
	}

	return false
}

// SetCanskipview gets a reference to the given bool and assigns it to the Canskipview field.
func (o *ModScormGetScormAccessInformation200Response) SetCanskipview(v bool) {
	o.Canskipview = &v
}

// GetCanviewreport returns the Canviewreport field value if set, zero value otherwise.
func (o *ModScormGetScormAccessInformation200Response) GetCanviewreport() bool {
	if o == nil || IsNil(o.Canviewreport) {
		var ret bool
		return ret
	}
	return *o.Canviewreport
}

// GetCanviewreportOk returns a tuple with the Canviewreport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModScormGetScormAccessInformation200Response) GetCanviewreportOk() (*bool, bool) {
	if o == nil || IsNil(o.Canviewreport) {
		return nil, false
	}
	return o.Canviewreport, true
}

// HasCanviewreport returns a boolean if a field has been set.
func (o *ModScormGetScormAccessInformation200Response) HasCanviewreport() bool {
	if o != nil && !IsNil(o.Canviewreport) {
		return true
	}

	return false
}

// SetCanviewreport gets a reference to the given bool and assigns it to the Canviewreport field.
func (o *ModScormGetScormAccessInformation200Response) SetCanviewreport(v bool) {
	o.Canviewreport = &v
}

// GetCanviewscores returns the Canviewscores field value if set, zero value otherwise.
func (o *ModScormGetScormAccessInformation200Response) GetCanviewscores() bool {
	if o == nil || IsNil(o.Canviewscores) {
		var ret bool
		return ret
	}
	return *o.Canviewscores
}

// GetCanviewscoresOk returns a tuple with the Canviewscores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModScormGetScormAccessInformation200Response) GetCanviewscoresOk() (*bool, bool) {
	if o == nil || IsNil(o.Canviewscores) {
		return nil, false
	}
	return o.Canviewscores, true
}

// HasCanviewscores returns a boolean if a field has been set.
func (o *ModScormGetScormAccessInformation200Response) HasCanviewscores() bool {
	if o != nil && !IsNil(o.Canviewscores) {
		return true
	}

	return false
}

// SetCanviewscores gets a reference to the given bool and assigns it to the Canviewscores field.
func (o *ModScormGetScormAccessInformation200Response) SetCanviewscores(v bool) {
	o.Canviewscores = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ModScormGetScormAccessInformation200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModScormGetScormAccessInformation200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ModScormGetScormAccessInformation200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *ModScormGetScormAccessInformation200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o ModScormGetScormAccessInformation200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModScormGetScormAccessInformation200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Canaddinstance) {
		toSerialize["canaddinstance"] = o.Canaddinstance
	}
	if !IsNil(o.Candeleteownresponses) {
		toSerialize["candeleteownresponses"] = o.Candeleteownresponses
	}
	if !IsNil(o.Candeleteresponses) {
		toSerialize["candeleteresponses"] = o.Candeleteresponses
	}
	if !IsNil(o.Cansavetrack) {
		toSerialize["cansavetrack"] = o.Cansavetrack
	}
	if !IsNil(o.Canskipview) {
		toSerialize["canskipview"] = o.Canskipview
	}
	if !IsNil(o.Canviewreport) {
		toSerialize["canviewreport"] = o.Canviewreport
	}
	if !IsNil(o.Canviewscores) {
		toSerialize["canviewscores"] = o.Canviewscores
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableModScormGetScormAccessInformation200Response struct {
	value *ModScormGetScormAccessInformation200Response
	isSet bool
}

func (v NullableModScormGetScormAccessInformation200Response) Get() *ModScormGetScormAccessInformation200Response {
	return v.value
}

func (v *NullableModScormGetScormAccessInformation200Response) Set(val *ModScormGetScormAccessInformation200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModScormGetScormAccessInformation200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModScormGetScormAccessInformation200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModScormGetScormAccessInformation200Response(val *ModScormGetScormAccessInformation200Response) *NullableModScormGetScormAccessInformation200Response {
	return &NullableModScormGetScormAccessInformation200Response{value: val, isSet: true}
}

func (v NullableModScormGetScormAccessInformation200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModScormGetScormAccessInformation200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


