package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateRoutingProfileConcurrencyCmd = &cobra.Command{
	Use:   "update-routing-profile-concurrency",
	Short: "Updates the channels that agents can handle in the Contact Control Panel (CCP) for a routing profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateRoutingProfileConcurrencyCmd).Standalone()

	connect_updateRoutingProfileConcurrencyCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_updateRoutingProfileConcurrencyCmd.Flags().String("media-concurrencies", "", "The channels that agents can handle in the Contact Control Panel (CCP).")
	connect_updateRoutingProfileConcurrencyCmd.Flags().String("routing-profile-id", "", "The identifier of the routing profile.")
	connect_updateRoutingProfileConcurrencyCmd.MarkFlagRequired("instance-id")
	connect_updateRoutingProfileConcurrencyCmd.MarkFlagRequired("media-concurrencies")
	connect_updateRoutingProfileConcurrencyCmd.MarkFlagRequired("routing-profile-id")
	connectCmd.AddCommand(connect_updateRoutingProfileConcurrencyCmd)
}
