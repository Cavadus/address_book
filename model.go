// model.go

package main

import (
	"database/sql"
	"errors"
)

type person struct {
	ID    int    `json:"id"`
	Fname string `json:"fname"`
	Lname string `json:"lname"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (p *person) getUser(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *person) updateUser(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *person) deleteUser(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *person) createUser(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *person) getUsers(db *sql.DB, start, count int) ([]person, error) {
	return nil, errors.New("Not implemented")
}
