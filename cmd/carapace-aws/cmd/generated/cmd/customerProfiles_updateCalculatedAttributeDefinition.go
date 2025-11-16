package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_updateCalculatedAttributeDefinitionCmd = &cobra.Command{
	Use:   "update-calculated-attribute-definition",
	Short: "Updates an existing calculated attribute definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_updateCalculatedAttributeDefinitionCmd).Standalone()

	customerProfiles_updateCalculatedAttributeDefinitionCmd.Flags().String("calculated-attribute-name", "", "The unique name of the calculated attribute.")
	customerProfiles_updateCalculatedAttributeDefinitionCmd.Flags().String("conditions", "", "The conditions including range, object count, and threshold for the calculated attribute.")
	customerProfiles_updateCalculatedAttributeDefinitionCmd.Flags().String("description", "", "The description of the calculated attribute.")
	customerProfiles_updateCalculatedAttributeDefinitionCmd.Flags().String("display-name", "", "The display name of the calculated attribute.")
	customerProfiles_updateCalculatedAttributeDefinitionCmd.Flags().String("domain-name", "", "The unique name of the domain.")
	customerProfiles_updateCalculatedAttributeDefinitionCmd.MarkFlagRequired("calculated-attribute-name")
	customerProfiles_updateCalculatedAttributeDefinitionCmd.MarkFlagRequired("domain-name")
	customerProfilesCmd.AddCommand(customerProfiles_updateCalculatedAttributeDefinitionCmd)
}
