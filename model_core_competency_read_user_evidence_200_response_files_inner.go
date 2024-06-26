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

// checks if the CoreCompetencyReadUserEvidence200ResponseFilesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCompetencyReadUserEvidence200ResponseFilesInner{}

// CoreCompetencyReadUserEvidence200ResponseFilesInner struct for CoreCompetencyReadUserEvidence200ResponseFilesInner
type CoreCompetencyReadUserEvidence200ResponseFilesInner struct {
	// author
	Author *string `json:"author,omitempty"`
	// component
	Component *string `json:"component,omitempty"`
	// contextid
	Contextid *int32 `json:"contextid,omitempty"`
	// filearea
	Filearea *string `json:"filearea,omitempty"`
	// filename
	Filename *string `json:"filename,omitempty"`
	// filenameshort
	Filenameshort *string `json:"filenameshort,omitempty"`
	// filepath
	Filepath *string `json:"filepath,omitempty"`
	// filesize
	Filesize *int32 `json:"filesize,omitempty"`
	// filesizeformatted
	Filesizeformatted *string `json:"filesizeformatted,omitempty"`
	// icon
	Icon *string `json:"icon,omitempty"`
	// isdir
	Isdir *bool `json:"isdir,omitempty"`
	// isimage
	Isimage *bool `json:"isimage,omitempty"`
	// itemid
	Itemid *int32 `json:"itemid,omitempty"`
	// license
	License *string `json:"license,omitempty"`
	// timecreated
	Timecreated *int32 `json:"timecreated,omitempty"`
	// timecreatedformatted
	Timecreatedformatted *string `json:"timecreatedformatted,omitempty"`
	// timemodified
	Timemodified *int32 `json:"timemodified,omitempty"`
	// timemodifiedformatted
	Timemodifiedformatted *string `json:"timemodifiedformatted,omitempty"`
	// url
	Url *string `json:"url,omitempty"`
}

// NewCoreCompetencyReadUserEvidence200ResponseFilesInner instantiates a new CoreCompetencyReadUserEvidence200ResponseFilesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCompetencyReadUserEvidence200ResponseFilesInner() *CoreCompetencyReadUserEvidence200ResponseFilesInner {
	this := CoreCompetencyReadUserEvidence200ResponseFilesInner{}
	var author string = "null"
	this.Author = &author
	var filearea string = "null"
	this.Filearea = &filearea
	var filename string = "null"
	this.Filename = &filename
	var filenameshort string = "null"
	this.Filenameshort = &filenameshort
	var filepath string = "null"
	this.Filepath = &filepath
	var filesize int32 = null
	this.Filesize = &filesize
	var filesizeformatted string = "null"
	this.Filesizeformatted = &filesizeformatted
	var icon string = "null"
	this.Icon = &icon
	var isdir bool = null
	this.Isdir = &isdir
	var isimage bool = null
	this.Isimage = &isimage
	var license string = "null"
	this.License = &license
	var timecreated int32 = null
	this.Timecreated = &timecreated
	var timecreatedformatted string = "null"
	this.Timecreatedformatted = &timecreatedformatted
	var timemodifiedformatted string = "null"
	this.Timemodifiedformatted = &timemodifiedformatted
	return &this
}

// NewCoreCompetencyReadUserEvidence200ResponseFilesInnerWithDefaults instantiates a new CoreCompetencyReadUserEvidence200ResponseFilesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCompetencyReadUserEvidence200ResponseFilesInnerWithDefaults() *CoreCompetencyReadUserEvidence200ResponseFilesInner {
	this := CoreCompetencyReadUserEvidence200ResponseFilesInner{}
	var author string = "null"
	this.Author = &author
	var filearea string = "null"
	this.Filearea = &filearea
	var filename string = "null"
	this.Filename = &filename
	var filenameshort string = "null"
	this.Filenameshort = &filenameshort
	var filepath string = "null"
	this.Filepath = &filepath
	var filesize int32 = null
	this.Filesize = &filesize
	var filesizeformatted string = "null"
	this.Filesizeformatted = &filesizeformatted
	var icon string = "null"
	this.Icon = &icon
	var isdir bool = null
	this.Isdir = &isdir
	var isimage bool = null
	this.Isimage = &isimage
	var license string = "null"
	this.License = &license
	var timecreated int32 = null
	this.Timecreated = &timecreated
	var timecreatedformatted string = "null"
	this.Timecreatedformatted = &timecreatedformatted
	var timemodifiedformatted string = "null"
	this.Timemodifiedformatted = &timemodifiedformatted
	return &this
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetAuthor() string {
	if o == nil || IsNil(o.Author) {
		var ret string
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetAuthorOk() (*string, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given string and assigns it to the Author field.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) SetAuthor(v string) {
	o.Author = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetComponent() string {
	if o == nil || IsNil(o.Component) {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetComponentOk() (*string, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) SetComponent(v string) {
	o.Component = &v
}

// GetContextid returns the Contextid field value if set, zero value otherwise.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetContextid() int32 {
	if o == nil || IsNil(o.Contextid) {
		var ret int32
		return ret
	}
	return *o.Contextid
}

// GetContextidOk returns a tuple with the Contextid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetContextidOk() (*int32, bool) {
	if o == nil || IsNil(o.Contextid) {
		return nil, false
	}
	return o.Contextid, true
}

// HasContextid returns a boolean if a field has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) HasContextid() bool {
	if o != nil && !IsNil(o.Contextid) {
		return true
	}

	return false
}

// SetContextid gets a reference to the given int32 and assigns it to the Contextid field.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) SetContextid(v int32) {
	o.Contextid = &v
}

// GetFilearea returns the Filearea field value if set, zero value otherwise.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetFilearea() string {
	if o == nil || IsNil(o.Filearea) {
		var ret string
		return ret
	}
	return *o.Filearea
}

