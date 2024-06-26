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

// checks if the ModWorkshopAddSubmissionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModWorkshopAddSubmissionRequest{}

// ModWorkshopAddSubmissionRequest struct for ModWorkshopAddSubmissionRequest
type ModWorkshopAddSubmissionRequest struct {
	// The draft file area id for attachments
	Attachmentsid *int32 `json:"attachmentsid,omitempty"`
	// Submission text content
	Content *string `json:"content,omitempty"`
	// The format used for the content
	Contentformat *int32 `json:"contentformat,omitempty"`
	// The draft file area id for inline attachments in the content
	Inlineattachmentsid *int32 `json:"inlineattachmentsid,omitempty"`
	// Submission title
	Title string `json:"title"`
	// Workshop id
	Workshopid int32 `json:"workshopid"`
}

type _ModWorkshopAddSubmissionRequest ModWorkshopAddSubmissionRequest

// NewModWorkshopAddSubmissionRequest instantiates a new ModWorkshopAddSubmissionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModWorkshopAddSubmissionRequest(title string, workshopid int32) *ModWorkshopAddSubmissionRequest {
	this := ModWorkshopAddSubmissionRequest{}
	var attachmentsid int32 = 0
	this.Attachmentsid = &attachmentsid
	var content string = ""
	this.Content = &content
	var contentformat int32 = 0
	this.Contentformat = &contentformat
	var inlineattachmentsid int32 = 0
	this.Inlineattachmentsid = &inlineattachmentsid
	this.Title = title
	this.Workshopid = workshopid
	return &this
}

// NewModWorkshopAddSubmissionRequestWithDefaults instantiates a new ModWorkshopAddSubmissionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModWorkshopAddSubmissionRequestWithDefaults() *ModWorkshopAddSubmissionRequest {
	this := ModWorkshopAddSubmissionRequest{}
	var attachmentsid int32 = 0
	this.Attachmentsid = &attachmentsid
	var content string = ""
	this.Content = &content
	var contentformat int32 = 0
	this.Contentformat = &contentformat
	var inlineattachmentsid int32 = 0
	this.Inlineattachmentsid = &inlineattachmentsid
	var title string = "null"
	this.Title = title
	var workshopid int32 = null
	this.Workshopid = workshopid
	return &this
}

// GetAttachmentsid returns the Attachmentsid field value if set, zero value otherwise.
func (o *ModWorkshopAddSubmissionRequest) GetAttachmentsid() int32 {
	if o == nil || IsNil(o.Attachmentsid) {
		var ret int32
		return ret
	}
	return *o.Attachmentsid
}

// GetAttachmentsidOk returns a tuple with the Attachmentsid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModWorkshopAddSubmissionRequest) GetAttachmentsidOk() (*int32, bool) {
	if o == nil || IsNil(o.Attachmentsid) {
		return nil, false
	}
	return o.Attachmentsid, true
}

// HasAttachmentsid returns a boolean if a field has been set.
func (o *ModWorkshopAddSubmissionRequest) HasAttachmentsid() bool {
	if o != nil && !IsNil(o.Attachmentsid) {
		return true
	}

	return false
}

// SetAttachmentsid gets a reference to the given int32 and assigns it to the Attachmentsid field.
func (o *ModWorkshopAddSubmissionRequest) SetAttachmentsid(v int32) {
	o.Attachmentsid = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *ModWorkshopAddSubmissionRequest) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModWorkshopAddSubmissionRequest) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *ModWorkshopAddSubmissionRequest) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *ModWorkshopAddSubmissionRequest) SetContent(v string) {
	o.Content = &v
}

// GetContentformat returns the Contentformat field value if set, zero value otherwise.
func (o *ModWorkshopAddSubmissionRequest) GetContentformat() int32 {
	if o == nil || IsNil(o.Contentformat) {
		var ret int32
		return ret
	}
	return *o.Contentformat
}

// GetContentformatOk returns a tuple with the Contentformat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModWorkshopAddSubmissionRequest) GetContentformatOk() (*int32, bool) {
	if o == nil || IsNil(o.Contentformat) {
		return nil, false
	}
	return o.Contentformat, true
}

