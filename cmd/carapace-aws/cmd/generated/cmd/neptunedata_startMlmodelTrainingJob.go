package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_startMlmodelTrainingJobCmd = &cobra.Command{
	Use:   "start-mlmodel-training-job",
	Short: "Creates a new Neptune ML model training job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_startMlmodelTrainingJobCmd).Standalone()

	neptunedata_startMlmodelTrainingJobCmd.Flags().String("base-processing-instance-type", "", "The type of ML instance used in preparing and managing training of ML models.")
	neptunedata_startMlmodelTrainingJobCmd.Flags().String("custom-model-training-parameters", "", "The configuration for custom model training.")
	neptunedata_startMlmodelTrainingJobCmd.Flags().String("data-processing-job-id", "", "The job ID of the completed data-processing job that has created the data that the training will work with.")
	neptunedata_startMlmodelTrainingJobCmd.Flags().Bool("enable-managed-spot-training", false, "Optimizes the cost of training machine-learning models by using Amazon Elastic Compute Cloud spot instances.")
	neptunedata_startMlmodelTrainingJobCmd.Flags().String("id", "", "A unique identifier for the new job.")
	neptunedata_startMlmodelTrainingJobCmd.Flags().String("max-hponumber-of-training-jobs", "", "Maximum total number of training jobs to start for the hyperparameter tuning job.")
	neptunedata_startMlmodelTrainingJobCmd.Flags().String("max-hpoparallel-training-jobs", "", "Maximum number of parallel training jobs to start for the hyperparameter tuning job.")
	neptunedata_startMlmodelTrainingJobCmd.Flags().String("neptune-iam-role-arn", "", "The ARN of an IAM role that provides Neptune access to SageMaker and Amazon S3 resources.")
	neptunedata_startMlmodelTrainingJobCmd.Flags().Bool("no-enable-managed-spot-training", false, "Optimizes the cost of training machine-learning models by using Amazon Elastic Compute Cloud spot instances.")
	neptunedata_startMlmodelTrainingJobCmd.Flags().String("previous-model-training-job-id", "", "The job ID of a completed model-training job that you want to update incrementally based on updated data.")
	neptunedata_startMlmodelTrainingJobCmd.Flags().String("s3-output-encryption-kmskey", "", "The Amazon Key Management Service (KMS) key that SageMaker uses to encrypt the output of the processing job.")
	neptunedata_startMlmodelTrainingJobCmd.Flags().String("sagemaker-iam-role-arn", "", "The ARN of an IAM role for SageMaker execution.This must be listed in your DB cluster parameter group or an error will occur.")
	neptunedata_startMlmodelTrainingJobCmd.Flags().String("security-group-ids", "", "The VPC security group IDs.")
	neptunedata_startMlmodelTrainingJobCmd.Flags().String("subnets", "", "The IDs of the subnets in the Neptune VPC.")
	neptunedata_startMlmodelTrainingJobCmd.Flags().String("train-model-s3-location", "", "The location in Amazon S3 where the model artifacts are to be stored.")
	neptunedata_startMlmodelTrainingJobCmd.Flags().String("training-instance-type", "", "The type of ML instance used for model training.")
	neptunedata_startMlmodelTrainingJobCmd.Flags().String("training-instance-volume-size-in-gb", "", "The disk volume size of the training instance.")
	neptunedata_startMlmodelTrainingJobCmd.Flags().String("training-time-out-in-seconds", "", "Timeout in seconds for the training job.")
	neptunedata_startMlmodelTrainingJobCmd.Flags().String("volume-encryption-kmskey", "", "The Amazon Key Management Service (KMS) key that SageMaker uses to encrypt data on the storage volume attached to the ML compute instances that run the training job.")
	neptunedata_startMlmodelTrainingJobCmd.MarkFlagRequired("data-processing-job-id")
	neptunedata_startMlmodelTrainingJobCmd.Flag("no-enable-managed-spot-training").Hidden = true
	neptunedata_startMlmodelTrainingJobCmd.MarkFlagRequired("train-model-s3-location")
	neptunedataCmd.AddCommand(neptunedata_startMlmodelTrainingJobCmd)
}
