package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_uploadLayerPartCmd = &cobra.Command{
	Use:   "upload-layer-part",
	Short: "Uploads an image layer part to Amazon ECR.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_uploadLayerPartCmd).Standalone()

	ecr_uploadLayerPartCmd.Flags().String("layer-part-blob", "", "The base64-encoded layer part payload.")
	ecr_uploadLayerPartCmd.Flags().String("part-first-byte", "", "The position of the first byte of the layer part witin the overall image layer.")
	ecr_uploadLayerPartCmd.Flags().String("part-last-byte", "", "The position of the last byte of the layer part within the overall image layer.")
	ecr_uploadLayerPartCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID associated with the registry to which you are uploading layer parts.")
	ecr_uploadLayerPartCmd.Flags().String("repository-name", "", "The name of the repository to which you are uploading layer parts.")
	ecr_uploadLayerPartCmd.Flags().String("upload-id", "", "The upload ID from a previous [InitiateLayerUpload]() operation to associate with the layer part upload.")
	ecr_uploadLayerPartCmd.MarkFlagRequired("layer-part-blob")
	ecr_uploadLayerPartCmd.MarkFlagRequired("part-first-byte")
	ecr_uploadLayerPartCmd.MarkFlagRequired("part-last-byte")
	ecr_uploadLayerPartCmd.MarkFlagRequired("repository-name")
	ecr_uploadLayerPartCmd.MarkFlagRequired("upload-id")
	ecrCmd.AddCommand(ecr_uploadLayerPartCmd)
}
