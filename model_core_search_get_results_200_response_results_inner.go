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

// checks if the CoreSearchGetResults200ResponseResultsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreSearchGetResults200ResponseResultsInner{}

// CoreSearchGetResults200ResponseResultsInner struct for CoreSearchGetResults200ResponseResultsInner
type CoreSearchGetResults200ResponseResultsInner struct {
	// search area name
	Areaname *string `json:"areaname,omitempty"`
	// component name
	Componentname *string `json:"componentname,omitempty"`
	// result contents
	Content *string `json:"content,omitempty"`
	// result context id
	Contextid *int32 `json:"contextid,omitempty"`
	// result context url
	Contexturl *string `json:"contexturl,omitempty"`
	// result course fullname
	Coursefullname *string `json:"coursefullname,omitempty"`
	// result course url
	Courseurl *string `json:"courseurl,omitempty"`
	// extra result contents, depends on the search area
	Description1 *string `json:"description1,omitempty"`
	// extra result contents, depends on the search area
	Description2 *string `json:"description2,omitempty"`
	// result url
	Docurl *string `json:"docurl,omitempty"`
	// result file name if present
	Filename *string `json:"filename,omitempty"`
	// result file names if present
	Filenames *string `json:"filenames,omitempty"`
	// icon url
	Iconurl *string `json:"iconurl,omitempty"`
	// unique id in the search area scope
	Itemid *int32 `json:"itemid,omitempty"`
	// whether multiple files are returned or not
	Multiplefiles *int32 `json:"multiplefiles,omitempty"`
	// text fields format, it is the same for all of them
	Textformat *int32 `json:"textformat,omitempty"`
	// result modified time
	Timemodified *int32 `json:"timemodified,omitempty"`
	// result title
	Title *string `json:"title,omitempty"`
	// user fullname
	Userfullname *string `json:"userfullname,omitempty"`
	// user id
	Userid *int32 `json:"userid,omitempty"`
	// user url
	Userurl *string `json:"userurl,omitempty"`
}

// NewCoreSearchGetResults200ResponseResultsInner instantiates a new CoreSearchGetResults200ResponseResultsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreSearchGetResults200ResponseResultsInner() *CoreSearchGetResults200ResponseResultsInner {
	this := CoreSearchGetResults200ResponseResultsInner{}
	var areaname string = "null"
	this.Areaname = &areaname
	var componentname string = "null"
	this.Componentname = &componentname
	var content string = ""
	this.Content = &content
	var contextid int32 = null
	this.Contextid = &contextid
	var contexturl string = "null"
	this.Contexturl = &contexturl
	var coursefullname string = "null"
	this.Coursefullname = &coursefullname
	var courseurl string = "null"
	this.Courseurl = &courseurl
	var description1 string = ""
	this.Description1 = &description1
	var description2 string = ""
	this.Description2 = &description2
	var docurl string = "null"
	this.Docurl = &docurl
	var filename string = "null"
	this.Filename = &filename
	var filenames string = "null"
	this.Filenames = &filenames
	var iconurl string = ""
	this.Iconurl = &iconurl
	var itemid int32 = null
	this.Itemid = &itemid
	var multiplefiles int32 = null
	this.Multiplefiles = &multiplefiles
	var textformat int32 = null
	this.Textformat = &textformat
	var timemodified int32 = null
	this.Timemodified = &timemodified
	var title string = "null"
	this.Title = &title
	var userfullname string = "null"
	this.Userfullname = &userfullname
	var userid int32 = null
	this.Userid = &userid
	var userurl string = "null"
	this.Userurl = &userurl
	return &this
}

