package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecrPublic_batchDeleteImageCmd = &cobra.Command{
	Use:   "batch-delete-image",
	Short: "Deletes a list of specified images that are within a repository in a public registry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecrPublic_batchDeleteImageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecrPublic_batchDeleteImageCmd).Standalone()

		ecrPublic_batchDeleteImageCmd.Flags().String("image-ids", "", "A list of image ID references that correspond to images to delete.")
		ecrPublic_batchDeleteImageCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID, or registry alias, that's associated with the registry that contains the image to delete.")
		ecrPublic_batchDeleteImageCmd.Flags().String("repository-name", "", "The repository in a public registry that contains the image to delete.")
		ecrPublic_batchDeleteImageCmd.MarkFlagRequired("image-ids")
		ecrPublic_batchDeleteImageCmd.MarkFlagRequired("repository-name")
	})
	ecrPublicCmd.AddCommand(ecrPublic_batchDeleteImageCmd)
}
