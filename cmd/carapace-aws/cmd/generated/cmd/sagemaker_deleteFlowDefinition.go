package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteFlowDefinitionCmd = &cobra.Command{
	Use:   "delete-flow-definition",
	Short: "Deletes the specified flow definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteFlowDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_deleteFlowDefinitionCmd).Standalone()

		sagemaker_deleteFlowDefinitionCmd.Flags().String("flow-definition-name", "", "The name of the flow definition you are deleting.")
		sagemaker_deleteFlowDefinitionCmd.MarkFlagRequired("flow-definition-name")
	})
	sagemakerCmd.AddCommand(sagemaker_deleteFlowDefinitionCmd)
}
