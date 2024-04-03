package repository

import (
	"database/sql"
	"errors"
	"go-blog-app/internal/domain"
	"log"
)

type UserRepository interface {
	CreateUser(user domain.User) (domain.User, error)
	FindUser(email string) (domain.User, error)
	FindUserById(id uint) (domain.User, error)
	UpdateUser(id uint, u domain.User) (domain.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r userRepository) CreateUser(user domain.User) (domain.User, error) {
	query := `INSERT INTO users (first_name, last_name, email, password, role) VALUES ($1, $2, $3, $4, $5) returning id`
	err := r.db.QueryRow(query, user.FirstName, user.LastName, user.Email, user.Password, user.Role).Scan(&user.ID)
	if err != nil {
		log.Printf("create user error %v", err)
		return domain.User{}, err
	}

	return user, nil
}

func (r userRepository) FindUser(email string) (domain.User, error) {
	var user domain.User
	query := `SELECT id, first_name, last_name, email, role, password FROM users WHERE email = $1`
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Role, &user.Password)
	if err != nil {
		log.Printf("find user error %v", err)
		return domain.User{}, errors.New("user does not exist")
	}
	return user, nil
}

func (r userRepository) FindUserById(id uint) (domain.User, error) {
	var user domain.User
	query := `SELECT id, first_name, last_name, email, role FROM users WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Role)
	if err != nil {
		log.Printf("find user error %v", err)
		return domain.User{}, errors.New("user does not exist")
	}
	return user, nil
}

func (r userRepository) UpdateUser(id uint, u domain.User) (domain.User, error) {
	var user domain.User
	query := `UPDATE users SET first_name = $1, last_name = $2, email = $3, password = $4 WHERE id = $5`
	_, err := r.db.Exec(query, u.FirstName, u.LastName, u.Email, u.Password, id)
	if err != nil {
		log.Printf("error on update %v", err)
		return domain.User{}, errors.New("failed update user")
	}
	return user, nil
}
