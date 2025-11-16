package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_deleteResourcePolicyStatementCmd = &cobra.Command{
	Use:   "delete-resource-policy-statement",
	Short: "Deletes a policy statement from a resource policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_deleteResourcePolicyStatementCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_deleteResourcePolicyStatementCmd).Standalone()

		lexv2Models_deleteResourcePolicyStatementCmd.Flags().String("expected-revision-id", "", "The identifier of the revision of the policy to delete the statement from.")
		lexv2Models_deleteResourcePolicyStatementCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the bot or bot alias that the resource policy is attached to.")
		lexv2Models_deleteResourcePolicyStatementCmd.Flags().String("statement-id", "", "The name of the statement (SID) to delete from the policy.")
		lexv2Models_deleteResourcePolicyStatementCmd.MarkFlagRequired("resource-arn")
		lexv2Models_deleteResourcePolicyStatementCmd.MarkFlagRequired("statement-id")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_deleteResourcePolicyStatementCmd)
}
