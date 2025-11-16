package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_getCoverageStatisticsCmd = &cobra.Command{
	Use:   "get-coverage-statistics",
	Short: "Retrieves aggregated statistics for your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_getCoverageStatisticsCmd).Standalone()

	guardduty_getCoverageStatisticsCmd.Flags().String("detector-id", "", "The unique ID of the GuardDuty detector.")
	guardduty_getCoverageStatisticsCmd.Flags().String("filter-criteria", "", "Represents the criteria used to filter the coverage statistics.")
	guardduty_getCoverageStatisticsCmd.Flags().String("statistics-type", "", "Represents the statistics type used to aggregate the coverage details.")
	guardduty_getCoverageStatisticsCmd.MarkFlagRequired("detector-id")
	guardduty_getCoverageStatisticsCmd.MarkFlagRequired("statistics-type")
	guarddutyCmd.AddCommand(guardduty_getCoverageStatisticsCmd)
}
