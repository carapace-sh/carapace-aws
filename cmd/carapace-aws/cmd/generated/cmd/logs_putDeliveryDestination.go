package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_putDeliveryDestinationCmd = &cobra.Command{
	Use:   "put-delivery-destination",
	Short: "Creates or updates a logical *delivery destination*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_putDeliveryDestinationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_putDeliveryDestinationCmd).Standalone()

		logs_putDeliveryDestinationCmd.Flags().String("delivery-destination-configuration", "", "A structure that contains the ARN of the Amazon Web Services resource that will receive the logs.")
		logs_putDeliveryDestinationCmd.Flags().String("delivery-destination-type", "", "The type of delivery destination.")
		logs_putDeliveryDestinationCmd.Flags().String("name", "", "A name for this delivery destination.")
		logs_putDeliveryDestinationCmd.Flags().String("output-format", "", "The format for the logs that this delivery destination will receive.")
		logs_putDeliveryDestinationCmd.Flags().String("tags", "", "An optional list of key-value pairs to associate with the resource.")
		logs_putDeliveryDestinationCmd.MarkFlagRequired("name")
	})
	logsCmd.AddCommand(logs_putDeliveryDestinationCmd)
}
