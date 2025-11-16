package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_deleteCalculatedAttributeDefinitionCmd = &cobra.Command{
	Use:   "delete-calculated-attribute-definition",
	Short: "Deletes an existing calculated attribute definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_deleteCalculatedAttributeDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_deleteCalculatedAttributeDefinitionCmd).Standalone()

		customerProfiles_deleteCalculatedAttributeDefinitionCmd.Flags().String("calculated-attribute-name", "", "The unique name of the calculated attribute.")
		customerProfiles_deleteCalculatedAttributeDefinitionCmd.Flags().String("domain-name", "", "The unique name of the domain.")
		customerProfiles_deleteCalculatedAttributeDefinitionCmd.MarkFlagRequired("calculated-attribute-name")
		customerProfiles_deleteCalculatedAttributeDefinitionCmd.MarkFlagRequired("domain-name")
	})
	customerProfilesCmd.AddCommand(customerProfiles_deleteCalculatedAttributeDefinitionCmd)
}
