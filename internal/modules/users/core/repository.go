package core

type UserRepository interface {
	FindByID(id int) (*User, error)
	FindByHotel(hotelID int) ([]User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id int) error
}
