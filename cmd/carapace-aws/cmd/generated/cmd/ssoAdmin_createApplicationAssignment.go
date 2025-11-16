package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_createApplicationAssignmentCmd = &cobra.Command{
	Use:   "create-application-assignment",
	Short: "Grant application access to a user or group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_createApplicationAssignmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoAdmin_createApplicationAssignmentCmd).Standalone()

		ssoAdmin_createApplicationAssignmentCmd.Flags().String("application-arn", "", "The ARN of the application for which the assignment is created.")
		ssoAdmin_createApplicationAssignmentCmd.Flags().String("principal-id", "", "An identifier for an object in IAM Identity Center, such as a user or group.")
		ssoAdmin_createApplicationAssignmentCmd.Flags().String("principal-type", "", "The entity type for which the assignment will be created.")
		ssoAdmin_createApplicationAssignmentCmd.MarkFlagRequired("application-arn")
		ssoAdmin_createApplicationAssignmentCmd.MarkFlagRequired("principal-id")
		ssoAdmin_createApplicationAssignmentCmd.MarkFlagRequired("principal-type")
	})
	ssoAdminCmd.AddCommand(ssoAdmin_createApplicationAssignmentCmd)
}
