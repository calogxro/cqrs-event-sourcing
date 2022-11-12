package domain

import (
	"fmt"
	"testing"
)

func TestUserContacts(t *testing.T) {
	uc := NewUserContacts()
	uc.AddContact(NewContact("phone", "123456"))
	contacts := uc.GetContactsByType("phone")
	fmt.Println(contacts)
	fmt.Println(uc)
}
