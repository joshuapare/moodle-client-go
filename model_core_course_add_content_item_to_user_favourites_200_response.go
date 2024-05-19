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

// checks if the CoreCourseAddContentItemToUserFavourites200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCourseAddContentItemToUserFavourites200Response{}

// CoreCourseAddContentItemToUserFavourites200Response struct for CoreCourseAddContentItemToUserFavourites200Response
type CoreCourseAddContentItemToUserFavourites200Response struct {
	// The archetype of the module exposing the content item
	Archetype string `json:"archetype"`
	// The name of the component exposing the content item
	Componentname string `json:"componentname"`
	// Has the user favourited the content item
	Favourite bool `json:"favourite"`
	// Html description / help for the content item
	Help string `json:"help"`
	// Html containing the icon for the content item
	Icon string `json:"icon"`
	// The id of the content item
	Id int32 `json:"id"`
	// If this item was pulled from the old callback and has no item id.
	Legacyitem bool `json:"legacyitem"`
	// The link to the content item creation page
	Link string `json:"link"`
	// Name of the content item
	Name string `json:"name"`
	// The purpose of the component exposing the content item
	Purpose string `json:"purpose"`
	// Has this item been recommended
	Recommended bool `json:"recommended"`
	// The string title of the content item, human readable
	Title string `json:"title"`
}

type _CoreCourseAddContentItemToUserFavourites200Response CoreCourseAddContentItemToUserFavourites200Response

// NewCoreCourseAddContentItemToUserFavourites200Response instantiates a new CoreCourseAddContentItemToUserFavourites200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCourseAddContentItemToUserFavourites200Response(archetype string, componentname string, favourite bool, help string, icon string, id int32, legacyitem bool, link string, name string, purpose string, recommended bool, title string) *CoreCourseAddContentItemToUserFavourites200Response {
	this := CoreCourseAddContentItemToUserFavourites200Response{}
	this.Archetype = archetype
	this.Componentname = componentname
	this.Favourite = favourite
	this.Help = help
	this.Icon = icon
	this.Id = id
	this.Legacyitem = legacyitem
	this.Link = link
	this.Name = name
	this.Purpose = purpose
	this.Recommended = recommended
	this.Title = title
	return &this
}

// NewCoreCourseAddContentItemToUserFavourites200ResponseWithDefaults instantiates a new CoreCourseAddContentItemToUserFavourites200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCourseAddContentItemToUserFavourites200ResponseWithDefaults() *CoreCourseAddContentItemToUserFavourites200Response {
	this := CoreCourseAddContentItemToUserFavourites200Response{}
	var archetype string = "null"
	this.Archetype = archetype
	var componentname string = "null"
	this.Componentname = componentname
	var favourite bool = null
	this.Favourite = favourite
	var help string = "null"
	this.Help = help
	var icon string = "null"
	this.Icon = icon
	var id int32 = null
	this.Id = id
	var legacyitem bool = null
	this.Legacyitem = legacyitem
	var link string = "null"
	this.Link = link
	var name string = "null"
	this.Name = name
	var purpose string = "null"
	this.Purpose = purpose
	var recommended bool = null
	this.Recommended = recommended
	var title string = "null"
	this.Title = title
	return &this
}

// GetArchetype returns the Archetype field value
func (o *CoreCourseAddContentItemToUserFavourites200Response) GetArchetype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Archetype
}

// GetArchetypeOk returns a tuple with the Archetype field value
// and a boolean to check if the value has been set.
func (o *CoreCourseAddContentItemToUserFavourites200Response) GetArchetypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Archetype, true
}

// SetArchetype sets field value
func (o *CoreCourseAddContentItemToUserFavourites200Response) SetArchetype(v string) {
	o.Archetype = v
}

// GetComponentname returns the Componentname field value
func (o *CoreCourseAddContentItemToUserFavourites200Response) GetComponentname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Componentname
}

// GetComponentnameOk returns a tuple with the Componentname field value
// and a boolean to check if the value has been set.
func (o *CoreCourseAddContentItemToUserFavourites200Response) GetComponentnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Componentname, true
}

// SetComponentname sets field value
func (o *CoreCourseAddContentItemToUserFavourites200Response) SetComponentname(v string) {
	o.Componentname = v
}

// GetFavourite returns the Favourite field value
func (o *CoreCourseAddContentItemToUserFavourites200Response) GetFavourite() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Favourite
}

// GetFavouriteOk returns a tuple with the Favourite field value
// and a boolean to check if the value has been set.
func (o *CoreCourseAddContentItemToUserFavourites200Response) GetFavouriteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Favourite, true
}

// SetFavourite sets field value
func (o *CoreCourseAddContentItemToUserFavourites200Response) SetFavourite(v bool) {
	o.Favourite = v
}

// GetHelp returns the Help field value
func (o *CoreCourseAddContentItemToUserFavourites200Response) GetHelp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Help
}

// GetHelpOk returns a tuple with the Help field value
// and a boolean to check if the value has been set.
func (o *CoreCourseAddContentItemToUserFavourites200Response) GetHelpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Help, true
}

// SetHelp sets field value
func (o *CoreCourseAddContentItemToUserFavourites200Response) SetHelp(v string) {
	o.Help = v
}

// GetIcon returns the Icon field value
func (o *CoreCourseAddContentItemToUserFavourites200Response) GetIcon() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Icon
}

