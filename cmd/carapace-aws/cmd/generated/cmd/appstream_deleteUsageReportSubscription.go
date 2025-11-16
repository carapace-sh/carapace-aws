package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_deleteUsageReportSubscriptionCmd = &cobra.Command{
	Use:   "delete-usage-report-subscription",
	Short: "Disables usage report generation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_deleteUsageReportSubscriptionCmd).Standalone()

	appstreamCmd.AddCommand(appstream_deleteUsageReportSubscriptionCmd)
}
