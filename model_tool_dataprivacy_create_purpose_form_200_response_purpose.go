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

// checks if the ToolDataprivacyCreatePurposeForm200ResponsePurpose type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ToolDataprivacyCreatePurposeForm200ResponsePurpose{}

// ToolDataprivacyCreatePurposeForm200ResponsePurpose struct for ToolDataprivacyCreatePurposeForm200ResponsePurpose
type ToolDataprivacyCreatePurposeForm200ResponsePurpose struct {
	// The purpose description.
	Description string `json:"description"`
	// description format (1 = HTML, 0 = MOODLE, 2 = PLAIN, or 4 = MARKDOWN)
	Descriptionformat *int32 `json:"descriptionformat,omitempty"`
	Formattedlawfulbases []ToolDataprivacyCreatePurposeForm200ResponsePurposeFormattedlawfulbasesInner `json:"formattedlawfulbases"`
	// formattedretentionperiod
	Formattedretentionperiod string `json:"formattedretentionperiod"`
	Formattedsensitivedatareasons []ToolDataprivacyCreatePurposeForm200ResponsePurposeFormattedlawfulbasesInner `json:"formattedsensitivedatareasons,omitempty"`
	// id
	Id int32 `json:"id"`
	// Comma-separated IDs matching records in tool_dataprivacy_lawfulbasis.
	Lawfulbases string `json:"lawfulbases"`
	// The purpose name.
	Name string `json:"name"`
	// Data retention with higher precedent over user's request to be forgotten.
	Protected int32 `json:"protected"`
	// Retention period. ISO_8601 durations format (as in DateInterval format).
	Retentionperiod string `json:"retentionperiod"`
	// roleoverrides
	Roleoverrides string `json:"roleoverrides"`
	// Comma-separated IDs matching records in tool_dataprivacy_sensitive
	Sensitivedatareasons string `json:"sensitivedatareasons"`
	// timecreated
	Timecreated int32 `json:"timecreated"`
	// timemodified
	Timemodified int32 `json:"timemodified"`
	// usermodified
	Usermodified int32 `json:"usermodified"`
}

type _ToolDataprivacyCreatePurposeForm200ResponsePurpose ToolDataprivacyCreatePurposeForm200ResponsePurpose

// NewToolDataprivacyCreatePurposeForm200ResponsePurpose instantiates a new ToolDataprivacyCreatePurposeForm200ResponsePurpose object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToolDataprivacyCreatePurposeForm200ResponsePurpose(description string, formattedlawfulbases []ToolDataprivacyCreatePurposeForm200ResponsePurposeFormattedlawfulbasesInner, formattedretentionperiod string, id int32, lawfulbases string, name string, protected int32, retentionperiod string, roleoverrides string, sensitivedatareasons string, timecreated int32, timemodified int32, usermodified int32) *ToolDataprivacyCreatePurposeForm200ResponsePurpose {
	this := ToolDataprivacyCreatePurposeForm200ResponsePurpose{}
	this.Description = description
	var descriptionformat int32 = 1
	this.Descriptionformat = &descriptionformat
	this.Formattedlawfulbases = formattedlawfulbases
	this.Formattedretentionperiod = formattedretentionperiod
	this.Id = id
	this.Lawfulbases = lawfulbases
	this.Name = name
	this.Protected = protected
	this.Retentionperiod = retentionperiod
	this.Roleoverrides = roleoverrides
	this.Sensitivedatareasons = sensitivedatareasons
	this.Timecreated = timecreated
	this.Timemodified = timemodified
	this.Usermodified = usermodified
	return &this
}

