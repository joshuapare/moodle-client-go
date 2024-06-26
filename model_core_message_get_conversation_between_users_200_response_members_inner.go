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

// checks if the CoreMessageGetConversationBetweenUsers200ResponseMembersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreMessageGetConversationBetweenUsers200ResponseMembersInner{}

// CoreMessageGetConversationBetweenUsers200ResponseMembersInner struct for CoreMessageGetConversationBetweenUsers200ResponseMembersInner
type CoreMessageGetConversationBetweenUsers200ResponseMembersInner struct {
	// If the user can be messaged
	Canmessage *bool `json:"canmessage,omitempty"`
	// If the user can still message even if they get blocked
	Canmessageevenifblocked *bool `json:"canmessageevenifblocked,omitempty"`
	Contactrequests []CoreMessageGetConversationBetweenUsers200ResponseMembersInnerContactrequestsInner `json:"contactrequests,omitempty"`
	Conversations []CoreMessageGetConversationBetweenUsers200ResponseMembersInnerConversationsInner `json:"conversations,omitempty"`
	// The user's name
	Fullname *string `json:"fullname,omitempty"`
	// The user id
	Id *int32 `json:"id,omitempty"`
	// If the user has been blocked
	Isblocked *bool `json:"isblocked,omitempty"`
	// Is the user a contact?
	Iscontact *bool `json:"iscontact,omitempty"`
	// Is the user deleted?
	Isdeleted *bool `json:"isdeleted,omitempty"`
	// The user's online status
	Isonline *bool `json:"isonline,omitempty"`
	// User picture URL
	Profileimageurl *string `json:"profileimageurl,omitempty"`
	// Small user picture URL
	Profileimageurlsmall *string `json:"profileimageurlsmall,omitempty"`
	// The link to the user's profile page
	Profileurl *string `json:"profileurl,omitempty"`
	// If the user requires to be contacts
	Requirescontact *bool `json:"requirescontact,omitempty"`
	// Show the user's online status?
	Showonlinestatus *bool `json:"showonlinestatus,omitempty"`
}

// NewCoreMessageGetConversationBetweenUsers200ResponseMembersInner instantiates a new CoreMessageGetConversationBetweenUsers200ResponseMembersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreMessageGetConversationBetweenUsers200ResponseMembersInner() *CoreMessageGetConversationBetweenUsers200ResponseMembersInner {
	this := CoreMessageGetConversationBetweenUsers200ResponseMembersInner{}
	return &this
}

// NewCoreMessageGetConversationBetweenUsers200ResponseMembersInnerWithDefaults instantiates a new CoreMessageGetConversationBetweenUsers200ResponseMembersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreMessageGetConversationBetweenUsers200ResponseMembersInnerWithDefaults() *CoreMessageGetConversationBetweenUsers200ResponseMembersInner {
	this := CoreMessageGetConversationBetweenUsers200ResponseMembersInner{}
	return &this
}

// GetCanmessage returns the Canmessage field value if set, zero value otherwise.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) GetCanmessage() bool {
	if o == nil || IsNil(o.Canmessage) {
		var ret bool
		return ret
	}
	return *o.Canmessage
}

// GetCanmessageOk returns a tuple with the Canmessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) GetCanmessageOk() (*bool, bool) {
	if o == nil || IsNil(o.Canmessage) {
		return nil, false
	}
	return o.Canmessage, true
}

// HasCanmessage returns a boolean if a field has been set.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) HasCanmessage() bool {
	if o != nil && !IsNil(o.Canmessage) {
		return true
	}

	return false
}

// SetCanmessage gets a reference to the given bool and assigns it to the Canmessage field.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) SetCanmessage(v bool) {
	o.Canmessage = &v
}

// GetCanmessageevenifblocked returns the Canmessageevenifblocked field value if set, zero value otherwise.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) GetCanmessageevenifblocked() bool {
	if o == nil || IsNil(o.Canmessageevenifblocked) {
		var ret bool
		return ret
	}
	return *o.Canmessageevenifblocked
}

// GetCanmessageevenifblockedOk returns a tuple with the Canmessageevenifblocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) GetCanmessageevenifblockedOk() (*bool, bool) {
	if o == nil || IsNil(o.Canmessageevenifblocked) {
		return nil, false
	}
	return o.Canmessageevenifblocked, true
}

// HasCanmessageevenifblocked returns a boolean if a field has been set.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) HasCanmessageevenifblocked() bool {
	if o != nil && !IsNil(o.Canmessageevenifblocked) {
		return true
	}

	return false
}

