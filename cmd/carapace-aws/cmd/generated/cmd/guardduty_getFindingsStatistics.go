package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_getFindingsStatisticsCmd = &cobra.Command{
	Use:   "get-findings-statistics",
	Short: "Lists GuardDuty findings statistics for the specified detector ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_getFindingsStatisticsCmd).Standalone()

	guardduty_getFindingsStatisticsCmd.Flags().String("detector-id", "", "The ID of the detector whose findings statistics you want to retrieve.")
	guardduty_getFindingsStatisticsCmd.Flags().String("finding-criteria", "", "Represents the criteria that is used for querying findings.")
	guardduty_getFindingsStatisticsCmd.Flags().String("finding-statistic-types", "", "The types of finding statistics to retrieve.")
	guardduty_getFindingsStatisticsCmd.Flags().String("group-by", "", "Displays the findings statistics grouped by one of the listed valid values.")
	guardduty_getFindingsStatisticsCmd.Flags().String("max-results", "", "The maximum number of results to be returned in the response.")
	guardduty_getFindingsStatisticsCmd.Flags().String("order-by", "", "Displays the sorted findings in the requested order.")
	guardduty_getFindingsStatisticsCmd.MarkFlagRequired("detector-id")
	guarddutyCmd.AddCommand(guardduty_getFindingsStatisticsCmd)
}
