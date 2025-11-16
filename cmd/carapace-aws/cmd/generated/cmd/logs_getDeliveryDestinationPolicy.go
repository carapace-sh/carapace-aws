package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_getDeliveryDestinationPolicyCmd = &cobra.Command{
	Use:   "get-delivery-destination-policy",
	Short: "Retrieves the delivery destination policy assigned to the delivery destination that you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_getDeliveryDestinationPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_getDeliveryDestinationPolicyCmd).Standalone()

		logs_getDeliveryDestinationPolicyCmd.Flags().String("delivery-destination-name", "", "The name of the delivery destination that you want to retrieve the policy of.")
		logs_getDeliveryDestinationPolicyCmd.MarkFlagRequired("delivery-destination-name")
	})
	logsCmd.AddCommand(logs_getDeliveryDestinationPolicyCmd)
}
