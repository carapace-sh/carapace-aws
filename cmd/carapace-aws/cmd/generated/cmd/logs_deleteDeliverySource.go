package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_deleteDeliverySourceCmd = &cobra.Command{
	Use:   "delete-delivery-source",
	Short: "Deletes a *delivery source*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_deleteDeliverySourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_deleteDeliverySourceCmd).Standalone()

		logs_deleteDeliverySourceCmd.Flags().String("name", "", "The name of the delivery source that you want to delete.")
		logs_deleteDeliverySourceCmd.MarkFlagRequired("name")
	})
	logsCmd.AddCommand(logs_deleteDeliverySourceCmd)
}
