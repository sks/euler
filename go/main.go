package main

import (
	"fmt"
	"os"

	"github.com/sks/euler/go/problem1"
	"github.com/sks/euler/go/problem2"
	"github.com/sks/euler/go/problem4"
)

type Question interface {
	Question()
}

var problems = make(map[string]Question)

func init() {
	problems["problem4"] = problem4.Problem{}
	problems["problem1"] = problem1.Problem{}
	problems["problem2"] = problem2.Problem{}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage %[1]s problem<NUMBER>\n eg: %[1]s problem4\n", os.Args[0])
		os.Exit(1)
	}
	var problem = os.Args[1]

	if question, ok := problems[problem]; !ok {
		fmt.Printf("Could not find the problem %s\n", problem)
		os.Exit(2)
	} else {
		question.Question()
	}
}
