package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_deleteImageCmd = &cobra.Command{
	Use:   "delete-image",
	Short: "Deletes an Image Builder image resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_deleteImageCmd).Standalone()

	imagebuilder_deleteImageCmd.Flags().String("image-build-version-arn", "", "The Amazon Resource Name (ARN) of the Image Builder image resource to delete.")
	imagebuilder_deleteImageCmd.MarkFlagRequired("image-build-version-arn")
	imagebuilderCmd.AddCommand(imagebuilder_deleteImageCmd)
}
