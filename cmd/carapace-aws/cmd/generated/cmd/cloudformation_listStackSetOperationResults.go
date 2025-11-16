package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_listStackSetOperationResultsCmd = &cobra.Command{
	Use:   "list-stack-set-operation-results",
	Short: "Returns summary information about the results of a StackSet operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_listStackSetOperationResultsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_listStackSetOperationResultsCmd).Standalone()

		cloudformation_listStackSetOperationResultsCmd.Flags().String("call-as", "", "\\[Service-managed permissions] Specifies whether you are acting as an account administrator in the organization's management account or as a delegated administrator in a member account.")
		cloudformation_listStackSetOperationResultsCmd.Flags().String("filters", "", "The filter to apply to operation results.")
		cloudformation_listStackSetOperationResultsCmd.Flags().String("max-results", "", "The maximum number of results to be returned with a single call.")
		cloudformation_listStackSetOperationResultsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		cloudformation_listStackSetOperationResultsCmd.Flags().String("operation-id", "", "The ID of the StackSet operation.")
		cloudformation_listStackSetOperationResultsCmd.Flags().String("stack-set-name", "", "The name or unique ID of the StackSet that you want to get operation results for.")
		cloudformation_listStackSetOperationResultsCmd.MarkFlagRequired("operation-id")
		cloudformation_listStackSetOperationResultsCmd.MarkFlagRequired("stack-set-name")
	})
	cloudformationCmd.AddCommand(cloudformation_listStackSetOperationResultsCmd)
}
