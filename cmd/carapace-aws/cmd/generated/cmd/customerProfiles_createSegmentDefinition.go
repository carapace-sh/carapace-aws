package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_createSegmentDefinitionCmd = &cobra.Command{
	Use:   "create-segment-definition",
	Short: "Creates a segment definition associated to the given domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_createSegmentDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_createSegmentDefinitionCmd).Standalone()

		customerProfiles_createSegmentDefinitionCmd.Flags().String("description", "", "The description of the segment definition.")
		customerProfiles_createSegmentDefinitionCmd.Flags().String("display-name", "", "The display name of the segment definition.")
		customerProfiles_createSegmentDefinitionCmd.Flags().String("domain-name", "", "The unique name of the domain.")
		customerProfiles_createSegmentDefinitionCmd.Flags().String("segment-definition-name", "", "The unique name of the segment definition.")
		customerProfiles_createSegmentDefinitionCmd.Flags().String("segment-groups", "", "Specifies the base segments and dimensions for a segment definition along with their respective relationship.")
		customerProfiles_createSegmentDefinitionCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
		customerProfiles_createSegmentDefinitionCmd.MarkFlagRequired("display-name")
		customerProfiles_createSegmentDefinitionCmd.MarkFlagRequired("domain-name")
		customerProfiles_createSegmentDefinitionCmd.MarkFlagRequired("segment-definition-name")
		customerProfiles_createSegmentDefinitionCmd.MarkFlagRequired("segment-groups")
	})
	customerProfilesCmd.AddCommand(customerProfiles_createSegmentDefinitionCmd)
}
