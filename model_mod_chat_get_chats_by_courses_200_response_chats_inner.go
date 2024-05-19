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

// checks if the ModChatGetChatsByCourses200ResponseChatsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModChatGetChatsByCourses200ResponseChatsInner{}

// ModChatGetChatsByCourses200ResponseChatsInner Chats
type ModChatGetChatsByCourses200ResponseChatsInner struct {
	// chat method (sockets, ajax, header_js)
	Chatmethod *string `json:"chatmethod,omitempty"`
	// chat time
	Chattime *int32 `json:"chattime,omitempty"`
	// Course id
	Course *int32 `json:"course,omitempty"`
	// Course module id
	Coursemodule *int32 `json:"coursemodule,omitempty"`
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
	// keep days
	Keepdays *int32 `json:"keepdays,omitempty"`
	// Forced activity language
	Lang *string `json:"lang,omitempty"`
	// Activity name
	Name *string `json:"name,omitempty"`
	// schedule type
	Schedule *int32 `json:"schedule,omitempty"`
	// Course section id
	Section *int32 `json:"section,omitempty"`
	// student logs visible to everyone
	Studentlogs *int32 `json:"studentlogs,omitempty"`
	// time of last modification
	Timemodified *int32 `json:"timemodified,omitempty"`
	// Visible
	Visible *bool `json:"visible,omitempty"`
}

// NewModChatGetChatsByCourses200ResponseChatsInner instantiates a new ModChatGetChatsByCourses200ResponseChatsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModChatGetChatsByCourses200ResponseChatsInner() *ModChatGetChatsByCourses200ResponseChatsInner {
	this := ModChatGetChatsByCourses200ResponseChatsInner{}
	var chatmethod string = "null"
	this.Chatmethod = &chatmethod
	var chattime int32 = null
	this.Chattime = &chattime
	var keepdays int32 = null
	this.Keepdays = &keepdays
	var schedule int32 = null
	this.Schedule = &schedule
	var studentlogs int32 = null
	this.Studentlogs = &studentlogs
	var timemodified int32 = null
	this.Timemodified = &timemodified
	return &this
}

// NewModChatGetChatsByCourses200ResponseChatsInnerWithDefaults instantiates a new ModChatGetChatsByCourses200ResponseChatsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModChatGetChatsByCourses200ResponseChatsInnerWithDefaults() *ModChatGetChatsByCourses200ResponseChatsInner {
	this := ModChatGetChatsByCourses200ResponseChatsInner{}
	var chatmethod string = "null"
	this.Chatmethod = &chatmethod
	var chattime int32 = null
	this.Chattime = &chattime
	var keepdays int32 = null
	this.Keepdays = &keepdays
	var schedule int32 = null
	this.Schedule = &schedule
	var studentlogs int32 = null
	this.Studentlogs = &studentlogs
	var timemodified int32 = null
	this.Timemodified = &timemodified
	return &this
}

// GetChatmethod returns the Chatmethod field value if set, zero value otherwise.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) GetChatmethod() string {
	if o == nil || IsNil(o.Chatmethod) {
		var ret string
		return ret
	}
	return *o.Chatmethod
}

// GetChatmethodOk returns a tuple with the Chatmethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) GetChatmethodOk() (*string, bool) {
	if o == nil || IsNil(o.Chatmethod) {
		return nil, false
	}
	return o.Chatmethod, true
}

// HasChatmethod returns a boolean if a field has been set.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) HasChatmethod() bool {
	if o != nil && !IsNil(o.Chatmethod) {
		return true
	}

	return false
}

// SetChatmethod gets a reference to the given string and assigns it to the Chatmethod field.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) SetChatmethod(v string) {
	o.Chatmethod = &v
}

// GetChattime returns the Chattime field value if set, zero value otherwise.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) GetChattime() int32 {
	if o == nil || IsNil(o.Chattime) {
		var ret int32
		return ret
	}
	return *o.Chattime
}

// GetChattimeOk returns a tuple with the Chattime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) GetChattimeOk() (*int32, bool) {
	if o == nil || IsNil(o.Chattime) {
		return nil, false
	}
	return o.Chattime, true
}

// HasChattime returns a boolean if a field has been set.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) HasChattime() bool {
	if o != nil && !IsNil(o.Chattime) {
		return true
	}

	return false
}

