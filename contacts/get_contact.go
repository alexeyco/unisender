package contacts

import (
	"github.com/alexeyco/unisender/api"
	"time"
)

// GetContactResultList list info, contact subscribed.
type GetContactResultList struct {
	ID      int64     `json:"id"`
	Title   string    `json:"title"`
	AddedAt time.Time `json:"added_at,omitempty"`
}

// GetContactResultEmail contact email info.
type GetContactResultEmail struct {
	Email        string                 `json:"email"`
	AddedAt      time.Time              `json:"added_at"`
	Status       string                 `json:"status"`
	Availability string                 `json:"availability"`
	LastSend     time.Time              `json:"last_send_datetime,omitempty"`
	LastDelivery time.Time              `json:"last_delivery_datetime,omitempty"`
	LastRead     time.Time              `json:"last_read_datetime,omitempty"`
	LastClick    time.Time              `json:"last_click_datetime,omitempty"`
	Rating       float64                `json:"rating,omitempty"`
	Lists        []GetContactResultList `json:"lists,omitempty"`
	Fields       map[string]string      `json:"fields,omitempty"`
}

// GetContactResult contact struct.
type GetContactResult struct {
	Email GetContactResultEmail `json:"email"`
}

// GetContactRequest request to getting information about a contact (one contact only).
type GetContactRequest struct {
	request *api.Request
}

// IncludeLists displays information about the lists to which the contact has been added.
func (r *GetContactRequest) IncludeLists() *GetContactRequest {
	r.request.Add("include_lists", "1")
	return r
}

// IncludeFields displays information about additional contact fields.
func (r *GetContactRequest) IncludeFields() *GetContactRequest {
	r.request.Add("include_fields", "1")
	return r
}

// IncludeDetails displays additional information about the contact.
func (r *GetContactRequest) IncludeDetails() *GetContactRequest {
	r.request.Add("include_details", "1")
	return r
}

// Execute sends request to UniSender API and returns result.
func (r *GetContactRequest) Execute() (person *GetContactResult, err error) {
	var p GetContactResult
	if err = r.request.Execute("getContact", &p); err != nil {
		return
	}

	person = &p

	return
}

// GetContact returns request to getting information about a contact (one contact only).
func GetContact(request *api.Request, email string) *GetContactRequest {
	request.Add("email", email)

	return &GetContactRequest{
		request: request,
	}
}
