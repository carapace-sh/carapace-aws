package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_startImageBuilderCmd = &cobra.Command{
	Use:   "start-image-builder",
	Short: "Starts the specified image builder.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_startImageBuilderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_startImageBuilderCmd).Standalone()

		appstream_startImageBuilderCmd.Flags().String("appstream-agent-version", "", "The version of the AppStream 2.0 agent to use for this image builder.")
		appstream_startImageBuilderCmd.Flags().String("name", "", "The name of the image builder.")
		appstream_startImageBuilderCmd.MarkFlagRequired("name")
	})
	appstreamCmd.AddCommand(appstream_startImageBuilderCmd)
}
