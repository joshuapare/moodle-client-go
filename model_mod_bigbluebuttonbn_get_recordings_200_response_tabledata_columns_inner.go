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

// checks if the ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner{}

// ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner struct for ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner
type ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner struct {
	// Whether this column contains HTML
	AllowHTML *bool `json:"allowHTML,omitempty"`
	// Formatter name
	Formatter *string `json:"formatter,omitempty"`
	Key *string `json:"key,omitempty"`
	Label *string `json:"label,omitempty"`
	// Whether this column is sortable
	Sortable *bool `json:"sortable,omitempty"`
	// Column type
	Type *string `json:"type,omitempty"`
	Width *string `json:"width,omitempty"`
}

// NewModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner instantiates a new ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner() *ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner {
	this := ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner{}
	var allowHTML bool = false
	this.AllowHTML = &allowHTML
	var formatter string = "null"
	this.Formatter = &formatter
	var sortable bool = false
	this.Sortable = &sortable
	var type_ string = "null"
	this.Type = &type_
	return &this
}

// NewModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInnerWithDefaults instantiates a new ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInnerWithDefaults() *ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner {
	this := ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner{}
	var allowHTML bool = false
	this.AllowHTML = &allowHTML
	var formatter string = "null"
	this.Formatter = &formatter
	var sortable bool = false
	this.Sortable = &sortable
	var type_ string = "null"
	this.Type = &type_
	return &this
}

// GetAllowHTML returns the AllowHTML field value if set, zero value otherwise.
func (o *ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) GetAllowHTML() bool {
	if o == nil || IsNil(o.AllowHTML) {
		var ret bool
		return ret
	}
	return *o.AllowHTML
}

// GetAllowHTMLOk returns a tuple with the AllowHTML field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) GetAllowHTMLOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowHTML) {
		return nil, false
	}
	return o.AllowHTML, true
}

// HasAllowHTML returns a boolean if a field has been set.
func (o *ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) HasAllowHTML() bool {
	if o != nil && !IsNil(o.AllowHTML) {
		return true
	}

	return false
}

// SetAllowHTML gets a reference to the given bool and assigns it to the AllowHTML field.
func (o *ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) SetAllowHTML(v bool) {
	o.AllowHTML = &v
}

// GetFormatter returns the Formatter field value if set, zero value otherwise.
func (o *ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) GetFormatter() string {
	if o == nil || IsNil(o.Formatter) {
		var ret string
		return ret
	}
	return *o.Formatter
}

// GetFormatterOk returns a tuple with the Formatter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) GetFormatterOk() (*string, bool) {
	if o == nil || IsNil(o.Formatter) {
		return nil, false
	}
	return o.Formatter, true
}

// HasFormatter returns a boolean if a field has been set.
func (o *ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) HasFormatter() bool {
	if o != nil && !IsNil(o.Formatter) {
		return true
	}

	return false
}

// SetFormatter gets a reference to the given string and assigns it to the Formatter field.
func (o *ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) SetFormatter(v string) {
	o.Formatter = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) SetKey(v string) {
	o.Key = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) SetLabel(v string) {
	o.Label = &v
}

// GetSortable returns the Sortable field value if set, zero value otherwise.
func (o *ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) GetSortable() bool {
	if o == nil || IsNil(o.Sortable) {
		var ret bool
		return ret
	}
	return *o.Sortable
}

// GetSortableOk returns a tuple with the Sortable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) GetSortableOk() (*bool, bool) {
	if o == nil || IsNil(o.Sortable) {
		return nil, false
	}
	return o.Sortable, true
}

// HasSortable returns a boolean if a field has been set.
func (o *ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) HasSortable() bool {
	if o != nil && !IsNil(o.Sortable) {
		return true
	}

	return false
}

// SetSortable gets a reference to the given bool and assigns it to the Sortable field.
func (o *ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) SetSortable(v bool) {
	o.Sortable = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) SetType(v string) {
	o.Type = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) GetWidth() string {
	if o == nil || IsNil(o.Width) {
		var ret string
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) GetWidthOk() (*string, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given string and assigns it to the Width field.
func (o *ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) SetWidth(v string) {
	o.Width = &v
}

func (o ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowHTML) {
		toSerialize["allowHTML"] = o.AllowHTML
	}
	if !IsNil(o.Formatter) {
		toSerialize["formatter"] = o.Formatter
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Sortable) {
		toSerialize["sortable"] = o.Sortable
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	return toSerialize, nil
}

type NullableModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner struct {
	value *ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner
	isSet bool
}

func (v NullableModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) Get() *ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner {
	return v.value
}

func (v *NullableModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) Set(val *ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner(val *ModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) *NullableModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner {
	return &NullableModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner{value: val, isSet: true}
}

func (v NullableModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModBigbluebuttonbnGetRecordings200ResponseTabledataColumnsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

