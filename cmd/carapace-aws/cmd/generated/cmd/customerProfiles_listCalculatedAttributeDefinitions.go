package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_listCalculatedAttributeDefinitionsCmd = &cobra.Command{
	Use:   "list-calculated-attribute-definitions",
	Short: "Lists calculated attribute definitions for Customer Profiles",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_listCalculatedAttributeDefinitionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_listCalculatedAttributeDefinitionsCmd).Standalone()

		customerProfiles_listCalculatedAttributeDefinitionsCmd.Flags().String("domain-name", "", "The unique name of the domain.")
		customerProfiles_listCalculatedAttributeDefinitionsCmd.Flags().String("max-results", "", "The maximum number of calculated attribute definitions returned per page.")
		customerProfiles_listCalculatedAttributeDefinitionsCmd.Flags().String("next-token", "", "The pagination token from the previous call to ListCalculatedAttributeDefinitions.")
		customerProfiles_listCalculatedAttributeDefinitionsCmd.MarkFlagRequired("domain-name")
	})
	customerProfilesCmd.AddCommand(customerProfiles_listCalculatedAttributeDefinitionsCmd)
}
