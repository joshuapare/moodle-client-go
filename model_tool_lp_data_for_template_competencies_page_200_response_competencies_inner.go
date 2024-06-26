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

// checks if the ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner{}

// ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner struct for ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner
type ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner struct {
	Competency *CoreCompetencyCreateCompetency200Response `json:"competency,omitempty"`
	Comppath *ToolLpDataForCompetencySummary200ResponseComppath `json:"comppath,omitempty"`
	Framework *CoreCompetencyDuplicateCompetencyFramework200Response `json:"framework,omitempty"`
	// hascourses
	Hascourses *bool `json:"hascourses,omitempty"`
	// hasrelatedcompetencies
	Hasrelatedcompetencies *bool `json:"hasrelatedcompetencies,omitempty"`
	Linkedcourses []CoreCourseGetEnrolledCoursesByTimelineClassification200ResponseCoursesInner `json:"linkedcourses,omitempty"`
	// pluginbaseurl
	Pluginbaseurl *string `json:"pluginbaseurl,omitempty"`
	Relatedcompetencies []CoreCompetencyReadUserEvidence200ResponseCompetenciesInner `json:"relatedcompetencies,omitempty"`
	// scaleconfiguration
	Scaleconfiguration *string `json:"scaleconfiguration,omitempty"`
	// scaleid
	Scaleid *int32 `json:"scaleid,omitempty"`
	// taxonomyterm
	Taxonomyterm *string `json:"taxonomyterm,omitempty"`
}

// NewToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner instantiates a new ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner() *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner {
	this := ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner{}
	return &this
}

// NewToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInnerWithDefaults instantiates a new ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInnerWithDefaults() *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner {
	this := ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner{}
	return &this
}

// GetCompetency returns the Competency field value if set, zero value otherwise.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) GetCompetency() CoreCompetencyCreateCompetency200Response {
	if o == nil || IsNil(o.Competency) {
		var ret CoreCompetencyCreateCompetency200Response
		return ret
	}
	return *o.Competency
}

// GetCompetencyOk returns a tuple with the Competency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) GetCompetencyOk() (*CoreCompetencyCreateCompetency200Response, bool) {
	if o == nil || IsNil(o.Competency) {
		return nil, false
	}
	return o.Competency, true
}

// HasCompetency returns a boolean if a field has been set.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) HasCompetency() bool {
	if o != nil && !IsNil(o.Competency) {
		return true
	}

	return false
}

// SetCompetency gets a reference to the given CoreCompetencyCreateCompetency200Response and assigns it to the Competency field.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) SetCompetency(v CoreCompetencyCreateCompetency200Response) {
	o.Competency = &v
}

// GetComppath returns the Comppath field value if set, zero value otherwise.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) GetComppath() ToolLpDataForCompetencySummary200ResponseComppath {
	if o == nil || IsNil(o.Comppath) {
		var ret ToolLpDataForCompetencySummary200ResponseComppath
		return ret
	}
	return *o.Comppath
}

// GetComppathOk returns a tuple with the Comppath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) GetComppathOk() (*ToolLpDataForCompetencySummary200ResponseComppath, bool) {
	if o == nil || IsNil(o.Comppath) {
		return nil, false
	}
	return o.Comppath, true
}

// HasComppath returns a boolean if a field has been set.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) HasComppath() bool {
	if o != nil && !IsNil(o.Comppath) {
		return true
	}

	return false
}

// SetComppath gets a reference to the given ToolLpDataForCompetencySummary200ResponseComppath and assigns it to the Comppath field.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) SetComppath(v ToolLpDataForCompetencySummary200ResponseComppath) {
	o.Comppath = &v
}

// GetFramework returns the Framework field value if set, zero value otherwise.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) GetFramework() CoreCompetencyDuplicateCompetencyFramework200Response {
	if o == nil || IsNil(o.Framework) {
		var ret CoreCompetencyDuplicateCompetencyFramework200Response
		return ret
	}
	return *o.Framework
}

// GetFrameworkOk returns a tuple with the Framework field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) GetFrameworkOk() (*CoreCompetencyDuplicateCompetencyFramework200Response, bool) {
	if o == nil || IsNil(o.Framework) {
		return nil, false
	}
	return o.Framework, true
}

// HasFramework returns a boolean if a field has been set.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) HasFramework() bool {
	if o != nil && !IsNil(o.Framework) {
		return true
	}

	return false
}

// SetFramework gets a reference to the given CoreCompetencyDuplicateCompetencyFramework200Response and assigns it to the Framework field.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) SetFramework(v CoreCompetencyDuplicateCompetencyFramework200Response) {
	o.Framework = &v
}

// GetHascourses returns the Hascourses field value if set, zero value otherwise.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) GetHascourses() bool {
	if o == nil || IsNil(o.Hascourses) {
		var ret bool
		return ret
	}
	return *o.Hascourses
}