// HasContentformat returns a boolean if a field has been set.
func (o *ModWorkshopAddSubmissionRequest) HasContentformat() bool {
	if o != nil && !IsNil(o.Contentformat) {
		return true
	}

	return false
}

// SetContentformat gets a reference to the given int32 and assigns it to the Contentformat field.
func (o *ModWorkshopAddSubmissionRequest) SetContentformat(v int32) {
	o.Contentformat = &v
}

// GetInlineattachmentsid returns the Inlineattachmentsid field value if set, zero value otherwise.
func (o *ModWorkshopAddSubmissionRequest) GetInlineattachmentsid() int32 {
	if o == nil || IsNil(o.Inlineattachmentsid) {
		var ret int32
		return ret
	}
	return *o.Inlineattachmentsid
}

// GetInlineattachmentsidOk returns a tuple with the Inlineattachmentsid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModWorkshopAddSubmissionRequest) GetInlineattachmentsidOk() (*int32, bool) {
	if o == nil || IsNil(o.Inlineattachmentsid) {
		return nil, false
	}
	return o.Inlineattachmentsid, true
}

// HasInlineattachmentsid returns a boolean if a field has been set.
func (o *ModWorkshopAddSubmissionRequest) HasInlineattachmentsid() bool {
	if o != nil && !IsNil(o.Inlineattachmentsid) {
		return true
	}

	return false
}

// SetInlineattachmentsid gets a reference to the given int32 and assigns it to the Inlineattachmentsid field.
func (o *ModWorkshopAddSubmissionRequest) SetInlineattachmentsid(v int32) {
	o.Inlineattachmentsid = &v
}

// GetTitle returns the Title field value
func (o *ModWorkshopAddSubmissionRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ModWorkshopAddSubmissionRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ModWorkshopAddSubmissionRequest) SetTitle(v string) {
	o.Title = v
}

// GetWorkshopid returns the Workshopid field value
func (o *ModWorkshopAddSubmissionRequest) GetWorkshopid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Workshopid
}

// GetWorkshopidOk returns a tuple with the Workshopid field value
// and a boolean to check if the value has been set.
func (o *ModWorkshopAddSubmissionRequest) GetWorkshopidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Workshopid, true
}

// SetWorkshopid sets field value
func (o *ModWorkshopAddSubmissionRequest) SetWorkshopid(v int32) {
	o.Workshopid = v
}

func (o ModWorkshopAddSubmissionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModWorkshopAddSubmissionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attachmentsid) {
		toSerialize["attachmentsid"] = o.Attachmentsid
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.Contentformat) {
		toSerialize["contentformat"] = o.Contentformat
	}
	if !IsNil(o.Inlineattachmentsid) {
		toSerialize["inlineattachmentsid"] = o.Inlineattachmentsid
	}
	toSerialize["title"] = o.Title
	toSerialize["workshopid"] = o.Workshopid
	return toSerialize, nil
}

func (o *ModWorkshopAddSubmissionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
		"workshopid",
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

	varModWorkshopAddSubmissionRequest := _ModWorkshopAddSubmissionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModWorkshopAddSubmissionRequest)

	if err != nil {
		return err
	}

	*o = ModWorkshopAddSubmissionRequest(varModWorkshopAddSubmissionRequest)

	return err
}

type NullableModWorkshopAddSubmissionRequest struct {
	value *ModWorkshopAddSubmissionRequest
	isSet bool
}

func (v NullableModWorkshopAddSubmissionRequest) Get() *ModWorkshopAddSubmissionRequest {
	return v.value
}

func (v *NullableModWorkshopAddSubmissionRequest) Set(val *ModWorkshopAddSubmissionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModWorkshopAddSubmissionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModWorkshopAddSubmissionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModWorkshopAddSubmissionRequest(val *ModWorkshopAddSubmissionRequest) *NullableModWorkshopAddSubmissionRequest {
	return &NullableModWorkshopAddSubmissionRequest{value: val, isSet: true}
}

func (v NullableModWorkshopAddSubmissionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModWorkshopAddSubmissionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


