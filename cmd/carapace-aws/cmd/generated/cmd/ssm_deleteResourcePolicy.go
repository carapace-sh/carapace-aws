package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_deleteResourcePolicyCmd = &cobra.Command{
	Use:   "delete-resource-policy",
	Short: "Deletes a Systems Manager resource policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_deleteResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_deleteResourcePolicyCmd).Standalone()

		ssm_deleteResourcePolicyCmd.Flags().String("policy-hash", "", "ID of the current policy version.")
		ssm_deleteResourcePolicyCmd.Flags().String("policy-id", "", "The policy ID.")
		ssm_deleteResourcePolicyCmd.Flags().String("resource-arn", "", "Amazon Resource Name (ARN) of the resource to which the policies are attached.")
		ssm_deleteResourcePolicyCmd.MarkFlagRequired("policy-hash")
		ssm_deleteResourcePolicyCmd.MarkFlagRequired("policy-id")
		ssm_deleteResourcePolicyCmd.MarkFlagRequired("resource-arn")
	})
	ssmCmd.AddCommand(ssm_deleteResourcePolicyCmd)
}
