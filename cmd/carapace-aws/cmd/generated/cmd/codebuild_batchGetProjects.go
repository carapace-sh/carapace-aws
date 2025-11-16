package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_batchGetProjectsCmd = &cobra.Command{
	Use:   "batch-get-projects",
	Short: "Gets information about one or more build projects.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_batchGetProjectsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codebuild_batchGetProjectsCmd).Standalone()

		codebuild_batchGetProjectsCmd.Flags().String("names", "", "The names or ARNs of the build projects.")
		codebuild_batchGetProjectsCmd.MarkFlagRequired("names")
	})
	codebuildCmd.AddCommand(codebuild_batchGetProjectsCmd)
}
