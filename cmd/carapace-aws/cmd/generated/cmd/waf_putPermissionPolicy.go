package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_putPermissionPolicyCmd = &cobra.Command{
	Use:   "put-permission-policy",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_putPermissionPolicyCmd).Standalone()

	waf_putPermissionPolicyCmd.Flags().String("policy", "", "The policy to attach to the specified RuleGroup.")
	waf_putPermissionPolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the RuleGroup to which you want to attach the policy.")
	waf_putPermissionPolicyCmd.MarkFlagRequired("policy")
	waf_putPermissionPolicyCmd.MarkFlagRequired("resource-arn")
	wafCmd.AddCommand(waf_putPermissionPolicyCmd)
}
