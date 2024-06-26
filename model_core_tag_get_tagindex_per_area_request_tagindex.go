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

// checks if the CoreTagGetTagindexPerAreaRequestTagindex type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreTagGetTagindexPerAreaRequestTagindex{}

// CoreTagGetTagindexPerAreaRequestTagindex struct for CoreTagGetTagindexPerAreaRequestTagindex
type CoreTagGetTagindexPerAreaRequestTagindex struct {
	// context id where to search for items
	Ctx *int32 `json:"ctx,omitempty"`
	// exlusive mode for this tag area
	Excl *bool `json:"excl,omitempty"`
	// context id where the link was displayed
	From *int32 `json:"from,omitempty"`
	// tag id
	Id *int32 `json:"id,omitempty"`
	// page number (0-based)
	Page *int32 `json:"page,omitempty"`
	// search in the context recursive
	Rec *int32 `json:"rec,omitempty"`
	// tag area id
	Ta *int32 `json:"ta,omitempty"`
	// tag name
	Tag *string `json:"tag,omitempty"`
	// tag collection id
	Tc *int32 `json:"tc,omitempty"`
}

// NewCoreTagGetTagindexPerAreaRequestTagindex instantiates a new CoreTagGetTagindexPerAreaRequestTagindex object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreTagGetTagindexPerAreaRequestTagindex() *CoreTagGetTagindexPerAreaRequestTagindex {
	this := CoreTagGetTagindexPerAreaRequestTagindex{}
	var ctx int32 = 0
	this.Ctx = &ctx
	var excl bool = 0
	this.Excl = &excl
	var from int32 = 0
	this.From = &from
	var id int32 = 0
	this.Id = &id
	var page int32 = 0
	this.Page = &page
	var rec int32 = 1
	this.Rec = &rec
	var ta int32 = 0
	this.Ta = &ta
	var tag string = ""
	this.Tag = &tag
	var tc int32 = 0
	this.Tc = &tc
	return &this
}

// NewCoreTagGetTagindexPerAreaRequestTagindexWithDefaults instantiates a new CoreTagGetTagindexPerAreaRequestTagindex object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreTagGetTagindexPerAreaRequestTagindexWithDefaults() *CoreTagGetTagindexPerAreaRequestTagindex {
	this := CoreTagGetTagindexPerAreaRequestTagindex{}
	var ctx int32 = 0
	this.Ctx = &ctx
	var excl bool = 0
	this.Excl = &excl
	var from int32 = 0
	this.From = &from
	var id int32 = 0
	this.Id = &id
	var page int32 = 0
	this.Page = &page
	var rec int32 = 1
	this.Rec = &rec
	var ta int32 = 0
	this.Ta = &ta
	var tag string = ""
	this.Tag = &tag
	var tc int32 = 0
	this.Tc = &tc
	return &this
}

// GetCtx returns the Ctx field value if set, zero value otherwise.
func (o *CoreTagGetTagindexPerAreaRequestTagindex) GetCtx() int32 {
	if o == nil || IsNil(o.Ctx) {
		var ret int32
		return ret
	}
	return *o.Ctx
}

// GetCtxOk returns a tuple with the Ctx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreTagGetTagindexPerAreaRequestTagindex) GetCtxOk() (*int32, bool) {
	if o == nil || IsNil(o.Ctx) {
		return nil, false
	}
	return o.Ctx, true
}

// HasCtx returns a boolean if a field has been set.
func (o *CoreTagGetTagindexPerAreaRequestTagindex) HasCtx() bool {
	if o != nil && !IsNil(o.Ctx) {
		return true
	}

	return false
}

// SetCtx gets a reference to the given int32 and assigns it to the Ctx field.
func (o *CoreTagGetTagindexPerAreaRequestTagindex) SetCtx(v int32) {
	o.Ctx = &v
}

// GetExcl returns the Excl field value if set, zero value otherwise.
func (o *CoreTagGetTagindexPerAreaRequestTagindex) GetExcl() bool {
	if o == nil || IsNil(o.Excl) {
		var ret bool
		return ret
	}
	return *o.Excl
}

// GetExclOk returns a tuple with the Excl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreTagGetTagindexPerAreaRequestTagindex) GetExclOk() (*bool, bool) {
	if o == nil || IsNil(o.Excl) {
		return nil, false
	}
	return o.Excl, true
}

// HasExcl returns a boolean if a field has been set.
func (o *CoreTagGetTagindexPerAreaRequestTagindex) HasExcl() bool {
	if o != nil && !IsNil(o.Excl) {
		return true
	}

	return false
}

// SetExcl gets a reference to the given bool and assigns it to the Excl field.
func (o *CoreTagGetTagindexPerAreaRequestTagindex) SetExcl(v bool) {
	o.Excl = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *CoreTagGetTagindexPerAreaRequestTagindex) GetFrom() int32 {
	if o == nil || IsNil(o.From) {
		var ret int32
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreTagGetTagindexPerAreaRequestTagindex) GetFromOk() (*int32, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *CoreTagGetTagindexPerAreaRequestTagindex) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given int32 and assigns it to the From field.
func (o *CoreTagGetTagindexPerAreaRequestTagindex) SetFrom(v int32) {
	o.From = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CoreTagGetTagindexPerAreaRequestTagindex) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreTagGetTagindexPerAreaRequestTagindex) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CoreTagGetTagindexPerAreaRequestTagindex) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CoreTagGetTagindexPerAreaRequestTagindex) SetId(v int32) {
	o.Id = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *CoreTagGetTagindexPerAreaRequestTagindex) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreTagGetTagindexPerAreaRequestTagindex) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *CoreTagGetTagindexPerAreaRequestTagindex) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *CoreTagGetTagindexPerAreaRequestTagindex) SetPage(v int32) {
	o.Page = &v
}

