package unisender

import "fmt"

type List struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

// GetLists returns all available campaign lists.
// see https://www.unisender.com/en/support/api/partners/getlists/
func (u *Unisender) GetLists() (lists []List, err error) {
	err = u.request("getLists", &lists)
	return
}

type createListResponse struct {
	ID int64 `json:"id"`
}

// CreateList creates a new contact list.
// see https://www.unisender.com/en/support/api/partners/createlist/
func (u *Unisender) CreateList(title string, options ...Option) (id int64, err error) {
	options = append(options, Option{
		name:  "title",
		value: title,
	})

	var res createListResponse
	if err = u.request("createList", &res, options...); err != nil {
		return
	}

	id = res.ID

	return
}

// UpdateList changes campaign list properties.
// see https://www.unisender.com/en/support/api/partners/updatelist/
func (u *Unisender) UpdateList(id int64, title string, options ...Option) (err error) {
	options = append(
		options,
		Option{
			name:  "id",
			value: fmt.Sprintf("%d", id),
		},
		Option{
			name:  "title",
			value: title,
		},
	)

	err = u.request("updateList", nil, options...)

	return
}

// UpdateList removes a list.
// see https://www.unisender.com/en/support/api/partners/deletelist/
func (u *Unisender) DeleteList(id int64) (err error) {
	option := Option{
		name:  "id",
		value: fmt.Sprintf("%d", id),
	}

	err = u.request("deleteList", nil, option)

	return
}