// GetFileareaOk returns a tuple with the Filearea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetFileareaOk() (*string, bool) {
	if o == nil || IsNil(o.Filearea) {
		return nil, false
	}
	return o.Filearea, true
}

// HasFilearea returns a boolean if a field has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) HasFilearea() bool {
	if o != nil && !IsNil(o.Filearea) {
		return true
	}

	return false
}

// SetFilearea gets a reference to the given string and assigns it to the Filearea field.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) SetFilearea(v string) {
	o.Filearea = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetFilename() string {
	if o == nil || IsNil(o.Filename) {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetFilenameOk() (*string, bool) {
	if o == nil || IsNil(o.Filename) {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) HasFilename() bool {
	if o != nil && !IsNil(o.Filename) {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) SetFilename(v string) {
	o.Filename = &v
}

// GetFilenameshort returns the Filenameshort field value if set, zero value otherwise.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetFilenameshort() string {
	if o == nil || IsNil(o.Filenameshort) {
		var ret string
		return ret
	}
	return *o.Filenameshort
}

// GetFilenameshortOk returns a tuple with the Filenameshort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetFilenameshortOk() (*string, bool) {
	if o == nil || IsNil(o.Filenameshort) {
		return nil, false
	}
	return o.Filenameshort, true
}

// HasFilenameshort returns a boolean if a field has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) HasFilenameshort() bool {
	if o != nil && !IsNil(o.Filenameshort) {
		return true
	}

	return false
}

// SetFilenameshort gets a reference to the given string and assigns it to the Filenameshort field.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) SetFilenameshort(v string) {
	o.Filenameshort = &v
}

// GetFilepath returns the Filepath field value if set, zero value otherwise.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetFilepath() string {
	if o == nil || IsNil(o.Filepath) {
		var ret string
		return ret
	}
	return *o.Filepath
}

// GetFilepathOk returns a tuple with the Filepath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetFilepathOk() (*string, bool) {
	if o == nil || IsNil(o.Filepath) {
		return nil, false
	}
	return o.Filepath, true
}

// HasFilepath returns a boolean if a field has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) HasFilepath() bool {
	if o != nil && !IsNil(o.Filepath) {
		return true
	}

	return false
}

// SetFilepath gets a reference to the given string and assigns it to the Filepath field.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) SetFilepath(v string) {
	o.Filepath = &v
}

// GetFilesize returns the Filesize field value if set, zero value otherwise.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetFilesize() int32 {
	if o == nil || IsNil(o.Filesize) {
		var ret int32
		return ret
	}
	return *o.Filesize
}

// GetFilesizeOk returns a tuple with the Filesize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetFilesizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Filesize) {
		return nil, false
	}
	return o.Filesize, true
}

// HasFilesize returns a boolean if a field has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) HasFilesize() bool {
	if o != nil && !IsNil(o.Filesize) {
		return true
	}

	return false
}

// SetFilesize gets a reference to the given int32 and assigns it to the Filesize field.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) SetFilesize(v int32) {
	o.Filesize = &v
}

// GetFilesizeformatted returns the Filesizeformatted field value if set, zero value otherwise.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetFilesizeformatted() string {
	if o == nil || IsNil(o.Filesizeformatted) {
		var ret string
		return ret
	}
	return *o.Filesizeformatted
}

// GetFilesizeformattedOk returns a tuple with the Filesizeformatted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetFilesizeformattedOk() (*string, bool) {
	if o == nil || IsNil(o.Filesizeformatted) {
		return nil, false
	}
	return o.Filesizeformatted, true
}

// HasFilesizeformatted returns a boolean if a field has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) HasFilesizeformatted() bool {
	if o != nil && !IsNil(o.Filesizeformatted) {
		return true
	}

	return false
}

