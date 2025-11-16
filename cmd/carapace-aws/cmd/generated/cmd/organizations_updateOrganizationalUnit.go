package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_updateOrganizationalUnitCmd = &cobra.Command{
	Use:   "update-organizational-unit",
	Short: "Renames the specified organizational unit (OU).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_updateOrganizationalUnitCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(organizations_updateOrganizationalUnitCmd).Standalone()

		organizations_updateOrganizationalUnitCmd.Flags().String("name", "", "The new name that you want to assign to the OU.")
		organizations_updateOrganizationalUnitCmd.Flags().String("organizational-unit-id", "", "The unique identifier (ID) of the OU that you want to rename.")
		organizations_updateOrganizationalUnitCmd.MarkFlagRequired("organizational-unit-id")
	})
	organizationsCmd.AddCommand(organizations_updateOrganizationalUnitCmd)
}
