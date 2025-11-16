package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_describeUsageReportSubscriptionsCmd = &cobra.Command{
	Use:   "describe-usage-report-subscriptions",
	Short: "Retrieves a list that describes one or more usage report subscriptions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_describeUsageReportSubscriptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_describeUsageReportSubscriptionsCmd).Standalone()

		appstream_describeUsageReportSubscriptionsCmd.Flags().String("max-results", "", "The maximum size of each page of results.")
		appstream_describeUsageReportSubscriptionsCmd.Flags().String("next-token", "", "The pagination token to use to retrieve the next page of results for this operation.")
	})
	appstreamCmd.AddCommand(appstream_describeUsageReportSubscriptionsCmd)
}
