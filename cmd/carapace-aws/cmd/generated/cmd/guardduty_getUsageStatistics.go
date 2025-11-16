package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_getUsageStatisticsCmd = &cobra.Command{
	Use:   "get-usage-statistics",
	Short: "Lists Amazon GuardDuty usage statistics over the last 30 days for the specified detector ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_getUsageStatisticsCmd).Standalone()

	guardduty_getUsageStatisticsCmd.Flags().String("detector-id", "", "The ID of the detector that specifies the GuardDuty service whose usage statistics you want to retrieve.")
	guardduty_getUsageStatisticsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
	guardduty_getUsageStatisticsCmd.Flags().String("next-token", "", "A token to use for paginating results that are returned in the response.")
	guardduty_getUsageStatisticsCmd.Flags().String("unit", "", "The currency unit you would like to view your usage statistics in.")
	guardduty_getUsageStatisticsCmd.Flags().String("usage-criteria", "", "Represents the criteria used for querying usage.")
	guardduty_getUsageStatisticsCmd.Flags().String("usage-statistic-type", "", "The type of usage statistics to retrieve.")
	guardduty_getUsageStatisticsCmd.MarkFlagRequired("detector-id")
	guardduty_getUsageStatisticsCmd.MarkFlagRequired("usage-criteria")
	guardduty_getUsageStatisticsCmd.MarkFlagRequired("usage-statistic-type")
	guarddutyCmd.AddCommand(guardduty_getUsageStatisticsCmd)
}
