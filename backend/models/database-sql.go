package models

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/arthurkushman/buildsqlx"
	_ "github.com/lib/pq"
)

type DatabaseSQL struct {
	db *buildsqlx.DB
}

func newDatabaseSQL(driverName string, dataSourceName string) *DatabaseSQL {
	db := buildsqlx.NewDb(buildsqlx.NewConnection(driverName, dataSourceName))
	// Check connection
	err := db.Sql().Ping()
	if err != nil {
		log.Println(driverName, dataSourceName)
		log.Fatalf("Database connection failed -> %v", err)
	}
	return &DatabaseSQL{db: db}
}

// GetUser returns user by email, if exists
func (m *DatabaseSQL) GetUser(email string) (*User, error) {
	email = strings.ToLower(email)

	res, err := m.db.Table("users").Select("id", "password").Where("email", "=", email).Get()
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, errors.New("user not found")
	}
	u := res[0]
	id := int(u["id"].(int64)) // la DB devuelve string, parseable a int64 y luego lo convertimos a int
	password := fmt.Sprint(u["password"])

	user := User{
		ID:       id,
		Email:    email,
		Password: password,
	}

	return &user, nil
}

// CreateUser inserts new user into the database, returns error if error
func (m *DatabaseSQL) CreateUser(email string, password string) error {
	email = strings.ToLower(email)

	_, err := m.GetUser(email)
	if err == nil {
		return errors.New("email is already taken")
	}

	err = m.db.Table("users").Insert(map[string]interface{}{"email": email, "password": password, "created_on": "NOW()"})
	if err != nil {
		return err
	}
	return nil
}
