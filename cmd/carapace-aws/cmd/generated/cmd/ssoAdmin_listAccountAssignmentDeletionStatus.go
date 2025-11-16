package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_listAccountAssignmentDeletionStatusCmd = &cobra.Command{
	Use:   "list-account-assignment-deletion-status",
	Short: "Lists the status of the Amazon Web Services account assignment deletion requests for a specified IAM Identity Center instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_listAccountAssignmentDeletionStatusCmd).Standalone()

	ssoAdmin_listAccountAssignmentDeletionStatusCmd.Flags().String("filter", "", "Filters results based on the passed attribute value.")
	ssoAdmin_listAccountAssignmentDeletionStatusCmd.Flags().String("instance-arn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
	ssoAdmin_listAccountAssignmentDeletionStatusCmd.Flags().String("max-results", "", "The maximum number of results to display for the assignment.")
	ssoAdmin_listAccountAssignmentDeletionStatusCmd.Flags().String("next-token", "", "The pagination token for the list API.")
	ssoAdmin_listAccountAssignmentDeletionStatusCmd.MarkFlagRequired("instance-arn")
	ssoAdminCmd.AddCommand(ssoAdmin_listAccountAssignmentDeletionStatusCmd)
}
