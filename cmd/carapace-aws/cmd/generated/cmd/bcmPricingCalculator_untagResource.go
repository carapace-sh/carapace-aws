package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmPricingCalculator_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmPricingCalculator_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bcmPricingCalculator_untagResourceCmd).Standalone()

		bcmPricingCalculator_untagResourceCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the resource to remove tags from.")
		bcmPricingCalculator_untagResourceCmd.Flags().String("tag-keys", "", "The keys of the tags to remove from the resource.")
		bcmPricingCalculator_untagResourceCmd.MarkFlagRequired("arn")
		bcmPricingCalculator_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	bcmPricingCalculatorCmd.AddCommand(bcmPricingCalculator_untagResourceCmd)
}