// GetIconOk returns a tuple with the Icon field value
// and a boolean to check if the value has been set.
func (o *CoreCourseAddContentItemToUserFavourites200Response) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Icon, true
}

// SetIcon sets field value
func (o *CoreCourseAddContentItemToUserFavourites200Response) SetIcon(v string) {
	o.Icon = v
}

// GetId returns the Id field value
func (o *CoreCourseAddContentItemToUserFavourites200Response) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CoreCourseAddContentItemToUserFavourites200Response) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CoreCourseAddContentItemToUserFavourites200Response) SetId(v int32) {
	o.Id = v
}

// GetLegacyitem returns the Legacyitem field value
func (o *CoreCourseAddContentItemToUserFavourites200Response) GetLegacyitem() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Legacyitem
}

// GetLegacyitemOk returns a tuple with the Legacyitem field value
// and a boolean to check if the value has been set.
func (o *CoreCourseAddContentItemToUserFavourites200Response) GetLegacyitemOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Legacyitem, true
}

// SetLegacyitem sets field value
func (o *CoreCourseAddContentItemToUserFavourites200Response) SetLegacyitem(v bool) {
	o.Legacyitem = v
}

// GetLink returns the Link field value
func (o *CoreCourseAddContentItemToUserFavourites200Response) GetLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Link
}

// GetLinkOk returns a tuple with the Link field value
// and a boolean to check if the value has been set.
func (o *CoreCourseAddContentItemToUserFavourites200Response) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Link, true
}

// SetLink sets field value
func (o *CoreCourseAddContentItemToUserFavourites200Response) SetLink(v string) {
	o.Link = v
}

// GetName returns the Name field value
func (o *CoreCourseAddContentItemToUserFavourites200Response) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CoreCourseAddContentItemToUserFavourites200Response) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CoreCourseAddContentItemToUserFavourites200Response) SetName(v string) {
	o.Name = v
}

// GetPurpose returns the Purpose field value
func (o *CoreCourseAddContentItemToUserFavourites200Response) GetPurpose() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value
// and a boolean to check if the value has been set.
func (o *CoreCourseAddContentItemToUserFavourites200Response) GetPurposeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Purpose, true
}

// SetPurpose sets field value
func (o *CoreCourseAddContentItemToUserFavourites200Response) SetPurpose(v string) {
	o.Purpose = v
}

// GetRecommended returns the Recommended field value
func (o *CoreCourseAddContentItemToUserFavourites200Response) GetRecommended() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Recommended
}

// GetRecommendedOk returns a tuple with the Recommended field value
// and a boolean to check if the value has been set.
func (o *CoreCourseAddContentItemToUserFavourites200Response) GetRecommendedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recommended, true
}

// SetRecommended sets field value
func (o *CoreCourseAddContentItemToUserFavourites200Response) SetRecommended(v bool) {
	o.Recommended = v
}

// GetTitle returns the Title field value
func (o *CoreCourseAddContentItemToUserFavourites200Response) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *CoreCourseAddContentItemToUserFavourites200Response) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *CoreCourseAddContentItemToUserFavourites200Response) SetTitle(v string) {
	o.Title = v
}

func (o CoreCourseAddContentItemToUserFavourites200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCourseAddContentItemToUserFavourites200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["archetype"] = o.Archetype
	toSerialize["componentname"] = o.Componentname
	toSerialize["favourite"] = o.Favourite
	toSerialize["help"] = o.Help
	toSerialize["icon"] = o.Icon
	toSerialize["id"] = o.Id
	toSerialize["legacyitem"] = o.Legacyitem
	toSerialize["link"] = o.Link
	toSerialize["name"] = o.Name
	toSerialize["purpose"] = o.Purpose
	toSerialize["recommended"] = o.Recommended
	toSerialize["title"] = o.Title
	return toSerialize, nil
}

func (o *CoreCourseAddContentItemToUserFavourites200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"archetype",
		"componentname",
		"favourite",
		"help",
		"icon",
		"id",
		"legacyitem",
		"link",
		"name",
		"purpose",
		"recommended",
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

	varCoreCourseAddContentItemToUserFavourites200Response := _CoreCourseAddContentItemToUserFavourites200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreCourseAddContentItemToUserFavourites200Response)

	if err != nil {
		return err
	}

	*o = CoreCourseAddContentItemToUserFavourites200Response(varCoreCourseAddContentItemToUserFavourites200Response)

	return err
}

type NullableCoreCourseAddContentItemToUserFavourites200Response struct {
	value *CoreCourseAddContentItemToUserFavourites200Response
	isSet bool
}

func (v NullableCoreCourseAddContentItemToUserFavourites200Response) Get() *CoreCourseAddContentItemToUserFavourites200Response {
	return v.value
}

func (v *NullableCoreCourseAddContentItemToUserFavourites200Response) Set(val *CoreCourseAddContentItemToUserFavourites200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCourseAddContentItemToUserFavourites200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCourseAddContentItemToUserFavourites200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCourseAddContentItemToUserFavourites200Response(val *CoreCourseAddContentItemToUserFavourites200Response) *NullableCoreCourseAddContentItemToUserFavourites200Response {
	return &NullableCoreCourseAddContentItemToUserFavourites200Response{value: val, isSet: true}
}

func (v NullableCoreCourseAddContentItemToUserFavourites200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCourseAddContentItemToUserFavourites200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


