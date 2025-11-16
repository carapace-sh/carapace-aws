package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_deleteResourcePolicyCmd = &cobra.Command{
	Use:   "delete-resource-policy",
	Short: "Deletes a resource policy from this account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_deleteResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_deleteResourcePolicyCmd).Standalone()

		logs_deleteResourcePolicyCmd.Flags().String("expected-revision-id", "", "The expected revision ID of the resource policy.")
		logs_deleteResourcePolicyCmd.Flags().String("policy-name", "", "The name of the policy to be revoked.")
		logs_deleteResourcePolicyCmd.Flags().String("resource-arn", "", "The ARN of the CloudWatch Logs resource for which the resource policy needs to be deleted")
	})
	logsCmd.AddCommand(logs_deleteResourcePolicyCmd)
}
