package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func Day1() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalCals := make([]int, 0)
	cals := 0

	for scanner.Scan() {
		if scanner.Text() != "" {
			i, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			cals += i
		} else if scanner.Text() == "" {
			totalCals = append(totalCals, cals)
			cals = 0
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	sort.Ints(totalCals)
	fmt.Printf("main mostCals: %v\n", totalCals[len(totalCals)-1]) // __AUTO_GENERATED_PRINT_VAR__
	fmt.Printf("main mostCals: %v\n", totalCals[len(totalCals)-2]) // __AUTO_GENERATED_PRINT_VAR__
	fmt.Printf("main mostCals: %v\n", totalCals[len(totalCals)-3]) // __AUTO_GENERATED_PRINT_VAR__

	total := totalCals[len(totalCals)-1] + totalCals[len(totalCals)-2] + totalCals[len(totalCals)-3]
	fmt.Printf("main total: %v\n", total) // __AUTO_GENERATED_PRINT_VAR__
}
