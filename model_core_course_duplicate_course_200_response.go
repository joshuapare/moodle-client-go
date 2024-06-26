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

// checks if the CoreCourseDuplicateCourse200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCourseDuplicateCourse200Response{}

// CoreCourseDuplicateCourse200Response struct for CoreCourseDuplicateCourse200Response
type CoreCourseDuplicateCourse200Response struct {
	// course id
	Id int32 `json:"id"`
	// short name
	Shortname string `json:"shortname"`
}

type _CoreCourseDuplicateCourse200Response CoreCourseDuplicateCourse200Response

// NewCoreCourseDuplicateCourse200Response instantiates a new CoreCourseDuplicateCourse200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCourseDuplicateCourse200Response(id int32, shortname string) *CoreCourseDuplicateCourse200Response {
	this := CoreCourseDuplicateCourse200Response{}
	this.Id = id
	this.Shortname = shortname
	return &this
}

// NewCoreCourseDuplicateCourse200ResponseWithDefaults instantiates a new CoreCourseDuplicateCourse200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCourseDuplicateCourse200ResponseWithDefaults() *CoreCourseDuplicateCourse200Response {
	this := CoreCourseDuplicateCourse200Response{}
	var shortname string = "null"
	this.Shortname = shortname
	return &this
}

// GetId returns the Id field value
func (o *CoreCourseDuplicateCourse200Response) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CoreCourseDuplicateCourse200Response) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CoreCourseDuplicateCourse200Response) SetId(v int32) {
	o.Id = v
}

// GetShortname returns the Shortname field value
func (o *CoreCourseDuplicateCourse200Response) GetShortname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Shortname
}

// GetShortnameOk returns a tuple with the Shortname field value
// and a boolean to check if the value has been set.
func (o *CoreCourseDuplicateCourse200Response) GetShortnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Shortname, true
}

// SetShortname sets field value
func (o *CoreCourseDuplicateCourse200Response) SetShortname(v string) {
	o.Shortname = v
}

func (o CoreCourseDuplicateCourse200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCourseDuplicateCourse200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["shortname"] = o.Shortname
	return toSerialize, nil
}

func (o *CoreCourseDuplicateCourse200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"shortname",
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

	varCoreCourseDuplicateCourse200Response := _CoreCourseDuplicateCourse200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreCourseDuplicateCourse200Response)

	if err != nil {
		return err
	}

	*o = CoreCourseDuplicateCourse200Response(varCoreCourseDuplicateCourse200Response)

	return err
}

type NullableCoreCourseDuplicateCourse200Response struct {
	value *CoreCourseDuplicateCourse200Response
	isSet bool
}

func (v NullableCoreCourseDuplicateCourse200Response) Get() *CoreCourseDuplicateCourse200Response {
	return v.value
}

func (v *NullableCoreCourseDuplicateCourse200Response) Set(val *CoreCourseDuplicateCourse200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCourseDuplicateCourse200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCourseDuplicateCourse200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCourseDuplicateCourse200Response(val *CoreCourseDuplicateCourse200Response) *NullableCoreCourseDuplicateCourse200Response {
	return &NullableCoreCourseDuplicateCourse200Response{value: val, isSet: true}
}

func (v NullableCoreCourseDuplicateCourse200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCourseDuplicateCourse200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


