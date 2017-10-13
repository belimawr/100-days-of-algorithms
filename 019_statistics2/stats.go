package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func countDaysWorked(commits commitsList) int64 {
	d := 24 * time.Hour
	counter := map[time.Time][]string{}

	for _, c := range commits {
		tmp := c.Commit.Committer.Date.Truncate(d)
		counter[tmp] = append(counter[tmp], c.Commit.Message)
	}

	return int64(len(counter))
}

func getCommitsList(baseURL, user, repo string) (commitsList, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/commits", baseURL, user, repo)

	lst, response, err := fetchCommits(url)

	if err != nil {
		return commitsList{}, err
	}

	urls := getPaginationLinks(response)

	for _, u := range urls {
		l, _, err := fetchCommits(u)
		if err != nil {
			return lst, err
		}
		lst = append(lst, l...)
	}

	return lst, nil
}

func fetchCommits(url string) (commitsList, *http.Response, error) {
	response, err := http.Get(url)

	if err != nil {
		return commitsList{}, response, err
	}

	if response.StatusCode != http.StatusOK {
		return commitsList{}, response, fmt.Errorf("unexpected response code: %d", response.StatusCode)
	}

	if response.Body != nil {
		defer response.Body.Close()
	}

	lst := commitsList{}

	if err := json.NewDecoder(response.Body).Decode(&lst); err != nil {
		return commitsList{}, response, err
	}

	return reversed(lst), response, nil
}

func getPaginationLinks(response *http.Response) []string {
	rawLinks := response.Header.Get("Link")

	var re = regexp.MustCompile(`<(?P<url>https:\/\/([a-zA-Z0-9\/.:])*)\?page=(?P<page>\d+)>;\srel=\"(?P<rel>\w+)\"`)
	var next, last int
	for _, link := range strings.Split(rawLinks, ",") {
		for _, match := range re.FindAllStringSubmatch(strings.TrimSpace(link), -1) {
			switch match[4] {
			case "next":
				next, _ = strconv.Atoi(match[3])
			case "last":
				last, _ = strconv.Atoi(match[3])
			}
		}
	}

	baseURL := response.Request.URL
	urls := []string{}
	for i := next; i <= last; i++ {
		url := fmt.Sprintf("%s?page=%d", baseURL, i)
		urls = append(urls, url)
	}

	return urls
}

func elapsedDays(lst commitsList) int64 {
	if lst == nil {
		return -1
	}

	day := 24 * time.Hour
	firstDay := lst[0].Commit.Committer.Date.Truncate(day)

	today := time.Now().Truncate(day)

	diff := today.Unix() - firstDay.Unix()

	return diff / (60 * 60 * 24)
}

func reversed(s commitsList) commitsList {
	cp := append([]gitHubCommit(nil), s...)

	for left, right := 0, len(cp)-1; left < right; left, right = left+1, right-1 {
		cp[left], cp[right] = cp[right], cp[left]
	}

	return cp
}
