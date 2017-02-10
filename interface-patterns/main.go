package main

import (
	"fmt"
	"net/http"
)

// 1. Define original typename for any data types that func, slice, map, other.
// Example name type for functionType.
type handlerFunc func(w http.ResponseWriter, r *http.Request)

// 2. Embed the Name() function in person struct.
type person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p *person) Name() string {
	return fmt.Sprintf("%s %s (%d)", p.FirstName, p.LastName, p.Age)
}

// 3. partial implement of interface.
// An error occurs if the return value does not satisfy the expected interface 'human'.
type human interface {
	Title() string
	Name() string
}

type gender int

const (
	fe = iota
	ma
)

type female struct {
	*person
}

type male struct {
	*person
}

func (f *female) Title() string {
	return "Ms."
}

func (m *male) Title() string {
	return "Mr."
}

func newPerson(g gender, FirstName, LastName string, Age int) human {
	p := &person{FirstName, LastName, Age}

	if g == fe {
		return &female{p}
	}

	return &male{p}
}

// 4. Dynamic embedding of interface.
// パッケージ外に公開されない型として宣言
type innerStringer fmt.Stringer

type humanA struct {
	innerStringer // fmt.Stringer型でありながらパッケージ外から代入できない匿名フィールド
	firstname     string
	lastname      string
	age           int
}

type stringerFunc func() string

func (sf stringerFunc) String() string {
	return sf()
}

func bindStringer(h *humanA, f func(h *humanA) string) fmt.Stringer {
	return stringerFunc(func() string {
		return f(h)
	})
}

func newHumanA(firstname, lastname string, age int) (h *humanA) { // 返り値の型だけでなく返り値としたい実変数名も指定できる
	h = &humanA{
		nil,
		firstname,
		lastname,
		age,
	}

	h.innerStringer = stringerFunc(func() string {
		return fmt.Sprintf("%s %s (%d)", h.firstname, h.lastname, h.age)
	})

	return
}

// 引数のエイリアスとそれに対応する方だけでなく関数型ならその関数型の引数と返り値も含めて引数として事前に定義できる
func (h *humanA) setStringer(sf func(h *humanA) string) {
	h.innerStringer = stringerFunc(func() string {
		return sf(h)
	})
}

func main() {
	// 2.
	p := &person{
		"kotaro",
		"yoshimatsu",
		31,
	}
	r := p.Name()
	fmt.Println(r)

	// 3.
	h := newPerson(0, "Kotaro", "Yoshimatsu", 32)
	fmt.Printf("Hello, %s %s\n", h.Title(), h.Name())

	// 4.
	ha := newHumanA("KOTARO", "YOSHIMATSU", 33)
	fmt.Printf("I am %s\n", ha.innerStringer) // 関数の実行ではなく関数の実行結果を返すプロパティフィールドなので()は要らない
}
