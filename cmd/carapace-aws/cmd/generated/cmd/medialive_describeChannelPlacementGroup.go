package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_describeChannelPlacementGroupCmd = &cobra.Command{
	Use:   "describe-channel-placement-group",
	Short: "Get details about a ChannelPlacementGroup.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_describeChannelPlacementGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_describeChannelPlacementGroupCmd).Standalone()

		medialive_describeChannelPlacementGroupCmd.Flags().String("channel-placement-group-id", "", "The ID of the channel placement group.")
		medialive_describeChannelPlacementGroupCmd.Flags().String("cluster-id", "", "The ID of the cluster.")
		medialive_describeChannelPlacementGroupCmd.MarkFlagRequired("channel-placement-group-id")
		medialive_describeChannelPlacementGroupCmd.MarkFlagRequired("cluster-id")
	})
	medialiveCmd.AddCommand(medialive_describeChannelPlacementGroupCmd)
}
