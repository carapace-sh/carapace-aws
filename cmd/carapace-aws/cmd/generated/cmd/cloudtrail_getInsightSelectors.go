package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_getInsightSelectorsCmd = &cobra.Command{
	Use:   "get-insight-selectors",
	Short: "Describes the settings for the Insights event selectors that you configured for your trail or event data store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_getInsightSelectorsCmd).Standalone()

	cloudtrail_getInsightSelectorsCmd.Flags().String("event-data-store", "", "Specifies the ARN (or ID suffix of the ARN) of the event data store for which you want to get Insights selectors.")
	cloudtrail_getInsightSelectorsCmd.Flags().String("trail-name", "", "Specifies the name of the trail or trail ARN.")
	cloudtrailCmd.AddCommand(cloudtrail_getInsightSelectorsCmd)
}
