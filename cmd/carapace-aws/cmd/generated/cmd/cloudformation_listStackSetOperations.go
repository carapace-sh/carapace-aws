package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_listStackSetOperationsCmd = &cobra.Command{
	Use:   "list-stack-set-operations",
	Short: "Returns summary information about operations performed on a StackSet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_listStackSetOperationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_listStackSetOperationsCmd).Standalone()

		cloudformation_listStackSetOperationsCmd.Flags().String("call-as", "", "\\[Service-managed permissions] Specifies whether you are acting as an account administrator in the organization's management account or as a delegated administrator in a member account.")
		cloudformation_listStackSetOperationsCmd.Flags().String("max-results", "", "The maximum number of results to be returned with a single call.")
		cloudformation_listStackSetOperationsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		cloudformation_listStackSetOperationsCmd.Flags().String("stack-set-name", "", "The name or unique ID of the StackSet that you want to get operation summaries for.")
		cloudformation_listStackSetOperationsCmd.MarkFlagRequired("stack-set-name")
	})
	cloudformationCmd.AddCommand(cloudformation_listStackSetOperationsCmd)
}
