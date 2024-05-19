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

// checks if the CoreBlockGetDashboardBlocks200ResponseBlocksInnerContents type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreBlockGetDashboardBlocks200ResponseBlocksInnerContents{}

// CoreBlockGetDashboardBlocks200ResponseBlocksInnerContents struct for CoreBlockGetDashboardBlocks200ResponseBlocksInnerContents
type CoreBlockGetDashboardBlocks200ResponseBlocksInnerContents struct {
	// Block contents.
	Content string `json:"content"`
	// content format (1 = HTML, 0 = MOODLE, 2 = PLAIN, or 4 = MARKDOWN)
	Contentformat int32 `json:"contentformat"`
	Files []CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner `json:"files"`
	// Block footer.
	Footer string `json:"footer"`
	// Block title.
	Title string `json:"title"`
}

type _CoreBlockGetDashboardBlocks200ResponseBlocksInnerContents CoreBlockGetDashboardBlocks200ResponseBlocksInnerContents

// NewCoreBlockGetDashboardBlocks200ResponseBlocksInnerContents instantiates a new CoreBlockGetDashboardBlocks200ResponseBlocksInnerContents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreBlockGetDashboardBlocks200ResponseBlocksInnerContents(content string, contentformat int32, files []CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner, footer string, title string) *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContents {
	this := CoreBlockGetDashboardBlocks200ResponseBlocksInnerContents{}
	this.Content = content
	this.Contentformat = contentformat
	this.Files = files
	this.Footer = footer
	this.Title = title
	return &this
}

// NewCoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsWithDefaults instantiates a new CoreBlockGetDashboardBlocks200ResponseBlocksInnerContents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsWithDefaults() *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContents {
	this := CoreBlockGetDashboardBlocks200ResponseBlocksInnerContents{}
	return &this
}

// GetContent returns the Content field value
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContents) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContents) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContents) SetContent(v string) {
	o.Content = v
}

// GetContentformat returns the Contentformat field value
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContents) GetContentformat() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Contentformat
}

// GetContentformatOk returns a tuple with the Contentformat field value
// and a boolean to check if the value has been set.
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContents) GetContentformatOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contentformat, true
}

// SetContentformat sets field value
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContents) SetContentformat(v int32) {
	o.Contentformat = v
}

// GetFiles returns the Files field value
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContents) GetFiles() []CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner {
	if o == nil {
		var ret []CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner
		return ret
	}

	return o.Files
}

// GetFilesOk returns a tuple with the Files field value
// and a boolean to check if the value has been set.
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContents) GetFilesOk() ([]CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Files, true
}

// SetFiles sets field value
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContents) SetFiles(v []CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) {
	o.Files = v
}

// GetFooter returns the Footer field value
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContents) GetFooter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Footer
}

// GetFooterOk returns a tuple with the Footer field value
// and a boolean to check if the value has been set.
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContents) GetFooterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Footer, true
}

// SetFooter sets field value
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContents) SetFooter(v string) {
	o.Footer = v
}

// GetTitle returns the Title field value
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContents) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContents) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContents) SetTitle(v string) {
	o.Title = v
}

func (o CoreBlockGetDashboardBlocks200ResponseBlocksInnerContents) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreBlockGetDashboardBlocks200ResponseBlocksInnerContents) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["content"] = o.Content
	toSerialize["contentformat"] = o.Contentformat
	toSerialize["files"] = o.Files
	toSerialize["footer"] = o.Footer
	toSerialize["title"] = o.Title
	return toSerialize, nil
}

func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContents) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"content",
		"contentformat",
		"files",
		"footer",
		"title",
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

	varCoreBlockGetDashboardBlocks200ResponseBlocksInnerContents := _CoreBlockGetDashboardBlocks200ResponseBlocksInnerContents{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreBlockGetDashboardBlocks200ResponseBlocksInnerContents)

	if err != nil {
		return err
	}

	*o = CoreBlockGetDashboardBlocks200ResponseBlocksInnerContents(varCoreBlockGetDashboardBlocks200ResponseBlocksInnerContents)

	return err
}

type NullableCoreBlockGetDashboardBlocks200ResponseBlocksInnerContents struct {
	value *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContents
	isSet bool
}

func (v NullableCoreBlockGetDashboardBlocks200ResponseBlocksInnerContents) Get() *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContents {
	return v.value
}

func (v *NullableCoreBlockGetDashboardBlocks200ResponseBlocksInnerContents) Set(val *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContents) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreBlockGetDashboardBlocks200ResponseBlocksInnerContents) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreBlockGetDashboardBlocks200ResponseBlocksInnerContents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreBlockGetDashboardBlocks200ResponseBlocksInnerContents(val *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContents) *NullableCoreBlockGetDashboardBlocks200ResponseBlocksInnerContents {
	return &NullableCoreBlockGetDashboardBlocks200ResponseBlocksInnerContents{value: val, isSet: true}
}

func (v NullableCoreBlockGetDashboardBlocks200ResponseBlocksInnerContents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreBlockGetDashboardBlocks200ResponseBlocksInnerContents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

