package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_createCalculatedAttributeDefinitionCmd = &cobra.Command{
	Use:   "create-calculated-attribute-definition",
	Short: "Creates a new calculated attribute definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_createCalculatedAttributeDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_createCalculatedAttributeDefinitionCmd).Standalone()

		customerProfiles_createCalculatedAttributeDefinitionCmd.Flags().String("attribute-details", "", "Mathematical expression and a list of attribute items specified in that expression.")
		customerProfiles_createCalculatedAttributeDefinitionCmd.Flags().String("calculated-attribute-name", "", "The unique name of the calculated attribute.")
		customerProfiles_createCalculatedAttributeDefinitionCmd.Flags().String("conditions", "", "The conditions including range, object count, and threshold for the calculated attribute.")
		customerProfiles_createCalculatedAttributeDefinitionCmd.Flags().String("description", "", "The description of the calculated attribute.")
		customerProfiles_createCalculatedAttributeDefinitionCmd.Flags().String("display-name", "", "The display name of the calculated attribute.")
		customerProfiles_createCalculatedAttributeDefinitionCmd.Flags().String("domain-name", "", "The unique name of the domain.")
		customerProfiles_createCalculatedAttributeDefinitionCmd.Flags().String("filter", "", "Defines how to filter incoming objects to include part of the Calculated Attribute.")
		customerProfiles_createCalculatedAttributeDefinitionCmd.Flags().String("statistic", "", "The aggregation operation to perform for the calculated attribute.")
		customerProfiles_createCalculatedAttributeDefinitionCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
		customerProfiles_createCalculatedAttributeDefinitionCmd.Flags().String("use-historical-data", "", "Whether historical data ingested before the Calculated Attribute was created should be included in calculations.")
		customerProfiles_createCalculatedAttributeDefinitionCmd.MarkFlagRequired("attribute-details")
		customerProfiles_createCalculatedAttributeDefinitionCmd.MarkFlagRequired("calculated-attribute-name")
		customerProfiles_createCalculatedAttributeDefinitionCmd.MarkFlagRequired("domain-name")
		customerProfiles_createCalculatedAttributeDefinitionCmd.MarkFlagRequired("statistic")
	})
	customerProfilesCmd.AddCommand(customerProfiles_createCalculatedAttributeDefinitionCmd)
}
