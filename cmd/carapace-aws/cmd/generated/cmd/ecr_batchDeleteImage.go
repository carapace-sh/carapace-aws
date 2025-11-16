package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_batchDeleteImageCmd = &cobra.Command{
	Use:   "batch-delete-image",
	Short: "Deletes a list of specified images within a repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_batchDeleteImageCmd).Standalone()

	ecr_batchDeleteImageCmd.Flags().String("image-ids", "", "A list of image ID references that correspond to images to delete.")
	ecr_batchDeleteImageCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID associated with the registry that contains the image to delete.")
	ecr_batchDeleteImageCmd.Flags().String("repository-name", "", "The repository that contains the image to delete.")
	ecr_batchDeleteImageCmd.MarkFlagRequired("image-ids")
	ecr_batchDeleteImageCmd.MarkFlagRequired("repository-name")
	ecrCmd.AddCommand(ecr_batchDeleteImageCmd)
}
