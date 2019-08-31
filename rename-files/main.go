// Package main provides ...
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	l, _ := ioutil.ReadDir("./")
	for _, v := range l {
		name := v.Name()
		if name == "main.go" {
			continue
		} else if strings.Contains(name, "01m") || strings.Contains(name, "02m") || strings.Contains(name, "03m") {
			continue
		}

		// fmt.Println(name)
		mw := strings.Split(name, "m_")
		w := strings.Split(mw[1], "wk")
		// fmt.Println(mw[0], w[0])
		month, _ := strconv.Atoi(mw[0])
		week, _ := strconv.Atoi(w[0])
		newWeek := (month-1)*4 + week
		newName := fmt.Sprintf("%02dm_%02dwk.md", month, newWeek)

		// fmt.Println(month, week, newWeek, name, newName)
		os.Rename(name, newName)

	}
}
