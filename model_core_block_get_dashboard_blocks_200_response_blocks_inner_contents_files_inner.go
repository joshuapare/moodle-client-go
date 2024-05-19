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

// checks if the CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner{}

// CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner File.
type CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner struct {
	// File name.
	Filename *string `json:"filename,omitempty"`
	// File path.
	Filepath *string `json:"filepath,omitempty"`
	// File size.
	Filesize *int32 `json:"filesize,omitempty"`
	// Downloadable file url.
	Fileurl *string `json:"fileurl,omitempty"`
	// Whether is an external file.
	Isexternalfile *bool `json:"isexternalfile,omitempty"`
	// File mime type.
	Mimetype *string `json:"mimetype,omitempty"`
	// The repository type for external files.
	Repositorytype *string `json:"repositorytype,omitempty"`
	// Time modified.
	Timemodified *int32 `json:"timemodified,omitempty"`
}

// NewCoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner instantiates a new CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner() *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner {
	this := CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner{}
	return &this
}

// NewCoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInnerWithDefaults instantiates a new CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInnerWithDefaults() *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner {
	this := CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner{}
	return &this
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) GetFilename() string {
	if o == nil || IsNil(o.Filename) {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) GetFilenameOk() (*string, bool) {
	if o == nil || IsNil(o.Filename) {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) HasFilename() bool {
	if o != nil && !IsNil(o.Filename) {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) SetFilename(v string) {
	o.Filename = &v
}

// GetFilepath returns the Filepath field value if set, zero value otherwise.
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) GetFilepath() string {
	if o == nil || IsNil(o.Filepath) {
		var ret string
		return ret
	}
	return *o.Filepath
}

// GetFilepathOk returns a tuple with the Filepath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) GetFilepathOk() (*string, bool) {
	if o == nil || IsNil(o.Filepath) {
		return nil, false
	}
	return o.Filepath, true
}

// HasFilepath returns a boolean if a field has been set.
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) HasFilepath() bool {
	if o != nil && !IsNil(o.Filepath) {
		return true
	}

	return false
}

// SetFilepath gets a reference to the given string and assigns it to the Filepath field.
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) SetFilepath(v string) {
	o.Filepath = &v
}

// GetFilesize returns the Filesize field value if set, zero value otherwise.
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) GetFilesize() int32 {
	if o == nil || IsNil(o.Filesize) {
		var ret int32
		return ret
	}
	return *o.Filesize
}

// GetFilesizeOk returns a tuple with the Filesize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) GetFilesizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Filesize) {
		return nil, false
	}
	return o.Filesize, true
}

// HasFilesize returns a boolean if a field has been set.
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) HasFilesize() bool {
	if o != nil && !IsNil(o.Filesize) {
		return true
	}

	return false
}

// SetFilesize gets a reference to the given int32 and assigns it to the Filesize field.
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) SetFilesize(v int32) {
	o.Filesize = &v
}

// GetFileurl returns the Fileurl field value if set, zero value otherwise.
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) GetFileurl() string {
	if o == nil || IsNil(o.Fileurl) {
		var ret string
		return ret
	}
	return *o.Fileurl
}

// GetFileurlOk returns a tuple with the Fileurl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) GetFileurlOk() (*string, bool) {
	if o == nil || IsNil(o.Fileurl) {
		return nil, false
	}
	return o.Fileurl, true
}

// HasFileurl returns a boolean if a field has been set.
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) HasFileurl() bool {
	if o != nil && !IsNil(o.Fileurl) {
		return true
	}

	return false
}

// SetFileurl gets a reference to the given string and assigns it to the Fileurl field.
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) SetFileurl(v string) {
	o.Fileurl = &v
}

// GetIsexternalfile returns the Isexternalfile field value if set, zero value otherwise.
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) GetIsexternalfile() bool {
	if o == nil || IsNil(o.Isexternalfile) {
		var ret bool
		return ret
	}
	return *o.Isexternalfile
}

// GetIsexternalfileOk returns a tuple with the Isexternalfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) GetIsexternalfileOk() (*bool, bool) {
	if o == nil || IsNil(o.Isexternalfile) {
		return nil, false
	}
	return o.Isexternalfile, true
}

