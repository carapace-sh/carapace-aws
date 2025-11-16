package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_getDomainLayoutCmd = &cobra.Command{
	Use:   "get-domain-layout",
	Short: "Gets the layout to view data for a specific domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_getDomainLayoutCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_getDomainLayoutCmd).Standalone()

		customerProfiles_getDomainLayoutCmd.Flags().String("domain-name", "", "The unique name of the domain.")
		customerProfiles_getDomainLayoutCmd.Flags().String("layout-definition-name", "", "The unique name of the layout.")
		customerProfiles_getDomainLayoutCmd.MarkFlagRequired("domain-name")
		customerProfiles_getDomainLayoutCmd.MarkFlagRequired("layout-definition-name")
	})
	customerProfilesCmd.AddCommand(customerProfiles_getDomainLayoutCmd)
}
