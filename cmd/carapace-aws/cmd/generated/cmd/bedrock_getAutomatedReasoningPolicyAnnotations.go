package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_getAutomatedReasoningPolicyAnnotationsCmd = &cobra.Command{
	Use:   "get-automated-reasoning-policy-annotations",
	Short: "Retrieves the current annotations for an Automated Reasoning policy build workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_getAutomatedReasoningPolicyAnnotationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_getAutomatedReasoningPolicyAnnotationsCmd).Standalone()

		bedrock_getAutomatedReasoningPolicyAnnotationsCmd.Flags().String("build-workflow-id", "", "The unique identifier of the build workflow whose annotations you want to retrieve.")
		bedrock_getAutomatedReasoningPolicyAnnotationsCmd.Flags().String("policy-arn", "", "The Amazon Resource Name (ARN) of the Automated Reasoning policy whose annotations you want to retrieve.")
		bedrock_getAutomatedReasoningPolicyAnnotationsCmd.MarkFlagRequired("build-workflow-id")
		bedrock_getAutomatedReasoningPolicyAnnotationsCmd.MarkFlagRequired("policy-arn")
	})
	bedrockCmd.AddCommand(bedrock_getAutomatedReasoningPolicyAnnotationsCmd)
}
