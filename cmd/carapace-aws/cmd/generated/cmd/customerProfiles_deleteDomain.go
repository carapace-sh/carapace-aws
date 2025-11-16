package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_deleteDomainCmd = &cobra.Command{
	Use:   "delete-domain",
	Short: "Deletes a specific domain and all of its customer data, such as customer profile attributes and their related objects.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_deleteDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_deleteDomainCmd).Standalone()

		customerProfiles_deleteDomainCmd.Flags().String("domain-name", "", "The unique name of the domain.")
		customerProfiles_deleteDomainCmd.MarkFlagRequired("domain-name")
	})
	customerProfilesCmd.AddCommand(customerProfiles_deleteDomainCmd)
}
