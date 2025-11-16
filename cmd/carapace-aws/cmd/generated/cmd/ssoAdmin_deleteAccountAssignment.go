package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_deleteAccountAssignmentCmd = &cobra.Command{
	Use:   "delete-account-assignment",
	Short: "Deletes a principal's access from a specified Amazon Web Services account using a specified permission set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_deleteAccountAssignmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoAdmin_deleteAccountAssignmentCmd).Standalone()

		ssoAdmin_deleteAccountAssignmentCmd.Flags().String("instance-arn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
		ssoAdmin_deleteAccountAssignmentCmd.Flags().String("permission-set-arn", "", "The ARN of the permission set that will be used to remove access.")
		ssoAdmin_deleteAccountAssignmentCmd.Flags().String("principal-id", "", "An identifier for an object in IAM Identity Center, such as a user or group.")
		ssoAdmin_deleteAccountAssignmentCmd.Flags().String("principal-type", "", "The entity type for which the assignment will be deleted.")
		ssoAdmin_deleteAccountAssignmentCmd.Flags().String("target-id", "", "TargetID is an Amazon Web Services account identifier, (For example, 123456789012).")
		ssoAdmin_deleteAccountAssignmentCmd.Flags().String("target-type", "", "The entity type for which the assignment will be deleted.")
		ssoAdmin_deleteAccountAssignmentCmd.MarkFlagRequired("instance-arn")
		ssoAdmin_deleteAccountAssignmentCmd.MarkFlagRequired("permission-set-arn")
		ssoAdmin_deleteAccountAssignmentCmd.MarkFlagRequired("principal-id")
		ssoAdmin_deleteAccountAssignmentCmd.MarkFlagRequired("principal-type")
		ssoAdmin_deleteAccountAssignmentCmd.MarkFlagRequired("target-id")
		ssoAdmin_deleteAccountAssignmentCmd.MarkFlagRequired("target-type")
	})
	ssoAdminCmd.AddCommand(ssoAdmin_deleteAccountAssignmentCmd)
}
