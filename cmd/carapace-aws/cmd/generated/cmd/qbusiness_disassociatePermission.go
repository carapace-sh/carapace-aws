package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_disassociatePermissionCmd = &cobra.Command{
	Use:   "disassociate-permission",
	Short: "Removes a permission policy from a Amazon Q Business application, revoking the cross-account access that was previously granted to an ISV.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_disassociatePermissionCmd).Standalone()

	qbusiness_disassociatePermissionCmd.Flags().String("application-id", "", "The unique identifier of the Amazon Q Business application.")
	qbusiness_disassociatePermissionCmd.Flags().String("statement-id", "", "The statement ID of the permission to remove.")
	qbusiness_disassociatePermissionCmd.MarkFlagRequired("application-id")
	qbusiness_disassociatePermissionCmd.MarkFlagRequired("statement-id")
	qbusinessCmd.AddCommand(qbusiness_disassociatePermissionCmd)
}
