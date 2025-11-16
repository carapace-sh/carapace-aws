package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shield_describeSubscriptionCmd = &cobra.Command{
	Use:   "describe-subscription",
	Short: "Provides details about the Shield Advanced subscription for an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shield_describeSubscriptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(shield_describeSubscriptionCmd).Standalone()

	})
	shieldCmd.AddCommand(shield_describeSubscriptionCmd)
}
