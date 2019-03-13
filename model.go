// model.go

package main

import (
	"database/sql"
	"fmt"
)

type Person struct {
	ID    int    `json:"id"`
	Fname string `json:"fname"`
	Lname string `json:"lname"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (p *Person) getUser(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT fname, lname, email, phone FROM address_book WHERE id=%d", p.ID)
	return db.QueryRow(statement).Scan(&p.Fname, &p.Lname, &p.Email, &p.Phone)
}

func (p *Person) updateUser(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE address_book SET fname='%s', lname='%s', email='%s', phone='%s' WHERE id=%d", p.Fname, p.Lname, p.Email, p.Phone, p.ID)
	_, err := db.Exec(statement)
	return err
}

func (p *Person) deleteUser(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM address_book WHERE id=%d", p.ID)
	_, err := db.Exec(statement)
	return err
}

func (p *Person) createUser(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO address_book(fname, lname, email, phone) VALUES('%s', '%s', '%s', '%s')", p.Fname, p.Lname, p.Email, p.Phone)
	_, err := db.Exec(statement)

	if err != nil {
		return err
	}

	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&p.ID)

	if err != nil {
		return err
	}

	return nil
}

func (p *Person) getUsers(db *sql.DB, start, count int) ([]Person, error) {
	statement := fmt.Sprintf("SELECT ID, fname, lname, email, phone FROM address_book LIMIT %d OFFSET %d", count, start)
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	people := []Person{}

	for rows.Next() {
		var p Person
		if err := rows.Scan(&p.ID, &p.Fname, &p.Lname, &p.Email, &p.Phone); err != nil {
			return nil, err
		}
		people = append(people, p)
	}

	return people, nil
}

func (p *Person) importCSV(db *sql.DB) {

}

func (p *Person) exportCSV(db *sql.DB) {

}
