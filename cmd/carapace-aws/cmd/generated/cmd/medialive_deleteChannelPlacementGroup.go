package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_deleteChannelPlacementGroupCmd = &cobra.Command{
	Use:   "delete-channel-placement-group",
	Short: "Delete the specified ChannelPlacementGroup that exists in the specified Cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_deleteChannelPlacementGroupCmd).Standalone()

	medialive_deleteChannelPlacementGroupCmd.Flags().String("channel-placement-group-id", "", "The ID of the channel placement group.")
	medialive_deleteChannelPlacementGroupCmd.Flags().String("cluster-id", "", "The ID of the cluster.")
	medialive_deleteChannelPlacementGroupCmd.MarkFlagRequired("channel-placement-group-id")
	medialive_deleteChannelPlacementGroupCmd.MarkFlagRequired("cluster-id")
	medialiveCmd.AddCommand(medialive_deleteChannelPlacementGroupCmd)
}
