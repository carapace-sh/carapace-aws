package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteModelQualityJobDefinitionCmd = &cobra.Command{
	Use:   "delete-model-quality-job-definition",
	Short: "Deletes the secified model quality monitoring job definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteModelQualityJobDefinitionCmd).Standalone()

	sagemaker_deleteModelQualityJobDefinitionCmd.Flags().String("job-definition-name", "", "The name of the model quality monitoring job definition to delete.")
	sagemaker_deleteModelQualityJobDefinitionCmd.MarkFlagRequired("job-definition-name")
	sagemakerCmd.AddCommand(sagemaker_deleteModelQualityJobDefinitionCmd)
}
