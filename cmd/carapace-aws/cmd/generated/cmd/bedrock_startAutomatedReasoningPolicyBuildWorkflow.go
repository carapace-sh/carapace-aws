package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_startAutomatedReasoningPolicyBuildWorkflowCmd = &cobra.Command{
	Use:   "start-automated-reasoning-policy-build-workflow",
	Short: "Starts a new build workflow for an Automated Reasoning policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_startAutomatedReasoningPolicyBuildWorkflowCmd).Standalone()

	bedrock_startAutomatedReasoningPolicyBuildWorkflowCmd.Flags().String("build-workflow-type", "", "The type of build workflow to start (e.g., DOCUMENT\\_INGESTION for processing new documents, POLICY\\_REPAIR for fixing existing policies).")
	bedrock_startAutomatedReasoningPolicyBuildWorkflowCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier to ensure that the operation completes no more than once.")
	bedrock_startAutomatedReasoningPolicyBuildWorkflowCmd.Flags().String("policy-arn", "", "The Amazon Resource Name (ARN) of the Automated Reasoning policy for which to start the build workflow.")
	bedrock_startAutomatedReasoningPolicyBuildWorkflowCmd.Flags().String("source-content", "", "The source content for the build workflow, such as documents to analyze or repair instructions for existing policies.")
	bedrock_startAutomatedReasoningPolicyBuildWorkflowCmd.MarkFlagRequired("build-workflow-type")
	bedrock_startAutomatedReasoningPolicyBuildWorkflowCmd.MarkFlagRequired("policy-arn")
	bedrock_startAutomatedReasoningPolicyBuildWorkflowCmd.MarkFlagRequired("source-content")
	bedrockCmd.AddCommand(bedrock_startAutomatedReasoningPolicyBuildWorkflowCmd)
}
