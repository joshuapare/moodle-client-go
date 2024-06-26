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

// checks if the CoreCompetencyReadUserEvidence200ResponseCompetenciesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCompetencyReadUserEvidence200ResponseCompetenciesInner{}

// CoreCompetencyReadUserEvidence200ResponseCompetenciesInner struct for CoreCompetencyReadUserEvidence200ResponseCompetenciesInner
type CoreCompetencyReadUserEvidence200ResponseCompetenciesInner struct {
	// competencyframeworkid
	Competencyframeworkid *int32 `json:"competencyframeworkid,omitempty"`
	// description
	Description *string `json:"description,omitempty"`
	// description format (1 = HTML, 0 = MOODLE, 2 = PLAIN, or 4 = MARKDOWN)
	Descriptionformat *int32 `json:"descriptionformat,omitempty"`
	// id
	Id *int32 `json:"id,omitempty"`
	// idnumber
	Idnumber *string `json:"idnumber,omitempty"`
	// parentid
	Parentid *int32 `json:"parentid,omitempty"`
	// path
	Path *string `json:"path,omitempty"`
	// ruleconfig
	Ruleconfig *string `json:"ruleconfig,omitempty"`
	// ruleoutcome
	Ruleoutcome *int32 `json:"ruleoutcome,omitempty"`
	// ruletype
	Ruletype *string `json:"ruletype,omitempty"`
	// scaleconfiguration
	Scaleconfiguration *string `json:"scaleconfiguration,omitempty"`
	// scaleid
	Scaleid *int32 `json:"scaleid,omitempty"`
	// shortname
	Shortname *string `json:"shortname,omitempty"`
	// sortorder
	Sortorder *int32 `json:"sortorder,omitempty"`
	// timecreated
	Timecreated *int32 `json:"timecreated,omitempty"`
	// timemodified
	Timemodified *int32 `json:"timemodified,omitempty"`
	// usermodified
	Usermodified *int32 `json:"usermodified,omitempty"`
}

// NewCoreCompetencyReadUserEvidence200ResponseCompetenciesInner instantiates a new CoreCompetencyReadUserEvidence200ResponseCompetenciesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCompetencyReadUserEvidence200ResponseCompetenciesInner() *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner {
	this := CoreCompetencyReadUserEvidence200ResponseCompetenciesInner{}
	var competencyframeworkid int32 = 0
	this.Competencyframeworkid = &competencyframeworkid
	var description string = ""
	this.Description = &description
	var descriptionformat int32 = 1
	this.Descriptionformat = &descriptionformat
	var id int32 = 0
	this.Id = &id
	var parentid int32 = 0
	this.Parentid = &parentid
	var path string = "/0/"
	this.Path = &path
	var ruleoutcome int32 = 0
	this.Ruleoutcome = &ruleoutcome
	var sortorder int32 = 0
	this.Sortorder = &sortorder
	var timecreated int32 = 0
	this.Timecreated = &timecreated
	var timemodified int32 = 0
	this.Timemodified = &timemodified
	var usermodified int32 = 0
	this.Usermodified = &usermodified
	return &this
}

// NewCoreCompetencyReadUserEvidence200ResponseCompetenciesInnerWithDefaults instantiates a new CoreCompetencyReadUserEvidence200ResponseCompetenciesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCompetencyReadUserEvidence200ResponseCompetenciesInnerWithDefaults() *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner {
	this := CoreCompetencyReadUserEvidence200ResponseCompetenciesInner{}
	var competencyframeworkid int32 = 0
	this.Competencyframeworkid = &competencyframeworkid
	var description string = ""
	this.Description = &description
	var descriptionformat int32 = 1
	this.Descriptionformat = &descriptionformat
	var id int32 = 0
	this.Id = &id
	var parentid int32 = 0
	this.Parentid = &parentid
	var path string = "/0/"
	this.Path = &path
	var ruleoutcome int32 = 0
	this.Ruleoutcome = &ruleoutcome
	var sortorder int32 = 0
	this.Sortorder = &sortorder
	var timecreated int32 = 0
	this.Timecreated = &timecreated
	var timemodified int32 = 0
	this.Timemodified = &timemodified
	var usermodified int32 = 0
	this.Usermodified = &usermodified
	return &this
}

// GetCompetencyframeworkid returns the Competencyframeworkid field value if set, zero value otherwise.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) GetCompetencyframeworkid() int32 {
	if o == nil || IsNil(o.Competencyframeworkid) {
		var ret int32
		return ret
	}
	return *o.Competencyframeworkid
}

