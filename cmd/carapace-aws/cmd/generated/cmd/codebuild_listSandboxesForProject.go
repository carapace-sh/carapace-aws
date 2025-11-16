package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_listSandboxesForProjectCmd = &cobra.Command{
	Use:   "list-sandboxes-for-project",
	Short: "Gets a list of sandboxes for a given project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_listSandboxesForProjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codebuild_listSandboxesForProjectCmd).Standalone()

		codebuild_listSandboxesForProjectCmd.Flags().String("max-results", "", "The maximum number of sandbox records to be retrieved.")
		codebuild_listSandboxesForProjectCmd.Flags().String("next-token", "", "The next token, if any, to get paginated results.")
		codebuild_listSandboxesForProjectCmd.Flags().String("project-name", "", "The CodeBuild project name.")
		codebuild_listSandboxesForProjectCmd.Flags().String("sort-order", "", "The order in which sandbox records should be retrieved.")
		codebuild_listSandboxesForProjectCmd.MarkFlagRequired("project-name")
	})
	codebuildCmd.AddCommand(codebuild_listSandboxesForProjectCmd)
}