// NewCoreSearchGetResults200ResponseResultsInnerWithDefaults instantiates a new CoreSearchGetResults200ResponseResultsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreSearchGetResults200ResponseResultsInnerWithDefaults() *CoreSearchGetResults200ResponseResultsInner {
	this := CoreSearchGetResults200ResponseResultsInner{}
	var areaname string = "null"
	this.Areaname = &areaname
	var componentname string = "null"
	this.Componentname = &componentname
	var content string = ""
	this.Content = &content
	var contextid int32 = null
	this.Contextid = &contextid
	var contexturl string = "null"
	this.Contexturl = &contexturl
	var coursefullname string = "null"
	this.Coursefullname = &coursefullname
	var courseurl string = "null"
	this.Courseurl = &courseurl
	var description1 string = ""
	this.Description1 = &description1
	var description2 string = ""
	this.Description2 = &description2
	var docurl string = "null"
	this.Docurl = &docurl
	var filename string = "null"
	this.Filename = &filename
	var filenames string = "null"
	this.Filenames = &filenames
	var iconurl string = ""
	this.Iconurl = &iconurl
	var itemid int32 = null
	this.Itemid = &itemid
	var multiplefiles int32 = null
	this.Multiplefiles = &multiplefiles
	var textformat int32 = null
	this.Textformat = &textformat
	var timemodified int32 = null
	this.Timemodified = &timemodified
	var title string = "null"
	this.Title = &title
	var userfullname string = "null"
	this.Userfullname = &userfullname
	var userid int32 = null
	this.Userid = &userid
	var userurl string = "null"
	this.Userurl = &userurl
	return &this
}

// GetAreaname returns the Areaname field value if set, zero value otherwise.
func (o *CoreSearchGetResults200ResponseResultsInner) GetAreaname() string {
	if o == nil || IsNil(o.Areaname) {
		var ret string
		return ret
	}
	return *o.Areaname
}

// GetAreanameOk returns a tuple with the Areaname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) GetAreanameOk() (*string, bool) {
	if o == nil || IsNil(o.Areaname) {
		return nil, false
	}
	return o.Areaname, true
}

// HasAreaname returns a boolean if a field has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) HasAreaname() bool {
	if o != nil && !IsNil(o.Areaname) {
		return true
	}

	return false
}

// SetAreaname gets a reference to the given string and assigns it to the Areaname field.
func (o *CoreSearchGetResults200ResponseResultsInner) SetAreaname(v string) {
	o.Areaname = &v
}

// GetComponentname returns the Componentname field value if set, zero value otherwise.
func (o *CoreSearchGetResults200ResponseResultsInner) GetComponentname() string {
	if o == nil || IsNil(o.Componentname) {
		var ret string
		return ret
	}
	return *o.Componentname
}

// GetComponentnameOk returns a tuple with the Componentname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) GetComponentnameOk() (*string, bool) {
	if o == nil || IsNil(o.Componentname) {
		return nil, false
	}
	return o.Componentname, true
}

// HasComponentname returns a boolean if a field has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) HasComponentname() bool {
	if o != nil && !IsNil(o.Componentname) {
		return true
	}

	return false
}

// SetComponentname gets a reference to the given string and assigns it to the Componentname field.
func (o *CoreSearchGetResults200ResponseResultsInner) SetComponentname(v string) {
	o.Componentname = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *CoreSearchGetResults200ResponseResultsInner) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *CoreSearchGetResults200ResponseResultsInner) SetContent(v string) {
	o.Content = &v
}

// GetContextid returns the Contextid field value if set, zero value otherwise.
func (o *CoreSearchGetResults200ResponseResultsInner) GetContextid() int32 {
	if o == nil || IsNil(o.Contextid) {
		var ret int32
		return ret
	}
	return *o.Contextid
}

// GetContextidOk returns a tuple with the Contextid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) GetContextidOk() (*int32, bool) {
	if o == nil || IsNil(o.Contextid) {
		return nil, false
	}
	return o.Contextid, true
}

// HasContextid returns a boolean if a field has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) HasContextid() bool {
	if o != nil && !IsNil(o.Contextid) {
		return true
	}

	return false
}

// SetContextid gets a reference to the given int32 and assigns it to the Contextid field.
func (o *CoreSearchGetResults200ResponseResultsInner) SetContextid(v int32) {
	o.Contextid = &v
}

// GetContexturl returns the Contexturl field value if set, zero value otherwise.
func (o *CoreSearchGetResults200ResponseResultsInner) GetContexturl() string {
	if o == nil || IsNil(o.Contexturl) {
		var ret string
		return ret
	}
	return *o.Contexturl
}

