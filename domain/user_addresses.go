package domain

import "github.com/calogxro/cqrs-es/utils"

type UserAddresses struct {
	addressesByRegion map[string]*utils.Set[Address]
}

func NewUserAddresses() *UserAddresses {
	return &UserAddresses{
		addressesByRegion: make(map[string]*utils.Set[Address]),
	}
}

func (c *UserAddresses) GetAddressesByRegion(state string) *utils.Set[Address] {
	return c.addressesByRegion[state]
}

func (c *UserAddresses) AddAddress(address Address) {
	if c.addressesByRegion[address.state] == nil {
		c.addressesByRegion[address.state] = utils.NewSet[Address]()
	}
	c.addressesByRegion[address.state].Add(address)
}

func (c *UserAddresses) RemoveAddress(address Address) {
	if c.addressesByRegion[address.state] != nil {
		c.addressesByRegion[address.state].Remove(address)
	}
}
