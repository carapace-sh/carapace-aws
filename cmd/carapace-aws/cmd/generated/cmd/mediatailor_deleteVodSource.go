package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_deleteVodSourceCmd = &cobra.Command{
	Use:   "delete-vod-source",
	Short: "The video on demand (VOD) source to delete.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_deleteVodSourceCmd).Standalone()

	mediatailor_deleteVodSourceCmd.Flags().String("source-location-name", "", "The name of the source location associated with this VOD Source.")
	mediatailor_deleteVodSourceCmd.Flags().String("vod-source-name", "", "The name of the VOD source.")
	mediatailor_deleteVodSourceCmd.MarkFlagRequired("source-location-name")
	mediatailor_deleteVodSourceCmd.MarkFlagRequired("vod-source-name")
	mediatailorCmd.AddCommand(mediatailor_deleteVodSourceCmd)
}