// SetChattime gets a reference to the given int32 and assigns it to the Chattime field.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) SetChattime(v int32) {
	o.Chattime = &v
}

// GetCourse returns the Course field value if set, zero value otherwise.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) GetCourse() int32 {
	if o == nil || IsNil(o.Course) {
		var ret int32
		return ret
	}
	return *o.Course
}

// GetCourseOk returns a tuple with the Course field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) GetCourseOk() (*int32, bool) {
	if o == nil || IsNil(o.Course) {
		return nil, false
	}
	return o.Course, true
}

// HasCourse returns a boolean if a field has been set.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) HasCourse() bool {
	if o != nil && !IsNil(o.Course) {
		return true
	}

	return false
}

// SetCourse gets a reference to the given int32 and assigns it to the Course field.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) SetCourse(v int32) {
	o.Course = &v
}

// GetCoursemodule returns the Coursemodule field value if set, zero value otherwise.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) GetCoursemodule() int32 {
	if o == nil || IsNil(o.Coursemodule) {
		var ret int32
		return ret
	}
	return *o.Coursemodule
}

// GetCoursemoduleOk returns a tuple with the Coursemodule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) GetCoursemoduleOk() (*int32, bool) {
	if o == nil || IsNil(o.Coursemodule) {
		return nil, false
	}
	return o.Coursemodule, true
}

// HasCoursemodule returns a boolean if a field has been set.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) HasCoursemodule() bool {
	if o != nil && !IsNil(o.Coursemodule) {
		return true
	}

	return false
}

// SetCoursemodule gets a reference to the given int32 and assigns it to the Coursemodule field.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) SetCoursemodule(v int32) {
	o.Coursemodule = &v
}

// GetGroupingid returns the Groupingid field value if set, zero value otherwise.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) GetGroupingid() int32 {
	if o == nil || IsNil(o.Groupingid) {
		var ret int32
		return ret
	}
	return *o.Groupingid
}

// GetGroupingidOk returns a tuple with the Groupingid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) GetGroupingidOk() (*int32, bool) {
	if o == nil || IsNil(o.Groupingid) {
		return nil, false
	}
	return o.Groupingid, true
}

// HasGroupingid returns a boolean if a field has been set.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) HasGroupingid() bool {
	if o != nil && !IsNil(o.Groupingid) {
		return true
	}

	return false
}

// SetGroupingid gets a reference to the given int32 and assigns it to the Groupingid field.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) SetGroupingid(v int32) {
	o.Groupingid = &v
}

// GetGroupmode returns the Groupmode field value if set, zero value otherwise.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) GetGroupmode() int32 {
	if o == nil || IsNil(o.Groupmode) {
		var ret int32
		return ret
	}
	return *o.Groupmode
}

// GetGroupmodeOk returns a tuple with the Groupmode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) GetGroupmodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Groupmode) {
		return nil, false
	}
	return o.Groupmode, true
}

// HasGroupmode returns a boolean if a field has been set.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) HasGroupmode() bool {
	if o != nil && !IsNil(o.Groupmode) {
		return true
	}

	return false
}

// SetGroupmode gets a reference to the given int32 and assigns it to the Groupmode field.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) SetGroupmode(v int32) {
	o.Groupmode = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) SetId(v int32) {
	o.Id = &v
}

// GetIntro returns the Intro field value if set, zero value otherwise.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) GetIntro() string {
	if o == nil || IsNil(o.Intro) {
		var ret string
		return ret
	}
	return *o.Intro
}

// GetIntroOk returns a tuple with the Intro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) GetIntroOk() (*string, bool) {
	if o == nil || IsNil(o.Intro) {
		return nil, false
	}
	return o.Intro, true
}

// HasIntro returns a boolean if a field has been set.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) HasIntro() bool {
	if o != nil && !IsNil(o.Intro) {
		return true
	}

	return false
}

// SetIntro gets a reference to the given string and assigns it to the Intro field.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) SetIntro(v string) {
	o.Intro = &v
}

// GetIntrofiles returns the Introfiles field value if set, zero value otherwise.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) GetIntrofiles() []CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner {
	if o == nil || IsNil(o.Introfiles) {
		var ret []CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner
		return ret
	}
	return o.Introfiles
}

