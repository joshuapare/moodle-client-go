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

// checks if the CoreUpdateInplaceEditable200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreUpdateInplaceEditable200Response{}

// CoreUpdateInplaceEditable200Response struct for CoreUpdateInplaceEditable200Response
type CoreUpdateInplaceEditable200Response struct {
	// component responsible for the update
	Component *string `json:"component,omitempty"`
	// display value (may contain link or other html tags)
	Displayvalue string `json:"displayvalue"`
	// hint for editing element
	Edithint *string `json:"edithint,omitempty"`
	Editicon *CoreUpdateInplaceEditable200ResponseEditicon `json:"editicon,omitempty"`
	// label for editing element
	Editlabel *string `json:"editlabel,omitempty"`
	// identifier of the updated item
	Itemid *string `json:"itemid,omitempty"`
	// itemtype
	Itemtype *string `json:"itemtype,omitempty"`
	// Should everything be wrapped in the edit link or link displayed separately
	Linkeverything *int32 `json:"linkeverything,omitempty"`
	// options of the element, format depends on type
	Options *string `json:"options,omitempty"`
	// type of the element (text, toggle, select)
	Type *string `json:"type,omitempty"`
	// value of the item as it is stored
	Value *string `json:"value,omitempty"`
}

type _CoreUpdateInplaceEditable200Response CoreUpdateInplaceEditable200Response

// NewCoreUpdateInplaceEditable200Response instantiates a new CoreUpdateInplaceEditable200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreUpdateInplaceEditable200Response(displayvalue string) *CoreUpdateInplaceEditable200Response {
	this := CoreUpdateInplaceEditable200Response{}
	this.Displayvalue = displayvalue
	var edithint string = "null"
	this.Edithint = &edithint
	var editlabel string = "null"
	this.Editlabel = &editlabel
	var linkeverything int32 = null
	this.Linkeverything = &linkeverything
	var options string = "null"
	this.Options = &options
	var type_ string = "null"
	this.Type = &type_
	var value string = "null"
	this.Value = &value
	return &this
}

// NewCoreUpdateInplaceEditable200ResponseWithDefaults instantiates a new CoreUpdateInplaceEditable200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreUpdateInplaceEditable200ResponseWithDefaults() *CoreUpdateInplaceEditable200Response {
	this := CoreUpdateInplaceEditable200Response{}
	var displayvalue string = "null"
	this.Displayvalue = displayvalue
	var edithint string = "null"
	this.Edithint = &edithint
	var editlabel string = "null"
	this.Editlabel = &editlabel
	var linkeverything int32 = null
	this.Linkeverything = &linkeverything
	var options string = "null"
	this.Options = &options
	var type_ string = "null"
	this.Type = &type_
	var value string = "null"
	this.Value = &value
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *CoreUpdateInplaceEditable200Response) GetComponent() string {
	if o == nil || IsNil(o.Component) {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreUpdateInplaceEditable200Response) GetComponentOk() (*string, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *CoreUpdateInplaceEditable200Response) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *CoreUpdateInplaceEditable200Response) SetComponent(v string) {
	o.Component = &v
}

// GetDisplayvalue returns the Displayvalue field value
func (o *CoreUpdateInplaceEditable200Response) GetDisplayvalue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Displayvalue
}

// GetDisplayvalueOk returns a tuple with the Displayvalue field value
// and a boolean to check if the value has been set.
func (o *CoreUpdateInplaceEditable200Response) GetDisplayvalueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Displayvalue, true
}

// SetDisplayvalue sets field value
func (o *CoreUpdateInplaceEditable200Response) SetDisplayvalue(v string) {
	o.Displayvalue = v
}

// GetEdithint returns the Edithint field value if set, zero value otherwise.
func (o *CoreUpdateInplaceEditable200Response) GetEdithint() string {
	if o == nil || IsNil(o.Edithint) {
		var ret string
		return ret
	}
	return *o.Edithint
}

// GetEdithintOk returns a tuple with the Edithint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreUpdateInplaceEditable200Response) GetEdithintOk() (*string, bool) {
	if o == nil || IsNil(o.Edithint) {
		return nil, false
	}
	return o.Edithint, true
}

