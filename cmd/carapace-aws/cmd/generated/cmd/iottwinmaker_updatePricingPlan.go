package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_updatePricingPlanCmd = &cobra.Command{
	Use:   "update-pricing-plan",
	Short: "Update the pricing plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_updatePricingPlanCmd).Standalone()

	iottwinmaker_updatePricingPlanCmd.Flags().String("bundle-names", "", "The bundle names.")
	iottwinmaker_updatePricingPlanCmd.Flags().String("pricing-mode", "", "The pricing mode.")
	iottwinmaker_updatePricingPlanCmd.MarkFlagRequired("pricing-mode")
	iottwinmakerCmd.AddCommand(iottwinmaker_updatePricingPlanCmd)
}
