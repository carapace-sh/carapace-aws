package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_getDomainCmd = &cobra.Command{
	Use:   "get-domain",
	Short: "Returns information about a specific domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_getDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_getDomainCmd).Standalone()

		customerProfiles_getDomainCmd.Flags().String("domain-name", "", "The unique name of the domain.")
		customerProfiles_getDomainCmd.MarkFlagRequired("domain-name")
	})
	customerProfilesCmd.AddCommand(customerProfiles_getDomainCmd)
}
