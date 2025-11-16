package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_createBatchImportJobCmd = &cobra.Command{
	Use:   "create-batch-import-job",
	Short: "Creates a batch import job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_createBatchImportJobCmd).Standalone()

	frauddetector_createBatchImportJobCmd.Flags().String("event-type-name", "", "The name of the event type.")
	frauddetector_createBatchImportJobCmd.Flags().String("iam-role-arn", "", "The ARN of the IAM role created for Amazon S3 bucket that holds your data file.")
	frauddetector_createBatchImportJobCmd.Flags().String("input-path", "", "The URI that points to the Amazon S3 location of your data file.")
	frauddetector_createBatchImportJobCmd.Flags().String("job-id", "", "The ID of the batch import job.")
	frauddetector_createBatchImportJobCmd.Flags().String("output-path", "", "The URI that points to the Amazon S3 location for storing your results.")
	frauddetector_createBatchImportJobCmd.Flags().String("tags", "", "A collection of key-value pairs associated with this request.")
	frauddetector_createBatchImportJobCmd.MarkFlagRequired("event-type-name")
	frauddetector_createBatchImportJobCmd.MarkFlagRequired("iam-role-arn")
	frauddetector_createBatchImportJobCmd.MarkFlagRequired("input-path")
	frauddetector_createBatchImportJobCmd.MarkFlagRequired("job-id")
	frauddetector_createBatchImportJobCmd.MarkFlagRequired("output-path")
	frauddetectorCmd.AddCommand(frauddetector_createBatchImportJobCmd)
}
