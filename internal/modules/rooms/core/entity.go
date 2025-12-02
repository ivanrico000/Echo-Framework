package domain

import "errors"

type Room struct {
	id          int
	number      int
	name        string
	description string
	singleBeds  int
	doubleBeds  int
}

var (
	ErrInvalidRoomID         = errors.New("Id invalido")
	ErrInvalidRoomNumber     = errors.New("Error consultando por numero de habitacion")
	ErrInvalidRoomName       = errors.New("Nombre de habitacion no puede estar vacio")
	ErrInvalidSingleBeds     = errors.New("El numero de camas sencillas debe ser igual o mayor a 0")
	ErrInvalidDoubleBeds     = errors.New("El numero de camas dobles debe ser igual o mayor a 0")
	ErrInvalidBedCombination = errors.New("Se requiere almenos una cama")
)

func NewRoom(number int, name, description string, singleBeds, doubleBeds int) (*Room, error) {

	if number <= 0 {
		return nil, ErrInvalidRoomNumber
	}

	if name == "" {
		return nil, ErrInvalidRoomName
	}

	if singleBeds < 0 {
		return nil, ErrInvalidSingleBeds
	}

	if doubleBeds < 0 {
		return nil, ErrInvalidDoubleBeds
	}

	if singleBeds+doubleBeds == 0 {
		return nil, ErrInvalidBedCombination
	}

	r := &Room{
		number:      number,
		name:        name,
		description: description,
		singleBeds:  singleBeds,
		doubleBeds:  doubleBeds,
	}

	return r, nil
}

func RebuildRoom(id, number int, name, description string, singleBeds, doubleBeds int) *Room {
	return &Room{
		id:          id,
		number:      number,
		name:        name,
		description: description,
		singleBeds:  singleBeds,
		doubleBeds:  doubleBeds,
	}
}

func (r *Room) ID() int {
	return r.id
}

func (r *Room) Number() int {
	return r.number
}

func (r *Room) Name() string {
	return r.name
}

func (r *Room) Description() string {
	return r.description
}

func (r *Room) SingleBeds() int {
	return r.singleBeds
}

func (r *Room) DoubleBeds() int {
	return r.doubleBeds
}

func (r *Room) UpdateName(name string) error {
	if name == "" {
		return ErrInvalidRoomName
	}
	r.name = name
	return nil
}

func (r *Room) UpdateDescription(desc string) {
	r.description = desc
}

func (r *Room) UpdateBeds(singleBeds, doubleBeds int) error {
	if singleBeds < 0 {
		return ErrInvalidSingleBeds
	}

	if doubleBeds < 0 {
		return ErrInvalidDoubleBeds
	}

	if singleBeds+doubleBeds == 0 {
		return ErrInvalidBedCombination
	}

	r.singleBeds = singleBeds
	r.doubleBeds = doubleBeds

	return nil
}

func (r *Room) UpdateNumber(number int) error {
	if number <= 0 {
		return ErrInvalidRoomNumber
	}

	r.number = number

	return nil
}
