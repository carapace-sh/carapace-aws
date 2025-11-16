package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var marketplacecommerceanalytics_generateDataSetCmd = &cobra.Command{
	Use:   "generate-data-set",
	Short: "Given a data set type and data set publication date, asynchronously publishes the requested data set to the specified S3 bucket and notifies the specified SNS topic once the data is available.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(marketplacecommerceanalytics_generateDataSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(marketplacecommerceanalytics_generateDataSetCmd).Standalone()

		marketplacecommerceanalytics_generateDataSetCmd.Flags().String("customer-defined-values", "", "(Optional) Key-value pairs which will be returned, unmodified, in the Amazon SNS notification message and the data set metadata file.")
		marketplacecommerceanalytics_generateDataSetCmd.Flags().String("data-set-publication-date", "", "The date a data set was published.")
		marketplacecommerceanalytics_generateDataSetCmd.Flags().String("data-set-type", "", "The desired data set type.")
		marketplacecommerceanalytics_generateDataSetCmd.Flags().String("destination-s3-bucket-name", "", "The name (friendly name, not ARN) of the destination S3 bucket.")
		marketplacecommerceanalytics_generateDataSetCmd.Flags().String("destination-s3-prefix", "", "(Optional) The desired S3 prefix for the published data set, similar to a directory path in standard file systems.")
		marketplacecommerceanalytics_generateDataSetCmd.Flags().String("role-name-arn", "", "The Amazon Resource Name (ARN) of the Role with an attached permissions policy to interact with the provided AWS services.")
		marketplacecommerceanalytics_generateDataSetCmd.Flags().String("sns-topic-arn", "", "Amazon Resource Name (ARN) for the SNS Topic that will be notified when the data set has been published or if an error has occurred.")
		marketplacecommerceanalytics_generateDataSetCmd.MarkFlagRequired("data-set-publication-date")
		marketplacecommerceanalytics_generateDataSetCmd.MarkFlagRequired("data-set-type")
		marketplacecommerceanalytics_generateDataSetCmd.MarkFlagRequired("destination-s3-bucket-name")
		marketplacecommerceanalytics_generateDataSetCmd.MarkFlagRequired("role-name-arn")
		marketplacecommerceanalytics_generateDataSetCmd.MarkFlagRequired("sns-topic-arn")
	})
	marketplacecommerceanalyticsCmd.AddCommand(marketplacecommerceanalytics_generateDataSetCmd)
}
