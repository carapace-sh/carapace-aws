package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_listDomainLayoutsCmd = &cobra.Command{
	Use:   "list-domain-layouts",
	Short: "Lists the existing layouts that can be used to view data for a specific domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_listDomainLayoutsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_listDomainLayoutsCmd).Standalone()

		customerProfiles_listDomainLayoutsCmd.Flags().String("domain-name", "", "The unique name of the domain.")
		customerProfiles_listDomainLayoutsCmd.Flags().String("max-results", "", "The maximum number of objects returned per page.")
		customerProfiles_listDomainLayoutsCmd.Flags().String("next-token", "", "Identifies the next page of results to return.")
		customerProfiles_listDomainLayoutsCmd.MarkFlagRequired("domain-name")
	})
	customerProfilesCmd.AddCommand(customerProfiles_listDomainLayoutsCmd)
}
