package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_deletePermissionPolicyCmd = &cobra.Command{
	Use:   "delete-permission-policy",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_deletePermissionPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_deletePermissionPolicyCmd).Standalone()

		waf_deletePermissionPolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the RuleGroup from which you want to delete the policy.")
		waf_deletePermissionPolicyCmd.MarkFlagRequired("resource-arn")
	})
	wafCmd.AddCommand(waf_deletePermissionPolicyCmd)
}