// GetContexturlOk returns a tuple with the Contexturl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) GetContexturlOk() (*string, bool) {
	if o == nil || IsNil(o.Contexturl) {
		return nil, false
	}
	return o.Contexturl, true
}

// HasContexturl returns a boolean if a field has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) HasContexturl() bool {
	if o != nil && !IsNil(o.Contexturl) {
		return true
	}

	return false
}

// SetContexturl gets a reference to the given string and assigns it to the Contexturl field.
func (o *CoreSearchGetResults200ResponseResultsInner) SetContexturl(v string) {
	o.Contexturl = &v
}

// GetCoursefullname returns the Coursefullname field value if set, zero value otherwise.
func (o *CoreSearchGetResults200ResponseResultsInner) GetCoursefullname() string {
	if o == nil || IsNil(o.Coursefullname) {
		var ret string
		return ret
	}
	return *o.Coursefullname
}

// GetCoursefullnameOk returns a tuple with the Coursefullname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) GetCoursefullnameOk() (*string, bool) {
	if o == nil || IsNil(o.Coursefullname) {
		return nil, false
	}
	return o.Coursefullname, true
}

// HasCoursefullname returns a boolean if a field has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) HasCoursefullname() bool {
	if o != nil && !IsNil(o.Coursefullname) {
		return true
	}

	return false
}

// SetCoursefullname gets a reference to the given string and assigns it to the Coursefullname field.
func (o *CoreSearchGetResults200ResponseResultsInner) SetCoursefullname(v string) {
	o.Coursefullname = &v
}

// GetCourseurl returns the Courseurl field value if set, zero value otherwise.
func (o *CoreSearchGetResults200ResponseResultsInner) GetCourseurl() string {
	if o == nil || IsNil(o.Courseurl) {
		var ret string
		return ret
	}
	return *o.Courseurl
}

// GetCourseurlOk returns a tuple with the Courseurl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) GetCourseurlOk() (*string, bool) {
	if o == nil || IsNil(o.Courseurl) {
		return nil, false
	}
	return o.Courseurl, true
}

// HasCourseurl returns a boolean if a field has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) HasCourseurl() bool {
	if o != nil && !IsNil(o.Courseurl) {
		return true
	}

	return false
}

// SetCourseurl gets a reference to the given string and assigns it to the Courseurl field.
func (o *CoreSearchGetResults200ResponseResultsInner) SetCourseurl(v string) {
	o.Courseurl = &v
}

// GetDescription1 returns the Description1 field value if set, zero value otherwise.
func (o *CoreSearchGetResults200ResponseResultsInner) GetDescription1() string {
	if o == nil || IsNil(o.Description1) {
		var ret string
		return ret
	}
	return *o.Description1
}

// GetDescription1Ok returns a tuple with the Description1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) GetDescription1Ok() (*string, bool) {
	if o == nil || IsNil(o.Description1) {
		return nil, false
	}
	return o.Description1, true
}

// HasDescription1 returns a boolean if a field has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) HasDescription1() bool {
	if o != nil && !IsNil(o.Description1) {
		return true
	}

	return false
}

// SetDescription1 gets a reference to the given string and assigns it to the Description1 field.
func (o *CoreSearchGetResults200ResponseResultsInner) SetDescription1(v string) {
	o.Description1 = &v
}

// GetDescription2 returns the Description2 field value if set, zero value otherwise.
func (o *CoreSearchGetResults200ResponseResultsInner) GetDescription2() string {
	if o == nil || IsNil(o.Description2) {
		var ret string
		return ret
	}
	return *o.Description2
}

// GetDescription2Ok returns a tuple with the Description2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) GetDescription2Ok() (*string, bool) {
	if o == nil || IsNil(o.Description2) {
		return nil, false
	}
	return o.Description2, true
}

// HasDescription2 returns a boolean if a field has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) HasDescription2() bool {
	if o != nil && !IsNil(o.Description2) {
		return true
	}

	return false
}

// SetDescription2 gets a reference to the given string and assigns it to the Description2 field.
func (o *CoreSearchGetResults200ResponseResultsInner) SetDescription2(v string) {
	o.Description2 = &v
}

