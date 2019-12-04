package contacts

import "time"

type PersonList struct {
	ID      int64     `json:"id"`
	Title   string    `json:"title"`
	AddedAt time.Time `json:"added_at"`
}

type Person struct {
	Email        string            `json:"email"`
	AddedAt      time.Time         `json:"added_at"`
	Status       string            `json:"status"`
	Availability string            `json:"availability"`
	LastSend     time.Time         `json:"last_send_datetime,omitempty"`
	LastDelivery time.Time         `json:"last_delivery_datetime,omitempty"`
	LastRead     time.Time         `json:"last_read_datetime,omitempty"`
	LastClick    time.Time         `json:"last_click_datetime,omitempty"`
	Rating       float64           `json:"rating,omitempty"`
	Lists        []Person          `json:"lists,omitempty"`
	Fields       map[string]string `json:"fields,omitempty"`
}
