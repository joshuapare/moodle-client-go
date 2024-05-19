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

// checks if the CoreCourseCreateCategoriesRequestCategoriesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCourseCreateCategoriesRequestCategoriesInner{}

// CoreCourseCreateCategoriesRequestCategoriesInner struct for CoreCourseCreateCategoriesRequestCategoriesInner
type CoreCourseCreateCategoriesRequestCategoriesInner struct {
	// the new category description
	Description *string `json:"description,omitempty"`
	// description format (1 = HTML, 0 = MOODLE, 2 = PLAIN, or 4 = MARKDOWN)
	Descriptionformat *int32 `json:"descriptionformat,omitempty"`
	// the new category idnumber
	Idnumber *string `json:"idnumber,omitempty"`
	// new category name
	Name *string `json:"name,omitempty"`
	// the parent category id inside which the new category will be created                                          - set to 0 for a root category
	Parent *int32 `json:"parent,omitempty"`
	// the new category theme. This option must be enabled on moodle
	Theme *string `json:"theme,omitempty"`
}

// NewCoreCourseCreateCategoriesRequestCategoriesInner instantiates a new CoreCourseCreateCategoriesRequestCategoriesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCourseCreateCategoriesRequestCategoriesInner() *CoreCourseCreateCategoriesRequestCategoriesInner {
	this := CoreCourseCreateCategoriesRequestCategoriesInner{}
	var description string = "null"
	this.Description = &description
	var descriptionformat int32 = 1
	this.Descriptionformat = &descriptionformat
	var idnumber string = "null"
	this.Idnumber = &idnumber
	var name string = "null"
	this.Name = &name
	var parent int32 = 0
	this.Parent = &parent
	var theme string = "null"
	this.Theme = &theme
	return &this
}

// NewCoreCourseCreateCategoriesRequestCategoriesInnerWithDefaults instantiates a new CoreCourseCreateCategoriesRequestCategoriesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCourseCreateCategoriesRequestCategoriesInnerWithDefaults() *CoreCourseCreateCategoriesRequestCategoriesInner {
	this := CoreCourseCreateCategoriesRequestCategoriesInner{}
	var description string = "null"
	this.Description = &description
	var descriptionformat int32 = 1
	this.Descriptionformat = &descriptionformat
	var idnumber string = "null"
	this.Idnumber = &idnumber
	var name string = "null"
	this.Name = &name
	var parent int32 = 0
	this.Parent = &parent
	var theme string = "null"
	this.Theme = &theme
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CoreCourseCreateCategoriesRequestCategoriesInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseCreateCategoriesRequestCategoriesInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CoreCourseCreateCategoriesRequestCategoriesInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CoreCourseCreateCategoriesRequestCategoriesInner) SetDescription(v string) {
	o.Description = &v
}

// GetDescriptionformat returns the Descriptionformat field value if set, zero value otherwise.
func (o *CoreCourseCreateCategoriesRequestCategoriesInner) GetDescriptionformat() int32 {
	if o == nil || IsNil(o.Descriptionformat) {
		var ret int32
		return ret
	}
	return *o.Descriptionformat
}

// GetDescriptionformatOk returns a tuple with the Descriptionformat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseCreateCategoriesRequestCategoriesInner) GetDescriptionformatOk() (*int32, bool) {
	if o == nil || IsNil(o.Descriptionformat) {
		return nil, false
	}
	return o.Descriptionformat, true
}

// HasDescriptionformat returns a boolean if a field has been set.
func (o *CoreCourseCreateCategoriesRequestCategoriesInner) HasDescriptionformat() bool {
	if o != nil && !IsNil(o.Descriptionformat) {
		return true
	}

	return false
}

// SetDescriptionformat gets a reference to the given int32 and assigns it to the Descriptionformat field.
func (o *CoreCourseCreateCategoriesRequestCategoriesInner) SetDescriptionformat(v int32) {
	o.Descriptionformat = &v
}

