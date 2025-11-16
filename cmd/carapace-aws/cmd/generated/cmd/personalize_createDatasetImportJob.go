package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_createDatasetImportJobCmd = &cobra.Command{
	Use:   "create-dataset-import-job",
	Short: "Creates a job that imports training data from your data source (an Amazon S3 bucket) to an Amazon Personalize dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_createDatasetImportJobCmd).Standalone()

	personalize_createDatasetImportJobCmd.Flags().String("data-source", "", "The Amazon S3 bucket that contains the training data to import.")
	personalize_createDatasetImportJobCmd.Flags().String("dataset-arn", "", "The ARN of the dataset that receives the imported data.")
	personalize_createDatasetImportJobCmd.Flags().String("import-mode", "", "Specify how to add the new records to an existing dataset.")
	personalize_createDatasetImportJobCmd.Flags().String("job-name", "", "The name for the dataset import job.")
	personalize_createDatasetImportJobCmd.Flags().Bool("no-publish-attribution-metrics-to-s3", false, "If you created a metric attribution, specify whether to publish metrics for this import job to Amazon S3")
	personalize_createDatasetImportJobCmd.Flags().Bool("publish-attribution-metrics-to-s3", false, "If you created a metric attribution, specify whether to publish metrics for this import job to Amazon S3")
	personalize_createDatasetImportJobCmd.Flags().String("role-arn", "", "The ARN of the IAM role that has permissions to read from the Amazon S3 data source.")
	personalize_createDatasetImportJobCmd.Flags().String("tags", "", "A list of [tags](https://docs.aws.amazon.com/personalize/latest/dg/tagging-resources.html) to apply to the dataset import job.")
	personalize_createDatasetImportJobCmd.MarkFlagRequired("data-source")
	personalize_createDatasetImportJobCmd.MarkFlagRequired("dataset-arn")
	personalize_createDatasetImportJobCmd.MarkFlagRequired("job-name")
	personalize_createDatasetImportJobCmd.Flag("no-publish-attribution-metrics-to-s3").Hidden = true
	personalize_createDatasetImportJobCmd.MarkFlagRequired("role-arn")
	personalizeCmd.AddCommand(personalize_createDatasetImportJobCmd)
}