// NewToolDataprivacyCreatePurposeForm200ResponsePurposeWithDefaults instantiates a new ToolDataprivacyCreatePurposeForm200ResponsePurpose object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToolDataprivacyCreatePurposeForm200ResponsePurposeWithDefaults() *ToolDataprivacyCreatePurposeForm200ResponsePurpose {
	this := ToolDataprivacyCreatePurposeForm200ResponsePurpose{}
	var description string = ""
	this.Description = description
	var descriptionformat int32 = 1
	this.Descriptionformat = &descriptionformat
	var formattedretentionperiod string = "null"
	this.Formattedretentionperiod = formattedretentionperiod
	var id int32 = 0
	this.Id = id
	var lawfulbases string = "null"
	this.Lawfulbases = lawfulbases
	var name string = "null"
	this.Name = name
	var protected int32 = 0
	this.Protected = protected
	var retentionperiod string = ""
	this.Retentionperiod = retentionperiod
	var roleoverrides string = "null"
	this.Roleoverrides = roleoverrides
	var sensitivedatareasons string = ""
	this.Sensitivedatareasons = sensitivedatareasons
	var timecreated int32 = 0
	this.Timecreated = timecreated
	var timemodified int32 = 0
	this.Timemodified = timemodified
	var usermodified int32 = 0
	this.Usermodified = usermodified
	return &this
}

// GetDescription returns the Description field value
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) SetDescription(v string) {
	o.Description = v
}

// GetDescriptionformat returns the Descriptionformat field value if set, zero value otherwise.
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) GetDescriptionformat() int32 {
	if o == nil || IsNil(o.Descriptionformat) {
		var ret int32
		return ret
	}
	return *o.Descriptionformat
}

// GetDescriptionformatOk returns a tuple with the Descriptionformat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) GetDescriptionformatOk() (*int32, bool) {
	if o == nil || IsNil(o.Descriptionformat) {
		return nil, false
	}
	return o.Descriptionformat, true
}

// HasDescriptionformat returns a boolean if a field has been set.
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) HasDescriptionformat() bool {
	if o != nil && !IsNil(o.Descriptionformat) {
		return true
	}

	return false
}

// SetDescriptionformat gets a reference to the given int32 and assigns it to the Descriptionformat field.
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) SetDescriptionformat(v int32) {
	o.Descriptionformat = &v
}

// GetFormattedlawfulbases returns the Formattedlawfulbases field value
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) GetFormattedlawfulbases() []ToolDataprivacyCreatePurposeForm200ResponsePurposeFormattedlawfulbasesInner {
	if o == nil {
		var ret []ToolDataprivacyCreatePurposeForm200ResponsePurposeFormattedlawfulbasesInner
		return ret
	}

	return o.Formattedlawfulbases
}

// GetFormattedlawfulbasesOk returns a tuple with the Formattedlawfulbases field value
// and a boolean to check if the value has been set.
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) GetFormattedlawfulbasesOk() ([]ToolDataprivacyCreatePurposeForm200ResponsePurposeFormattedlawfulbasesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Formattedlawfulbases, true
}

// SetFormattedlawfulbases sets field value
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) SetFormattedlawfulbases(v []ToolDataprivacyCreatePurposeForm200ResponsePurposeFormattedlawfulbasesInner) {
	o.Formattedlawfulbases = v
}

// GetFormattedretentionperiod returns the Formattedretentionperiod field value
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) GetFormattedretentionperiod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Formattedretentionperiod
}

// GetFormattedretentionperiodOk returns a tuple with the Formattedretentionperiod field value
// and a boolean to check if the value has been set.
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) GetFormattedretentionperiodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Formattedretentionperiod, true
}

// SetFormattedretentionperiod sets field value
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) SetFormattedretentionperiod(v string) {
	o.Formattedretentionperiod = v
}

// GetFormattedsensitivedatareasons returns the Formattedsensitivedatareasons field value if set, zero value otherwise.
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) GetFormattedsensitivedatareasons() []ToolDataprivacyCreatePurposeForm200ResponsePurposeFormattedlawfulbasesInner {
	if o == nil || IsNil(o.Formattedsensitivedatareasons) {
		var ret []ToolDataprivacyCreatePurposeForm200ResponsePurposeFormattedlawfulbasesInner
		return ret
	}
	return o.Formattedsensitivedatareasons
}