// SetCanmessageevenifblocked gets a reference to the given bool and assigns it to the Canmessageevenifblocked field.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) SetCanmessageevenifblocked(v bool) {
	o.Canmessageevenifblocked = &v
}

// GetContactrequests returns the Contactrequests field value if set, zero value otherwise.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) GetContactrequests() []CoreMessageGetConversationBetweenUsers200ResponseMembersInnerContactrequestsInner {
	if o == nil || IsNil(o.Contactrequests) {
		var ret []CoreMessageGetConversationBetweenUsers200ResponseMembersInnerContactrequestsInner
		return ret
	}
	return o.Contactrequests
}

// GetContactrequestsOk returns a tuple with the Contactrequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) GetContactrequestsOk() ([]CoreMessageGetConversationBetweenUsers200ResponseMembersInnerContactrequestsInner, bool) {
	if o == nil || IsNil(o.Contactrequests) {
		return nil, false
	}
	return o.Contactrequests, true
}

// HasContactrequests returns a boolean if a field has been set.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) HasContactrequests() bool {
	if o != nil && !IsNil(o.Contactrequests) {
		return true
	}

	return false
}

// SetContactrequests gets a reference to the given []CoreMessageGetConversationBetweenUsers200ResponseMembersInnerContactrequestsInner and assigns it to the Contactrequests field.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) SetContactrequests(v []CoreMessageGetConversationBetweenUsers200ResponseMembersInnerContactrequestsInner) {
	o.Contactrequests = v
}

// GetConversations returns the Conversations field value if set, zero value otherwise.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) GetConversations() []CoreMessageGetConversationBetweenUsers200ResponseMembersInnerConversationsInner {
	if o == nil || IsNil(o.Conversations) {
		var ret []CoreMessageGetConversationBetweenUsers200ResponseMembersInnerConversationsInner
		return ret
	}
	return o.Conversations
}

// GetConversationsOk returns a tuple with the Conversations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) GetConversationsOk() ([]CoreMessageGetConversationBetweenUsers200ResponseMembersInnerConversationsInner, bool) {
	if o == nil || IsNil(o.Conversations) {
		return nil, false
	}
	return o.Conversations, true
}

// HasConversations returns a boolean if a field has been set.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) HasConversations() bool {
	if o != nil && !IsNil(o.Conversations) {
		return true
	}

	return false
}

// SetConversations gets a reference to the given []CoreMessageGetConversationBetweenUsers200ResponseMembersInnerConversationsInner and assigns it to the Conversations field.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) SetConversations(v []CoreMessageGetConversationBetweenUsers200ResponseMembersInnerConversationsInner) {
	o.Conversations = v
}

// GetFullname returns the Fullname field value if set, zero value otherwise.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) GetFullname() string {
	if o == nil || IsNil(o.Fullname) {
		var ret string
		return ret
	}
	return *o.Fullname
}

// GetFullnameOk returns a tuple with the Fullname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) GetFullnameOk() (*string, bool) {
	if o == nil || IsNil(o.Fullname) {
		return nil, false
	}
	return o.Fullname, true
}

// HasFullname returns a boolean if a field has been set.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) HasFullname() bool {
	if o != nil && !IsNil(o.Fullname) {
		return true
	}

	return false
}

// SetFullname gets a reference to the given string and assigns it to the Fullname field.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) SetFullname(v string) {
	o.Fullname = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) SetId(v int32) {
	o.Id = &v
}

// GetIsblocked returns the Isblocked field value if set, zero value otherwise.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) GetIsblocked() bool {
	if o == nil || IsNil(o.Isblocked) {
		var ret bool
		return ret
	}
	return *o.Isblocked
}

// GetIsblockedOk returns a tuple with the Isblocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) GetIsblockedOk() (*bool, bool) {
	if o == nil || IsNil(o.Isblocked) {
		return nil, false
	}
	return o.Isblocked, true
}

// HasIsblocked returns a boolean if a field has been set.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) HasIsblocked() bool {
	if o != nil && !IsNil(o.Isblocked) {
		return true
	}

	return false
}

// SetIsblocked gets a reference to the given bool and assigns it to the Isblocked field.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) SetIsblocked(v bool) {
	o.Isblocked = &v
}

// GetIscontact returns the Iscontact field value if set, zero value otherwise.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) GetIscontact() bool {
	if o == nil || IsNil(o.Iscontact) {
		var ret bool
		return ret
	}
	return *o.Iscontact
}

// GetIscontactOk returns a tuple with the Iscontact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) GetIscontactOk() (*bool, bool) {
	if o == nil || IsNil(o.Iscontact) {
		return nil, false
	}
	return o.Iscontact, true
}

