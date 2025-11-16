package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shield_deleteSubscriptionCmd = &cobra.Command{
	Use:   "delete-subscription",
	Short: "Removes Shield Advanced from an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shield_deleteSubscriptionCmd).Standalone()

	shieldCmd.AddCommand(shield_deleteSubscriptionCmd)
}
