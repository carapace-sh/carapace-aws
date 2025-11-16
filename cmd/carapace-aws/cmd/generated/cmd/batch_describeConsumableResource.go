package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_describeConsumableResourceCmd = &cobra.Command{
	Use:   "describe-consumable-resource",
	Short: "Returns a description of the specified consumable resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_describeConsumableResourceCmd).Standalone()

	batch_describeConsumableResourceCmd.Flags().String("consumable-resource", "", "The name or ARN of the consumable resource whose description will be returned.")
	batch_describeConsumableResourceCmd.MarkFlagRequired("consumable-resource")
	batchCmd.AddCommand(batch_describeConsumableResourceCmd)
}