// GetFormattedsensitivedatareasonsOk returns a tuple with the Formattedsensitivedatareasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) GetFormattedsensitivedatareasonsOk() ([]ToolDataprivacyCreatePurposeForm200ResponsePurposeFormattedlawfulbasesInner, bool) {
	if o == nil || IsNil(o.Formattedsensitivedatareasons) {
		return nil, false
	}
	return o.Formattedsensitivedatareasons, true
}

// HasFormattedsensitivedatareasons returns a boolean if a field has been set.
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) HasFormattedsensitivedatareasons() bool {
	if o != nil && !IsNil(o.Formattedsensitivedatareasons) {
		return true
	}

	return false
}

// SetFormattedsensitivedatareasons gets a reference to the given []ToolDataprivacyCreatePurposeForm200ResponsePurposeFormattedlawfulbasesInner and assigns it to the Formattedsensitivedatareasons field.
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) SetFormattedsensitivedatareasons(v []ToolDataprivacyCreatePurposeForm200ResponsePurposeFormattedlawfulbasesInner) {
	o.Formattedsensitivedatareasons = v
}

// GetId returns the Id field value
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) SetId(v int32) {
	o.Id = v
}

// GetLawfulbases returns the Lawfulbases field value
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) GetLawfulbases() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Lawfulbases
}

// GetLawfulbasesOk returns a tuple with the Lawfulbases field value
// and a boolean to check if the value has been set.
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) GetLawfulbasesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lawfulbases, true
}

// SetLawfulbases sets field value
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) SetLawfulbases(v string) {
	o.Lawfulbases = v
}

// GetName returns the Name field value
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) SetName(v string) {
	o.Name = v
}

// GetProtected returns the Protected field value
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) GetProtected() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Protected
}

// GetProtectedOk returns a tuple with the Protected field value
// and a boolean to check if the value has been set.
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) GetProtectedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protected, true
}

// SetProtected sets field value
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) SetProtected(v int32) {
	o.Protected = v
}

// GetRetentionperiod returns the Retentionperiod field value
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) GetRetentionperiod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Retentionperiod
}

// GetRetentionperiodOk returns a tuple with the Retentionperiod field value
// and a boolean to check if the value has been set.
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) GetRetentionperiodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Retentionperiod, true
}

// SetRetentionperiod sets field value
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) SetRetentionperiod(v string) {
	o.Retentionperiod = v
}

// GetRoleoverrides returns the Roleoverrides field value
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) GetRoleoverrides() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Roleoverrides
}

// GetRoleoverridesOk returns a tuple with the Roleoverrides field value
// and a boolean to check if the value has been set.
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) GetRoleoverridesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Roleoverrides, true
}

// SetRoleoverrides sets field value
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) SetRoleoverrides(v string) {
	o.Roleoverrides = v
}

// GetSensitivedatareasons returns the Sensitivedatareasons field value
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) GetSensitivedatareasons() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sensitivedatareasons
}

// GetSensitivedatareasonsOk returns a tuple with the Sensitivedatareasons field value
// and a boolean to check if the value has been set.
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) GetSensitivedatareasonsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sensitivedatareasons, true
}

// SetSensitivedatareasons sets field value
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) SetSensitivedatareasons(v string) {
	o.Sensitivedatareasons = v
}

// GetTimecreated returns the Timecreated field value
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) GetTimecreated() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timecreated
}

// GetTimecreatedOk returns a tuple with the Timecreated field value
// and a boolean to check if the value has been set.
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) GetTimecreatedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timecreated, true
}

// SetTimecreated sets field value
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) SetTimecreated(v int32) {
	o.Timecreated = v
}

// GetTimemodified returns the Timemodified field value
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) GetTimemodified() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timemodified
}

