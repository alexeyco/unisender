package contacts_test

import (
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/alexeyco/unisender/contacts"
	"github.com/alexeyco/unisender/test"
)

func TestImportContactsContact_Delete(t *testing.T) {
	collection := contacts.NewImportContactsCollection()
	collection.Email(test.RandomString(12, 36)).
		Delete()

	testImportContactsContactField(t, collection, "delete", "1")
}

func TestImportContactsContact_SetTags(t *testing.T) {
	expectedTags := test.RandomStringSlice(12, 36)
	givenTags := strings.Join(expectedTags, ",")

	collection := contacts.NewImportContactsCollection()
	collection.Email(test.RandomString(12, 36)).
		SetTags(expectedTags...)

	testImportContactsContactField(t, collection, "tags", givenTags)
}

func TestImportContactsContact_SetStatusNew(t *testing.T) {
	collection := contacts.NewImportContactsCollection()
	collection.Email(test.RandomString(12, 36)).
		SetStatusNew()

	testImportContactsContactField(t, collection, "email_status", "new")
}

func TestImportContactsContact_SetStatusActive(t *testing.T) {
	collection := contacts.NewImportContactsCollection()
	collection.Email(test.RandomString(12, 36)).
		SetStatusActive()

	testImportContactsContactField(t, collection, "email_status", "active")
}

func TestImportContactsContact_SetStatusInactive(t *testing.T) {
	collection := contacts.NewImportContactsCollection()
	collection.Email(test.RandomString(12, 36)).
		SetStatusInactive()

	testImportContactsContactField(t, collection, "email_status", "inactive")
}

func TestImportContactsContact_SetStatusUnsubscribed(t *testing.T) {
	collection := contacts.NewImportContactsCollection()
	collection.Email(test.RandomString(12, 36)).
		SetStatusUnsubscribed()

	testImportContactsContactField(t, collection, "email_status", "unsubscribed")
}

func TestImportContactsContact_SetAvailabilityAvailable(t *testing.T) {
	collection := contacts.NewImportContactsCollection()
	collection.Email(test.RandomString(12, 36)).
		SetAvailabilityAvailable()

	testImportContactsContactField(t, collection, "email_availability", "available")
}

func TestImportContactsContact_SetAvailabilityUnreachable(t *testing.T) {
	collection := contacts.NewImportContactsCollection()
	collection.Email(test.RandomString(12, 36)).
		SetAvailabilityUnreachable()

	testImportContactsContactField(t, collection, "email_availability", "unreachable")
}

func TestImportContactsContact_SetAvailabilityTempUnreachable(t *testing.T) {
	collection := contacts.NewImportContactsCollection()
	collection.Email(test.RandomString(12, 36)).
		SetAvailabilityTempUnreachable()

	testImportContactsContactField(t, collection, "email_availability", "temp_unreachable")
}

func TestImportContactsContact_SetAvailabilityMailboxFull(t *testing.T) {
	collection := contacts.NewImportContactsCollection()
	collection.Email(test.RandomString(12, 36)).
		SetAvailabilityMailboxFull()

	testImportContactsContactField(t, collection, "email_availability", "mailbox_full")
}

func TestImportContactsContact_SetAvailabilitySpamRejected(t *testing.T) {
	collection := contacts.NewImportContactsCollection()
	collection.Email(test.RandomString(12, 36)).
		SetAvailabilitySpamRejected()

	testImportContactsContactField(t, collection, "email_availability", "spam_rejected")
}

func TestImportContactsContact_SetAvailabilitySpamFolder(t *testing.T) {
	collection := contacts.NewImportContactsCollection()
	collection.Email(test.RandomString(12, 36)).
		SetAvailabilitySpamFolder()

	testImportContactsContactField(t, collection, "email_availability", "spam_folder")
}

func TestImportContactsContact_SetAddTime(t *testing.T) {
	addTime := test.RandomTime(12, 365)

	collection := contacts.NewImportContactsCollection()
	collection.Email(test.RandomString(12, 36)).
		SetAddTime(addTime)

	testImportContactsContactField(t, collection, "email_add_time", addTime.Format(time.RFC3339))
}

func TestImportContactsContact_SetConfirmTime(t *testing.T) {
	confirmedTime := test.RandomTime(12, 365)

	collection := contacts.NewImportContactsCollection()
	collection.Email(test.RandomString(12, 36)).
		SetConfirmTime(confirmedTime)

	testImportContactsContactField(t, collection, "email_confirm_time", confirmedTime.Format(time.RFC3339))
}

func TestImportContactsContact_AddListID(t *testing.T) {
	l := test.RandomInt(12, 36)

	expectedListIDs := make([]string, l)
	var givenListIDs []string

	expectedSubscribeTimes := make([]string, l)
	var givenSubscribeTimes []string

	collection := contacts.NewImportContactsCollection()
	contact := collection.Email(test.RandomString(12, 36))

	for i := 0; i < l; i++ {
		listID := test.RandomInt64(9999, 999999)
		subscribeTime := test.RandomTime(12, 365)

		expectedListIDs[i] = strconv.FormatInt(listID, 10)
		expectedSubscribeTimes[i] = subscribeTime.Format(time.RFC3339)

		contact.AddListID(listID, subscribeTime)
	}

	fieldNames := collection.FieldNames()
	data := collection.Data()

	c, ok := data[0]
	if !ok {
		t.Fatal("Contact should exist")
	}

	n, ok := hasField("email_list_ids", fieldNames)
	if !ok {
		t.Fatalf(`Field names should have "%s" field`, "email_list_ids")
	}

	listIDs, ok := c[n]
	if !ok {
		t.Fatalf(`Contact should have "%s" field value`, "email_list_ids")
	}

	n, ok = hasField("email_subscribe_times", fieldNames)
	if !ok {
		t.Fatalf(`Field names should have "%s" field`, "email_subscribe_times")
	}

	subscribeTimes, ok := c[n]
	if !ok {
		t.Fatalf(`Contact should have "%s" field value`, "email_subscribe_times")
	}

	givenListIDs = strings.Split(listIDs, ",")
	givenSubscribeTimes = strings.Split(subscribeTimes, ",")

	sort.Strings(expectedListIDs)
	sort.Strings(givenListIDs)

	if !reflect.DeepEqual(expectedListIDs, givenListIDs) {
		t.Fatal("List IDs should be equal")
	}

	sort.Strings(expectedSubscribeTimes)
	sort.Strings(givenSubscribeTimes)

	if !reflect.DeepEqual(expectedSubscribeTimes, givenSubscribeTimes) {
		t.Fatal("Subscribe times should be equal")
	}
}

