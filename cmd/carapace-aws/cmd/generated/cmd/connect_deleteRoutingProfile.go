package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_deleteRoutingProfileCmd = &cobra.Command{
	Use:   "delete-routing-profile",
	Short: "Deletes a routing profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_deleteRoutingProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_deleteRoutingProfileCmd).Standalone()

		connect_deleteRoutingProfileCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_deleteRoutingProfileCmd.Flags().String("routing-profile-id", "", "The identifier of the routing profile.")
		connect_deleteRoutingProfileCmd.MarkFlagRequired("instance-id")
		connect_deleteRoutingProfileCmd.MarkFlagRequired("routing-profile-id")
	})
	connectCmd.AddCommand(connect_deleteRoutingProfileCmd)
}
