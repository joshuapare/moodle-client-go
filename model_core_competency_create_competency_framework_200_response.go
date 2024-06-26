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

// checks if the CoreCompetencyCreateCompetencyFramework200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCompetencyCreateCompetencyFramework200Response{}

// CoreCompetencyCreateCompetencyFramework200Response struct for CoreCompetencyCreateCompetencyFramework200Response
type CoreCompetencyCreateCompetencyFramework200Response struct {
	// canmanage
	Canmanage bool `json:"canmanage"`
	// competenciescount
	Competenciescount int32 `json:"competenciescount"`
	// contextid
	Contextid int32 `json:"contextid"`
	// contextname
	Contextname string `json:"contextname"`
	// contextnamenoprefix
	Contextnamenoprefix string `json:"contextnamenoprefix"`
	// description
	Description string `json:"description"`
	// description format (1 = HTML, 0 = MOODLE, 2 = PLAIN, or 4 = MARKDOWN)
	Descriptionformat *int32 `json:"descriptionformat,omitempty"`
	// id
	Id int32 `json:"id"`
	// idnumber
	Idnumber string `json:"idnumber"`
	// scaleconfiguration
	Scaleconfiguration string `json:"scaleconfiguration"`
	// scaleid
	Scaleid int32 `json:"scaleid"`
	// shortname
	Shortname string `json:"shortname"`
	// taxonomies
	Taxonomies string `json:"taxonomies"`
	// timecreated
	Timecreated int32 `json:"timecreated"`
	// timemodified
	Timemodified int32 `json:"timemodified"`
	// usermodified
	Usermodified int32 `json:"usermodified"`
	// visible
	Visible bool `json:"visible"`
}

type _CoreCompetencyCreateCompetencyFramework200Response CoreCompetencyCreateCompetencyFramework200Response

// NewCoreCompetencyCreateCompetencyFramework200Response instantiates a new CoreCompetencyCreateCompetencyFramework200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCompetencyCreateCompetencyFramework200Response(canmanage bool, competenciescount int32, contextid int32, contextname string, contextnamenoprefix string, description string, id int32, idnumber string, scaleconfiguration string, scaleid int32, shortname string, taxonomies string, timecreated int32, timemodified int32, usermodified int32, visible bool) *CoreCompetencyCreateCompetencyFramework200Response {
	this := CoreCompetencyCreateCompetencyFramework200Response{}
	this.Canmanage = canmanage
	this.Competenciescount = competenciescount
	this.Contextid = contextid
	this.Contextname = contextname
	this.Contextnamenoprefix = contextnamenoprefix
	this.Description = description
	var descriptionformat int32 = 1
	this.Descriptionformat = &descriptionformat
	this.Id = id
	this.Idnumber = idnumber
	this.Scaleconfiguration = scaleconfiguration
	this.Scaleid = scaleid
	this.Shortname = shortname
	this.Taxonomies = taxonomies
	this.Timecreated = timecreated
	this.Timemodified = timemodified
	this.Usermodified = usermodified
	this.Visible = visible
	return &this
}

// NewCoreCompetencyCreateCompetencyFramework200ResponseWithDefaults instantiates a new CoreCompetencyCreateCompetencyFramework200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCompetencyCreateCompetencyFramework200ResponseWithDefaults() *CoreCompetencyCreateCompetencyFramework200Response {
	this := CoreCompetencyCreateCompetencyFramework200Response{}
	var canmanage bool = null
	this.Canmanage = canmanage
	var competenciescount int32 = null
	this.Competenciescount = competenciescount
	var contextid int32 = null
	this.Contextid = contextid
	var contextname string = "null"
	this.Contextname = contextname
	var contextnamenoprefix string = "null"
	this.Contextnamenoprefix = contextnamenoprefix
	var description string = ""
	this.Description = description
	var descriptionformat int32 = 1
	this.Descriptionformat = &descriptionformat
	var id int32 = 0
	this.Id = id
	var taxonomies string = ""
	this.Taxonomies = taxonomies
	var timecreated int32 = 0
	this.Timecreated = timecreated
	var timemodified int32 = 0
	this.Timemodified = timemodified
	var usermodified int32 = 0
	this.Usermodified = usermodified
	var visible bool = 1
	this.Visible = visible
	return &this
}

