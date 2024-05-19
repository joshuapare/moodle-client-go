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

// checks if the ModWikiEditPageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModWikiEditPageRequest{}

// ModWikiEditPageRequest struct for ModWikiEditPageRequest
type ModWikiEditPageRequest struct {
	// Page contents.
	Content string `json:"content"`
	// Page ID.
	Pageid int32 `json:"pageid"`
	// Section page title.
	Section *string `json:"section,omitempty"`
}

type _ModWikiEditPageRequest ModWikiEditPageRequest

// NewModWikiEditPageRequest instantiates a new ModWikiEditPageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModWikiEditPageRequest(content string, pageid int32) *ModWikiEditPageRequest {
	this := ModWikiEditPageRequest{}
	this.Content = content
	this.Pageid = pageid
	var section string = "null"
	this.Section = &section
	return &this
}

// NewModWikiEditPageRequestWithDefaults instantiates a new ModWikiEditPageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModWikiEditPageRequestWithDefaults() *ModWikiEditPageRequest {
	this := ModWikiEditPageRequest{}
	var pageid int32 = null
	this.Pageid = pageid
	var section string = "null"
	this.Section = &section
	return &this
}

// GetContent returns the Content field value
func (o *ModWikiEditPageRequest) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *ModWikiEditPageRequest) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *ModWikiEditPageRequest) SetContent(v string) {
	o.Content = v
}

// GetPageid returns the Pageid field value
func (o *ModWikiEditPageRequest) GetPageid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pageid
}

// GetPageidOk returns a tuple with the Pageid field value
// and a boolean to check if the value has been set.
func (o *ModWikiEditPageRequest) GetPageidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pageid, true
}

// SetPageid sets field value
func (o *ModWikiEditPageRequest) SetPageid(v int32) {
	o.Pageid = v
}

// GetSection returns the Section field value if set, zero value otherwise.
func (o *ModWikiEditPageRequest) GetSection() string {
	if o == nil || IsNil(o.Section) {
		var ret string
		return ret
	}
	return *o.Section
}

// GetSectionOk returns a tuple with the Section field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModWikiEditPageRequest) GetSectionOk() (*string, bool) {
	if o == nil || IsNil(o.Section) {
		return nil, false
	}
	return o.Section, true
}

// HasSection returns a boolean if a field has been set.
func (o *ModWikiEditPageRequest) HasSection() bool {
	if o != nil && !IsNil(o.Section) {
		return true
	}

	return false
}

// SetSection gets a reference to the given string and assigns it to the Section field.
func (o *ModWikiEditPageRequest) SetSection(v string) {
	o.Section = &v
}

func (o ModWikiEditPageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModWikiEditPageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["content"] = o.Content
	toSerialize["pageid"] = o.Pageid
	if !IsNil(o.Section) {
		toSerialize["section"] = o.Section
	}
	return toSerialize, nil
}

func (o *ModWikiEditPageRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"content",
		"pageid",
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

	varModWikiEditPageRequest := _ModWikiEditPageRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModWikiEditPageRequest)

	if err != nil {
		return err
	}

	*o = ModWikiEditPageRequest(varModWikiEditPageRequest)

	return err
}

type NullableModWikiEditPageRequest struct {
	value *ModWikiEditPageRequest
	isSet bool
}

func (v NullableModWikiEditPageRequest) Get() *ModWikiEditPageRequest {
	return v.value
}

func (v *NullableModWikiEditPageRequest) Set(val *ModWikiEditPageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModWikiEditPageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModWikiEditPageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModWikiEditPageRequest(val *ModWikiEditPageRequest) *NullableModWikiEditPageRequest {
	return &NullableModWikiEditPageRequest{value: val, isSet: true}
}

func (v NullableModWikiEditPageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModWikiEditPageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

