package cqrs

import (
	"testing"

	d "github.com/calogxro/cqrs-es/domain"
	"github.com/calogxro/cqrs-es/utils"
)

func TestCQRS(t *testing.T) {
	// Setup

	wr := NewWriteRepository()
	rr := NewReadRepository()
	ua := NewUserAggregate(wr)
	up := NewUserProjection(rr)
	projector := NewProjector(rr)

	userId := "1"

	c1 := d.NewContact("EMAIL", "email_1@example.com")
	c2 := d.NewContact("PHONE", "1234567890")
	c3 := d.NewContact("EMAIL", "email_2@example.com")

	a1 := d.NewAddress("Rome", "Italy")
	a2 := d.NewAddress("London", "UK")
	a3 := d.NewAddress("Palermo", "Italy")

	var user *d.User
	var contactByType *utils.Set[d.Contact]
	var addressByRegion *utils.Set[d.Address]

	// Create

	user = ua.handleCreateUserCommand(CreateUserCommand{
		UserId:    userId,
		FirstName: "X",
		LastName:  "Y",
	})
	projector.project(user)

	// Update 1

	user, _ = ua.handleUpdateUserCommand(UpdateUserCommand{
		UserId:    userId,
		Contacts:  utils.ArraytoSet([]d.Contact{c1, c2, c3}),
		Addresses: utils.ArraytoSet([]d.Address{a1, a2, a3}),
	})
	projector.project(user)

	// Queries 1

	contactByType, _ = up.HandleContactByTypeQuery(ContactByTypeQuery{
		UserId:      userId,
		ContactType: "EMAIL",
	})

	addressByRegion, _ = up.HandleAddressByRegionQuery(AddressByRegionQuery{
		UserId: userId,
		State:  "Italy",
	})

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

	user, _ = ua.handleUpdateUserCommand(UpdateUserCommand{
		UserId:    userId,
		Contacts:  utils.ArraytoSet([]d.Contact{c2, c3}),
		Addresses: utils.ArraytoSet([]d.Address{a2, a3}),
	})
	projector.project(user)

	// Queries 2

	contactByType, _ = up.HandleContactByTypeQuery(ContactByTypeQuery{
		UserId:      userId,
		ContactType: "EMAIL",
	})

	addressByRegion, _ = up.HandleAddressByRegionQuery(AddressByRegionQuery{
		UserId: userId,
		State:  "Italy",
	})

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
