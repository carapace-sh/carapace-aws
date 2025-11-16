package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_startCacheReportCmd = &cobra.Command{
	Use:   "start-cache-report",
	Short: "Starts generating a report of the file metadata currently cached by an S3 File Gateway for a specific file share.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_startCacheReportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_startCacheReportCmd).Standalone()

		storagegateway_startCacheReportCmd.Flags().String("bucket-region", "", "The Amazon Web Services Region of the Amazon S3 bucket where you want to save the cache report.")
		storagegateway_startCacheReportCmd.Flags().String("client-token", "", "A unique identifier that you use to ensure idempotent report generation if you need to retry an unsuccessful `StartCacheReport` request.")
		storagegateway_startCacheReportCmd.Flags().String("exclusion-filters", "", "The list of filters and parameters that determine which files are excluded from the report.")
		storagegateway_startCacheReportCmd.Flags().String("file-share-arn", "", "")
		storagegateway_startCacheReportCmd.Flags().String("inclusion-filters", "", "The list of filters and parameters that determine which files are included in the report.")
		storagegateway_startCacheReportCmd.Flags().String("location-arn", "", "The ARN of the Amazon S3 bucket where you want to save the cache report.")
		storagegateway_startCacheReportCmd.Flags().String("role", "", "The ARN of the IAM role used when saving the cache report to Amazon S3.")
		storagegateway_startCacheReportCmd.Flags().String("tags", "", "A list of up to 50 key/value tags that you can assign to the cache report.")
		storagegateway_startCacheReportCmd.Flags().String("vpcendpoint-dnsname", "", "The DNS name of the VPC endpoint associated with the Amazon S3 where you want to save the cache report.")
		storagegateway_startCacheReportCmd.MarkFlagRequired("bucket-region")
		storagegateway_startCacheReportCmd.MarkFlagRequired("client-token")
		storagegateway_startCacheReportCmd.MarkFlagRequired("file-share-arn")
		storagegateway_startCacheReportCmd.MarkFlagRequired("location-arn")
		storagegateway_startCacheReportCmd.MarkFlagRequired("role")
	})
	storagegatewayCmd.AddCommand(storagegateway_startCacheReportCmd)
}
