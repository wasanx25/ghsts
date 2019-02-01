package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/github"
)

func main() {
	client := github.NewClient(nil)
	opt := &github.RepositoryListByOrgOptions{Type: "public"}
	repos, _, err := client.Repositories.ListByOrg(context.Background(), "github", opt)
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range repos {
		fmt.Println(r.GetFullName())
	}
}
