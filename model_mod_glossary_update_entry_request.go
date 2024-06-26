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

// checks if the ModGlossaryUpdateEntryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModGlossaryUpdateEntryRequest{}

// ModGlossaryUpdateEntryRequest struct for ModGlossaryUpdateEntryRequest
type ModGlossaryUpdateEntryRequest struct {
	// Glossary concept
	Concept string `json:"concept"`
	// Glossary concept definition
	Definition string `json:"definition"`
	// definition format (1 = HTML, 0 = MOODLE, 2 = PLAIN, or 4 = MARKDOWN)
	Definitionformat int32 `json:"definitionformat"`
	// Glossary entry id to update
	Entryid int32 `json:"entryid"`
	Options []ModGlossaryUpdateEntryRequestOptionsInner `json:"options,omitempty"`
}

type _ModGlossaryUpdateEntryRequest ModGlossaryUpdateEntryRequest

// NewModGlossaryUpdateEntryRequest instantiates a new ModGlossaryUpdateEntryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModGlossaryUpdateEntryRequest(concept string, definition string, definitionformat int32, entryid int32) *ModGlossaryUpdateEntryRequest {
	this := ModGlossaryUpdateEntryRequest{}
	this.Concept = concept
	this.Definition = definition
	this.Definitionformat = definitionformat
	this.Entryid = entryid
	return &this
}

// NewModGlossaryUpdateEntryRequestWithDefaults instantiates a new ModGlossaryUpdateEntryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModGlossaryUpdateEntryRequestWithDefaults() *ModGlossaryUpdateEntryRequest {
	this := ModGlossaryUpdateEntryRequest{}
	return &this
}

// GetConcept returns the Concept field value
func (o *ModGlossaryUpdateEntryRequest) GetConcept() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Concept
}

// GetConceptOk returns a tuple with the Concept field value
// and a boolean to check if the value has been set.
func (o *ModGlossaryUpdateEntryRequest) GetConceptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Concept, true
}

// SetConcept sets field value
func (o *ModGlossaryUpdateEntryRequest) SetConcept(v string) {
	o.Concept = v
}

// GetDefinition returns the Definition field value
func (o *ModGlossaryUpdateEntryRequest) GetDefinition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value
// and a boolean to check if the value has been set.
func (o *ModGlossaryUpdateEntryRequest) GetDefinitionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Definition, true
}

// SetDefinition sets field value
func (o *ModGlossaryUpdateEntryRequest) SetDefinition(v string) {
	o.Definition = v
}

// GetDefinitionformat returns the Definitionformat field value
func (o *ModGlossaryUpdateEntryRequest) GetDefinitionformat() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Definitionformat
}

// GetDefinitionformatOk returns a tuple with the Definitionformat field value
// and a boolean to check if the value has been set.
func (o *ModGlossaryUpdateEntryRequest) GetDefinitionformatOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Definitionformat, true
}

// SetDefinitionformat sets field value
func (o *ModGlossaryUpdateEntryRequest) SetDefinitionformat(v int32) {
	o.Definitionformat = v
}

// GetEntryid returns the Entryid field value
func (o *ModGlossaryUpdateEntryRequest) GetEntryid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Entryid
}

// GetEntryidOk returns a tuple with the Entryid field value
// and a boolean to check if the value has been set.
func (o *ModGlossaryUpdateEntryRequest) GetEntryidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entryid, true
}

// SetEntryid sets field value
func (o *ModGlossaryUpdateEntryRequest) SetEntryid(v int32) {
	o.Entryid = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ModGlossaryUpdateEntryRequest) GetOptions() []ModGlossaryUpdateEntryRequestOptionsInner {
	if o == nil || IsNil(o.Options) {
		var ret []ModGlossaryUpdateEntryRequestOptionsInner
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModGlossaryUpdateEntryRequest) GetOptionsOk() ([]ModGlossaryUpdateEntryRequestOptionsInner, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ModGlossaryUpdateEntryRequest) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []ModGlossaryUpdateEntryRequestOptionsInner and assigns it to the Options field.
func (o *ModGlossaryUpdateEntryRequest) SetOptions(v []ModGlossaryUpdateEntryRequestOptionsInner) {
	o.Options = v
}

func (o ModGlossaryUpdateEntryRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModGlossaryUpdateEntryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["concept"] = o.Concept
	toSerialize["definition"] = o.Definition
	toSerialize["definitionformat"] = o.Definitionformat
	toSerialize["entryid"] = o.Entryid
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

func (o *ModGlossaryUpdateEntryRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"concept",
		"definition",
		"definitionformat",
		"entryid",
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

	varModGlossaryUpdateEntryRequest := _ModGlossaryUpdateEntryRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModGlossaryUpdateEntryRequest)

	if err != nil {
		return err
	}

	*o = ModGlossaryUpdateEntryRequest(varModGlossaryUpdateEntryRequest)

	return err
}

type NullableModGlossaryUpdateEntryRequest struct {
	value *ModGlossaryUpdateEntryRequest
	isSet bool
}

func (v NullableModGlossaryUpdateEntryRequest) Get() *ModGlossaryUpdateEntryRequest {
	return v.value
}

func (v *NullableModGlossaryUpdateEntryRequest) Set(val *ModGlossaryUpdateEntryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModGlossaryUpdateEntryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModGlossaryUpdateEntryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModGlossaryUpdateEntryRequest(val *ModGlossaryUpdateEntryRequest) *NullableModGlossaryUpdateEntryRequest {
	return &NullableModGlossaryUpdateEntryRequest{value: val, isSet: true}
}

func (v NullableModGlossaryUpdateEntryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModGlossaryUpdateEntryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


