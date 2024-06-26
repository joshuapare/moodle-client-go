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

// checks if the CoreCompetencyUpdateTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCompetencyUpdateTemplateRequest{}

// CoreCompetencyUpdateTemplateRequest struct for CoreCompetencyUpdateTemplateRequest
type CoreCompetencyUpdateTemplateRequest struct {
	Template CoreCompetencyUpdateTemplateRequestTemplate `json:"template"`
}

type _CoreCompetencyUpdateTemplateRequest CoreCompetencyUpdateTemplateRequest

// NewCoreCompetencyUpdateTemplateRequest instantiates a new CoreCompetencyUpdateTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCompetencyUpdateTemplateRequest(template CoreCompetencyUpdateTemplateRequestTemplate) *CoreCompetencyUpdateTemplateRequest {
	this := CoreCompetencyUpdateTemplateRequest{}
	this.Template = template
	return &this
}

// NewCoreCompetencyUpdateTemplateRequestWithDefaults instantiates a new CoreCompetencyUpdateTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCompetencyUpdateTemplateRequestWithDefaults() *CoreCompetencyUpdateTemplateRequest {
	this := CoreCompetencyUpdateTemplateRequest{}
	return &this
}

// GetTemplate returns the Template field value
func (o *CoreCompetencyUpdateTemplateRequest) GetTemplate() CoreCompetencyUpdateTemplateRequestTemplate {
	if o == nil {
		var ret CoreCompetencyUpdateTemplateRequestTemplate
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyUpdateTemplateRequest) GetTemplateOk() (*CoreCompetencyUpdateTemplateRequestTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *CoreCompetencyUpdateTemplateRequest) SetTemplate(v CoreCompetencyUpdateTemplateRequestTemplate) {
	o.Template = v
}

func (o CoreCompetencyUpdateTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCompetencyUpdateTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["template"] = o.Template
	return toSerialize, nil
}

func (o *CoreCompetencyUpdateTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varCoreCompetencyUpdateTemplateRequest := _CoreCompetencyUpdateTemplateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreCompetencyUpdateTemplateRequest)

	if err != nil {
		return err
	}

	*o = CoreCompetencyUpdateTemplateRequest(varCoreCompetencyUpdateTemplateRequest)

	return err
}

type NullableCoreCompetencyUpdateTemplateRequest struct {
	value *CoreCompetencyUpdateTemplateRequest
	isSet bool
}

func (v NullableCoreCompetencyUpdateTemplateRequest) Get() *CoreCompetencyUpdateTemplateRequest {
	return v.value
}

func (v *NullableCoreCompetencyUpdateTemplateRequest) Set(val *CoreCompetencyUpdateTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCompetencyUpdateTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCompetencyUpdateTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCompetencyUpdateTemplateRequest(val *CoreCompetencyUpdateTemplateRequest) *NullableCoreCompetencyUpdateTemplateRequest {
	return &NullableCoreCompetencyUpdateTemplateRequest{value: val, isSet: true}
}

func (v NullableCoreCompetencyUpdateTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCompetencyUpdateTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


