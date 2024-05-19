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

// checks if the ModGlossaryAddEntryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModGlossaryAddEntryRequest{}

// ModGlossaryAddEntryRequest struct for ModGlossaryAddEntryRequest
type ModGlossaryAddEntryRequest struct {
	// Glossary concept
	Concept string `json:"concept"`
	// Glossary concept definition
	Definition string `json:"definition"`
	// definition format (1 = HTML, 0 = MOODLE, 2 = PLAIN, or 4 = MARKDOWN)
	Definitionformat int32 `json:"definitionformat"`
	// Glossary id
	Glossaryid int32 `json:"glossaryid"`
	Options []ModGlossaryAddEntryRequestOptionsInner `json:"options,omitempty"`
}

type _ModGlossaryAddEntryRequest ModGlossaryAddEntryRequest

// NewModGlossaryAddEntryRequest instantiates a new ModGlossaryAddEntryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModGlossaryAddEntryRequest(concept string, definition string, definitionformat int32, glossaryid int32) *ModGlossaryAddEntryRequest {
	this := ModGlossaryAddEntryRequest{}
	this.Concept = concept
	this.Definition = definition
	this.Definitionformat = definitionformat
	this.Glossaryid = glossaryid
	return &this
}

// NewModGlossaryAddEntryRequestWithDefaults instantiates a new ModGlossaryAddEntryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModGlossaryAddEntryRequestWithDefaults() *ModGlossaryAddEntryRequest {
	this := ModGlossaryAddEntryRequest{}
	var concept string = "null"
	this.Concept = concept
	var definition string = "null"
	this.Definition = definition
	var glossaryid int32 = null
	this.Glossaryid = glossaryid
	return &this
}

// GetConcept returns the Concept field value
func (o *ModGlossaryAddEntryRequest) GetConcept() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Concept
}

// GetConceptOk returns a tuple with the Concept field value
// and a boolean to check if the value has been set.
func (o *ModGlossaryAddEntryRequest) GetConceptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Concept, true
}

// SetConcept sets field value
func (o *ModGlossaryAddEntryRequest) SetConcept(v string) {
	o.Concept = v
}

// GetDefinition returns the Definition field value
func (o *ModGlossaryAddEntryRequest) GetDefinition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value
// and a boolean to check if the value has been set.
func (o *ModGlossaryAddEntryRequest) GetDefinitionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Definition, true
}

// SetDefinition sets field value
func (o *ModGlossaryAddEntryRequest) SetDefinition(v string) {
	o.Definition = v
}

// GetDefinitionformat returns the Definitionformat field value
func (o *ModGlossaryAddEntryRequest) GetDefinitionformat() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Definitionformat
}

// GetDefinitionformatOk returns a tuple with the Definitionformat field value
// and a boolean to check if the value has been set.
func (o *ModGlossaryAddEntryRequest) GetDefinitionformatOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Definitionformat, true
}

// SetDefinitionformat sets field value
func (o *ModGlossaryAddEntryRequest) SetDefinitionformat(v int32) {
	o.Definitionformat = v
}

// GetGlossaryid returns the Glossaryid field value
func (o *ModGlossaryAddEntryRequest) GetGlossaryid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Glossaryid
}

// GetGlossaryidOk returns a tuple with the Glossaryid field value
// and a boolean to check if the value has been set.
func (o *ModGlossaryAddEntryRequest) GetGlossaryidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Glossaryid, true
}

// SetGlossaryid sets field value
func (o *ModGlossaryAddEntryRequest) SetGlossaryid(v int32) {
	o.Glossaryid = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ModGlossaryAddEntryRequest) GetOptions() []ModGlossaryAddEntryRequestOptionsInner {
	if o == nil || IsNil(o.Options) {
		var ret []ModGlossaryAddEntryRequestOptionsInner
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModGlossaryAddEntryRequest) GetOptionsOk() ([]ModGlossaryAddEntryRequestOptionsInner, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ModGlossaryAddEntryRequest) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []ModGlossaryAddEntryRequestOptionsInner and assigns it to the Options field.
func (o *ModGlossaryAddEntryRequest) SetOptions(v []ModGlossaryAddEntryRequestOptionsInner) {
	o.Options = v
}

func (o ModGlossaryAddEntryRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModGlossaryAddEntryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["concept"] = o.Concept
	toSerialize["definition"] = o.Definition
	toSerialize["definitionformat"] = o.Definitionformat
	toSerialize["glossaryid"] = o.Glossaryid
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

func (o *ModGlossaryAddEntryRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"concept",
		"definition",
		"definitionformat",
		"glossaryid",
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

	varModGlossaryAddEntryRequest := _ModGlossaryAddEntryRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModGlossaryAddEntryRequest)

	if err != nil {
		return err
	}

	*o = ModGlossaryAddEntryRequest(varModGlossaryAddEntryRequest)

	return err
}

type NullableModGlossaryAddEntryRequest struct {
	value *ModGlossaryAddEntryRequest
	isSet bool
}

func (v NullableModGlossaryAddEntryRequest) Get() *ModGlossaryAddEntryRequest {
	return v.value
}

func (v *NullableModGlossaryAddEntryRequest) Set(val *ModGlossaryAddEntryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModGlossaryAddEntryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModGlossaryAddEntryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModGlossaryAddEntryRequest(val *ModGlossaryAddEntryRequest) *NullableModGlossaryAddEntryRequest {
	return &NullableModGlossaryAddEntryRequest{value: val, isSet: true}
}

func (v NullableModGlossaryAddEntryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModGlossaryAddEntryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

