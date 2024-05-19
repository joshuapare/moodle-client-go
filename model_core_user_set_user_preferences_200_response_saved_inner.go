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

// checks if the CoreUserSetUserPreferences200ResponseSavedInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreUserSetUserPreferences200ResponseSavedInner{}

// CoreUserSetUserPreferences200ResponseSavedInner struct for CoreUserSetUserPreferences200ResponseSavedInner
type CoreUserSetUserPreferences200ResponseSavedInner struct {
	// The name of the preference
	Name *string `json:"name,omitempty"`
	// The user the preference was set for
	Userid *int32 `json:"userid,omitempty"`
}

// NewCoreUserSetUserPreferences200ResponseSavedInner instantiates a new CoreUserSetUserPreferences200ResponseSavedInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreUserSetUserPreferences200ResponseSavedInner() *CoreUserSetUserPreferences200ResponseSavedInner {
	this := CoreUserSetUserPreferences200ResponseSavedInner{}
	var userid int32 = null
	this.Userid = &userid
	return &this
}

// NewCoreUserSetUserPreferences200ResponseSavedInnerWithDefaults instantiates a new CoreUserSetUserPreferences200ResponseSavedInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreUserSetUserPreferences200ResponseSavedInnerWithDefaults() *CoreUserSetUserPreferences200ResponseSavedInner {
	this := CoreUserSetUserPreferences200ResponseSavedInner{}
	var userid int32 = null
	this.Userid = &userid
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CoreUserSetUserPreferences200ResponseSavedInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreUserSetUserPreferences200ResponseSavedInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CoreUserSetUserPreferences200ResponseSavedInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CoreUserSetUserPreferences200ResponseSavedInner) SetName(v string) {
	o.Name = &v
}

// GetUserid returns the Userid field value if set, zero value otherwise.
func (o *CoreUserSetUserPreferences200ResponseSavedInner) GetUserid() int32 {
	if o == nil || IsNil(o.Userid) {
		var ret int32
		return ret
	}
	return *o.Userid
}

// GetUseridOk returns a tuple with the Userid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreUserSetUserPreferences200ResponseSavedInner) GetUseridOk() (*int32, bool) {
	if o == nil || IsNil(o.Userid) {
		return nil, false
	}
	return o.Userid, true
}

// HasUserid returns a boolean if a field has been set.
func (o *CoreUserSetUserPreferences200ResponseSavedInner) HasUserid() bool {
	if o != nil && !IsNil(o.Userid) {
		return true
	}

	return false
}

// SetUserid gets a reference to the given int32 and assigns it to the Userid field.
func (o *CoreUserSetUserPreferences200ResponseSavedInner) SetUserid(v int32) {
	o.Userid = &v
}

func (o CoreUserSetUserPreferences200ResponseSavedInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreUserSetUserPreferences200ResponseSavedInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Userid) {
		toSerialize["userid"] = o.Userid
	}
	return toSerialize, nil
}

type NullableCoreUserSetUserPreferences200ResponseSavedInner struct {
	value *CoreUserSetUserPreferences200ResponseSavedInner
	isSet bool
}

func (v NullableCoreUserSetUserPreferences200ResponseSavedInner) Get() *CoreUserSetUserPreferences200ResponseSavedInner {
	return v.value
}

func (v *NullableCoreUserSetUserPreferences200ResponseSavedInner) Set(val *CoreUserSetUserPreferences200ResponseSavedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreUserSetUserPreferences200ResponseSavedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreUserSetUserPreferences200ResponseSavedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreUserSetUserPreferences200ResponseSavedInner(val *CoreUserSetUserPreferences200ResponseSavedInner) *NullableCoreUserSetUserPreferences200ResponseSavedInner {
	return &NullableCoreUserSetUserPreferences200ResponseSavedInner{value: val, isSet: true}
}

func (v NullableCoreUserSetUserPreferences200ResponseSavedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreUserSetUserPreferences200ResponseSavedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

