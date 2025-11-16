package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_associatePermissionCmd = &cobra.Command{
	Use:   "associate-permission",
	Short: "Adds or updates a permission policy for a Amazon Q Business application, allowing cross-account access for an ISV.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_associatePermissionCmd).Standalone()

	qbusiness_associatePermissionCmd.Flags().String("actions", "", "The list of Amazon Q Business actions that the ISV is allowed to perform.")
	qbusiness_associatePermissionCmd.Flags().String("application-id", "", "The unique identifier of the Amazon Q Business application.")
	qbusiness_associatePermissionCmd.Flags().String("conditions", "", "The conditions that restrict when the permission is effective.")
	qbusiness_associatePermissionCmd.Flags().String("principal", "", "The Amazon Resource Name of the IAM role for the ISV that is being granted permission.")
	qbusiness_associatePermissionCmd.Flags().String("statement-id", "", "A unique identifier for the policy statement.")
	qbusiness_associatePermissionCmd.MarkFlagRequired("actions")
	qbusiness_associatePermissionCmd.MarkFlagRequired("application-id")
	qbusiness_associatePermissionCmd.MarkFlagRequired("principal")
	qbusiness_associatePermissionCmd.MarkFlagRequired("statement-id")
	qbusinessCmd.AddCommand(qbusiness_associatePermissionCmd)
}