// GetIntrofilesOk returns a tuple with the Introfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) GetIntrofilesOk() ([]CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner, bool) {
	if o == nil || IsNil(o.Introfiles) {
		return nil, false
	}
	return o.Introfiles, true
}

// HasIntrofiles returns a boolean if a field has been set.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) HasIntrofiles() bool {
	if o != nil && !IsNil(o.Introfiles) {
		return true
	}

	return false
}

// SetIntrofiles gets a reference to the given []CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner and assigns it to the Introfiles field.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) SetIntrofiles(v []CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) {
	o.Introfiles = v
}

// GetIntroformat returns the Introformat field value if set, zero value otherwise.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) GetIntroformat() int32 {
	if o == nil || IsNil(o.Introformat) {
		var ret int32
		return ret
	}
	return *o.Introformat
}

// GetIntroformatOk returns a tuple with the Introformat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) GetIntroformatOk() (*int32, bool) {
	if o == nil || IsNil(o.Introformat) {
		return nil, false
	}
	return o.Introformat, true
}

// HasIntroformat returns a boolean if a field has been set.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) HasIntroformat() bool {
	if o != nil && !IsNil(o.Introformat) {
		return true
	}

	return false
}

// SetIntroformat gets a reference to the given int32 and assigns it to the Introformat field.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) SetIntroformat(v int32) {
	o.Introformat = &v
}

// GetKeepdays returns the Keepdays field value if set, zero value otherwise.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) GetKeepdays() int32 {
	if o == nil || IsNil(o.Keepdays) {
		var ret int32
		return ret
	}
	return *o.Keepdays
}

// GetKeepdaysOk returns a tuple with the Keepdays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) GetKeepdaysOk() (*int32, bool) {
	if o == nil || IsNil(o.Keepdays) {
		return nil, false
	}
	return o.Keepdays, true
}

// HasKeepdays returns a boolean if a field has been set.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) HasKeepdays() bool {
	if o != nil && !IsNil(o.Keepdays) {
		return true
	}

	return false
}

// SetKeepdays gets a reference to the given int32 and assigns it to the Keepdays field.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) SetKeepdays(v int32) {
	o.Keepdays = &v
}

// GetLang returns the Lang field value if set, zero value otherwise.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) GetLang() string {
	if o == nil || IsNil(o.Lang) {
		var ret string
		return ret
	}
	return *o.Lang
}

// GetLangOk returns a tuple with the Lang field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) GetLangOk() (*string, bool) {
	if o == nil || IsNil(o.Lang) {
		return nil, false
	}
	return o.Lang, true
}

// HasLang returns a boolean if a field has been set.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) HasLang() bool {
	if o != nil && !IsNil(o.Lang) {
		return true
	}

	return false
}

// SetLang gets a reference to the given string and assigns it to the Lang field.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) SetLang(v string) {
	o.Lang = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) SetName(v string) {
	o.Name = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) GetSchedule() int32 {
	if o == nil || IsNil(o.Schedule) {
		var ret int32
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) GetScheduleOk() (*int32, bool) {
	if o == nil || IsNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) HasSchedule() bool {
	if o != nil && !IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given int32 and assigns it to the Schedule field.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) SetSchedule(v int32) {
	o.Schedule = &v
}

// GetSection returns the Section field value if set, zero value otherwise.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) GetSection() int32 {
	if o == nil || IsNil(o.Section) {
		var ret int32
		return ret
	}
	return *o.Section
}

// GetSectionOk returns a tuple with the Section field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) GetSectionOk() (*int32, bool) {
	if o == nil || IsNil(o.Section) {
		return nil, false
	}
	return o.Section, true
}

// HasSection returns a boolean if a field has been set.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) HasSection() bool {
	if o != nil && !IsNil(o.Section) {
		return true
	}

	return false
}

// SetSection gets a reference to the given int32 and assigns it to the Section field.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) SetSection(v int32) {
	o.Section = &v
}

// GetStudentlogs returns the Studentlogs field value if set, zero value otherwise.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) GetStudentlogs() int32 {
	if o == nil || IsNil(o.Studentlogs) {
		var ret int32
		return ret
	}
	return *o.Studentlogs
}

