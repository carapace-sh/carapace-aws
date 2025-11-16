package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_getPricingPlanCmd = &cobra.Command{
	Use:   "get-pricing-plan",
	Short: "Gets the pricing plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_getPricingPlanCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iottwinmaker_getPricingPlanCmd).Standalone()

	})
	iottwinmakerCmd.AddCommand(iottwinmaker_getPricingPlanCmd)
}
