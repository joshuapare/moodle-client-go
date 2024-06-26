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

// checks if the CoreBlockFetchAddableBlocksRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreBlockFetchAddableBlocksRequest{}

// CoreBlockFetchAddableBlocksRequest struct for CoreBlockFetchAddableBlocksRequest
type CoreBlockFetchAddableBlocksRequest struct {
	// The context ID of the page.
	Pagecontextid int32 `json:"pagecontextid"`
	// Page hash
	Pagehash *string `json:"pagehash,omitempty"`
	// The layout of the page.
	Pagelayout string `json:"pagelayout"`
	// The type of the page.
	Pagetype string `json:"pagetype"`
	// The subpage identifier
	Subpage *string `json:"subpage,omitempty"`
}

type _CoreBlockFetchAddableBlocksRequest CoreBlockFetchAddableBlocksRequest

// NewCoreBlockFetchAddableBlocksRequest instantiates a new CoreBlockFetchAddableBlocksRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreBlockFetchAddableBlocksRequest(pagecontextid int32, pagelayout string, pagetype string) *CoreBlockFetchAddableBlocksRequest {
	this := CoreBlockFetchAddableBlocksRequest{}
	this.Pagecontextid = pagecontextid
	var pagehash string = ""
	this.Pagehash = &pagehash
	this.Pagelayout = pagelayout
	this.Pagetype = pagetype
	var subpage string = ""
	this.Subpage = &subpage
	return &this
}

// NewCoreBlockFetchAddableBlocksRequestWithDefaults instantiates a new CoreBlockFetchAddableBlocksRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreBlockFetchAddableBlocksRequestWithDefaults() *CoreBlockFetchAddableBlocksRequest {
	this := CoreBlockFetchAddableBlocksRequest{}
	var pagecontextid int32 = null
	this.Pagecontextid = pagecontextid
	var pagehash string = ""
	this.Pagehash = &pagehash
	var pagelayout string = "null"
	this.Pagelayout = pagelayout
	var pagetype string = "null"
	this.Pagetype = pagetype
	var subpage string = ""
	this.Subpage = &subpage
	return &this
}

// GetPagecontextid returns the Pagecontextid field value
func (o *CoreBlockFetchAddableBlocksRequest) GetPagecontextid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pagecontextid
}

// GetPagecontextidOk returns a tuple with the Pagecontextid field value
// and a boolean to check if the value has been set.
func (o *CoreBlockFetchAddableBlocksRequest) GetPagecontextidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagecontextid, true
}

// SetPagecontextid sets field value
func (o *CoreBlockFetchAddableBlocksRequest) SetPagecontextid(v int32) {
	o.Pagecontextid = v
}

// GetPagehash returns the Pagehash field value if set, zero value otherwise.
func (o *CoreBlockFetchAddableBlocksRequest) GetPagehash() string {
	if o == nil || IsNil(o.Pagehash) {
		var ret string
		return ret
	}
	return *o.Pagehash
}

// GetPagehashOk returns a tuple with the Pagehash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreBlockFetchAddableBlocksRequest) GetPagehashOk() (*string, bool) {
	if o == nil || IsNil(o.Pagehash) {
		return nil, false
	}
	return o.Pagehash, true
}

// HasPagehash returns a boolean if a field has been set.
func (o *CoreBlockFetchAddableBlocksRequest) HasPagehash() bool {
	if o != nil && !IsNil(o.Pagehash) {
		return true
	}

	return false
}

// SetPagehash gets a reference to the given string and assigns it to the Pagehash field.
func (o *CoreBlockFetchAddableBlocksRequest) SetPagehash(v string) {
	o.Pagehash = &v
}

// GetPagelayout returns the Pagelayout field value
func (o *CoreBlockFetchAddableBlocksRequest) GetPagelayout() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pagelayout
}

// GetPagelayoutOk returns a tuple with the Pagelayout field value
// and a boolean to check if the value has been set.
func (o *CoreBlockFetchAddableBlocksRequest) GetPagelayoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagelayout, true
}

// SetPagelayout sets field value
func (o *CoreBlockFetchAddableBlocksRequest) SetPagelayout(v string) {
	o.Pagelayout = v
}

// GetPagetype returns the Pagetype field value
func (o *CoreBlockFetchAddableBlocksRequest) GetPagetype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pagetype
}

// GetPagetypeOk returns a tuple with the Pagetype field value
// and a boolean to check if the value has been set.
func (o *CoreBlockFetchAddableBlocksRequest) GetPagetypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagetype, true
}

// SetPagetype sets field value
func (o *CoreBlockFetchAddableBlocksRequest) SetPagetype(v string) {
	o.Pagetype = v
}

// GetSubpage returns the Subpage field value if set, zero value otherwise.
func (o *CoreBlockFetchAddableBlocksRequest) GetSubpage() string {
	if o == nil || IsNil(o.Subpage) {
		var ret string
		return ret
	}
	return *o.Subpage
}

// GetSubpageOk returns a tuple with the Subpage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreBlockFetchAddableBlocksRequest) GetSubpageOk() (*string, bool) {
	if o == nil || IsNil(o.Subpage) {
		return nil, false
	}
	return o.Subpage, true
}

// HasSubpage returns a boolean if a field has been set.
func (o *CoreBlockFetchAddableBlocksRequest) HasSubpage() bool {
	if o != nil && !IsNil(o.Subpage) {
		return true
	}

	return false
}

// SetSubpage gets a reference to the given string and assigns it to the Subpage field.
func (o *CoreBlockFetchAddableBlocksRequest) SetSubpage(v string) {
	o.Subpage = &v
}

func (o CoreBlockFetchAddableBlocksRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreBlockFetchAddableBlocksRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pagecontextid"] = o.Pagecontextid
	if !IsNil(o.Pagehash) {
		toSerialize["pagehash"] = o.Pagehash
	}
	toSerialize["pagelayout"] = o.Pagelayout
	toSerialize["pagetype"] = o.Pagetype
	if !IsNil(o.Subpage) {
		toSerialize["subpage"] = o.Subpage
	}
	return toSerialize, nil
}

func (o *CoreBlockFetchAddableBlocksRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pagecontextid",
		"pagelayout",
		"pagetype",
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

	varCoreBlockFetchAddableBlocksRequest := _CoreBlockFetchAddableBlocksRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreBlockFetchAddableBlocksRequest)

	if err != nil {
		return err
	}

	*o = CoreBlockFetchAddableBlocksRequest(varCoreBlockFetchAddableBlocksRequest)

	return err
}

type NullableCoreBlockFetchAddableBlocksRequest struct {
	value *CoreBlockFetchAddableBlocksRequest
	isSet bool
}

func (v NullableCoreBlockFetchAddableBlocksRequest) Get() *CoreBlockFetchAddableBlocksRequest {
	return v.value
}

func (v *NullableCoreBlockFetchAddableBlocksRequest) Set(val *CoreBlockFetchAddableBlocksRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreBlockFetchAddableBlocksRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreBlockFetchAddableBlocksRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreBlockFetchAddableBlocksRequest(val *CoreBlockFetchAddableBlocksRequest) *NullableCoreBlockFetchAddableBlocksRequest {
	return &NullableCoreBlockFetchAddableBlocksRequest{value: val, isSet: true}
}

func (v NullableCoreBlockFetchAddableBlocksRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreBlockFetchAddableBlocksRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


