// main.go

package main

func main() {
	a := App{}
	// Set your username, password, and database name here
	a.Initialize("tester", "password", "pizza_hut")

	a.Run(":8080")
}
