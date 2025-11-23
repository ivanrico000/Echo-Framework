package persistence

import (
	"errors"

	"Echo/internal/modules/rooms/domain"

	"gorm.io/gorm"
)

// -------------------------
// MODELO GORM
// -------------------------
type RoomModel struct {
	ID          int `gorm:"primaryKey;autoIncrement"`
	Number      int
	Name        string
	Description string
	SingleBeds  int
	DoubleBeds  int
}

func (m *RoomModel) ToDomain() *domain.Room {
	return domain.RebuildRoom(
		m.ID,
		m.Number,
		m.Name,
		m.Description,
		m.SingleBeds,
		m.DoubleBeds,
	)
}

// Convierte dominio → GORM model
func FromDomain(r *domain.Room) *RoomModel {
	return &RoomModel{
		ID:          r.ID(),
		Number:      r.Number(),
		Name:        r.Name(),
		Description: r.Description(),
		SingleBeds:  r.SingleBeds(),
		DoubleBeds:  r.DoubleBeds(),
	}
}

// -------------------------
// REPOSITORY IMPLEMENTATION
// -------------------------
type RoomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) *RoomRepository {
	return &RoomRepository{db: db}
}

// -------------------------
// Create
// -------------------------
func (r *RoomRepository) Create(room *domain.Room) error {
	model := FromDomain(room)
	return r.db.Create(model).Error
}

// -------------------------
// Update
// -------------------------
func (r *RoomRepository) Update(room *domain.Room) error {
	model := FromDomain(room)
	return r.db.Save(model).Error
}

// -------------------------
// Delete
// -------------------------
func (r *RoomRepository) Delete(id int) error {
	return r.db.Delete(&RoomModel{}, id).Error
}

// -------------------------
// GetById
// -------------------------
func (r *RoomRepository) GetById(id int) (*domain.Room, error) {
	var model RoomModel

	err := r.db.First(&model, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrInvalidRoomID
	}
	if err != nil {
		return nil, err
	}

	return model.ToDomain(), nil
}

// -------------------------
// GetByNumber
// -------------------------
func (r *RoomRepository) GetByNumber(number int) (*domain.Room, error) {
	var model RoomModel

	err := r.db.Where("number = ?", number).First(&model).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrInvalidRoomNumber
	}
	if err != nil {
		return nil, err
	}

	return model.ToDomain(), nil
}

// -------------------------
// List
// -------------------------
func (r *RoomRepository) List() ([]*domain.Room, error) {
	var models []RoomModel
	err := r.db.Find(&models).Error
	if err != nil {
		return nil, err
	}

	rooms := make([]*domain.Room, 0, len(models))

	for _, m := range models {
		room := m.ToDomain()
		if room == nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}
