package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_deleteApplicationAssignmentCmd = &cobra.Command{
	Use:   "delete-application-assignment",
	Short: "Revoke application access to an application by deleting application assignments for a user or group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_deleteApplicationAssignmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoAdmin_deleteApplicationAssignmentCmd).Standalone()

		ssoAdmin_deleteApplicationAssignmentCmd.Flags().String("application-arn", "", "Specifies the ARN of the application.")
		ssoAdmin_deleteApplicationAssignmentCmd.Flags().String("principal-id", "", "An identifier for an object in IAM Identity Center, such as a user or group.")
		ssoAdmin_deleteApplicationAssignmentCmd.Flags().String("principal-type", "", "The entity type for which the assignment will be deleted.")
		ssoAdmin_deleteApplicationAssignmentCmd.MarkFlagRequired("application-arn")
		ssoAdmin_deleteApplicationAssignmentCmd.MarkFlagRequired("principal-id")
		ssoAdmin_deleteApplicationAssignmentCmd.MarkFlagRequired("principal-type")
	})
	ssoAdminCmd.AddCommand(ssoAdmin_deleteApplicationAssignmentCmd)
}
