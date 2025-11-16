package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_searchProfilesCmd = &cobra.Command{
	Use:   "search-profiles",
	Short: "Searches for profiles within a specific domain using one or more predefined search keys (e.g., \\_fullName, \\_phone, \\_email, \\_account, etc.) and/or custom-defined search keys.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_searchProfilesCmd).Standalone()

	customerProfiles_searchProfilesCmd.Flags().String("additional-search-keys", "", "A list of `AdditionalSearchKey` objects that are each searchable identifiers of a profile.")
	customerProfiles_searchProfilesCmd.Flags().String("domain-name", "", "The unique name of the domain.")
	customerProfiles_searchProfilesCmd.Flags().String("key-name", "", "A searchable identifier of a customer profile.")
	customerProfiles_searchProfilesCmd.Flags().String("logical-operator", "", "Relationship between all specified search keys that will be used to search for profiles.")
	customerProfiles_searchProfilesCmd.Flags().String("max-results", "", "The maximum number of objects returned per page.")
	customerProfiles_searchProfilesCmd.Flags().String("next-token", "", "The pagination token from the previous SearchProfiles API call.")
	customerProfiles_searchProfilesCmd.Flags().String("values", "", "A list of key values.")
	customerProfiles_searchProfilesCmd.MarkFlagRequired("domain-name")
	customerProfiles_searchProfilesCmd.MarkFlagRequired("key-name")
	customerProfiles_searchProfilesCmd.MarkFlagRequired("values")
	customerProfilesCmd.AddCommand(customerProfiles_searchProfilesCmd)
}
