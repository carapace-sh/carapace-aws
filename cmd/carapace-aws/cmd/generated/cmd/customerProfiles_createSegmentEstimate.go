package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_createSegmentEstimateCmd = &cobra.Command{
	Use:   "create-segment-estimate",
	Short: "Creates a segment estimate query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_createSegmentEstimateCmd).Standalone()

	customerProfiles_createSegmentEstimateCmd.Flags().String("domain-name", "", "The unique name of the domain.")
	customerProfiles_createSegmentEstimateCmd.Flags().String("segment-query", "", "The segment query for calculating a segment estimate.")
	customerProfiles_createSegmentEstimateCmd.MarkFlagRequired("domain-name")
	customerProfiles_createSegmentEstimateCmd.MarkFlagRequired("segment-query")
	customerProfilesCmd.AddCommand(customerProfiles_createSegmentEstimateCmd)
}
