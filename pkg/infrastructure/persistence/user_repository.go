package persistence

import (
	"Echo/internal/modules/users/core"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) core.UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) FindByID(id int) (*core.User, error) {
	var user core.User
	err := r.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByHotel(hotelID int) ([]core.User, error) {
	var users []core.User
	err := r.DB.Where("hotel_id = ?", hotelID).Find(&users).Error
	return users, err
}

func (r *UserRepository) Create(user *core.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) Update(user *core.User) error {
	return r.DB.Save(user).Error
}

func (r *UserRepository) Delete(id int) error {
	return r.DB.Delete(&core.User{}, id).Error
}
