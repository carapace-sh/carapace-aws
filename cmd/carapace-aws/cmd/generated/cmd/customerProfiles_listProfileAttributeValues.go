package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_listProfileAttributeValuesCmd = &cobra.Command{
	Use:   "list-profile-attribute-values",
	Short: "Fetch the possible attribute values given the attribute name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_listProfileAttributeValuesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_listProfileAttributeValuesCmd).Standalone()

		customerProfiles_listProfileAttributeValuesCmd.Flags().String("attribute-name", "", "The attribute name.")
		customerProfiles_listProfileAttributeValuesCmd.Flags().String("domain-name", "", "The unique identifier of the domain.")
		customerProfiles_listProfileAttributeValuesCmd.MarkFlagRequired("attribute-name")
		customerProfiles_listProfileAttributeValuesCmd.MarkFlagRequired("domain-name")
	})
	customerProfilesCmd.AddCommand(customerProfiles_listProfileAttributeValuesCmd)
}
