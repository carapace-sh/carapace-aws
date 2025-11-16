package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_describeAccountAssignmentDeletionStatusCmd = &cobra.Command{
	Use:   "describe-account-assignment-deletion-status",
	Short: "Describes the status of the assignment deletion request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_describeAccountAssignmentDeletionStatusCmd).Standalone()

	ssoAdmin_describeAccountAssignmentDeletionStatusCmd.Flags().String("account-assignment-deletion-request-id", "", "The identifier that is used to track the request operation progress.")
	ssoAdmin_describeAccountAssignmentDeletionStatusCmd.Flags().String("instance-arn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
	ssoAdmin_describeAccountAssignmentDeletionStatusCmd.MarkFlagRequired("account-assignment-deletion-request-id")
	ssoAdmin_describeAccountAssignmentDeletionStatusCmd.MarkFlagRequired("instance-arn")
	ssoAdminCmd.AddCommand(ssoAdmin_describeAccountAssignmentDeletionStatusCmd)
}
