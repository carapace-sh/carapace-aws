package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmPricingCalculator_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all tags associated with a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmPricingCalculator_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bcmPricingCalculator_listTagsForResourceCmd).Standalone()

		bcmPricingCalculator_listTagsForResourceCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the resource to list tags for.")
		bcmPricingCalculator_listTagsForResourceCmd.MarkFlagRequired("arn")
	})
	bcmPricingCalculatorCmd.AddCommand(bcmPricingCalculator_listTagsForResourceCmd)
}