// GetTimemodifiedOk returns a tuple with the Timemodified field value
// and a boolean to check if the value has been set.
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) GetTimemodifiedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timemodified, true
}

// SetTimemodified sets field value
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) SetTimemodified(v int32) {
	o.Timemodified = v
}

// GetUsermodified returns the Usermodified field value
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) GetUsermodified() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Usermodified
}

// GetUsermodifiedOk returns a tuple with the Usermodified field value
// and a boolean to check if the value has been set.
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) GetUsermodifiedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Usermodified, true
}

// SetUsermodified sets field value
func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) SetUsermodified(v int32) {
	o.Usermodified = v
}

func (o ToolDataprivacyCreatePurposeForm200ResponsePurpose) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToolDataprivacyCreatePurposeForm200ResponsePurpose) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	if !IsNil(o.Descriptionformat) {
		toSerialize["descriptionformat"] = o.Descriptionformat
	}
	toSerialize["formattedlawfulbases"] = o.Formattedlawfulbases
	toSerialize["formattedretentionperiod"] = o.Formattedretentionperiod
	if !IsNil(o.Formattedsensitivedatareasons) {
		toSerialize["formattedsensitivedatareasons"] = o.Formattedsensitivedatareasons
	}
	toSerialize["id"] = o.Id
	toSerialize["lawfulbases"] = o.Lawfulbases
	toSerialize["name"] = o.Name
	toSerialize["protected"] = o.Protected
	toSerialize["retentionperiod"] = o.Retentionperiod
	toSerialize["roleoverrides"] = o.Roleoverrides
	toSerialize["sensitivedatareasons"] = o.Sensitivedatareasons
	toSerialize["timecreated"] = o.Timecreated
	toSerialize["timemodified"] = o.Timemodified
	toSerialize["usermodified"] = o.Usermodified
	return toSerialize, nil
}

func (o *ToolDataprivacyCreatePurposeForm200ResponsePurpose) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"formattedlawfulbases",
		"formattedretentionperiod",
		"id",
		"lawfulbases",
		"name",
		"protected",
		"retentionperiod",
		"roleoverrides",
		"sensitivedatareasons",
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

	varToolDataprivacyCreatePurposeForm200ResponsePurpose := _ToolDataprivacyCreatePurposeForm200ResponsePurpose{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varToolDataprivacyCreatePurposeForm200ResponsePurpose)

	if err != nil {
		return err
	}

	*o = ToolDataprivacyCreatePurposeForm200ResponsePurpose(varToolDataprivacyCreatePurposeForm200ResponsePurpose)

	return err
}

type NullableToolDataprivacyCreatePurposeForm200ResponsePurpose struct {
	value *ToolDataprivacyCreatePurposeForm200ResponsePurpose
	isSet bool
}

func (v NullableToolDataprivacyCreatePurposeForm200ResponsePurpose) Get() *ToolDataprivacyCreatePurposeForm200ResponsePurpose {
	return v.value
}

func (v *NullableToolDataprivacyCreatePurposeForm200ResponsePurpose) Set(val *ToolDataprivacyCreatePurposeForm200ResponsePurpose) {
	v.value = val
	v.isSet = true
}

func (v NullableToolDataprivacyCreatePurposeForm200ResponsePurpose) IsSet() bool {
	return v.isSet
}

func (v *NullableToolDataprivacyCreatePurposeForm200ResponsePurpose) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToolDataprivacyCreatePurposeForm200ResponsePurpose(val *ToolDataprivacyCreatePurposeForm200ResponsePurpose) *NullableToolDataprivacyCreatePurposeForm200ResponsePurpose {
	return &NullableToolDataprivacyCreatePurposeForm200ResponsePurpose{value: val, isSet: true}
}

func (v NullableToolDataprivacyCreatePurposeForm200ResponsePurpose) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToolDataprivacyCreatePurposeForm200ResponsePurpose) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

