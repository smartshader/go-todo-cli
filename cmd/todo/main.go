package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/smartshader/go-todo-cli"
)

const todoFileName = ".todo.json"

func errorExit(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func main() {
	l := &todo.List{}

	if err := l.Get(todoFileName); err != nil {
		errorExit(err)
	}

	switch {
	case len(os.Args) == 1:
		for _, item := range *l {
			fmt.Println(item.Task)
		}
	default:
		item := strings.Join(os.Args[1:], " ")
		l.Add(item)
		if err := l.Save(todoFileName); err != nil {
			errorExit(err)
		}
	}
}
