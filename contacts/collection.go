package contacts

import "time"

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

type Collection struct {
	contacts   []*Contact
	fieldNames []string
}

func (c *Collection) Email(email string) *Contact {
	contact := &Contact{
		kind: "email",
		value: &contactValue{
			value: email,
		},
	}

	c.contacts = append(c.contacts, contact)

	return contact
}

func (c *Collection) Phone(phone string) *Contact {
	contact := &Contact{
		kind: "phone",
		value: &contactValue{
			value: phone,
		},
	}

	c.contacts = append(c.contacts, contact)

	return contact
}

func (c *Collection) SetFieldNames(fieldNames ...string) {
	c.fieldNames = fieldNames
}

func NewCollection() *Collection {
	return &Collection{}
}
