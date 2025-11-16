package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecrPublic_initiateLayerUploadCmd = &cobra.Command{
	Use:   "initiate-layer-upload",
	Short: "Notifies Amazon ECR that you intend to upload an image layer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecrPublic_initiateLayerUploadCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecrPublic_initiateLayerUploadCmd).Standalone()

		ecrPublic_initiateLayerUploadCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID, or registry alias, that's associated with the registry to which you intend to upload layers.")
		ecrPublic_initiateLayerUploadCmd.Flags().String("repository-name", "", "The name of the repository that you want to upload layers to.")
		ecrPublic_initiateLayerUploadCmd.MarkFlagRequired("repository-name")
	})
	ecrPublicCmd.AddCommand(ecrPublic_initiateLayerUploadCmd)
}
