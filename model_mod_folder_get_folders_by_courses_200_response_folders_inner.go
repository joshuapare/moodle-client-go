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

// checks if the ModFolderGetFoldersByCourses200ResponseFoldersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModFolderGetFoldersByCourses200ResponseFoldersInner{}

// ModFolderGetFoldersByCourses200ResponseFoldersInner struct for ModFolderGetFoldersByCourses200ResponseFoldersInner
type ModFolderGetFoldersByCourses200ResponseFoldersInner struct {
	// Course id
	Course *int32 `json:"course,omitempty"`
	// Course module id
	Coursemodule *int32 `json:"coursemodule,omitempty"`
	// Display type of folder contents on a separate page or inline
	Display *int32 `json:"display,omitempty"`
	// Whether file download is forced
	Forcedownload *int32 `json:"forcedownload,omitempty"`
	// Group id
	Groupingid *int32 `json:"groupingid,omitempty"`
	// Group mode
	Groupmode *int32 `json:"groupmode,omitempty"`
	// Activity instance id
	Id *int32 `json:"id,omitempty"`
	// Activity introduction
	Intro *string `json:"intro,omitempty"`
	Introfiles []CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner `json:"introfiles,omitempty"`
	// intro format (1 = HTML, 0 = MOODLE, 2 = PLAIN, or 4 = MARKDOWN)
	Introformat *int32 `json:"introformat,omitempty"`
	// Forced activity language
	Lang *string `json:"lang,omitempty"`
	// Activity name
	Name *string `json:"name,omitempty"`
	// Incremented when after each file changes, to avoid cache
	Revision *int32 `json:"revision,omitempty"`
	// Course section id
	Section *int32 `json:"section,omitempty"`
	// Whether to show the download folder button
	Showdownloadfolder *int32 `json:"showdownloadfolder,omitempty"`
	// 1 = expanded, 0 = collapsed for sub-folders
	Showexpanded *int32 `json:"showexpanded,omitempty"`
	// Last time the folder was modified
	Timemodified *int32 `json:"timemodified,omitempty"`
	// Visible
	Visible *bool `json:"visible,omitempty"`
}

// NewModFolderGetFoldersByCourses200ResponseFoldersInner instantiates a new ModFolderGetFoldersByCourses200ResponseFoldersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModFolderGetFoldersByCourses200ResponseFoldersInner() *ModFolderGetFoldersByCourses200ResponseFoldersInner {
	this := ModFolderGetFoldersByCourses200ResponseFoldersInner{}
	var display int32 = null
	this.Display = &display
	var forcedownload int32 = null
	this.Forcedownload = &forcedownload
	var revision int32 = null
	this.Revision = &revision
	var showdownloadfolder int32 = null
	this.Showdownloadfolder = &showdownloadfolder
	var showexpanded int32 = null
	this.Showexpanded = &showexpanded
	var timemodified int32 = null
	this.Timemodified = &timemodified
	return &this
}

// NewModFolderGetFoldersByCourses200ResponseFoldersInnerWithDefaults instantiates a new ModFolderGetFoldersByCourses200ResponseFoldersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModFolderGetFoldersByCourses200ResponseFoldersInnerWithDefaults() *ModFolderGetFoldersByCourses200ResponseFoldersInner {
	this := ModFolderGetFoldersByCourses200ResponseFoldersInner{}
	var display int32 = null
	this.Display = &display
	var forcedownload int32 = null
	this.Forcedownload = &forcedownload
	var revision int32 = null
	this.Revision = &revision
	var showdownloadfolder int32 = null
	this.Showdownloadfolder = &showdownloadfolder
	var showexpanded int32 = null
	this.Showexpanded = &showexpanded
	var timemodified int32 = null
	this.Timemodified = &timemodified
	return &this
}

// GetCourse returns the Course field value if set, zero value otherwise.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) GetCourse() int32 {
	if o == nil || IsNil(o.Course) {
		var ret int32
		return ret
	}
	return *o.Course
}

// GetCourseOk returns a tuple with the Course field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) GetCourseOk() (*int32, bool) {
	if o == nil || IsNil(o.Course) {
		return nil, false
	}
	return o.Course, true
}

// HasCourse returns a boolean if a field has been set.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) HasCourse() bool {
	if o != nil && !IsNil(o.Course) {
		return true
	}

	return false
}

// SetCourse gets a reference to the given int32 and assigns it to the Course field.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) SetCourse(v int32) {
	o.Course = &v
}

// GetCoursemodule returns the Coursemodule field value if set, zero value otherwise.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) GetCoursemodule() int32 {
	if o == nil || IsNil(o.Coursemodule) {
		var ret int32
		return ret
	}
	return *o.Coursemodule
}

