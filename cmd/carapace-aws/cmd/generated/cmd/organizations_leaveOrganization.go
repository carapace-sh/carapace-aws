package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_leaveOrganizationCmd = &cobra.Command{
	Use:   "leave-organization",
	Short: "Removes a member account from its parent organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_leaveOrganizationCmd).Standalone()

	organizationsCmd.AddCommand(organizations_leaveOrganizationCmd)
}
