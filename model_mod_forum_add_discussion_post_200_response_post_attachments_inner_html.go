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

// checks if the ModForumAddDiscussionPost200ResponsePostAttachmentsInnerHtml type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModForumAddDiscussionPost200ResponsePostAttachmentsInnerHtml{}

// ModForumAddDiscussionPost200ResponsePostAttachmentsInnerHtml struct for ModForumAddDiscussionPost200ResponsePostAttachmentsInnerHtml
type ModForumAddDiscussionPost200ResponsePostAttachmentsInnerHtml struct {
	// The HTML source for the Plagiarism Response
	Plagiarism *string `json:"plagiarism,omitempty"`
}

// NewModForumAddDiscussionPost200ResponsePostAttachmentsInnerHtml instantiates a new ModForumAddDiscussionPost200ResponsePostAttachmentsInnerHtml object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModForumAddDiscussionPost200ResponsePostAttachmentsInnerHtml() *ModForumAddDiscussionPost200ResponsePostAttachmentsInnerHtml {
	this := ModForumAddDiscussionPost200ResponsePostAttachmentsInnerHtml{}
	var plagiarism string = "null"
	this.Plagiarism = &plagiarism
	return &this
}

// NewModForumAddDiscussionPost200ResponsePostAttachmentsInnerHtmlWithDefaults instantiates a new ModForumAddDiscussionPost200ResponsePostAttachmentsInnerHtml object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModForumAddDiscussionPost200ResponsePostAttachmentsInnerHtmlWithDefaults() *ModForumAddDiscussionPost200ResponsePostAttachmentsInnerHtml {
	this := ModForumAddDiscussionPost200ResponsePostAttachmentsInnerHtml{}
	var plagiarism string = "null"
	this.Plagiarism = &plagiarism
	return &this
}

// GetPlagiarism returns the Plagiarism field value if set, zero value otherwise.
func (o *ModForumAddDiscussionPost200ResponsePostAttachmentsInnerHtml) GetPlagiarism() string {
	if o == nil || IsNil(o.Plagiarism) {
		var ret string
		return ret
	}
	return *o.Plagiarism
}

// GetPlagiarismOk returns a tuple with the Plagiarism field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumAddDiscussionPost200ResponsePostAttachmentsInnerHtml) GetPlagiarismOk() (*string, bool) {
	if o == nil || IsNil(o.Plagiarism) {
		return nil, false
	}
	return o.Plagiarism, true
}

// HasPlagiarism returns a boolean if a field has been set.
func (o *ModForumAddDiscussionPost200ResponsePostAttachmentsInnerHtml) HasPlagiarism() bool {
	if o != nil && !IsNil(o.Plagiarism) {
		return true
	}

	return false
}

// SetPlagiarism gets a reference to the given string and assigns it to the Plagiarism field.
func (o *ModForumAddDiscussionPost200ResponsePostAttachmentsInnerHtml) SetPlagiarism(v string) {
	o.Plagiarism = &v
}

func (o ModForumAddDiscussionPost200ResponsePostAttachmentsInnerHtml) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModForumAddDiscussionPost200ResponsePostAttachmentsInnerHtml) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Plagiarism) {
		toSerialize["plagiarism"] = o.Plagiarism
	}
	return toSerialize, nil
}

type NullableModForumAddDiscussionPost200ResponsePostAttachmentsInnerHtml struct {
	value *ModForumAddDiscussionPost200ResponsePostAttachmentsInnerHtml
	isSet bool
}

func (v NullableModForumAddDiscussionPost200ResponsePostAttachmentsInnerHtml) Get() *ModForumAddDiscussionPost200ResponsePostAttachmentsInnerHtml {
	return v.value
}

func (v *NullableModForumAddDiscussionPost200ResponsePostAttachmentsInnerHtml) Set(val *ModForumAddDiscussionPost200ResponsePostAttachmentsInnerHtml) {
	v.value = val
	v.isSet = true
}

func (v NullableModForumAddDiscussionPost200ResponsePostAttachmentsInnerHtml) IsSet() bool {
	return v.isSet
}

func (v *NullableModForumAddDiscussionPost200ResponsePostAttachmentsInnerHtml) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModForumAddDiscussionPost200ResponsePostAttachmentsInnerHtml(val *ModForumAddDiscussionPost200ResponsePostAttachmentsInnerHtml) *NullableModForumAddDiscussionPost200ResponsePostAttachmentsInnerHtml {
	return &NullableModForumAddDiscussionPost200ResponsePostAttachmentsInnerHtml{value: val, isSet: true}
}

func (v NullableModForumAddDiscussionPost200ResponsePostAttachmentsInnerHtml) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModForumAddDiscussionPost200ResponsePostAttachmentsInnerHtml) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


