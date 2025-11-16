package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_deleteDomainLayoutCmd = &cobra.Command{
	Use:   "delete-domain-layout",
	Short: "Deletes the layout used to view data for a specific domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_deleteDomainLayoutCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_deleteDomainLayoutCmd).Standalone()

		customerProfiles_deleteDomainLayoutCmd.Flags().String("domain-name", "", "The unique name of the domain.")
		customerProfiles_deleteDomainLayoutCmd.Flags().String("layout-definition-name", "", "The unique name of the layout.")
		customerProfiles_deleteDomainLayoutCmd.MarkFlagRequired("domain-name")
		customerProfiles_deleteDomainLayoutCmd.MarkFlagRequired("layout-definition-name")
	})
	customerProfilesCmd.AddCommand(customerProfiles_deleteDomainLayoutCmd)
}
