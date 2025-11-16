package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_getImageCmd = &cobra.Command{
	Use:   "get-image",
	Short: "Gets an image.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_getImageCmd).Standalone()

	imagebuilder_getImageCmd.Flags().String("image-build-version-arn", "", "The Amazon Resource Name (ARN) of the image that you want to get.")
	imagebuilder_getImageCmd.MarkFlagRequired("image-build-version-arn")
	imagebuilderCmd.AddCommand(imagebuilder_getImageCmd)
}
