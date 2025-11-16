package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_listAccountAssignmentsCmd = &cobra.Command{
	Use:   "list-account-assignments",
	Short: "Lists the assignee of the specified Amazon Web Services account with the specified permission set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_listAccountAssignmentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoAdmin_listAccountAssignmentsCmd).Standalone()

		ssoAdmin_listAccountAssignmentsCmd.Flags().String("account-id", "", "The identifier of the Amazon Web Services account from which to list the assignments.")
		ssoAdmin_listAccountAssignmentsCmd.Flags().String("instance-arn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
		ssoAdmin_listAccountAssignmentsCmd.Flags().String("max-results", "", "The maximum number of results to display for the assignment.")
		ssoAdmin_listAccountAssignmentsCmd.Flags().String("next-token", "", "The pagination token for the list API.")
		ssoAdmin_listAccountAssignmentsCmd.Flags().String("permission-set-arn", "", "The ARN of the permission set from which to list assignments.")
		ssoAdmin_listAccountAssignmentsCmd.MarkFlagRequired("account-id")
		ssoAdmin_listAccountAssignmentsCmd.MarkFlagRequired("instance-arn")
		ssoAdmin_listAccountAssignmentsCmd.MarkFlagRequired("permission-set-arn")
	})
	ssoAdminCmd.AddCommand(ssoAdmin_listAccountAssignmentsCmd)
}
