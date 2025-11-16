package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_createVodSourceCmd = &cobra.Command{
	Use:   "create-vod-source",
	Short: "The VOD source configuration parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_createVodSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediatailor_createVodSourceCmd).Standalone()

		mediatailor_createVodSourceCmd.Flags().String("http-package-configurations", "", "A list of HTTP package configuration parameters for this VOD source.")
		mediatailor_createVodSourceCmd.Flags().String("source-location-name", "", "The name of the source location for this VOD source.")
		mediatailor_createVodSourceCmd.Flags().String("tags", "", "The tags to assign to the VOD source.")
		mediatailor_createVodSourceCmd.Flags().String("vod-source-name", "", "The name associated with the VOD source.&gt;")
		mediatailor_createVodSourceCmd.MarkFlagRequired("http-package-configurations")
		mediatailor_createVodSourceCmd.MarkFlagRequired("source-location-name")
		mediatailor_createVodSourceCmd.MarkFlagRequired("vod-source-name")
	})
	mediatailorCmd.AddCommand(mediatailor_createVodSourceCmd)
}
