package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_removeAccountFromOrganizationCmd = &cobra.Command{
	Use:   "remove-account-from-organization",
	Short: "Removes the specified account from the organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_removeAccountFromOrganizationCmd).Standalone()

	organizations_removeAccountFromOrganizationCmd.Flags().String("account-id", "", "The unique identifier (ID) of the member account that you want to remove from the organization.")
	organizations_removeAccountFromOrganizationCmd.MarkFlagRequired("account-id")
	organizationsCmd.AddCommand(organizations_removeAccountFromOrganizationCmd)
}
