package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_putResourcePolicyCmd = &cobra.Command{
	Use:   "put-resource-policy",
	Short: "Creates or updates a Systems Manager resource policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_putResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_putResourcePolicyCmd).Standalone()

		ssm_putResourcePolicyCmd.Flags().String("policy", "", "A policy you want to associate with a resource.")
		ssm_putResourcePolicyCmd.Flags().String("policy-hash", "", "ID of the current policy version.")
		ssm_putResourcePolicyCmd.Flags().String("policy-id", "", "The policy ID.")
		ssm_putResourcePolicyCmd.Flags().String("resource-arn", "", "Amazon Resource Name (ARN) of the resource to which you want to attach a policy.")
		ssm_putResourcePolicyCmd.MarkFlagRequired("policy")
		ssm_putResourcePolicyCmd.MarkFlagRequired("resource-arn")
	})
	ssmCmd.AddCommand(ssm_putResourcePolicyCmd)
}