// GetDocurl returns the Docurl field value if set, zero value otherwise.
func (o *CoreSearchGetResults200ResponseResultsInner) GetDocurl() string {
	if o == nil || IsNil(o.Docurl) {
		var ret string
		return ret
	}
	return *o.Docurl
}

// GetDocurlOk returns a tuple with the Docurl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) GetDocurlOk() (*string, bool) {
	if o == nil || IsNil(o.Docurl) {
		return nil, false
	}
	return o.Docurl, true
}

// HasDocurl returns a boolean if a field has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) HasDocurl() bool {
	if o != nil && !IsNil(o.Docurl) {
		return true
	}

	return false
}

// SetDocurl gets a reference to the given string and assigns it to the Docurl field.
func (o *CoreSearchGetResults200ResponseResultsInner) SetDocurl(v string) {
	o.Docurl = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *CoreSearchGetResults200ResponseResultsInner) GetFilename() string {
	if o == nil || IsNil(o.Filename) {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) GetFilenameOk() (*string, bool) {
	if o == nil || IsNil(o.Filename) {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) HasFilename() bool {
	if o != nil && !IsNil(o.Filename) {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *CoreSearchGetResults200ResponseResultsInner) SetFilename(v string) {
	o.Filename = &v
}

// GetFilenames returns the Filenames field value if set, zero value otherwise.
func (o *CoreSearchGetResults200ResponseResultsInner) GetFilenames() string {
	if o == nil || IsNil(o.Filenames) {
		var ret string
		return ret
	}
	return *o.Filenames
}

// GetFilenamesOk returns a tuple with the Filenames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) GetFilenamesOk() (*string, bool) {
	if o == nil || IsNil(o.Filenames) {
		return nil, false
	}
	return o.Filenames, true
}

// HasFilenames returns a boolean if a field has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) HasFilenames() bool {
	if o != nil && !IsNil(o.Filenames) {
		return true
	}

	return false
}

// SetFilenames gets a reference to the given string and assigns it to the Filenames field.
func (o *CoreSearchGetResults200ResponseResultsInner) SetFilenames(v string) {
	o.Filenames = &v
}

// GetIconurl returns the Iconurl field value if set, zero value otherwise.
func (o *CoreSearchGetResults200ResponseResultsInner) GetIconurl() string {
	if o == nil || IsNil(o.Iconurl) {
		var ret string
		return ret
	}
	return *o.Iconurl
}

// GetIconurlOk returns a tuple with the Iconurl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) GetIconurlOk() (*string, bool) {
	if o == nil || IsNil(o.Iconurl) {
		return nil, false
	}
	return o.Iconurl, true
}

// HasIconurl returns a boolean if a field has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) HasIconurl() bool {
	if o != nil && !IsNil(o.Iconurl) {
		return true
	}

	return false
}

// SetIconurl gets a reference to the given string and assigns it to the Iconurl field.
func (o *CoreSearchGetResults200ResponseResultsInner) SetIconurl(v string) {
	o.Iconurl = &v
}

// GetItemid returns the Itemid field value if set, zero value otherwise.
func (o *CoreSearchGetResults200ResponseResultsInner) GetItemid() int32 {
	if o == nil || IsNil(o.Itemid) {
		var ret int32
		return ret
	}
	return *o.Itemid
}

// GetItemidOk returns a tuple with the Itemid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) GetItemidOk() (*int32, bool) {
	if o == nil || IsNil(o.Itemid) {
		return nil, false
	}
	return o.Itemid, true
}

// HasItemid returns a boolean if a field has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) HasItemid() bool {
	if o != nil && !IsNil(o.Itemid) {
		return true
	}

	return false
}

// SetItemid gets a reference to the given int32 and assigns it to the Itemid field.
func (o *CoreSearchGetResults200ResponseResultsInner) SetItemid(v int32) {
	o.Itemid = &v
}

// GetMultiplefiles returns the Multiplefiles field value if set, zero value otherwise.
func (o *CoreSearchGetResults200ResponseResultsInner) GetMultiplefiles() int32 {
	if o == nil || IsNil(o.Multiplefiles) {
		var ret int32
		return ret
	}
	return *o.Multiplefiles
}

