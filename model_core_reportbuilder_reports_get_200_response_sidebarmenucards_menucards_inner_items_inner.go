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

// checks if the CoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner{}

// CoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner struct for CoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner
type CoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner struct {
	// action
	Action *string `json:"action,omitempty"`
	// disabled
	Disabled *bool `json:"disabled,omitempty"`
	// identifier
	Identifier *string `json:"identifier,omitempty"`
	// name
	Name *string `json:"name,omitempty"`
	// title
	Title *string `json:"title,omitempty"`
}

// NewCoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner instantiates a new CoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner() *CoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner {
	this := CoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner{}
	var action string = "null"
	this.Action = &action
	var disabled bool = false
	this.Disabled = &disabled
	var identifier string = "null"
	this.Identifier = &identifier
	return &this
}

// NewCoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInnerWithDefaults instantiates a new CoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInnerWithDefaults() *CoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner {
	this := CoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner{}
	var action string = "null"
	this.Action = &action
	var disabled bool = false
	this.Disabled = &disabled
	var identifier string = "null"
	this.Identifier = &identifier
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *CoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *CoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *CoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner) SetAction(v string) {
	o.Action = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *CoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *CoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *CoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *CoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner) GetIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *CoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *CoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner) SetName(v string) {
	o.Name = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *CoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *CoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *CoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner) SetTitle(v string) {
	o.Title = &v
}

func (o CoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableCoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner struct {
	value *CoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner
	isSet bool
}

func (v NullableCoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner) Get() *CoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner {
	return v.value
}

func (v *NullableCoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner) Set(val *CoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner(val *CoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner) *NullableCoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner {
	return &NullableCoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner{value: val, isSet: true}
}

func (v NullableCoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreReportbuilderReportsGet200ResponseSidebarmenucardsMenucardsInnerItemsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

