package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var detective_enableOrganizationAdminAccountCmd = &cobra.Command{
	Use:   "enable-organization-admin-account",
	Short: "Designates the Detective administrator account for the organization in the current Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(detective_enableOrganizationAdminAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(detective_enableOrganizationAdminAccountCmd).Standalone()

		detective_enableOrganizationAdminAccountCmd.Flags().String("account-id", "", "The Amazon Web Services account identifier of the account to designate as the Detective administrator account for the organization.")
		detective_enableOrganizationAdminAccountCmd.MarkFlagRequired("account-id")
	})
	detectiveCmd.AddCommand(detective_enableOrganizationAdminAccountCmd)
}
