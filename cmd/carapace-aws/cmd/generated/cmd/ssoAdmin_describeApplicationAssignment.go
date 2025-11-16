package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_describeApplicationAssignmentCmd = &cobra.Command{
	Use:   "describe-application-assignment",
	Short: "Retrieves a direct assignment of a user or group to an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_describeApplicationAssignmentCmd).Standalone()

	ssoAdmin_describeApplicationAssignmentCmd.Flags().String("application-arn", "", "Specifies the ARN of the application.")
	ssoAdmin_describeApplicationAssignmentCmd.Flags().String("principal-id", "", "An identifier for an object in IAM Identity Center, such as a user or group.")
	ssoAdmin_describeApplicationAssignmentCmd.Flags().String("principal-type", "", "The entity type for which the assignment will be created.")
	ssoAdmin_describeApplicationAssignmentCmd.MarkFlagRequired("application-arn")
	ssoAdmin_describeApplicationAssignmentCmd.MarkFlagRequired("principal-id")
	ssoAdmin_describeApplicationAssignmentCmd.MarkFlagRequired("principal-type")
	ssoAdminCmd.AddCommand(ssoAdmin_describeApplicationAssignmentCmd)
}
