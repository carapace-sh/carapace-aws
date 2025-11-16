package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_restartChannelPipelinesCmd = &cobra.Command{
	Use:   "restart-channel-pipelines",
	Short: "Restart pipelines in one channel that is currently running.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_restartChannelPipelinesCmd).Standalone()

	medialive_restartChannelPipelinesCmd.Flags().String("channel-id", "", "ID of channel")
	medialive_restartChannelPipelinesCmd.Flags().String("pipeline-ids", "", "An array of pipelines to restart in this channel.")
	medialive_restartChannelPipelinesCmd.MarkFlagRequired("channel-id")
	medialiveCmd.AddCommand(medialive_restartChannelPipelinesCmd)
}
