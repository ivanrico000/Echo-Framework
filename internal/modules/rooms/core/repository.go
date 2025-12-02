package domain

type RoomRepository interface {
	Create(room *Room) error
	Update(room *Room) error
	Delete(id int) error

	GetById(id int) (*Room, error)
	GetByNumber(number int) (*Room, error)
	List() ([]*Room, error)
}
