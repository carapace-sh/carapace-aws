package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocsCmd = &cobra.Command{
	Use:   "workdocs",
	Short: "The Amazon WorkDocs API is designed for the following use cases:",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workdocsCmd).Standalone()

	})
	rootCmd.AddCommand(workdocsCmd)
}
