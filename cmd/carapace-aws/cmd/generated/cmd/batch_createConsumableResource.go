package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_createConsumableResourceCmd = &cobra.Command{
	Use:   "create-consumable-resource",
	Short: "Creates an Batch consumable resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_createConsumableResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(batch_createConsumableResourceCmd).Standalone()

		batch_createConsumableResourceCmd.Flags().String("consumable-resource-name", "", "The name of the consumable resource.")
		batch_createConsumableResourceCmd.Flags().String("resource-type", "", "Indicates whether the resource is available to be re-used after a job completes.")
		batch_createConsumableResourceCmd.Flags().String("tags", "", "The tags that you apply to the consumable resource to help you categorize and organize your resources.")
		batch_createConsumableResourceCmd.Flags().String("total-quantity", "", "The total amount of the consumable resource that is available.")
		batch_createConsumableResourceCmd.MarkFlagRequired("consumable-resource-name")
	})
	batchCmd.AddCommand(batch_createConsumableResourceCmd)
}
