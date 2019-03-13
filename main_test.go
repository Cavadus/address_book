// main_test.go

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a = App{}
	a.Initialize("tester", "password", "pizza_hut")
	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM address_book")
	a.DB.Exec("ALTER TABLE address_book AUTO_INCREMENT = 1")
}

const tableCreationQuery = `
CREATE TABLE IF NOT EXISTS address_book
(
	id INT AUTO_INCREMENT PRIMARY KEY,
	fname VARCHAR(255) NOT NULL,
	lname VARCHAR(255) NOT NULL,
	email VARCHAR(255),
	phone VARCHAR(255)
)`

func TestEmptyTable(t *testing.T) {
	clearTable()
	req, _ := http.NewRequest("GET", "/address_book", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array.  Got %s", body)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d.  Got %d\n", expected, actual)
	}
}

func TestGetNonExistentPerson(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/address_book/45", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "User not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'User not found'.  Got '%s'", m["error"])
	}
}

func TestCreateUser(t *testing.T) {
	clearTable()

	payload := []byte(`{"fname":"Dave", "lname":"Coulier", "email":"dcoulier@stepbystep.biz", "phone":"1235556789"}`)

	req, _ := http.NewRequest("POST", "/address_book", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["fname"] != "Dave" {
		t.Errorf("Expected first name to be 'Dave'.  Got '%v'", m["fname"])
	}

	if m["lname"] != "Coulier" {
		t.Errorf("Expected last name to be 'Coulier'.  Got '%v'", m["lname"])
	}

	if m["email"] != "dcoulier@stepbystep.biz" {
		t.Errorf("Expected email to be 'dcoulier@stepbystep.biz'.  Got '%v'", m["email"])
	}

	if m["phone"] != "1235556789" {
		t.Errorf("Expected phone number to be '1235556789'.  Got '%v'", m["phone"])
	}

	if m["id"] != 1.0 {
		t.Errorf("Expected person's ID to be '1'.  Got '%v'", m["id"])
	}
}

func TestGetUser(t *testing.T) {
	clearTable()
	addUsers(1)

	req, _ := http.NewRequest("GET", "/address_book/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func addUsers(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		statement := fmt.Sprintf("INSERT INTO address_book(fname, lname, email, phone) VALUES ('%s', '%s', '%s', '%s')", ("Fname" + strconv.Itoa(i+1)), ("Lname" + strconv.Itoa(i+1)), ("email" + strconv.Itoa(i+1) + "@email.com"), ("Phone" + strconv.Itoa(i+1)))
		a.DB.Exec(statement)
	}
}

func TestUpdateUser(t *testing.T) {
	clearTable()
	addUsers(1)

	req, _ := http.NewRequest("GET", "/address_book/1", nil)
	response := executeRequest(req)

	var originalUser map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalUser)

	payload := []byte(`{"fname":"updated name", "phone":"9999998888"}`)

	req, _ = http.NewRequest("PUT", "/address_book/1", bytes.NewBuffer(payload))
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["id"] != originalUser["id"] {
		t.Errorf("Expected the id to remain the same (%v).  Got %v", originalUser["id"], m["id"])
	}

	if m["fname"] == originalUser["fname"] {
		t.Errorf("Expected the name to change from '%v' to '%v'.  Got '%v'", originalUser["fname"], m["fname"], m["fname"])
	}

	if m["phone"] == originalUser["phone"] {
		t.Errorf("Expected the phone number to change from '%v' to '%v'.  Got '%v'", originalUser["phone"], m["phone"], m["phone"])
	}
}

func TestDeleteUser(t *testing.T) {
	clearTable()
	addUsers(1)

	req, _ := http.NewRequest("GET", "/address_book/1", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("DELETE", "/address_book/1", nil)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("GET", "/address_book/1", nil)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}
