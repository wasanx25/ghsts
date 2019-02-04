package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ghodss/yaml"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Setting struct {
	Owner             string             `json:"owner"`
	Repos             []string           `json:"repos"`
	BranchProtections []BranchProtection `json:"branch_protections"`
}

type BranchProtection struct {
	Name string `json:"name"`
	Protection *github.ProtectionRequest `json:"protection_request"`
}

func main() {
	var setting Setting

	buf, err := ioutil.ReadFile("./settings.yml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(buf, &setting)
	if err != nil {
		log.Fatal(err)
	}

	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("No github token")
	}

	ts := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: token,
	})
	ctx := context.Background()
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	for _, repo := range setting.Repos {
		for _, protection := range setting.BranchProtections {
			pro, _, err := client.Repositories.UpdateBranchProtection(ctx, setting.Owner, repo, protection.Name, protection.Protection)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(pro)
		}
	}
}