// GetCompetencyframeworkidOk returns a tuple with the Competencyframeworkid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) GetCompetencyframeworkidOk() (*int32, bool) {
	if o == nil || IsNil(o.Competencyframeworkid) {
		return nil, false
	}
	return o.Competencyframeworkid, true
}

// HasCompetencyframeworkid returns a boolean if a field has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) HasCompetencyframeworkid() bool {
	if o != nil && !IsNil(o.Competencyframeworkid) {
		return true
	}

	return false
}

// SetCompetencyframeworkid gets a reference to the given int32 and assigns it to the Competencyframeworkid field.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) SetCompetencyframeworkid(v int32) {
	o.Competencyframeworkid = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) SetDescription(v string) {
	o.Description = &v
}

// GetDescriptionformat returns the Descriptionformat field value if set, zero value otherwise.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) GetDescriptionformat() int32 {
	if o == nil || IsNil(o.Descriptionformat) {
		var ret int32
		return ret
	}
	return *o.Descriptionformat
}

// GetDescriptionformatOk returns a tuple with the Descriptionformat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) GetDescriptionformatOk() (*int32, bool) {
	if o == nil || IsNil(o.Descriptionformat) {
		return nil, false
	}
	return o.Descriptionformat, true
}

// HasDescriptionformat returns a boolean if a field has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) HasDescriptionformat() bool {
	if o != nil && !IsNil(o.Descriptionformat) {
		return true
	}

	return false
}

// SetDescriptionformat gets a reference to the given int32 and assigns it to the Descriptionformat field.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) SetDescriptionformat(v int32) {
	o.Descriptionformat = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) SetId(v int32) {
	o.Id = &v
}

// GetIdnumber returns the Idnumber field value if set, zero value otherwise.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) GetIdnumber() string {
	if o == nil || IsNil(o.Idnumber) {
		var ret string
		return ret
	}
	return *o.Idnumber
}

// GetIdnumberOk returns a tuple with the Idnumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) GetIdnumberOk() (*string, bool) {
	if o == nil || IsNil(o.Idnumber) {
		return nil, false
	}
	return o.Idnumber, true
}

// HasIdnumber returns a boolean if a field has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) HasIdnumber() bool {
	if o != nil && !IsNil(o.Idnumber) {
		return true
	}

	return false
}

// SetIdnumber gets a reference to the given string and assigns it to the Idnumber field.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) SetIdnumber(v string) {
	o.Idnumber = &v
}

// GetParentid returns the Parentid field value if set, zero value otherwise.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) GetParentid() int32 {
	if o == nil || IsNil(o.Parentid) {
		var ret int32
		return ret
	}
	return *o.Parentid
}

// GetParentidOk returns a tuple with the Parentid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) GetParentidOk() (*int32, bool) {
	if o == nil || IsNil(o.Parentid) {
		return nil, false
	}
	return o.Parentid, true
}

// HasParentid returns a boolean if a field has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) HasParentid() bool {
	if o != nil && !IsNil(o.Parentid) {
		return true
	}

	return false
}

// SetParentid gets a reference to the given int32 and assigns it to the Parentid field.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) SetParentid(v int32) {
	o.Parentid = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) SetPath(v string) {
	o.Path = &v
}

// GetRuleconfig returns the Ruleconfig field value if set, zero value otherwise.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) GetRuleconfig() string {
	if o == nil || IsNil(o.Ruleconfig) {
		var ret string
		return ret
	}
	return *o.Ruleconfig
}

// GetRuleconfigOk returns a tuple with the Ruleconfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) GetRuleconfigOk() (*string, bool) {
	if o == nil || IsNil(o.Ruleconfig) {
		return nil, false
	}
	return o.Ruleconfig, true
}

// HasRuleconfig returns a boolean if a field has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) HasRuleconfig() bool {
	if o != nil && !IsNil(o.Ruleconfig) {
		return true
	}

	return false
}

// SetRuleconfig gets a reference to the given string and assigns it to the Ruleconfig field.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) SetRuleconfig(v string) {
	o.Ruleconfig = &v
}

// GetRuleoutcome returns the Ruleoutcome field value if set, zero value otherwise.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) GetRuleoutcome() int32 {
	if o == nil || IsNil(o.Ruleoutcome) {
		var ret int32
		return ret
	}
	return *o.Ruleoutcome
}

// GetRuleoutcomeOk returns a tuple with the Ruleoutcome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) GetRuleoutcomeOk() (*int32, bool) {
	if o == nil || IsNil(o.Ruleoutcome) {
		return nil, false
	}
	return o.Ruleoutcome, true
}

