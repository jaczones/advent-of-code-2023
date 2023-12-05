package main

import (
	"os"

	"github.com/jaczones/advent-of-code-2023/solutions"
)

func main() {
	//Get session cookie to get puzzle input
	sessionCookie := os.Getenv("SESSION_COOKIE")

	solutions.SolveDayOne(sessionCookie)

}
