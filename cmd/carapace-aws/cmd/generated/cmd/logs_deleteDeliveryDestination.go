package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_deleteDeliveryDestinationCmd = &cobra.Command{
	Use:   "delete-delivery-destination",
	Short: "Deletes a *delivery destination*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_deleteDeliveryDestinationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_deleteDeliveryDestinationCmd).Standalone()

		logs_deleteDeliveryDestinationCmd.Flags().String("name", "", "The name of the delivery destination that you want to delete.")
		logs_deleteDeliveryDestinationCmd.MarkFlagRequired("name")
	})
	logsCmd.AddCommand(logs_deleteDeliveryDestinationCmd)
}
