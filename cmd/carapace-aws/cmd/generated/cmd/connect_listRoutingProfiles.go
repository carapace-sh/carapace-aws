package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listRoutingProfilesCmd = &cobra.Command{
	Use:   "list-routing-profiles",
	Short: "Provides summary information about the routing profiles for the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listRoutingProfilesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_listRoutingProfilesCmd).Standalone()

		connect_listRoutingProfilesCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_listRoutingProfilesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_listRoutingProfilesCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_listRoutingProfilesCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_listRoutingProfilesCmd)
}
