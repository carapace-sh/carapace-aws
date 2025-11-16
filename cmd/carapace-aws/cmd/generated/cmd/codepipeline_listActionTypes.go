package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_listActionTypesCmd = &cobra.Command{
	Use:   "list-action-types",
	Short: "Gets a summary of all CodePipeline action types associated with your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_listActionTypesCmd).Standalone()

	codepipeline_listActionTypesCmd.Flags().String("action-owner-filter", "", "Filters the list of action types to those created by a specified entity.")
	codepipeline_listActionTypesCmd.Flags().String("next-token", "", "An identifier that was returned from the previous list action types call, which can be used to return the next set of action types in the list.")
	codepipeline_listActionTypesCmd.Flags().String("region-filter", "", "The Region to filter on for the list of action types.")
	codepipelineCmd.AddCommand(codepipeline_listActionTypesCmd)
}
