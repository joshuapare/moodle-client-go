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

// checks if the CoreTagGetTagindex200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreTagGetTagindex200Response{}

// CoreTagGetTagindex200Response struct for CoreTagGetTagindex200Response
type CoreTagGetTagindex200Response struct {
	// name of anchor
	Anchor *string `json:"anchor,omitempty"`
	// component
	Component string `json:"component"`
	// title
	Content string `json:"content"`
	// text for exclusive link
	Exclusivetext *string `json:"exclusivetext,omitempty"`
	// URL for exclusive link
	Exclusiveurl *string `json:"exclusiveurl,omitempty"`
	// whether the content is present
	Hascontent int32 `json:"hascontent"`
	// itemtype
	Itemtype string `json:"itemtype"`
	// URL for the next page
	Nextpageurl *string `json:"nextpageurl,omitempty"`
	// URL for the next page
	Prevpageurl *string `json:"prevpageurl,omitempty"`
	// tag area id
	Ta int32 `json:"ta"`
	// tag id
	Tagid int32 `json:"tagid"`
	// title
	Title string `json:"title"`
}

type _CoreTagGetTagindex200Response CoreTagGetTagindex200Response

// NewCoreTagGetTagindex200Response instantiates a new CoreTagGetTagindex200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreTagGetTagindex200Response(component string, content string, hascontent int32, itemtype string, ta int32, tagid int32, title string) *CoreTagGetTagindex200Response {
	this := CoreTagGetTagindex200Response{}
	var anchor string = "null"
	this.Anchor = &anchor
	this.Component = component
	this.Content = content
	var exclusivetext string = "null"
	this.Exclusivetext = &exclusivetext
	var exclusiveurl string = "null"
	this.Exclusiveurl = &exclusiveurl
	this.Hascontent = hascontent
	this.Itemtype = itemtype
	var nextpageurl string = "null"
	this.Nextpageurl = &nextpageurl
	this.Ta = ta
	this.Tagid = tagid
	this.Title = title
	return &this
}

// NewCoreTagGetTagindex200ResponseWithDefaults instantiates a new CoreTagGetTagindex200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreTagGetTagindex200ResponseWithDefaults() *CoreTagGetTagindex200Response {
	this := CoreTagGetTagindex200Response{}
	var anchor string = "null"
	this.Anchor = &anchor
	var content string = "null"
	this.Content = content
	var exclusivetext string = "null"
	this.Exclusivetext = &exclusivetext
	var exclusiveurl string = "null"
	this.Exclusiveurl = &exclusiveurl
	var hascontent int32 = null
	this.Hascontent = hascontent
	var itemtype string = "null"
	this.Itemtype = itemtype
	var nextpageurl string = "null"
	this.Nextpageurl = &nextpageurl
	var tagid int32 = null
	this.Tagid = tagid
	return &this
}

// GetAnchor returns the Anchor field value if set, zero value otherwise.
func (o *CoreTagGetTagindex200Response) GetAnchor() string {
	if o == nil || IsNil(o.Anchor) {
		var ret string
		return ret
	}
	return *o.Anchor
}

// GetAnchorOk returns a tuple with the Anchor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreTagGetTagindex200Response) GetAnchorOk() (*string, bool) {
	if o == nil || IsNil(o.Anchor) {
		return nil, false
	}
	return o.Anchor, true
}

// HasAnchor returns a boolean if a field has been set.
func (o *CoreTagGetTagindex200Response) HasAnchor() bool {
	if o != nil && !IsNil(o.Anchor) {
		return true
	}

	return false
}

// SetAnchor gets a reference to the given string and assigns it to the Anchor field.
func (o *CoreTagGetTagindex200Response) SetAnchor(v string) {
	o.Anchor = &v
}

// GetComponent returns the Component field value
func (o *CoreTagGetTagindex200Response) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *CoreTagGetTagindex200Response) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *CoreTagGetTagindex200Response) SetComponent(v string) {
	o.Component = v
}

// GetContent returns the Content field value
func (o *CoreTagGetTagindex200Response) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *CoreTagGetTagindex200Response) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *CoreTagGetTagindex200Response) SetContent(v string) {
	o.Content = v
}

