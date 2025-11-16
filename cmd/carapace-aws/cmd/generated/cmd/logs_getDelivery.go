package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_getDeliveryCmd = &cobra.Command{
	Use:   "get-delivery",
	Short: "Returns complete information about one logical *delivery*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_getDeliveryCmd).Standalone()

	logs_getDeliveryCmd.Flags().String("id", "", "The ID of the delivery that you want to retrieve.")
	logs_getDeliveryCmd.MarkFlagRequired("id")
	logsCmd.AddCommand(logs_getDeliveryCmd)
}
