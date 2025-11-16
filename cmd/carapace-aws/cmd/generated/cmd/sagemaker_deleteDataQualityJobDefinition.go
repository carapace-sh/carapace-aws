package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteDataQualityJobDefinitionCmd = &cobra.Command{
	Use:   "delete-data-quality-job-definition",
	Short: "Deletes a data quality monitoring job definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteDataQualityJobDefinitionCmd).Standalone()

	sagemaker_deleteDataQualityJobDefinitionCmd.Flags().String("job-definition-name", "", "The name of the data quality monitoring job definition to delete.")
	sagemaker_deleteDataQualityJobDefinitionCmd.MarkFlagRequired("job-definition-name")
	sagemakerCmd.AddCommand(sagemaker_deleteDataQualityJobDefinitionCmd)
}
