package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_sendWorkflowStepActionCmd = &cobra.Command{
	Use:   "send-workflow-step-action",
	Short: "Pauses or resumes image creation when the associated workflow runs a `WaitForAction` step.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_sendWorkflowStepActionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_sendWorkflowStepActionCmd).Standalone()

		imagebuilder_sendWorkflowStepActionCmd.Flags().String("action", "", "The action for the image creation process to take while a workflow `WaitForAction` step waits for an asynchronous action to complete.")
		imagebuilder_sendWorkflowStepActionCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure idempotency of the request.")
		imagebuilder_sendWorkflowStepActionCmd.Flags().String("image-build-version-arn", "", "The Amazon Resource Name (ARN) of the image build version to send action for.")
		imagebuilder_sendWorkflowStepActionCmd.Flags().String("reason", "", "The reason why this action is sent.")
		imagebuilder_sendWorkflowStepActionCmd.Flags().String("step-execution-id", "", "Uniquely identifies the workflow step that sent the step action.")
		imagebuilder_sendWorkflowStepActionCmd.MarkFlagRequired("action")
		imagebuilder_sendWorkflowStepActionCmd.MarkFlagRequired("client-token")
		imagebuilder_sendWorkflowStepActionCmd.MarkFlagRequired("image-build-version-arn")
		imagebuilder_sendWorkflowStepActionCmd.MarkFlagRequired("step-execution-id")
	})
	imagebuilderCmd.AddCommand(imagebuilder_sendWorkflowStepActionCmd)
}
