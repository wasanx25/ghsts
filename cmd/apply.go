package cmd

import (
	"github.com/spf13/cobra"
)

var file string
var dryRun bool

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "apply setting from yaml",
	Run: func(cmd *cobra.Command, args []string) {
		run(file, dryRun)
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)
	applyCmd.Flags().StringVarP(&file, "file", "f", "settings.yml", "Filename that contains the configuration to apply")
	applyCmd.Flags().BoolVarP(&dryRun, "dry-run", "n", false, "Do a dry run without executing actions")
}
