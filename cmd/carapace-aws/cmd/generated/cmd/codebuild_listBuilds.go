package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_listBuildsCmd = &cobra.Command{
	Use:   "list-builds",
	Short: "Gets a list of build IDs, with each build ID representing a single build.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_listBuildsCmd).Standalone()

	codebuild_listBuildsCmd.Flags().String("next-token", "", "During a previous call, if there are more than 100 items in the list, only the first 100 items are returned, along with a unique string called a *nextToken*.")
	codebuild_listBuildsCmd.Flags().String("sort-order", "", "The order to list build IDs.")
	codebuildCmd.AddCommand(codebuild_listBuildsCmd)
}
