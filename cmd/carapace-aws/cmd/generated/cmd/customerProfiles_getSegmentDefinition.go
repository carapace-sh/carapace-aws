package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_getSegmentDefinitionCmd = &cobra.Command{
	Use:   "get-segment-definition",
	Short: "Gets a segment definition from the domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_getSegmentDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_getSegmentDefinitionCmd).Standalone()

		customerProfiles_getSegmentDefinitionCmd.Flags().String("domain-name", "", "The unique name of the domain.")
		customerProfiles_getSegmentDefinitionCmd.Flags().String("segment-definition-name", "", "The unique name of the segment definition.")
		customerProfiles_getSegmentDefinitionCmd.MarkFlagRequired("domain-name")
		customerProfiles_getSegmentDefinitionCmd.MarkFlagRequired("segment-definition-name")
	})
	customerProfilesCmd.AddCommand(customerProfiles_getSegmentDefinitionCmd)
}
