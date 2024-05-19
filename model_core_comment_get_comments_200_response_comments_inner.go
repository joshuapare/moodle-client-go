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

// checks if the CoreCommentGetComments200ResponseCommentsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCommentGetComments200ResponseCommentsInner{}

// CoreCommentGetComments200ResponseCommentsInner comment
type CoreCommentGetComments200ResponseCommentsInner struct {
	// HTML user picture
	Avatar *string `json:"avatar,omitempty"`
	// The content text formatted
	Content *string `json:"content,omitempty"`
	// Permission to delete=true/false
	Delete *bool `json:"delete,omitempty"`
	// content format (1 = HTML, 0 = MOODLE, 2 = PLAIN, or 4 = MARKDOWN)
	Format *int32 `json:"format,omitempty"`
	// fullname
	Fullname *string `json:"fullname,omitempty"`
	// Comment ID
	Id *int32 `json:"id,omitempty"`
	// URL profile
	Profileurl *string `json:"profileurl,omitempty"`
	// Time format
	Strftimeformat *string `json:"strftimeformat,omitempty"`
	// Time in human format
	Time *string `json:"time,omitempty"`
	// Time created (timestamp)
	Timecreated *int32 `json:"timecreated,omitempty"`
	// User ID
	Userid *int32 `json:"userid,omitempty"`
}

// NewCoreCommentGetComments200ResponseCommentsInner instantiates a new CoreCommentGetComments200ResponseCommentsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCommentGetComments200ResponseCommentsInner() *CoreCommentGetComments200ResponseCommentsInner {
	this := CoreCommentGetComments200ResponseCommentsInner{}
	var avatar string = "null"
	this.Avatar = &avatar
	var content string = "null"
	this.Content = &content
	var delete bool = null
	this.Delete = &delete
	var fullname string = "null"
	this.Fullname = &fullname
	var id int32 = null
	this.Id = &id
	var profileurl string = "null"
	this.Profileurl = &profileurl
	var strftimeformat string = "null"
	this.Strftimeformat = &strftimeformat
	var time string = "null"
	this.Time = &time
	var timecreated int32 = null
	this.Timecreated = &timecreated
	return &this
}

// NewCoreCommentGetComments200ResponseCommentsInnerWithDefaults instantiates a new CoreCommentGetComments200ResponseCommentsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCommentGetComments200ResponseCommentsInnerWithDefaults() *CoreCommentGetComments200ResponseCommentsInner {
	this := CoreCommentGetComments200ResponseCommentsInner{}
	var avatar string = "null"
	this.Avatar = &avatar
	var content string = "null"
	this.Content = &content
	var delete bool = null
	this.Delete = &delete
	var fullname string = "null"
	this.Fullname = &fullname
	var id int32 = null
	this.Id = &id
	var profileurl string = "null"
	this.Profileurl = &profileurl
	var strftimeformat string = "null"
	this.Strftimeformat = &strftimeformat
	var time string = "null"
	this.Time = &time
	var timecreated int32 = null
	this.Timecreated = &timecreated
	return &this
}

// GetAvatar returns the Avatar field value if set, zero value otherwise.
func (o *CoreCommentGetComments200ResponseCommentsInner) GetAvatar() string {
	if o == nil || IsNil(o.Avatar) {
		var ret string
		return ret
	}
	return *o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCommentGetComments200ResponseCommentsInner) GetAvatarOk() (*string, bool) {
	if o == nil || IsNil(o.Avatar) {
		return nil, false
	}
	return o.Avatar, true
}

// HasAvatar returns a boolean if a field has been set.
func (o *CoreCommentGetComments200ResponseCommentsInner) HasAvatar() bool {
	if o != nil && !IsNil(o.Avatar) {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given string and assigns it to the Avatar field.
func (o *CoreCommentGetComments200ResponseCommentsInner) SetAvatar(v string) {
	o.Avatar = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *CoreCommentGetComments200ResponseCommentsInner) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCommentGetComments200ResponseCommentsInner) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *CoreCommentGetComments200ResponseCommentsInner) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *CoreCommentGetComments200ResponseCommentsInner) SetContent(v string) {
	o.Content = &v
}

// GetDelete returns the Delete field value if set, zero value otherwise.
func (o *CoreCommentGetComments200ResponseCommentsInner) GetDelete() bool {
	if o == nil || IsNil(o.Delete) {
		var ret bool
		return ret
	}
	return *o.Delete
}

