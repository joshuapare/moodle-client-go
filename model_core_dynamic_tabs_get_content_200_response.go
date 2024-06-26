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

// checks if the CoreDynamicTabsGetContent200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreDynamicTabsGetContent200Response{}

// CoreDynamicTabsGetContent200Response struct for CoreDynamicTabsGetContent200Response
type CoreDynamicTabsGetContent200Response struct {
	// JSON-encoded data for template
	Content string `json:"content"`
	// JavaScript fragment
	Javascript string `json:"javascript"`
	// Template name
	Template string `json:"template"`
}

type _CoreDynamicTabsGetContent200Response CoreDynamicTabsGetContent200Response

// NewCoreDynamicTabsGetContent200Response instantiates a new CoreDynamicTabsGetContent200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreDynamicTabsGetContent200Response(content string, javascript string, template string) *CoreDynamicTabsGetContent200Response {
	this := CoreDynamicTabsGetContent200Response{}
	this.Content = content
	this.Javascript = javascript
	this.Template = template
	return &this
}

// NewCoreDynamicTabsGetContent200ResponseWithDefaults instantiates a new CoreDynamicTabsGetContent200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreDynamicTabsGetContent200ResponseWithDefaults() *CoreDynamicTabsGetContent200Response {
	this := CoreDynamicTabsGetContent200Response{}
	var content string = "null"
	this.Content = content
	var javascript string = "null"
	this.Javascript = javascript
	var template string = "null"
	this.Template = template
	return &this
}

// GetContent returns the Content field value
func (o *CoreDynamicTabsGetContent200Response) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *CoreDynamicTabsGetContent200Response) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *CoreDynamicTabsGetContent200Response) SetContent(v string) {
	o.Content = v
}

// GetJavascript returns the Javascript field value
func (o *CoreDynamicTabsGetContent200Response) GetJavascript() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Javascript
}

// GetJavascriptOk returns a tuple with the Javascript field value
// and a boolean to check if the value has been set.
func (o *CoreDynamicTabsGetContent200Response) GetJavascriptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Javascript, true
}

// SetJavascript sets field value
func (o *CoreDynamicTabsGetContent200Response) SetJavascript(v string) {
	o.Javascript = v
}

// GetTemplate returns the Template field value
func (o *CoreDynamicTabsGetContent200Response) GetTemplate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *CoreDynamicTabsGetContent200Response) GetTemplateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *CoreDynamicTabsGetContent200Response) SetTemplate(v string) {
	o.Template = v
}

func (o CoreDynamicTabsGetContent200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreDynamicTabsGetContent200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["content"] = o.Content
	toSerialize["javascript"] = o.Javascript
	toSerialize["template"] = o.Template
	return toSerialize, nil
}

func (o *CoreDynamicTabsGetContent200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"content",
		"javascript",
		"template",
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

	varCoreDynamicTabsGetContent200Response := _CoreDynamicTabsGetContent200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreDynamicTabsGetContent200Response)

	if err != nil {
		return err
	}

	*o = CoreDynamicTabsGetContent200Response(varCoreDynamicTabsGetContent200Response)

	return err
}

type NullableCoreDynamicTabsGetContent200Response struct {
	value *CoreDynamicTabsGetContent200Response
	isSet bool
}

func (v NullableCoreDynamicTabsGetContent200Response) Get() *CoreDynamicTabsGetContent200Response {
	return v.value
}

func (v *NullableCoreDynamicTabsGetContent200Response) Set(val *CoreDynamicTabsGetContent200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreDynamicTabsGetContent200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreDynamicTabsGetContent200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreDynamicTabsGetContent200Response(val *CoreDynamicTabsGetContent200Response) *NullableCoreDynamicTabsGetContent200Response {
	return &NullableCoreDynamicTabsGetContent200Response{value: val, isSet: true}
}

func (v NullableCoreDynamicTabsGetContent200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreDynamicTabsGetContent200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


