package contacts

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type ImportContactsContact struct {
	collection *ImportContactsCollection
	kind       string
	fields     map[string]string
	listIDs    map[int64]time.Time
}

func (c *ImportContactsContact) Delete() *ImportContactsContact {
	return c.setField("delete", "1")
}

func (c *ImportContactsContact) SetTags(tags ...string) *ImportContactsContact {
	return c.setField("tags", strings.Join(tags, ","))
}

func (c *ImportContactsContact) SetStatusNew() *ImportContactsContact {
	return c.setField("status", "new", true)
}

func (c *ImportContactsContact) SetStatusActive() *ImportContactsContact {
	return c.setField("status", "active", true)
}

func (c *ImportContactsContact) SetStatusInactive() *ImportContactsContact {
	return c.setField("status", "inactive", true)
}

func (c *ImportContactsContact) SetStatusUnsubscribed() *ImportContactsContact {
	return c.setField("status", "unsubscribed", true)
}

func (c *ImportContactsContact) SetAvailabilityAvailable() *ImportContactsContact {
	return c.setField("availability", "available", true)
}

func (c *ImportContactsContact) SetAvailabilityUnreachable() *ImportContactsContact {
	return c.setField("availability", "unreachable", true)
}

func (c *ImportContactsContact) SetAvailabilityTempUnreachable() *ImportContactsContact {
	return c.setField("availability", "temp_unreachable", true)
}

func (c *ImportContactsContact) SetAvailabilityMailboxFull() *ImportContactsContact {
	return c.setField("availability", "mailbox_full", true)
}

func (c *ImportContactsContact) SetAvailabilitySpamRejected() *ImportContactsContact {
	return c.setField("availability", "spam_rejected", true)
}

func (c *ImportContactsContact) SetAvailabilitySpamFolder() *ImportContactsContact {
	return c.setField("availability", "spam_folder", true)
}

func (c *ImportContactsContact) SetAddTime(addTime time.Time) *ImportContactsContact {
	return c.setField("add_time", addTime.Format(time.RFC3339), true)
}

func (c *ImportContactsContact) SetConfirmTime(confirmTime time.Time) *ImportContactsContact {
	return c.setField("confirm_time", confirmTime.Format(time.RFC3339), true)
}

func (c *ImportContactsContact) AddListID(listID int64, subscribeTime time.Time) *ImportContactsContact {
	c.listIDs[listID] = subscribeTime
	return c
}

func (c *ImportContactsContact) SetUnsubscribedListIDs(listIDs ...int64) *ImportContactsContact {
	return c.setField("unsubscribed_list_ids", c.int64SliceToString(listIDs...), true)
}

func (c *ImportContactsContact) SetExcludedListIDs(listIDs ...int64) *ImportContactsContact {
	return c.setField("excluded_list_ids", c.int64SliceToString(listIDs...), true)
}

func (c *ImportContactsContact) SetField(name, value string) *ImportContactsContact {
	return c.setField(name, value)
}

func (c *ImportContactsContact) setField(name, value string, withKind ...bool) *ImportContactsContact {
	if len(withKind) > 0 && withKind[0] {
		name = fmt.Sprintf("%s_%s", c.kind, name)
	}

	c.fields[name] = value
	c.collection.addFieldName(name)

	return c
}

func (c *ImportContactsContact) int64SliceToString(v ...int64) string {
	s := make([]string, len(v))
	for n, i := range v {
		s[n] = strconv.FormatInt(i, 10)
	}

	return strings.Join(s, ",")
}

func (c *ImportContactsContact) prepare() {
	l := len(c.listIDs)
	if l == 0 {
		return
	}

	listIDs := make([]string, l)
	subscribeTimes := make([]string, l)

	i := 0
	for listID, subscribeTime := range c.listIDs {
		listIDs[i] = strconv.FormatInt(listID, 10)
		subscribeTimes[i] = subscribeTime.Format(time.RFC3339)

		i++
	}

	c.setField("list_ids", strings.Join(listIDs, ","), true).
		setField("subscribe_times", strings.Join(subscribeTimes, ","), true)
}

func (c *ImportContactsContact) data() (data map[int]string) {
	data = map[int]string{}
	for n, fieldName := range c.collection.fieldNames {
		if v, ok := c.fields[fieldName]; ok {
			data[n] = v
		}
	}

	return data
}

type ImportContactsCollection struct {
	contacts   []*ImportContactsContact
	fieldNames []string
	prepared   bool
}

func (c *ImportContactsCollection) Email(email string) *ImportContactsContact {
	return c.newContact("email", email)
}

func (c *ImportContactsCollection) Phone(phone string) *ImportContactsContact {
	return c.newContact("phone", phone)
}

func (c *ImportContactsCollection) newContact(kind, contact string) (cnt *ImportContactsContact) {
	c.addFieldName(kind)

	cnt = &ImportContactsContact{
		collection: c,
		kind:       kind,
		fields: map[string]string{
			kind: contact,
		},
		listIDs: map[int64]time.Time{},
	}

	c.contacts = append(c.contacts, cnt)
	c.prepared = false

	return
}

func (c *ImportContactsCollection) addFieldName(fieldName string) {
	// Prevent duplication
	for _, name := range c.fieldNames {
		if name == fieldName {
			return
		}
	}

	c.fieldNames = append(c.fieldNames, fieldName)
	c.prepared = false
}

func (c *ImportContactsCollection) prepare() {
	if c.prepared {
		return
	}

	for _, contact := range c.contacts {
		contact.prepare()
	}

	c.prepared = true
}

func (c *ImportContactsCollection) FieldNames() []string {
	c.prepare()
	return c.fieldNames
}

func (c *ImportContactsCollection) Data() (data map[int]map[int]string) {
	c.prepare()

	data = map[int]map[int]string{}
	for n, contact := range c.contacts {
		data[n] = contact.data()
	}

	return
}

func NewImportContactsCollection() *ImportContactsCollection {
	return &ImportContactsCollection{}
}
