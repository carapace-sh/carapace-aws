package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecrPublic_putImageCmd = &cobra.Command{
	Use:   "put-image",
	Short: "Creates or updates the image manifest and tags that are associated with an image.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecrPublic_putImageCmd).Standalone()

	ecrPublic_putImageCmd.Flags().String("image-digest", "", "The image digest of the image manifest that corresponds to the image.")
	ecrPublic_putImageCmd.Flags().String("image-manifest", "", "The image manifest that corresponds to the image to be uploaded.")
	ecrPublic_putImageCmd.Flags().String("image-manifest-media-type", "", "The media type of the image manifest.")
	ecrPublic_putImageCmd.Flags().String("image-tag", "", "The tag to associate with the image.")
	ecrPublic_putImageCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID, or registry alias, that's associated with the public registry that contains the repository where the image is put.")
	ecrPublic_putImageCmd.Flags().String("repository-name", "", "The name of the repository where the image is put.")
	ecrPublic_putImageCmd.MarkFlagRequired("image-manifest")
	ecrPublic_putImageCmd.MarkFlagRequired("repository-name")
	ecrPublicCmd.AddCommand(ecrPublic_putImageCmd)
}
