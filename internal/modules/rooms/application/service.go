package application

import (
	"Echo/internal/modules/rooms/domain"
)

type RoomService struct {
	repo domain.RoomRepository
}

func NewRoomService(r domain.RoomRepository) *RoomService {
	return &RoomService{repo: r}
}

func (s *RoomService) CreateRoom(dto RoomCreateRequest) (*RoomResponse, error) {
	room, err := domain.NewRoom(
		dto.Number,
		dto.Name,
		dto.Description,
		dto.SingleBeds,
		dto.DoubleBeds,
	)

	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(room); err != nil {
		return nil, err
	}

	return toRoomResponse(room), nil
}

func (s *RoomService) UpdateRoom(id int, dto RoomUpdateRequest) (*RoomResponse, error) {
	room, err := s.repo.GetById(id)

	if err != nil {
		return nil, err
	}

	if dto.Number != nil {
		if err := room.UpdateNumber(*dto.Number); err != nil {
			return nil, err
		}
	}

	if dto.Name != nil {
		if err := room.UpdateName(*dto.Name); err != nil {
			return nil, err
		}
	}

	if dto.Description != nil {
		room.UpdateDescription(*dto.Description)
	}

	if dto.SingleBeds != nil || dto.DoubleBeds != nil {
		single := room.SingleBeds()
		double := room.DoubleBeds()

		if dto.SingleBeds != nil {
			single = *dto.SingleBeds
		}
		if dto.DoubleBeds != nil {
			double = *dto.DoubleBeds
		}

		if err := room.UpdateBeds(single, double); err != nil {
			return nil, err
		}
	}

	if err := s.repo.Update(room); err != nil {
		return nil, err
	}

	return toRoomResponse(room), nil
}

func (s *RoomService) GetRoomById(id int) (*RoomResponse, error) {
	room, err := s.repo.GetById(id)

	if err != nil {
		return nil, err
	}

	return toRoomResponse(room), nil
}

func (s *RoomService) GetRoomByNumber(number int) (*RoomResponse, error) {
	room, err := s.repo.GetByNumber(number)

	if err != nil {
		return nil, err
	}

	return toRoomResponse(room), nil
}

func (s *RoomService) ListRooms() ([]*RoomResponse, error) {
	rooms, err := s.repo.List()
	if err != nil {
		return nil, err
	}

	result := make([]*RoomResponse, 0, len(rooms))
	for _, r := range rooms {
		result = append(result, toRoomResponse(r))
	}

	return result, nil
}

func (s *RoomService) DeleteRoom(id int) error {
	return s.repo.Delete(id)
}

func toRoomResponse(r *domain.Room) *RoomResponse {
	return &RoomResponse{
		ID:          r.ID(),
		Number:      r.Number(),
		Name:        r.Name(),
		Description: r.Description(),
		SingleBeds:  r.SingleBeds(),
		DoubleBeds:  r.DoubleBeds(),
	}
}
