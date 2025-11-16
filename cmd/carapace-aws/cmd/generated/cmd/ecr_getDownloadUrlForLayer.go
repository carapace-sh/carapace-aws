package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_getDownloadUrlForLayerCmd = &cobra.Command{
	Use:   "get-download-url-for-layer",
	Short: "Retrieves the pre-signed Amazon S3 download URL corresponding to an image layer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_getDownloadUrlForLayerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecr_getDownloadUrlForLayerCmd).Standalone()

		ecr_getDownloadUrlForLayerCmd.Flags().String("layer-digest", "", "The digest of the image layer to download.")
		ecr_getDownloadUrlForLayerCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID associated with the registry that contains the image layer to download.")
		ecr_getDownloadUrlForLayerCmd.Flags().String("repository-name", "", "The name of the repository that is associated with the image layer to download.")
		ecr_getDownloadUrlForLayerCmd.MarkFlagRequired("layer-digest")
		ecr_getDownloadUrlForLayerCmd.MarkFlagRequired("repository-name")
	})
	ecrCmd.AddCommand(ecr_getDownloadUrlForLayerCmd)
}
