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

// checks if the ModForumAddDiscussionPost200ResponsePost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModForumAddDiscussionPost200ResponsePost{}

// ModForumAddDiscussionPost200ResponsePost struct for ModForumAddDiscussionPost200ResponsePost
type ModForumAddDiscussionPost200ResponsePost struct {
	Attachments []ModForumAddDiscussionPost200ResponsePostAttachmentsInner `json:"attachments"`
	Author ModForumAddDiscussionPost200ResponsePostAuthor `json:"author"`
	Capabilities ModForumAddDiscussionPost200ResponsePostCapabilities `json:"capabilities"`
	// charcount
	Charcount *int32 `json:"charcount,omitempty"`
	// discussionid
	Discussionid int32 `json:"discussionid"`
	// hasparent
	Hasparent bool `json:"hasparent"`
	// haswordcount
	Haswordcount bool `json:"haswordcount"`
	Html *ModForumAddDiscussionPost200ResponsePostHtml `json:"html,omitempty"`
	// id
	Id int32 `json:"id"`
	// isdeleted
	Isdeleted bool `json:"isdeleted"`
	// isprivatereply
	Isprivatereply bool `json:"isprivatereply"`
	// message
	Message string `json:"message"`
	// message format (1 = HTML, 0 = MOODLE, 2 = PLAIN, or 4 = MARKDOWN)
	Messageformat int32 `json:"messageformat"`
	Messageinlinefiles []ModFeedbackGetAnalysis200ResponseItemsdataInnerItemItemfilesInner `json:"messageinlinefiles,omitempty"`
	// parentid
	Parentid *int32 `json:"parentid,omitempty"`
	// replysubject
	Replysubject string `json:"replysubject"`
	// subject
	Subject string `json:"subject"`
	Tags []ModForumAddDiscussionPost200ResponsePostTagsInner `json:"tags,omitempty"`
	// timecreated
	Timecreated int32 `json:"timecreated"`
	// timemodified
	Timemodified int32 `json:"timemodified"`
	// unread
	Unread *bool `json:"unread,omitempty"`
	Urls *ModForumAddDiscussionPost200ResponsePostUrls `json:"urls,omitempty"`
	// wordcount
	Wordcount *int32 `json:"wordcount,omitempty"`
}

type _ModForumAddDiscussionPost200ResponsePost ModForumAddDiscussionPost200ResponsePost

// NewModForumAddDiscussionPost200ResponsePost instantiates a new ModForumAddDiscussionPost200ResponsePost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModForumAddDiscussionPost200ResponsePost(attachments []ModForumAddDiscussionPost200ResponsePostAttachmentsInner, author ModForumAddDiscussionPost200ResponsePostAuthor, capabilities ModForumAddDiscussionPost200ResponsePostCapabilities, discussionid int32, hasparent bool, haswordcount bool, id int32, isdeleted bool, isprivatereply bool, message string, messageformat int32, replysubject string, subject string, timecreated int32, timemodified int32) *ModForumAddDiscussionPost200ResponsePost {
	this := ModForumAddDiscussionPost200ResponsePost{}
	this.Attachments = attachments
	this.Author = author
	this.Capabilities = capabilities
	var charcount int32 = null
	this.Charcount = &charcount
	this.Discussionid = discussionid
	this.Hasparent = hasparent
	this.Haswordcount = haswordcount
	this.Id = id
	this.Isdeleted = isdeleted
	this.Isprivatereply = isprivatereply
	this.Message = message
	this.Messageformat = messageformat
	var parentid int32 = null
	this.Parentid = &parentid
	this.Replysubject = replysubject
	this.Subject = subject
	this.Timecreated = timecreated
	this.Timemodified = timemodified
	var unread bool = null
	this.Unread = &unread
	var wordcount int32 = null
	this.Wordcount = &wordcount
	return &this
}

