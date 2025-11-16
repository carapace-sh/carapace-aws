package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_importHubContentCmd = &cobra.Command{
	Use:   "import-hub-content",
	Short: "Import hub content.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_importHubContentCmd).Standalone()

	sagemaker_importHubContentCmd.Flags().String("document-schema-version", "", "The version of the hub content schema to import.")
	sagemaker_importHubContentCmd.Flags().String("hub-content-description", "", "A description of the hub content to import.")
	sagemaker_importHubContentCmd.Flags().String("hub-content-display-name", "", "The display name of the hub content to import.")
	sagemaker_importHubContentCmd.Flags().String("hub-content-document", "", "The hub content document that describes information about the hub content such as type, associated containers, scripts, and more.")
	sagemaker_importHubContentCmd.Flags().String("hub-content-markdown", "", "A string that provides a description of the hub content.")
	sagemaker_importHubContentCmd.Flags().String("hub-content-name", "", "The name of the hub content to import.")
	sagemaker_importHubContentCmd.Flags().String("hub-content-search-keywords", "", "The searchable keywords of the hub content.")
	sagemaker_importHubContentCmd.Flags().String("hub-content-type", "", "The type of hub content to import.")
	sagemaker_importHubContentCmd.Flags().String("hub-content-version", "", "The version of the hub content to import.")
	sagemaker_importHubContentCmd.Flags().String("hub-name", "", "The name of the hub to import content into.")
	sagemaker_importHubContentCmd.Flags().String("support-status", "", "The status of the hub content resource.")
	sagemaker_importHubContentCmd.Flags().String("tags", "", "Any tags associated with the hub content.")
	sagemaker_importHubContentCmd.MarkFlagRequired("document-schema-version")
	sagemaker_importHubContentCmd.MarkFlagRequired("hub-content-document")
	sagemaker_importHubContentCmd.MarkFlagRequired("hub-content-name")
	sagemaker_importHubContentCmd.MarkFlagRequired("hub-content-type")
	sagemaker_importHubContentCmd.MarkFlagRequired("hub-name")
	sagemakerCmd.AddCommand(sagemaker_importHubContentCmd)
}
