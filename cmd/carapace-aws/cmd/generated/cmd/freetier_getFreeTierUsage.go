package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var freetier_getFreeTierUsageCmd = &cobra.Command{
	Use:   "get-free-tier-usage",
	Short: "Returns a list of all Free Tier usage objects that match your filters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(freetier_getFreeTierUsageCmd).Standalone()

	freetier_getFreeTierUsageCmd.Flags().String("filter", "", "An expression that specifies the conditions that you want each `FreeTierUsage` object to meet.")
	freetier_getFreeTierUsageCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
	freetier_getFreeTierUsageCmd.Flags().String("next-token", "", "The pagination token that indicates the next set of results to retrieve.")
	freetierCmd.AddCommand(freetier_getFreeTierUsageCmd)
}
