package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteHubContentReferenceCmd = &cobra.Command{
	Use:   "delete-hub-content-reference",
	Short: "Delete a hub content reference in order to remove a model from a private hub.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteHubContentReferenceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_deleteHubContentReferenceCmd).Standalone()

		sagemaker_deleteHubContentReferenceCmd.Flags().String("hub-content-name", "", "The name of the hub content to delete.")
		sagemaker_deleteHubContentReferenceCmd.Flags().String("hub-content-type", "", "The type of hub content reference to delete.")
		sagemaker_deleteHubContentReferenceCmd.Flags().String("hub-name", "", "The name of the hub to delete the hub content reference from.")
		sagemaker_deleteHubContentReferenceCmd.MarkFlagRequired("hub-content-name")
		sagemaker_deleteHubContentReferenceCmd.MarkFlagRequired("hub-content-type")
		sagemaker_deleteHubContentReferenceCmd.MarkFlagRequired("hub-name")
	})
	sagemakerCmd.AddCommand(sagemaker_deleteHubContentReferenceCmd)
}
