// main.go

package main

func main() {
	a := App{}
	// Set your username, password, and database name here
	a.Initialize("username", "password", "database")

	a.Run(":8080")
}