// HasRuleoutcome returns a boolean if a field has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) HasRuleoutcome() bool {
	if o != nil && !IsNil(o.Ruleoutcome) {
		return true
	}

	return false
}

// SetRuleoutcome gets a reference to the given int32 and assigns it to the Ruleoutcome field.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) SetRuleoutcome(v int32) {
	o.Ruleoutcome = &v
}

// GetRuletype returns the Ruletype field value if set, zero value otherwise.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) GetRuletype() string {
	if o == nil || IsNil(o.Ruletype) {
		var ret string
		return ret
	}
	return *o.Ruletype
}

// GetRuletypeOk returns a tuple with the Ruletype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) GetRuletypeOk() (*string, bool) {
	if o == nil || IsNil(o.Ruletype) {
		return nil, false
	}
	return o.Ruletype, true
}

// HasRuletype returns a boolean if a field has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) HasRuletype() bool {
	if o != nil && !IsNil(o.Ruletype) {
		return true
	}

	return false
}

// SetRuletype gets a reference to the given string and assigns it to the Ruletype field.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) SetRuletype(v string) {
	o.Ruletype = &v
}

// GetScaleconfiguration returns the Scaleconfiguration field value if set, zero value otherwise.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) GetScaleconfiguration() string {
	if o == nil || IsNil(o.Scaleconfiguration) {
		var ret string
		return ret
	}
	return *o.Scaleconfiguration
}

// GetScaleconfigurationOk returns a tuple with the Scaleconfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) GetScaleconfigurationOk() (*string, bool) {
	if o == nil || IsNil(o.Scaleconfiguration) {
		return nil, false
	}
	return o.Scaleconfiguration, true
}

// HasScaleconfiguration returns a boolean if a field has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) HasScaleconfiguration() bool {
	if o != nil && !IsNil(o.Scaleconfiguration) {
		return true
	}

	return false
}

// SetScaleconfiguration gets a reference to the given string and assigns it to the Scaleconfiguration field.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) SetScaleconfiguration(v string) {
	o.Scaleconfiguration = &v
}

// GetScaleid returns the Scaleid field value if set, zero value otherwise.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) GetScaleid() int32 {
	if o == nil || IsNil(o.Scaleid) {
		var ret int32
		return ret
	}
	return *o.Scaleid
}

// GetScaleidOk returns a tuple with the Scaleid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) GetScaleidOk() (*int32, bool) {
	if o == nil || IsNil(o.Scaleid) {
		return nil, false
	}
	return o.Scaleid, true
}

// HasScaleid returns a boolean if a field has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) HasScaleid() bool {
	if o != nil && !IsNil(o.Scaleid) {
		return true
	}

	return false
}

// SetScaleid gets a reference to the given int32 and assigns it to the Scaleid field.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) SetScaleid(v int32) {
	o.Scaleid = &v
}

// GetShortname returns the Shortname field value if set, zero value otherwise.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) GetShortname() string {
	if o == nil || IsNil(o.Shortname) {
		var ret string
		return ret
	}
	return *o.Shortname
}

// GetShortnameOk returns a tuple with the Shortname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) GetShortnameOk() (*string, bool) {
	if o == nil || IsNil(o.Shortname) {
		return nil, false
	}
	return o.Shortname, true
}

// HasShortname returns a boolean if a field has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) HasShortname() bool {
	if o != nil && !IsNil(o.Shortname) {
		return true
	}

	return false
}

// SetShortname gets a reference to the given string and assigns it to the Shortname field.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) SetShortname(v string) {
	o.Shortname = &v
}

// GetSortorder returns the Sortorder field value if set, zero value otherwise.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) GetSortorder() int32 {
	if o == nil || IsNil(o.Sortorder) {
		var ret int32
		return ret
	}
	return *o.Sortorder
}

// GetSortorderOk returns a tuple with the Sortorder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) GetSortorderOk() (*int32, bool) {
	if o == nil || IsNil(o.Sortorder) {
		return nil, false
	}
	return o.Sortorder, true
}

// HasSortorder returns a boolean if a field has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) HasSortorder() bool {
	if o != nil && !IsNil(o.Sortorder) {
		return true
	}

	return false
}

// SetSortorder gets a reference to the given int32 and assigns it to the Sortorder field.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) SetSortorder(v int32) {
	o.Sortorder = &v
}

// GetTimecreated returns the Timecreated field value if set, zero value otherwise.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) GetTimecreated() int32 {
	if o == nil || IsNil(o.Timecreated) {
		var ret int32
		return ret
	}
	return *o.Timecreated
}

