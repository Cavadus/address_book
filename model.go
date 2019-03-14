// model.go

package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
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

func (p *Person) getUsers(db *sql.DB, start int, count int) ([]Person, error) {
	//func (p *Person) getUsers(db *sql.DB) ([]Person, error) {
	statement := fmt.Sprintf("SELECT ID, fname, lname, email, phone FROM address_book LIMIT %d OFFSET %d", count, start)
	//statement := fmt.Sprintf("SELECT * FROM address_book")
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

func (p *Person) importCSV(db *sql.DB) error {
	filePath := "/tmp/import.csv"
	mysql.RegisterLocalFile(filePath)
	statement := fmt.Sprintf("LOAD DATA LOCAL INFILE '" + filePath + "' INTO TABLE pizza_hut FIELDS TERMINATED BY ',' LINES TERMINATED BY '\n' IGNORE 1 ROWS (id, fname, lname, email, phone)")
	_, err := db.Exec(statement)
	if err != nil {
		return err
	}
	return nil
}

func (p *Person) exportCSV(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT * FROM address_book")
	rows, err := db.Query(statement)
	if err != nil {
		return err
	}

	colNames, err := rows.Columns()
	if err != nil {
		return err
	}

	file, err := os.Create("exported.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = '\t'
	readCols := make([]interface{}, len(colNames))
	writeCols := make([]string, len(colNames))

	for i := range writeCols {
		readCols[i] = &writeCols[i]
	}

	for rows.Next() {
		err := rows.Scan(readCols...)
		if err != nil {
			return err
		}
		writer.Write(writeCols)
	}

	if err = rows.Err(); err != nil {
		return err
	}
	writer.Flush()
	return nil
}
