package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_describeVodSourceCmd = &cobra.Command{
	Use:   "describe-vod-source",
	Short: "Provides details about a specific video on demand (VOD) source in a specific source location.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_describeVodSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediatailor_describeVodSourceCmd).Standalone()

		mediatailor_describeVodSourceCmd.Flags().String("source-location-name", "", "The name of the source location associated with this VOD Source.")
		mediatailor_describeVodSourceCmd.Flags().String("vod-source-name", "", "The name of the VOD Source.")
		mediatailor_describeVodSourceCmd.MarkFlagRequired("source-location-name")
		mediatailor_describeVodSourceCmd.MarkFlagRequired("vod-source-name")
	})
	mediatailorCmd.AddCommand(mediatailor_describeVodSourceCmd)
}
