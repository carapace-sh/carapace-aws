package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_stopImageBuilderCmd = &cobra.Command{
	Use:   "stop-image-builder",
	Short: "Stops the specified image builder.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_stopImageBuilderCmd).Standalone()

	appstream_stopImageBuilderCmd.Flags().String("name", "", "The name of the image builder.")
	appstream_stopImageBuilderCmd.MarkFlagRequired("name")
	appstreamCmd.AddCommand(appstream_stopImageBuilderCmd)
}