// GetMultiplefilesOk returns a tuple with the Multiplefiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) GetMultiplefilesOk() (*int32, bool) {
	if o == nil || IsNil(o.Multiplefiles) {
		return nil, false
	}
	return o.Multiplefiles, true
}

// HasMultiplefiles returns a boolean if a field has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) HasMultiplefiles() bool {
	if o != nil && !IsNil(o.Multiplefiles) {
		return true
	}

	return false
}

// SetMultiplefiles gets a reference to the given int32 and assigns it to the Multiplefiles field.
func (o *CoreSearchGetResults200ResponseResultsInner) SetMultiplefiles(v int32) {
	o.Multiplefiles = &v
}

// GetTextformat returns the Textformat field value if set, zero value otherwise.
func (o *CoreSearchGetResults200ResponseResultsInner) GetTextformat() int32 {
	if o == nil || IsNil(o.Textformat) {
		var ret int32
		return ret
	}
	return *o.Textformat
}

// GetTextformatOk returns a tuple with the Textformat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) GetTextformatOk() (*int32, bool) {
	if o == nil || IsNil(o.Textformat) {
		return nil, false
	}
	return o.Textformat, true
}

// HasTextformat returns a boolean if a field has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) HasTextformat() bool {
	if o != nil && !IsNil(o.Textformat) {
		return true
	}

	return false
}

// SetTextformat gets a reference to the given int32 and assigns it to the Textformat field.
func (o *CoreSearchGetResults200ResponseResultsInner) SetTextformat(v int32) {
	o.Textformat = &v
}

// GetTimemodified returns the Timemodified field value if set, zero value otherwise.
func (o *CoreSearchGetResults200ResponseResultsInner) GetTimemodified() int32 {
	if o == nil || IsNil(o.Timemodified) {
		var ret int32
		return ret
	}
	return *o.Timemodified
}

// GetTimemodifiedOk returns a tuple with the Timemodified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) GetTimemodifiedOk() (*int32, bool) {
	if o == nil || IsNil(o.Timemodified) {
		return nil, false
	}
	return o.Timemodified, true
}

// HasTimemodified returns a boolean if a field has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) HasTimemodified() bool {
	if o != nil && !IsNil(o.Timemodified) {
		return true
	}

	return false
}

// SetTimemodified gets a reference to the given int32 and assigns it to the Timemodified field.
func (o *CoreSearchGetResults200ResponseResultsInner) SetTimemodified(v int32) {
	o.Timemodified = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *CoreSearchGetResults200ResponseResultsInner) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *CoreSearchGetResults200ResponseResultsInner) SetTitle(v string) {
	o.Title = &v
}

// GetUserfullname returns the Userfullname field value if set, zero value otherwise.
func (o *CoreSearchGetResults200ResponseResultsInner) GetUserfullname() string {
	if o == nil || IsNil(o.Userfullname) {
		var ret string
		return ret
	}
	return *o.Userfullname
}

// GetUserfullnameOk returns a tuple with the Userfullname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) GetUserfullnameOk() (*string, bool) {
	if o == nil || IsNil(o.Userfullname) {
		return nil, false
	}
	return o.Userfullname, true
}

// HasUserfullname returns a boolean if a field has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) HasUserfullname() bool {
	if o != nil && !IsNil(o.Userfullname) {
		return true
	}

	return false
}

// SetUserfullname gets a reference to the given string and assigns it to the Userfullname field.
func (o *CoreSearchGetResults200ResponseResultsInner) SetUserfullname(v string) {
	o.Userfullname = &v
}

// GetUserid returns the Userid field value if set, zero value otherwise.
func (o *CoreSearchGetResults200ResponseResultsInner) GetUserid() int32 {
	if o == nil || IsNil(o.Userid) {
		var ret int32
		return ret
	}
	return *o.Userid
}

// GetUseridOk returns a tuple with the Userid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) GetUseridOk() (*int32, bool) {
	if o == nil || IsNil(o.Userid) {
		return nil, false
	}
	return o.Userid, true
}

