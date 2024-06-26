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

// checks if the ToolDataprivacyCreateCategoryForm200ResponseCategory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ToolDataprivacyCreateCategoryForm200ResponseCategory{}

// ToolDataprivacyCreateCategoryForm200ResponseCategory struct for ToolDataprivacyCreateCategoryForm200ResponseCategory
type ToolDataprivacyCreateCategoryForm200ResponseCategory struct {
	// The category description.
	Description string `json:"description"`
	// description format (1 = HTML, 0 = MOODLE, 2 = PLAIN, or 4 = MARKDOWN)
	Descriptionformat *int32 `json:"descriptionformat,omitempty"`
	// id
	Id int32 `json:"id"`
	// The category name.
	Name string `json:"name"`
	// timecreated
	Timecreated int32 `json:"timecreated"`
	// timemodified
	Timemodified int32 `json:"timemodified"`
	// usermodified
	Usermodified int32 `json:"usermodified"`
}

type _ToolDataprivacyCreateCategoryForm200ResponseCategory ToolDataprivacyCreateCategoryForm200ResponseCategory

// NewToolDataprivacyCreateCategoryForm200ResponseCategory instantiates a new ToolDataprivacyCreateCategoryForm200ResponseCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToolDataprivacyCreateCategoryForm200ResponseCategory(description string, id int32, name string, timecreated int32, timemodified int32, usermodified int32) *ToolDataprivacyCreateCategoryForm200ResponseCategory {
	this := ToolDataprivacyCreateCategoryForm200ResponseCategory{}
	this.Description = description
	var descriptionformat int32 = 1
	this.Descriptionformat = &descriptionformat
	this.Id = id
	this.Name = name
	this.Timecreated = timecreated
	this.Timemodified = timemodified
	this.Usermodified = usermodified
	return &this
}

// NewToolDataprivacyCreateCategoryForm200ResponseCategoryWithDefaults instantiates a new ToolDataprivacyCreateCategoryForm200ResponseCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToolDataprivacyCreateCategoryForm200ResponseCategoryWithDefaults() *ToolDataprivacyCreateCategoryForm200ResponseCategory {
	this := ToolDataprivacyCreateCategoryForm200ResponseCategory{}
	var description string = ""
	this.Description = description
	var descriptionformat int32 = 1
	this.Descriptionformat = &descriptionformat
	var id int32 = 0
	this.Id = id
	var name string = "null"
	this.Name = name
	var timecreated int32 = 0
	this.Timecreated = timecreated
	var timemodified int32 = 0
	this.Timemodified = timemodified
	var usermodified int32 = 0
	this.Usermodified = usermodified
	return &this
}

// GetDescription returns the Description field value
func (o *ToolDataprivacyCreateCategoryForm200ResponseCategory) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ToolDataprivacyCreateCategoryForm200ResponseCategory) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ToolDataprivacyCreateCategoryForm200ResponseCategory) SetDescription(v string) {
	o.Description = v
}

// GetDescriptionformat returns the Descriptionformat field value if set, zero value otherwise.
func (o *ToolDataprivacyCreateCategoryForm200ResponseCategory) GetDescriptionformat() int32 {
	if o == nil || IsNil(o.Descriptionformat) {
		var ret int32
		return ret
	}
	return *o.Descriptionformat
}

// GetDescriptionformatOk returns a tuple with the Descriptionformat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolDataprivacyCreateCategoryForm200ResponseCategory) GetDescriptionformatOk() (*int32, bool) {
	if o == nil || IsNil(o.Descriptionformat) {
		return nil, false
	}
	return o.Descriptionformat, true
}

// HasDescriptionformat returns a boolean if a field has been set.
func (o *ToolDataprivacyCreateCategoryForm200ResponseCategory) HasDescriptionformat() bool {
	if o != nil && !IsNil(o.Descriptionformat) {
		return true
	}

	return false
}

// SetDescriptionformat gets a reference to the given int32 and assigns it to the Descriptionformat field.
func (o *ToolDataprivacyCreateCategoryForm200ResponseCategory) SetDescriptionformat(v int32) {
	o.Descriptionformat = &v
}

