package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_getPermissionPolicyCmd = &cobra.Command{
	Use:   "get-permission-policy",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_getPermissionPolicyCmd).Standalone()

	waf_getPermissionPolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the RuleGroup for which you want to get the policy.")
	waf_getPermissionPolicyCmd.MarkFlagRequired("resource-arn")
	wafCmd.AddCommand(waf_getPermissionPolicyCmd)
}
