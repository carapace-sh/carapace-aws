package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackage_deleteChannelCmd = &cobra.Command{
	Use:   "delete-channel",
	Short: "Deletes an existing Channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackage_deleteChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackage_deleteChannelCmd).Standalone()

		mediapackage_deleteChannelCmd.Flags().String("id", "", "The ID of the Channel to delete.")
		mediapackage_deleteChannelCmd.MarkFlagRequired("id")
	})
	mediapackageCmd.AddCommand(mediapackage_deleteChannelCmd)
}
