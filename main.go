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
	Rules []Rule `json:"rules"`
	Owner string `json:"owner"`
	Repos []Repo `json:"repos"`
}

type Rule struct {
	Name              string                    `json:"name"`
	BranchName        string                    `json:"branch_name"`
	ProtectionRequest *github.ProtectionRequest `json:"protection_request"`
}

type Repo struct {
	Name string `json:"name"`
	Rule string `json:"rule"`
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
		rule := setting.findRule(repo.Rule)
		pro, _, err := client.Repositories.UpdateBranchProtection(ctx, setting.Owner, repo.Name, rule.BranchName, rule.ProtectionRequest)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(pro)
	}
}

func (s *Setting) findRule(name string) Rule {
	for _, rule := range s.Rules {
		if rule.Name == name {
			return rule
		}
	}
	panic(fmt.Sprintf("it has not Rule(%s)", name))
}