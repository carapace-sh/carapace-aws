package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_startMlmodelTransformJobCmd = &cobra.Command{
	Use:   "start-mlmodel-transform-job",
	Short: "Creates a new model transform job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_startMlmodelTransformJobCmd).Standalone()

	neptunedata_startMlmodelTransformJobCmd.Flags().String("base-processing-instance-type", "", "The type of ML instance used in preparing and managing training of ML models.")
	neptunedata_startMlmodelTransformJobCmd.Flags().String("base-processing-instance-volume-size-in-gb", "", "The disk volume size of the training instance in gigabytes.")
	neptunedata_startMlmodelTransformJobCmd.Flags().String("custom-model-transform-parameters", "", "Configuration information for a model transform using a custom model.")
	neptunedata_startMlmodelTransformJobCmd.Flags().String("data-processing-job-id", "", "The job ID of a completed data-processing job.")
	neptunedata_startMlmodelTransformJobCmd.Flags().String("id", "", "A unique identifier for the new job.")
	neptunedata_startMlmodelTransformJobCmd.Flags().String("ml-model-training-job-id", "", "The job ID of a completed model-training job.")
	neptunedata_startMlmodelTransformJobCmd.Flags().String("model-transform-output-s3-location", "", "The location in Amazon S3 where the model artifacts are to be stored.")
	neptunedata_startMlmodelTransformJobCmd.Flags().String("neptune-iam-role-arn", "", "The ARN of an IAM role that provides Neptune access to SageMaker and Amazon S3 resources.")
	neptunedata_startMlmodelTransformJobCmd.Flags().String("s3-output-encryption-kmskey", "", "The Amazon Key Management Service (KMS) key that SageMaker uses to encrypt the output of the processing job.")
	neptunedata_startMlmodelTransformJobCmd.Flags().String("sagemaker-iam-role-arn", "", "The ARN of an IAM role for SageMaker execution.")
	neptunedata_startMlmodelTransformJobCmd.Flags().String("security-group-ids", "", "The VPC security group IDs.")
	neptunedata_startMlmodelTransformJobCmd.Flags().String("subnets", "", "The IDs of the subnets in the Neptune VPC.")
	neptunedata_startMlmodelTransformJobCmd.Flags().String("training-job-name", "", "The name of a completed SageMaker training job.")
	neptunedata_startMlmodelTransformJobCmd.Flags().String("volume-encryption-kmskey", "", "The Amazon Key Management Service (KMS) key that SageMaker uses to encrypt data on the storage volume attached to the ML compute instances that run the training job.")
	neptunedata_startMlmodelTransformJobCmd.MarkFlagRequired("model-transform-output-s3-location")
	neptunedataCmd.AddCommand(neptunedata_startMlmodelTransformJobCmd)
}
