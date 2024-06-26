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

// checks if the CoreGradesGetFeedback200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreGradesGetFeedback200Response{}

// CoreGradesGetFeedback200Response struct for CoreGradesGetFeedback200Response
type CoreGradesGetFeedback200Response struct {
	// Additional field for the user (email or ID number, for example)
	Additionalfield string `json:"additionalfield"`
	// The full feedback text
	Feedbacktext string `json:"feedbacktext"`
	// Students name
	Fullname string `json:"fullname"`
	// Students picture
	Picture string `json:"picture"`
	// Title of the grade item that the feedback is for
	Title string `json:"title"`
}

type _CoreGradesGetFeedback200Response CoreGradesGetFeedback200Response

// NewCoreGradesGetFeedback200Response instantiates a new CoreGradesGetFeedback200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreGradesGetFeedback200Response(additionalfield string, feedbacktext string, fullname string, picture string, title string) *CoreGradesGetFeedback200Response {
	this := CoreGradesGetFeedback200Response{}
	this.Additionalfield = additionalfield
	this.Feedbacktext = feedbacktext
	this.Fullname = fullname
	this.Picture = picture
	this.Title = title
	return &this
}

// NewCoreGradesGetFeedback200ResponseWithDefaults instantiates a new CoreGradesGetFeedback200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreGradesGetFeedback200ResponseWithDefaults() *CoreGradesGetFeedback200Response {
	this := CoreGradesGetFeedback200Response{}
	var additionalfield string = "null"
	this.Additionalfield = additionalfield
	var feedbacktext string = "null"
	this.Feedbacktext = feedbacktext
	var fullname string = "null"
	this.Fullname = fullname
	var picture string = "null"
	this.Picture = picture
	var title string = "null"
	this.Title = title
	return &this
}

// GetAdditionalfield returns the Additionalfield field value
func (o *CoreGradesGetFeedback200Response) GetAdditionalfield() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Additionalfield
}

// GetAdditionalfieldOk returns a tuple with the Additionalfield field value
// and a boolean to check if the value has been set.
func (o *CoreGradesGetFeedback200Response) GetAdditionalfieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Additionalfield, true
}

// SetAdditionalfield sets field value
func (o *CoreGradesGetFeedback200Response) SetAdditionalfield(v string) {
	o.Additionalfield = v
}

// GetFeedbacktext returns the Feedbacktext field value
func (o *CoreGradesGetFeedback200Response) GetFeedbacktext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Feedbacktext
}

// GetFeedbacktextOk returns a tuple with the Feedbacktext field value
// and a boolean to check if the value has been set.
func (o *CoreGradesGetFeedback200Response) GetFeedbacktextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Feedbacktext, true
}

// SetFeedbacktext sets field value
func (o *CoreGradesGetFeedback200Response) SetFeedbacktext(v string) {
	o.Feedbacktext = v
}

// GetFullname returns the Fullname field value
func (o *CoreGradesGetFeedback200Response) GetFullname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fullname
}

// GetFullnameOk returns a tuple with the Fullname field value
// and a boolean to check if the value has been set.
func (o *CoreGradesGetFeedback200Response) GetFullnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fullname, true
}

// SetFullname sets field value
func (o *CoreGradesGetFeedback200Response) SetFullname(v string) {
	o.Fullname = v
}

// GetPicture returns the Picture field value
func (o *CoreGradesGetFeedback200Response) GetPicture() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Picture
}

// GetPictureOk returns a tuple with the Picture field value
// and a boolean to check if the value has been set.
func (o *CoreGradesGetFeedback200Response) GetPictureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Picture, true
}

// SetPicture sets field value
func (o *CoreGradesGetFeedback200Response) SetPicture(v string) {
	o.Picture = v
}

// GetTitle returns the Title field value
func (o *CoreGradesGetFeedback200Response) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *CoreGradesGetFeedback200Response) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *CoreGradesGetFeedback200Response) SetTitle(v string) {
	o.Title = v
}

func (o CoreGradesGetFeedback200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreGradesGetFeedback200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["additionalfield"] = o.Additionalfield
	toSerialize["feedbacktext"] = o.Feedbacktext
	toSerialize["fullname"] = o.Fullname
	toSerialize["picture"] = o.Picture
	toSerialize["title"] = o.Title
	return toSerialize, nil
}

func (o *CoreGradesGetFeedback200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"additionalfield",
		"feedbacktext",
		"fullname",
		"picture",
		"title",
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

	varCoreGradesGetFeedback200Response := _CoreGradesGetFeedback200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreGradesGetFeedback200Response)

	if err != nil {
		return err
	}

	*o = CoreGradesGetFeedback200Response(varCoreGradesGetFeedback200Response)

	return err
}

type NullableCoreGradesGetFeedback200Response struct {
	value *CoreGradesGetFeedback200Response
	isSet bool
}

func (v NullableCoreGradesGetFeedback200Response) Get() *CoreGradesGetFeedback200Response {
	return v.value
}

func (v *NullableCoreGradesGetFeedback200Response) Set(val *CoreGradesGetFeedback200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreGradesGetFeedback200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreGradesGetFeedback200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreGradesGetFeedback200Response(val *CoreGradesGetFeedback200Response) *NullableCoreGradesGetFeedback200Response {
	return &NullableCoreGradesGetFeedback200Response{value: val, isSet: true}
}

func (v NullableCoreGradesGetFeedback200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreGradesGetFeedback200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


