package service

type CreateUserDTO struct {
	HotelID  int    `msgpack:"hotel_id"`
	RoleID   int    `msgpack:"role_id"`
	Name     string `msgpack:"username"`
	Email    string `msgpack:"email"`
	Password string `msgpack:"password"`
	Phone    string `msgpack:"phone"`
}

type UpdateUserDTO struct {
	Email  string `msgpack:"email"`
	RoleID int    `msgpack:"role_id"`
	Phone  string `msgpack:"phone"`
}
