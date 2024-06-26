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

// checks if the CoreOutputLoadTemplateWithDependencies200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreOutputLoadTemplateWithDependencies200Response{}

// CoreOutputLoadTemplateWithDependencies200Response struct for CoreOutputLoadTemplateWithDependencies200Response
type CoreOutputLoadTemplateWithDependencies200Response struct {
	Strings []CoreOutputLoadTemplateWithDependencies200ResponseStringsInner `json:"strings"`
	Templates []CoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner `json:"templates"`
}

type _CoreOutputLoadTemplateWithDependencies200Response CoreOutputLoadTemplateWithDependencies200Response

// NewCoreOutputLoadTemplateWithDependencies200Response instantiates a new CoreOutputLoadTemplateWithDependencies200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreOutputLoadTemplateWithDependencies200Response(strings []CoreOutputLoadTemplateWithDependencies200ResponseStringsInner, templates []CoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner) *CoreOutputLoadTemplateWithDependencies200Response {
	this := CoreOutputLoadTemplateWithDependencies200Response{}
	this.Strings = strings
	this.Templates = templates
	return &this
}

// NewCoreOutputLoadTemplateWithDependencies200ResponseWithDefaults instantiates a new CoreOutputLoadTemplateWithDependencies200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreOutputLoadTemplateWithDependencies200ResponseWithDefaults() *CoreOutputLoadTemplateWithDependencies200Response {
	this := CoreOutputLoadTemplateWithDependencies200Response{}
	return &this
}

// GetStrings returns the Strings field value
func (o *CoreOutputLoadTemplateWithDependencies200Response) GetStrings() []CoreOutputLoadTemplateWithDependencies200ResponseStringsInner {
	if o == nil {
		var ret []CoreOutputLoadTemplateWithDependencies200ResponseStringsInner
		return ret
	}

	return o.Strings
}

// GetStringsOk returns a tuple with the Strings field value
// and a boolean to check if the value has been set.
func (o *CoreOutputLoadTemplateWithDependencies200Response) GetStringsOk() ([]CoreOutputLoadTemplateWithDependencies200ResponseStringsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Strings, true
}

// SetStrings sets field value
func (o *CoreOutputLoadTemplateWithDependencies200Response) SetStrings(v []CoreOutputLoadTemplateWithDependencies200ResponseStringsInner) {
	o.Strings = v
}

// GetTemplates returns the Templates field value
func (o *CoreOutputLoadTemplateWithDependencies200Response) GetTemplates() []CoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner {
	if o == nil {
		var ret []CoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner
		return ret
	}

	return o.Templates
}

// GetTemplatesOk returns a tuple with the Templates field value
// and a boolean to check if the value has been set.
func (o *CoreOutputLoadTemplateWithDependencies200Response) GetTemplatesOk() ([]CoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Templates, true
}

// SetTemplates sets field value
func (o *CoreOutputLoadTemplateWithDependencies200Response) SetTemplates(v []CoreOutputLoadTemplateWithDependencies200ResponseTemplatesInner) {
	o.Templates = v
}

func (o CoreOutputLoadTemplateWithDependencies200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreOutputLoadTemplateWithDependencies200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["strings"] = o.Strings
	toSerialize["templates"] = o.Templates
	return toSerialize, nil
}

func (o *CoreOutputLoadTemplateWithDependencies200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"strings",
		"templates",
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

	varCoreOutputLoadTemplateWithDependencies200Response := _CoreOutputLoadTemplateWithDependencies200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreOutputLoadTemplateWithDependencies200Response)

	if err != nil {
		return err
	}

	*o = CoreOutputLoadTemplateWithDependencies200Response(varCoreOutputLoadTemplateWithDependencies200Response)

	return err
}

type NullableCoreOutputLoadTemplateWithDependencies200Response struct {
	value *CoreOutputLoadTemplateWithDependencies200Response
	isSet bool
}

func (v NullableCoreOutputLoadTemplateWithDependencies200Response) Get() *CoreOutputLoadTemplateWithDependencies200Response {
	return v.value
}

func (v *NullableCoreOutputLoadTemplateWithDependencies200Response) Set(val *CoreOutputLoadTemplateWithDependencies200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreOutputLoadTemplateWithDependencies200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreOutputLoadTemplateWithDependencies200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreOutputLoadTemplateWithDependencies200Response(val *CoreOutputLoadTemplateWithDependencies200Response) *NullableCoreOutputLoadTemplateWithDependencies200Response {
	return &NullableCoreOutputLoadTemplateWithDependencies200Response{value: val, isSet: true}
}

func (v NullableCoreOutputLoadTemplateWithDependencies200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreOutputLoadTemplateWithDependencies200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


