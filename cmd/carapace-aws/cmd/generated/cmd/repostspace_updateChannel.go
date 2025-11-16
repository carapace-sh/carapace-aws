package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repostspace_updateChannelCmd = &cobra.Command{
	Use:   "update-channel",
	Short: "Modifies an existing channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repostspace_updateChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(repostspace_updateChannelCmd).Standalone()

		repostspace_updateChannelCmd.Flags().String("channel-description", "", "A description for the channel.")
		repostspace_updateChannelCmd.Flags().String("channel-id", "", "The unique ID of the private re:Post channel.")
		repostspace_updateChannelCmd.Flags().String("channel-name", "", "The name for the channel.")
		repostspace_updateChannelCmd.Flags().String("space-id", "", "The unique ID of the private re:Post.")
		repostspace_updateChannelCmd.MarkFlagRequired("channel-id")
		repostspace_updateChannelCmd.MarkFlagRequired("channel-name")
		repostspace_updateChannelCmd.MarkFlagRequired("space-id")
	})
	repostspaceCmd.AddCommand(repostspace_updateChannelCmd)
}
