package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackage_createChannelCmd = &cobra.Command{
	Use:   "create-channel",
	Short: "Creates a new Channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackage_createChannelCmd).Standalone()

	mediapackage_createChannelCmd.Flags().String("description", "", "A short text description of the Channel.")
	mediapackage_createChannelCmd.Flags().String("id", "", "The ID of the Channel.")
	mediapackage_createChannelCmd.Flags().String("tags", "", "")
	mediapackage_createChannelCmd.MarkFlagRequired("id")
	mediapackageCmd.AddCommand(mediapackage_createChannelCmd)
}
