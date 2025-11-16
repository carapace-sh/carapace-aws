package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_describeThumbnailsCmd = &cobra.Command{
	Use:   "describe-thumbnails",
	Short: "Describe the latest thumbnails data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_describeThumbnailsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_describeThumbnailsCmd).Standalone()

		medialive_describeThumbnailsCmd.Flags().String("channel-id", "", "Unique ID of the channel")
		medialive_describeThumbnailsCmd.Flags().String("pipeline-id", "", "Pipeline ID (\"0\" or \"1\")")
		medialive_describeThumbnailsCmd.Flags().String("thumbnail-type", "", "thumbnail type")
		medialive_describeThumbnailsCmd.MarkFlagRequired("channel-id")
		medialive_describeThumbnailsCmd.MarkFlagRequired("pipeline-id")
		medialive_describeThumbnailsCmd.MarkFlagRequired("thumbnail-type")
	})
	medialiveCmd.AddCommand(medialive_describeThumbnailsCmd)
}
