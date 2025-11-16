package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_describeDeliverySourcesCmd = &cobra.Command{
	Use:   "describe-delivery-sources",
	Short: "Retrieves a list of the delivery sources that have been created in the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_describeDeliverySourcesCmd).Standalone()

	logs_describeDeliverySourcesCmd.Flags().String("limit", "", "Optionally specify the maximum number of delivery sources to return in the response.")
	logs_describeDeliverySourcesCmd.Flags().String("next-token", "", "")
	logsCmd.AddCommand(logs_describeDeliverySourcesCmd)
}
