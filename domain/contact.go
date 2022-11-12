package domain

type Contact struct {
	contactType string
	detail      string
}

func NewContact(contactType, detail string) Contact {
	return Contact{
		contactType: contactType,
		detail:      detail,
	}
}

func (c *Contact) GetType() string {
	return c.contactType
}

func (c *Contact) GetDetail() string {
	return c.detail
}

func (c *Contact) SetDetail(detail string) {
	c.detail = detail
}
