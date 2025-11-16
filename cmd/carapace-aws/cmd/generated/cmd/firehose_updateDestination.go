package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var firehose_updateDestinationCmd = &cobra.Command{
	Use:   "update-destination",
	Short: "Updates the specified destination of the specified Firehose stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(firehose_updateDestinationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(firehose_updateDestinationCmd).Standalone()

		firehose_updateDestinationCmd.Flags().String("amazon-open-search-serverless-destination-update", "", "Describes an update for a destination in the Serverless offering for Amazon OpenSearch Service.")
		firehose_updateDestinationCmd.Flags().String("amazonopensearchservice-destination-update", "", "Describes an update for a destination in Amazon OpenSearch Service.")
		firehose_updateDestinationCmd.Flags().String("current-delivery-stream-version-id", "", "Obtain this value from the `VersionId` result of [DeliveryStreamDescription]().")
		firehose_updateDestinationCmd.Flags().String("delivery-stream-name", "", "The name of the Firehose stream.")
		firehose_updateDestinationCmd.Flags().String("destination-id", "", "The ID of the destination.")
		firehose_updateDestinationCmd.Flags().String("elasticsearch-destination-update", "", "Describes an update for a destination in Amazon OpenSearch Service.")
		firehose_updateDestinationCmd.Flags().String("extended-s3-destination-update", "", "Describes an update for a destination in Amazon S3.")
		firehose_updateDestinationCmd.Flags().String("http-endpoint-destination-update", "", "Describes an update to the specified HTTP endpoint destination.")
		firehose_updateDestinationCmd.Flags().String("iceberg-destination-update", "", "Describes an update for a destination in Apache Iceberg Tables.")
		firehose_updateDestinationCmd.Flags().String("redshift-destination-update", "", "Describes an update for a destination in Amazon Redshift.")
		firehose_updateDestinationCmd.Flags().String("s3-destination-update", "", "\\[Deprecated] Describes an update for a destination in Amazon S3.")
		firehose_updateDestinationCmd.Flags().String("snowflake-destination-update", "", "Update to the Snowflake destination configuration settings.")
		firehose_updateDestinationCmd.Flags().String("splunk-destination-update", "", "Describes an update for a destination in Splunk.")
		firehose_updateDestinationCmd.MarkFlagRequired("current-delivery-stream-version-id")
		firehose_updateDestinationCmd.MarkFlagRequired("delivery-stream-name")
		firehose_updateDestinationCmd.MarkFlagRequired("destination-id")
	})
	firehoseCmd.AddCommand(firehose_updateDestinationCmd)
}
