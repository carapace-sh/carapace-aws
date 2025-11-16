package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_putDeliverySourceCmd = &cobra.Command{
	Use:   "put-delivery-source",
	Short: "Creates or updates a logical *delivery source*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_putDeliverySourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_putDeliverySourceCmd).Standalone()

		logs_putDeliverySourceCmd.Flags().String("log-type", "", "Defines the type of log that the source is sending.")
		logs_putDeliverySourceCmd.Flags().String("name", "", "A name for this delivery source.")
		logs_putDeliverySourceCmd.Flags().String("resource-arn", "", "The ARN of the Amazon Web Services resource that is generating and sending logs.")
		logs_putDeliverySourceCmd.Flags().String("tags", "", "An optional list of key-value pairs to associate with the resource.")
		logs_putDeliverySourceCmd.MarkFlagRequired("log-type")
		logs_putDeliverySourceCmd.MarkFlagRequired("name")
		logs_putDeliverySourceCmd.MarkFlagRequired("resource-arn")
	})
	logsCmd.AddCommand(logs_putDeliverySourceCmd)
}
