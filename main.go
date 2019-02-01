package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/google/go-github/github"
	"gopkg.in/yaml.v2"
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

	fmt.Println(setting.Owner)
	fmt.Println(setting.Repos)

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
