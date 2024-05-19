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

// checks if the CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner{}

// CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner struct for CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner
type CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner struct {
	// The id of the conversation
	Conversationid *int32 `json:"conversationid,omitempty"`
	// The user's name
	Fullname *string `json:"fullname,omitempty"`
	// If the user has been blocked
	Isblocked *bool `json:"isblocked,omitempty"`
	// If we are messaging the user
	Ismessaging *bool `json:"ismessaging,omitempty"`
	// The user's online status
	Isonline *bool `json:"isonline,omitempty"`
	// If the user has read the message
	Isread *bool `json:"isread,omitempty"`
	// The user's last message
	Lastmessage *string `json:"lastmessage,omitempty"`
	// Timestamp for last message
	Lastmessagedate *int32 `json:"lastmessagedate,omitempty"`
	// The unique search message id
	Messageid *int32 `json:"messageid,omitempty"`
	// User picture URL
	Profileimageurl *string `json:"profileimageurl,omitempty"`
	// Small user picture URL
	Profileimageurlsmall *string `json:"profileimageurlsmall,omitempty"`
	// Was the last message sent from the current user?
	Sentfromcurrentuser *bool `json:"sentfromcurrentuser,omitempty"`
	// Show the user's online status?
	Showonlinestatus *bool `json:"showonlinestatus,omitempty"`
	// The number of unread messages in this conversation
	Unreadcount *int32 `json:"unreadcount,omitempty"`
	// The user's id
	Userid *int32 `json:"userid,omitempty"`
}

// NewCoreMessageDataForMessageareaSearchMessages200ResponseContactsInner instantiates a new CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreMessageDataForMessageareaSearchMessages200ResponseContactsInner() *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner {
	this := CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner{}
	var conversationid int32 = null
	this.Conversationid = &conversationid
	var fullname string = "null"
	this.Fullname = &fullname
	var isblocked bool = null
	this.Isblocked = &isblocked
	var ismessaging bool = null
	this.Ismessaging = &ismessaging
	var isonline bool = null
	this.Isonline = &isonline
	var isread bool = null
	this.Isread = &isread
	var lastmessage string = "null"
	this.Lastmessage = &lastmessage
	var lastmessagedate int32 = null
	this.Lastmessagedate = &lastmessagedate
	var messageid int32 = null
	this.Messageid = &messageid
	var profileimageurl string = "null"
	this.Profileimageurl = &profileimageurl
	var profileimageurlsmall string = "null"
	this.Profileimageurlsmall = &profileimageurlsmall
	var sentfromcurrentuser bool = null
	this.Sentfromcurrentuser = &sentfromcurrentuser
	var showonlinestatus bool = null
	this.Showonlinestatus = &showonlinestatus
	var unreadcount int32 = null
	this.Unreadcount = &unreadcount
	var userid int32 = null
	this.Userid = &userid
	return &this
}

// NewCoreMessageDataForMessageareaSearchMessages200ResponseContactsInnerWithDefaults instantiates a new CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreMessageDataForMessageareaSearchMessages200ResponseContactsInnerWithDefaults() *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner {
	this := CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner{}
	var conversationid int32 = null
	this.Conversationid = &conversationid
	var fullname string = "null"
	this.Fullname = &fullname
	var isblocked bool = null
	this.Isblocked = &isblocked
	var ismessaging bool = null
	this.Ismessaging = &ismessaging
	var isonline bool = null
	this.Isonline = &isonline
	var isread bool = null
	this.Isread = &isread
	var lastmessage string = "null"
	this.Lastmessage = &lastmessage
	var lastmessagedate int32 = null
	this.Lastmessagedate = &lastmessagedate
	var messageid int32 = null
	this.Messageid = &messageid
	var profileimageurl string = "null"
	this.Profileimageurl = &profileimageurl
	var profileimageurlsmall string = "null"
	this.Profileimageurlsmall = &profileimageurlsmall
	var sentfromcurrentuser bool = null
	this.Sentfromcurrentuser = &sentfromcurrentuser
	var showonlinestatus bool = null
	this.Showonlinestatus = &showonlinestatus
	var unreadcount int32 = null
	this.Unreadcount = &unreadcount
	var userid int32 = null
	this.Userid = &userid
	return &this
}

// GetConversationid returns the Conversationid field value if set, zero value otherwise.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) GetConversationid() int32 {
	if o == nil || IsNil(o.Conversationid) {
		var ret int32
		return ret
	}
	return *o.Conversationid
}

// GetConversationidOk returns a tuple with the Conversationid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) GetConversationidOk() (*int32, bool) {
	if o == nil || IsNil(o.Conversationid) {
		return nil, false
	}
	return o.Conversationid, true
}

