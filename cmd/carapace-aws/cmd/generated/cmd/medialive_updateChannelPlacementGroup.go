package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_updateChannelPlacementGroupCmd = &cobra.Command{
	Use:   "update-channel-placement-group",
	Short: "Change the settings for a ChannelPlacementGroup.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_updateChannelPlacementGroupCmd).Standalone()

	medialive_updateChannelPlacementGroupCmd.Flags().String("channel-placement-group-id", "", "The ID of the channel placement group.")
	medialive_updateChannelPlacementGroupCmd.Flags().String("cluster-id", "", "The ID of the cluster.")
	medialive_updateChannelPlacementGroupCmd.Flags().String("name", "", "Include this parameter only if you want to change the current name of the ChannelPlacementGroup.")
	medialive_updateChannelPlacementGroupCmd.Flags().String("nodes", "", "Include this parameter only if you want to change the list of Nodes that are associated with the ChannelPlacementGroup.")
	medialive_updateChannelPlacementGroupCmd.MarkFlagRequired("channel-placement-group-id")
	medialive_updateChannelPlacementGroupCmd.MarkFlagRequired("cluster-id")
	medialiveCmd.AddCommand(medialive_updateChannelPlacementGroupCmd)
}
