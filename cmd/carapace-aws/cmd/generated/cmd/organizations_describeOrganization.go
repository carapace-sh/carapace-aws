package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_describeOrganizationCmd = &cobra.Command{
	Use:   "describe-organization",
	Short: "Retrieves information about the organization that the user's account belongs to.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_describeOrganizationCmd).Standalone()

	organizationsCmd.AddCommand(organizations_describeOrganizationCmd)
}