// HasConversationid returns a boolean if a field has been set.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) HasConversationid() bool {
	if o != nil && !IsNil(o.Conversationid) {
		return true
	}

	return false
}

// SetConversationid gets a reference to the given int32 and assigns it to the Conversationid field.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) SetConversationid(v int32) {
	o.Conversationid = &v
}

// GetFullname returns the Fullname field value if set, zero value otherwise.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) GetFullname() string {
	if o == nil || IsNil(o.Fullname) {
		var ret string
		return ret
	}
	return *o.Fullname
}

// GetFullnameOk returns a tuple with the Fullname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) GetFullnameOk() (*string, bool) {
	if o == nil || IsNil(o.Fullname) {
		return nil, false
	}
	return o.Fullname, true
}

// HasFullname returns a boolean if a field has been set.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) HasFullname() bool {
	if o != nil && !IsNil(o.Fullname) {
		return true
	}

	return false
}

// SetFullname gets a reference to the given string and assigns it to the Fullname field.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) SetFullname(v string) {
	o.Fullname = &v
}

// GetIsblocked returns the Isblocked field value if set, zero value otherwise.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) GetIsblocked() bool {
	if o == nil || IsNil(o.Isblocked) {
		var ret bool
		return ret
	}
	return *o.Isblocked
}

// GetIsblockedOk returns a tuple with the Isblocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) GetIsblockedOk() (*bool, bool) {
	if o == nil || IsNil(o.Isblocked) {
		return nil, false
	}
	return o.Isblocked, true
}

// HasIsblocked returns a boolean if a field has been set.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) HasIsblocked() bool {
	if o != nil && !IsNil(o.Isblocked) {
		return true
	}

	return false
}

// SetIsblocked gets a reference to the given bool and assigns it to the Isblocked field.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) SetIsblocked(v bool) {
	o.Isblocked = &v
}

// GetIsmessaging returns the Ismessaging field value if set, zero value otherwise.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) GetIsmessaging() bool {
	if o == nil || IsNil(o.Ismessaging) {
		var ret bool
		return ret
	}
	return *o.Ismessaging
}

// GetIsmessagingOk returns a tuple with the Ismessaging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) GetIsmessagingOk() (*bool, bool) {
	if o == nil || IsNil(o.Ismessaging) {
		return nil, false
	}
	return o.Ismessaging, true
}

// HasIsmessaging returns a boolean if a field has been set.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) HasIsmessaging() bool {
	if o != nil && !IsNil(o.Ismessaging) {
		return true
	}

	return false
}

// SetIsmessaging gets a reference to the given bool and assigns it to the Ismessaging field.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) SetIsmessaging(v bool) {
	o.Ismessaging = &v
}

// GetIsonline returns the Isonline field value if set, zero value otherwise.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) GetIsonline() bool {
	if o == nil || IsNil(o.Isonline) {
		var ret bool
		return ret
	}
	return *o.Isonline
}

// GetIsonlineOk returns a tuple with the Isonline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) GetIsonlineOk() (*bool, bool) {
	if o == nil || IsNil(o.Isonline) {
		return nil, false
	}
	return o.Isonline, true
}

// HasIsonline returns a boolean if a field has been set.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) HasIsonline() bool {
	if o != nil && !IsNil(o.Isonline) {
		return true
	}

	return false
}

// SetIsonline gets a reference to the given bool and assigns it to the Isonline field.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) SetIsonline(v bool) {
	o.Isonline = &v
}

// GetIsread returns the Isread field value if set, zero value otherwise.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) GetIsread() bool {
	if o == nil || IsNil(o.Isread) {
		var ret bool
		return ret
	}
	return *o.Isread
}

// GetIsreadOk returns a tuple with the Isread field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) GetIsreadOk() (*bool, bool) {
	if o == nil || IsNil(o.Isread) {
		return nil, false
	}
	return o.Isread, true
}

// HasIsread returns a boolean if a field has been set.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) HasIsread() bool {
	if o != nil && !IsNil(o.Isread) {
		return true
	}

	return false
}

// SetIsread gets a reference to the given bool and assigns it to the Isread field.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) SetIsread(v bool) {
	o.Isread = &v
}

// GetLastmessage returns the Lastmessage field value if set, zero value otherwise.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) GetLastmessage() string {
	if o == nil || IsNil(o.Lastmessage) {
		var ret string
		return ret
	}
	return *o.Lastmessage
}

// GetLastmessageOk returns a tuple with the Lastmessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) GetLastmessageOk() (*string, bool) {
	if o == nil || IsNil(o.Lastmessage) {
		return nil, false
	}
	return o.Lastmessage, true
}

// HasLastmessage returns a boolean if a field has been set.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) HasLastmessage() bool {
	if o != nil && !IsNil(o.Lastmessage) {
		return true
	}

	return false
}

