package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_describeOrganizationalUnitCmd = &cobra.Command{
	Use:   "describe-organizational-unit",
	Short: "Retrieves information about an organizational unit (OU).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_describeOrganizationalUnitCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(organizations_describeOrganizationalUnitCmd).Standalone()

		organizations_describeOrganizationalUnitCmd.Flags().String("organizational-unit-id", "", "The unique identifier (ID) of the organizational unit that you want details about.")
		organizations_describeOrganizationalUnitCmd.MarkFlagRequired("organizational-unit-id")
	})
	organizationsCmd.AddCommand(organizations_describeOrganizationalUnitCmd)
}
