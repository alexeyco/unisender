package contacts

import "time"

// List mailing list info.
type List struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

// PersonList list info, contact subscribed.
type PersonList struct {
	ID      int64     `json:"id"`
	Title   string    `json:"title"`
	AddedAt time.Time `json:"added_at,omitempty"`
}

// PersonEmail contact email info.
type PersonEmail struct {
	Email        string            `json:"email"`
	AddedAt      time.Time         `json:"added_at"`
	Status       string            `json:"status"`
	Availability string            `json:"availability"`
	LastSend     time.Time         `json:"last_send_datetime,omitempty"`
	LastDelivery time.Time         `json:"last_delivery_datetime,omitempty"`
	LastRead     time.Time         `json:"last_read_datetime,omitempty"`
	LastClick    time.Time         `json:"last_click_datetime,omitempty"`
	Rating       float64           `json:"rating,omitempty"`
	Lists        []PersonList      `json:"lists,omitempty"`
	Fields       map[string]string `json:"fields,omitempty"`
}

// Person contact struct.
type Person struct {
	Email PersonEmail `json:"email"`
}