// HasUserid returns a boolean if a field has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) HasUserid() bool {
	if o != nil && !IsNil(o.Userid) {
		return true
	}

	return false
}

// SetUserid gets a reference to the given int32 and assigns it to the Userid field.
func (o *CoreSearchGetResults200ResponseResultsInner) SetUserid(v int32) {
	o.Userid = &v
}

// GetUserurl returns the Userurl field value if set, zero value otherwise.
func (o *CoreSearchGetResults200ResponseResultsInner) GetUserurl() string {
	if o == nil || IsNil(o.Userurl) {
		var ret string
		return ret
	}
	return *o.Userurl
}

// GetUserurlOk returns a tuple with the Userurl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) GetUserurlOk() (*string, bool) {
	if o == nil || IsNil(o.Userurl) {
		return nil, false
	}
	return o.Userurl, true
}

// HasUserurl returns a boolean if a field has been set.
func (o *CoreSearchGetResults200ResponseResultsInner) HasUserurl() bool {
	if o != nil && !IsNil(o.Userurl) {
		return true
	}

	return false
}

// SetUserurl gets a reference to the given string and assigns it to the Userurl field.
func (o *CoreSearchGetResults200ResponseResultsInner) SetUserurl(v string) {
	o.Userurl = &v
}

func (o CoreSearchGetResults200ResponseResultsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreSearchGetResults200ResponseResultsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Areaname) {
		toSerialize["areaname"] = o.Areaname
	}
	if !IsNil(o.Componentname) {
		toSerialize["componentname"] = o.Componentname
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.Contextid) {
		toSerialize["contextid"] = o.Contextid
	}
	if !IsNil(o.Contexturl) {
		toSerialize["contexturl"] = o.Contexturl
	}
	if !IsNil(o.Coursefullname) {
		toSerialize["coursefullname"] = o.Coursefullname
	}
	if !IsNil(o.Courseurl) {
		toSerialize["courseurl"] = o.Courseurl
	}
	if !IsNil(o.Description1) {
		toSerialize["description1"] = o.Description1
	}
	if !IsNil(o.Description2) {
		toSerialize["description2"] = o.Description2
	}
	if !IsNil(o.Docurl) {
		toSerialize["docurl"] = o.Docurl
	}
	if !IsNil(o.Filename) {
		toSerialize["filename"] = o.Filename
	}
	if !IsNil(o.Filenames) {
		toSerialize["filenames"] = o.Filenames
	}
	if !IsNil(o.Iconurl) {
		toSerialize["iconurl"] = o.Iconurl
	}
	if !IsNil(o.Itemid) {
		toSerialize["itemid"] = o.Itemid
	}
	if !IsNil(o.Multiplefiles) {
		toSerialize["multiplefiles"] = o.Multiplefiles
	}
	if !IsNil(o.Textformat) {
		toSerialize["textformat"] = o.Textformat
	}
	if !IsNil(o.Timemodified) {
		toSerialize["timemodified"] = o.Timemodified
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Userfullname) {
		toSerialize["userfullname"] = o.Userfullname
	}
	if !IsNil(o.Userid) {
		toSerialize["userid"] = o.Userid
	}
	if !IsNil(o.Userurl) {
		toSerialize["userurl"] = o.Userurl
	}
	return toSerialize, nil
}

type NullableCoreSearchGetResults200ResponseResultsInner struct {
	value *CoreSearchGetResults200ResponseResultsInner
	isSet bool
}

func (v NullableCoreSearchGetResults200ResponseResultsInner) Get() *CoreSearchGetResults200ResponseResultsInner {
	return v.value
}

func (v *NullableCoreSearchGetResults200ResponseResultsInner) Set(val *CoreSearchGetResults200ResponseResultsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreSearchGetResults200ResponseResultsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreSearchGetResults200ResponseResultsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreSearchGetResults200ResponseResultsInner(val *CoreSearchGetResults200ResponseResultsInner) *NullableCoreSearchGetResults200ResponseResultsInner {
	return &NullableCoreSearchGetResults200ResponseResultsInner{value: val, isSet: true}
}

func (v NullableCoreSearchGetResults200ResponseResultsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreSearchGetResults200ResponseResultsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

