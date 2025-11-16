package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_updateConsumableResourceCmd = &cobra.Command{
	Use:   "update-consumable-resource",
	Short: "Updates a consumable resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_updateConsumableResourceCmd).Standalone()

	batch_updateConsumableResourceCmd.Flags().String("client-token", "", "If this parameter is specified and two update requests with identical payloads and `clientToken`s are received, these requests are considered the same request.")
	batch_updateConsumableResourceCmd.Flags().String("consumable-resource", "", "The name or ARN of the consumable resource to be updated.")
	batch_updateConsumableResourceCmd.Flags().String("operation", "", "Indicates how the quantity of the consumable resource will be updated.")
	batch_updateConsumableResourceCmd.Flags().String("quantity", "", "The change in the total quantity of the consumable resource.")
	batch_updateConsumableResourceCmd.MarkFlagRequired("consumable-resource")
	batchCmd.AddCommand(batch_updateConsumableResourceCmd)
}
