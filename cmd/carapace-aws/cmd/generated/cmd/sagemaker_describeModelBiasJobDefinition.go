package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeModelBiasJobDefinitionCmd = &cobra.Command{
	Use:   "describe-model-bias-job-definition",
	Short: "Returns a description of a model bias job definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeModelBiasJobDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_describeModelBiasJobDefinitionCmd).Standalone()

		sagemaker_describeModelBiasJobDefinitionCmd.Flags().String("job-definition-name", "", "The name of the model bias job definition.")
		sagemaker_describeModelBiasJobDefinitionCmd.MarkFlagRequired("job-definition-name")
	})
	sagemakerCmd.AddCommand(sagemaker_describeModelBiasJobDefinitionCmd)
}
