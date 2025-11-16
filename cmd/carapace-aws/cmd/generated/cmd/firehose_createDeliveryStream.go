package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var firehose_createDeliveryStreamCmd = &cobra.Command{
	Use:   "create-delivery-stream",
	Short: "Creates a Firehose stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(firehose_createDeliveryStreamCmd).Standalone()

	firehose_createDeliveryStreamCmd.Flags().String("amazon-open-search-serverless-destination-configuration", "", "The destination in the Serverless offering for Amazon OpenSearch Service.")
	firehose_createDeliveryStreamCmd.Flags().String("amazonopensearchservice-destination-configuration", "", "The destination in Amazon OpenSearch Service.")
	firehose_createDeliveryStreamCmd.Flags().String("database-source-configuration", "", "The top level object for configuring streams with database as a source.")
	firehose_createDeliveryStreamCmd.Flags().String("delivery-stream-encryption-configuration-input", "", "Used to specify the type and Amazon Resource Name (ARN) of the KMS key needed for Server-Side Encryption (SSE).")
	firehose_createDeliveryStreamCmd.Flags().String("delivery-stream-name", "", "The name of the Firehose stream.")
	firehose_createDeliveryStreamCmd.Flags().String("delivery-stream-type", "", "The Firehose stream type.")
	firehose_createDeliveryStreamCmd.Flags().String("direct-put-source-configuration", "", "The structure that configures parameters such as `ThroughputHintInMBs` for a stream configured with Direct PUT as a source.")
	firehose_createDeliveryStreamCmd.Flags().String("elasticsearch-destination-configuration", "", "The destination in Amazon OpenSearch Service.")
	firehose_createDeliveryStreamCmd.Flags().String("extended-s3-destination-configuration", "", "The destination in Amazon S3.")
	firehose_createDeliveryStreamCmd.Flags().String("http-endpoint-destination-configuration", "", "Enables configuring Kinesis Firehose to deliver data to any HTTP endpoint destination.")
	firehose_createDeliveryStreamCmd.Flags().String("iceberg-destination-configuration", "", "Configure Apache Iceberg Tables destination.")
	firehose_createDeliveryStreamCmd.Flags().String("kinesis-stream-source-configuration", "", "When a Kinesis data stream is used as the source for the Firehose stream, a [KinesisStreamSourceConfiguration]() containing the Kinesis data stream Amazon Resource Name (ARN) and the role ARN for the source stream.")
	firehose_createDeliveryStreamCmd.Flags().String("msksource-configuration", "", "")
	firehose_createDeliveryStreamCmd.Flags().String("redshift-destination-configuration", "", "The destination in Amazon Redshift.")
	firehose_createDeliveryStreamCmd.Flags().String("s3-destination-configuration", "", "\\[Deprecated] The destination in Amazon S3.")
	firehose_createDeliveryStreamCmd.Flags().String("snowflake-destination-configuration", "", "Configure Snowflake destination")
	firehose_createDeliveryStreamCmd.Flags().String("splunk-destination-configuration", "", "The destination in Splunk.")
	firehose_createDeliveryStreamCmd.Flags().String("tags", "", "A set of tags to assign to the Firehose stream.")
	firehose_createDeliveryStreamCmd.MarkFlagRequired("delivery-stream-name")
	firehoseCmd.AddCommand(firehose_createDeliveryStreamCmd)
}