// HasIsexternalfile returns a boolean if a field has been set.
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) HasIsexternalfile() bool {
	if o != nil && !IsNil(o.Isexternalfile) {
		return true
	}

	return false
}

// SetIsexternalfile gets a reference to the given bool and assigns it to the Isexternalfile field.
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) SetIsexternalfile(v bool) {
	o.Isexternalfile = &v
}

// GetMimetype returns the Mimetype field value if set, zero value otherwise.
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) GetMimetype() string {
	if o == nil || IsNil(o.Mimetype) {
		var ret string
		return ret
	}
	return *o.Mimetype
}

// GetMimetypeOk returns a tuple with the Mimetype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) GetMimetypeOk() (*string, bool) {
	if o == nil || IsNil(o.Mimetype) {
		return nil, false
	}
	return o.Mimetype, true
}

// HasMimetype returns a boolean if a field has been set.
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) HasMimetype() bool {
	if o != nil && !IsNil(o.Mimetype) {
		return true
	}

	return false
}

// SetMimetype gets a reference to the given string and assigns it to the Mimetype field.
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) SetMimetype(v string) {
	o.Mimetype = &v
}

// GetRepositorytype returns the Repositorytype field value if set, zero value otherwise.
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) GetRepositorytype() string {
	if o == nil || IsNil(o.Repositorytype) {
		var ret string
		return ret
	}
	return *o.Repositorytype
}

// GetRepositorytypeOk returns a tuple with the Repositorytype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) GetRepositorytypeOk() (*string, bool) {
	if o == nil || IsNil(o.Repositorytype) {
		return nil, false
	}
	return o.Repositorytype, true
}

// HasRepositorytype returns a boolean if a field has been set.
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) HasRepositorytype() bool {
	if o != nil && !IsNil(o.Repositorytype) {
		return true
	}

	return false
}

// SetRepositorytype gets a reference to the given string and assigns it to the Repositorytype field.
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) SetRepositorytype(v string) {
	o.Repositorytype = &v
}

// GetTimemodified returns the Timemodified field value if set, zero value otherwise.
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) GetTimemodified() int32 {
	if o == nil || IsNil(o.Timemodified) {
		var ret int32
		return ret
	}
	return *o.Timemodified
}

// GetTimemodifiedOk returns a tuple with the Timemodified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) GetTimemodifiedOk() (*int32, bool) {
	if o == nil || IsNil(o.Timemodified) {
		return nil, false
	}
	return o.Timemodified, true
}

// HasTimemodified returns a boolean if a field has been set.
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) HasTimemodified() bool {
	if o != nil && !IsNil(o.Timemodified) {
		return true
	}

	return false
}

// SetTimemodified gets a reference to the given int32 and assigns it to the Timemodified field.
func (o *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) SetTimemodified(v int32) {
	o.Timemodified = &v
}

func (o CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Filename) {
		toSerialize["filename"] = o.Filename
	}
	if !IsNil(o.Filepath) {
		toSerialize["filepath"] = o.Filepath
	}
	if !IsNil(o.Filesize) {
		toSerialize["filesize"] = o.Filesize
	}
	if !IsNil(o.Fileurl) {
		toSerialize["fileurl"] = o.Fileurl
	}
	if !IsNil(o.Isexternalfile) {
		toSerialize["isexternalfile"] = o.Isexternalfile
	}
	if !IsNil(o.Mimetype) {
		toSerialize["mimetype"] = o.Mimetype
	}
	if !IsNil(o.Repositorytype) {
		toSerialize["repositorytype"] = o.Repositorytype
	}
	if !IsNil(o.Timemodified) {
		toSerialize["timemodified"] = o.Timemodified
	}
	return toSerialize, nil
}

type NullableCoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner struct {
	value *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner
	isSet bool
}

func (v NullableCoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) Get() *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner {
	return v.value
}

func (v *NullableCoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) Set(val *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner(val *CoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) *NullableCoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner {
	return &NullableCoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner{value: val, isSet: true}
}

func (v NullableCoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreBlockGetDashboardBlocks200ResponseBlocksInnerContentsFilesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


