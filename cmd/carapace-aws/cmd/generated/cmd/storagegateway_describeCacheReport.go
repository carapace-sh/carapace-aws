package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_describeCacheReportCmd = &cobra.Command{
	Use:   "describe-cache-report",
	Short: "Returns information about the specified cache report, including completion status and generation progress.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_describeCacheReportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_describeCacheReportCmd).Standalone()

		storagegateway_describeCacheReportCmd.Flags().String("cache-report-arn", "", "The Amazon Resource Name (ARN) of the cache report you want to describe.")
		storagegateway_describeCacheReportCmd.MarkFlagRequired("cache-report-arn")
	})
	storagegatewayCmd.AddCommand(storagegateway_describeCacheReportCmd)
}
