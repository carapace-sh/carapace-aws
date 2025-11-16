package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_listPermissionSetsCmd = &cobra.Command{
	Use:   "list-permission-sets",
	Short: "Lists the [PermissionSet]()s in an IAM Identity Center instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_listPermissionSetsCmd).Standalone()

	ssoAdmin_listPermissionSetsCmd.Flags().String("instance-arn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
	ssoAdmin_listPermissionSetsCmd.Flags().String("max-results", "", "The maximum number of results to display for the assignment.")
	ssoAdmin_listPermissionSetsCmd.Flags().String("next-token", "", "The pagination token for the list API.")
	ssoAdmin_listPermissionSetsCmd.MarkFlagRequired("instance-arn")
	ssoAdminCmd.AddCommand(ssoAdmin_listPermissionSetsCmd)
}
