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

// checks if the CoreFilesGetUnusedDraftItemid200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreFilesGetUnusedDraftItemid200Response{}

// CoreFilesGetUnusedDraftItemid200Response struct for CoreFilesGetUnusedDraftItemid200Response
type CoreFilesGetUnusedDraftItemid200Response struct {
	// File area component.
	Component string `json:"component"`
	// File area context.
	Contextid int32 `json:"contextid"`
	// File area name.
	Filearea string `json:"filearea"`
	// File are item id.
	Itemid int32 `json:"itemid"`
	// File area user id.
	Userid int32 `json:"userid"`
	Warnings []AuthEmailSignupUser200ResponseWarningsInner `json:"warnings,omitempty"`
}

type _CoreFilesGetUnusedDraftItemid200Response CoreFilesGetUnusedDraftItemid200Response

// NewCoreFilesGetUnusedDraftItemid200Response instantiates a new CoreFilesGetUnusedDraftItemid200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreFilesGetUnusedDraftItemid200Response(component string, contextid int32, filearea string, itemid int32, userid int32) *CoreFilesGetUnusedDraftItemid200Response {
	this := CoreFilesGetUnusedDraftItemid200Response{}
	this.Component = component
	this.Contextid = contextid
	this.Filearea = filearea
	this.Itemid = itemid
	this.Userid = userid
	return &this
}

// NewCoreFilesGetUnusedDraftItemid200ResponseWithDefaults instantiates a new CoreFilesGetUnusedDraftItemid200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreFilesGetUnusedDraftItemid200ResponseWithDefaults() *CoreFilesGetUnusedDraftItemid200Response {
	this := CoreFilesGetUnusedDraftItemid200Response{}
	var component string = "null"
	this.Component = component
	var contextid int32 = null
	this.Contextid = contextid
	var filearea string = "null"
	this.Filearea = filearea
	var itemid int32 = null
	this.Itemid = itemid
	var userid int32 = null
	this.Userid = userid
	return &this
}

// GetComponent returns the Component field value
func (o *CoreFilesGetUnusedDraftItemid200Response) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *CoreFilesGetUnusedDraftItemid200Response) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *CoreFilesGetUnusedDraftItemid200Response) SetComponent(v string) {
	o.Component = v
}

// GetContextid returns the Contextid field value
func (o *CoreFilesGetUnusedDraftItemid200Response) GetContextid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Contextid
}

// GetContextidOk returns a tuple with the Contextid field value
// and a boolean to check if the value has been set.
func (o *CoreFilesGetUnusedDraftItemid200Response) GetContextidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contextid, true
}

// SetContextid sets field value
func (o *CoreFilesGetUnusedDraftItemid200Response) SetContextid(v int32) {
	o.Contextid = v
}

// GetFilearea returns the Filearea field value
func (o *CoreFilesGetUnusedDraftItemid200Response) GetFilearea() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Filearea
}

// GetFileareaOk returns a tuple with the Filearea field value
// and a boolean to check if the value has been set.
func (o *CoreFilesGetUnusedDraftItemid200Response) GetFileareaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filearea, true
}

// SetFilearea sets field value
func (o *CoreFilesGetUnusedDraftItemid200Response) SetFilearea(v string) {
	o.Filearea = v
}

// GetItemid returns the Itemid field value
func (o *CoreFilesGetUnusedDraftItemid200Response) GetItemid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Itemid
}

// GetItemidOk returns a tuple with the Itemid field value
// and a boolean to check if the value has been set.
func (o *CoreFilesGetUnusedDraftItemid200Response) GetItemidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Itemid, true
}

// SetItemid sets field value
func (o *CoreFilesGetUnusedDraftItemid200Response) SetItemid(v int32) {
	o.Itemid = v
}

// GetUserid returns the Userid field value
func (o *CoreFilesGetUnusedDraftItemid200Response) GetUserid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Userid
}

// GetUseridOk returns a tuple with the Userid field value
// and a boolean to check if the value has been set.
func (o *CoreFilesGetUnusedDraftItemid200Response) GetUseridOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Userid, true
}

// SetUserid sets field value
func (o *CoreFilesGetUnusedDraftItemid200Response) SetUserid(v int32) {
	o.Userid = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *CoreFilesGetUnusedDraftItemid200Response) GetWarnings() []AuthEmailSignupUser200ResponseWarningsInner {
	if o == nil || IsNil(o.Warnings) {
		var ret []AuthEmailSignupUser200ResponseWarningsInner
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreFilesGetUnusedDraftItemid200Response) GetWarningsOk() ([]AuthEmailSignupUser200ResponseWarningsInner, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *CoreFilesGetUnusedDraftItemid200Response) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []AuthEmailSignupUser200ResponseWarningsInner and assigns it to the Warnings field.
func (o *CoreFilesGetUnusedDraftItemid200Response) SetWarnings(v []AuthEmailSignupUser200ResponseWarningsInner) {
	o.Warnings = v
}

func (o CoreFilesGetUnusedDraftItemid200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreFilesGetUnusedDraftItemid200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["component"] = o.Component
	toSerialize["contextid"] = o.Contextid
	toSerialize["filearea"] = o.Filearea
	toSerialize["itemid"] = o.Itemid
	toSerialize["userid"] = o.Userid
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

func (o *CoreFilesGetUnusedDraftItemid200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"component",
		"contextid",
		"filearea",
		"itemid",
		"userid",
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

	varCoreFilesGetUnusedDraftItemid200Response := _CoreFilesGetUnusedDraftItemid200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreFilesGetUnusedDraftItemid200Response)

	if err != nil {
		return err
	}

	*o = CoreFilesGetUnusedDraftItemid200Response(varCoreFilesGetUnusedDraftItemid200Response)

	return err
}

type NullableCoreFilesGetUnusedDraftItemid200Response struct {
	value *CoreFilesGetUnusedDraftItemid200Response
	isSet bool
}

func (v NullableCoreFilesGetUnusedDraftItemid200Response) Get() *CoreFilesGetUnusedDraftItemid200Response {
	return v.value
}

func (v *NullableCoreFilesGetUnusedDraftItemid200Response) Set(val *CoreFilesGetUnusedDraftItemid200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreFilesGetUnusedDraftItemid200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreFilesGetUnusedDraftItemid200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreFilesGetUnusedDraftItemid200Response(val *CoreFilesGetUnusedDraftItemid200Response) *NullableCoreFilesGetUnusedDraftItemid200Response {
	return &NullableCoreFilesGetUnusedDraftItemid200Response{value: val, isSet: true}
}

func (v NullableCoreFilesGetUnusedDraftItemid200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreFilesGetUnusedDraftItemid200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