// GetStudentlogsOk returns a tuple with the Studentlogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) GetStudentlogsOk() (*int32, bool) {
	if o == nil || IsNil(o.Studentlogs) {
		return nil, false
	}
	return o.Studentlogs, true
}

// HasStudentlogs returns a boolean if a field has been set.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) HasStudentlogs() bool {
	if o != nil && !IsNil(o.Studentlogs) {
		return true
	}

	return false
}

// SetStudentlogs gets a reference to the given int32 and assigns it to the Studentlogs field.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) SetStudentlogs(v int32) {
	o.Studentlogs = &v
}

// GetTimemodified returns the Timemodified field value if set, zero value otherwise.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) GetTimemodified() int32 {
	if o == nil || IsNil(o.Timemodified) {
		var ret int32
		return ret
	}
	return *o.Timemodified
}

// GetTimemodifiedOk returns a tuple with the Timemodified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) GetTimemodifiedOk() (*int32, bool) {
	if o == nil || IsNil(o.Timemodified) {
		return nil, false
	}
	return o.Timemodified, true
}

// HasTimemodified returns a boolean if a field has been set.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) HasTimemodified() bool {
	if o != nil && !IsNil(o.Timemodified) {
		return true
	}

	return false
}

// SetTimemodified gets a reference to the given int32 and assigns it to the Timemodified field.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) SetTimemodified(v int32) {
	o.Timemodified = &v
}

// GetVisible returns the Visible field value if set, zero value otherwise.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) GetVisible() bool {
	if o == nil || IsNil(o.Visible) {
		var ret bool
		return ret
	}
	return *o.Visible
}

// GetVisibleOk returns a tuple with the Visible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) GetVisibleOk() (*bool, bool) {
	if o == nil || IsNil(o.Visible) {
		return nil, false
	}
	return o.Visible, true
}

// HasVisible returns a boolean if a field has been set.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) HasVisible() bool {
	if o != nil && !IsNil(o.Visible) {
		return true
	}

	return false
}

// SetVisible gets a reference to the given bool and assigns it to the Visible field.
func (o *ModChatGetChatsByCourses200ResponseChatsInner) SetVisible(v bool) {
	o.Visible = &v
}

func (o ModChatGetChatsByCourses200ResponseChatsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModChatGetChatsByCourses200ResponseChatsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Chatmethod) {
		toSerialize["chatmethod"] = o.Chatmethod
	}
	if !IsNil(o.Chattime) {
		toSerialize["chattime"] = o.Chattime
	}
	if !IsNil(o.Course) {
		toSerialize["course"] = o.Course
	}
	if !IsNil(o.Coursemodule) {
		toSerialize["coursemodule"] = o.Coursemodule
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
	if !IsNil(o.Keepdays) {
		toSerialize["keepdays"] = o.Keepdays
	}
	if !IsNil(o.Lang) {
		toSerialize["lang"] = o.Lang
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	if !IsNil(o.Section) {
		toSerialize["section"] = o.Section
	}
	if !IsNil(o.Studentlogs) {
		toSerialize["studentlogs"] = o.Studentlogs
	}
	if !IsNil(o.Timemodified) {
		toSerialize["timemodified"] = o.Timemodified
	}
	if !IsNil(o.Visible) {
		toSerialize["visible"] = o.Visible
	}
	return toSerialize, nil
}

type NullableModChatGetChatsByCourses200ResponseChatsInner struct {
	value *ModChatGetChatsByCourses200ResponseChatsInner
	isSet bool
}

func (v NullableModChatGetChatsByCourses200ResponseChatsInner) Get() *ModChatGetChatsByCourses200ResponseChatsInner {
	return v.value
}

func (v *NullableModChatGetChatsByCourses200ResponseChatsInner) Set(val *ModChatGetChatsByCourses200ResponseChatsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableModChatGetChatsByCourses200ResponseChatsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableModChatGetChatsByCourses200ResponseChatsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModChatGetChatsByCourses200ResponseChatsInner(val *ModChatGetChatsByCourses200ResponseChatsInner) *NullableModChatGetChatsByCourses200ResponseChatsInner {
	return &NullableModChatGetChatsByCourses200ResponseChatsInner{value: val, isSet: true}
}

func (v NullableModChatGetChatsByCourses200ResponseChatsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModChatGetChatsByCourses200ResponseChatsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