// GetCanmanage returns the Canmanage field value
func (o *CoreCompetencyCreateCompetencyFramework200Response) GetCanmanage() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Canmanage
}

// GetCanmanageOk returns a tuple with the Canmanage field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyCreateCompetencyFramework200Response) GetCanmanageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Canmanage, true
}

// SetCanmanage sets field value
func (o *CoreCompetencyCreateCompetencyFramework200Response) SetCanmanage(v bool) {
	o.Canmanage = v
}

// GetCompetenciescount returns the Competenciescount field value
func (o *CoreCompetencyCreateCompetencyFramework200Response) GetCompetenciescount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Competenciescount
}

// GetCompetenciescountOk returns a tuple with the Competenciescount field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyCreateCompetencyFramework200Response) GetCompetenciescountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Competenciescount, true
}

// SetCompetenciescount sets field value
func (o *CoreCompetencyCreateCompetencyFramework200Response) SetCompetenciescount(v int32) {
	o.Competenciescount = v
}

// GetContextid returns the Contextid field value
func (o *CoreCompetencyCreateCompetencyFramework200Response) GetContextid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Contextid
}

// GetContextidOk returns a tuple with the Contextid field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyCreateCompetencyFramework200Response) GetContextidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contextid, true
}

// SetContextid sets field value
func (o *CoreCompetencyCreateCompetencyFramework200Response) SetContextid(v int32) {
	o.Contextid = v
}

// GetContextname returns the Contextname field value
func (o *CoreCompetencyCreateCompetencyFramework200Response) GetContextname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Contextname
}

// GetContextnameOk returns a tuple with the Contextname field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyCreateCompetencyFramework200Response) GetContextnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contextname, true
}

// SetContextname sets field value
func (o *CoreCompetencyCreateCompetencyFramework200Response) SetContextname(v string) {
	o.Contextname = v
}

// GetContextnamenoprefix returns the Contextnamenoprefix field value
func (o *CoreCompetencyCreateCompetencyFramework200Response) GetContextnamenoprefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Contextnamenoprefix
}

// GetContextnamenoprefixOk returns a tuple with the Contextnamenoprefix field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyCreateCompetencyFramework200Response) GetContextnamenoprefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contextnamenoprefix, true
}

// SetContextnamenoprefix sets field value
func (o *CoreCompetencyCreateCompetencyFramework200Response) SetContextnamenoprefix(v string) {
	o.Contextnamenoprefix = v
}

// GetDescription returns the Description field value
func (o *CoreCompetencyCreateCompetencyFramework200Response) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyCreateCompetencyFramework200Response) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CoreCompetencyCreateCompetencyFramework200Response) SetDescription(v string) {
	o.Description = v
}

// GetDescriptionformat returns the Descriptionformat field value if set, zero value otherwise.
func (o *CoreCompetencyCreateCompetencyFramework200Response) GetDescriptionformat() int32 {
	if o == nil || IsNil(o.Descriptionformat) {
		var ret int32
		return ret
	}
	return *o.Descriptionformat
}

// GetDescriptionformatOk returns a tuple with the Descriptionformat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyCreateCompetencyFramework200Response) GetDescriptionformatOk() (*int32, bool) {
	if o == nil || IsNil(o.Descriptionformat) {
		return nil, false
	}
	return o.Descriptionformat, true
}

// HasDescriptionformat returns a boolean if a field has been set.
func (o *CoreCompetencyCreateCompetencyFramework200Response) HasDescriptionformat() bool {
	if o != nil && !IsNil(o.Descriptionformat) {
		return true
	}

	return false
}

