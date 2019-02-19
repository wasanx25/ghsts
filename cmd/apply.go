package cmd

import (
	"context"
	"io/ioutil"
	"log"
	"os"

	"github.com/ghodss/yaml"
	"github.com/google/go-github/github"
	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
	"golang.org/x/sync/errgroup"
)

var file string
var dryRun bool

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "apply setting from yaml",
	Run: func(cmd *cobra.Command, args []string) {
		applySetting(file, dryRun)
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)
	applyCmd.Flags().StringVarP(&file, "file", "f", "settings.yml", "Filename that contains the configuration to apply")
	applyCmd.Flags().BoolVarP(&dryRun, "dry-run", "n", false, "Do a dry run without executing actions")
}

type Setting struct {
	Rules []*Rule  `json:"rules"`
	Owner string   `json:"owner"`
	Repos []string `json:"repos"`
}

type Rule struct {
	BranchName        string                    `json:"branch_name"`
	ProtectionRequest *github.ProtectionRequest `json:"protection_request"`
}

func applySetting(file string, dryRun bool) {
	var setting Setting

	buf, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(buf, &setting)
	if err != nil {
		log.Fatal(err)
	}

	if dryRun {
		pp.Println(setting)
		os.Exit(0)
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

	eg := errgroup.Group{}

	for _, repo := range setting.Repos {
		eg.Go(func() error {
			for _, rule := range setting.Rules {
				_, _, err = client.Repositories.UpdateBranchProtection(ctx, setting.Owner, repo, rule.BranchName, rule.ProtectionRequest)
				return err
			}
			return nil
		})
	}

	if err = eg.Wait(); err != nil {
		log.Fatal(err)
	}

	log.Println("Success!")
}
