package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_updateDeliveryConfigurationCmd = &cobra.Command{
	Use:   "update-delivery-configuration",
	Short: "Use this operation to update the configuration of a [delivery](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_Delivery.html) to change either the S3 path pattern or the format of the delivered logs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_updateDeliveryConfigurationCmd).Standalone()

	logs_updateDeliveryConfigurationCmd.Flags().String("field-delimiter", "", "The field delimiter to use between record fields when the final output format of a delivery is in `Plain`, `W3C`, or `Raw` format.")
	logs_updateDeliveryConfigurationCmd.Flags().String("id", "", "The ID of the delivery to be updated by this request.")
	logs_updateDeliveryConfigurationCmd.Flags().String("record-fields", "", "The list of record fields to be delivered to the destination, in order.")
	logs_updateDeliveryConfigurationCmd.Flags().String("s3-delivery-configuration", "", "This structure contains parameters that are valid only when the delivery's delivery destination is an S3 bucket.")
	logs_updateDeliveryConfigurationCmd.MarkFlagRequired("id")
	logsCmd.AddCommand(logs_updateDeliveryConfigurationCmd)
}
