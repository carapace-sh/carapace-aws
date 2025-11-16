package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_getDeliverySourceCmd = &cobra.Command{
	Use:   "get-delivery-source",
	Short: "Retrieves complete information about one delivery source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_getDeliverySourceCmd).Standalone()

	logs_getDeliverySourceCmd.Flags().String("name", "", "The name of the delivery source that you want to retrieve.")
	logs_getDeliverySourceCmd.MarkFlagRequired("name")
	logsCmd.AddCommand(logs_getDeliverySourceCmd)
}
