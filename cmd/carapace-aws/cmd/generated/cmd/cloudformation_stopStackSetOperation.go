package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_stopStackSetOperationCmd = &cobra.Command{
	Use:   "stop-stack-set-operation",
	Short: "Stops an in-progress operation on a StackSet and its associated stack instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_stopStackSetOperationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_stopStackSetOperationCmd).Standalone()

		cloudformation_stopStackSetOperationCmd.Flags().String("call-as", "", "Specifies whether you are acting as an account administrator in the organization's management account or as a delegated administrator in a member account.")
		cloudformation_stopStackSetOperationCmd.Flags().String("operation-id", "", "The ID of the stack operation.")
		cloudformation_stopStackSetOperationCmd.Flags().String("stack-set-name", "", "The name or unique ID of the StackSet that you want to stop the operation for.")
		cloudformation_stopStackSetOperationCmd.MarkFlagRequired("operation-id")
		cloudformation_stopStackSetOperationCmd.MarkFlagRequired("stack-set-name")
	})
	cloudformationCmd.AddCommand(cloudformation_stopStackSetOperationCmd)
}
