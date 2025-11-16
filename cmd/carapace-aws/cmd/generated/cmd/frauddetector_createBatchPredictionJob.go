package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_createBatchPredictionJobCmd = &cobra.Command{
	Use:   "create-batch-prediction-job",
	Short: "Creates a batch prediction job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_createBatchPredictionJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_createBatchPredictionJobCmd).Standalone()

		frauddetector_createBatchPredictionJobCmd.Flags().String("detector-name", "", "The name of the detector.")
		frauddetector_createBatchPredictionJobCmd.Flags().String("detector-version", "", "The detector version.")
		frauddetector_createBatchPredictionJobCmd.Flags().String("event-type-name", "", "The name of the event type.")
		frauddetector_createBatchPredictionJobCmd.Flags().String("iam-role-arn", "", "The ARN of the IAM role to use for this job request.")
		frauddetector_createBatchPredictionJobCmd.Flags().String("input-path", "", "The Amazon S3 location of your training file.")
		frauddetector_createBatchPredictionJobCmd.Flags().String("job-id", "", "The ID of the batch prediction job.")
		frauddetector_createBatchPredictionJobCmd.Flags().String("output-path", "", "The Amazon S3 location of your output file.")
		frauddetector_createBatchPredictionJobCmd.Flags().String("tags", "", "A collection of key and value pairs.")
		frauddetector_createBatchPredictionJobCmd.MarkFlagRequired("detector-name")
		frauddetector_createBatchPredictionJobCmd.MarkFlagRequired("event-type-name")
		frauddetector_createBatchPredictionJobCmd.MarkFlagRequired("iam-role-arn")
		frauddetector_createBatchPredictionJobCmd.MarkFlagRequired("input-path")
		frauddetector_createBatchPredictionJobCmd.MarkFlagRequired("job-id")
		frauddetector_createBatchPredictionJobCmd.MarkFlagRequired("output-path")
	})
	frauddetectorCmd.AddCommand(frauddetector_createBatchPredictionJobCmd)
}
