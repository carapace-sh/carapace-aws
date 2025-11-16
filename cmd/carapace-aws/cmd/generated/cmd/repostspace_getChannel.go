package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repostspace_getChannelCmd = &cobra.Command{
	Use:   "get-channel",
	Short: "Displays information about a channel in a private re:Post.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repostspace_getChannelCmd).Standalone()

	repostspace_getChannelCmd.Flags().String("channel-id", "", "The unique ID of the private re:Post channel.")
	repostspace_getChannelCmd.Flags().String("space-id", "", "The unique ID of the private re:Post.")
	repostspace_getChannelCmd.MarkFlagRequired("channel-id")
	repostspace_getChannelCmd.MarkFlagRequired("space-id")
	repostspaceCmd.AddCommand(repostspace_getChannelCmd)
}
