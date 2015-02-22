package main

import (
	"fmt"
	"github.com/"
)

//https://us.api.battle.net/d3/profile/brackets%231829/?locale=en_US&apikey=
func main() {
	m := martini.Classic()

	m.Get("/", func() string {
		return "hi"
	})

	fmt.Println("Hello")
	m.Run()
}
