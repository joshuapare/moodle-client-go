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

// checks if the CoreCourseSearchCourses200ResponseCoursesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCourseSearchCourses200ResponseCoursesInner{}

// CoreCourseSearchCourses200ResponseCoursesInner struct for CoreCourseSearchCourses200ResponseCoursesInner
type CoreCourseSearchCourses200ResponseCoursesInner struct {
	// category id
	Categoryid *int32 `json:"categoryid,omitempty"`
	// category name
	Categoryname *string `json:"categoryname,omitempty"`
	Contacts []CoreCourseSearchCourses200ResponseCoursesInnerContactsInner `json:"contacts,omitempty"`
	// Course image
	Courseimage *string `json:"courseimage,omitempty"`
	Customfields []CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner `json:"customfields,omitempty"`
	// course display name
	Displayname *string `json:"displayname,omitempty"`
	Enrollmentmethods []map[string]interface{} `json:"enrollmentmethods,omitempty"`
	// course full name
	Fullname *string `json:"fullname,omitempty"`
	// course id
	Id *int32 `json:"id,omitempty"`
	Overviewfiles []CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner `json:"overviewfiles,omitempty"`
	// course short name
	Shortname *string `json:"shortname,omitempty"`
	// Whether the activity dates are shown or not
	Showactivitydates *bool `json:"showactivitydates,omitempty"`
	// Whether the activity completion conditions are shown or not
	Showcompletionconditions *bool `json:"showcompletionconditions,omitempty"`
	// Sort order in the category
	Sortorder *int32 `json:"sortorder,omitempty"`
	// summary
	Summary *string `json:"summary,omitempty"`
	Summaryfiles []CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner `json:"summaryfiles,omitempty"`
	// summary format (1 = HTML, 0 = MOODLE, 2 = PLAIN, or 4 = MARKDOWN)
	Summaryformat *int32 `json:"summaryformat,omitempty"`
}

// NewCoreCourseSearchCourses200ResponseCoursesInner instantiates a new CoreCourseSearchCourses200ResponseCoursesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCourseSearchCourses200ResponseCoursesInner() *CoreCourseSearchCourses200ResponseCoursesInner {
	this := CoreCourseSearchCourses200ResponseCoursesInner{}
	return &this
}

// NewCoreCourseSearchCourses200ResponseCoursesInnerWithDefaults instantiates a new CoreCourseSearchCourses200ResponseCoursesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCourseSearchCourses200ResponseCoursesInnerWithDefaults() *CoreCourseSearchCourses200ResponseCoursesInner {
	this := CoreCourseSearchCourses200ResponseCoursesInner{}
	return &this
}

// GetCategoryid returns the Categoryid field value if set, zero value otherwise.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) GetCategoryid() int32 {
	if o == nil || IsNil(o.Categoryid) {
		var ret int32
		return ret
	}
	return *o.Categoryid
}

// GetCategoryidOk returns a tuple with the Categoryid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) GetCategoryidOk() (*int32, bool) {
	if o == nil || IsNil(o.Categoryid) {
		return nil, false
	}
	return o.Categoryid, true
}

// HasCategoryid returns a boolean if a field has been set.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) HasCategoryid() bool {
	if o != nil && !IsNil(o.Categoryid) {
		return true
	}

	return false
}

// SetCategoryid gets a reference to the given int32 and assigns it to the Categoryid field.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) SetCategoryid(v int32) {
	o.Categoryid = &v
}

// GetCategoryname returns the Categoryname field value if set, zero value otherwise.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) GetCategoryname() string {
	if o == nil || IsNil(o.Categoryname) {
		var ret string
		return ret
	}
	return *o.Categoryname
}

// GetCategorynameOk returns a tuple with the Categoryname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) GetCategorynameOk() (*string, bool) {
	if o == nil || IsNil(o.Categoryname) {
		return nil, false
	}
	return o.Categoryname, true
}

// HasCategoryname returns a boolean if a field has been set.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) HasCategoryname() bool {
	if o != nil && !IsNil(o.Categoryname) {
		return true
	}

	return false
}

// SetCategoryname gets a reference to the given string and assigns it to the Categoryname field.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) SetCategoryname(v string) {
	o.Categoryname = &v
}

// GetContacts returns the Contacts field value if set, zero value otherwise.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) GetContacts() []CoreCourseSearchCourses200ResponseCoursesInnerContactsInner {
	if o == nil || IsNil(o.Contacts) {
		var ret []CoreCourseSearchCourses200ResponseCoursesInnerContactsInner
		return ret
	}
	return o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) GetContactsOk() ([]CoreCourseSearchCourses200ResponseCoursesInnerContactsInner, bool) {
	if o == nil || IsNil(o.Contacts) {
		return nil, false
	}
	return o.Contacts, true
}

