package contacts

import (
	"fmt"
	"strings"
	"time"
)

const (
	defaultStatus       = "new"
	defaultAvailability = "available"
)

type contactValue struct {
	value               string
	status              string
	availability        string
	addTime             time.Time
	requestIP           string
	confirmTime         time.Time
	listIDs             []int64
	subscribeTimes      []time.Time
	unsubscribedListIDs []int64
	excludedListIDs     []int64
}

func (v *contactValue) toMap() map[string]string {
	return map[string]string{
		"status":                v.status,
		"availability":          v.availability,
		"add_time":              v.addTime.String(),
		"request_ip":            v.requestIP,
		"confirm_time":          v.confirmTime.String(),
		"list_ids":              v.int64SliceToString(v.listIDs...),
		"subscribe_times":       v.timeSliceToString(v.subscribeTimes...),
		"unsubscribed_list_ids": v.int64SliceToString(v.unsubscribedListIDs...),
		"excluded_list_ids":     v.int64SliceToString(v.excludedListIDs...),
	}
}

func (v *contactValue) int64SliceToString(slice ...int64) string {
	strSlice := make([]string, len(slice))
	for n, v := range slice {
		strSlice[n] = fmt.Sprintf("%d", v)
	}

	return strings.Join(strSlice, ",")
}

func (v *contactValue) timeSliceToString(slice ...time.Time) string {
	strSlice := make([]string, len(slice))
	for n, v := range slice {
		strSlice[n] = v.String()
	}

	return strings.Join(strSlice, ",")
}

type Contact struct {
	kind   string
	value  *contactValue
	tags   []string
	delete bool
}

func (c *Contact) Delete() {
	c.delete = true
}

func (c *Contact) SetTags(tags ...string) {
	c.tags = tags
}

func (c *Contact) SetStatusNew() {
	c.value.status = "new"
}

func (c *Contact) SetStatusActive() {
	c.value.status = "active"
}

func (c *Contact) SetStatusInactive() {
	c.value.status = "inactive"
}

func (c *Contact) SetStatusUnsubscribed() {
	c.value.status = "ubsubscribed"
}

func (c *Contact) SetAvailabilityAvailable() {
	c.value.availability = "available"
}

func (c *Contact) SetAvailabilityUnreachable() {
	c.value.availability = "unreachable"
}

func (c *Contact) SetAvailabilityTempUnreachable() {
	c.value.availability = "temp_unreachable"
}

func (c *Contact) SetAvailabilityMailboxFull() {
	c.value.availability = "mailbox_full"
}

func (c *Contact) SetAvailabilitySpamRejected() {
	c.value.availability = "spam_rejected"
}

func (c *Contact) SetAvailabilitySpamFolder() {
	c.value.availability = "spam_folder"
}

func (c *Contact) SetEmailAddTime(addTime time.Time) {
	c.value.addTime = addTime
}

func (c *Contact) SetEmailRequestIP(requestIP string) {
	c.value.requestIP = requestIP
}

func (c *Contact) SetConfirmTime(confirmTime time.Time) {
	c.value.confirmTime = confirmTime
}

func (c *Contact) AddListID(listID int64, subscribeTime time.Time) {
	v := c.value

	v.listIDs = append(v.listIDs, listID)
	v.subscribeTimes = append(v.subscribeTimes, subscribeTime)
}

func (c *Contact) SetUnsubscribedListIDs(unsubscribedListIDs ...int64) {
	c.value.unsubscribedListIDs = unsubscribedListIDs
}

func (c *Contact) SetEmailExcludedListIDs(excludedListIDs ...int64) {
	c.value.excludedListIDs = excludedListIDs
}

func (c *Contact) toMap() map[string]string {
	var del string
	if c.delete {
		del = "1"
	} else {
		del = "0"
	}

	m := map[string]string{
		"delete": del,
		"tags":   strings.Join(c.tags, ","),
	}

	m[c.kind] = c.value.value
	for k, v := range c.value.toMap() {
		m[c.kind+"_"+k] = v
	}

	return m
}

type Collection struct {
	contacts   []*Contact
	fieldNames []string
}

func (c *Collection) Email(email string) *Contact {
	contact := &Contact{
		kind:  "email",
		value: c.newContactValue(email),
	}

	c.contacts = append(c.contacts, contact)

	return contact
}

func (c *Collection) Phone(phone string) *Contact {
	contact := &Contact{
		kind:  "phone",
		value: c.newContactValue(phone),
	}

	c.contacts = append(c.contacts, contact)

	return contact
}

func (c *Collection) newContactValue(value string) *contactValue {
	return &contactValue{
		value:        value,
		status:       defaultStatus,
		availability: defaultAvailability,
	}
}

func (c *Collection) SetFieldNames(fieldNames ...string) {
	c.fieldNames = fieldNames
}

func NewCollection() *Collection {
	return &Collection{}
}
