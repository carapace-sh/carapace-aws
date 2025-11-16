package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_deleteDeliveryCmd = &cobra.Command{
	Use:   "delete-delivery",
	Short: "Deletes a *delivery*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_deleteDeliveryCmd).Standalone()

	logs_deleteDeliveryCmd.Flags().String("id", "", "The unique ID of the delivery to delete.")
	logs_deleteDeliveryCmd.MarkFlagRequired("id")
	logsCmd.AddCommand(logs_deleteDeliveryCmd)
}
