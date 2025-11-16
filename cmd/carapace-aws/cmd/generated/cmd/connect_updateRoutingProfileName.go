package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateRoutingProfileNameCmd = &cobra.Command{
	Use:   "update-routing-profile-name",
	Short: "Updates the name and description of a routing profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateRoutingProfileNameCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_updateRoutingProfileNameCmd).Standalone()

		connect_updateRoutingProfileNameCmd.Flags().String("description", "", "The description of the routing profile.")
		connect_updateRoutingProfileNameCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_updateRoutingProfileNameCmd.Flags().String("name", "", "The name of the routing profile.")
		connect_updateRoutingProfileNameCmd.Flags().String("routing-profile-id", "", "The identifier of the routing profile.")
		connect_updateRoutingProfileNameCmd.MarkFlagRequired("instance-id")
		connect_updateRoutingProfileNameCmd.MarkFlagRequired("routing-profile-id")
	})
	connectCmd.AddCommand(connect_updateRoutingProfileNameCmd)
}
