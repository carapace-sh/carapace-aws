package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shield_updateSubscriptionCmd = &cobra.Command{
	Use:   "update-subscription",
	Short: "Updates the details of an existing subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shield_updateSubscriptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(shield_updateSubscriptionCmd).Standalone()

		shield_updateSubscriptionCmd.Flags().String("auto-renew", "", "When you initally create a subscription, `AutoRenew` is set to `ENABLED`.")
	})
	shieldCmd.AddCommand(shield_updateSubscriptionCmd)
}