// NewModForumAddDiscussionPost200ResponsePostWithDefaults instantiates a new ModForumAddDiscussionPost200ResponsePost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModForumAddDiscussionPost200ResponsePostWithDefaults() *ModForumAddDiscussionPost200ResponsePost {
	this := ModForumAddDiscussionPost200ResponsePost{}
	var charcount int32 = null
	this.Charcount = &charcount
	var discussionid int32 = null
	this.Discussionid = discussionid
	var hasparent bool = null
	this.Hasparent = hasparent
	var haswordcount bool = null
	this.Haswordcount = haswordcount
	var isdeleted bool = null
	this.Isdeleted = isdeleted
	var isprivatereply bool = null
	this.Isprivatereply = isprivatereply
	var message string = "null"
	this.Message = message
	var messageformat int32 = null
	this.Messageformat = messageformat
	var parentid int32 = null
	this.Parentid = &parentid
	var replysubject string = "null"
	this.Replysubject = replysubject
	var subject string = "null"
	this.Subject = subject
	var timecreated int32 = null
	this.Timecreated = timecreated
	var timemodified int32 = null
	this.Timemodified = timemodified
	var unread bool = null
	this.Unread = &unread
	var wordcount int32 = null
	this.Wordcount = &wordcount
	return &this
}

// GetAttachments returns the Attachments field value
func (o *ModForumAddDiscussionPost200ResponsePost) GetAttachments() []ModForumAddDiscussionPost200ResponsePostAttachmentsInner {
	if o == nil {
		var ret []ModForumAddDiscussionPost200ResponsePostAttachmentsInner
		return ret
	}

	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value
// and a boolean to check if the value has been set.
func (o *ModForumAddDiscussionPost200ResponsePost) GetAttachmentsOk() ([]ModForumAddDiscussionPost200ResponsePostAttachmentsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachments, true
}

// SetAttachments sets field value
func (o *ModForumAddDiscussionPost200ResponsePost) SetAttachments(v []ModForumAddDiscussionPost200ResponsePostAttachmentsInner) {
	o.Attachments = v
}

// GetAuthor returns the Author field value
func (o *ModForumAddDiscussionPost200ResponsePost) GetAuthor() ModForumAddDiscussionPost200ResponsePostAuthor {
	if o == nil {
		var ret ModForumAddDiscussionPost200ResponsePostAuthor
		return ret
	}

	return o.Author
}

// GetAuthorOk returns a tuple with the Author field value
// and a boolean to check if the value has been set.
func (o *ModForumAddDiscussionPost200ResponsePost) GetAuthorOk() (*ModForumAddDiscussionPost200ResponsePostAuthor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Author, true
}

// SetAuthor sets field value
func (o *ModForumAddDiscussionPost200ResponsePost) SetAuthor(v ModForumAddDiscussionPost200ResponsePostAuthor) {
	o.Author = v
}

// GetCapabilities returns the Capabilities field value
func (o *ModForumAddDiscussionPost200ResponsePost) GetCapabilities() ModForumAddDiscussionPost200ResponsePostCapabilities {
	if o == nil {
		var ret ModForumAddDiscussionPost200ResponsePostCapabilities
		return ret
	}

	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value
// and a boolean to check if the value has been set.
func (o *ModForumAddDiscussionPost200ResponsePost) GetCapabilitiesOk() (*ModForumAddDiscussionPost200ResponsePostCapabilities, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Capabilities, true
}

// SetCapabilities sets field value
func (o *ModForumAddDiscussionPost200ResponsePost) SetCapabilities(v ModForumAddDiscussionPost200ResponsePostCapabilities) {
	o.Capabilities = v
}

// GetCharcount returns the Charcount field value if set, zero value otherwise.
func (o *ModForumAddDiscussionPost200ResponsePost) GetCharcount() int32 {
	if o == nil || IsNil(o.Charcount) {
		var ret int32
		return ret
	}
	return *o.Charcount
}

// GetCharcountOk returns a tuple with the Charcount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumAddDiscussionPost200ResponsePost) GetCharcountOk() (*int32, bool) {
	if o == nil || IsNil(o.Charcount) {
		return nil, false
	}
	return o.Charcount, true
}

// HasCharcount returns a boolean if a field has been set.
func (o *ModForumAddDiscussionPost200ResponsePost) HasCharcount() bool {
	if o != nil && !IsNil(o.Charcount) {
		return true
	}

	return false
}

// SetCharcount gets a reference to the given int32 and assigns it to the Charcount field.
func (o *ModForumAddDiscussionPost200ResponsePost) SetCharcount(v int32) {
	o.Charcount = &v
}

// GetDiscussionid returns the Discussionid field value
func (o *ModForumAddDiscussionPost200ResponsePost) GetDiscussionid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Discussionid
}