// HasEdithint returns a boolean if a field has been set.
func (o *CoreUpdateInplaceEditable200Response) HasEdithint() bool {
	if o != nil && !IsNil(o.Edithint) {
		return true
	}

	return false
}

// SetEdithint gets a reference to the given string and assigns it to the Edithint field.
func (o *CoreUpdateInplaceEditable200Response) SetEdithint(v string) {
	o.Edithint = &v
}

// GetEditicon returns the Editicon field value if set, zero value otherwise.
func (o *CoreUpdateInplaceEditable200Response) GetEditicon() CoreUpdateInplaceEditable200ResponseEditicon {
	if o == nil || IsNil(o.Editicon) {
		var ret CoreUpdateInplaceEditable200ResponseEditicon
		return ret
	}
	return *o.Editicon
}

// GetEditiconOk returns a tuple with the Editicon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreUpdateInplaceEditable200Response) GetEditiconOk() (*CoreUpdateInplaceEditable200ResponseEditicon, bool) {
	if o == nil || IsNil(o.Editicon) {
		return nil, false
	}
	return o.Editicon, true
}

// HasEditicon returns a boolean if a field has been set.
func (o *CoreUpdateInplaceEditable200Response) HasEditicon() bool {
	if o != nil && !IsNil(o.Editicon) {
		return true
	}

	return false
}

// SetEditicon gets a reference to the given CoreUpdateInplaceEditable200ResponseEditicon and assigns it to the Editicon field.
func (o *CoreUpdateInplaceEditable200Response) SetEditicon(v CoreUpdateInplaceEditable200ResponseEditicon) {
	o.Editicon = &v
}

// GetEditlabel returns the Editlabel field value if set, zero value otherwise.
func (o *CoreUpdateInplaceEditable200Response) GetEditlabel() string {
	if o == nil || IsNil(o.Editlabel) {
		var ret string
		return ret
	}
	return *o.Editlabel
}

// GetEditlabelOk returns a tuple with the Editlabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreUpdateInplaceEditable200Response) GetEditlabelOk() (*string, bool) {
	if o == nil || IsNil(o.Editlabel) {
		return nil, false
	}
	return o.Editlabel, true
}

// HasEditlabel returns a boolean if a field has been set.
func (o *CoreUpdateInplaceEditable200Response) HasEditlabel() bool {
	if o != nil && !IsNil(o.Editlabel) {
		return true
	}

	return false
}

// SetEditlabel gets a reference to the given string and assigns it to the Editlabel field.
func (o *CoreUpdateInplaceEditable200Response) SetEditlabel(v string) {
	o.Editlabel = &v
}

// GetItemid returns the Itemid field value if set, zero value otherwise.
func (o *CoreUpdateInplaceEditable200Response) GetItemid() string {
	if o == nil || IsNil(o.Itemid) {
		var ret string
		return ret
	}
	return *o.Itemid
}

// GetItemidOk returns a tuple with the Itemid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreUpdateInplaceEditable200Response) GetItemidOk() (*string, bool) {
	if o == nil || IsNil(o.Itemid) {
		return nil, false
	}
	return o.Itemid, true
}

// HasItemid returns a boolean if a field has been set.
func (o *CoreUpdateInplaceEditable200Response) HasItemid() bool {
	if o != nil && !IsNil(o.Itemid) {
		return true
	}

	return false
}

// SetItemid gets a reference to the given string and assigns it to the Itemid field.
func (o *CoreUpdateInplaceEditable200Response) SetItemid(v string) {
	o.Itemid = &v
}

// GetItemtype returns the Itemtype field value if set, zero value otherwise.
func (o *CoreUpdateInplaceEditable200Response) GetItemtype() string {
	if o == nil || IsNil(o.Itemtype) {
		var ret string
		return ret
	}
	return *o.Itemtype
}

// GetItemtypeOk returns a tuple with the Itemtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreUpdateInplaceEditable200Response) GetItemtypeOk() (*string, bool) {
	if o == nil || IsNil(o.Itemtype) {
		return nil, false
	}
	return o.Itemtype, true
}

// HasItemtype returns a boolean if a field has been set.
func (o *CoreUpdateInplaceEditable200Response) HasItemtype() bool {
	if o != nil && !IsNil(o.Itemtype) {
		return true
	}

	return false
}

