package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_listChannelPlacementGroupsCmd = &cobra.Command{
	Use:   "list-channel-placement-groups",
	Short: "Retrieve the list of ChannelPlacementGroups in the specified Cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_listChannelPlacementGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_listChannelPlacementGroupsCmd).Standalone()

		medialive_listChannelPlacementGroupsCmd.Flags().String("cluster-id", "", "The ID of the cluster")
		medialive_listChannelPlacementGroupsCmd.Flags().String("max-results", "", "The maximum number of items to return.")
		medialive_listChannelPlacementGroupsCmd.Flags().String("next-token", "", "The token to retrieve the next page of results.")
		medialive_listChannelPlacementGroupsCmd.MarkFlagRequired("cluster-id")
	})
	medialiveCmd.AddCommand(medialive_listChannelPlacementGroupsCmd)
}
