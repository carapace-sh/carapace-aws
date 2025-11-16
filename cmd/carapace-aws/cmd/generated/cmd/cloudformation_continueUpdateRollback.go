package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_continueUpdateRollbackCmd = &cobra.Command{
	Use:   "continue-update-rollback",
	Short: "Continues rolling back a stack from `UPDATE_ROLLBACK_FAILED` to `UPDATE_ROLLBACK_COMPLETE` state.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_continueUpdateRollbackCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_continueUpdateRollbackCmd).Standalone()

		cloudformation_continueUpdateRollbackCmd.Flags().String("client-request-token", "", "A unique identifier for this `ContinueUpdateRollback` request.")
		cloudformation_continueUpdateRollbackCmd.Flags().String("resources-to-skip", "", "A list of the logical IDs of the resources that CloudFormation skips during the continue update rollback operation.")
		cloudformation_continueUpdateRollbackCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role that CloudFormation assumes to roll back the stack.")
		cloudformation_continueUpdateRollbackCmd.Flags().String("stack-name", "", "The name or the unique ID of the stack that you want to continue rolling back.")
		cloudformation_continueUpdateRollbackCmd.MarkFlagRequired("stack-name")
	})
	cloudformationCmd.AddCommand(cloudformation_continueUpdateRollbackCmd)
}