// SetDescriptionformat gets a reference to the given int32 and assigns it to the Descriptionformat field.
func (o *CoreCompetencyCreateCompetencyFramework200Response) SetDescriptionformat(v int32) {
	o.Descriptionformat = &v
}

// GetId returns the Id field value
func (o *CoreCompetencyCreateCompetencyFramework200Response) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyCreateCompetencyFramework200Response) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CoreCompetencyCreateCompetencyFramework200Response) SetId(v int32) {
	o.Id = v
}

// GetIdnumber returns the Idnumber field value
func (o *CoreCompetencyCreateCompetencyFramework200Response) GetIdnumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Idnumber
}

// GetIdnumberOk returns a tuple with the Idnumber field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyCreateCompetencyFramework200Response) GetIdnumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Idnumber, true
}

// SetIdnumber sets field value
func (o *CoreCompetencyCreateCompetencyFramework200Response) SetIdnumber(v string) {
	o.Idnumber = v
}

// GetScaleconfiguration returns the Scaleconfiguration field value
func (o *CoreCompetencyCreateCompetencyFramework200Response) GetScaleconfiguration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scaleconfiguration
}

// GetScaleconfigurationOk returns a tuple with the Scaleconfiguration field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyCreateCompetencyFramework200Response) GetScaleconfigurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scaleconfiguration, true
}

// SetScaleconfiguration sets field value
func (o *CoreCompetencyCreateCompetencyFramework200Response) SetScaleconfiguration(v string) {
	o.Scaleconfiguration = v
}

// GetScaleid returns the Scaleid field value
func (o *CoreCompetencyCreateCompetencyFramework200Response) GetScaleid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Scaleid
}

// GetScaleidOk returns a tuple with the Scaleid field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyCreateCompetencyFramework200Response) GetScaleidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scaleid, true
}

// SetScaleid sets field value
func (o *CoreCompetencyCreateCompetencyFramework200Response) SetScaleid(v int32) {
	o.Scaleid = v
}

// GetShortname returns the Shortname field value
func (o *CoreCompetencyCreateCompetencyFramework200Response) GetShortname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Shortname
}

// GetShortnameOk returns a tuple with the Shortname field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyCreateCompetencyFramework200Response) GetShortnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Shortname, true
}

// SetShortname sets field value
func (o *CoreCompetencyCreateCompetencyFramework200Response) SetShortname(v string) {
	o.Shortname = v
}

// GetTaxonomies returns the Taxonomies field value
func (o *CoreCompetencyCreateCompetencyFramework200Response) GetTaxonomies() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Taxonomies
}

// GetTaxonomiesOk returns a tuple with the Taxonomies field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyCreateCompetencyFramework200Response) GetTaxonomiesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Taxonomies, true
}

// SetTaxonomies sets field value
func (o *CoreCompetencyCreateCompetencyFramework200Response) SetTaxonomies(v string) {
	o.Taxonomies = v
}

// GetTimecreated returns the Timecreated field value
func (o *CoreCompetencyCreateCompetencyFramework200Response) GetTimecreated() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timecreated
}

// GetTimecreatedOk returns a tuple with the Timecreated field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyCreateCompetencyFramework200Response) GetTimecreatedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timecreated, true
}

// SetTimecreated sets field value
func (o *CoreCompetencyCreateCompetencyFramework200Response) SetTimecreated(v int32) {
	o.Timecreated = v
}

// GetTimemodified returns the Timemodified field value
func (o *CoreCompetencyCreateCompetencyFramework200Response) GetTimemodified() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timemodified
}

// GetTimemodifiedOk returns a tuple with the Timemodified field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyCreateCompetencyFramework200Response) GetTimemodifiedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timemodified, true
}

// SetTimemodified sets field value
func (o *CoreCompetencyCreateCompetencyFramework200Response) SetTimemodified(v int32) {
	o.Timemodified = v
}

