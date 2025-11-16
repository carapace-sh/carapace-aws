package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_createImageBuilderStreamingUrlCmd = &cobra.Command{
	Use:   "create-image-builder-streaming-url",
	Short: "Creates a URL to start an image builder streaming session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_createImageBuilderStreamingUrlCmd).Standalone()

	appstream_createImageBuilderStreamingUrlCmd.Flags().String("name", "", "The name of the image builder.")
	appstream_createImageBuilderStreamingUrlCmd.Flags().String("validity", "", "The time that the streaming URL will be valid, in seconds.")
	appstream_createImageBuilderStreamingUrlCmd.MarkFlagRequired("name")
	appstreamCmd.AddCommand(appstream_createImageBuilderStreamingUrlCmd)
}
