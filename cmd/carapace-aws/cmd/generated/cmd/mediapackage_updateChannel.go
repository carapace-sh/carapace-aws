package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackage_updateChannelCmd = &cobra.Command{
	Use:   "update-channel",
	Short: "Updates an existing Channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackage_updateChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackage_updateChannelCmd).Standalone()

		mediapackage_updateChannelCmd.Flags().String("description", "", "A short text description of the Channel.")
		mediapackage_updateChannelCmd.Flags().String("id", "", "The ID of the Channel to update.")
		mediapackage_updateChannelCmd.MarkFlagRequired("id")
	})
	mediapackageCmd.AddCommand(mediapackage_updateChannelCmd)
}