func TestImportContactsContact_SetUnsubscribedListIDs(t *testing.T) {
	listIDs := test.RandomInt64Slice(12, 36)
	l := make([]string, len(listIDs))
	for n, id := range listIDs {
		l[n] = strconv.FormatInt(id, 10)
	}

	collection := contacts.NewImportContactsCollection()
	collection.Email(test.RandomString(12, 36)).
		SetUnsubscribedListIDs(listIDs...)

	testImportContactsContactField(t, collection, "email_unsubscribed_list_ids", strings.Join(l, ","))
}

func TestImportContactsContact_SetExcludedListIDs(t *testing.T) {
	listIDs := test.RandomInt64Slice(12, 36)
	l := make([]string, len(listIDs))
	for n, id := range listIDs {
		l[n] = strconv.FormatInt(id, 10)
	}

	collection := contacts.NewImportContactsCollection()
	collection.Email(test.RandomString(12, 36)).
		SetExcludedListIDs(listIDs...)

	testImportContactsContactField(t, collection, "email_excluded_list_ids", strings.Join(l, ","))
}

func TestImportContactsContact_SetField(t *testing.T) {
	fieldName := test.RandomString(12, 36)
	fieldValue := test.RandomString(12, 36)

	collection := contacts.NewImportContactsCollection()
	collection.Email(test.RandomString(12, 36)).
		SetField(fieldName, fieldValue)

	testImportContactsContactField(t, collection, fieldName, fieldValue)
}

func testImportContactsContactField(
	t *testing.T,
	collection *contacts.ImportContactsCollection,
	field, expectedValue string,
) {
	var givenValue string

	fieldNames := collection.FieldNames()
	data := collection.Data()

	n, ok := hasField(field, fieldNames)
	if !ok {
		t.Fatalf(`Field names should have "%s" field`, field)
	}

	contact, ok := data[0]
	if !ok {
		t.Fatal("Contact should exist")
	}

	givenValue, ok = contact[n]
	if !ok {
		t.Fatalf(`Contact should have "%s" field value`, field)
	}

	if givenValue != expectedValue {
		t.Fatalf(`Field "%s" should be "%s", "%s" given`, field, expectedValue, givenValue)
	}
}

func TestImportContactsCollection_Email(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)
	var givenEmail string

	collection := contacts.NewImportContactsCollection()
	collection.Email(expectedEmail)

	fieldNames := collection.FieldNames()
	data := collection.Data()

	n, ok := hasField("email", fieldNames)
	if !ok {
		t.Fatalf(`Field names should have "%s"`, "email")
	}

	contact, ok := data[0]
	if !ok {
		t.Fatal("Contact should exist")
	}

	givenEmail, ok = contact[n]
	if !ok {
		t.Fatalf(`Contact should have "%s" field value`, "email")
	}

	if givenEmail != expectedEmail {
		t.Fatalf(`Email should be "%s", "%s" given`, expectedEmail, givenEmail)
	}
}

func TestImportContactsCollection_Phone(t *testing.T) {
	expectedPhone := test.RandomString(12, 36)
	var givenPhone string

	collection := contacts.NewImportContactsCollection()
	collection.Phone(expectedPhone)

	fieldNames := collection.FieldNames()
	data := collection.Data()

	n, ok := hasField("phone", fieldNames)
	if !ok {
		t.Fatalf(`Field names should have "%s"`, "phone")
	}

	contact, ok := data[0]
	if !ok {
		t.Fatal("Contact should exist")
	}

	givenPhone, ok = contact[n]
	if !ok {
		t.Fatalf(`Contact should have "%s" field value`, "phone")
	}

	if givenPhone != expectedPhone {
		t.Fatalf(`Phone should be "%s", "%s" given`, expectedPhone, givenPhone)
	}
}

func TestImportContactsCollection_FieldNames(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)

	expectedFieldNames := append([]string{"email"}, test.RandomStringSlice(12, 36)...)
	var givenFieldNames []string

	collection := contacts.NewImportContactsCollection()
	contact := collection.Email(expectedEmail)

	for _, fieldName := range expectedFieldNames {
		if fieldName == "email" {
			continue
		}

		contact.SetField(fieldName, test.RandomString(12, 36))
	}

	givenFieldNames = collection.FieldNames()
	if !reflect.DeepEqual(expectedFieldNames, givenFieldNames) {
		t.Fatal("Field names should be equal")
	}
}

func hasField(fieldName string, fieldNames []string) (n int, ok bool) {
	for i, name := range fieldNames {
		if name == fieldName {
			n = i
			ok = true

			return
		}
	}

	return
}
