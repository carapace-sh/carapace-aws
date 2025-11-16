package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteModelBiasJobDefinitionCmd = &cobra.Command{
	Use:   "delete-model-bias-job-definition",
	Short: "Deletes an Amazon SageMaker AI model bias job definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteModelBiasJobDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_deleteModelBiasJobDefinitionCmd).Standalone()

		sagemaker_deleteModelBiasJobDefinitionCmd.Flags().String("job-definition-name", "", "The name of the model bias job definition to delete.")
		sagemaker_deleteModelBiasJobDefinitionCmd.MarkFlagRequired("job-definition-name")
	})
	sagemakerCmd.AddCommand(sagemaker_deleteModelBiasJobDefinitionCmd)
}
