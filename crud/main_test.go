package crud

import (
	"testing"

	d "github.com/calogxro/cqrs-es/domain"
	"github.com/calogxro/cqrs-es/utils"
)

func TestCRUD(t *testing.T) {
	repo := NewUserRepository()
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

	var email_contacts *utils.Set[d.Contact]
	var phone_contacts *utils.Set[d.Contact]
	var italy_addresses *utils.Set[d.Address]
	var uk_addresses *utils.Set[d.Address]

	// Create

	service.createUser(userId, "fistName1", "lastName1")

	// Update 1

	contacts = utils.ArraytoSet([]d.Contact{c1, c2, c3})
	addresses = utils.ArraytoSet([]d.Address{a1, a2, a3})

	service.updateUser(userId, contacts, addresses)

	if len(repo.store) != 1 {
		t.Fatal("")
	}

	// Queries 1 - contact by type

	email_contacts, _ = service.getContactByType(userId, "EMAIL")
	phone_contacts, _ = service.getContactByType(userId, "PHONE")

	if email_contacts.Size() != 2 {
		t.Fatal("")
	}

	if !email_contacts.Contains(c1) || !email_contacts.Contains(c3) {
		t.Fatal("")
	}

	if phone_contacts.Size() != 1 {
		t.Fatal("")
	}

	if !phone_contacts.Contains(c2) {
		t.Fatal("")
	}

	// Queries 1 - address by region

	italy_addresses, _ = service.getAddressByRegion(userId, "Italy")
	uk_addresses, _ = service.getAddressByRegion(userId, "UK")

	if italy_addresses.Size() != 2 {
		t.Fatal("")
	}

	if !italy_addresses.Contains(a1) || !italy_addresses.Contains(a3) {
		t.Fatal("")
	}

	if uk_addresses.Size() != 1 {
		t.Fatal("")
	}

	if !uk_addresses.Contains(a2) {
		t.Fatal("")
	}

	// Update 2

	contacts = utils.ArraytoSet([]d.Contact{c2, c3})
	addresses = utils.ArraytoSet([]d.Address{a2, a3})

	service.updateUser(userId, contacts, addresses)

	if len(repo.store) != 1 {
		t.Fatal("")
	}

	// Queries 2 - contact by type

	email_contacts, _ = service.getContactByType(userId, "EMAIL")
	phone_contacts, _ = service.getContactByType(userId, "PHONE")

	if email_contacts.Size() != 1 {
		t.Fatal("")
	}

	if !email_contacts.Contains(c3) {
		t.Fatal("")
	}

	if phone_contacts.Size() != 1 {
		t.Fatal("")
	}

	if !phone_contacts.Contains(c2) {
		t.Fatal("")
	}

	// Queries 2 - address by region

	italy_addresses, _ = service.getAddressByRegion(userId, "Italy")
	uk_addresses, _ = service.getAddressByRegion(userId, "UK")

	if italy_addresses.Size() != 1 {
		t.Fatal("")
	}

	if !italy_addresses.Contains(a3) {
		t.Fatal("")
	}

	if uk_addresses.Size() != 1 {
		t.Fatal("")
	}

	if !uk_addresses.Contains(a2) {
		t.Fatal("")
	}
}