// GetDiscussionidOk returns a tuple with the Discussionid field value
// and a boolean to check if the value has been set.
func (o *ModForumAddDiscussionPost200ResponsePost) GetDiscussionidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Discussionid, true
}

// SetDiscussionid sets field value
func (o *ModForumAddDiscussionPost200ResponsePost) SetDiscussionid(v int32) {
	o.Discussionid = v
}

// GetHasparent returns the Hasparent field value
func (o *ModForumAddDiscussionPost200ResponsePost) GetHasparent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Hasparent
}

// GetHasparentOk returns a tuple with the Hasparent field value
// and a boolean to check if the value has been set.
func (o *ModForumAddDiscussionPost200ResponsePost) GetHasparentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hasparent, true
}

// SetHasparent sets field value
func (o *ModForumAddDiscussionPost200ResponsePost) SetHasparent(v bool) {
	o.Hasparent = v
}

// GetHaswordcount returns the Haswordcount field value
func (o *ModForumAddDiscussionPost200ResponsePost) GetHaswordcount() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Haswordcount
}

// GetHaswordcountOk returns a tuple with the Haswordcount field value
// and a boolean to check if the value has been set.
func (o *ModForumAddDiscussionPost200ResponsePost) GetHaswordcountOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Haswordcount, true
}

// SetHaswordcount sets field value
func (o *ModForumAddDiscussionPost200ResponsePost) SetHaswordcount(v bool) {
	o.Haswordcount = v
}

// GetHtml returns the Html field value if set, zero value otherwise.
func (o *ModForumAddDiscussionPost200ResponsePost) GetHtml() ModForumAddDiscussionPost200ResponsePostHtml {
	if o == nil || IsNil(o.Html) {
		var ret ModForumAddDiscussionPost200ResponsePostHtml
		return ret
	}
	return *o.Html
}

// GetHtmlOk returns a tuple with the Html field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumAddDiscussionPost200ResponsePost) GetHtmlOk() (*ModForumAddDiscussionPost200ResponsePostHtml, bool) {
	if o == nil || IsNil(o.Html) {
		return nil, false
	}
	return o.Html, true
}

// HasHtml returns a boolean if a field has been set.
func (o *ModForumAddDiscussionPost200ResponsePost) HasHtml() bool {
	if o != nil && !IsNil(o.Html) {
		return true
	}

	return false
}

// SetHtml gets a reference to the given ModForumAddDiscussionPost200ResponsePostHtml and assigns it to the Html field.
func (o *ModForumAddDiscussionPost200ResponsePost) SetHtml(v ModForumAddDiscussionPost200ResponsePostHtml) {
	o.Html = &v
}

// GetId returns the Id field value
func (o *ModForumAddDiscussionPost200ResponsePost) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ModForumAddDiscussionPost200ResponsePost) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ModForumAddDiscussionPost200ResponsePost) SetId(v int32) {
	o.Id = v
}

// GetIsdeleted returns the Isdeleted field value
func (o *ModForumAddDiscussionPost200ResponsePost) GetIsdeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Isdeleted
}

// GetIsdeletedOk returns a tuple with the Isdeleted field value
// and a boolean to check if the value has been set.
func (o *ModForumAddDiscussionPost200ResponsePost) GetIsdeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Isdeleted, true
}

// SetIsdeleted sets field value
func (o *ModForumAddDiscussionPost200ResponsePost) SetIsdeleted(v bool) {
	o.Isdeleted = v
}

// GetIsprivatereply returns the Isprivatereply field value
func (o *ModForumAddDiscussionPost200ResponsePost) GetIsprivatereply() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Isprivatereply
}

// GetIsprivatereplyOk returns a tuple with the Isprivatereply field value
// and a boolean to check if the value has been set.
func (o *ModForumAddDiscussionPost200ResponsePost) GetIsprivatereplyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Isprivatereply, true
}

