package cqrs

import (
	"errors"

	d "github.com/calogxro/cqrs-es/domain"
	"github.com/calogxro/cqrs-es/utils"
)

type UserProjection struct {
	repository *ReadRepository
}

func NewUserProjection(repository *ReadRepository) *UserProjection {
	return &UserProjection{
		repository: repository,
	}
}

// Get contact by type

func (s *UserProjection) HandleContactByTypeQuery(query ContactByTypeQuery) (*utils.Set[d.Contact], error) {
	userContacts := s.repository.GetUserContacts(query.UserId)
	if userContacts == nil {
		return nil, errors.New("")
	}
	contactsByType := userContacts.GetContactsByType(query.ContactType)
	return contactsByType, nil
}

// Get address by region/state

func (s *UserProjection) HandleAddressByRegionQuery(query AddressByRegionQuery) (*utils.Set[d.Address], error) {
	userAddresses := s.repository.GetUserAddresses(query.UserId)
	if userAddresses == nil {
		return nil, errors.New("")
	}
	addressesByRegion := userAddresses.GetAddressesByRegion(query.State)
	return addressesByRegion, nil
}