// GetExclusivetext returns the Exclusivetext field value if set, zero value otherwise.
func (o *CoreTagGetTagindex200Response) GetExclusivetext() string {
	if o == nil || IsNil(o.Exclusivetext) {
		var ret string
		return ret
	}
	return *o.Exclusivetext
}

// GetExclusivetextOk returns a tuple with the Exclusivetext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreTagGetTagindex200Response) GetExclusivetextOk() (*string, bool) {
	if o == nil || IsNil(o.Exclusivetext) {
		return nil, false
	}
	return o.Exclusivetext, true
}

// HasExclusivetext returns a boolean if a field has been set.
func (o *CoreTagGetTagindex200Response) HasExclusivetext() bool {
	if o != nil && !IsNil(o.Exclusivetext) {
		return true
	}

	return false
}

// SetExclusivetext gets a reference to the given string and assigns it to the Exclusivetext field.
func (o *CoreTagGetTagindex200Response) SetExclusivetext(v string) {
	o.Exclusivetext = &v
}

// GetExclusiveurl returns the Exclusiveurl field value if set, zero value otherwise.
func (o *CoreTagGetTagindex200Response) GetExclusiveurl() string {
	if o == nil || IsNil(o.Exclusiveurl) {
		var ret string
		return ret
	}
	return *o.Exclusiveurl
}

// GetExclusiveurlOk returns a tuple with the Exclusiveurl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreTagGetTagindex200Response) GetExclusiveurlOk() (*string, bool) {
	if o == nil || IsNil(o.Exclusiveurl) {
		return nil, false
	}
	return o.Exclusiveurl, true
}

// HasExclusiveurl returns a boolean if a field has been set.
func (o *CoreTagGetTagindex200Response) HasExclusiveurl() bool {
	if o != nil && !IsNil(o.Exclusiveurl) {
		return true
	}

	return false
}

// SetExclusiveurl gets a reference to the given string and assigns it to the Exclusiveurl field.
func (o *CoreTagGetTagindex200Response) SetExclusiveurl(v string) {
	o.Exclusiveurl = &v
}

// GetHascontent returns the Hascontent field value
func (o *CoreTagGetTagindex200Response) GetHascontent() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Hascontent
}

// GetHascontentOk returns a tuple with the Hascontent field value
// and a boolean to check if the value has been set.
func (o *CoreTagGetTagindex200Response) GetHascontentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hascontent, true
}

// SetHascontent sets field value
func (o *CoreTagGetTagindex200Response) SetHascontent(v int32) {
	o.Hascontent = v
}

// GetItemtype returns the Itemtype field value
func (o *CoreTagGetTagindex200Response) GetItemtype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Itemtype
}

// GetItemtypeOk returns a tuple with the Itemtype field value
// and a boolean to check if the value has been set.
func (o *CoreTagGetTagindex200Response) GetItemtypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Itemtype, true
}

// SetItemtype sets field value
func (o *CoreTagGetTagindex200Response) SetItemtype(v string) {
	o.Itemtype = v
}

// GetNextpageurl returns the Nextpageurl field value if set, zero value otherwise.
func (o *CoreTagGetTagindex200Response) GetNextpageurl() string {
	if o == nil || IsNil(o.Nextpageurl) {
		var ret string
		return ret
	}
	return *o.Nextpageurl
}

// GetNextpageurlOk returns a tuple with the Nextpageurl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreTagGetTagindex200Response) GetNextpageurlOk() (*string, bool) {
	if o == nil || IsNil(o.Nextpageurl) {
		return nil, false
	}
	return o.Nextpageurl, true
}

// HasNextpageurl returns a boolean if a field has been set.
func (o *CoreTagGetTagindex200Response) HasNextpageurl() bool {
	if o != nil && !IsNil(o.Nextpageurl) {
		return true
	}

	return false
}

// SetNextpageurl gets a reference to the given string and assigns it to the Nextpageurl field.
func (o *CoreTagGetTagindex200Response) SetNextpageurl(v string) {
	o.Nextpageurl = &v
}

