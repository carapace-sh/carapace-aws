package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_getDeliveryDestinationCmd = &cobra.Command{
	Use:   "get-delivery-destination",
	Short: "Retrieves complete information about one delivery destination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_getDeliveryDestinationCmd).Standalone()

	logs_getDeliveryDestinationCmd.Flags().String("name", "", "The name of the delivery destination that you want to retrieve.")
	logs_getDeliveryDestinationCmd.MarkFlagRequired("name")
	logsCmd.AddCommand(logs_getDeliveryDestinationCmd)
}
