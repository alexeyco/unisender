package contacts

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Contact struct {
	mu         sync.RWMutex
	collection *Collection
	kind       string
	fields     map[string]string
	listIDs    map[int64]time.Time
}

func (c *Contact) Delete() *Contact {
	return c.setField("delete", "1")
}

func (c *Contact) SetTags(tags ...string) *Contact {
	return c.setField("tags", strings.Join(tags, ","))
}

func (c *Contact) SetStatusNew() *Contact {
	return c.setField("status", "new", true)
}

func (c *Contact) SetStatusActive() *Contact {
	return c.setField("status", "active", true)
}

func (c *Contact) SetStatusInactive() *Contact {
	return c.setField("status", "inactive", true)
}

func (c *Contact) SetStatusUnsubscribed() *Contact {
	return c.setField("status", "unsubscribed", true)
}

func (c *Contact) SetAvailabilityAvailable() *Contact {
	return c.setField("availability", "available", true)
}

func (c *Contact) SetAvailabilityUnreachable() *Contact {
	return c.setField("availability", "unreachable", true)
}

func (c *Contact) SetAvailabilityTempUnreachable() *Contact {
	return c.setField("availability", "temp_unreachable", true)
}

func (c *Contact) SetAvailabilityMailboxFull() *Contact {
	return c.setField("availability", "mailbox_full", true)
}

func (c *Contact) SetAvailabilitySpamRejected() *Contact {
	return c.setField("availability", "spam_rejected", true)
}

func (c *Contact) SetAvailabilitySpamFolder() *Contact {
	return c.setField("availability", "spam_folder", true)
}

func (c *Contact) SetAddTime(addTime time.Time) *Contact {
	return c.setField("add_time", addTime.Format(time.RFC3339), true)
}

func (c *Contact) SetConfirmTime(confirmTime time.Time) *Contact {
	return c.setField("confirm_time", confirmTime.Format(time.RFC3339), true)
}

func (c *Contact) AddListID(listID int64, subscribeTime time.Time) *Contact {
	c.mu.Lock()
	c.listIDs[listID] = subscribeTime

	l := len(c.listIDs)
	listIDs := make([]string, l)
	subscribeTimes := make([]string, l)

	i := 0
	for listID, subscribeTime := range c.listIDs {
		listIDs[i] = strconv.FormatInt(listID, 10)
		subscribeTimes[i] = subscribeTime.Format(time.RFC3339)

		i++
	}

	c.mu.Unlock()

	c.setField("list_ids", strings.Join(listIDs, ","), true)
	c.setField("subscribe_times", strings.Join(subscribeTimes, ","), true)

	return c
}

func (c *Contact) SetUnsubscribedListIDs(listIDs ...int64) *Contact {
	return c.setField("unsubscribed_list_ids", c.int64SliceToString(listIDs...), true)
}

func (c *Contact) SetExcludedListIDs(listIDs ...int64) *Contact {
	return c.setField("excluded_list_ids", c.int64SliceToString(listIDs...), true)
}

func (c *Contact) setField(name, value string, withKind ...bool) *Contact {
	c.mu.Lock()
	defer c.mu.Unlock()

	if len(withKind) > 0 && withKind[0] {
		name = fmt.Sprintf("%s_%s", c.kind, name)
	}

	c.fields[name] = value
	c.collection.addFieldName(name)

	return c
}

func (c *Contact) int64SliceToString(v ...int64) string {
	s := make([]string, len(v))
	for n, i := range v {
		s[n] = strconv.FormatInt(i, 10)
	}

	return strings.Join(s, ",")
}

func (c *Contact) data() (data map[int]string) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	data = map[int]string{}
	for n, fieldName := range c.collection.fieldNames {
		if v, ok := c.fields[fieldName]; ok {
			data[n] = v
		}
	}

	return data
}

type Collection struct {
	mu         sync.RWMutex
	contacts   []*Contact
	fieldNames []string
}

func (c *Collection) Email(email string) *Contact {
	return c.newContact("email", email)
}

func (c *Collection) Phone(phone string) *Contact {
	return c.newContact("phone", phone)
}

func (c *Collection) newContact(kind, contact string) (cnt *Contact) {
	c.addFieldName(kind)

	cnt = &Contact{
		collection: c,
		kind:       kind,
		fields: map[string]string{
			kind: contact,
		},
		listIDs: map[int64]time.Time{},
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	c.contacts = append(c.contacts, cnt)

	return
}

func (c *Collection) addFieldName(fieldName string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Prevent duplication
	for _, name := range c.fieldNames {
		if name == fieldName {
			return
		}
	}

	c.fieldNames = append(c.fieldNames, fieldName)
}

func (c *Collection) FieldNames() []string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.fieldNames
}

func (c *Collection) Data() (data map[int]map[int]string) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	data = map[int]map[int]string{}
	for n, contact := range c.contacts {
		data[n] = contact.data()
	}

	return
}

func NewCollection() *Collection {
	return &Collection{}
}
