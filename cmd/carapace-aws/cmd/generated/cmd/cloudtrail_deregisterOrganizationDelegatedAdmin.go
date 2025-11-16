package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_deregisterOrganizationDelegatedAdminCmd = &cobra.Command{
	Use:   "deregister-organization-delegated-admin",
	Short: "Removes CloudTrail delegated administrator permissions from a member account in an organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_deregisterOrganizationDelegatedAdminCmd).Standalone()

	cloudtrail_deregisterOrganizationDelegatedAdminCmd.Flags().String("delegated-admin-account-id", "", "A delegated administrator account ID.")
	cloudtrail_deregisterOrganizationDelegatedAdminCmd.MarkFlagRequired("delegated-admin-account-id")
	cloudtrailCmd.AddCommand(cloudtrail_deregisterOrganizationDelegatedAdminCmd)
}
