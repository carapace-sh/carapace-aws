package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_deleteImageCmd = &cobra.Command{
	Use:   "delete-image",
	Short: "Deletes the specified image.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_deleteImageCmd).Standalone()

	appstream_deleteImageCmd.Flags().String("name", "", "The name of the image.")
	appstream_deleteImageCmd.MarkFlagRequired("name")
	appstreamCmd.AddCommand(appstream_deleteImageCmd)
}