// GetCoursemoduleOk returns a tuple with the Coursemodule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) GetCoursemoduleOk() (*int32, bool) {
	if o == nil || IsNil(o.Coursemodule) {
		return nil, false
	}
	return o.Coursemodule, true
}

// HasCoursemodule returns a boolean if a field has been set.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) HasCoursemodule() bool {
	if o != nil && !IsNil(o.Coursemodule) {
		return true
	}

	return false
}

// SetCoursemodule gets a reference to the given int32 and assigns it to the Coursemodule field.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) SetCoursemodule(v int32) {
	o.Coursemodule = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) GetDisplay() int32 {
	if o == nil || IsNil(o.Display) {
		var ret int32
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) GetDisplayOk() (*int32, bool) {
	if o == nil || IsNil(o.Display) {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) HasDisplay() bool {
	if o != nil && !IsNil(o.Display) {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given int32 and assigns it to the Display field.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) SetDisplay(v int32) {
	o.Display = &v
}

// GetForcedownload returns the Forcedownload field value if set, zero value otherwise.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) GetForcedownload() int32 {
	if o == nil || IsNil(o.Forcedownload) {
		var ret int32
		return ret
	}
	return *o.Forcedownload
}

// GetForcedownloadOk returns a tuple with the Forcedownload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) GetForcedownloadOk() (*int32, bool) {
	if o == nil || IsNil(o.Forcedownload) {
		return nil, false
	}
	return o.Forcedownload, true
}

// HasForcedownload returns a boolean if a field has been set.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) HasForcedownload() bool {
	if o != nil && !IsNil(o.Forcedownload) {
		return true
	}

	return false
}

// SetForcedownload gets a reference to the given int32 and assigns it to the Forcedownload field.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) SetForcedownload(v int32) {
	o.Forcedownload = &v
}

// GetGroupingid returns the Groupingid field value if set, zero value otherwise.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) GetGroupingid() int32 {
	if o == nil || IsNil(o.Groupingid) {
		var ret int32
		return ret
	}
	return *o.Groupingid
}

// GetGroupingidOk returns a tuple with the Groupingid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) GetGroupingidOk() (*int32, bool) {
	if o == nil || IsNil(o.Groupingid) {
		return nil, false
	}
	return o.Groupingid, true
}

// HasGroupingid returns a boolean if a field has been set.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) HasGroupingid() bool {
	if o != nil && !IsNil(o.Groupingid) {
		return true
	}

	return false
}

// SetGroupingid gets a reference to the given int32 and assigns it to the Groupingid field.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) SetGroupingid(v int32) {
	o.Groupingid = &v
}

// GetGroupmode returns the Groupmode field value if set, zero value otherwise.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) GetGroupmode() int32 {
	if o == nil || IsNil(o.Groupmode) {
		var ret int32
		return ret
	}
	return *o.Groupmode
}

// GetGroupmodeOk returns a tuple with the Groupmode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) GetGroupmodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Groupmode) {
		return nil, false
	}
	return o.Groupmode, true
}

// HasGroupmode returns a boolean if a field has been set.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) HasGroupmode() bool {
	if o != nil && !IsNil(o.Groupmode) {
		return true
	}

	return false
}

// SetGroupmode gets a reference to the given int32 and assigns it to the Groupmode field.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) SetGroupmode(v int32) {
	o.Groupmode = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) SetId(v int32) {
	o.Id = &v
}

// GetIntro returns the Intro field value if set, zero value otherwise.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) GetIntro() string {
	if o == nil || IsNil(o.Intro) {
		var ret string
		return ret
	}
	return *o.Intro
}

// GetIntroOk returns a tuple with the Intro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) GetIntroOk() (*string, bool) {
	if o == nil || IsNil(o.Intro) {
		return nil, false
	}
	return o.Intro, true
}

// HasIntro returns a boolean if a field has been set.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) HasIntro() bool {
	if o != nil && !IsNil(o.Intro) {
		return true
	}

	return false
}

// SetIntro gets a reference to the given string and assigns it to the Intro field.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) SetIntro(v string) {
	o.Intro = &v
}

// GetIntrofiles returns the Introfiles field value if set, zero value otherwise.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) GetIntrofiles() []CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner {
	if o == nil || IsNil(o.Introfiles) {
		var ret []CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner
		return ret
	}
	return o.Introfiles
}

// GetIntrofilesOk returns a tuple with the Introfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) GetIntrofilesOk() ([]CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner, bool) {
	if o == nil || IsNil(o.Introfiles) {
		return nil, false
	}
	return o.Introfiles, true
}

