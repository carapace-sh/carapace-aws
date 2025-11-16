package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeDataQualityJobDefinitionCmd = &cobra.Command{
	Use:   "describe-data-quality-job-definition",
	Short: "Gets the details of a data quality monitoring job definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeDataQualityJobDefinitionCmd).Standalone()

	sagemaker_describeDataQualityJobDefinitionCmd.Flags().String("job-definition-name", "", "The name of the data quality monitoring job definition to describe.")
	sagemaker_describeDataQualityJobDefinitionCmd.MarkFlagRequired("job-definition-name")
	sagemakerCmd.AddCommand(sagemaker_describeDataQualityJobDefinitionCmd)
}
