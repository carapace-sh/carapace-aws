package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_createDeliveryCmd = &cobra.Command{
	Use:   "create-delivery",
	Short: "Creates a *delivery*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_createDeliveryCmd).Standalone()

	logs_createDeliveryCmd.Flags().String("delivery-destination-arn", "", "The ARN of the delivery destination to use for this delivery.")
	logs_createDeliveryCmd.Flags().String("delivery-source-name", "", "The name of the delivery source to use for this delivery.")
	logs_createDeliveryCmd.Flags().String("field-delimiter", "", "The field delimiter to use between record fields when the final output format of a delivery is in `Plain`, `W3C`, or `Raw` format.")
	logs_createDeliveryCmd.Flags().String("record-fields", "", "The list of record fields to be delivered to the destination, in order.")
	logs_createDeliveryCmd.Flags().String("s3-delivery-configuration", "", "This structure contains parameters that are valid only when the delivery's delivery destination is an S3 bucket.")
	logs_createDeliveryCmd.Flags().String("tags", "", "An optional list of key-value pairs to associate with the resource.")
	logs_createDeliveryCmd.MarkFlagRequired("delivery-destination-arn")
	logs_createDeliveryCmd.MarkFlagRequired("delivery-source-name")
	logsCmd.AddCommand(logs_createDeliveryCmd)
}
