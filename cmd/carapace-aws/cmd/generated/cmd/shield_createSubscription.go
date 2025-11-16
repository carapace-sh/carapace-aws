package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shield_createSubscriptionCmd = &cobra.Command{
	Use:   "create-subscription",
	Short: "Activates Shield Advanced for an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shield_createSubscriptionCmd).Standalone()

	shieldCmd.AddCommand(shield_createSubscriptionCmd)
}