// GetRec returns the Rec field value if set, zero value otherwise.
func (o *CoreTagGetTagindexPerAreaRequestTagindex) GetRec() int32 {
	if o == nil || IsNil(o.Rec) {
		var ret int32
		return ret
	}
	return *o.Rec
}

// GetRecOk returns a tuple with the Rec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreTagGetTagindexPerAreaRequestTagindex) GetRecOk() (*int32, bool) {
	if o == nil || IsNil(o.Rec) {
		return nil, false
	}
	return o.Rec, true
}

// HasRec returns a boolean if a field has been set.
func (o *CoreTagGetTagindexPerAreaRequestTagindex) HasRec() bool {
	if o != nil && !IsNil(o.Rec) {
		return true
	}

	return false
}

// SetRec gets a reference to the given int32 and assigns it to the Rec field.
func (o *CoreTagGetTagindexPerAreaRequestTagindex) SetRec(v int32) {
	o.Rec = &v
}

// GetTa returns the Ta field value if set, zero value otherwise.
func (o *CoreTagGetTagindexPerAreaRequestTagindex) GetTa() int32 {
	if o == nil || IsNil(o.Ta) {
		var ret int32
		return ret
	}
	return *o.Ta
}

// GetTaOk returns a tuple with the Ta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreTagGetTagindexPerAreaRequestTagindex) GetTaOk() (*int32, bool) {
	if o == nil || IsNil(o.Ta) {
		return nil, false
	}
	return o.Ta, true
}

// HasTa returns a boolean if a field has been set.
func (o *CoreTagGetTagindexPerAreaRequestTagindex) HasTa() bool {
	if o != nil && !IsNil(o.Ta) {
		return true
	}

	return false
}

// SetTa gets a reference to the given int32 and assigns it to the Ta field.
func (o *CoreTagGetTagindexPerAreaRequestTagindex) SetTa(v int32) {
	o.Ta = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *CoreTagGetTagindexPerAreaRequestTagindex) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreTagGetTagindexPerAreaRequestTagindex) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *CoreTagGetTagindexPerAreaRequestTagindex) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *CoreTagGetTagindexPerAreaRequestTagindex) SetTag(v string) {
	o.Tag = &v
}

// GetTc returns the Tc field value if set, zero value otherwise.
func (o *CoreTagGetTagindexPerAreaRequestTagindex) GetTc() int32 {
	if o == nil || IsNil(o.Tc) {
		var ret int32
		return ret
	}
	return *o.Tc
}

// GetTcOk returns a tuple with the Tc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreTagGetTagindexPerAreaRequestTagindex) GetTcOk() (*int32, bool) {
	if o == nil || IsNil(o.Tc) {
		return nil, false
	}
	return o.Tc, true
}

// HasTc returns a boolean if a field has been set.
func (o *CoreTagGetTagindexPerAreaRequestTagindex) HasTc() bool {
	if o != nil && !IsNil(o.Tc) {
		return true
	}

	return false
}

// SetTc gets a reference to the given int32 and assigns it to the Tc field.
func (o *CoreTagGetTagindexPerAreaRequestTagindex) SetTc(v int32) {
	o.Tc = &v
}

func (o CoreTagGetTagindexPerAreaRequestTagindex) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreTagGetTagindexPerAreaRequestTagindex) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ctx) {
		toSerialize["ctx"] = o.Ctx
	}
	if !IsNil(o.Excl) {
		toSerialize["excl"] = o.Excl
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.Rec) {
		toSerialize["rec"] = o.Rec
	}
	if !IsNil(o.Ta) {
		toSerialize["ta"] = o.Ta
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.Tc) {
		toSerialize["tc"] = o.Tc
	}
	return toSerialize, nil
}

type NullableCoreTagGetTagindexPerAreaRequestTagindex struct {
	value *CoreTagGetTagindexPerAreaRequestTagindex
	isSet bool
}

func (v NullableCoreTagGetTagindexPerAreaRequestTagindex) Get() *CoreTagGetTagindexPerAreaRequestTagindex {
	return v.value
}

func (v *NullableCoreTagGetTagindexPerAreaRequestTagindex) Set(val *CoreTagGetTagindexPerAreaRequestTagindex) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreTagGetTagindexPerAreaRequestTagindex) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreTagGetTagindexPerAreaRequestTagindex) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreTagGetTagindexPerAreaRequestTagindex(val *CoreTagGetTagindexPerAreaRequestTagindex) *NullableCoreTagGetTagindexPerAreaRequestTagindex {
	return &NullableCoreTagGetTagindexPerAreaRequestTagindex{value: val, isSet: true}
}

func (v NullableCoreTagGetTagindexPerAreaRequestTagindex) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreTagGetTagindexPerAreaRequestTagindex) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


