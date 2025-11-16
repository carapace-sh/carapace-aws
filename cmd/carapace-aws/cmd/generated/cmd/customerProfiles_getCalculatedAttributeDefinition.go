package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_getCalculatedAttributeDefinitionCmd = &cobra.Command{
	Use:   "get-calculated-attribute-definition",
	Short: "Provides more information on a calculated attribute definition for Customer Profiles.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_getCalculatedAttributeDefinitionCmd).Standalone()

	customerProfiles_getCalculatedAttributeDefinitionCmd.Flags().String("calculated-attribute-name", "", "The unique name of the calculated attribute.")
	customerProfiles_getCalculatedAttributeDefinitionCmd.Flags().String("domain-name", "", "The unique name of the domain.")
	customerProfiles_getCalculatedAttributeDefinitionCmd.MarkFlagRequired("calculated-attribute-name")
	customerProfiles_getCalculatedAttributeDefinitionCmd.MarkFlagRequired("domain-name")
	customerProfilesCmd.AddCommand(customerProfiles_getCalculatedAttributeDefinitionCmd)
}
