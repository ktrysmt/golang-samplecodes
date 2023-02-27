package main

import (
	"fmt"
	"os"
	"strconv"

	"golang.org/x/xerrors"
)

func main() {
	if err := f(os.Args[1]); err != nil {
		fmt.Printf("%+v\n", xerrors.Errorf(": %w", err))
	}
}

func f(s string) error {
	n, err := strconv.Atoi(s)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	if n < 10 {
		return xerrors.Errorf("n is %d", n)
	}
	return nil
}
