package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_createChannelPlacementGroupCmd = &cobra.Command{
	Use:   "create-channel-placement-group",
	Short: "Create a ChannelPlacementGroup in the specified Cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_createChannelPlacementGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_createChannelPlacementGroupCmd).Standalone()

		medialive_createChannelPlacementGroupCmd.Flags().String("cluster-id", "", "The ID of the cluster.")
		medialive_createChannelPlacementGroupCmd.Flags().String("name", "", "Specify a name that is unique in the Cluster.")
		medialive_createChannelPlacementGroupCmd.Flags().String("nodes", "", "An array of one ID for the Node that you want to associate with the ChannelPlacementGroup.")
		medialive_createChannelPlacementGroupCmd.Flags().String("request-id", "", "An ID that you assign to a create request.")
		medialive_createChannelPlacementGroupCmd.Flags().String("tags", "", "A collection of key-value pairs.")
		medialive_createChannelPlacementGroupCmd.MarkFlagRequired("cluster-id")
	})
	medialiveCmd.AddCommand(medialive_createChannelPlacementGroupCmd)
}