// HasContacts returns a boolean if a field has been set.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) HasContacts() bool {
	if o != nil && !IsNil(o.Contacts) {
		return true
	}

	return false
}

// SetContacts gets a reference to the given []CoreCourseSearchCourses200ResponseCoursesInnerContactsInner and assigns it to the Contacts field.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) SetContacts(v []CoreCourseSearchCourses200ResponseCoursesInnerContactsInner) {
	o.Contacts = v
}

// GetCourseimage returns the Courseimage field value if set, zero value otherwise.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) GetCourseimage() string {
	if o == nil || IsNil(o.Courseimage) {
		var ret string
		return ret
	}
	return *o.Courseimage
}

// GetCourseimageOk returns a tuple with the Courseimage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) GetCourseimageOk() (*string, bool) {
	if o == nil || IsNil(o.Courseimage) {
		return nil, false
	}
	return o.Courseimage, true
}

// HasCourseimage returns a boolean if a field has been set.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) HasCourseimage() bool {
	if o != nil && !IsNil(o.Courseimage) {
		return true
	}

	return false
}

// SetCourseimage gets a reference to the given string and assigns it to the Courseimage field.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) SetCourseimage(v string) {
	o.Courseimage = &v
}

// GetCustomfields returns the Customfields field value if set, zero value otherwise.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) GetCustomfields() []CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner {
	if o == nil || IsNil(o.Customfields) {
		var ret []CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner
		return ret
	}
	return o.Customfields
}

// GetCustomfieldsOk returns a tuple with the Customfields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) GetCustomfieldsOk() ([]CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner, bool) {
	if o == nil || IsNil(o.Customfields) {
		return nil, false
	}
	return o.Customfields, true
}

// HasCustomfields returns a boolean if a field has been set.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) HasCustomfields() bool {
	if o != nil && !IsNil(o.Customfields) {
		return true
	}

	return false
}

// SetCustomfields gets a reference to the given []CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner and assigns it to the Customfields field.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) SetCustomfields(v []CoreCourseGetCoursesByField200ResponseCoursesInnerCustomfieldsInner) {
	o.Customfields = v
}

// GetDisplayname returns the Displayname field value if set, zero value otherwise.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) GetDisplayname() string {
	if o == nil || IsNil(o.Displayname) {
		var ret string
		return ret
	}
	return *o.Displayname
}

// GetDisplaynameOk returns a tuple with the Displayname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) GetDisplaynameOk() (*string, bool) {
	if o == nil || IsNil(o.Displayname) {
		return nil, false
	}
	return o.Displayname, true
}

// HasDisplayname returns a boolean if a field has been set.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) HasDisplayname() bool {
	if o != nil && !IsNil(o.Displayname) {
		return true
	}

	return false
}

// SetDisplayname gets a reference to the given string and assigns it to the Displayname field.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) SetDisplayname(v string) {
	o.Displayname = &v
}

// GetEnrollmentmethods returns the Enrollmentmethods field value if set, zero value otherwise.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) GetEnrollmentmethods() []map[string]interface{} {
	if o == nil || IsNil(o.Enrollmentmethods) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Enrollmentmethods
}

// GetEnrollmentmethodsOk returns a tuple with the Enrollmentmethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) GetEnrollmentmethodsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Enrollmentmethods) {
		return nil, false
	}
	return o.Enrollmentmethods, true
}

// HasEnrollmentmethods returns a boolean if a field has been set.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) HasEnrollmentmethods() bool {
	if o != nil && !IsNil(o.Enrollmentmethods) {
		return true
	}

	return false
}

// SetEnrollmentmethods gets a reference to the given []map[string]interface{} and assigns it to the Enrollmentmethods field.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) SetEnrollmentmethods(v []map[string]interface{}) {
	o.Enrollmentmethods = v
}

// GetFullname returns the Fullname field value if set, zero value otherwise.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) GetFullname() string {
	if o == nil || IsNil(o.Fullname) {
		var ret string
		return ret
	}
	return *o.Fullname
}

// GetFullnameOk returns a tuple with the Fullname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) GetFullnameOk() (*string, bool) {
	if o == nil || IsNil(o.Fullname) {
		return nil, false
	}
	return o.Fullname, true
}

// HasFullname returns a boolean if a field has been set.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) HasFullname() bool {
	if o != nil && !IsNil(o.Fullname) {
		return true
	}

	return false
}

