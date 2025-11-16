package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateHubContentCmd = &cobra.Command{
	Use:   "update-hub-content",
	Short: "Updates SageMaker hub content (either a `Model` or `Notebook` resource).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateHubContentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_updateHubContentCmd).Standalone()

		sagemaker_updateHubContentCmd.Flags().String("hub-content-description", "", "The description of the hub content.")
		sagemaker_updateHubContentCmd.Flags().String("hub-content-display-name", "", "The display name of the hub content.")
		sagemaker_updateHubContentCmd.Flags().String("hub-content-markdown", "", "A string that provides a description of the hub content.")
		sagemaker_updateHubContentCmd.Flags().String("hub-content-name", "", "The name of the hub content resource that you want to update.")
		sagemaker_updateHubContentCmd.Flags().String("hub-content-search-keywords", "", "The searchable keywords of the hub content.")
		sagemaker_updateHubContentCmd.Flags().String("hub-content-type", "", "The content type of the resource that you want to update.")
		sagemaker_updateHubContentCmd.Flags().String("hub-content-version", "", "The hub content version that you want to update.")
		sagemaker_updateHubContentCmd.Flags().String("hub-name", "", "The name of the SageMaker hub that contains the hub content you want to update.")
		sagemaker_updateHubContentCmd.Flags().String("support-status", "", "Indicates the current status of the hub content resource.")
		sagemaker_updateHubContentCmd.MarkFlagRequired("hub-content-name")
		sagemaker_updateHubContentCmd.MarkFlagRequired("hub-content-type")
		sagemaker_updateHubContentCmd.MarkFlagRequired("hub-content-version")
		sagemaker_updateHubContentCmd.MarkFlagRequired("hub-name")
	})
	sagemakerCmd.AddCommand(sagemaker_updateHubContentCmd)
}
