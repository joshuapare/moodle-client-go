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

// checks if the ModWorkshopGetUserPlan200ResponseUserplanPhasesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModWorkshopGetUserPlan200ResponseUserplanPhasesInner{}

// ModWorkshopGetUserPlan200ResponseUserplanPhasesInner struct for ModWorkshopGetUserPlan200ResponseUserplanPhasesInner
type ModWorkshopGetUserPlan200ResponseUserplanPhasesInner struct {
	Actions []ModWorkshopGetUserPlan200ResponseUserplanPhasesInnerActionsInner `json:"actions,omitempty"`
	// Whether is the active task.
	Active *bool `json:"active,omitempty"`
	// Phase code.
	Code *int32 `json:"code,omitempty"`
	Tasks []ModWorkshopGetUserPlan200ResponseUserplanPhasesInnerTasksInner `json:"tasks,omitempty"`
	// Phase title.
	Title *string `json:"title,omitempty"`
}

// NewModWorkshopGetUserPlan200ResponseUserplanPhasesInner instantiates a new ModWorkshopGetUserPlan200ResponseUserplanPhasesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModWorkshopGetUserPlan200ResponseUserplanPhasesInner() *ModWorkshopGetUserPlan200ResponseUserplanPhasesInner {
	this := ModWorkshopGetUserPlan200ResponseUserplanPhasesInner{}
	var active bool = null
	this.Active = &active
	var code int32 = null
	this.Code = &code
	var title string = "null"
	this.Title = &title
	return &this
}

// NewModWorkshopGetUserPlan200ResponseUserplanPhasesInnerWithDefaults instantiates a new ModWorkshopGetUserPlan200ResponseUserplanPhasesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModWorkshopGetUserPlan200ResponseUserplanPhasesInnerWithDefaults() *ModWorkshopGetUserPlan200ResponseUserplanPhasesInner {
	this := ModWorkshopGetUserPlan200ResponseUserplanPhasesInner{}
	var active bool = null
	this.Active = &active
	var code int32 = null
	this.Code = &code
	var title string = "null"
	this.Title = &title
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *ModWorkshopGetUserPlan200ResponseUserplanPhasesInner) GetActions() []ModWorkshopGetUserPlan200ResponseUserplanPhasesInnerActionsInner {
	if o == nil || IsNil(o.Actions) {
		var ret []ModWorkshopGetUserPlan200ResponseUserplanPhasesInnerActionsInner
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModWorkshopGetUserPlan200ResponseUserplanPhasesInner) GetActionsOk() ([]ModWorkshopGetUserPlan200ResponseUserplanPhasesInnerActionsInner, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *ModWorkshopGetUserPlan200ResponseUserplanPhasesInner) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []ModWorkshopGetUserPlan200ResponseUserplanPhasesInnerActionsInner and assigns it to the Actions field.
func (o *ModWorkshopGetUserPlan200ResponseUserplanPhasesInner) SetActions(v []ModWorkshopGetUserPlan200ResponseUserplanPhasesInnerActionsInner) {
	o.Actions = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ModWorkshopGetUserPlan200ResponseUserplanPhasesInner) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModWorkshopGetUserPlan200ResponseUserplanPhasesInner) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ModWorkshopGetUserPlan200ResponseUserplanPhasesInner) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ModWorkshopGetUserPlan200ResponseUserplanPhasesInner) SetActive(v bool) {
	o.Active = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ModWorkshopGetUserPlan200ResponseUserplanPhasesInner) GetCode() int32 {
	if o == nil || IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModWorkshopGetUserPlan200ResponseUserplanPhasesInner) GetCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ModWorkshopGetUserPlan200ResponseUserplanPhasesInner) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *ModWorkshopGetUserPlan200ResponseUserplanPhasesInner) SetCode(v int32) {
	o.Code = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *ModWorkshopGetUserPlan200ResponseUserplanPhasesInner) GetTasks() []ModWorkshopGetUserPlan200ResponseUserplanPhasesInnerTasksInner {
	if o == nil || IsNil(o.Tasks) {
		var ret []ModWorkshopGetUserPlan200ResponseUserplanPhasesInnerTasksInner
		return ret
	}
	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModWorkshopGetUserPlan200ResponseUserplanPhasesInner) GetTasksOk() ([]ModWorkshopGetUserPlan200ResponseUserplanPhasesInnerTasksInner, bool) {
	if o == nil || IsNil(o.Tasks) {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *ModWorkshopGetUserPlan200ResponseUserplanPhasesInner) HasTasks() bool {
	if o != nil && !IsNil(o.Tasks) {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []ModWorkshopGetUserPlan200ResponseUserplanPhasesInnerTasksInner and assigns it to the Tasks field.
func (o *ModWorkshopGetUserPlan200ResponseUserplanPhasesInner) SetTasks(v []ModWorkshopGetUserPlan200ResponseUserplanPhasesInnerTasksInner) {
	o.Tasks = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ModWorkshopGetUserPlan200ResponseUserplanPhasesInner) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModWorkshopGetUserPlan200ResponseUserplanPhasesInner) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ModWorkshopGetUserPlan200ResponseUserplanPhasesInner) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ModWorkshopGetUserPlan200ResponseUserplanPhasesInner) SetTitle(v string) {
	o.Title = &v
}

func (o ModWorkshopGetUserPlan200ResponseUserplanPhasesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModWorkshopGetUserPlan200ResponseUserplanPhasesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Actions) {
		toSerialize["actions"] = o.Actions
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Tasks) {
		toSerialize["tasks"] = o.Tasks
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableModWorkshopGetUserPlan200ResponseUserplanPhasesInner struct {
	value *ModWorkshopGetUserPlan200ResponseUserplanPhasesInner
	isSet bool
}

func (v NullableModWorkshopGetUserPlan200ResponseUserplanPhasesInner) Get() *ModWorkshopGetUserPlan200ResponseUserplanPhasesInner {
	return v.value
}

func (v *NullableModWorkshopGetUserPlan200ResponseUserplanPhasesInner) Set(val *ModWorkshopGetUserPlan200ResponseUserplanPhasesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableModWorkshopGetUserPlan200ResponseUserplanPhasesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableModWorkshopGetUserPlan200ResponseUserplanPhasesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModWorkshopGetUserPlan200ResponseUserplanPhasesInner(val *ModWorkshopGetUserPlan200ResponseUserplanPhasesInner) *NullableModWorkshopGetUserPlan200ResponseUserplanPhasesInner {
	return &NullableModWorkshopGetUserPlan200ResponseUserplanPhasesInner{value: val, isSet: true}
}

func (v NullableModWorkshopGetUserPlan200ResponseUserplanPhasesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModWorkshopGetUserPlan200ResponseUserplanPhasesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


