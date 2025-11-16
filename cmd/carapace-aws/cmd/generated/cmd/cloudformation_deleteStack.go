package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_deleteStackCmd = &cobra.Command{
	Use:   "delete-stack",
	Short: "Deletes a specified stack.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_deleteStackCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_deleteStackCmd).Standalone()

		cloudformation_deleteStackCmd.Flags().String("client-request-token", "", "A unique identifier for this `DeleteStack` request.")
		cloudformation_deleteStackCmd.Flags().String("deletion-mode", "", "Specifies the deletion mode for the stack.")
		cloudformation_deleteStackCmd.Flags().String("retain-resources", "", "For stacks in the `DELETE_FAILED` state, a list of resource logical IDs that are associated with the resources you want to retain.")
		cloudformation_deleteStackCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role that CloudFormation assumes to delete the stack.")
		cloudformation_deleteStackCmd.Flags().String("stack-name", "", "The name or the unique stack ID that's associated with the stack.")
		cloudformation_deleteStackCmd.MarkFlagRequired("stack-name")
	})
	cloudformationCmd.AddCommand(cloudformation_deleteStackCmd)
}
