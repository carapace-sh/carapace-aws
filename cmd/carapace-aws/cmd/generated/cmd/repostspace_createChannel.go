package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repostspace_createChannelCmd = &cobra.Command{
	Use:   "create-channel",
	Short: "Creates a channel in an AWS re:Post Private private re:Post.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repostspace_createChannelCmd).Standalone()

	repostspace_createChannelCmd.Flags().String("channel-description", "", "A description for the channel.")
	repostspace_createChannelCmd.Flags().String("channel-name", "", "The name for the channel.")
	repostspace_createChannelCmd.Flags().String("space-id", "", "The unique ID of the private re:Post.")
	repostspace_createChannelCmd.MarkFlagRequired("channel-name")
	repostspace_createChannelCmd.MarkFlagRequired("space-id")
	repostspaceCmd.AddCommand(repostspace_createChannelCmd)
}
