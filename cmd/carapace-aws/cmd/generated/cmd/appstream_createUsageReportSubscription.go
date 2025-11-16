package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_createUsageReportSubscriptionCmd = &cobra.Command{
	Use:   "create-usage-report-subscription",
	Short: "Creates a usage report subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_createUsageReportSubscriptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_createUsageReportSubscriptionCmd).Standalone()

	})
	appstreamCmd.AddCommand(appstream_createUsageReportSubscriptionCmd)
}
