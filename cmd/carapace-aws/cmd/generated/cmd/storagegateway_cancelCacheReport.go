package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_cancelCacheReportCmd = &cobra.Command{
	Use:   "cancel-cache-report",
	Short: "Cancels generation of a specified cache report.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_cancelCacheReportCmd).Standalone()

	storagegateway_cancelCacheReportCmd.Flags().String("cache-report-arn", "", "The Amazon Resource Name (ARN) of the cache report you want to cancel.")
	storagegateway_cancelCacheReportCmd.MarkFlagRequired("cache-report-arn")
	storagegatewayCmd.AddCommand(storagegateway_cancelCacheReportCmd)
}
