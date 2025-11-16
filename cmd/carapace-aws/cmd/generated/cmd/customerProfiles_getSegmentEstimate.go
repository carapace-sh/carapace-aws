package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_getSegmentEstimateCmd = &cobra.Command{
	Use:   "get-segment-estimate",
	Short: "Gets the result of a segment estimate query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_getSegmentEstimateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_getSegmentEstimateCmd).Standalone()

		customerProfiles_getSegmentEstimateCmd.Flags().String("domain-name", "", "The unique name of the domain.")
		customerProfiles_getSegmentEstimateCmd.Flags().String("estimate-id", "", "The query Id passed by a previous `CreateSegmentEstimate` operation.")
		customerProfiles_getSegmentEstimateCmd.MarkFlagRequired("domain-name")
		customerProfiles_getSegmentEstimateCmd.MarkFlagRequired("estimate-id")
	})
	customerProfilesCmd.AddCommand(customerProfiles_getSegmentEstimateCmd)
}
