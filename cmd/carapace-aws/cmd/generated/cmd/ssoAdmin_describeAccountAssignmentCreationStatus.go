package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_describeAccountAssignmentCreationStatusCmd = &cobra.Command{
	Use:   "describe-account-assignment-creation-status",
	Short: "Describes the status of the assignment creation request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_describeAccountAssignmentCreationStatusCmd).Standalone()

	ssoAdmin_describeAccountAssignmentCreationStatusCmd.Flags().String("account-assignment-creation-request-id", "", "The identifier that is used to track the request operation progress.")
	ssoAdmin_describeAccountAssignmentCreationStatusCmd.Flags().String("instance-arn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
	ssoAdmin_describeAccountAssignmentCreationStatusCmd.MarkFlagRequired("account-assignment-creation-request-id")
	ssoAdmin_describeAccountAssignmentCreationStatusCmd.MarkFlagRequired("instance-arn")
	ssoAdminCmd.AddCommand(ssoAdmin_describeAccountAssignmentCreationStatusCmd)
}
