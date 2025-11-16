package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_putPermissionPolicyCmd = &cobra.Command{
	Use:   "put-permission-policy",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_putPermissionPolicyCmd).Standalone()

	wafRegional_putPermissionPolicyCmd.Flags().String("policy", "", "The policy to attach to the specified RuleGroup.")
	wafRegional_putPermissionPolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the RuleGroup to which you want to attach the policy.")
	wafRegional_putPermissionPolicyCmd.MarkFlagRequired("policy")
	wafRegional_putPermissionPolicyCmd.MarkFlagRequired("resource-arn")
	wafRegionalCmd.AddCommand(wafRegional_putPermissionPolicyCmd)
}
