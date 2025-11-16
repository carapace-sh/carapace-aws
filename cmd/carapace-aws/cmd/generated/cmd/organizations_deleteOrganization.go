package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_deleteOrganizationCmd = &cobra.Command{
	Use:   "delete-organization",
	Short: "Deletes the organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_deleteOrganizationCmd).Standalone()

	organizationsCmd.AddCommand(organizations_deleteOrganizationCmd)
}
