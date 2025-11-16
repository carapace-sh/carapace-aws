package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeFlowDefinitionCmd = &cobra.Command{
	Use:   "describe-flow-definition",
	Short: "Returns information about the specified flow definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeFlowDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_describeFlowDefinitionCmd).Standalone()

		sagemaker_describeFlowDefinitionCmd.Flags().String("flow-definition-name", "", "The name of the flow definition.")
		sagemaker_describeFlowDefinitionCmd.MarkFlagRequired("flow-definition-name")
	})
	sagemakerCmd.AddCommand(sagemaker_describeFlowDefinitionCmd)
}
