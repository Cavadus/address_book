// main.go

package main

func main() {
	a := App{}
	a.Initialize("tester", "password", "pizza_hut")

	a.Run(":8080")
}