// SetFilesizeformatted gets a reference to the given string and assigns it to the Filesizeformatted field.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) SetFilesizeformatted(v string) {
	o.Filesizeformatted = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetIcon() string {
	if o == nil || IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) SetIcon(v string) {
	o.Icon = &v
}

// GetIsdir returns the Isdir field value if set, zero value otherwise.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetIsdir() bool {
	if o == nil || IsNil(o.Isdir) {
		var ret bool
		return ret
	}
	return *o.Isdir
}

// GetIsdirOk returns a tuple with the Isdir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetIsdirOk() (*bool, bool) {
	if o == nil || IsNil(o.Isdir) {
		return nil, false
	}
	return o.Isdir, true
}

// HasIsdir returns a boolean if a field has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) HasIsdir() bool {
	if o != nil && !IsNil(o.Isdir) {
		return true
	}

	return false
}

// SetIsdir gets a reference to the given bool and assigns it to the Isdir field.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) SetIsdir(v bool) {
	o.Isdir = &v
}

// GetIsimage returns the Isimage field value if set, zero value otherwise.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetIsimage() bool {
	if o == nil || IsNil(o.Isimage) {
		var ret bool
		return ret
	}
	return *o.Isimage
}

// GetIsimageOk returns a tuple with the Isimage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetIsimageOk() (*bool, bool) {
	if o == nil || IsNil(o.Isimage) {
		return nil, false
	}
	return o.Isimage, true
}

// HasIsimage returns a boolean if a field has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) HasIsimage() bool {
	if o != nil && !IsNil(o.Isimage) {
		return true
	}

	return false
}

// SetIsimage gets a reference to the given bool and assigns it to the Isimage field.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) SetIsimage(v bool) {
	o.Isimage = &v
}

// GetItemid returns the Itemid field value if set, zero value otherwise.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetItemid() int32 {
	if o == nil || IsNil(o.Itemid) {
		var ret int32
		return ret
	}
	return *o.Itemid
}

// GetItemidOk returns a tuple with the Itemid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetItemidOk() (*int32, bool) {
	if o == nil || IsNil(o.Itemid) {
		return nil, false
	}
	return o.Itemid, true
}

// HasItemid returns a boolean if a field has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) HasItemid() bool {
	if o != nil && !IsNil(o.Itemid) {
		return true
	}

	return false
}

// SetItemid gets a reference to the given int32 and assigns it to the Itemid field.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) SetItemid(v int32) {
	o.Itemid = &v
}

// GetLicense returns the License field value if set, zero value otherwise.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetLicense() string {
	if o == nil || IsNil(o.License) {
		var ret string
		return ret
	}
	return *o.License
}

// GetLicenseOk returns a tuple with the License field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetLicenseOk() (*string, bool) {
	if o == nil || IsNil(o.License) {
		return nil, false
	}
	return o.License, true
}

// HasLicense returns a boolean if a field has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) HasLicense() bool {
	if o != nil && !IsNil(o.License) {
		return true
	}

	return false
}

// SetLicense gets a reference to the given string and assigns it to the License field.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) SetLicense(v string) {
	o.License = &v
}

// GetTimecreated returns the Timecreated field value if set, zero value otherwise.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetTimecreated() int32 {
	if o == nil || IsNil(o.Timecreated) {
		var ret int32
		return ret
	}
	return *o.Timecreated
}

// GetTimecreatedOk returns a tuple with the Timecreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetTimecreatedOk() (*int32, bool) {
	if o == nil || IsNil(o.Timecreated) {
		return nil, false
	}
	return o.Timecreated, true
}

// HasTimecreated returns a boolean if a field has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) HasTimecreated() bool {
	if o != nil && !IsNil(o.Timecreated) {
		return true
	}

	return false
}

// SetTimecreated gets a reference to the given int32 and assigns it to the Timecreated field.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) SetTimecreated(v int32) {
	o.Timecreated = &v
}

// GetTimecreatedformatted returns the Timecreatedformatted field value if set, zero value otherwise.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetTimecreatedformatted() string {
	if o == nil || IsNil(o.Timecreatedformatted) {
		var ret string
		return ret
	}
	return *o.Timecreatedformatted
}

// GetTimecreatedformattedOk returns a tuple with the Timecreatedformatted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetTimecreatedformattedOk() (*string, bool) {
	if o == nil || IsNil(o.Timecreatedformatted) {
		return nil, false
	}
	return o.Timecreatedformatted, true
}

// HasTimecreatedformatted returns a boolean if a field has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) HasTimecreatedformatted() bool {
	if o != nil && !IsNil(o.Timecreatedformatted) {
		return true
	}

	return false
}

// SetTimecreatedformatted gets a reference to the given string and assigns it to the Timecreatedformatted field.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) SetTimecreatedformatted(v string) {
	o.Timecreatedformatted = &v
}

