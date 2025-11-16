package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeModelQualityJobDefinitionCmd = &cobra.Command{
	Use:   "describe-model-quality-job-definition",
	Short: "Returns a description of a model quality job definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeModelQualityJobDefinitionCmd).Standalone()

	sagemaker_describeModelQualityJobDefinitionCmd.Flags().String("job-definition-name", "", "The name of the model quality job.")
	sagemaker_describeModelQualityJobDefinitionCmd.MarkFlagRequired("job-definition-name")
	sagemakerCmd.AddCommand(sagemaker_describeModelQualityJobDefinitionCmd)
}
