package cmd

import (
	"github.com/spf13/cobra"
)

var file string

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "apply setting from yaml",
	Run: func(cmd *cobra.Command, args []string) {
		run(file)
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)
	applyCmd.Flags().StringVarP(&file, "file", "f", "settings.yml", "Set file path")
}