// GetDeleteOk returns a tuple with the Delete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCommentGetComments200ResponseCommentsInner) GetDeleteOk() (*bool, bool) {
	if o == nil || IsNil(o.Delete) {
		return nil, false
	}
	return o.Delete, true
}

// HasDelete returns a boolean if a field has been set.
func (o *CoreCommentGetComments200ResponseCommentsInner) HasDelete() bool {
	if o != nil && !IsNil(o.Delete) {
		return true
	}

	return false
}

// SetDelete gets a reference to the given bool and assigns it to the Delete field.
func (o *CoreCommentGetComments200ResponseCommentsInner) SetDelete(v bool) {
	o.Delete = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *CoreCommentGetComments200ResponseCommentsInner) GetFormat() int32 {
	if o == nil || IsNil(o.Format) {
		var ret int32
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCommentGetComments200ResponseCommentsInner) GetFormatOk() (*int32, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *CoreCommentGetComments200ResponseCommentsInner) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given int32 and assigns it to the Format field.
func (o *CoreCommentGetComments200ResponseCommentsInner) SetFormat(v int32) {
	o.Format = &v
}

// GetFullname returns the Fullname field value if set, zero value otherwise.
func (o *CoreCommentGetComments200ResponseCommentsInner) GetFullname() string {
	if o == nil || IsNil(o.Fullname) {
		var ret string
		return ret
	}
	return *o.Fullname
}

// GetFullnameOk returns a tuple with the Fullname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCommentGetComments200ResponseCommentsInner) GetFullnameOk() (*string, bool) {
	if o == nil || IsNil(o.Fullname) {
		return nil, false
	}
	return o.Fullname, true
}

// HasFullname returns a boolean if a field has been set.
func (o *CoreCommentGetComments200ResponseCommentsInner) HasFullname() bool {
	if o != nil && !IsNil(o.Fullname) {
		return true
	}

	return false
}

// SetFullname gets a reference to the given string and assigns it to the Fullname field.
func (o *CoreCommentGetComments200ResponseCommentsInner) SetFullname(v string) {
	o.Fullname = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CoreCommentGetComments200ResponseCommentsInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCommentGetComments200ResponseCommentsInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CoreCommentGetComments200ResponseCommentsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CoreCommentGetComments200ResponseCommentsInner) SetId(v int32) {
	o.Id = &v
}

// GetProfileurl returns the Profileurl field value if set, zero value otherwise.
func (o *CoreCommentGetComments200ResponseCommentsInner) GetProfileurl() string {
	if o == nil || IsNil(o.Profileurl) {
		var ret string
		return ret
	}
	return *o.Profileurl
}

// GetProfileurlOk returns a tuple with the Profileurl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCommentGetComments200ResponseCommentsInner) GetProfileurlOk() (*string, bool) {
	if o == nil || IsNil(o.Profileurl) {
		return nil, false
	}
	return o.Profileurl, true
}

// HasProfileurl returns a boolean if a field has been set.
func (o *CoreCommentGetComments200ResponseCommentsInner) HasProfileurl() bool {
	if o != nil && !IsNil(o.Profileurl) {
		return true
	}

	return false
}

// SetProfileurl gets a reference to the given string and assigns it to the Profileurl field.
func (o *CoreCommentGetComments200ResponseCommentsInner) SetProfileurl(v string) {
	o.Profileurl = &v
}

// GetStrftimeformat returns the Strftimeformat field value if set, zero value otherwise.
func (o *CoreCommentGetComments200ResponseCommentsInner) GetStrftimeformat() string {
	if o == nil || IsNil(o.Strftimeformat) {
		var ret string
		return ret
	}
	return *o.Strftimeformat
}

// GetStrftimeformatOk returns a tuple with the Strftimeformat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCommentGetComments200ResponseCommentsInner) GetStrftimeformatOk() (*string, bool) {
	if o == nil || IsNil(o.Strftimeformat) {
		return nil, false
	}
	return o.Strftimeformat, true
}

// HasStrftimeformat returns a boolean if a field has been set.
func (o *CoreCommentGetComments200ResponseCommentsInner) HasStrftimeformat() bool {
	if o != nil && !IsNil(o.Strftimeformat) {
		return true
	}

	return false
}

