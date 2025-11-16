package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_deleteOrganizationalUnitCmd = &cobra.Command{
	Use:   "delete-organizational-unit",
	Short: "Deletes an organizational unit (OU) from a root or another OU.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_deleteOrganizationalUnitCmd).Standalone()

	organizations_deleteOrganizationalUnitCmd.Flags().String("organizational-unit-id", "", "The unique identifier (ID) of the organizational unit that you want to delete.")
	organizations_deleteOrganizationalUnitCmd.MarkFlagRequired("organizational-unit-id")
	organizationsCmd.AddCommand(organizations_deleteOrganizationalUnitCmd)
}
