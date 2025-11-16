package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createFlowDefinitionCmd = &cobra.Command{
	Use:   "create-flow-definition",
	Short: "Creates a flow definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createFlowDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_createFlowDefinitionCmd).Standalone()

		sagemaker_createFlowDefinitionCmd.Flags().String("flow-definition-name", "", "The name of your flow definition.")
		sagemaker_createFlowDefinitionCmd.Flags().String("human-loop-activation-config", "", "An object containing information about the events that trigger a human workflow.")
		sagemaker_createFlowDefinitionCmd.Flags().String("human-loop-config", "", "An object containing information about the tasks the human reviewers will perform.")
		sagemaker_createFlowDefinitionCmd.Flags().String("human-loop-request-source", "", "Container for configuring the source of human task requests.")
		sagemaker_createFlowDefinitionCmd.Flags().String("output-config", "", "An object containing information about where the human review results will be uploaded.")
		sagemaker_createFlowDefinitionCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the role needed to call other services on your behalf.")
		sagemaker_createFlowDefinitionCmd.Flags().String("tags", "", "An array of key-value pairs that contain metadata to help you categorize and organize a flow definition.")
		sagemaker_createFlowDefinitionCmd.MarkFlagRequired("flow-definition-name")
		sagemaker_createFlowDefinitionCmd.MarkFlagRequired("output-config")
		sagemaker_createFlowDefinitionCmd.MarkFlagRequired("role-arn")
	})
	sagemakerCmd.AddCommand(sagemaker_createFlowDefinitionCmd)
}
