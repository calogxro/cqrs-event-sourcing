package crud

import (
	"errors"

	d "github.com/calogxro/cqrs-es/domain"
	"github.com/calogxro/cqrs-es/utils"
)

type UserService struct {
	repository *UserRepository
}

func NewUserService(repository *UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) createUser(userId, firstName, lastName string) {
	user := d.NewUser(userId, firstName, lastName)
	s.repository.addUser(user)
}

func (s *UserService) updateUser(userId string, contacts *utils.Set[d.Contact], addresses *utils.Set[d.Address]) error {
	user := s.repository.getUser(userId)
	if user == nil {
		return errors.New("")
	}
	user.SetContacts(contacts)
	user.SetAddresses(addresses)
	//s.repository.updateUser(user)
	return nil
}

func (s *UserService) getContactByType(userId string, contactType string) (*utils.Set[d.Contact], error) {
	user := s.repository.getUser(userId)
	if user == nil {
		return nil, errors.New("")
	}
	contacts := utils.NewSet[d.Contact]()
	for _, contact := range user.GetContacts().Values() {
		if contact.GetType() == contactType {
			contacts.Add(contact)
		}
	}
	return contacts, nil
}

func (s *UserService) getAddressByRegion(userId string, state string) (*utils.Set[d.Address], error) {
	user := s.repository.getUser(userId)
	if user == nil {
		return nil, errors.New("")
	}
	addresses := utils.NewSet[d.Address]()
	for _, address := range user.GetAddresses().Values() {
		if address.GetState() == state {
			addresses.Add(address)
		}
	}
	return addresses, nil
}