// GetHascoursesOk returns a tuple with the Hascourses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) GetHascoursesOk() (*bool, bool) {
	if o == nil || IsNil(o.Hascourses) {
		return nil, false
	}
	return o.Hascourses, true
}

// HasHascourses returns a boolean if a field has been set.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) HasHascourses() bool {
	if o != nil && !IsNil(o.Hascourses) {
		return true
	}

	return false
}

// SetHascourses gets a reference to the given bool and assigns it to the Hascourses field.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) SetHascourses(v bool) {
	o.Hascourses = &v
}

// GetHasrelatedcompetencies returns the Hasrelatedcompetencies field value if set, zero value otherwise.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) GetHasrelatedcompetencies() bool {
	if o == nil || IsNil(o.Hasrelatedcompetencies) {
		var ret bool
		return ret
	}
	return *o.Hasrelatedcompetencies
}

// GetHasrelatedcompetenciesOk returns a tuple with the Hasrelatedcompetencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) GetHasrelatedcompetenciesOk() (*bool, bool) {
	if o == nil || IsNil(o.Hasrelatedcompetencies) {
		return nil, false
	}
	return o.Hasrelatedcompetencies, true
}

// HasHasrelatedcompetencies returns a boolean if a field has been set.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) HasHasrelatedcompetencies() bool {
	if o != nil && !IsNil(o.Hasrelatedcompetencies) {
		return true
	}

	return false
}

// SetHasrelatedcompetencies gets a reference to the given bool and assigns it to the Hasrelatedcompetencies field.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) SetHasrelatedcompetencies(v bool) {
	o.Hasrelatedcompetencies = &v
}

// GetLinkedcourses returns the Linkedcourses field value if set, zero value otherwise.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) GetLinkedcourses() []CoreCourseGetEnrolledCoursesByTimelineClassification200ResponseCoursesInner {
	if o == nil || IsNil(o.Linkedcourses) {
		var ret []CoreCourseGetEnrolledCoursesByTimelineClassification200ResponseCoursesInner
		return ret
	}
	return o.Linkedcourses
}

// GetLinkedcoursesOk returns a tuple with the Linkedcourses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) GetLinkedcoursesOk() ([]CoreCourseGetEnrolledCoursesByTimelineClassification200ResponseCoursesInner, bool) {
	if o == nil || IsNil(o.Linkedcourses) {
		return nil, false
	}
	return o.Linkedcourses, true
}

// HasLinkedcourses returns a boolean if a field has been set.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) HasLinkedcourses() bool {
	if o != nil && !IsNil(o.Linkedcourses) {
		return true
	}

	return false
}

// SetLinkedcourses gets a reference to the given []CoreCourseGetEnrolledCoursesByTimelineClassification200ResponseCoursesInner and assigns it to the Linkedcourses field.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) SetLinkedcourses(v []CoreCourseGetEnrolledCoursesByTimelineClassification200ResponseCoursesInner) {
	o.Linkedcourses = v
}

// GetPluginbaseurl returns the Pluginbaseurl field value if set, zero value otherwise.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) GetPluginbaseurl() string {
	if o == nil || IsNil(o.Pluginbaseurl) {
		var ret string
		return ret
	}
	return *o.Pluginbaseurl
}

// GetPluginbaseurlOk returns a tuple with the Pluginbaseurl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) GetPluginbaseurlOk() (*string, bool) {
	if o == nil || IsNil(o.Pluginbaseurl) {
		return nil, false
	}
	return o.Pluginbaseurl, true
}

// HasPluginbaseurl returns a boolean if a field has been set.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) HasPluginbaseurl() bool {
	if o != nil && !IsNil(o.Pluginbaseurl) {
		return true
	}

	return false
}

// SetPluginbaseurl gets a reference to the given string and assigns it to the Pluginbaseurl field.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) SetPluginbaseurl(v string) {
	o.Pluginbaseurl = &v
}

// GetRelatedcompetencies returns the Relatedcompetencies field value if set, zero value otherwise.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) GetRelatedcompetencies() []CoreCompetencyReadUserEvidence200ResponseCompetenciesInner {
	if o == nil || IsNil(o.Relatedcompetencies) {
		var ret []CoreCompetencyReadUserEvidence200ResponseCompetenciesInner
		return ret
	}
	return o.Relatedcompetencies
}

// GetRelatedcompetenciesOk returns a tuple with the Relatedcompetencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) GetRelatedcompetenciesOk() ([]CoreCompetencyReadUserEvidence200ResponseCompetenciesInner, bool) {
	if o == nil || IsNil(o.Relatedcompetencies) {
		return nil, false
	}
	return o.Relatedcompetencies, true
}

