package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowball_createLongTermPricingCmd = &cobra.Command{
	Use:   "create-long-term-pricing",
	Short: "Creates a job with the long-term usage option for a device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowball_createLongTermPricingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(snowball_createLongTermPricingCmd).Standalone()

		snowball_createLongTermPricingCmd.Flags().String("is-long-term-pricing-auto-renew", "", "Specifies whether the current long-term pricing type for the device should be renewed.")
		snowball_createLongTermPricingCmd.Flags().String("long-term-pricing-type", "", "The type of long-term pricing option you want for the device, either 1-year or 3-year long-term pricing.")
		snowball_createLongTermPricingCmd.Flags().String("snowball-type", "", "The type of Snow Family devices to use for the long-term pricing job.")
		snowball_createLongTermPricingCmd.MarkFlagRequired("long-term-pricing-type")
		snowball_createLongTermPricingCmd.MarkFlagRequired("snowball-type")
	})
	snowballCmd.AddCommand(snowball_createLongTermPricingCmd)
}
