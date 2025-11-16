package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_createCustomModelDeploymentCmd = &cobra.Command{
	Use:   "create-custom-model-deployment",
	Short: "Deploys a custom model for on-demand inference in Amazon Bedrock.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_createCustomModelDeploymentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_createCustomModelDeploymentCmd).Standalone()

		bedrock_createCustomModelDeploymentCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier to ensure that the operation completes no more than one time.")
		bedrock_createCustomModelDeploymentCmd.Flags().String("description", "", "A description for the custom model deployment to help you identify its purpose.")
		bedrock_createCustomModelDeploymentCmd.Flags().String("model-arn", "", "The Amazon Resource Name (ARN) of the custom model to deploy for on-demand inference.")
		bedrock_createCustomModelDeploymentCmd.Flags().String("model-deployment-name", "", "The name for the custom model deployment.")
		bedrock_createCustomModelDeploymentCmd.Flags().String("tags", "", "Tags to assign to the custom model deployment.")
		bedrock_createCustomModelDeploymentCmd.MarkFlagRequired("model-arn")
		bedrock_createCustomModelDeploymentCmd.MarkFlagRequired("model-deployment-name")
	})
	bedrockCmd.AddCommand(bedrock_createCustomModelDeploymentCmd)
}
