package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_listDomainsCmd = &cobra.Command{
	Use:   "list-domains",
	Short: "Returns a list of all the domains for an AWS account that have been created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_listDomainsCmd).Standalone()

	customerProfiles_listDomainsCmd.Flags().String("max-results", "", "The maximum number of objects returned per page.")
	customerProfiles_listDomainsCmd.Flags().String("next-token", "", "The pagination token from the previous ListDomain API call.")
	customerProfilesCmd.AddCommand(customerProfiles_listDomainsCmd)
}
