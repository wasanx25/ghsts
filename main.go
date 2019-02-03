package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"github.com/ghodss/yaml"
)

type Setting struct {
	Owner string   `yaml:"owner"`
	Repos []string `yaml:"repos"`
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
		info, _, err := client.Repositories.Get(ctx, setting.Owner, repo)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(info.GetFullName())
	}
}
