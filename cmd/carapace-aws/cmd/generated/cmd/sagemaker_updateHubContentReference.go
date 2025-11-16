package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateHubContentReferenceCmd = &cobra.Command{
	Use:   "update-hub-content-reference",
	Short: "Updates the contents of a SageMaker hub for a `ModelReference` resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateHubContentReferenceCmd).Standalone()

	sagemaker_updateHubContentReferenceCmd.Flags().String("hub-content-name", "", "The name of the hub content resource that you want to update.")
	sagemaker_updateHubContentReferenceCmd.Flags().String("hub-content-type", "", "The content type of the resource that you want to update.")
	sagemaker_updateHubContentReferenceCmd.Flags().String("hub-name", "", "The name of the SageMaker hub that contains the hub content you want to update.")
	sagemaker_updateHubContentReferenceCmd.Flags().String("min-version", "", "The minimum hub content version of the referenced model that you want to use.")
	sagemaker_updateHubContentReferenceCmd.MarkFlagRequired("hub-content-name")
	sagemaker_updateHubContentReferenceCmd.MarkFlagRequired("hub-content-type")
	sagemaker_updateHubContentReferenceCmd.MarkFlagRequired("hub-name")
	sagemakerCmd.AddCommand(sagemaker_updateHubContentReferenceCmd)
}
