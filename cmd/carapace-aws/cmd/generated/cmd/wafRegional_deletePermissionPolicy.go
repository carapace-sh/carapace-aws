package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_deletePermissionPolicyCmd = &cobra.Command{
	Use:   "delete-permission-policy",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_deletePermissionPolicyCmd).Standalone()

	wafRegional_deletePermissionPolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the RuleGroup from which you want to delete the policy.")
	wafRegional_deletePermissionPolicyCmd.MarkFlagRequired("resource-arn")
	wafRegionalCmd.AddCommand(wafRegional_deletePermissionPolicyCmd)
}
