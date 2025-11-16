package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_getAutomatedReasoningPolicyBuildWorkflowResultAssetsCmd = &cobra.Command{
	Use:   "get-automated-reasoning-policy-build-workflow-result-assets",
	Short: "Retrieves the resulting assets from a completed Automated Reasoning policy build workflow, including build logs, quality reports, and generated policy artifacts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_getAutomatedReasoningPolicyBuildWorkflowResultAssetsCmd).Standalone()

	bedrock_getAutomatedReasoningPolicyBuildWorkflowResultAssetsCmd.Flags().String("asset-type", "", "The type of asset to retrieve (e.g., BUILD\\_LOG, QUALITY\\_REPORT, POLICY\\_DEFINITION).")
	bedrock_getAutomatedReasoningPolicyBuildWorkflowResultAssetsCmd.Flags().String("build-workflow-id", "", "The unique identifier of the build workflow whose result assets you want to retrieve.")
	bedrock_getAutomatedReasoningPolicyBuildWorkflowResultAssetsCmd.Flags().String("policy-arn", "", "The Amazon Resource Name (ARN) of the Automated Reasoning policy whose build workflow assets you want to retrieve.")
	bedrock_getAutomatedReasoningPolicyBuildWorkflowResultAssetsCmd.MarkFlagRequired("asset-type")
	bedrock_getAutomatedReasoningPolicyBuildWorkflowResultAssetsCmd.MarkFlagRequired("build-workflow-id")
	bedrock_getAutomatedReasoningPolicyBuildWorkflowResultAssetsCmd.MarkFlagRequired("policy-arn")
	bedrockCmd.AddCommand(bedrock_getAutomatedReasoningPolicyBuildWorkflowResultAssetsCmd)
}
