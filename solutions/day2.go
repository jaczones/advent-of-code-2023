package solutions

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/jaczones/advent-of-code-2023/util"
)

func SolveDayTwo(sessionCookie string) {
	url := "https://adventofcode.com/2023/day/2/input"

	content, err := util.FetchWebPage(url, sessionCookie)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	games := strings.Split(content, "\n")

	re := regexp.MustCompile(`(\d+)\s+(\w+)`)

	var totalSum int

	for _, game := range games {
		// Skip empty strings
		if game == "" {
			continue
		}

		// Find all matches in the game text
		matches := re.FindAllStringSubmatch(game, -1)

		colorCounts := make(map[string]int)

		for _, match := range matches {
			// Extract count and color
			count, err := strconv.Atoi(match[1])
			if err != nil {
				fmt.Println("Error converting count to integer:", err)
				continue
			}
			color := match[2]

			// Track the highest count for each color
			if count > colorCounts[color] {
				colorCounts[color] = count
			}
		}

		// Multiply the highest counts for each color and add to the total sum
		product := 1
		for _, count := range colorCounts {
			product *= count
		}
		totalSum += product
	}

	// Print the sum of products for all games
	fmt.Printf("Sum of Products for All Games: %d\n", totalSum)
}
