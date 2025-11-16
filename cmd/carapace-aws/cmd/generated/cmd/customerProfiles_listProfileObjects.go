package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_listProfileObjectsCmd = &cobra.Command{
	Use:   "list-profile-objects",
	Short: "Returns a list of objects associated with a profile of a given ProfileObjectType.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_listProfileObjectsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_listProfileObjectsCmd).Standalone()

		customerProfiles_listProfileObjectsCmd.Flags().String("domain-name", "", "The unique name of the domain.")
		customerProfiles_listProfileObjectsCmd.Flags().String("max-results", "", "The maximum number of objects returned per page.")
		customerProfiles_listProfileObjectsCmd.Flags().String("next-token", "", "The pagination token from the previous call to ListProfileObjects.")
		customerProfiles_listProfileObjectsCmd.Flags().String("object-filter", "", "Applies a filter to the response to include profile objects with the specified index values.")
		customerProfiles_listProfileObjectsCmd.Flags().String("object-type-name", "", "The name of the profile object type.")
		customerProfiles_listProfileObjectsCmd.Flags().String("profile-id", "", "The unique identifier of a customer profile.")
		customerProfiles_listProfileObjectsCmd.MarkFlagRequired("domain-name")
		customerProfiles_listProfileObjectsCmd.MarkFlagRequired("object-type-name")
		customerProfiles_listProfileObjectsCmd.MarkFlagRequired("profile-id")
	})
	customerProfilesCmd.AddCommand(customerProfiles_listProfileObjectsCmd)
}
