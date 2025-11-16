package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_listCalculatedAttributesForProfileCmd = &cobra.Command{
	Use:   "list-calculated-attributes-for-profile",
	Short: "Retrieve a list of calculated attributes for a customer profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_listCalculatedAttributesForProfileCmd).Standalone()

	customerProfiles_listCalculatedAttributesForProfileCmd.Flags().String("domain-name", "", "The unique name of the domain.")
	customerProfiles_listCalculatedAttributesForProfileCmd.Flags().String("max-results", "", "The maximum number of calculated attributes returned per page.")
	customerProfiles_listCalculatedAttributesForProfileCmd.Flags().String("next-token", "", "The pagination token from the previous call to ListCalculatedAttributesForProfile.")
	customerProfiles_listCalculatedAttributesForProfileCmd.Flags().String("profile-id", "", "The unique identifier of a customer profile.")
	customerProfiles_listCalculatedAttributesForProfileCmd.MarkFlagRequired("domain-name")
	customerProfiles_listCalculatedAttributesForProfileCmd.MarkFlagRequired("profile-id")
	customerProfilesCmd.AddCommand(customerProfiles_listCalculatedAttributesForProfileCmd)
}
