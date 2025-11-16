package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var marketplacecommerceanalytics_startSupportDataExportCmd = &cobra.Command{
	Use:   "start-support-data-export",
	Short: "*This target has been deprecated.* Given a data set type and a from date, asynchronously publishes the requested customer support data to the specified S3 bucket and notifies the specified SNS topic once the data is available.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(marketplacecommerceanalytics_startSupportDataExportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(marketplacecommerceanalytics_startSupportDataExportCmd).Standalone()

		marketplacecommerceanalytics_startSupportDataExportCmd.Flags().String("customer-defined-values", "", "*This target has been deprecated.* (Optional) Key-value pairs which will be returned, unmodified, in the Amazon SNS notification message and the data set metadata file.")
		marketplacecommerceanalytics_startSupportDataExportCmd.Flags().String("data-set-type", "", "*This target has been deprecated.* Specifies the data set type to be written to the output csv file.")
		marketplacecommerceanalytics_startSupportDataExportCmd.Flags().String("destination-s3-bucket-name", "", "*This target has been deprecated.* The name (friendly name, not ARN) of the destination S3 bucket.")
		marketplacecommerceanalytics_startSupportDataExportCmd.Flags().String("destination-s3-prefix", "", "*This target has been deprecated.* (Optional) The desired S3 prefix for the published data set, similar to a directory path in standard file systems.")
		marketplacecommerceanalytics_startSupportDataExportCmd.Flags().String("from-date", "", "*This target has been deprecated.* The start date from which to retrieve the data set in UTC.")
		marketplacecommerceanalytics_startSupportDataExportCmd.Flags().String("role-name-arn", "", "*This target has been deprecated.* The Amazon Resource Name (ARN) of the Role with an attached permissions policy to interact with the provided AWS services.")
		marketplacecommerceanalytics_startSupportDataExportCmd.Flags().String("sns-topic-arn", "", "*This target has been deprecated.* Amazon Resource Name (ARN) for the SNS Topic that will be notified when the data set has been published or if an error has occurred.")
		marketplacecommerceanalytics_startSupportDataExportCmd.MarkFlagRequired("data-set-type")
		marketplacecommerceanalytics_startSupportDataExportCmd.MarkFlagRequired("destination-s3-bucket-name")
		marketplacecommerceanalytics_startSupportDataExportCmd.MarkFlagRequired("from-date")
		marketplacecommerceanalytics_startSupportDataExportCmd.MarkFlagRequired("role-name-arn")
		marketplacecommerceanalytics_startSupportDataExportCmd.MarkFlagRequired("sns-topic-arn")
	})
	marketplacecommerceanalyticsCmd.AddCommand(marketplacecommerceanalytics_startSupportDataExportCmd)
}
