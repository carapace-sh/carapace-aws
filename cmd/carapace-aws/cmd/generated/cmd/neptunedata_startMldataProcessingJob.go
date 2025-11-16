package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_startMldataProcessingJobCmd = &cobra.Command{
	Use:   "start-mldata-processing-job",
	Short: "Creates a new Neptune ML data processing job for processing the graph data exported from Neptune for training.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_startMldataProcessingJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptunedata_startMldataProcessingJobCmd).Standalone()

		neptunedata_startMldataProcessingJobCmd.Flags().String("config-file-name", "", "A data specification file that describes how to load the exported graph data for training.")
		neptunedata_startMldataProcessingJobCmd.Flags().String("id", "", "A unique identifier for the new job.")
		neptunedata_startMldataProcessingJobCmd.Flags().String("input-data-s3-location", "", "The URI of the Amazon S3 location where you want SageMaker to download the data needed to run the data processing job.")
		neptunedata_startMldataProcessingJobCmd.Flags().String("model-type", "", "One of the two model types that Neptune ML currently supports: heterogeneous graph models (`heterogeneous`), and knowledge graph (`kge`).")
		neptunedata_startMldataProcessingJobCmd.Flags().String("neptune-iam-role-arn", "", "The Amazon Resource Name (ARN) of an IAM role that SageMaker can assume to perform tasks on your behalf.")
		neptunedata_startMldataProcessingJobCmd.Flags().String("previous-data-processing-job-id", "", "The job ID of a completed data processing job run on an earlier version of the data.")
		neptunedata_startMldataProcessingJobCmd.Flags().String("processed-data-s3-location", "", "The URI of the Amazon S3 location where you want SageMaker to save the results of a data processing job.")
		neptunedata_startMldataProcessingJobCmd.Flags().String("processing-instance-type", "", "The type of ML instance used during data processing.")
		neptunedata_startMldataProcessingJobCmd.Flags().String("processing-instance-volume-size-in-gb", "", "The disk volume size of the processing instance.")
		neptunedata_startMldataProcessingJobCmd.Flags().String("processing-time-out-in-seconds", "", "Timeout in seconds for the data processing job.")
		neptunedata_startMldataProcessingJobCmd.Flags().String("s3-output-encryption-kmskey", "", "The Amazon Key Management Service (Amazon KMS) key that SageMaker uses to encrypt the output of the processing job.")
		neptunedata_startMldataProcessingJobCmd.Flags().String("sagemaker-iam-role-arn", "", "The ARN of an IAM role for SageMaker execution.")
		neptunedata_startMldataProcessingJobCmd.Flags().String("security-group-ids", "", "The VPC security group IDs.")
		neptunedata_startMldataProcessingJobCmd.Flags().String("subnets", "", "The IDs of the subnets in the Neptune VPC.")
		neptunedata_startMldataProcessingJobCmd.Flags().String("volume-encryption-kmskey", "", "The Amazon Key Management Service (Amazon KMS) key that SageMaker uses to encrypt data on the storage volume attached to the ML compute instances that run the training job.")
		neptunedata_startMldataProcessingJobCmd.MarkFlagRequired("input-data-s3-location")
		neptunedata_startMldataProcessingJobCmd.MarkFlagRequired("processed-data-s3-location")
	})
	neptunedataCmd.AddCommand(neptunedata_startMldataProcessingJobCmd)
}
