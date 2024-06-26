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

// checks if the ModLessonGetPages200ResponsePagesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModLessonGetPages200ResponsePagesInner{}

// ModLessonGetPages200ResponsePagesInner The lesson pages
type ModLessonGetPages200ResponsePagesInner struct {
	Answerids []map[string]interface{} `json:"answerids,omitempty"`
	// The total number of files attached to the page
	Filescount *int32 `json:"filescount,omitempty"`
	// The total size of the files
	Filessizetotal *int32 `json:"filessizetotal,omitempty"`
	Jumps []map[string]interface{} `json:"jumps,omitempty"`
	Page *ModLessonGetPages200ResponsePagesInnerPage `json:"page,omitempty"`
}

// NewModLessonGetPages200ResponsePagesInner instantiates a new ModLessonGetPages200ResponsePagesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModLessonGetPages200ResponsePagesInner() *ModLessonGetPages200ResponsePagesInner {
	this := ModLessonGetPages200ResponsePagesInner{}
	var filescount int32 = null
	this.Filescount = &filescount
	var filessizetotal int32 = null
	this.Filessizetotal = &filessizetotal
	return &this
}

// NewModLessonGetPages200ResponsePagesInnerWithDefaults instantiates a new ModLessonGetPages200ResponsePagesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModLessonGetPages200ResponsePagesInnerWithDefaults() *ModLessonGetPages200ResponsePagesInner {
	this := ModLessonGetPages200ResponsePagesInner{}
	var filescount int32 = null
	this.Filescount = &filescount
	var filessizetotal int32 = null
	this.Filessizetotal = &filessizetotal
	return &this
}

// GetAnswerids returns the Answerids field value if set, zero value otherwise.
func (o *ModLessonGetPages200ResponsePagesInner) GetAnswerids() []map[string]interface{} {
	if o == nil || IsNil(o.Answerids) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Answerids
}

// GetAnsweridsOk returns a tuple with the Answerids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModLessonGetPages200ResponsePagesInner) GetAnsweridsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Answerids) {
		return nil, false
	}
	return o.Answerids, true
}

// HasAnswerids returns a boolean if a field has been set.
func (o *ModLessonGetPages200ResponsePagesInner) HasAnswerids() bool {
	if o != nil && !IsNil(o.Answerids) {
		return true
	}

	return false
}

// SetAnswerids gets a reference to the given []map[string]interface{} and assigns it to the Answerids field.
func (o *ModLessonGetPages200ResponsePagesInner) SetAnswerids(v []map[string]interface{}) {
	o.Answerids = v
}

// GetFilescount returns the Filescount field value if set, zero value otherwise.
func (o *ModLessonGetPages200ResponsePagesInner) GetFilescount() int32 {
	if o == nil || IsNil(o.Filescount) {
		var ret int32
		return ret
	}
	return *o.Filescount
}

// GetFilescountOk returns a tuple with the Filescount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModLessonGetPages200ResponsePagesInner) GetFilescountOk() (*int32, bool) {
	if o == nil || IsNil(o.Filescount) {
		return nil, false
	}
	return o.Filescount, true
}

// HasFilescount returns a boolean if a field has been set.
func (o *ModLessonGetPages200ResponsePagesInner) HasFilescount() bool {
	if o != nil && !IsNil(o.Filescount) {
		return true
	}

	return false
}

// SetFilescount gets a reference to the given int32 and assigns it to the Filescount field.
func (o *ModLessonGetPages200ResponsePagesInner) SetFilescount(v int32) {
	o.Filescount = &v
}

// GetFilessizetotal returns the Filessizetotal field value if set, zero value otherwise.
func (o *ModLessonGetPages200ResponsePagesInner) GetFilessizetotal() int32 {
	if o == nil || IsNil(o.Filessizetotal) {
		var ret int32
		return ret
	}
	return *o.Filessizetotal
}

// GetFilessizetotalOk returns a tuple with the Filessizetotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModLessonGetPages200ResponsePagesInner) GetFilessizetotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Filessizetotal) {
		return nil, false
	}
	return o.Filessizetotal, true
}

// HasFilessizetotal returns a boolean if a field has been set.
func (o *ModLessonGetPages200ResponsePagesInner) HasFilessizetotal() bool {
	if o != nil && !IsNil(o.Filessizetotal) {
		return true
	}

	return false
}

// SetFilessizetotal gets a reference to the given int32 and assigns it to the Filessizetotal field.
func (o *ModLessonGetPages200ResponsePagesInner) SetFilessizetotal(v int32) {
	o.Filessizetotal = &v
}

// GetJumps returns the Jumps field value if set, zero value otherwise.
func (o *ModLessonGetPages200ResponsePagesInner) GetJumps() []map[string]interface{} {
	if o == nil || IsNil(o.Jumps) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Jumps
}

// GetJumpsOk returns a tuple with the Jumps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModLessonGetPages200ResponsePagesInner) GetJumpsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Jumps) {
		return nil, false
	}
	return o.Jumps, true
}

// HasJumps returns a boolean if a field has been set.
func (o *ModLessonGetPages200ResponsePagesInner) HasJumps() bool {
	if o != nil && !IsNil(o.Jumps) {
		return true
	}

	return false
}

// SetJumps gets a reference to the given []map[string]interface{} and assigns it to the Jumps field.
func (o *ModLessonGetPages200ResponsePagesInner) SetJumps(v []map[string]interface{}) {
	o.Jumps = v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *ModLessonGetPages200ResponsePagesInner) GetPage() ModLessonGetPages200ResponsePagesInnerPage {
	if o == nil || IsNil(o.Page) {
		var ret ModLessonGetPages200ResponsePagesInnerPage
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModLessonGetPages200ResponsePagesInner) GetPageOk() (*ModLessonGetPages200ResponsePagesInnerPage, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *ModLessonGetPages200ResponsePagesInner) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given ModLessonGetPages200ResponsePagesInnerPage and assigns it to the Page field.
func (o *ModLessonGetPages200ResponsePagesInner) SetPage(v ModLessonGetPages200ResponsePagesInnerPage) {
	o.Page = &v
}

func (o ModLessonGetPages200ResponsePagesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModLessonGetPages200ResponsePagesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Answerids) {
		toSerialize["answerids"] = o.Answerids
	}
	if !IsNil(o.Filescount) {
		toSerialize["filescount"] = o.Filescount
	}
	if !IsNil(o.Filessizetotal) {
		toSerialize["filessizetotal"] = o.Filessizetotal
	}
	if !IsNil(o.Jumps) {
		toSerialize["jumps"] = o.Jumps
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	return toSerialize, nil
}

type NullableModLessonGetPages200ResponsePagesInner struct {
	value *ModLessonGetPages200ResponsePagesInner
	isSet bool
}

func (v NullableModLessonGetPages200ResponsePagesInner) Get() *ModLessonGetPages200ResponsePagesInner {
	return v.value
}

func (v *NullableModLessonGetPages200ResponsePagesInner) Set(val *ModLessonGetPages200ResponsePagesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableModLessonGetPages200ResponsePagesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableModLessonGetPages200ResponsePagesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModLessonGetPages200ResponsePagesInner(val *ModLessonGetPages200ResponsePagesInner) *NullableModLessonGetPages200ResponsePagesInner {
	return &NullableModLessonGetPages200ResponsePagesInner{value: val, isSet: true}
}

func (v NullableModLessonGetPages200ResponsePagesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModLessonGetPages200ResponsePagesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