// SetLastmessage gets a reference to the given string and assigns it to the Lastmessage field.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) SetLastmessage(v string) {
	o.Lastmessage = &v
}

// GetLastmessagedate returns the Lastmessagedate field value if set, zero value otherwise.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) GetLastmessagedate() int32 {
	if o == nil || IsNil(o.Lastmessagedate) {
		var ret int32
		return ret
	}
	return *o.Lastmessagedate
}

// GetLastmessagedateOk returns a tuple with the Lastmessagedate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) GetLastmessagedateOk() (*int32, bool) {
	if o == nil || IsNil(o.Lastmessagedate) {
		return nil, false
	}
	return o.Lastmessagedate, true
}

// HasLastmessagedate returns a boolean if a field has been set.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) HasLastmessagedate() bool {
	if o != nil && !IsNil(o.Lastmessagedate) {
		return true
	}

	return false
}

// SetLastmessagedate gets a reference to the given int32 and assigns it to the Lastmessagedate field.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) SetLastmessagedate(v int32) {
	o.Lastmessagedate = &v
}

// GetMessageid returns the Messageid field value if set, zero value otherwise.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) GetMessageid() int32 {
	if o == nil || IsNil(o.Messageid) {
		var ret int32
		return ret
	}
	return *o.Messageid
}

// GetMessageidOk returns a tuple with the Messageid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) GetMessageidOk() (*int32, bool) {
	if o == nil || IsNil(o.Messageid) {
		return nil, false
	}
	return o.Messageid, true
}

// HasMessageid returns a boolean if a field has been set.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) HasMessageid() bool {
	if o != nil && !IsNil(o.Messageid) {
		return true
	}

	return false
}

// SetMessageid gets a reference to the given int32 and assigns it to the Messageid field.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) SetMessageid(v int32) {
	o.Messageid = &v
}

// GetProfileimageurl returns the Profileimageurl field value if set, zero value otherwise.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) GetProfileimageurl() string {
	if o == nil || IsNil(o.Profileimageurl) {
		var ret string
		return ret
	}
	return *o.Profileimageurl
}

// GetProfileimageurlOk returns a tuple with the Profileimageurl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) GetProfileimageurlOk() (*string, bool) {
	if o == nil || IsNil(o.Profileimageurl) {
		return nil, false
	}
	return o.Profileimageurl, true
}

// HasProfileimageurl returns a boolean if a field has been set.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) HasProfileimageurl() bool {
	if o != nil && !IsNil(o.Profileimageurl) {
		return true
	}

	return false
}

// SetProfileimageurl gets a reference to the given string and assigns it to the Profileimageurl field.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) SetProfileimageurl(v string) {
	o.Profileimageurl = &v
}

// GetProfileimageurlsmall returns the Profileimageurlsmall field value if set, zero value otherwise.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) GetProfileimageurlsmall() string {
	if o == nil || IsNil(o.Profileimageurlsmall) {
		var ret string
		return ret
	}
	return *o.Profileimageurlsmall
}

// GetProfileimageurlsmallOk returns a tuple with the Profileimageurlsmall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) GetProfileimageurlsmallOk() (*string, bool) {
	if o == nil || IsNil(o.Profileimageurlsmall) {
		return nil, false
	}
	return o.Profileimageurlsmall, true
}

// HasProfileimageurlsmall returns a boolean if a field has been set.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) HasProfileimageurlsmall() bool {
	if o != nil && !IsNil(o.Profileimageurlsmall) {
		return true
	}

	return false
}

// SetProfileimageurlsmall gets a reference to the given string and assigns it to the Profileimageurlsmall field.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) SetProfileimageurlsmall(v string) {
	o.Profileimageurlsmall = &v
}

// GetSentfromcurrentuser returns the Sentfromcurrentuser field value if set, zero value otherwise.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) GetSentfromcurrentuser() bool {
	if o == nil || IsNil(o.Sentfromcurrentuser) {
		var ret bool
		return ret
	}
	return *o.Sentfromcurrentuser
}

// GetSentfromcurrentuserOk returns a tuple with the Sentfromcurrentuser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) GetSentfromcurrentuserOk() (*bool, bool) {
	if o == nil || IsNil(o.Sentfromcurrentuser) {
		return nil, false
	}
	return o.Sentfromcurrentuser, true
}

// HasSentfromcurrentuser returns a boolean if a field has been set.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) HasSentfromcurrentuser() bool {
	if o != nil && !IsNil(o.Sentfromcurrentuser) {
		return true
	}

	return false
}

// SetSentfromcurrentuser gets a reference to the given bool and assigns it to the Sentfromcurrentuser field.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) SetSentfromcurrentuser(v bool) {
	o.Sentfromcurrentuser = &v
}

