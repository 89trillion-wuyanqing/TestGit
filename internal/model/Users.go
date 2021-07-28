package model

type Users struct {
	id    string
	Name  string
	Age   int
	Phone string
}

func (u *Users) GetId() string {
	return u.id
}

func (u *Users) SetId(id string) {
	u.id = id
}