// GetPrevpageurl returns the Prevpageurl field value if set, zero value otherwise.
func (o *CoreTagGetTagindex200Response) GetPrevpageurl() string {
	if o == nil || IsNil(o.Prevpageurl) {
		var ret string
		return ret
	}
	return *o.Prevpageurl
}

// GetPrevpageurlOk returns a tuple with the Prevpageurl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreTagGetTagindex200Response) GetPrevpageurlOk() (*string, bool) {
	if o == nil || IsNil(o.Prevpageurl) {
		return nil, false
	}
	return o.Prevpageurl, true
}

// HasPrevpageurl returns a boolean if a field has been set.
func (o *CoreTagGetTagindex200Response) HasPrevpageurl() bool {
	if o != nil && !IsNil(o.Prevpageurl) {
		return true
	}

	return false
}

// SetPrevpageurl gets a reference to the given string and assigns it to the Prevpageurl field.
func (o *CoreTagGetTagindex200Response) SetPrevpageurl(v string) {
	o.Prevpageurl = &v
}

// GetTa returns the Ta field value
func (o *CoreTagGetTagindex200Response) GetTa() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Ta
}

// GetTaOk returns a tuple with the Ta field value
// and a boolean to check if the value has been set.
func (o *CoreTagGetTagindex200Response) GetTaOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ta, true
}

// SetTa sets field value
func (o *CoreTagGetTagindex200Response) SetTa(v int32) {
	o.Ta = v
}

// GetTagid returns the Tagid field value
func (o *CoreTagGetTagindex200Response) GetTagid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Tagid
}

// GetTagidOk returns a tuple with the Tagid field value
// and a boolean to check if the value has been set.
func (o *CoreTagGetTagindex200Response) GetTagidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tagid, true
}

// SetTagid sets field value
func (o *CoreTagGetTagindex200Response) SetTagid(v int32) {
	o.Tagid = v
}

// GetTitle returns the Title field value
func (o *CoreTagGetTagindex200Response) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *CoreTagGetTagindex200Response) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *CoreTagGetTagindex200Response) SetTitle(v string) {
	o.Title = v
}

func (o CoreTagGetTagindex200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreTagGetTagindex200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Anchor) {
		toSerialize["anchor"] = o.Anchor
	}
	toSerialize["component"] = o.Component
	toSerialize["content"] = o.Content
	if !IsNil(o.Exclusivetext) {
		toSerialize["exclusivetext"] = o.Exclusivetext
	}
	if !IsNil(o.Exclusiveurl) {
		toSerialize["exclusiveurl"] = o.Exclusiveurl
	}
	toSerialize["hascontent"] = o.Hascontent
	toSerialize["itemtype"] = o.Itemtype
	if !IsNil(o.Nextpageurl) {
		toSerialize["nextpageurl"] = o.Nextpageurl
	}
	if !IsNil(o.Prevpageurl) {
		toSerialize["prevpageurl"] = o.Prevpageurl
	}
	toSerialize["ta"] = o.Ta
	toSerialize["tagid"] = o.Tagid
	toSerialize["title"] = o.Title
	return toSerialize, nil
}

func (o *CoreTagGetTagindex200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"component",
		"content",
		"hascontent",
		"itemtype",
		"ta",
		"tagid",
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

	varCoreTagGetTagindex200Response := _CoreTagGetTagindex200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreTagGetTagindex200Response)

	if err != nil {
		return err
	}

	*o = CoreTagGetTagindex200Response(varCoreTagGetTagindex200Response)

	return err
}

type NullableCoreTagGetTagindex200Response struct {
	value *CoreTagGetTagindex200Response
	isSet bool
}

func (v NullableCoreTagGetTagindex200Response) Get() *CoreTagGetTagindex200Response {
	return v.value
}

func (v *NullableCoreTagGetTagindex200Response) Set(val *CoreTagGetTagindex200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreTagGetTagindex200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreTagGetTagindex200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreTagGetTagindex200Response(val *CoreTagGetTagindex200Response) *NullableCoreTagGetTagindex200Response {
	return &NullableCoreTagGetTagindex200Response{value: val, isSet: true}
}

func (v NullableCoreTagGetTagindex200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreTagGetTagindex200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