// GetShowonlinestatus returns the Showonlinestatus field value if set, zero value otherwise.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) GetShowonlinestatus() bool {
	if o == nil || IsNil(o.Showonlinestatus) {
		var ret bool
		return ret
	}
	return *o.Showonlinestatus
}

// GetShowonlinestatusOk returns a tuple with the Showonlinestatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) GetShowonlinestatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Showonlinestatus) {
		return nil, false
	}
	return o.Showonlinestatus, true
}

// HasShowonlinestatus returns a boolean if a field has been set.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) HasShowonlinestatus() bool {
	if o != nil && !IsNil(o.Showonlinestatus) {
		return true
	}

	return false
}

// SetShowonlinestatus gets a reference to the given bool and assigns it to the Showonlinestatus field.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) SetShowonlinestatus(v bool) {
	o.Showonlinestatus = &v
}

// GetUnreadcount returns the Unreadcount field value if set, zero value otherwise.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) GetUnreadcount() int32 {
	if o == nil || IsNil(o.Unreadcount) {
		var ret int32
		return ret
	}
	return *o.Unreadcount
}

// GetUnreadcountOk returns a tuple with the Unreadcount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) GetUnreadcountOk() (*int32, bool) {
	if o == nil || IsNil(o.Unreadcount) {
		return nil, false
	}
	return o.Unreadcount, true
}

// HasUnreadcount returns a boolean if a field has been set.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) HasUnreadcount() bool {
	if o != nil && !IsNil(o.Unreadcount) {
		return true
	}

	return false
}

// SetUnreadcount gets a reference to the given int32 and assigns it to the Unreadcount field.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) SetUnreadcount(v int32) {
	o.Unreadcount = &v
}

// GetUserid returns the Userid field value if set, zero value otherwise.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) GetUserid() int32 {
	if o == nil || IsNil(o.Userid) {
		var ret int32
		return ret
	}
	return *o.Userid
}

// GetUseridOk returns a tuple with the Userid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) GetUseridOk() (*int32, bool) {
	if o == nil || IsNil(o.Userid) {
		return nil, false
	}
	return o.Userid, true
}

// HasUserid returns a boolean if a field has been set.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) HasUserid() bool {
	if o != nil && !IsNil(o.Userid) {
		return true
	}

	return false
}

// SetUserid gets a reference to the given int32 and assigns it to the Userid field.
func (o *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) SetUserid(v int32) {
	o.Userid = &v
}

func (o CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Conversationid) {
		toSerialize["conversationid"] = o.Conversationid
	}
	if !IsNil(o.Fullname) {
		toSerialize["fullname"] = o.Fullname
	}
	if !IsNil(o.Isblocked) {
		toSerialize["isblocked"] = o.Isblocked
	}
	if !IsNil(o.Ismessaging) {
		toSerialize["ismessaging"] = o.Ismessaging
	}
	if !IsNil(o.Isonline) {
		toSerialize["isonline"] = o.Isonline
	}
	if !IsNil(o.Isread) {
		toSerialize["isread"] = o.Isread
	}
	if !IsNil(o.Lastmessage) {
		toSerialize["lastmessage"] = o.Lastmessage
	}
	if !IsNil(o.Lastmessagedate) {
		toSerialize["lastmessagedate"] = o.Lastmessagedate
	}
	if !IsNil(o.Messageid) {
		toSerialize["messageid"] = o.Messageid
	}
	if !IsNil(o.Profileimageurl) {
		toSerialize["profileimageurl"] = o.Profileimageurl
	}
	if !IsNil(o.Profileimageurlsmall) {
		toSerialize["profileimageurlsmall"] = o.Profileimageurlsmall
	}
	if !IsNil(o.Sentfromcurrentuser) {
		toSerialize["sentfromcurrentuser"] = o.Sentfromcurrentuser
	}
	if !IsNil(o.Showonlinestatus) {
		toSerialize["showonlinestatus"] = o.Showonlinestatus
	}
	if !IsNil(o.Unreadcount) {
		toSerialize["unreadcount"] = o.Unreadcount
	}
	if !IsNil(o.Userid) {
		toSerialize["userid"] = o.Userid
	}
	return toSerialize, nil
}

type NullableCoreMessageDataForMessageareaSearchMessages200ResponseContactsInner struct {
	value *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner
	isSet bool
}

func (v NullableCoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) Get() *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner {
	return v.value
}

func (v *NullableCoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) Set(val *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreMessageDataForMessageareaSearchMessages200ResponseContactsInner(val *CoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) *NullableCoreMessageDataForMessageareaSearchMessages200ResponseContactsInner {
	return &NullableCoreMessageDataForMessageareaSearchMessages200ResponseContactsInner{value: val, isSet: true}
}

func (v NullableCoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreMessageDataForMessageareaSearchMessages200ResponseContactsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

