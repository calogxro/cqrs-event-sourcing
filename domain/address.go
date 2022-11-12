package domain

type Address struct {
	city  string
	state string
	//postCode string
}

func NewAddress(city, state string) Address {
	return Address{
		city:  city,
		state: state,
	}
}

func (a *Address) GetCity() string {
	return a.city
}

func (a *Address) GetState() string {
	return a.state
}
