package unisender

import (
	"net/url"
)

type List struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

// ListOptions
type ListOptions struct {
	BeforeSubscribeUrl string `json:"before_subscribe_url"`
	AfterSubscribeUrl  string `json:"after_subscribe_url"`
}

// GetLists returns all available campaign lists.
// see https://www.unisender.com/en/support/api/partners/getlists/
func (u *Unisender) GetLists() (lists []List, err error) {
	err = u.request("getLists", url.Values{}, &lists)
	return
}

// CreateList creates a new contact list.
// see https://www.unisender.com/en/support/api/partners/createlist/
func (u *Unisender) CreateList(title string, options ...ListOptions) (id int64, err error) {
	return
}

// UpdateList changes campaign list properties.
// see https://www.unisender.com/en/support/api/partners/updatelist/
func (u *Unisender) UpdateList(id int64, title string, options ...ListOptions) (err error) {
	return
}

// UpdateList removes a list.
// see https://www.unisender.com/en/support/api/partners/deletelist/
func (u *Unisender) DeleteList(id int64) (err error) {
	return
}

//func (u *Unisender) CreateList(title string, options ...ListOptions) {
//
//}