// HasRelatedcompetencies returns a boolean if a field has been set.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) HasRelatedcompetencies() bool {
	if o != nil && !IsNil(o.Relatedcompetencies) {
		return true
	}

	return false
}

// SetRelatedcompetencies gets a reference to the given []CoreCompetencyReadUserEvidence200ResponseCompetenciesInner and assigns it to the Relatedcompetencies field.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) SetRelatedcompetencies(v []CoreCompetencyReadUserEvidence200ResponseCompetenciesInner) {
	o.Relatedcompetencies = v
}

// GetScaleconfiguration returns the Scaleconfiguration field value if set, zero value otherwise.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) GetScaleconfiguration() string {
	if o == nil || IsNil(o.Scaleconfiguration) {
		var ret string
		return ret
	}
	return *o.Scaleconfiguration
}

// GetScaleconfigurationOk returns a tuple with the Scaleconfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) GetScaleconfigurationOk() (*string, bool) {
	if o == nil || IsNil(o.Scaleconfiguration) {
		return nil, false
	}
	return o.Scaleconfiguration, true
}

// HasScaleconfiguration returns a boolean if a field has been set.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) HasScaleconfiguration() bool {
	if o != nil && !IsNil(o.Scaleconfiguration) {
		return true
	}

	return false
}

// SetScaleconfiguration gets a reference to the given string and assigns it to the Scaleconfiguration field.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) SetScaleconfiguration(v string) {
	o.Scaleconfiguration = &v
}

// GetScaleid returns the Scaleid field value if set, zero value otherwise.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) GetScaleid() int32 {
	if o == nil || IsNil(o.Scaleid) {
		var ret int32
		return ret
	}
	return *o.Scaleid
}

// GetScaleidOk returns a tuple with the Scaleid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) GetScaleidOk() (*int32, bool) {
	if o == nil || IsNil(o.Scaleid) {
		return nil, false
	}
	return o.Scaleid, true
}

// HasScaleid returns a boolean if a field has been set.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) HasScaleid() bool {
	if o != nil && !IsNil(o.Scaleid) {
		return true
	}

	return false
}

// SetScaleid gets a reference to the given int32 and assigns it to the Scaleid field.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) SetScaleid(v int32) {
	o.Scaleid = &v
}

// GetTaxonomyterm returns the Taxonomyterm field value if set, zero value otherwise.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) GetTaxonomyterm() string {
	if o == nil || IsNil(o.Taxonomyterm) {
		var ret string
		return ret
	}
	return *o.Taxonomyterm
}

// GetTaxonomytermOk returns a tuple with the Taxonomyterm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) GetTaxonomytermOk() (*string, bool) {
	if o == nil || IsNil(o.Taxonomyterm) {
		return nil, false
	}
	return o.Taxonomyterm, true
}

// HasTaxonomyterm returns a boolean if a field has been set.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) HasTaxonomyterm() bool {
	if o != nil && !IsNil(o.Taxonomyterm) {
		return true
	}

	return false
}

// SetTaxonomyterm gets a reference to the given string and assigns it to the Taxonomyterm field.
func (o *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) SetTaxonomyterm(v string) {
	o.Taxonomyterm = &v
}

func (o ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Competency) {
		toSerialize["competency"] = o.Competency
	}
	if !IsNil(o.Comppath) {
		toSerialize["comppath"] = o.Comppath
	}
	if !IsNil(o.Framework) {
		toSerialize["framework"] = o.Framework
	}
	if !IsNil(o.Hascourses) {
		toSerialize["hascourses"] = o.Hascourses
	}
	if !IsNil(o.Hasrelatedcompetencies) {
		toSerialize["hasrelatedcompetencies"] = o.Hasrelatedcompetencies
	}
	if !IsNil(o.Linkedcourses) {
		toSerialize["linkedcourses"] = o.Linkedcourses
	}
	if !IsNil(o.Pluginbaseurl) {
		toSerialize["pluginbaseurl"] = o.Pluginbaseurl
	}
	if !IsNil(o.Relatedcompetencies) {
		toSerialize["relatedcompetencies"] = o.Relatedcompetencies
	}
	if !IsNil(o.Scaleconfiguration) {
		toSerialize["scaleconfiguration"] = o.Scaleconfiguration
	}
	if !IsNil(o.Scaleid) {
		toSerialize["scaleid"] = o.Scaleid
	}
	if !IsNil(o.Taxonomyterm) {
		toSerialize["taxonomyterm"] = o.Taxonomyterm
	}
	return toSerialize, nil
}

type NullableToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner struct {
	value *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner
	isSet bool
}

func (v NullableToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) Get() *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner {
	return v.value
}

func (v *NullableToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) Set(val *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner(val *ToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) *NullableToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner {
	return &NullableToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner{value: val, isSet: true}
}

func (v NullableToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToolLpDataForTemplateCompetenciesPage200ResponseCompetenciesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


