package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createHubContentReferenceCmd = &cobra.Command{
	Use:   "create-hub-content-reference",
	Short: "Create a hub content reference in order to add a model in the JumpStart public hub to a private hub.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createHubContentReferenceCmd).Standalone()

	sagemaker_createHubContentReferenceCmd.Flags().String("hub-content-name", "", "The name of the hub content to reference.")
	sagemaker_createHubContentReferenceCmd.Flags().String("hub-name", "", "The name of the hub to add the hub content reference to.")
	sagemaker_createHubContentReferenceCmd.Flags().String("min-version", "", "The minimum version of the hub content to reference.")
	sagemaker_createHubContentReferenceCmd.Flags().String("sage-maker-public-hub-content-arn", "", "The ARN of the public hub content to reference.")
	sagemaker_createHubContentReferenceCmd.Flags().String("tags", "", "Any tags associated with the hub content to reference.")
	sagemaker_createHubContentReferenceCmd.MarkFlagRequired("hub-name")
	sagemaker_createHubContentReferenceCmd.MarkFlagRequired("sage-maker-public-hub-content-arn")
	sagemakerCmd.AddCommand(sagemaker_createHubContentReferenceCmd)
}
