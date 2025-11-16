package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_deleteCustomModelDeploymentCmd = &cobra.Command{
	Use:   "delete-custom-model-deployment",
	Short: "Deletes a custom model deployment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_deleteCustomModelDeploymentCmd).Standalone()

	bedrock_deleteCustomModelDeploymentCmd.Flags().String("custom-model-deployment-identifier", "", "The Amazon Resource Name (ARN) or name of the custom model deployment to delete.")
	bedrock_deleteCustomModelDeploymentCmd.MarkFlagRequired("custom-model-deployment-identifier")
	bedrockCmd.AddCommand(bedrock_deleteCustomModelDeploymentCmd)
}
