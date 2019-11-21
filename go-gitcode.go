package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type RepoFound struct {
	TotalCount        int  `json:"total_count"`
	IncompleteResults bool `json:"incomplete_results"`
	Items             []struct {
		Name       string `json:"name"`
		Path       string `json:"path"`
		Sha        string `json:"sha"`
		URL        string `json:"url"`
		GitURL     string `json:"git_url"`
		HTMLURL    string `json:"html_url"`
		Repository struct {
			ID       int    `json:"id"`
			NodeID   string `json:"node_id"`
			Name     string `json:"name"`
			FullName string `json:"full_name"`
			Private  bool   `json:"private"`
			Owner    struct {
				Login             string `json:"login"`
				ID                int    `json:"id"`
				NodeID            string `json:"node_id"`
				AvatarURL         string `json:"avatar_url"`
				GravatarID        string `json:"gravatar_id"`
				URL               string `json:"url"`
				HTMLURL           string `json:"html_url"`
				FollowersURL      string `json:"followers_url"`
				FollowingURL      string `json:"following_url"`
				GistsURL          string `json:"gists_url"`
				StarredURL        string `json:"starred_url"`
				SubscriptionsURL  string `json:"subscriptions_url"`
				OrganizationsURL  string `json:"organizations_url"`
				ReposURL          string `json:"repos_url"`
				EventsURL         string `json:"events_url"`
				ReceivedEventsURL string `json:"received_events_url"`
				Type              string `json:"type"`
				SiteAdmin         bool   `json:"site_admin"`
			} `json:"owner"`
			HTMLURL          string      `json:"html_url"`
			Description      interface{} `json:"description"`
			Fork             bool        `json:"fork"`
			URL              string      `json:"url"`
			ForksURL         string      `json:"forks_url"`
			KeysURL          string      `json:"keys_url"`
			CollaboratorsURL string      `json:"collaborators_url"`
			TeamsURL         string      `json:"teams_url"`
			HooksURL         string      `json:"hooks_url"`
			IssueEventsURL   string      `json:"issue_events_url"`
			EventsURL        string      `json:"events_url"`
			AssigneesURL     string      `json:"assignees_url"`
			BranchesURL      string      `json:"branches_url"`
			TagsURL          string      `json:"tags_url"`
			BlobsURL         string      `json:"blobs_url"`
			GitTagsURL       string      `json:"git_tags_url"`
			GitRefsURL       string      `json:"git_refs_url"`
			TreesURL         string      `json:"trees_url"`
			StatusesURL      string      `json:"statuses_url"`
			LanguagesURL     string      `json:"languages_url"`
			StargazersURL    string      `json:"stargazers_url"`
			ContributorsURL  string      `json:"contributors_url"`
			SubscribersURL   string      `json:"subscribers_url"`
			SubscriptionURL  string      `json:"subscription_url"`
			CommitsURL       string      `json:"commits_url"`
			GitCommitsURL    string      `json:"git_commits_url"`
			CommentsURL      string      `json:"comments_url"`
			IssueCommentURL  string      `json:"issue_comment_url"`
			ContentsURL      string      `json:"contents_url"`
			CompareURL       string      `json:"compare_url"`
			MergesURL        string      `json:"merges_url"`
			ArchiveURL       string      `json:"archive_url"`
			DownloadsURL     string      `json:"downloads_url"`
			IssuesURL        string      `json:"issues_url"`
			PullsURL         string      `json:"pulls_url"`
			MilestonesURL    string      `json:"milestones_url"`
			NotificationsURL string      `json:"notifications_url"`
			LabelsURL        string      `json:"labels_url"`
			ReleasesURL      string      `json:"releases_url"`
			DeploymentsURL   string      `json:"deployments_url"`
		} `json:"repository"`
		Score       float64 `json:"score"`
		TextMatches []struct {
			ObjectURL  string `json:"object_url"`
			ObjectType string `json:"object_type"`
			Property   string `json:"property"`
			Fragment   string `json:"fragment"`
			Matches    []struct {
				Text    string `json:"text"`
				Indices []int  `json:"indices"`
			} `json:"matches"`
		} `json:"text_matches"`
	} `json:"items"`
}

var (
	url = "https://api.github.com/search/code?q="
)

func SearchCode(query string, key string, token string, orderby string, sort string, debug bool) (RepoFound, error) {
	URI := url + query + "%20" + key + "&sort=" + orderby + "&o=" + sort
	req, err := http.NewRequest("GET", URI, nil)
	key = "token " + token + ""
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	req.Header.Set("Authorization", key)

	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
		errors.New("Rate limit")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}

	data := RepoFound{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, errors.New("Error when marshall")
	}

	if resp.StatusCode != http.StatusOK {
		// This time to change your token
		return data, errors.New("Rate limit")
	}
	if debug {
		PrintData(data)
	}

	return data, err

}

func PrintData(data RepoFound) {

	for _, i := range data.Items {
		blob := i.HTMLURL
		owner := i.Repository.Owner.Login
		fmt.Println("\nBlob :", blob)
		fmt.Println("Owner :", owner)
	}

}
