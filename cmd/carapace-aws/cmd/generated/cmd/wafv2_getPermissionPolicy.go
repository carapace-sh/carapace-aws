package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_getPermissionPolicyCmd = &cobra.Command{
	Use:   "get-permission-policy",
	Short: "Returns the IAM policy that is attached to the specified rule group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_getPermissionPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafv2_getPermissionPolicyCmd).Standalone()

		wafv2_getPermissionPolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the rule group for which you want to get the policy.")
		wafv2_getPermissionPolicyCmd.MarkFlagRequired("resource-arn")
	})
	wafv2Cmd.AddCommand(wafv2_getPermissionPolicyCmd)
}
