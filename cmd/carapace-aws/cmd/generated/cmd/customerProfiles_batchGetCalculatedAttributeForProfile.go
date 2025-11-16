package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_batchGetCalculatedAttributeForProfileCmd = &cobra.Command{
	Use:   "batch-get-calculated-attribute-for-profile",
	Short: "Fetch the possible attribute values given the attribute name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_batchGetCalculatedAttributeForProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_batchGetCalculatedAttributeForProfileCmd).Standalone()

		customerProfiles_batchGetCalculatedAttributeForProfileCmd.Flags().String("calculated-attribute-name", "", "The unique name of the calculated attribute.")
		customerProfiles_batchGetCalculatedAttributeForProfileCmd.Flags().String("condition-overrides", "", "Overrides the condition block within the original calculated attribute definition.")
		customerProfiles_batchGetCalculatedAttributeForProfileCmd.Flags().String("domain-name", "", "The unique name of the domain.")
		customerProfiles_batchGetCalculatedAttributeForProfileCmd.Flags().String("profile-ids", "", "List of unique identifiers for customer profiles to retrieve.")
		customerProfiles_batchGetCalculatedAttributeForProfileCmd.MarkFlagRequired("calculated-attribute-name")
		customerProfiles_batchGetCalculatedAttributeForProfileCmd.MarkFlagRequired("domain-name")
		customerProfiles_batchGetCalculatedAttributeForProfileCmd.MarkFlagRequired("profile-ids")
	})
	customerProfilesCmd.AddCommand(customerProfiles_batchGetCalculatedAttributeForProfileCmd)
}
