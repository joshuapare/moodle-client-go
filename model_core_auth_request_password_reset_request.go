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

// checks if the CoreAuthRequestPasswordResetRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreAuthRequestPasswordResetRequest{}

// CoreAuthRequestPasswordResetRequest struct for CoreAuthRequestPasswordResetRequest
type CoreAuthRequestPasswordResetRequest struct {
	// User email
	Email *string `json:"email,omitempty"`
	// User name
	Username *string `json:"username,omitempty"`
}

// NewCoreAuthRequestPasswordResetRequest instantiates a new CoreAuthRequestPasswordResetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreAuthRequestPasswordResetRequest() *CoreAuthRequestPasswordResetRequest {
	this := CoreAuthRequestPasswordResetRequest{}
	var email string = ""
	this.Email = &email
	var username string = ""
	this.Username = &username
	return &this
}

// NewCoreAuthRequestPasswordResetRequestWithDefaults instantiates a new CoreAuthRequestPasswordResetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreAuthRequestPasswordResetRequestWithDefaults() *CoreAuthRequestPasswordResetRequest {
	this := CoreAuthRequestPasswordResetRequest{}
	var email string = ""
	this.Email = &email
	var username string = ""
	this.Username = &username
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CoreAuthRequestPasswordResetRequest) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreAuthRequestPasswordResetRequest) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CoreAuthRequestPasswordResetRequest) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CoreAuthRequestPasswordResetRequest) SetEmail(v string) {
	o.Email = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *CoreAuthRequestPasswordResetRequest) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreAuthRequestPasswordResetRequest) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *CoreAuthRequestPasswordResetRequest) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *CoreAuthRequestPasswordResetRequest) SetUsername(v string) {
	o.Username = &v
}

func (o CoreAuthRequestPasswordResetRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreAuthRequestPasswordResetRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableCoreAuthRequestPasswordResetRequest struct {
	value *CoreAuthRequestPasswordResetRequest
	isSet bool
}

func (v NullableCoreAuthRequestPasswordResetRequest) Get() *CoreAuthRequestPasswordResetRequest {
	return v.value
}

func (v *NullableCoreAuthRequestPasswordResetRequest) Set(val *CoreAuthRequestPasswordResetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreAuthRequestPasswordResetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreAuthRequestPasswordResetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreAuthRequestPasswordResetRequest(val *CoreAuthRequestPasswordResetRequest) *NullableCoreAuthRequestPasswordResetRequest {
	return &NullableCoreAuthRequestPasswordResetRequest{value: val, isSet: true}
}

func (v NullableCoreAuthRequestPasswordResetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreAuthRequestPasswordResetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


