package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/docopt/docopt-go"
	libgithub "github.com/google/go-github/github"
)

const (
	usage = `GitHub Watchers
Usage:
    github-watchers <username> [<repository>...]
`
)

type Watcher struct {
	Login string
	URL   string
}

func main() {
	args, err := docopt.Parse(usage, nil, true, "1.0", false, true)
	if err != nil {
		log.Fatal(err)
	}

	var (
		username     = args["<username>"].(string)
		repositories = args["<repository>"].([]string)
	)

	github := libgithub.NewClient(nil)

	if len(repositories) == 0 {
		repositories, err = getRepositories(github, username)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, repository := range repositories {
		fmt.Printf(
			"%s:\n",
			"https://github.com/"+username+"/"+repository,
		)

		watchers, err := getWatchers(github, username, repository)
		if err != nil {
			log.Fatal(err)
		}

		columnLength := getMaximumLoginLength(watchers) + 2

		for _, watcher := range watchers {
			fmt.Printf(
				"%-"+strconv.Itoa(columnLength)+"s %s\n",
				watcher.Login, watcher.URL,
			)
		}

		fmt.Println()
	}

}

func getWatchers(
	github *libgithub.Client, username string, repository string,
) ([]Watcher, error) {
	users, _, err := github.Activity.ListWatchers(
		username, repository, nil,
	)
	if err != nil {
		return []Watcher{}, err
	}

	watchers := []Watcher{}
	for _, user := range users {
		watchers = append(watchers, Watcher{
			Login: string(*user.Login),
			URL:   string(*user.HTMLURL),
		})
	}

	return watchers, nil
}

func getRepositories(
	github *libgithub.Client, username string,
) ([]string, error) {
	repositories, _, err := github.Repositories.List(
		username,
		&libgithub.RepositoryListOptions{
			Type:        "owner",
			Sort:        "updated",
			Direction:   "desc",
			ListOptions: libgithub.ListOptions{PerPage: 99999},
		},
	)
	if err != nil {
		return []string{}, err
	}

	names := []string{}
	for _, repository := range repositories {
		names = append(names, *repository.Name)
	}

	return names, nil
}

func getMaximumLoginLength(watchers []Watcher) int {
	maxLength := 0
	for _, watcher := range watchers {
		if len(watcher.Login) > maxLength {
			maxLength = len(watcher.Login)
		}
	}

	return maxLength
}