// SetStrftimeformat gets a reference to the given string and assigns it to the Strftimeformat field.
func (o *CoreCommentGetComments200ResponseCommentsInner) SetStrftimeformat(v string) {
	o.Strftimeformat = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *CoreCommentGetComments200ResponseCommentsInner) GetTime() string {
	if o == nil || IsNil(o.Time) {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCommentGetComments200ResponseCommentsInner) GetTimeOk() (*string, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *CoreCommentGetComments200ResponseCommentsInner) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *CoreCommentGetComments200ResponseCommentsInner) SetTime(v string) {
	o.Time = &v
}

// GetTimecreated returns the Timecreated field value if set, zero value otherwise.
func (o *CoreCommentGetComments200ResponseCommentsInner) GetTimecreated() int32 {
	if o == nil || IsNil(o.Timecreated) {
		var ret int32
		return ret
	}
	return *o.Timecreated
}

// GetTimecreatedOk returns a tuple with the Timecreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCommentGetComments200ResponseCommentsInner) GetTimecreatedOk() (*int32, bool) {
	if o == nil || IsNil(o.Timecreated) {
		return nil, false
	}
	return o.Timecreated, true
}

// HasTimecreated returns a boolean if a field has been set.
func (o *CoreCommentGetComments200ResponseCommentsInner) HasTimecreated() bool {
	if o != nil && !IsNil(o.Timecreated) {
		return true
	}

	return false
}

// SetTimecreated gets a reference to the given int32 and assigns it to the Timecreated field.
func (o *CoreCommentGetComments200ResponseCommentsInner) SetTimecreated(v int32) {
	o.Timecreated = &v
}

// GetUserid returns the Userid field value if set, zero value otherwise.
func (o *CoreCommentGetComments200ResponseCommentsInner) GetUserid() int32 {
	if o == nil || IsNil(o.Userid) {
		var ret int32
		return ret
	}
	return *o.Userid
}

// GetUseridOk returns a tuple with the Userid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCommentGetComments200ResponseCommentsInner) GetUseridOk() (*int32, bool) {
	if o == nil || IsNil(o.Userid) {
		return nil, false
	}
	return o.Userid, true
}

// HasUserid returns a boolean if a field has been set.
func (o *CoreCommentGetComments200ResponseCommentsInner) HasUserid() bool {
	if o != nil && !IsNil(o.Userid) {
		return true
	}

	return false
}

// SetUserid gets a reference to the given int32 and assigns it to the Userid field.
func (o *CoreCommentGetComments200ResponseCommentsInner) SetUserid(v int32) {
	o.Userid = &v
}

func (o CoreCommentGetComments200ResponseCommentsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCommentGetComments200ResponseCommentsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Avatar) {
		toSerialize["avatar"] = o.Avatar
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.Delete) {
		toSerialize["delete"] = o.Delete
	}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !IsNil(o.Fullname) {
		toSerialize["fullname"] = o.Fullname
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Profileurl) {
		toSerialize["profileurl"] = o.Profileurl
	}
	if !IsNil(o.Strftimeformat) {
		toSerialize["strftimeformat"] = o.Strftimeformat
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.Timecreated) {
		toSerialize["timecreated"] = o.Timecreated
	}
	if !IsNil(o.Userid) {
		toSerialize["userid"] = o.Userid
	}
	return toSerialize, nil
}

type NullableCoreCommentGetComments200ResponseCommentsInner struct {
	value *CoreCommentGetComments200ResponseCommentsInner
	isSet bool
}

func (v NullableCoreCommentGetComments200ResponseCommentsInner) Get() *CoreCommentGetComments200ResponseCommentsInner {
	return v.value
}

func (v *NullableCoreCommentGetComments200ResponseCommentsInner) Set(val *CoreCommentGetComments200ResponseCommentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCommentGetComments200ResponseCommentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCommentGetComments200ResponseCommentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCommentGetComments200ResponseCommentsInner(val *CoreCommentGetComments200ResponseCommentsInner) *NullableCoreCommentGetComments200ResponseCommentsInner {
	return &NullableCoreCommentGetComments200ResponseCommentsInner{value: val, isSet: true}
}

func (v NullableCoreCommentGetComments200ResponseCommentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCommentGetComments200ResponseCommentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


