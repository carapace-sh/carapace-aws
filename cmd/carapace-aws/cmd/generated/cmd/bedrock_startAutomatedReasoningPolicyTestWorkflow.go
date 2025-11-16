package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_startAutomatedReasoningPolicyTestWorkflowCmd = &cobra.Command{
	Use:   "start-automated-reasoning-policy-test-workflow",
	Short: "Initiates a test workflow to validate Automated Reasoning policy tests.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_startAutomatedReasoningPolicyTestWorkflowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_startAutomatedReasoningPolicyTestWorkflowCmd).Standalone()

		bedrock_startAutomatedReasoningPolicyTestWorkflowCmd.Flags().String("build-workflow-id", "", "The build workflow identifier.")
		bedrock_startAutomatedReasoningPolicyTestWorkflowCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier to ensure that the operation completes no more than one time.")
		bedrock_startAutomatedReasoningPolicyTestWorkflowCmd.Flags().String("policy-arn", "", "The Amazon Resource Name (ARN) of the Automated Reasoning policy to test.")
		bedrock_startAutomatedReasoningPolicyTestWorkflowCmd.Flags().String("test-case-ids", "", "The list of test identifiers to run.")
		bedrock_startAutomatedReasoningPolicyTestWorkflowCmd.MarkFlagRequired("build-workflow-id")
		bedrock_startAutomatedReasoningPolicyTestWorkflowCmd.MarkFlagRequired("policy-arn")
	})
	bedrockCmd.AddCommand(bedrock_startAutomatedReasoningPolicyTestWorkflowCmd)
}
