package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_createAccountAssignmentCmd = &cobra.Command{
	Use:   "create-account-assignment",
	Short: "Assigns access to a principal for a specified Amazon Web Services account using a specified permission set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_createAccountAssignmentCmd).Standalone()

	ssoAdmin_createAccountAssignmentCmd.Flags().String("instance-arn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
	ssoAdmin_createAccountAssignmentCmd.Flags().String("permission-set-arn", "", "The ARN of the permission set that the admin wants to grant the principal access to.")
	ssoAdmin_createAccountAssignmentCmd.Flags().String("principal-id", "", "An identifier for an object in IAM Identity Center, such as a user or group.")
	ssoAdmin_createAccountAssignmentCmd.Flags().String("principal-type", "", "The entity type for which the assignment will be created.")
	ssoAdmin_createAccountAssignmentCmd.Flags().String("target-id", "", "TargetID is an Amazon Web Services account identifier, (For example, 123456789012).")
	ssoAdmin_createAccountAssignmentCmd.Flags().String("target-type", "", "The entity type for which the assignment will be created.")
	ssoAdmin_createAccountAssignmentCmd.MarkFlagRequired("instance-arn")
	ssoAdmin_createAccountAssignmentCmd.MarkFlagRequired("permission-set-arn")
	ssoAdmin_createAccountAssignmentCmd.MarkFlagRequired("principal-id")
	ssoAdmin_createAccountAssignmentCmd.MarkFlagRequired("principal-type")
	ssoAdmin_createAccountAssignmentCmd.MarkFlagRequired("target-id")
	ssoAdmin_createAccountAssignmentCmd.MarkFlagRequired("target-type")
	ssoAdminCmd.AddCommand(ssoAdmin_createAccountAssignmentCmd)
}