// GetUsermodified returns the Usermodified field value
func (o *CoreCompetencyCreateCompetencyFramework200Response) GetUsermodified() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Usermodified
}

// GetUsermodifiedOk returns a tuple with the Usermodified field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyCreateCompetencyFramework200Response) GetUsermodifiedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Usermodified, true
}

// SetUsermodified sets field value
func (o *CoreCompetencyCreateCompetencyFramework200Response) SetUsermodified(v int32) {
	o.Usermodified = v
}

// GetVisible returns the Visible field value
func (o *CoreCompetencyCreateCompetencyFramework200Response) GetVisible() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Visible
}

// GetVisibleOk returns a tuple with the Visible field value
// and a boolean to check if the value has been set.
func (o *CoreCompetencyCreateCompetencyFramework200Response) GetVisibleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Visible, true
}

// SetVisible sets field value
func (o *CoreCompetencyCreateCompetencyFramework200Response) SetVisible(v bool) {
	o.Visible = v
}

func (o CoreCompetencyCreateCompetencyFramework200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCompetencyCreateCompetencyFramework200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["canmanage"] = o.Canmanage
	toSerialize["competenciescount"] = o.Competenciescount
	toSerialize["contextid"] = o.Contextid
	toSerialize["contextname"] = o.Contextname
	toSerialize["contextnamenoprefix"] = o.Contextnamenoprefix
	toSerialize["description"] = o.Description
	if !IsNil(o.Descriptionformat) {
		toSerialize["descriptionformat"] = o.Descriptionformat
	}
	toSerialize["id"] = o.Id
	toSerialize["idnumber"] = o.Idnumber
	toSerialize["scaleconfiguration"] = o.Scaleconfiguration
	toSerialize["scaleid"] = o.Scaleid
	toSerialize["shortname"] = o.Shortname
	toSerialize["taxonomies"] = o.Taxonomies
	toSerialize["timecreated"] = o.Timecreated
	toSerialize["timemodified"] = o.Timemodified
	toSerialize["usermodified"] = o.Usermodified
	toSerialize["visible"] = o.Visible
	return toSerialize, nil
}

func (o *CoreCompetencyCreateCompetencyFramework200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"canmanage",
		"competenciescount",
		"contextid",
		"contextname",
		"contextnamenoprefix",
		"description",
		"id",
		"idnumber",
		"scaleconfiguration",
		"scaleid",
		"shortname",
		"taxonomies",
		"timecreated",
		"timemodified",
		"usermodified",
		"visible",
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

	varCoreCompetencyCreateCompetencyFramework200Response := _CoreCompetencyCreateCompetencyFramework200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreCompetencyCreateCompetencyFramework200Response)

	if err != nil {
		return err
	}

	*o = CoreCompetencyCreateCompetencyFramework200Response(varCoreCompetencyCreateCompetencyFramework200Response)

	return err
}

type NullableCoreCompetencyCreateCompetencyFramework200Response struct {
	value *CoreCompetencyCreateCompetencyFramework200Response
	isSet bool
}

func (v NullableCoreCompetencyCreateCompetencyFramework200Response) Get() *CoreCompetencyCreateCompetencyFramework200Response {
	return v.value
}

func (v *NullableCoreCompetencyCreateCompetencyFramework200Response) Set(val *CoreCompetencyCreateCompetencyFramework200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCompetencyCreateCompetencyFramework200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCompetencyCreateCompetencyFramework200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCompetencyCreateCompetencyFramework200Response(val *CoreCompetencyCreateCompetencyFramework200Response) *NullableCoreCompetencyCreateCompetencyFramework200Response {
	return &NullableCoreCompetencyCreateCompetencyFramework200Response{value: val, isSet: true}
}

func (v NullableCoreCompetencyCreateCompetencyFramework200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCompetencyCreateCompetencyFramework200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


