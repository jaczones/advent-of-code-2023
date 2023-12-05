package solutions

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/jaczones/advent-of-code-2023/util"
)

func SolveDayOne(sessionCookie string) {
	url := "https://adventofcode.com/2023/day/1/input"

	content, err := util.FetchWebPage(url, sessionCookie)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	lines := strings.Split(content, "\n")

	digitRegex := regexp.MustCompile(`\d`)

	var total int

	for _, line := range lines {
		if line == "" {
			continue
		}

		digits := digitRegex.FindAllString(line, -1)

		if len(digits) >= 1 {
			firstDigit, err := strconv.Atoi(digits[0])

			if err != nil {
				fmt.Println("Error converting first digit:", err)
				continue
			}

			lastDigit, err := strconv.Atoi(digits[len(digits)-1])
			if err != nil {
				fmt.Println("Error converting last digit:", err)
				continue
			}

			fmt.Println("\nLine", line, "first", firstDigit, "last", lastDigit)
			twoDigit := firstDigit*10 + lastDigit
			total += twoDigit

		}
	}

	fmt.Println(total)
}
