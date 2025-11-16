package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteHubContentCmd = &cobra.Command{
	Use:   "delete-hub-content",
	Short: "Delete the contents of a hub.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteHubContentCmd).Standalone()

	sagemaker_deleteHubContentCmd.Flags().String("hub-content-name", "", "The name of the content that you want to delete from a hub.")
	sagemaker_deleteHubContentCmd.Flags().String("hub-content-type", "", "The type of content that you want to delete from a hub.")
	sagemaker_deleteHubContentCmd.Flags().String("hub-content-version", "", "The version of the content that you want to delete from a hub.")
	sagemaker_deleteHubContentCmd.Flags().String("hub-name", "", "The name of the hub that you want to delete content in.")
	sagemaker_deleteHubContentCmd.MarkFlagRequired("hub-content-name")
	sagemaker_deleteHubContentCmd.MarkFlagRequired("hub-content-type")
	sagemaker_deleteHubContentCmd.MarkFlagRequired("hub-content-version")
	sagemaker_deleteHubContentCmd.MarkFlagRequired("hub-name")
	sagemakerCmd.AddCommand(sagemaker_deleteHubContentCmd)
}
