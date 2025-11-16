package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_describeDeliveryDestinationsCmd = &cobra.Command{
	Use:   "describe-delivery-destinations",
	Short: "Retrieves a list of the delivery destinations that have been created in the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_describeDeliveryDestinationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_describeDeliveryDestinationsCmd).Standalone()

		logs_describeDeliveryDestinationsCmd.Flags().String("limit", "", "Optionally specify the maximum number of delivery destinations to return in the response.")
		logs_describeDeliveryDestinationsCmd.Flags().String("next-token", "", "")
	})
	logsCmd.AddCommand(logs_describeDeliveryDestinationsCmd)
}
