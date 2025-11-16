package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_rollbackStackCmd = &cobra.Command{
	Use:   "rollback-stack",
	Short: "When specifying `RollbackStack`, you preserve the state of previously provisioned resources when an operation fails.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_rollbackStackCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_rollbackStackCmd).Standalone()

		cloudformation_rollbackStackCmd.Flags().String("client-request-token", "", "A unique identifier for this `RollbackStack` request.")
		cloudformation_rollbackStackCmd.Flags().String("retain-except-on-create", "", "When set to `true`, newly created resources are deleted when the operation rolls back.")
		cloudformation_rollbackStackCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role that CloudFormation assumes to rollback the stack.")
		cloudformation_rollbackStackCmd.Flags().String("stack-name", "", "The name that's associated with the stack.")
		cloudformation_rollbackStackCmd.MarkFlagRequired("stack-name")
	})
	cloudformationCmd.AddCommand(cloudformation_rollbackStackCmd)
}
