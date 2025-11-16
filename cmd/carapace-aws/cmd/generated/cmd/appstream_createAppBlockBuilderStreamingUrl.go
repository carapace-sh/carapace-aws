package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_createAppBlockBuilderStreamingUrlCmd = &cobra.Command{
	Use:   "create-app-block-builder-streaming-url",
	Short: "Creates a URL to start a create app block builder streaming session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_createAppBlockBuilderStreamingUrlCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_createAppBlockBuilderStreamingUrlCmd).Standalone()

		appstream_createAppBlockBuilderStreamingUrlCmd.Flags().String("app-block-builder-name", "", "The name of the app block builder.")
		appstream_createAppBlockBuilderStreamingUrlCmd.Flags().String("validity", "", "The time that the streaming URL will be valid, in seconds.")
		appstream_createAppBlockBuilderStreamingUrlCmd.MarkFlagRequired("app-block-builder-name")
	})
	appstreamCmd.AddCommand(appstream_createAppBlockBuilderStreamingUrlCmd)
}
