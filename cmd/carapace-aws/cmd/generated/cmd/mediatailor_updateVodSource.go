package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_updateVodSourceCmd = &cobra.Command{
	Use:   "update-vod-source",
	Short: "Updates a VOD source's configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_updateVodSourceCmd).Standalone()

	mediatailor_updateVodSourceCmd.Flags().String("http-package-configurations", "", "A list of HTTP package configurations for the VOD source on this account.")
	mediatailor_updateVodSourceCmd.Flags().String("source-location-name", "", "The name of the source location associated with this VOD Source.")
	mediatailor_updateVodSourceCmd.Flags().String("vod-source-name", "", "The name of the VOD source.")
	mediatailor_updateVodSourceCmd.MarkFlagRequired("http-package-configurations")
	mediatailor_updateVodSourceCmd.MarkFlagRequired("source-location-name")
	mediatailor_updateVodSourceCmd.MarkFlagRequired("vod-source-name")
	mediatailorCmd.AddCommand(mediatailor_updateVodSourceCmd)
}
