package solutions

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/jaczones/advent-of-code-2023/util"
)

func DayOne(input string) uint32 {
	text := input
	textArray := []rune(text)
	numWords := [...]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for i := 0; i < 10; i++ {
		loc := strings.Index(text, numWords[i])
		for loc != -1 {
			textArray[loc+1] = []rune(strconv.Itoa(i))[0]
			text = string(textArray)
			loc = strings.Index(text, numWords[i])
		}
	}

	full := strings.Split(text, "\n")

	var total uint32
	for _, input := range full {
		nums := getDigitsFromString(input)
		if len(nums) > 0 {
			total += nums[0]*10 + nums[len(nums)-1]
		}
	}

	return total
}

func getDigitsFromString(input string) []uint32 {
	var nums []uint32
	digitsRegex := regexp.MustCompile(`\d`)
	digitRunes := digitsRegex.FindAll([]byte(input), -1)

	for _, r := range digitRunes {
		num, _ := strconv.Atoi(string(r))
		nums = append(nums, uint32(num))
	}

	return nums
}

func SolveDayOne(sessionCookie string) {
	url := "https://adventofcode.com/2023/day/1/input"

	content, err := util.FetchWebPage(url, sessionCookie)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	result := DayOne(content)
	fmt.Println("Total:", result)
}
