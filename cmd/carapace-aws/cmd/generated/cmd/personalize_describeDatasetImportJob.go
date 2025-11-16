package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_describeDatasetImportJobCmd = &cobra.Command{
	Use:   "describe-dataset-import-job",
	Short: "Describes the dataset import job created by [CreateDatasetImportJob](https://docs.aws.amazon.com/personalize/latest/dg/API_CreateDatasetImportJob.html), including the import job status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_describeDatasetImportJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(personalize_describeDatasetImportJobCmd).Standalone()

		personalize_describeDatasetImportJobCmd.Flags().String("dataset-import-job-arn", "", "The Amazon Resource Name (ARN) of the dataset import job to describe.")
		personalize_describeDatasetImportJobCmd.MarkFlagRequired("dataset-import-job-arn")
	})
	personalizeCmd.AddCommand(personalize_describeDatasetImportJobCmd)
}