// SetIsprivatereply sets field value
func (o *ModForumAddDiscussionPost200ResponsePost) SetIsprivatereply(v bool) {
	o.Isprivatereply = v
}

// GetMessage returns the Message field value
func (o *ModForumAddDiscussionPost200ResponsePost) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ModForumAddDiscussionPost200ResponsePost) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ModForumAddDiscussionPost200ResponsePost) SetMessage(v string) {
	o.Message = v
}

// GetMessageformat returns the Messageformat field value
func (o *ModForumAddDiscussionPost200ResponsePost) GetMessageformat() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Messageformat
}

// GetMessageformatOk returns a tuple with the Messageformat field value
// and a boolean to check if the value has been set.
func (o *ModForumAddDiscussionPost200ResponsePost) GetMessageformatOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Messageformat, true
}

// SetMessageformat sets field value
func (o *ModForumAddDiscussionPost200ResponsePost) SetMessageformat(v int32) {
	o.Messageformat = v
}

// GetMessageinlinefiles returns the Messageinlinefiles field value if set, zero value otherwise.
func (o *ModForumAddDiscussionPost200ResponsePost) GetMessageinlinefiles() []ModFeedbackGetAnalysis200ResponseItemsdataInnerItemItemfilesInner {
	if o == nil || IsNil(o.Messageinlinefiles) {
		var ret []ModFeedbackGetAnalysis200ResponseItemsdataInnerItemItemfilesInner
		return ret
	}
	return o.Messageinlinefiles
}

// GetMessageinlinefilesOk returns a tuple with the Messageinlinefiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumAddDiscussionPost200ResponsePost) GetMessageinlinefilesOk() ([]ModFeedbackGetAnalysis200ResponseItemsdataInnerItemItemfilesInner, bool) {
	if o == nil || IsNil(o.Messageinlinefiles) {
		return nil, false
	}
	return o.Messageinlinefiles, true
}

// HasMessageinlinefiles returns a boolean if a field has been set.
func (o *ModForumAddDiscussionPost200ResponsePost) HasMessageinlinefiles() bool {
	if o != nil && !IsNil(o.Messageinlinefiles) {
		return true
	}

	return false
}

// SetMessageinlinefiles gets a reference to the given []ModFeedbackGetAnalysis200ResponseItemsdataInnerItemItemfilesInner and assigns it to the Messageinlinefiles field.
func (o *ModForumAddDiscussionPost200ResponsePost) SetMessageinlinefiles(v []ModFeedbackGetAnalysis200ResponseItemsdataInnerItemItemfilesInner) {
	o.Messageinlinefiles = v
}

// GetParentid returns the Parentid field value if set, zero value otherwise.
func (o *ModForumAddDiscussionPost200ResponsePost) GetParentid() int32 {
	if o == nil || IsNil(o.Parentid) {
		var ret int32
		return ret
	}
	return *o.Parentid
}

// GetParentidOk returns a tuple with the Parentid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumAddDiscussionPost200ResponsePost) GetParentidOk() (*int32, bool) {
	if o == nil || IsNil(o.Parentid) {
		return nil, false
	}
	return o.Parentid, true
}

// HasParentid returns a boolean if a field has been set.
func (o *ModForumAddDiscussionPost200ResponsePost) HasParentid() bool {
	if o != nil && !IsNil(o.Parentid) {
		return true
	}

	return false
}

// SetParentid gets a reference to the given int32 and assigns it to the Parentid field.
func (o *ModForumAddDiscussionPost200ResponsePost) SetParentid(v int32) {
	o.Parentid = &v
}

// GetReplysubject returns the Replysubject field value
func (o *ModForumAddDiscussionPost200ResponsePost) GetReplysubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Replysubject
}

// GetReplysubjectOk returns a tuple with the Replysubject field value
// and a boolean to check if the value has been set.
func (o *ModForumAddDiscussionPost200ResponsePost) GetReplysubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Replysubject, true
}

// SetReplysubject sets field value
func (o *ModForumAddDiscussionPost200ResponsePost) SetReplysubject(v string) {
	o.Replysubject = v
}

