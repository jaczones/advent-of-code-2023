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

	maximumCounts := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	noExceedMap := make(map[int]bool)

	games := strings.Split(content, "\n")

	re := regexp.MustCompile(`(\d+)\s+(\w+)`)

gameLoop:
	for i, game := range games {
		// Skip empty strings
		if game == "" {
			continue
		}

		// Find all matches in the game text
		matches := re.FindAllStringSubmatch(game, -1)

		for _, match := range matches {
			// Extract count and color
			count, err := strconv.Atoi(match[1])
			if err != nil {
				fmt.Println("Error converting count to integer:", err)
				continue
			}
			color := match[2]

			// Check if the count exceeds the maximum for the specific color
			if count > maximumCounts[color] {
				// Disqualify the entire game
				continue gameLoop
			}
		}

		// If no iteration exceeds the maximum, mark the game number as eligible
		noExceedMap[i+1] = true
	}

	// Calculate the sum of eligible game numbers
	var sum int
	for gameNumber := range noExceedMap {
		sum += gameNumber
	}

	// Print the sum of game numbers with no exceedance
	fmt.Printf("Sum of Game Numbers with No Exceedance: %d\n", sum)
}
