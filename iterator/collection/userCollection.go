package collection

type UserCollection struct {
	Users []*User
}

func (u *UserCollection) CreateIterator() Iterator {
	return &userIterator{
		Users: u.Users,
	}
}