// GetTimecreatedOk returns a tuple with the Timecreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) GetTimecreatedOk() (*int32, bool) {
	if o == nil || IsNil(o.Timecreated) {
		return nil, false
	}
	return o.Timecreated, true
}

// HasTimecreated returns a boolean if a field has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) HasTimecreated() bool {
	if o != nil && !IsNil(o.Timecreated) {
		return true
	}

	return false
}

// SetTimecreated gets a reference to the given int32 and assigns it to the Timecreated field.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) SetTimecreated(v int32) {
	o.Timecreated = &v
}

// GetTimemodified returns the Timemodified field value if set, zero value otherwise.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) GetTimemodified() int32 {
	if o == nil || IsNil(o.Timemodified) {
		var ret int32
		return ret
	}
	return *o.Timemodified
}

// GetTimemodifiedOk returns a tuple with the Timemodified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) GetTimemodifiedOk() (*int32, bool) {
	if o == nil || IsNil(o.Timemodified) {
		return nil, false
	}
	return o.Timemodified, true
}

// HasTimemodified returns a boolean if a field has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) HasTimemodified() bool {
	if o != nil && !IsNil(o.Timemodified) {
		return true
	}

	return false
}

// SetTimemodified gets a reference to the given int32 and assigns it to the Timemodified field.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) SetTimemodified(v int32) {
	o.Timemodified = &v
}

// GetUsermodified returns the Usermodified field value if set, zero value otherwise.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) GetUsermodified() int32 {
	if o == nil || IsNil(o.Usermodified) {
		var ret int32
		return ret
	}
	return *o.Usermodified
}

// GetUsermodifiedOk returns a tuple with the Usermodified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) GetUsermodifiedOk() (*int32, bool) {
	if o == nil || IsNil(o.Usermodified) {
		return nil, false
	}
	return o.Usermodified, true
}

// HasUsermodified returns a boolean if a field has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) HasUsermodified() bool {
	if o != nil && !IsNil(o.Usermodified) {
		return true
	}

	return false
}

// SetUsermodified gets a reference to the given int32 and assigns it to the Usermodified field.
func (o *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) SetUsermodified(v int32) {
	o.Usermodified = &v
}

func (o CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Competencyframeworkid) {
		toSerialize["competencyframeworkid"] = o.Competencyframeworkid
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Descriptionformat) {
		toSerialize["descriptionformat"] = o.Descriptionformat
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Idnumber) {
		toSerialize["idnumber"] = o.Idnumber
	}
	if !IsNil(o.Parentid) {
		toSerialize["parentid"] = o.Parentid
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Ruleconfig) {
		toSerialize["ruleconfig"] = o.Ruleconfig
	}
	if !IsNil(o.Ruleoutcome) {
		toSerialize["ruleoutcome"] = o.Ruleoutcome
	}
	if !IsNil(o.Ruletype) {
		toSerialize["ruletype"] = o.Ruletype
	}
	if !IsNil(o.Scaleconfiguration) {
		toSerialize["scaleconfiguration"] = o.Scaleconfiguration
	}
	if !IsNil(o.Scaleid) {
		toSerialize["scaleid"] = o.Scaleid
	}
	if !IsNil(o.Shortname) {
		toSerialize["shortname"] = o.Shortname
	}
	if !IsNil(o.Sortorder) {
		toSerialize["sortorder"] = o.Sortorder
	}
	if !IsNil(o.Timecreated) {
		toSerialize["timecreated"] = o.Timecreated
	}
	if !IsNil(o.Timemodified) {
		toSerialize["timemodified"] = o.Timemodified
	}
	if !IsNil(o.Usermodified) {
		toSerialize["usermodified"] = o.Usermodified
	}
	return toSerialize, nil
}

type NullableCoreCompetencyReadUserEvidence200ResponseCompetenciesInner struct {
	value *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner
	isSet bool
}

func (v NullableCoreCompetencyReadUserEvidence200ResponseCompetenciesInner) Get() *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner {
	return v.value
}

func (v *NullableCoreCompetencyReadUserEvidence200ResponseCompetenciesInner) Set(val *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCompetencyReadUserEvidence200ResponseCompetenciesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCompetencyReadUserEvidence200ResponseCompetenciesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCompetencyReadUserEvidence200ResponseCompetenciesInner(val *CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) *NullableCoreCompetencyReadUserEvidence200ResponseCompetenciesInner {
	return &NullableCoreCompetencyReadUserEvidence200ResponseCompetenciesInner{value: val, isSet: true}
}

func (v NullableCoreCompetencyReadUserEvidence200ResponseCompetenciesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCompetencyReadUserEvidence200ResponseCompetenciesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


