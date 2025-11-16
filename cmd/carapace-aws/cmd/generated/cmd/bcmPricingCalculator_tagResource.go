package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmPricingCalculator_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds one or more tags to a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmPricingCalculator_tagResourceCmd).Standalone()

	bcmPricingCalculator_tagResourceCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the resource to add tags to.")
	bcmPricingCalculator_tagResourceCmd.Flags().String("tags", "", "The tags to add to the resource.")
	bcmPricingCalculator_tagResourceCmd.MarkFlagRequired("arn")
	bcmPricingCalculator_tagResourceCmd.MarkFlagRequired("tags")
	bcmPricingCalculatorCmd.AddCommand(bcmPricingCalculator_tagResourceCmd)
}
