package main

import (
	"flag"
	"fmt"
	"os"
)

var apiURL = "https://api.github.com"

func main() {
	user := flag.String("user", "belimawr", "GutHub username")
	repo := flag.String("repo", "100-days-of-algorithms", "Repository")
	maxDays := flag.Int64("maxDays", 150, "Maximum number of days")

	flag.Parse()

	lst, err := getCommitList(apiURL, *user, *repo)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	printStats(lst, *maxDays)
}

func printStats(lst commitsList, maxDays int64) {
	daysWorked := countDaysWorked(lst)
	elapsed := elapsedDays(lst)
	remaining := 100 - daysWorked
	freeDays := maxDays - elapsed - remaining

	fmt.Println("Days worked:", daysWorked)
	fmt.Println("Days since the beginning:", elapsed)
	fmt.Println("Remaining days to 100 days of code:", remaining)
	fmt.Println("Remaining days without code:", freeDays)
}
