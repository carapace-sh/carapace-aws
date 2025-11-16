package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_updateFindingAggregatorCmd = &cobra.Command{
	Use:   "update-finding-aggregator",
	Short: "The *aggregation Region* is now called the *home Region*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_updateFindingAggregatorCmd).Standalone()

	securityhub_updateFindingAggregatorCmd.Flags().String("finding-aggregator-arn", "", "The ARN of the finding aggregator.")
	securityhub_updateFindingAggregatorCmd.Flags().String("region-linking-mode", "", "Indicates whether to aggregate findings from all of the available Regions in the current partition.")
	securityhub_updateFindingAggregatorCmd.Flags().String("regions", "", "If `RegionLinkingMode` is `ALL_REGIONS_EXCEPT_SPECIFIED`, then this is a space-separated list of Regions that don't replicate and send findings to the home Region.")
	securityhub_updateFindingAggregatorCmd.MarkFlagRequired("finding-aggregator-arn")
	securityhub_updateFindingAggregatorCmd.MarkFlagRequired("region-linking-mode")
	securityhubCmd.AddCommand(securityhub_updateFindingAggregatorCmd)
}
