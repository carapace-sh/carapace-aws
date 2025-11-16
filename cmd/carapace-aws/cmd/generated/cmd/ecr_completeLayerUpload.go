package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_completeLayerUploadCmd = &cobra.Command{
	Use:   "complete-layer-upload",
	Short: "Informs Amazon ECR that the image layer upload has completed for a specified registry, repository name, and upload ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_completeLayerUploadCmd).Standalone()

	ecr_completeLayerUploadCmd.Flags().String("layer-digests", "", "The `sha256` digest of the image layer.")
	ecr_completeLayerUploadCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID associated with the registry to which to upload layers.")
	ecr_completeLayerUploadCmd.Flags().String("repository-name", "", "The name of the repository to associate with the image layer.")
	ecr_completeLayerUploadCmd.Flags().String("upload-id", "", "The upload ID from a previous [InitiateLayerUpload]() operation to associate with the image layer.")
	ecr_completeLayerUploadCmd.MarkFlagRequired("layer-digests")
	ecr_completeLayerUploadCmd.MarkFlagRequired("repository-name")
	ecr_completeLayerUploadCmd.MarkFlagRequired("upload-id")
	ecrCmd.AddCommand(ecr_completeLayerUploadCmd)
}
