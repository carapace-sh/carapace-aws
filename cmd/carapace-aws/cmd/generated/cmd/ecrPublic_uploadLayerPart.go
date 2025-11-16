package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecrPublic_uploadLayerPartCmd = &cobra.Command{
	Use:   "upload-layer-part",
	Short: "Uploads an image layer part to Amazon ECR.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecrPublic_uploadLayerPartCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecrPublic_uploadLayerPartCmd).Standalone()

		ecrPublic_uploadLayerPartCmd.Flags().String("layer-part-blob", "", "The base64-encoded layer part payload.")
		ecrPublic_uploadLayerPartCmd.Flags().String("part-first-byte", "", "The position of the first byte of the layer part witin the overall image layer.")
		ecrPublic_uploadLayerPartCmd.Flags().String("part-last-byte", "", "The position of the last byte of the layer part within the overall image layer.")
		ecrPublic_uploadLayerPartCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID, or registry alias, that's associated with the registry that you're uploading layer parts to.")
		ecrPublic_uploadLayerPartCmd.Flags().String("repository-name", "", "The name of the repository that you're uploading layer parts to.")
		ecrPublic_uploadLayerPartCmd.Flags().String("upload-id", "", "The upload ID from a previous [InitiateLayerUpload]() operation to associate with the layer part upload.")
		ecrPublic_uploadLayerPartCmd.MarkFlagRequired("layer-part-blob")
		ecrPublic_uploadLayerPartCmd.MarkFlagRequired("part-first-byte")
		ecrPublic_uploadLayerPartCmd.MarkFlagRequired("part-last-byte")
		ecrPublic_uploadLayerPartCmd.MarkFlagRequired("repository-name")
		ecrPublic_uploadLayerPartCmd.MarkFlagRequired("upload-id")
	})
	ecrPublicCmd.AddCommand(ecrPublic_uploadLayerPartCmd)
}
