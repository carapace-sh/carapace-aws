package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_getCalculatedAttributeForProfileCmd = &cobra.Command{
	Use:   "get-calculated-attribute-for-profile",
	Short: "Retrieve a calculated attribute for a customer profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_getCalculatedAttributeForProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_getCalculatedAttributeForProfileCmd).Standalone()

		customerProfiles_getCalculatedAttributeForProfileCmd.Flags().String("calculated-attribute-name", "", "The unique name of the calculated attribute.")
		customerProfiles_getCalculatedAttributeForProfileCmd.Flags().String("domain-name", "", "The unique name of the domain.")
		customerProfiles_getCalculatedAttributeForProfileCmd.Flags().String("profile-id", "", "The unique identifier of a customer profile.")
		customerProfiles_getCalculatedAttributeForProfileCmd.MarkFlagRequired("calculated-attribute-name")
		customerProfiles_getCalculatedAttributeForProfileCmd.MarkFlagRequired("domain-name")
		customerProfiles_getCalculatedAttributeForProfileCmd.MarkFlagRequired("profile-id")
	})
	customerProfilesCmd.AddCommand(customerProfiles_getCalculatedAttributeForProfileCmd)
}
