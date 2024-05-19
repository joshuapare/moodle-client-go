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

// checks if the CoreTagGetTagCollections200ResponseCollectionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreTagGetTagCollections200ResponseCollectionsInner{}

// CoreTagGetTagCollections200ResponseCollectionsInner struct for CoreTagGetTagCollections200ResponseCollectionsInner
type CoreTagGetTagCollections200ResponseCollectionsInner struct {
	// Component the collection is related to.
	Component *string `json:"component,omitempty"`
	// Custom URL for the tag page instead of /tag/index.php.
	Customurl *string `json:"customurl,omitempty"`
	// Collection id.
	Id *int32 `json:"id,omitempty"`
	// Whether is the default collection.
	Isdefault *bool `json:"isdefault,omitempty"`
	// Collection name.
	Name *string `json:"name,omitempty"`
	// Whether the tag collection is searchable.
	Searchable *bool `json:"searchable,omitempty"`
	// Collection ordering in the list.
	Sortorder *int32 `json:"sortorder,omitempty"`
}

// NewCoreTagGetTagCollections200ResponseCollectionsInner instantiates a new CoreTagGetTagCollections200ResponseCollectionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreTagGetTagCollections200ResponseCollectionsInner() *CoreTagGetTagCollections200ResponseCollectionsInner {
	this := CoreTagGetTagCollections200ResponseCollectionsInner{}
	var component string = "null"
	this.Component = &component
	var customurl string = "null"
	this.Customurl = &customurl
	var id int32 = null
	this.Id = &id
	var isdefault bool = false
	this.Isdefault = &isdefault
	var name string = "null"
	this.Name = &name
	var searchable bool = true
	this.Searchable = &searchable
	var sortorder int32 = null
	this.Sortorder = &sortorder
	return &this
}

// NewCoreTagGetTagCollections200ResponseCollectionsInnerWithDefaults instantiates a new CoreTagGetTagCollections200ResponseCollectionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreTagGetTagCollections200ResponseCollectionsInnerWithDefaults() *CoreTagGetTagCollections200ResponseCollectionsInner {
	this := CoreTagGetTagCollections200ResponseCollectionsInner{}
	var component string = "null"
	this.Component = &component
	var customurl string = "null"
	this.Customurl = &customurl
	var id int32 = null
	this.Id = &id
	var isdefault bool = false
	this.Isdefault = &isdefault
	var name string = "null"
	this.Name = &name
	var searchable bool = true
	this.Searchable = &searchable
	var sortorder int32 = null
	this.Sortorder = &sortorder
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *CoreTagGetTagCollections200ResponseCollectionsInner) GetComponent() string {
	if o == nil || IsNil(o.Component) {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreTagGetTagCollections200ResponseCollectionsInner) GetComponentOk() (*string, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *CoreTagGetTagCollections200ResponseCollectionsInner) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *CoreTagGetTagCollections200ResponseCollectionsInner) SetComponent(v string) {
	o.Component = &v
}

// GetCustomurl returns the Customurl field value if set, zero value otherwise.
func (o *CoreTagGetTagCollections200ResponseCollectionsInner) GetCustomurl() string {
	if o == nil || IsNil(o.Customurl) {
		var ret string
		return ret
	}
	return *o.Customurl
}

// GetCustomurlOk returns a tuple with the Customurl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreTagGetTagCollections200ResponseCollectionsInner) GetCustomurlOk() (*string, bool) {
	if o == nil || IsNil(o.Customurl) {
		return nil, false
	}
	return o.Customurl, true
}

// HasCustomurl returns a boolean if a field has been set.
func (o *CoreTagGetTagCollections200ResponseCollectionsInner) HasCustomurl() bool {
	if o != nil && !IsNil(o.Customurl) {
		return true
	}

	return false
}

// SetCustomurl gets a reference to the given string and assigns it to the Customurl field.
func (o *CoreTagGetTagCollections200ResponseCollectionsInner) SetCustomurl(v string) {
	o.Customurl = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CoreTagGetTagCollections200ResponseCollectionsInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreTagGetTagCollections200ResponseCollectionsInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CoreTagGetTagCollections200ResponseCollectionsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CoreTagGetTagCollections200ResponseCollectionsInner) SetId(v int32) {
	o.Id = &v
}

// GetIsdefault returns the Isdefault field value if set, zero value otherwise.
func (o *CoreTagGetTagCollections200ResponseCollectionsInner) GetIsdefault() bool {
	if o == nil || IsNil(o.Isdefault) {
		var ret bool
		return ret
	}
	return *o.Isdefault
}

