package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_copyImageCmd = &cobra.Command{
	Use:   "copy-image",
	Short: "Copies the image within the same region or to a new region within the same AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_copyImageCmd).Standalone()

	appstream_copyImageCmd.Flags().String("destination-image-description", "", "The description that the image will have when it is copied to the destination.")
	appstream_copyImageCmd.Flags().String("destination-image-name", "", "The name that the image will have when it is copied to the destination.")
	appstream_copyImageCmd.Flags().String("destination-region", "", "The destination region to which the image will be copied.")
	appstream_copyImageCmd.Flags().String("source-image-name", "", "The name of the image to copy.")
	appstream_copyImageCmd.MarkFlagRequired("destination-image-name")
	appstream_copyImageCmd.MarkFlagRequired("destination-region")
	appstream_copyImageCmd.MarkFlagRequired("source-image-name")
	appstreamCmd.AddCommand(appstream_copyImageCmd)
}
