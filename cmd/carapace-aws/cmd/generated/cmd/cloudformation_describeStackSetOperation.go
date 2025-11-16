package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_describeStackSetOperationCmd = &cobra.Command{
	Use:   "describe-stack-set-operation",
	Short: "Returns the description of the specified StackSet operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_describeStackSetOperationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_describeStackSetOperationCmd).Standalone()

		cloudformation_describeStackSetOperationCmd.Flags().String("call-as", "", "\\[Service-managed permissions] Specifies whether you are acting as an account administrator in the organization's management account or as a delegated administrator in a member account.")
		cloudformation_describeStackSetOperationCmd.Flags().String("operation-id", "", "The unique ID of the StackSet operation.")
		cloudformation_describeStackSetOperationCmd.Flags().String("stack-set-name", "", "The name or the unique stack ID of the StackSet for the stack operation.")
		cloudformation_describeStackSetOperationCmd.MarkFlagRequired("operation-id")
		cloudformation_describeStackSetOperationCmd.MarkFlagRequired("stack-set-name")
	})
	cloudformationCmd.AddCommand(cloudformation_describeStackSetOperationCmd)
}
