package table

import (
	"crypto/sha1"
	"database/sql"
	"encoding/base64"
	"github.com/chessnok/GoCalculator/orchestrator/internal/user"
	"github.com/google/uuid"
)

type Users struct {
	db *sql.DB
}

func generateID() string {
	return uuid.New().String()
}

func hashPassword(password string) string {
	hasher := sha1.New()
	hasher.Write([]byte(password))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}
func (u *Users) NewUser(username, password string) (*user.User, error) {
	id := generateID()
	password = hashPassword(password)
	usr := &user.User{
		ID:       id,
		Username: username,
		Password: password,
	}
	_, err := u.db.Exec("INSERT INTO users (id, username, password) VALUES ($1, $2, $3)", id, username, password)
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func (u *Users) SetIsAdmin(username string, isAdmin bool) error {
	_, err := u.db.Exec("UPDATE users SET is_admin = $1 WHERE username = $2", isAdmin, username)
	if err != nil {
		return err
	}
	return nil
}

func NewUsers(db *sql.DB) *Users {
	return &Users{
		db: db,
	}
}

func (u *Users) GetUserByUsername(username string) (*user.User, error) {
	row := u.db.QueryRow("SELECT id, password, is_admin FROM users WHERE username = $1", username)
	usr := &user.User{Username: username}
	err := row.Scan(&usr.ID, &usr.Password, &usr.IsAdmin)
	if err != nil {
		return nil, err
	}
	return usr, nil
}
func (u *Users) GetUserById(id string) (*user.User, error) {
	row := u.db.QueryRow("SELECT username, password, is_admin FROM users WHERE id = $1", id)
	usr := &user.User{ID: id}
	err := row.Scan(&usr.Username, &usr.Password, &usr.IsAdmin)
	if err != nil {
		return nil, err
	}
	return usr, nil
}
func ComparePasswords(hashedPassword, password string) bool {
	return hashedPassword == hashPassword(password)
}
