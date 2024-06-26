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

// checks if the CoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner{}

// CoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner struct for CoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner
type CoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner struct {
	// id
	Id *int32 `json:"id,omitempty"`
	// name
	Name *string `json:"name,omitempty"`
	// shortname
	Shortname *string `json:"shortname,omitempty"`
	// type
	Type *string `json:"type,omitempty"`
}

// NewCoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner instantiates a new CoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner() *CoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner {
	this := CoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner{}
	var id int32 = null
	this.Id = &id
	var name string = "null"
	this.Name = &name
	var shortname string = "null"
	this.Shortname = &shortname
	var type_ string = "null"
	this.Type = &type_
	return &this
}

// NewCoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInnerWithDefaults instantiates a new CoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInnerWithDefaults() *CoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner {
	this := CoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner{}
	var id int32 = null
	this.Id = &id
	var name string = "null"
	this.Name = &name
	var shortname string = "null"
	this.Shortname = &shortname
	var type_ string = "null"
	this.Type = &type_
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner) SetName(v string) {
	o.Name = &v
}

// GetShortname returns the Shortname field value if set, zero value otherwise.
func (o *CoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner) GetShortname() string {
	if o == nil || IsNil(o.Shortname) {
		var ret string
		return ret
	}
	return *o.Shortname
}

// GetShortnameOk returns a tuple with the Shortname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner) GetShortnameOk() (*string, bool) {
	if o == nil || IsNil(o.Shortname) {
		return nil, false
	}
	return o.Shortname, true
}

// HasShortname returns a boolean if a field has been set.
func (o *CoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner) HasShortname() bool {
	if o != nil && !IsNil(o.Shortname) {
		return true
	}

	return false
}

// SetShortname gets a reference to the given string and assigns it to the Shortname field.
func (o *CoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner) SetShortname(v string) {
	o.Shortname = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner) SetType(v string) {
	o.Type = &v
}

func (o CoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Shortname) {
		toSerialize["shortname"] = o.Shortname
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableCoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner struct {
	value *CoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner
	isSet bool
}

func (v NullableCoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner) Get() *CoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner {
	return v.value
}

func (v *NullableCoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner) Set(val *CoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner(val *CoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner) *NullableCoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner {
	return &NullableCoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner{value: val, isSet: true}
}

func (v NullableCoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCustomfieldReloadTemplate200ResponseCategoriesInnerFieldsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


