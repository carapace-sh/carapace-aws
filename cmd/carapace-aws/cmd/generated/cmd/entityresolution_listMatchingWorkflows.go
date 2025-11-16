package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_listMatchingWorkflowsCmd = &cobra.Command{
	Use:   "list-matching-workflows",
	Short: "Returns a list of all the `MatchingWorkflows` that have been created for an Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_listMatchingWorkflowsCmd).Standalone()

	entityresolution_listMatchingWorkflowsCmd.Flags().String("max-results", "", "The maximum number of objects returned per page.")
	entityresolution_listMatchingWorkflowsCmd.Flags().String("next-token", "", "The pagination token from the previous API call.")
	entityresolutionCmd.AddCommand(entityresolution_listMatchingWorkflowsCmd)
}
