package psql

import (
	"database/sql"
	"fmt"

	"github.com/vbetsun/todo-app/internal/core"
)

// Auth represents repository for authorization and authentication
type Auth struct {
	db *sql.DB
}

// NewAuth return instance of auth repository
func NewAuth(db *sql.DB) *Auth {
	return &Auth{db}
}

// CreateUser creates new user in DB
func (r *Auth) CreateUser(u core.User) (core.User, error) {
	var user core.User
	err := r.db.QueryRow(createUserQuery(), u.Name, u.Username, u.Password).
		Scan(&user.ID, &user.Name, &user.Username)
	return user, err
}

// GetUser returns user from DB by username and password
func (r *Auth) GetUser(username, pwd string) (core.User, error) {
	var user core.User
	err := r.db.QueryRow(getUserQuery(), username, pwd).Scan(&user.ID)

	return user, err
}

func createUserQuery() string {
	return fmt.Sprintf(`--sql
		INSERT INTO %s (name, username, password_hash) 
		VALUES ($1, $2, $3) 
		RETURNING id, name, username
	`, usersTable)
}

func getUserQuery() string {
	return fmt.Sprintf(`--sql
		SELECT id FROM %s 
		WHERE username = $1 
		AND password_hash = $2
	`, usersTable)
}
