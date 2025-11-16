package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_listSubscriptionsCmd = &cobra.Command{
	Use:   "list-subscriptions",
	Short: "Returns a list of the requester's subscriptions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_listSubscriptionsCmd).Standalone()

	sns_listSubscriptionsCmd.Flags().String("next-token", "", "Token returned by the previous `ListSubscriptions` request.")
	snsCmd.AddCommand(sns_listSubscriptionsCmd)
}
