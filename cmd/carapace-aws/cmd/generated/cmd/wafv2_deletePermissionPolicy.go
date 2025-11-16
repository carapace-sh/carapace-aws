package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_deletePermissionPolicyCmd = &cobra.Command{
	Use:   "delete-permission-policy",
	Short: "Permanently deletes an IAM policy from the specified rule group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_deletePermissionPolicyCmd).Standalone()

	wafv2_deletePermissionPolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the rule group from which you want to delete the policy.")
	wafv2_deletePermissionPolicyCmd.MarkFlagRequired("resource-arn")
	wafv2Cmd.AddCommand(wafv2_deletePermissionPolicyCmd)
}
