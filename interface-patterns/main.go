package main

import (
	"fmt"
	"net/http"
)

type handlerFunc func(w http.ResponseWriter, r *http.Request)

type person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p *person) stringify() string {
	return fmt.Sprintf("%s %s (%d)", p.FirstName, p.LastName, p.Age)
}

func main() {
	p := &person{
		"kotaro",
		"yoshimatsu",
		31,
	}
	res := p.stringify()
	fmt.Sprintln(res)
}
