package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_putResourcePolicyCmd = &cobra.Command{
	Use:   "put-resource-policy",
	Short: "Creates or updates a resource policy allowing other Amazon Web Services services to put log events to this account, such as Amazon Route 53. An account can have up to 10 resource policies per Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_putResourcePolicyCmd).Standalone()

	logs_putResourcePolicyCmd.Flags().String("expected-revision-id", "", "The expected revision ID of the resource policy.")
	logs_putResourcePolicyCmd.Flags().String("policy-document", "", "Details of the new policy, including the identity of the principal that is enabled to put logs to this account.")
	logs_putResourcePolicyCmd.Flags().String("policy-name", "", "Name of the new policy.")
	logs_putResourcePolicyCmd.Flags().String("resource-arn", "", "The ARN of the CloudWatch Logs resource to which the resource policy needs to be added or attached.")
	logsCmd.AddCommand(logs_putResourcePolicyCmd)
}
