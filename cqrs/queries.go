package cqrs

type ContactByTypeQuery struct {
	UserId      string
	ContactType string
}

type AddressByRegionQuery struct {
	UserId string
	State  string
}
