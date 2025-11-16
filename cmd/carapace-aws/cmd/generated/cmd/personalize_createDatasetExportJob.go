package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_createDatasetExportJobCmd = &cobra.Command{
	Use:   "create-dataset-export-job",
	Short: "Creates a job that exports data from your dataset to an Amazon S3 bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_createDatasetExportJobCmd).Standalone()

	personalize_createDatasetExportJobCmd.Flags().String("dataset-arn", "", "The Amazon Resource Name (ARN) of the dataset that contains the data to export.")
	personalize_createDatasetExportJobCmd.Flags().String("ingestion-mode", "", "The data to export, based on how you imported the data.")
	personalize_createDatasetExportJobCmd.Flags().String("job-name", "", "The name for the dataset export job.")
	personalize_createDatasetExportJobCmd.Flags().String("job-output", "", "The path to the Amazon S3 bucket where the job's output is stored.")
	personalize_createDatasetExportJobCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the IAM service role that has permissions to add data to your output Amazon S3 bucket.")
	personalize_createDatasetExportJobCmd.Flags().String("tags", "", "A list of [tags](https://docs.aws.amazon.com/personalize/latest/dg/tagging-resources.html) to apply to the dataset export job.")
	personalize_createDatasetExportJobCmd.MarkFlagRequired("dataset-arn")
	personalize_createDatasetExportJobCmd.MarkFlagRequired("job-name")
	personalize_createDatasetExportJobCmd.MarkFlagRequired("job-output")
	personalize_createDatasetExportJobCmd.MarkFlagRequired("role-arn")
	personalizeCmd.AddCommand(personalize_createDatasetExportJobCmd)
}
