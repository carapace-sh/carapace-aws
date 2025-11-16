package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_putInsightSelectorsCmd = &cobra.Command{
	Use:   "put-insight-selectors",
	Short: "Lets you enable Insights event logging by specifying the Insights selectors that you want to enable on an existing trail or event data store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_putInsightSelectorsCmd).Standalone()

	cloudtrail_putInsightSelectorsCmd.Flags().String("event-data-store", "", "The ARN (or ID suffix of the ARN) of the source event data store for which you want to change or add Insights selectors.")
	cloudtrail_putInsightSelectorsCmd.Flags().String("insight-selectors", "", "A JSON string that contains the Insights types you want to log on a trail or event data store.")
	cloudtrail_putInsightSelectorsCmd.Flags().String("insights-destination", "", "The ARN (or ID suffix of the ARN) of the destination event data store that logs Insights events.")
	cloudtrail_putInsightSelectorsCmd.Flags().String("trail-name", "", "The name of the CloudTrail trail for which you want to change or add Insights selectors.")
	cloudtrail_putInsightSelectorsCmd.MarkFlagRequired("insight-selectors")
	cloudtrailCmd.AddCommand(cloudtrail_putInsightSelectorsCmd)
}
