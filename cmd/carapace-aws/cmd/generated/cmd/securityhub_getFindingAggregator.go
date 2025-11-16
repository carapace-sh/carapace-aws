package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_getFindingAggregatorCmd = &cobra.Command{
	Use:   "get-finding-aggregator",
	Short: "The *aggregation Region* is now called the *home Region*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_getFindingAggregatorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_getFindingAggregatorCmd).Standalone()

		securityhub_getFindingAggregatorCmd.Flags().String("finding-aggregator-arn", "", "The ARN of the finding aggregator to return details for.")
		securityhub_getFindingAggregatorCmd.MarkFlagRequired("finding-aggregator-arn")
	})
	securityhubCmd.AddCommand(securityhub_getFindingAggregatorCmd)
}