// SetFullname gets a reference to the given string and assigns it to the Fullname field.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) SetFullname(v string) {
	o.Fullname = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) SetId(v int32) {
	o.Id = &v
}

// GetOverviewfiles returns the Overviewfiles field value if set, zero value otherwise.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) GetOverviewfiles() []CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner {
	if o == nil || IsNil(o.Overviewfiles) {
		var ret []CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner
		return ret
	}
	return o.Overviewfiles
}

// GetOverviewfilesOk returns a tuple with the Overviewfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) GetOverviewfilesOk() ([]CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner, bool) {
	if o == nil || IsNil(o.Overviewfiles) {
		return nil, false
	}
	return o.Overviewfiles, true
}

// HasOverviewfiles returns a boolean if a field has been set.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) HasOverviewfiles() bool {
	if o != nil && !IsNil(o.Overviewfiles) {
		return true
	}

	return false
}

// SetOverviewfiles gets a reference to the given []CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner and assigns it to the Overviewfiles field.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) SetOverviewfiles(v []CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) {
	o.Overviewfiles = v
}

// GetShortname returns the Shortname field value if set, zero value otherwise.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) GetShortname() string {
	if o == nil || IsNil(o.Shortname) {
		var ret string
		return ret
	}
	return *o.Shortname
}

// GetShortnameOk returns a tuple with the Shortname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) GetShortnameOk() (*string, bool) {
	if o == nil || IsNil(o.Shortname) {
		return nil, false
	}
	return o.Shortname, true
}

// HasShortname returns a boolean if a field has been set.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) HasShortname() bool {
	if o != nil && !IsNil(o.Shortname) {
		return true
	}

	return false
}

// SetShortname gets a reference to the given string and assigns it to the Shortname field.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) SetShortname(v string) {
	o.Shortname = &v
}

// GetShowactivitydates returns the Showactivitydates field value if set, zero value otherwise.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) GetShowactivitydates() bool {
	if o == nil || IsNil(o.Showactivitydates) {
		var ret bool
		return ret
	}
	return *o.Showactivitydates
}

// GetShowactivitydatesOk returns a tuple with the Showactivitydates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) GetShowactivitydatesOk() (*bool, bool) {
	if o == nil || IsNil(o.Showactivitydates) {
		return nil, false
	}
	return o.Showactivitydates, true
}

// HasShowactivitydates returns a boolean if a field has been set.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) HasShowactivitydates() bool {
	if o != nil && !IsNil(o.Showactivitydates) {
		return true
	}

	return false
}

// SetShowactivitydates gets a reference to the given bool and assigns it to the Showactivitydates field.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) SetShowactivitydates(v bool) {
	o.Showactivitydates = &v
}

// GetShowcompletionconditions returns the Showcompletionconditions field value if set, zero value otherwise.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) GetShowcompletionconditions() bool {
	if o == nil || IsNil(o.Showcompletionconditions) {
		var ret bool
		return ret
	}
	return *o.Showcompletionconditions
}

// GetShowcompletionconditionsOk returns a tuple with the Showcompletionconditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) GetShowcompletionconditionsOk() (*bool, bool) {
	if o == nil || IsNil(o.Showcompletionconditions) {
		return nil, false
	}
	return o.Showcompletionconditions, true
}

// HasShowcompletionconditions returns a boolean if a field has been set.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) HasShowcompletionconditions() bool {
	if o != nil && !IsNil(o.Showcompletionconditions) {
		return true
	}

	return false
}

// SetShowcompletionconditions gets a reference to the given bool and assigns it to the Showcompletionconditions field.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) SetShowcompletionconditions(v bool) {
	o.Showcompletionconditions = &v
}

// GetSortorder returns the Sortorder field value if set, zero value otherwise.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) GetSortorder() int32 {
	if o == nil || IsNil(o.Sortorder) {
		var ret int32
		return ret
	}
	return *o.Sortorder
}

// GetSortorderOk returns a tuple with the Sortorder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) GetSortorderOk() (*int32, bool) {
	if o == nil || IsNil(o.Sortorder) {
		return nil, false
	}
	return o.Sortorder, true
}

// HasSortorder returns a boolean if a field has been set.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) HasSortorder() bool {
	if o != nil && !IsNil(o.Sortorder) {
		return true
	}

	return false
}

// SetSortorder gets a reference to the given int32 and assigns it to the Sortorder field.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) SetSortorder(v int32) {
	o.Sortorder = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) SetSummary(v string) {
	o.Summary = &v
}

