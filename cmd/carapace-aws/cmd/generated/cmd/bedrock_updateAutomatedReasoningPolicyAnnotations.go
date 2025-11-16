package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_updateAutomatedReasoningPolicyAnnotationsCmd = &cobra.Command{
	Use:   "update-automated-reasoning-policy-annotations",
	Short: "Updates the annotations for an Automated Reasoning policy build workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_updateAutomatedReasoningPolicyAnnotationsCmd).Standalone()

	bedrock_updateAutomatedReasoningPolicyAnnotationsCmd.Flags().String("annotations", "", "The updated annotations containing modified rules, variables, and types for the policy.")
	bedrock_updateAutomatedReasoningPolicyAnnotationsCmd.Flags().String("build-workflow-id", "", "The unique identifier of the build workflow whose annotations you want to update.")
	bedrock_updateAutomatedReasoningPolicyAnnotationsCmd.Flags().String("last-updated-annotation-set-hash", "", "The hash value of the annotation set that you're updating.")
	bedrock_updateAutomatedReasoningPolicyAnnotationsCmd.Flags().String("policy-arn", "", "The Amazon Resource Name (ARN) of the Automated Reasoning policy whose annotations you want to update.")
	bedrock_updateAutomatedReasoningPolicyAnnotationsCmd.MarkFlagRequired("annotations")
	bedrock_updateAutomatedReasoningPolicyAnnotationsCmd.MarkFlagRequired("build-workflow-id")
	bedrock_updateAutomatedReasoningPolicyAnnotationsCmd.MarkFlagRequired("last-updated-annotation-set-hash")
	bedrock_updateAutomatedReasoningPolicyAnnotationsCmd.MarkFlagRequired("policy-arn")
	bedrockCmd.AddCommand(bedrock_updateAutomatedReasoningPolicyAnnotationsCmd)
}
