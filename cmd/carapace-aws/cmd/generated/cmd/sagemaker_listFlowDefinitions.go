package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listFlowDefinitionsCmd = &cobra.Command{
	Use:   "list-flow-definitions",
	Short: "Returns information about the flow definitions in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listFlowDefinitionsCmd).Standalone()

	sagemaker_listFlowDefinitionsCmd.Flags().String("creation-time-after", "", "A filter that returns only flow definitions with a creation time greater than or equal to the specified timestamp.")
	sagemaker_listFlowDefinitionsCmd.Flags().String("creation-time-before", "", "A filter that returns only flow definitions that were created before the specified timestamp.")
	sagemaker_listFlowDefinitionsCmd.Flags().String("max-results", "", "The total number of items to return.")
	sagemaker_listFlowDefinitionsCmd.Flags().String("next-token", "", "A token to resume pagination.")
	sagemaker_listFlowDefinitionsCmd.Flags().String("sort-order", "", "An optional value that specifies whether you want the results sorted in `Ascending` or `Descending` order.")
	sagemakerCmd.AddCommand(sagemaker_listFlowDefinitionsCmd)
}
