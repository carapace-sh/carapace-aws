package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_deleteImageBuilderCmd = &cobra.Command{
	Use:   "delete-image-builder",
	Short: "Deletes the specified image builder and releases the capacity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_deleteImageBuilderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_deleteImageBuilderCmd).Standalone()

		appstream_deleteImageBuilderCmd.Flags().String("name", "", "The name of the image builder.")
		appstream_deleteImageBuilderCmd.MarkFlagRequired("name")
	})
	appstreamCmd.AddCommand(appstream_deleteImageBuilderCmd)
}
