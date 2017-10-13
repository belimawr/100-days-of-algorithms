package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Test_reversed(t *testing.T) {
	lst := commitsList{
		gitHubCommit{
			SHA: "0",
		},
		gitHubCommit{
			SHA: "1",
		},
		gitHubCommit{
			SHA: "2",
		},
		gitHubCommit{
			SHA: "3",
		},
	}

	rev := reversed(lst)

	n := len(lst)
	for i := 0; i < n; i++ {
		if lst[i].SHA != rev[n-1-i].SHA {
			t.Errorf("lst[%d] = %q, rev[%d] = %q",
				i,
				lst[i].SHA,
				i,
				rev[n-1-i].SHA)
		}
	}
}

func Test_countDaysWorked(t *testing.T) {
	lst := commitsList{
		gitHubCommit{
			SHA: "0",
			Commit: commit{
				Committer: commiter{
					Date: time.Date(2009,
						time.November,
						10, 23, 0, 0, 0, time.UTC),
				},
			},
		},
		gitHubCommit{
			SHA: "0",
			Commit: commit{
				Committer: commiter{
					Date: time.Date(2009,
						time.November,
						10, 23, 0, 0, 0, time.UTC),
				},
			},
		},
		gitHubCommit{
			SHA: "0",
			Commit: commit{
				Committer: commiter{
					Date: time.Date(2009,
						time.November,
						10, 23, 0, 0, 0, time.UTC),
				},
			},
		},
		gitHubCommit{
			SHA: "0",
			Commit: commit{
				Committer: commiter{
					Date: time.Date(2009,
						time.December,
						10, 23, 0, 0, 0, time.UTC),
				},
			},
		},
		gitHubCommit{
			SHA: "0",
			Commit: commit{
				Committer: commiter{
					Date: time.Date(2009,
						time.December,
						10, 23, 0, 0, 0, time.UTC),
				},
			},
		},
	}

	c := countDaysWorked(lst)
	expected := 2
	if c != 2 {
		t.Errorf("expecting %d, got %d", expected, c)

	}
}

func Test_getCommitList(t *testing.T) {
	lst := commitsList{
		gitHubCommit{
			SHA: "0",
		},
		gitHubCommit{
			SHA: "1",
		},
		gitHubCommit{
			SHA: "2",
		},
		gitHubCommit{
			SHA: "3",
		},
	}
	lstJSON, _ := json.Marshal(lst)
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			requestCount++

			baseURL := "https://" + r.Host + r.URL.String()
			mockLinkFmt := `<%s?page=2>; rel="next",<%s?page=3>; rel="last"`
			mockLinks := fmt.Sprintf(mockLinkFmt, baseURL, baseURL)

			w.Header().Set("Link", mockLinks)
			w.Write(lstJSON)
		}))
	defer server.Close()

	got, err := getCommitsList(server.URL, "user", "repo")

	if err != nil {
		t.Errorf("did not expect an error: %q", err.Error())
	}

	expected := len(lst) * 3
	gotLen := len(got)

	if gotLen != expected {
		t.Errorf("expecting %d, got %d", expected, gotLen)
	}
}

func Test_getCommitList_pagination_fetch_error(t *testing.T) {
	lst := commitsList{
		gitHubCommit{
			SHA: "0",
		},
		gitHubCommit{
			SHA: "1",
		},
		gitHubCommit{
			SHA: "2",
		},
		gitHubCommit{
			SHA: "3",
		},
	}
	lstJSON, _ := json.Marshal(lst)
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			requestCount++

			baseURL := "https://" + r.Host + r.URL.String()
			mockLinkFmt := `<%s?page=2>; rel="next",<%s?page=3>; rel="last"`
			mockLinks := fmt.Sprintf(mockLinkFmt, baseURL, baseURL)

			w.Header().Set("Link", mockLinks)

			if requestCount == 2 {
				w.WriteHeader(http.StatusTeapot)
			}
			w.Write(lstJSON)
		}))
	defer server.Close()

	got, err := getCommitsList(server.URL, "user", "repo")

	if err == nil {
		t.Error("expecting an error")
	}

	expected := len(lst)
	gotLen := len(got)

	if gotLen != expected {
		t.Errorf("expecting %d, got %d", expected, gotLen)
	}
}

func Test_getCommitList_invalid_URL(t *testing.T) {
	got, err := getCommitsList("ht\t://INVALID", "user", "repo")

	if err == nil {
		t.Error("expecting an error")
	}

	if len(got) != 0 {
		t.Error("commit list must be empty")
	}
}

func Test_getCommitList_invalid_statusCode(t *testing.T) {
	lst := commitsList{
		gitHubCommit{
			SHA: "0",
		},
		gitHubCommit{
			SHA: "1",
		},
		gitHubCommit{
			SHA: "2",
		},
		gitHubCommit{
			SHA: "3",
		},
	}
	lstJSON, _ := json.Marshal(lst)

	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusTeapot)
			w.Write(lstJSON)
		}))
	defer server.Close()

	got, err := getCommitsList(server.URL, "user", "repo")

	if err == nil {
		t.Error("expecting an error")
	}

	if len(got) != 0 {
		t.Error("commit list must be empty")
	}
}

func Test_getCommitList_invalid_JSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("[[ INVALID"))
		}))
	defer server.Close()

	got, err := getCommitsList(server.URL, "user", "repo")

	if err == nil {
		t.Error("expecting an error")
	}

	if len(got) != 0 {
		t.Error("commit list must be empty")
	}
}

func Test_elapsedDays(t *testing.T) {
	lst := commitsList{
		gitHubCommit{
			SHA: "0",
			Commit: commit{
				Committer: commiter{
					Date: time.Now().Add(-36 * time.Hour),
				},
			},
		},
	}

	e := elapsedDays(lst)
	expected := int64(1)
	if e != expected {
		t.Errorf("expecting %d, got %d", expected, e)
	}
}

func Test_elapsedDays_nil(t *testing.T) {
	e := elapsedDays(nil)

	expected := int64(-1)
	if e != expected {
		t.Errorf("expecting %d, got %d", expected, e)
	}
}

func Test_getPaginationLinks(t *testing.T) {
	expectedLen := 2
	request, _ := http.NewRequest("GET", "https://api.github.com", nil)
	response := &http.Response{}
	response.Request = request
	response.Header = http.Header(map[string][]string{})
	response.Header.Set("Link", `<https://api.github.com/repositories/65779870/commits?page=2>; rel="next", <https://api.github.com/repositories/65779870/commits?page=3>; rel="last"`)

	lst := getPaginationLinks(response)

	if len(lst) != expectedLen {
		t.Errorf("expecting len(lst) = %d, got %d\n%#v", len(lst), expectedLen, lst)
	}
}
