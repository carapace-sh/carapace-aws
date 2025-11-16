package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createLabelingJobCmd = &cobra.Command{
	Use:   "create-labeling-job",
	Short: "Creates a job that uses workers to label the data objects in your input dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createLabelingJobCmd).Standalone()

	sagemaker_createLabelingJobCmd.Flags().String("human-task-config", "", "Configures the labeling task and how it is presented to workers; including, but not limited to price, keywords, and batch size (task count).")
	sagemaker_createLabelingJobCmd.Flags().String("input-config", "", "Input data for the labeling job, such as the Amazon S3 location of the data objects and the location of the manifest file that describes the data objects.")
	sagemaker_createLabelingJobCmd.Flags().String("label-attribute-name", "", "The attribute name to use for the label in the output manifest file.")
	sagemaker_createLabelingJobCmd.Flags().String("label-category-config-s3-uri", "", "The S3 URI of the file, referred to as a *label category configuration file*, that defines the categories used to label the data objects.")
	sagemaker_createLabelingJobCmd.Flags().String("labeling-job-algorithms-config", "", "Configures the information required to perform automated data labeling.")
	sagemaker_createLabelingJobCmd.Flags().String("labeling-job-name", "", "The name of the labeling job.")
	sagemaker_createLabelingJobCmd.Flags().String("output-config", "", "The location of the output data and the Amazon Web Services Key Management Service key ID for the key used to encrypt the output data, if any.")
	sagemaker_createLabelingJobCmd.Flags().String("role-arn", "", "The Amazon Resource Number (ARN) that Amazon SageMaker assumes to perform tasks on your behalf during data labeling.")
	sagemaker_createLabelingJobCmd.Flags().String("stopping-conditions", "", "A set of conditions for stopping the labeling job.")
	sagemaker_createLabelingJobCmd.Flags().String("tags", "", "An array of key/value pairs.")
	sagemaker_createLabelingJobCmd.MarkFlagRequired("human-task-config")
	sagemaker_createLabelingJobCmd.MarkFlagRequired("input-config")
	sagemaker_createLabelingJobCmd.MarkFlagRequired("label-attribute-name")
	sagemaker_createLabelingJobCmd.MarkFlagRequired("labeling-job-name")
	sagemaker_createLabelingJobCmd.MarkFlagRequired("output-config")
	sagemaker_createLabelingJobCmd.MarkFlagRequired("role-arn")
	sagemakerCmd.AddCommand(sagemaker_createLabelingJobCmd)
}
