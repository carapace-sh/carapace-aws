package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_listBuildsForProjectCmd = &cobra.Command{
	Use:   "list-builds-for-project",
	Short: "Gets a list of build identifiers for the specified build project, with each build identifier representing a single build.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_listBuildsForProjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codebuild_listBuildsForProjectCmd).Standalone()

		codebuild_listBuildsForProjectCmd.Flags().String("next-token", "", "During a previous call, if there are more than 100 items in the list, only the first 100 items are returned, along with a unique string called a *nextToken*.")
		codebuild_listBuildsForProjectCmd.Flags().String("project-name", "", "The name of the CodeBuild project.")
		codebuild_listBuildsForProjectCmd.Flags().String("sort-order", "", "The order to sort the results in.")
		codebuild_listBuildsForProjectCmd.MarkFlagRequired("project-name")
	})
	codebuildCmd.AddCommand(codebuild_listBuildsForProjectCmd)
}