// GetSubject returns the Subject field value
func (o *ModForumAddDiscussionPost200ResponsePost) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *ModForumAddDiscussionPost200ResponsePost) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *ModForumAddDiscussionPost200ResponsePost) SetSubject(v string) {
	o.Subject = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ModForumAddDiscussionPost200ResponsePost) GetTags() []ModForumAddDiscussionPost200ResponsePostTagsInner {
	if o == nil || IsNil(o.Tags) {
		var ret []ModForumAddDiscussionPost200ResponsePostTagsInner
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumAddDiscussionPost200ResponsePost) GetTagsOk() ([]ModForumAddDiscussionPost200ResponsePostTagsInner, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ModForumAddDiscussionPost200ResponsePost) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ModForumAddDiscussionPost200ResponsePostTagsInner and assigns it to the Tags field.
func (o *ModForumAddDiscussionPost200ResponsePost) SetTags(v []ModForumAddDiscussionPost200ResponsePostTagsInner) {
	o.Tags = v
}

// GetTimecreated returns the Timecreated field value
func (o *ModForumAddDiscussionPost200ResponsePost) GetTimecreated() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timecreated
}

// GetTimecreatedOk returns a tuple with the Timecreated field value
// and a boolean to check if the value has been set.
func (o *ModForumAddDiscussionPost200ResponsePost) GetTimecreatedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timecreated, true
}

// SetTimecreated sets field value
func (o *ModForumAddDiscussionPost200ResponsePost) SetTimecreated(v int32) {
	o.Timecreated = v
}

// GetTimemodified returns the Timemodified field value
func (o *ModForumAddDiscussionPost200ResponsePost) GetTimemodified() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timemodified
}

// GetTimemodifiedOk returns a tuple with the Timemodified field value
// and a boolean to check if the value has been set.
func (o *ModForumAddDiscussionPost200ResponsePost) GetTimemodifiedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timemodified, true
}

// SetTimemodified sets field value
func (o *ModForumAddDiscussionPost200ResponsePost) SetTimemodified(v int32) {
	o.Timemodified = v
}

// GetUnread returns the Unread field value if set, zero value otherwise.
func (o *ModForumAddDiscussionPost200ResponsePost) GetUnread() bool {
	if o == nil || IsNil(o.Unread) {
		var ret bool
		return ret
	}
	return *o.Unread
}

// GetUnreadOk returns a tuple with the Unread field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumAddDiscussionPost200ResponsePost) GetUnreadOk() (*bool, bool) {
	if o == nil || IsNil(o.Unread) {
		return nil, false
	}
	return o.Unread, true
}

// HasUnread returns a boolean if a field has been set.
func (o *ModForumAddDiscussionPost200ResponsePost) HasUnread() bool {
	if o != nil && !IsNil(o.Unread) {
		return true
	}

	return false
}

// SetUnread gets a reference to the given bool and assigns it to the Unread field.
func (o *ModForumAddDiscussionPost200ResponsePost) SetUnread(v bool) {
	o.Unread = &v
}

// GetUrls returns the Urls field value if set, zero value otherwise.
func (o *ModForumAddDiscussionPost200ResponsePost) GetUrls() ModForumAddDiscussionPost200ResponsePostUrls {
	if o == nil || IsNil(o.Urls) {
		var ret ModForumAddDiscussionPost200ResponsePostUrls
		return ret
	}
	return *o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumAddDiscussionPost200ResponsePost) GetUrlsOk() (*ModForumAddDiscussionPost200ResponsePostUrls, bool) {
	if o == nil || IsNil(o.Urls) {
		return nil, false
	}
	return o.Urls, true
}

// HasUrls returns a boolean if a field has been set.
func (o *ModForumAddDiscussionPost200ResponsePost) HasUrls() bool {
	if o != nil && !IsNil(o.Urls) {
		return true
	}

	return false
}

// SetUrls gets a reference to the given ModForumAddDiscussionPost200ResponsePostUrls and assigns it to the Urls field.
func (o *ModForumAddDiscussionPost200ResponsePost) SetUrls(v ModForumAddDiscussionPost200ResponsePostUrls) {
	o.Urls = &v
}

// GetWordcount returns the Wordcount field value if set, zero value otherwise.
func (o *ModForumAddDiscussionPost200ResponsePost) GetWordcount() int32 {
	if o == nil || IsNil(o.Wordcount) {
		var ret int32
		return ret
	}
	return *o.Wordcount
}