// GetIdnumber returns the Idnumber field value if set, zero value otherwise.
func (o *CoreCourseCreateCategoriesRequestCategoriesInner) GetIdnumber() string {
	if o == nil || IsNil(o.Idnumber) {
		var ret string
		return ret
	}
	return *o.Idnumber
}

// GetIdnumberOk returns a tuple with the Idnumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseCreateCategoriesRequestCategoriesInner) GetIdnumberOk() (*string, bool) {
	if o == nil || IsNil(o.Idnumber) {
		return nil, false
	}
	return o.Idnumber, true
}

// HasIdnumber returns a boolean if a field has been set.
func (o *CoreCourseCreateCategoriesRequestCategoriesInner) HasIdnumber() bool {
	if o != nil && !IsNil(o.Idnumber) {
		return true
	}

	return false
}

// SetIdnumber gets a reference to the given string and assigns it to the Idnumber field.
func (o *CoreCourseCreateCategoriesRequestCategoriesInner) SetIdnumber(v string) {
	o.Idnumber = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CoreCourseCreateCategoriesRequestCategoriesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseCreateCategoriesRequestCategoriesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CoreCourseCreateCategoriesRequestCategoriesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CoreCourseCreateCategoriesRequestCategoriesInner) SetName(v string) {
	o.Name = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *CoreCourseCreateCategoriesRequestCategoriesInner) GetParent() int32 {
	if o == nil || IsNil(o.Parent) {
		var ret int32
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseCreateCategoriesRequestCategoriesInner) GetParentOk() (*int32, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *CoreCourseCreateCategoriesRequestCategoriesInner) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given int32 and assigns it to the Parent field.
func (o *CoreCourseCreateCategoriesRequestCategoriesInner) SetParent(v int32) {
	o.Parent = &v
}

// GetTheme returns the Theme field value if set, zero value otherwise.
func (o *CoreCourseCreateCategoriesRequestCategoriesInner) GetTheme() string {
	if o == nil || IsNil(o.Theme) {
		var ret string
		return ret
	}
	return *o.Theme
}

// GetThemeOk returns a tuple with the Theme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseCreateCategoriesRequestCategoriesInner) GetThemeOk() (*string, bool) {
	if o == nil || IsNil(o.Theme) {
		return nil, false
	}
	return o.Theme, true
}

// HasTheme returns a boolean if a field has been set.
func (o *CoreCourseCreateCategoriesRequestCategoriesInner) HasTheme() bool {
	if o != nil && !IsNil(o.Theme) {
		return true
	}

	return false
}

// SetTheme gets a reference to the given string and assigns it to the Theme field.
func (o *CoreCourseCreateCategoriesRequestCategoriesInner) SetTheme(v string) {
	o.Theme = &v
}

func (o CoreCourseCreateCategoriesRequestCategoriesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCourseCreateCategoriesRequestCategoriesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Descriptionformat) {
		toSerialize["descriptionformat"] = o.Descriptionformat
	}
	if !IsNil(o.Idnumber) {
		toSerialize["idnumber"] = o.Idnumber
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !IsNil(o.Theme) {
		toSerialize["theme"] = o.Theme
	}
	return toSerialize, nil
}

type NullableCoreCourseCreateCategoriesRequestCategoriesInner struct {
	value *CoreCourseCreateCategoriesRequestCategoriesInner
	isSet bool
}

func (v NullableCoreCourseCreateCategoriesRequestCategoriesInner) Get() *CoreCourseCreateCategoriesRequestCategoriesInner {
	return v.value
}

func (v *NullableCoreCourseCreateCategoriesRequestCategoriesInner) Set(val *CoreCourseCreateCategoriesRequestCategoriesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCourseCreateCategoriesRequestCategoriesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCourseCreateCategoriesRequestCategoriesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCourseCreateCategoriesRequestCategoriesInner(val *CoreCourseCreateCategoriesRequestCategoriesInner) *NullableCoreCourseCreateCategoriesRequestCategoriesInner {
	return &NullableCoreCourseCreateCategoriesRequestCategoriesInner{value: val, isSet: true}
}

func (v NullableCoreCourseCreateCategoriesRequestCategoriesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCourseCreateCategoriesRequestCategoriesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

