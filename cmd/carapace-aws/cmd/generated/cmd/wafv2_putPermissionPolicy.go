package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_putPermissionPolicyCmd = &cobra.Command{
	Use:   "put-permission-policy",
	Short: "Use this to share a rule group with other accounts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_putPermissionPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafv2_putPermissionPolicyCmd).Standalone()

		wafv2_putPermissionPolicyCmd.Flags().String("policy", "", "The policy to attach to the specified rule group.")
		wafv2_putPermissionPolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the [RuleGroup]() to which you want to attach the policy.")
		wafv2_putPermissionPolicyCmd.MarkFlagRequired("policy")
		wafv2_putPermissionPolicyCmd.MarkFlagRequired("resource-arn")
	})
	wafv2Cmd.AddCommand(wafv2_putPermissionPolicyCmd)
}
