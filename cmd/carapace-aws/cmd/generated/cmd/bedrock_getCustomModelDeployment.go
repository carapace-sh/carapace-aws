package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_getCustomModelDeploymentCmd = &cobra.Command{
	Use:   "get-custom-model-deployment",
	Short: "Retrieves information about a custom model deployment, including its status, configuration, and metadata.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_getCustomModelDeploymentCmd).Standalone()

	bedrock_getCustomModelDeploymentCmd.Flags().String("custom-model-deployment-identifier", "", "The Amazon Resource Name (ARN) or name of the custom model deployment to retrieve information about.")
	bedrock_getCustomModelDeploymentCmd.MarkFlagRequired("custom-model-deployment-identifier")
	bedrockCmd.AddCommand(bedrock_getCustomModelDeploymentCmd)
}