// GetId returns the Id field value
func (o *ToolDataprivacyCreateCategoryForm200ResponseCategory) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ToolDataprivacyCreateCategoryForm200ResponseCategory) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ToolDataprivacyCreateCategoryForm200ResponseCategory) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ToolDataprivacyCreateCategoryForm200ResponseCategory) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ToolDataprivacyCreateCategoryForm200ResponseCategory) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ToolDataprivacyCreateCategoryForm200ResponseCategory) SetName(v string) {
	o.Name = v
}

// GetTimecreated returns the Timecreated field value
func (o *ToolDataprivacyCreateCategoryForm200ResponseCategory) GetTimecreated() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timecreated
}

// GetTimecreatedOk returns a tuple with the Timecreated field value
// and a boolean to check if the value has been set.
func (o *ToolDataprivacyCreateCategoryForm200ResponseCategory) GetTimecreatedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timecreated, true
}

// SetTimecreated sets field value
func (o *ToolDataprivacyCreateCategoryForm200ResponseCategory) SetTimecreated(v int32) {
	o.Timecreated = v
}

// GetTimemodified returns the Timemodified field value
func (o *ToolDataprivacyCreateCategoryForm200ResponseCategory) GetTimemodified() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timemodified
}

// GetTimemodifiedOk returns a tuple with the Timemodified field value
// and a boolean to check if the value has been set.
func (o *ToolDataprivacyCreateCategoryForm200ResponseCategory) GetTimemodifiedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timemodified, true
}

// SetTimemodified sets field value
func (o *ToolDataprivacyCreateCategoryForm200ResponseCategory) SetTimemodified(v int32) {
	o.Timemodified = v
}

// GetUsermodified returns the Usermodified field value
func (o *ToolDataprivacyCreateCategoryForm200ResponseCategory) GetUsermodified() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Usermodified
}

// GetUsermodifiedOk returns a tuple with the Usermodified field value
// and a boolean to check if the value has been set.
func (o *ToolDataprivacyCreateCategoryForm200ResponseCategory) GetUsermodifiedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Usermodified, true
}

// SetUsermodified sets field value
func (o *ToolDataprivacyCreateCategoryForm200ResponseCategory) SetUsermodified(v int32) {
	o.Usermodified = v
}

func (o ToolDataprivacyCreateCategoryForm200ResponseCategory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToolDataprivacyCreateCategoryForm200ResponseCategory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	if !IsNil(o.Descriptionformat) {
		toSerialize["descriptionformat"] = o.Descriptionformat
	}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["timecreated"] = o.Timecreated
	toSerialize["timemodified"] = o.Timemodified
	toSerialize["usermodified"] = o.Usermodified
	return toSerialize, nil
}

func (o *ToolDataprivacyCreateCategoryForm200ResponseCategory) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"id",
		"name",
		"timecreated",
		"timemodified",
		"usermodified",
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

	varToolDataprivacyCreateCategoryForm200ResponseCategory := _ToolDataprivacyCreateCategoryForm200ResponseCategory{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varToolDataprivacyCreateCategoryForm200ResponseCategory)

	if err != nil {
		return err
	}

	*o = ToolDataprivacyCreateCategoryForm200ResponseCategory(varToolDataprivacyCreateCategoryForm200ResponseCategory)

	return err
}

type NullableToolDataprivacyCreateCategoryForm200ResponseCategory struct {
	value *ToolDataprivacyCreateCategoryForm200ResponseCategory
	isSet bool
}

func (v NullableToolDataprivacyCreateCategoryForm200ResponseCategory) Get() *ToolDataprivacyCreateCategoryForm200ResponseCategory {
	return v.value
}

func (v *NullableToolDataprivacyCreateCategoryForm200ResponseCategory) Set(val *ToolDataprivacyCreateCategoryForm200ResponseCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableToolDataprivacyCreateCategoryForm200ResponseCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableToolDataprivacyCreateCategoryForm200ResponseCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToolDataprivacyCreateCategoryForm200ResponseCategory(val *ToolDataprivacyCreateCategoryForm200ResponseCategory) *NullableToolDataprivacyCreateCategoryForm200ResponseCategory {
	return &NullableToolDataprivacyCreateCategoryForm200ResponseCategory{value: val, isSet: true}
}

func (v NullableToolDataprivacyCreateCategoryForm200ResponseCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToolDataprivacyCreateCategoryForm200ResponseCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


