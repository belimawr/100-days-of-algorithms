package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func getCommitList(baseURL, user, repo string) (commitsList, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/commits", baseURL, user, repo)

	response, err := http.Get(url)

	if err != nil {
		return commitsList{}, err
	}

	if response.StatusCode != http.StatusOK {
		return commitsList{}, fmt.Errorf("unexpected response code: %d", response.StatusCode)
	}

	if response.Body != nil {
		defer response.Body.Close()
	}

	lst := commitsList{}

	if err := json.NewDecoder(response.Body).Decode(&lst); err != nil {
		return commitsList{}, err
	}

	return reversed(lst), nil
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
