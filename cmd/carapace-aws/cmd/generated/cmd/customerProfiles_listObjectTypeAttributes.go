package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_listObjectTypeAttributesCmd = &cobra.Command{
	Use:   "list-object-type-attributes",
	Short: "Fetch the possible attribute values given the attribute name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_listObjectTypeAttributesCmd).Standalone()

	customerProfiles_listObjectTypeAttributesCmd.Flags().String("domain-name", "", "The unique identifier of the domain.")
	customerProfiles_listObjectTypeAttributesCmd.Flags().String("max-results", "", "The maximum number of objects returned per page.")
	customerProfiles_listObjectTypeAttributesCmd.Flags().String("next-token", "", "The pagination token from the previous call.")
	customerProfiles_listObjectTypeAttributesCmd.Flags().String("object-type-name", "", "The name of the profile object type.")
	customerProfiles_listObjectTypeAttributesCmd.MarkFlagRequired("domain-name")
	customerProfiles_listObjectTypeAttributesCmd.MarkFlagRequired("object-type-name")
	customerProfilesCmd.AddCommand(customerProfiles_listObjectTypeAttributesCmd)
}
