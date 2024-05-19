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

// checks if the QbankEditquestionSetStatus200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QbankEditquestionSetStatus200Response{}

// QbankEditquestionSetStatus200Response struct for QbankEditquestionSetStatus200Response
type QbankEditquestionSetStatus200Response struct {
	// Error message if error exists
	Error string `json:"error"`
	// status: true if success
	Status bool `json:"status"`
	// statusname: name of the status
	Statusname string `json:"statusname"`
}

type _QbankEditquestionSetStatus200Response QbankEditquestionSetStatus200Response

// NewQbankEditquestionSetStatus200Response instantiates a new QbankEditquestionSetStatus200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQbankEditquestionSetStatus200Response(error_ string, status bool, statusname string) *QbankEditquestionSetStatus200Response {
	this := QbankEditquestionSetStatus200Response{}
	this.Error = error_
	this.Status = status
	this.Statusname = statusname
	return &this
}

// NewQbankEditquestionSetStatus200ResponseWithDefaults instantiates a new QbankEditquestionSetStatus200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQbankEditquestionSetStatus200ResponseWithDefaults() *QbankEditquestionSetStatus200Response {
	this := QbankEditquestionSetStatus200Response{}
	var error_ string = "null"
	this.Error = error_
	var statusname string = "null"
	this.Statusname = statusname
	return &this
}

// GetError returns the Error field value
func (o *QbankEditquestionSetStatus200Response) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *QbankEditquestionSetStatus200Response) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *QbankEditquestionSetStatus200Response) SetError(v string) {
	o.Error = v
}

// GetStatus returns the Status field value
func (o *QbankEditquestionSetStatus200Response) GetStatus() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *QbankEditquestionSetStatus200Response) GetStatusOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *QbankEditquestionSetStatus200Response) SetStatus(v bool) {
	o.Status = v
}

// GetStatusname returns the Statusname field value
func (o *QbankEditquestionSetStatus200Response) GetStatusname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Statusname
}

// GetStatusnameOk returns a tuple with the Statusname field value
// and a boolean to check if the value has been set.
func (o *QbankEditquestionSetStatus200Response) GetStatusnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Statusname, true
}

// SetStatusname sets field value
func (o *QbankEditquestionSetStatus200Response) SetStatusname(v string) {
	o.Statusname = v
}

func (o QbankEditquestionSetStatus200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QbankEditquestionSetStatus200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	toSerialize["status"] = o.Status
	toSerialize["statusname"] = o.Statusname
	return toSerialize, nil
}

func (o *QbankEditquestionSetStatus200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"error",
		"status",
		"statusname",
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

	varQbankEditquestionSetStatus200Response := _QbankEditquestionSetStatus200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQbankEditquestionSetStatus200Response)

	if err != nil {
		return err
	}

	*o = QbankEditquestionSetStatus200Response(varQbankEditquestionSetStatus200Response)

	return err
}

type NullableQbankEditquestionSetStatus200Response struct {
	value *QbankEditquestionSetStatus200Response
	isSet bool
}

func (v NullableQbankEditquestionSetStatus200Response) Get() *QbankEditquestionSetStatus200Response {
	return v.value
}

func (v *NullableQbankEditquestionSetStatus200Response) Set(val *QbankEditquestionSetStatus200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableQbankEditquestionSetStatus200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableQbankEditquestionSetStatus200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQbankEditquestionSetStatus200Response(val *QbankEditquestionSetStatus200Response) *NullableQbankEditquestionSetStatus200Response {
	return &NullableQbankEditquestionSetStatus200Response{value: val, isSet: true}
}

func (v NullableQbankEditquestionSetStatus200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQbankEditquestionSetStatus200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