// GetIsdefaultOk returns a tuple with the Isdefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreTagGetTagCollections200ResponseCollectionsInner) GetIsdefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Isdefault) {
		return nil, false
	}
	return o.Isdefault, true
}

// HasIsdefault returns a boolean if a field has been set.
func (o *CoreTagGetTagCollections200ResponseCollectionsInner) HasIsdefault() bool {
	if o != nil && !IsNil(o.Isdefault) {
		return true
	}

	return false
}

// SetIsdefault gets a reference to the given bool and assigns it to the Isdefault field.
func (o *CoreTagGetTagCollections200ResponseCollectionsInner) SetIsdefault(v bool) {
	o.Isdefault = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CoreTagGetTagCollections200ResponseCollectionsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreTagGetTagCollections200ResponseCollectionsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CoreTagGetTagCollections200ResponseCollectionsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CoreTagGetTagCollections200ResponseCollectionsInner) SetName(v string) {
	o.Name = &v
}

// GetSearchable returns the Searchable field value if set, zero value otherwise.
func (o *CoreTagGetTagCollections200ResponseCollectionsInner) GetSearchable() bool {
	if o == nil || IsNil(o.Searchable) {
		var ret bool
		return ret
	}
	return *o.Searchable
}

// GetSearchableOk returns a tuple with the Searchable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreTagGetTagCollections200ResponseCollectionsInner) GetSearchableOk() (*bool, bool) {
	if o == nil || IsNil(o.Searchable) {
		return nil, false
	}
	return o.Searchable, true
}

// HasSearchable returns a boolean if a field has been set.
func (o *CoreTagGetTagCollections200ResponseCollectionsInner) HasSearchable() bool {
	if o != nil && !IsNil(o.Searchable) {
		return true
	}

	return false
}

// SetSearchable gets a reference to the given bool and assigns it to the Searchable field.
func (o *CoreTagGetTagCollections200ResponseCollectionsInner) SetSearchable(v bool) {
	o.Searchable = &v
}

// GetSortorder returns the Sortorder field value if set, zero value otherwise.
func (o *CoreTagGetTagCollections200ResponseCollectionsInner) GetSortorder() int32 {
	if o == nil || IsNil(o.Sortorder) {
		var ret int32
		return ret
	}
	return *o.Sortorder
}

// GetSortorderOk returns a tuple with the Sortorder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreTagGetTagCollections200ResponseCollectionsInner) GetSortorderOk() (*int32, bool) {
	if o == nil || IsNil(o.Sortorder) {
		return nil, false
	}
	return o.Sortorder, true
}

// HasSortorder returns a boolean if a field has been set.
func (o *CoreTagGetTagCollections200ResponseCollectionsInner) HasSortorder() bool {
	if o != nil && !IsNil(o.Sortorder) {
		return true
	}

	return false
}

// SetSortorder gets a reference to the given int32 and assigns it to the Sortorder field.
func (o *CoreTagGetTagCollections200ResponseCollectionsInner) SetSortorder(v int32) {
	o.Sortorder = &v
}

func (o CoreTagGetTagCollections200ResponseCollectionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreTagGetTagCollections200ResponseCollectionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Component) {
		toSerialize["component"] = o.Component
	}
	if !IsNil(o.Customurl) {
		toSerialize["customurl"] = o.Customurl
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Isdefault) {
		toSerialize["isdefault"] = o.Isdefault
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Searchable) {
		toSerialize["searchable"] = o.Searchable
	}
	if !IsNil(o.Sortorder) {
		toSerialize["sortorder"] = o.Sortorder
	}
	return toSerialize, nil
}

type NullableCoreTagGetTagCollections200ResponseCollectionsInner struct {
	value *CoreTagGetTagCollections200ResponseCollectionsInner
	isSet bool
}

func (v NullableCoreTagGetTagCollections200ResponseCollectionsInner) Get() *CoreTagGetTagCollections200ResponseCollectionsInner {
	return v.value
}

func (v *NullableCoreTagGetTagCollections200ResponseCollectionsInner) Set(val *CoreTagGetTagCollections200ResponseCollectionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreTagGetTagCollections200ResponseCollectionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreTagGetTagCollections200ResponseCollectionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreTagGetTagCollections200ResponseCollectionsInner(val *CoreTagGetTagCollections200ResponseCollectionsInner) *NullableCoreTagGetTagCollections200ResponseCollectionsInner {
	return &NullableCoreTagGetTagCollections200ResponseCollectionsInner{value: val, isSet: true}
}

func (v NullableCoreTagGetTagCollections200ResponseCollectionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreTagGetTagCollections200ResponseCollectionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

