package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_batchGetImageCmd = &cobra.Command{
	Use:   "batch-get-image",
	Short: "Gets detailed information for an image.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_batchGetImageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecr_batchGetImageCmd).Standalone()

		ecr_batchGetImageCmd.Flags().String("accepted-media-types", "", "The accepted media types for the request.")
		ecr_batchGetImageCmd.Flags().String("image-ids", "", "A list of image ID references that correspond to images to describe.")
		ecr_batchGetImageCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID associated with the registry that contains the images to describe.")
		ecr_batchGetImageCmd.Flags().String("repository-name", "", "The repository that contains the images to describe.")
		ecr_batchGetImageCmd.MarkFlagRequired("image-ids")
		ecr_batchGetImageCmd.MarkFlagRequired("repository-name")
	})
	ecrCmd.AddCommand(ecr_batchGetImageCmd)
}
