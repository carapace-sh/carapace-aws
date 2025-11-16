package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_deleteFindingAggregatorCmd = &cobra.Command{
	Use:   "delete-finding-aggregator",
	Short: "The *aggregation Region* is now called the *home Region*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_deleteFindingAggregatorCmd).Standalone()

	securityhub_deleteFindingAggregatorCmd.Flags().String("finding-aggregator-arn", "", "The ARN of the finding aggregator to delete.")
	securityhub_deleteFindingAggregatorCmd.MarkFlagRequired("finding-aggregator-arn")
	securityhubCmd.AddCommand(securityhub_deleteFindingAggregatorCmd)
}
