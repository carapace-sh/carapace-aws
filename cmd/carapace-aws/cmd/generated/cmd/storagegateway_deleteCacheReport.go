package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_deleteCacheReportCmd = &cobra.Command{
	Use:   "delete-cache-report",
	Short: "Deletes the specified cache report and any associated tags from the Storage Gateway database.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_deleteCacheReportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_deleteCacheReportCmd).Standalone()

		storagegateway_deleteCacheReportCmd.Flags().String("cache-report-arn", "", "The Amazon Resource Name (ARN) of the cache report you want to delete.")
		storagegateway_deleteCacheReportCmd.MarkFlagRequired("cache-report-arn")
	})
	storagegatewayCmd.AddCommand(storagegateway_deleteCacheReportCmd)
}
