package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_deleteSegmentDefinitionCmd = &cobra.Command{
	Use:   "delete-segment-definition",
	Short: "Deletes a segment definition from the domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_deleteSegmentDefinitionCmd).Standalone()

	customerProfiles_deleteSegmentDefinitionCmd.Flags().String("domain-name", "", "The unique name of the domain.")
	customerProfiles_deleteSegmentDefinitionCmd.Flags().String("segment-definition-name", "", "The unique name of the segment definition.")
	customerProfiles_deleteSegmentDefinitionCmd.MarkFlagRequired("domain-name")
	customerProfiles_deleteSegmentDefinitionCmd.MarkFlagRequired("segment-definition-name")
	customerProfilesCmd.AddCommand(customerProfiles_deleteSegmentDefinitionCmd)
}