// GetWordcountOk returns a tuple with the Wordcount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModForumAddDiscussionPost200ResponsePost) GetWordcountOk() (*int32, bool) {
	if o == nil || IsNil(o.Wordcount) {
		return nil, false
	}
	return o.Wordcount, true
}

// HasWordcount returns a boolean if a field has been set.
func (o *ModForumAddDiscussionPost200ResponsePost) HasWordcount() bool {
	if o != nil && !IsNil(o.Wordcount) {
		return true
	}

	return false
}

// SetWordcount gets a reference to the given int32 and assigns it to the Wordcount field.
func (o *ModForumAddDiscussionPost200ResponsePost) SetWordcount(v int32) {
	o.Wordcount = &v
}

func (o ModForumAddDiscussionPost200ResponsePost) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModForumAddDiscussionPost200ResponsePost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attachments"] = o.Attachments
	toSerialize["author"] = o.Author
	toSerialize["capabilities"] = o.Capabilities
	if !IsNil(o.Charcount) {
		toSerialize["charcount"] = o.Charcount
	}
	toSerialize["discussionid"] = o.Discussionid
	toSerialize["hasparent"] = o.Hasparent
	toSerialize["haswordcount"] = o.Haswordcount
	if !IsNil(o.Html) {
		toSerialize["html"] = o.Html
	}
	toSerialize["id"] = o.Id
	toSerialize["isdeleted"] = o.Isdeleted
	toSerialize["isprivatereply"] = o.Isprivatereply
	toSerialize["message"] = o.Message
	toSerialize["messageformat"] = o.Messageformat
	if !IsNil(o.Messageinlinefiles) {
		toSerialize["messageinlinefiles"] = o.Messageinlinefiles
	}
	if !IsNil(o.Parentid) {
		toSerialize["parentid"] = o.Parentid
	}
	toSerialize["replysubject"] = o.Replysubject
	toSerialize["subject"] = o.Subject
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["timecreated"] = o.Timecreated
	toSerialize["timemodified"] = o.Timemodified
	if !IsNil(o.Unread) {
		toSerialize["unread"] = o.Unread
	}
	if !IsNil(o.Urls) {
		toSerialize["urls"] = o.Urls
	}
	if !IsNil(o.Wordcount) {
		toSerialize["wordcount"] = o.Wordcount
	}
	return toSerialize, nil
}

func (o *ModForumAddDiscussionPost200ResponsePost) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"attachments",
		"author",
		"capabilities",
		"discussionid",
		"hasparent",
		"haswordcount",
		"id",
		"isdeleted",
		"isprivatereply",
		"message",
		"messageformat",
		"replysubject",
		"subject",
		"timecreated",
		"timemodified",
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

	varModForumAddDiscussionPost200ResponsePost := _ModForumAddDiscussionPost200ResponsePost{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModForumAddDiscussionPost200ResponsePost)

	if err != nil {
		return err
	}

	*o = ModForumAddDiscussionPost200ResponsePost(varModForumAddDiscussionPost200ResponsePost)

	return err
}

type NullableModForumAddDiscussionPost200ResponsePost struct {
	value *ModForumAddDiscussionPost200ResponsePost
	isSet bool
}

func (v NullableModForumAddDiscussionPost200ResponsePost) Get() *ModForumAddDiscussionPost200ResponsePost {
	return v.value
}

func (v *NullableModForumAddDiscussionPost200ResponsePost) Set(val *ModForumAddDiscussionPost200ResponsePost) {
	v.value = val
	v.isSet = true
}

func (v NullableModForumAddDiscussionPost200ResponsePost) IsSet() bool {
	return v.isSet
}

func (v *NullableModForumAddDiscussionPost200ResponsePost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModForumAddDiscussionPost200ResponsePost(val *ModForumAddDiscussionPost200ResponsePost) *NullableModForumAddDiscussionPost200ResponsePost {
	return &NullableModForumAddDiscussionPost200ResponsePost{value: val, isSet: true}
}

func (v NullableModForumAddDiscussionPost200ResponsePost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModForumAddDiscussionPost200ResponsePost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


