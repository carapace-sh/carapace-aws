package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_getPermissionPolicyCmd = &cobra.Command{
	Use:   "get-permission-policy",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_getPermissionPolicyCmd).Standalone()

	wafRegional_getPermissionPolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the RuleGroup for which you want to get the policy.")
	wafRegional_getPermissionPolicyCmd.MarkFlagRequired("resource-arn")
	wafRegionalCmd.AddCommand(wafRegional_getPermissionPolicyCmd)
}
