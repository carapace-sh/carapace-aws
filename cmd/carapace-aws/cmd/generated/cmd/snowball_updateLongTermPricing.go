package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowball_updateLongTermPricingCmd = &cobra.Command{
	Use:   "update-long-term-pricing",
	Short: "Updates the long-term pricing type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowball_updateLongTermPricingCmd).Standalone()

	snowball_updateLongTermPricingCmd.Flags().String("is-long-term-pricing-auto-renew", "", "If set to `true`, specifies that the current long-term pricing type for the device should be automatically renewed before the long-term pricing contract expires.")
	snowball_updateLongTermPricingCmd.Flags().String("long-term-pricing-id", "", "The ID of the long-term pricing type for the device.")
	snowball_updateLongTermPricingCmd.Flags().String("replacement-job", "", "Specifies that a device that is ordered with long-term pricing should be replaced with a new device.")
	snowball_updateLongTermPricingCmd.MarkFlagRequired("long-term-pricing-id")
	snowballCmd.AddCommand(snowball_updateLongTermPricingCmd)
}
