package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var detective_disableOrganizationAdminAccountCmd = &cobra.Command{
	Use:   "disable-organization-admin-account",
	Short: "Removes the Detective administrator account in the current Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(detective_disableOrganizationAdminAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(detective_disableOrganizationAdminAccountCmd).Standalone()

	})
	detectiveCmd.AddCommand(detective_disableOrganizationAdminAccountCmd)
}
