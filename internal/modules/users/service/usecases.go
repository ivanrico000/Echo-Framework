package service

import (
	"Echo/internal/modules/users/core"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"

	"golang.org/x/crypto/argon2"
)

type UserUseCases struct {
	Repo core.UserRepository
}

func NewUserUseCases(r core.UserRepository) *UserUseCases {
	return &UserUseCases{Repo: r}
}

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16) // OWASP recomienda entre 16 y 32 bytes
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func hashPassword(password string) string {
	// Parámetros recomendados para 2025 (balance seguridad/velocidad)
	const time = 1           // nro de iteraciones
	const memory = 64 * 1024 // 64 MB
	const threads = 4
	const keyLength = 32

	salt, err := generateSalt()
	if err != nil {
		return ""
	}
	hash := argon2.IDKey([]byte(password), salt, time, memory, uint8(threads), keyLength)

	b64Hash := base64.RawStdEncoding.EncodeToString(hash)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)

	encoded := fmt.Sprintf(
		"$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s",
		memory, time, threads,
		b64Salt, b64Hash,
	)

	return encoded
}

func (uc *UserUseCases) CreateUser(dto CreateUserDTO) error {
	hash := hashPassword(dto.Password)

	hotelID := int64(dto.HotelID)
	roleID := int64(dto.RoleID)

	user := core.User{
		HotelID:  &hotelID,
		RoleID:   &roleID,
		Name:     dto.Name,
		Email:    dto.Email,
		Phone:    dto.Phone,
		Password: hash,
	}

	return uc.Repo.Create(&user)
}

func (uc *UserUseCases) UpdateUser(id int, dto UpdateUserDTO) error {
	user, err := uc.Repo.FindByID(id)
	if err != nil || user == nil {
		return errors.New("user not found")
	}

	roleID := int64(dto.RoleID)

	user.Email = dto.Email
	user.Phone = dto.Phone
	user.RoleID = &roleID

	return uc.Repo.Update(user)
}

func (uc *UserUseCases) DeleteUser(id int) error {
	return uc.Repo.Delete(id)
}

func (uc *UserUseCases) GetByID(id int) (*core.User, error) {
	user, err := uc.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (uc *UserUseCases) ListByHotel(hotelID int) ([]core.User, error) {
	return uc.Repo.FindByHotel(hotelID)
}
