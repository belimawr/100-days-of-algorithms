package main

import "time"

type commitsList []gitHubCommit

type gitHubCommit struct {
	SHA    string `json:"sha"`
	Commit commit `json:"commit"`
}

type commit struct {
	Author    commiter `json:"author"`
	Committer commiter `json:"committer"`
	Message   string   `json:"message"`
}

type commiter struct {
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Date  time.Time `json:"date"`
}
