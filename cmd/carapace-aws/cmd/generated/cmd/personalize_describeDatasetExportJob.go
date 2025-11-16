package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_describeDatasetExportJobCmd = &cobra.Command{
	Use:   "describe-dataset-export-job",
	Short: "Describes the dataset export job created by [CreateDatasetExportJob](https://docs.aws.amazon.com/personalize/latest/dg/API_CreateDatasetExportJob.html), including the export job status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_describeDatasetExportJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(personalize_describeDatasetExportJobCmd).Standalone()

		personalize_describeDatasetExportJobCmd.Flags().String("dataset-export-job-arn", "", "The Amazon Resource Name (ARN) of the dataset export job to describe.")
		personalize_describeDatasetExportJobCmd.MarkFlagRequired("dataset-export-job-arn")
	})
	personalizeCmd.AddCommand(personalize_describeDatasetExportJobCmd)
}