// HasIntrofiles returns a boolean if a field has been set.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) HasIntrofiles() bool {
	if o != nil && !IsNil(o.Introfiles) {
		return true
	}

	return false
}

// SetIntrofiles gets a reference to the given []CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner and assigns it to the Introfiles field.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) SetIntrofiles(v []CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) {
	o.Introfiles = v
}

// GetIntroformat returns the Introformat field value if set, zero value otherwise.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) GetIntroformat() int32 {
	if o == nil || IsNil(o.Introformat) {
		var ret int32
		return ret
	}
	return *o.Introformat
}

// GetIntroformatOk returns a tuple with the Introformat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) GetIntroformatOk() (*int32, bool) {
	if o == nil || IsNil(o.Introformat) {
		return nil, false
	}
	return o.Introformat, true
}

// HasIntroformat returns a boolean if a field has been set.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) HasIntroformat() bool {
	if o != nil && !IsNil(o.Introformat) {
		return true
	}

	return false
}

// SetIntroformat gets a reference to the given int32 and assigns it to the Introformat field.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) SetIntroformat(v int32) {
	o.Introformat = &v
}

// GetLang returns the Lang field value if set, zero value otherwise.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) GetLang() string {
	if o == nil || IsNil(o.Lang) {
		var ret string
		return ret
	}
	return *o.Lang
}

// GetLangOk returns a tuple with the Lang field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) GetLangOk() (*string, bool) {
	if o == nil || IsNil(o.Lang) {
		return nil, false
	}
	return o.Lang, true
}

// HasLang returns a boolean if a field has been set.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) HasLang() bool {
	if o != nil && !IsNil(o.Lang) {
		return true
	}

	return false
}

// SetLang gets a reference to the given string and assigns it to the Lang field.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) SetLang(v string) {
	o.Lang = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) SetName(v string) {
	o.Name = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) GetRevision() int32 {
	if o == nil || IsNil(o.Revision) {
		var ret int32
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) GetRevisionOk() (*int32, bool) {
	if o == nil || IsNil(o.Revision) {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) HasRevision() bool {
	if o != nil && !IsNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given int32 and assigns it to the Revision field.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) SetRevision(v int32) {
	o.Revision = &v
}

// GetSection returns the Section field value if set, zero value otherwise.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) GetSection() int32 {
	if o == nil || IsNil(o.Section) {
		var ret int32
		return ret
	}
	return *o.Section
}

// GetSectionOk returns a tuple with the Section field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) GetSectionOk() (*int32, bool) {
	if o == nil || IsNil(o.Section) {
		return nil, false
	}
	return o.Section, true
}

// HasSection returns a boolean if a field has been set.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) HasSection() bool {
	if o != nil && !IsNil(o.Section) {
		return true
	}

	return false
}

// SetSection gets a reference to the given int32 and assigns it to the Section field.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) SetSection(v int32) {
	o.Section = &v
}

// GetShowdownloadfolder returns the Showdownloadfolder field value if set, zero value otherwise.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) GetShowdownloadfolder() int32 {
	if o == nil || IsNil(o.Showdownloadfolder) {
		var ret int32
		return ret
	}
	return *o.Showdownloadfolder
}

// GetShowdownloadfolderOk returns a tuple with the Showdownloadfolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) GetShowdownloadfolderOk() (*int32, bool) {
	if o == nil || IsNil(o.Showdownloadfolder) {
		return nil, false
	}
	return o.Showdownloadfolder, true
}

// HasShowdownloadfolder returns a boolean if a field has been set.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) HasShowdownloadfolder() bool {
	if o != nil && !IsNil(o.Showdownloadfolder) {
		return true
	}

	return false
}

// SetShowdownloadfolder gets a reference to the given int32 and assigns it to the Showdownloadfolder field.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) SetShowdownloadfolder(v int32) {
	o.Showdownloadfolder = &v
}

// GetShowexpanded returns the Showexpanded field value if set, zero value otherwise.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) GetShowexpanded() int32 {
	if o == nil || IsNil(o.Showexpanded) {
		var ret int32
		return ret
	}
	return *o.Showexpanded
}

// GetShowexpandedOk returns a tuple with the Showexpanded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) GetShowexpandedOk() (*int32, bool) {
	if o == nil || IsNil(o.Showexpanded) {
		return nil, false
	}
	return o.Showexpanded, true
}

// HasShowexpanded returns a boolean if a field has been set.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) HasShowexpanded() bool {
	if o != nil && !IsNil(o.Showexpanded) {
		return true
	}

	return false
}

