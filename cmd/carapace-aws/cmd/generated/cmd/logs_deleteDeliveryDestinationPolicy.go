package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_deleteDeliveryDestinationPolicyCmd = &cobra.Command{
	Use:   "delete-delivery-destination-policy",
	Short: "Deletes a delivery destination policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_deleteDeliveryDestinationPolicyCmd).Standalone()

	logs_deleteDeliveryDestinationPolicyCmd.Flags().String("delivery-destination-name", "", "The name of the delivery destination that you want to delete the policy for.")
	logs_deleteDeliveryDestinationPolicyCmd.MarkFlagRequired("delivery-destination-name")
	logsCmd.AddCommand(logs_deleteDeliveryDestinationPolicyCmd)
}
