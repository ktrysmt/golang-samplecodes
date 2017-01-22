package main

type str struct {
	v string
}

func hitValue() (string, error) {
	c := str{"xxx"}
	return c.v, nil
}

func hitPointer() (string, error) {
	c := &str{"yyy"}
	return c.v, nil
}