// HasIscontact returns a boolean if a field has been set.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) HasIscontact() bool {
	if o != nil && !IsNil(o.Iscontact) {
		return true
	}

	return false
}

// SetIscontact gets a reference to the given bool and assigns it to the Iscontact field.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) SetIscontact(v bool) {
	o.Iscontact = &v
}

// GetIsdeleted returns the Isdeleted field value if set, zero value otherwise.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) GetIsdeleted() bool {
	if o == nil || IsNil(o.Isdeleted) {
		var ret bool
		return ret
	}
	return *o.Isdeleted
}

// GetIsdeletedOk returns a tuple with the Isdeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) GetIsdeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.Isdeleted) {
		return nil, false
	}
	return o.Isdeleted, true
}

// HasIsdeleted returns a boolean if a field has been set.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) HasIsdeleted() bool {
	if o != nil && !IsNil(o.Isdeleted) {
		return true
	}

	return false
}

// SetIsdeleted gets a reference to the given bool and assigns it to the Isdeleted field.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) SetIsdeleted(v bool) {
	o.Isdeleted = &v
}

// GetIsonline returns the Isonline field value if set, zero value otherwise.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) GetIsonline() bool {
	if o == nil || IsNil(o.Isonline) {
		var ret bool
		return ret
	}
	return *o.Isonline
}

// GetIsonlineOk returns a tuple with the Isonline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) GetIsonlineOk() (*bool, bool) {
	if o == nil || IsNil(o.Isonline) {
		return nil, false
	}
	return o.Isonline, true
}

// HasIsonline returns a boolean if a field has been set.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) HasIsonline() bool {
	if o != nil && !IsNil(o.Isonline) {
		return true
	}

	return false
}

// SetIsonline gets a reference to the given bool and assigns it to the Isonline field.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) SetIsonline(v bool) {
	o.Isonline = &v
}

// GetProfileimageurl returns the Profileimageurl field value if set, zero value otherwise.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) GetProfileimageurl() string {
	if o == nil || IsNil(o.Profileimageurl) {
		var ret string
		return ret
	}
	return *o.Profileimageurl
}

// GetProfileimageurlOk returns a tuple with the Profileimageurl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) GetProfileimageurlOk() (*string, bool) {
	if o == nil || IsNil(o.Profileimageurl) {
		return nil, false
	}
	return o.Profileimageurl, true
}

// HasProfileimageurl returns a boolean if a field has been set.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) HasProfileimageurl() bool {
	if o != nil && !IsNil(o.Profileimageurl) {
		return true
	}

	return false
}

// SetProfileimageurl gets a reference to the given string and assigns it to the Profileimageurl field.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) SetProfileimageurl(v string) {
	o.Profileimageurl = &v
}

// GetProfileimageurlsmall returns the Profileimageurlsmall field value if set, zero value otherwise.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) GetProfileimageurlsmall() string {
	if o == nil || IsNil(o.Profileimageurlsmall) {
		var ret string
		return ret
	}
	return *o.Profileimageurlsmall
}

// GetProfileimageurlsmallOk returns a tuple with the Profileimageurlsmall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) GetProfileimageurlsmallOk() (*string, bool) {
	if o == nil || IsNil(o.Profileimageurlsmall) {
		return nil, false
	}
	return o.Profileimageurlsmall, true
}

// HasProfileimageurlsmall returns a boolean if a field has been set.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) HasProfileimageurlsmall() bool {
	if o != nil && !IsNil(o.Profileimageurlsmall) {
		return true
	}

	return false
}

// SetProfileimageurlsmall gets a reference to the given string and assigns it to the Profileimageurlsmall field.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) SetProfileimageurlsmall(v string) {
	o.Profileimageurlsmall = &v
}

// GetProfileurl returns the Profileurl field value if set, zero value otherwise.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) GetProfileurl() string {
	if o == nil || IsNil(o.Profileurl) {
		var ret string
		return ret
	}
	return *o.Profileurl
}

// GetProfileurlOk returns a tuple with the Profileurl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) GetProfileurlOk() (*string, bool) {
	if o == nil || IsNil(o.Profileurl) {
		return nil, false
	}
	return o.Profileurl, true
}

// HasProfileurl returns a boolean if a field has been set.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) HasProfileurl() bool {
	if o != nil && !IsNil(o.Profileurl) {
		return true
	}

	return false
}

// SetProfileurl gets a reference to the given string and assigns it to the Profileurl field.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) SetProfileurl(v string) {
	o.Profileurl = &v
}