// GetSummaryfiles returns the Summaryfiles field value if set, zero value otherwise.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) GetSummaryfiles() []CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner {
	if o == nil || IsNil(o.Summaryfiles) {
		var ret []CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner
		return ret
	}
	return o.Summaryfiles
}

// GetSummaryfilesOk returns a tuple with the Summaryfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) GetSummaryfilesOk() ([]CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner, bool) {
	if o == nil || IsNil(o.Summaryfiles) {
		return nil, false
	}
	return o.Summaryfiles, true
}

// HasSummaryfiles returns a boolean if a field has been set.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) HasSummaryfiles() bool {
	if o != nil && !IsNil(o.Summaryfiles) {
		return true
	}

	return false
}

// SetSummaryfiles gets a reference to the given []CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner and assigns it to the Summaryfiles field.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) SetSummaryfiles(v []CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) {
	o.Summaryfiles = v
}

// GetSummaryformat returns the Summaryformat field value if set, zero value otherwise.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) GetSummaryformat() int32 {
	if o == nil || IsNil(o.Summaryformat) {
		var ret int32
		return ret
	}
	return *o.Summaryformat
}

// GetSummaryformatOk returns a tuple with the Summaryformat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) GetSummaryformatOk() (*int32, bool) {
	if o == nil || IsNil(o.Summaryformat) {
		return nil, false
	}
	return o.Summaryformat, true
}

// HasSummaryformat returns a boolean if a field has been set.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) HasSummaryformat() bool {
	if o != nil && !IsNil(o.Summaryformat) {
		return true
	}

	return false
}

// SetSummaryformat gets a reference to the given int32 and assigns it to the Summaryformat field.
func (o *CoreCourseSearchCourses200ResponseCoursesInner) SetSummaryformat(v int32) {
	o.Summaryformat = &v
}

func (o CoreCourseSearchCourses200ResponseCoursesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCourseSearchCourses200ResponseCoursesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Categoryid) {
		toSerialize["categoryid"] = o.Categoryid
	}
	if !IsNil(o.Categoryname) {
		toSerialize["categoryname"] = o.Categoryname
	}
	if !IsNil(o.Contacts) {
		toSerialize["contacts"] = o.Contacts
	}
	if !IsNil(o.Courseimage) {
		toSerialize["courseimage"] = o.Courseimage
	}
	if !IsNil(o.Customfields) {
		toSerialize["customfields"] = o.Customfields
	}
	if !IsNil(o.Displayname) {
		toSerialize["displayname"] = o.Displayname
	}
	if !IsNil(o.Enrollmentmethods) {
		toSerialize["enrollmentmethods"] = o.Enrollmentmethods
	}
	if !IsNil(o.Fullname) {
		toSerialize["fullname"] = o.Fullname
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Overviewfiles) {
		toSerialize["overviewfiles"] = o.Overviewfiles
	}
	if !IsNil(o.Shortname) {
		toSerialize["shortname"] = o.Shortname
	}
	if !IsNil(o.Showactivitydates) {
		toSerialize["showactivitydates"] = o.Showactivitydates
	}
	if !IsNil(o.Showcompletionconditions) {
		toSerialize["showcompletionconditions"] = o.Showcompletionconditions
	}
	if !IsNil(o.Sortorder) {
		toSerialize["sortorder"] = o.Sortorder
	}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	if !IsNil(o.Summaryfiles) {
		toSerialize["summaryfiles"] = o.Summaryfiles
	}
	if !IsNil(o.Summaryformat) {
		toSerialize["summaryformat"] = o.Summaryformat
	}
	return toSerialize, nil
}

type NullableCoreCourseSearchCourses200ResponseCoursesInner struct {
	value *CoreCourseSearchCourses200ResponseCoursesInner
	isSet bool
}

func (v NullableCoreCourseSearchCourses200ResponseCoursesInner) Get() *CoreCourseSearchCourses200ResponseCoursesInner {
	return v.value
}

func (v *NullableCoreCourseSearchCourses200ResponseCoursesInner) Set(val *CoreCourseSearchCourses200ResponseCoursesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCourseSearchCourses200ResponseCoursesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCourseSearchCourses200ResponseCoursesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCourseSearchCourses200ResponseCoursesInner(val *CoreCourseSearchCourses200ResponseCoursesInner) *NullableCoreCourseSearchCourses200ResponseCoursesInner {
	return &NullableCoreCourseSearchCourses200ResponseCoursesInner{value: val, isSet: true}
}

func (v NullableCoreCourseSearchCourses200ResponseCoursesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCourseSearchCourses200ResponseCoursesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