// SetShowexpanded gets a reference to the given int32 and assigns it to the Showexpanded field.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) SetShowexpanded(v int32) {
	o.Showexpanded = &v
}

// GetTimemodified returns the Timemodified field value if set, zero value otherwise.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) GetTimemodified() int32 {
	if o == nil || IsNil(o.Timemodified) {
		var ret int32
		return ret
	}
	return *o.Timemodified
}

// GetTimemodifiedOk returns a tuple with the Timemodified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) GetTimemodifiedOk() (*int32, bool) {
	if o == nil || IsNil(o.Timemodified) {
		return nil, false
	}
	return o.Timemodified, true
}

// HasTimemodified returns a boolean if a field has been set.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) HasTimemodified() bool {
	if o != nil && !IsNil(o.Timemodified) {
		return true
	}

	return false
}

// SetTimemodified gets a reference to the given int32 and assigns it to the Timemodified field.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) SetTimemodified(v int32) {
	o.Timemodified = &v
}

// GetVisible returns the Visible field value if set, zero value otherwise.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) GetVisible() bool {
	if o == nil || IsNil(o.Visible) {
		var ret bool
		return ret
	}
	return *o.Visible
}

// GetVisibleOk returns a tuple with the Visible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) GetVisibleOk() (*bool, bool) {
	if o == nil || IsNil(o.Visible) {
		return nil, false
	}
	return o.Visible, true
}

// HasVisible returns a boolean if a field has been set.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) HasVisible() bool {
	if o != nil && !IsNil(o.Visible) {
		return true
	}

	return false
}

// SetVisible gets a reference to the given bool and assigns it to the Visible field.
func (o *ModFolderGetFoldersByCourses200ResponseFoldersInner) SetVisible(v bool) {
	o.Visible = &v
}

func (o ModFolderGetFoldersByCourses200ResponseFoldersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModFolderGetFoldersByCourses200ResponseFoldersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Course) {
		toSerialize["course"] = o.Course
	}
	if !IsNil(o.Coursemodule) {
		toSerialize["coursemodule"] = o.Coursemodule
	}
	if !IsNil(o.Display) {
		toSerialize["display"] = o.Display
	}
	if !IsNil(o.Forcedownload) {
		toSerialize["forcedownload"] = o.Forcedownload
	}
	if !IsNil(o.Groupingid) {
		toSerialize["groupingid"] = o.Groupingid
	}
	if !IsNil(o.Groupmode) {
		toSerialize["groupmode"] = o.Groupmode
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Intro) {
		toSerialize["intro"] = o.Intro
	}
	if !IsNil(o.Introfiles) {
		toSerialize["introfiles"] = o.Introfiles
	}
	if !IsNil(o.Introformat) {
		toSerialize["introformat"] = o.Introformat
	}
	if !IsNil(o.Lang) {
		toSerialize["lang"] = o.Lang
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Revision) {
		toSerialize["revision"] = o.Revision
	}
	if !IsNil(o.Section) {
		toSerialize["section"] = o.Section
	}
	if !IsNil(o.Showdownloadfolder) {
		toSerialize["showdownloadfolder"] = o.Showdownloadfolder
	}
	if !IsNil(o.Showexpanded) {
		toSerialize["showexpanded"] = o.Showexpanded
	}
	if !IsNil(o.Timemodified) {
		toSerialize["timemodified"] = o.Timemodified
	}
	if !IsNil(o.Visible) {
		toSerialize["visible"] = o.Visible
	}
	return toSerialize, nil
}

type NullableModFolderGetFoldersByCourses200ResponseFoldersInner struct {
	value *ModFolderGetFoldersByCourses200ResponseFoldersInner
	isSet bool
}

func (v NullableModFolderGetFoldersByCourses200ResponseFoldersInner) Get() *ModFolderGetFoldersByCourses200ResponseFoldersInner {
	return v.value
}

func (v *NullableModFolderGetFoldersByCourses200ResponseFoldersInner) Set(val *ModFolderGetFoldersByCourses200ResponseFoldersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableModFolderGetFoldersByCourses200ResponseFoldersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableModFolderGetFoldersByCourses200ResponseFoldersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModFolderGetFoldersByCourses200ResponseFoldersInner(val *ModFolderGetFoldersByCourses200ResponseFoldersInner) *NullableModFolderGetFoldersByCourses200ResponseFoldersInner {
	return &NullableModFolderGetFoldersByCourses200ResponseFoldersInner{value: val, isSet: true}
}

func (v NullableModFolderGetFoldersByCourses200ResponseFoldersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModFolderGetFoldersByCourses200ResponseFoldersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

