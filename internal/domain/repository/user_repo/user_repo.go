package user_repo

import (
	"errors"
	"time"
	"url_shortening/infra/config/environment"
	"url_shortening/infra/db/postgres"

	"github.com/google/uuid"
)

type User struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepository struct {
	db     *postgres.Postgres
	config *environment.Config
}

func NewUserRepository(db *postgres.Postgres, config *environment.Config) *UserRepository {
	return &UserRepository{db: db, config: config}
}

func (r *UserRepository) RegisterUser(user *User) (string, string, error) {

	uniqueID, err := uuid.NewV7()
	if err != nil {
		return "", "", err
	}

	var count int
	err = r.db.Db.Raw(`SELECT COUNT(*) FROM users WHERE email = ?`, user.Email).Row().Scan(&count)
	if err != nil {
		return "", "", err
	}

	if count > 0 {
		return "", "", errors.New("user already exists")
	}

	query := `INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4)`
	response, err := r.db.Db.Raw(query, uniqueID, user.Name, user.Email, user.Password).Rows()
	if err != nil {
		return "", "", err
	}

	defer response.Close()

	return uniqueID.String(), user.Email, nil
}

func (r *UserRepository) GetUserByEmail(email string) (User, error) {

	var user User
	err := r.db.Db.Raw(`SELECT * FROM users WHERE email = $1`, email).Row().Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
