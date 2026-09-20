package repository

import (
	"github.com/aruncs31s/go_gin_api_starterkit/internal/domain/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	// CreateUser creates a new user in the database.
	CreateUser(user *models.User) error

	// GetUserByID retrieves a user by their ID.
	GetUserByID(id int) (*models.User, error)

	// UpdateUser updates an exi:w
	// sting user's information.
	UpdateUser(user *models.User) error

	// DeleteUser deletes a user from the database by their ID.
	DeleteUser(id int) error

	// GetAllUsers retrieves all users from the database.
	GetAllUsers() ([]*models.User, error)
}

type userRepository struct {
	// db is the database connection or ORM instance.
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetUserByID(id int) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) DeleteUser(id int) error {
	return r.db.Delete(&models.User{}, id).Error
}

func (r *userRepository) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
