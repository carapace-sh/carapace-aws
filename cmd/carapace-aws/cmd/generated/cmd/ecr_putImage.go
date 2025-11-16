package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_putImageCmd = &cobra.Command{
	Use:   "put-image",
	Short: "Creates or updates the image manifest and tags associated with an image.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_putImageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecr_putImageCmd).Standalone()

		ecr_putImageCmd.Flags().String("image-digest", "", "The image digest of the image manifest corresponding to the image.")
		ecr_putImageCmd.Flags().String("image-manifest", "", "The image manifest corresponding to the image to be uploaded.")
		ecr_putImageCmd.Flags().String("image-manifest-media-type", "", "The media type of the image manifest.")
		ecr_putImageCmd.Flags().String("image-tag", "", "The tag to associate with the image.")
		ecr_putImageCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID associated with the registry that contains the repository in which to put the image.")
		ecr_putImageCmd.Flags().String("repository-name", "", "The name of the repository in which to put the image.")
		ecr_putImageCmd.MarkFlagRequired("image-manifest")
		ecr_putImageCmd.MarkFlagRequired("repository-name")
	})
	ecrCmd.AddCommand(ecr_putImageCmd)
}
