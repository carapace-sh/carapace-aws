package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_initiateLayerUploadCmd = &cobra.Command{
	Use:   "initiate-layer-upload",
	Short: "Notifies Amazon ECR that you intend to upload an image layer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_initiateLayerUploadCmd).Standalone()

	ecr_initiateLayerUploadCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID associated with the registry to which you intend to upload layers.")
	ecr_initiateLayerUploadCmd.Flags().String("repository-name", "", "The name of the repository to which you intend to upload layers.")
	ecr_initiateLayerUploadCmd.MarkFlagRequired("repository-name")
	ecrCmd.AddCommand(ecr_initiateLayerUploadCmd)
}
