package es

import (
	"testing"

	d "github.com/calogxro/cqrs-es/domain"
	"github.com/calogxro/cqrs-es/utils"
)

func TestSetOfContacts(t *testing.T) {
	contact1 := d.NewContact("EMAIL", "abc@example.com")
	contact2 := d.NewContact("PHONE", "123456")
	contact3 := d.NewContact("EMAIL", "abc@example.com")
	contact4 := d.NewContact("EMAIL", "xyz@example.com")

	set := utils.NewSet[d.Contact]()
	set.Add(contact1)
	set.Add(contact2)

	if !set.Contains(contact1) {
		t.Fatal("")
	}

	if !set.Contains(contact2) {
		t.Fatal("")
	}

	if !set.Contains(contact3) {
		t.Fatal("")
	}

	if set.Contains(contact4) {
		t.Fatal("")
	}
}

func TestEventSourcing(t *testing.T) {
	repo := NewEventStore()
	service := NewUserService(repo)

	userId := "1"

	c1 := d.NewContact("EMAIL", "email_1@example.com")
	c2 := d.NewContact("PHONE", "1234567890")
	c3 := d.NewContact("EMAIL", "email_2@example.com")

	a1 := d.NewAddress("Rome", "Italy")
	a2 := d.NewAddress("London", "UK")
	a3 := d.NewAddress("Palermo", "Italy")

	var contacts *utils.Set[d.Contact]
	var addresses *utils.Set[d.Address]
	var contactByType *utils.Set[d.Contact]
	var addressByRegion *utils.Set[d.Address]

	// Create

	service.createUser(userId, "fistName1", "lastName1")

	// Update 1

	contacts = utils.ArraytoSet([]d.Contact{c1, c2, c3})
	addresses = utils.ArraytoSet([]d.Address{a1, a2, a3})

	service.updateUser(userId, contacts, addresses)

	// Queries 1

	contactByType, _ = service.getContactByType(userId, "EMAIL")
	addressByRegion, _ = service.getAddressByRegion(userId, "Italy")

	if contactByType.Size() != 2 {
		t.Fatal("")
	}

	if !contactByType.Contains(c1) || !contactByType.Contains(c3) {
		t.Fatal("")
	}

	if addressByRegion.Size() != 2 {
		t.Fatal("")
	}

	if !addressByRegion.Contains(a1) || !addressByRegion.Contains(a3) {
		t.Fatal("")
	}

	// Update 2

	contacts = utils.ArraytoSet([]d.Contact{c2, c3})
	addresses = utils.ArraytoSet([]d.Address{a2, a3})

	service.updateUser(userId, contacts, addresses)

	// Queries 2

	contactByType, _ = service.getContactByType(userId, "EMAIL")
	addressByRegion, _ = service.getAddressByRegion(userId, "Italy")

	if contactByType.Size() != 1 {
		t.Fatal("")
	}

	if !contactByType.Contains(c3) {
		t.Fatal("")
	}

	if addressByRegion.Size() != 1 {
		t.Fatal("")
	}

	if !addressByRegion.Contains(a3) {
		t.Fatal("")
	}
}