// GetRequirescontact returns the Requirescontact field value if set, zero value otherwise.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) GetRequirescontact() bool {
	if o == nil || IsNil(o.Requirescontact) {
		var ret bool
		return ret
	}
	return *o.Requirescontact
}

// GetRequirescontactOk returns a tuple with the Requirescontact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) GetRequirescontactOk() (*bool, bool) {
	if o == nil || IsNil(o.Requirescontact) {
		return nil, false
	}
	return o.Requirescontact, true
}

// HasRequirescontact returns a boolean if a field has been set.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) HasRequirescontact() bool {
	if o != nil && !IsNil(o.Requirescontact) {
		return true
	}

	return false
}

// SetRequirescontact gets a reference to the given bool and assigns it to the Requirescontact field.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) SetRequirescontact(v bool) {
	o.Requirescontact = &v
}

// GetShowonlinestatus returns the Showonlinestatus field value if set, zero value otherwise.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) GetShowonlinestatus() bool {
	if o == nil || IsNil(o.Showonlinestatus) {
		var ret bool
		return ret
	}
	return *o.Showonlinestatus
}

// GetShowonlinestatusOk returns a tuple with the Showonlinestatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) GetShowonlinestatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Showonlinestatus) {
		return nil, false
	}
	return o.Showonlinestatus, true
}

// HasShowonlinestatus returns a boolean if a field has been set.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) HasShowonlinestatus() bool {
	if o != nil && !IsNil(o.Showonlinestatus) {
		return true
	}

	return false
}

// SetShowonlinestatus gets a reference to the given bool and assigns it to the Showonlinestatus field.
func (o *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) SetShowonlinestatus(v bool) {
	o.Showonlinestatus = &v
}

func (o CoreMessageGetConversationBetweenUsers200ResponseMembersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreMessageGetConversationBetweenUsers200ResponseMembersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Canmessage) {
		toSerialize["canmessage"] = o.Canmessage
	}
	if !IsNil(o.Canmessageevenifblocked) {
		toSerialize["canmessageevenifblocked"] = o.Canmessageevenifblocked
	}
	if !IsNil(o.Contactrequests) {
		toSerialize["contactrequests"] = o.Contactrequests
	}
	if !IsNil(o.Conversations) {
		toSerialize["conversations"] = o.Conversations
	}
	if !IsNil(o.Fullname) {
		toSerialize["fullname"] = o.Fullname
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Isblocked) {
		toSerialize["isblocked"] = o.Isblocked
	}
	if !IsNil(o.Iscontact) {
		toSerialize["iscontact"] = o.Iscontact
	}
	if !IsNil(o.Isdeleted) {
		toSerialize["isdeleted"] = o.Isdeleted
	}
	if !IsNil(o.Isonline) {
		toSerialize["isonline"] = o.Isonline
	}
	if !IsNil(o.Profileimageurl) {
		toSerialize["profileimageurl"] = o.Profileimageurl
	}
	if !IsNil(o.Profileimageurlsmall) {
		toSerialize["profileimageurlsmall"] = o.Profileimageurlsmall
	}
	if !IsNil(o.Profileurl) {
		toSerialize["profileurl"] = o.Profileurl
	}
	if !IsNil(o.Requirescontact) {
		toSerialize["requirescontact"] = o.Requirescontact
	}
	if !IsNil(o.Showonlinestatus) {
		toSerialize["showonlinestatus"] = o.Showonlinestatus
	}
	return toSerialize, nil
}

type NullableCoreMessageGetConversationBetweenUsers200ResponseMembersInner struct {
	value *CoreMessageGetConversationBetweenUsers200ResponseMembersInner
	isSet bool
}

func (v NullableCoreMessageGetConversationBetweenUsers200ResponseMembersInner) Get() *CoreMessageGetConversationBetweenUsers200ResponseMembersInner {
	return v.value
}

func (v *NullableCoreMessageGetConversationBetweenUsers200ResponseMembersInner) Set(val *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreMessageGetConversationBetweenUsers200ResponseMembersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreMessageGetConversationBetweenUsers200ResponseMembersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreMessageGetConversationBetweenUsers200ResponseMembersInner(val *CoreMessageGetConversationBetweenUsers200ResponseMembersInner) *NullableCoreMessageGetConversationBetweenUsers200ResponseMembersInner {
	return &NullableCoreMessageGetConversationBetweenUsers200ResponseMembersInner{value: val, isSet: true}
}

func (v NullableCoreMessageGetConversationBetweenUsers200ResponseMembersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreMessageGetConversationBetweenUsers200ResponseMembersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


