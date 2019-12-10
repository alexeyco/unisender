package contacts2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/alexeyco/unisender/api"
)

// SubscribeResponse response of subscribe request.
type SubscribeResponse struct {
	PersonID int64 `json:"person_id"`
}

// SubscribeRequest request that adds contacts (email address and/or mobile phone) of a contact to one
// or several lists.
type SubscribeRequest struct {
	request *api.Request
}

// Field sets an additional contact field.
func (r *SubscribeRequest) Field(field, value string) *SubscribeRequest {
	r.request.Add(fmt.Sprintf("fields[%s]", field), value)
	return r
}

// Email sets contact email.
func (r *SubscribeRequest) Email(email string) *SubscribeRequest {
	return r.Field("email", email)
}

// Phone sets contact phone.
func (r *SubscribeRequest) Phone(phone string) *SubscribeRequest {
	return r.Field("phone", phone)
}

// Tags sets contact tags.
func (r *SubscribeRequest) Tags(tags ...string) *SubscribeRequest {
	r.request.Add("tags", strings.Join(tags, ","))
	return r
}

// DoubleOptinUnconfirmed if used, we believe that the contact only expressed a desire to subscribe,
// but has not yet confirmed the subscription. In this case, an invitation to subscribe will be sent to the contact.
// The text of the message will be taken from the properties of the first list from list_ids. By the way,
// the text can be changed using the updateOptInEmail method or through the web interface.
//
// See https://www.unisender.com/en/support/api/messages/updateoptinemail/
func (r *SubscribeRequest) DoubleOptinUnconfirmed() *SubscribeRequest {
	r.request.Add("double_optin", "0")
	return r
}

// DoubleOptinConfirmed if used, it is also considered that you already have the consent of the contact,
// the contact is added with the status “new”.
//
// See https://www.unisender.com/en/support/api/messages/updateoptinemail/
func (r *SubscribeRequest) DoubleOptinConfirmed() *SubscribeRequest {
	r.request.Add("double_optin", "3")
	return r
}

// DoubleOptinConfirmedIfActiveOrNew if used, the system checks for the presence of a contact in your lists.
// If the contact is already in your lists with the status “new” or “active”, then the address will simply be added
// to the list you specified. If the contact is not in your lists or its status is different from “new” or “active”,
// an invitation letter will be sent to him to subscribe. The text of this letter can be configured for each list
// using the updateOptInEmail method or through the web interface.
//
// See https://www.unisender.com/en/support/api/messages/updateoptinemail/
func (r *SubscribeRequest) DoubleOptinConfirmedIfActiveOrNew() *SubscribeRequest {
	r.request.Add("double_optin", "4")
	return r
}

// DoNotOverwrite if used, only the addition of new fields and labels takes place, existing fields retain their value.
func (r *SubscribeRequest) DoNotOverwrite() *SubscribeRequest {
	r.request.Add("overwrite", "0")
	return r
}

// OverwriteAll if used, all old fields are deleted and replaced with new ones, all old labels are also deleted
// and replaced with new ones. The contact will be deleted from all lists except those passed in the list_ids parameter.
func (r *SubscribeRequest) OverwriteAll() *SubscribeRequest {
	r.request.Add("overwrite", "1")
	return r
}

// OverwritePartially if used, the values ​​of the transferred fields are replaced, if the existing contact
// has other fields, then they retain their value. In the case of transferring labels, they are overwritten;
// if labels are not transmitted, the old values ​​of the labels are saved.
func (r *SubscribeRequest) OverwritePartially() *SubscribeRequest {
	r.request.Add("overwrite", "2")
	return r
}

// Execute sends request to UniSender API and returns result.
func (r *SubscribeRequest) Execute() (personID int64, err error) {
	var res SubscribeResponse
	if err = r.request.Execute("subscribe", &res); err != nil {
		return
	}

	personID = res.PersonID

	return
}

// Subscribe returns request that adds contacts (email address and/or mobile phone) of a contact to one
// or several lists, and also allows you to add/change the values ​​of additional fields and labels.
//
// See https://www.unisender.com/en/support/api/contacts/subscribe/
func Subscribe(request *api.Request, listIDs ...int64) *SubscribeRequest {
	ids := make([]string, len(listIDs))
	for n, id := range listIDs {
		ids[n] = strconv.FormatInt(id, 10)
	}

	request.Add("list_ids", strings.Join(ids, ","))

	return &SubscribeRequest{
		request: request,
	}
}
