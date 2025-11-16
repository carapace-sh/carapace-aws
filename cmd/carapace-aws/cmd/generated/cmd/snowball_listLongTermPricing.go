package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowball_listLongTermPricingCmd = &cobra.Command{
	Use:   "list-long-term-pricing",
	Short: "Lists all long-term pricing types.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowball_listLongTermPricingCmd).Standalone()

	snowball_listLongTermPricingCmd.Flags().String("max-results", "", "The maximum number of `ListLongTermPricing` objects to return.")
	snowball_listLongTermPricingCmd.Flags().String("next-token", "", "Because HTTP requests are stateless, this is the starting point for your next list of `ListLongTermPricing` to return.")
	snowballCmd.AddCommand(snowball_listLongTermPricingCmd)
}