// GetTimemodified returns the Timemodified field value if set, zero value otherwise.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetTimemodified() int32 {
	if o == nil || IsNil(o.Timemodified) {
		var ret int32
		return ret
	}
	return *o.Timemodified
}

// GetTimemodifiedOk returns a tuple with the Timemodified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetTimemodifiedOk() (*int32, bool) {
	if o == nil || IsNil(o.Timemodified) {
		return nil, false
	}
	return o.Timemodified, true
}

// HasTimemodified returns a boolean if a field has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) HasTimemodified() bool {
	if o != nil && !IsNil(o.Timemodified) {
		return true
	}

	return false
}

// SetTimemodified gets a reference to the given int32 and assigns it to the Timemodified field.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) SetTimemodified(v int32) {
	o.Timemodified = &v
}

// GetTimemodifiedformatted returns the Timemodifiedformatted field value if set, zero value otherwise.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetTimemodifiedformatted() string {
	if o == nil || IsNil(o.Timemodifiedformatted) {
		var ret string
		return ret
	}
	return *o.Timemodifiedformatted
}

// GetTimemodifiedformattedOk returns a tuple with the Timemodifiedformatted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetTimemodifiedformattedOk() (*string, bool) {
	if o == nil || IsNil(o.Timemodifiedformatted) {
		return nil, false
	}
	return o.Timemodifiedformatted, true
}

// HasTimemodifiedformatted returns a boolean if a field has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) HasTimemodifiedformatted() bool {
	if o != nil && !IsNil(o.Timemodifiedformatted) {
		return true
	}

	return false
}

// SetTimemodifiedformatted gets a reference to the given string and assigns it to the Timemodifiedformatted field.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) SetTimemodifiedformatted(v string) {
	o.Timemodifiedformatted = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CoreCompetencyReadUserEvidence200ResponseFilesInner) SetUrl(v string) {
	o.Url = &v
}

func (o CoreCompetencyReadUserEvidence200ResponseFilesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCompetencyReadUserEvidence200ResponseFilesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !IsNil(o.Component) {
		toSerialize["component"] = o.Component
	}
	if !IsNil(o.Contextid) {
		toSerialize["contextid"] = o.Contextid
	}
	if !IsNil(o.Filearea) {
		toSerialize["filearea"] = o.Filearea
	}
	if !IsNil(o.Filename) {
		toSerialize["filename"] = o.Filename
	}
	if !IsNil(o.Filenameshort) {
		toSerialize["filenameshort"] = o.Filenameshort
	}
	if !IsNil(o.Filepath) {
		toSerialize["filepath"] = o.Filepath
	}
	if !IsNil(o.Filesize) {
		toSerialize["filesize"] = o.Filesize
	}
	if !IsNil(o.Filesizeformatted) {
		toSerialize["filesizeformatted"] = o.Filesizeformatted
	}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	if !IsNil(o.Isdir) {
		toSerialize["isdir"] = o.Isdir
	}
	if !IsNil(o.Isimage) {
		toSerialize["isimage"] = o.Isimage
	}
	if !IsNil(o.Itemid) {
		toSerialize["itemid"] = o.Itemid
	}
	if !IsNil(o.License) {
		toSerialize["license"] = o.License
	}
	if !IsNil(o.Timecreated) {
		toSerialize["timecreated"] = o.Timecreated
	}
	if !IsNil(o.Timecreatedformatted) {
		toSerialize["timecreatedformatted"] = o.Timecreatedformatted
	}
	if !IsNil(o.Timemodified) {
		toSerialize["timemodified"] = o.Timemodified
	}
	if !IsNil(o.Timemodifiedformatted) {
		toSerialize["timemodifiedformatted"] = o.Timemodifiedformatted
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableCoreCompetencyReadUserEvidence200ResponseFilesInner struct {
	value *CoreCompetencyReadUserEvidence200ResponseFilesInner
	isSet bool
}

func (v NullableCoreCompetencyReadUserEvidence200ResponseFilesInner) Get() *CoreCompetencyReadUserEvidence200ResponseFilesInner {
	return v.value
}

func (v *NullableCoreCompetencyReadUserEvidence200ResponseFilesInner) Set(val *CoreCompetencyReadUserEvidence200ResponseFilesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCompetencyReadUserEvidence200ResponseFilesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCompetencyReadUserEvidence200ResponseFilesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCompetencyReadUserEvidence200ResponseFilesInner(val *CoreCompetencyReadUserEvidence200ResponseFilesInner) *NullableCoreCompetencyReadUserEvidence200ResponseFilesInner {
	return &NullableCoreCompetencyReadUserEvidence200ResponseFilesInner{value: val, isSet: true}
}

func (v NullableCoreCompetencyReadUserEvidence200ResponseFilesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCompetencyReadUserEvidence200ResponseFilesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


