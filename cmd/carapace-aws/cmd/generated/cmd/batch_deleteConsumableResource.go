package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_deleteConsumableResourceCmd = &cobra.Command{
	Use:   "delete-consumable-resource",
	Short: "Deletes the specified consumable resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_deleteConsumableResourceCmd).Standalone()

	batch_deleteConsumableResourceCmd.Flags().String("consumable-resource", "", "The name or ARN of the consumable resource that will be deleted.")
	batch_deleteConsumableResourceCmd.MarkFlagRequired("consumable-resource")
	batchCmd.AddCommand(batch_deleteConsumableResourceCmd)
}
