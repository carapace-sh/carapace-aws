package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_listFleetsCmd = &cobra.Command{
	Use:   "list-fleets",
	Short: "Gets a list of compute fleet names with each compute fleet name representing a single compute fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_listFleetsCmd).Standalone()

	codebuild_listFleetsCmd.Flags().String("max-results", "", "The maximum number of paginated compute fleets returned per response.")
	codebuild_listFleetsCmd.Flags().String("next-token", "", "During a previous call, if there are more than 100 items in the list, only the first 100 items are returned, along with a unique string called a *nextToken*.")
	codebuild_listFleetsCmd.Flags().String("sort-by", "", "The criterion to be used to list compute fleet names.")
	codebuild_listFleetsCmd.Flags().String("sort-order", "", "The order in which to list compute fleets.")
	codebuildCmd.AddCommand(codebuild_listFleetsCmd)
}
