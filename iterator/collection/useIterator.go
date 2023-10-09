package collection

type userIterator struct {
	Index int
	Users []*User
}

func (u *userIterator) HasNext() bool {
	if u.Index < len(u.Users) {
		return true
	}
	return false
}

func (u *userIterator) GetNext() *User {
	if u.HasNext() {
		user := u.Users[u.Index]
		u.Index++
		return user
	}
	return nil
}