// SetItemtype gets a reference to the given string and assigns it to the Itemtype field.
func (o *CoreUpdateInplaceEditable200Response) SetItemtype(v string) {
	o.Itemtype = &v
}

// GetLinkeverything returns the Linkeverything field value if set, zero value otherwise.
func (o *CoreUpdateInplaceEditable200Response) GetLinkeverything() int32 {
	if o == nil || IsNil(o.Linkeverything) {
		var ret int32
		return ret
	}
	return *o.Linkeverything
}

// GetLinkeverythingOk returns a tuple with the Linkeverything field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreUpdateInplaceEditable200Response) GetLinkeverythingOk() (*int32, bool) {
	if o == nil || IsNil(o.Linkeverything) {
		return nil, false
	}
	return o.Linkeverything, true
}

// HasLinkeverything returns a boolean if a field has been set.
func (o *CoreUpdateInplaceEditable200Response) HasLinkeverything() bool {
	if o != nil && !IsNil(o.Linkeverything) {
		return true
	}

	return false
}

// SetLinkeverything gets a reference to the given int32 and assigns it to the Linkeverything field.
func (o *CoreUpdateInplaceEditable200Response) SetLinkeverything(v int32) {
	o.Linkeverything = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *CoreUpdateInplaceEditable200Response) GetOptions() string {
	if o == nil || IsNil(o.Options) {
		var ret string
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreUpdateInplaceEditable200Response) GetOptionsOk() (*string, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *CoreUpdateInplaceEditable200Response) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given string and assigns it to the Options field.
func (o *CoreUpdateInplaceEditable200Response) SetOptions(v string) {
	o.Options = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CoreUpdateInplaceEditable200Response) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreUpdateInplaceEditable200Response) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CoreUpdateInplaceEditable200Response) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CoreUpdateInplaceEditable200Response) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CoreUpdateInplaceEditable200Response) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreUpdateInplaceEditable200Response) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CoreUpdateInplaceEditable200Response) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CoreUpdateInplaceEditable200Response) SetValue(v string) {
	o.Value = &v
}

func (o CoreUpdateInplaceEditable200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreUpdateInplaceEditable200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Component) {
		toSerialize["component"] = o.Component
	}
	toSerialize["displayvalue"] = o.Displayvalue
	if !IsNil(o.Edithint) {
		toSerialize["edithint"] = o.Edithint
	}
	if !IsNil(o.Editicon) {
		toSerialize["editicon"] = o.Editicon
	}
	if !IsNil(o.Editlabel) {
		toSerialize["editlabel"] = o.Editlabel
	}
	if !IsNil(o.Itemid) {
		toSerialize["itemid"] = o.Itemid
	}
	if !IsNil(o.Itemtype) {
		toSerialize["itemtype"] = o.Itemtype
	}
	if !IsNil(o.Linkeverything) {
		toSerialize["linkeverything"] = o.Linkeverything
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

func (o *CoreUpdateInplaceEditable200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"displayvalue",
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

	varCoreUpdateInplaceEditable200Response := _CoreUpdateInplaceEditable200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreUpdateInplaceEditable200Response)

	if err != nil {
		return err
	}

	*o = CoreUpdateInplaceEditable200Response(varCoreUpdateInplaceEditable200Response)

	return err
}

type NullableCoreUpdateInplaceEditable200Response struct {
	value *CoreUpdateInplaceEditable200Response
	isSet bool
}

func (v NullableCoreUpdateInplaceEditable200Response) Get() *CoreUpdateInplaceEditable200Response {
	return v.value
}

func (v *NullableCoreUpdateInplaceEditable200Response) Set(val *CoreUpdateInplaceEditable200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreUpdateInplaceEditable200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreUpdateInplaceEditable200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreUpdateInplaceEditable200Response(val *CoreUpdateInplaceEditable200Response) *NullableCoreUpdateInplaceEditable200Response {
	return &NullableCoreUpdateInplaceEditable200Response{value: val, isSet: true}
}

func (v NullableCoreUpdateInplaceEditable200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreUpdateInplaceEditable200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

