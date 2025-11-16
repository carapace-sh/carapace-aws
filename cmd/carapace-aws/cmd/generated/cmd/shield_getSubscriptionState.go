package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shield_getSubscriptionStateCmd = &cobra.Command{
	Use:   "get-subscription-state",
	Short: "Returns the `SubscriptionState`, either `Active` or `Inactive`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shield_getSubscriptionStateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(shield_getSubscriptionStateCmd).Standalone()

	})
	shieldCmd.AddCommand(shield_getSubscriptionStateCmd)
}
