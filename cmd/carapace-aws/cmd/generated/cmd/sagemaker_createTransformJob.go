package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createTransformJobCmd = &cobra.Command{
	Use:   "create-transform-job",
	Short: "Starts a transform job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createTransformJobCmd).Standalone()

	sagemaker_createTransformJobCmd.Flags().String("batch-strategy", "", "Specifies the number of records to include in a mini-batch for an HTTP inference request.")
	sagemaker_createTransformJobCmd.Flags().String("data-capture-config", "", "Configuration to control how SageMaker captures inference data.")
	sagemaker_createTransformJobCmd.Flags().String("data-processing", "", "The data structure used to specify the data to be used for inference in a batch transform job and to associate the data that is relevant to the prediction results in the output.")
	sagemaker_createTransformJobCmd.Flags().String("environment", "", "The environment variables to set in the Docker container.")
	sagemaker_createTransformJobCmd.Flags().String("experiment-config", "", "")
	sagemaker_createTransformJobCmd.Flags().String("max-concurrent-transforms", "", "The maximum number of parallel requests that can be sent to each instance in a transform job.")
	sagemaker_createTransformJobCmd.Flags().String("max-payload-in-mb", "", "The maximum allowed size of the payload, in MB.")
	sagemaker_createTransformJobCmd.Flags().String("model-client-config", "", "Configures the timeout and maximum number of retries for processing a transform job invocation.")
	sagemaker_createTransformJobCmd.Flags().String("model-name", "", "The name of the model that you want to use for the transform job.")
	sagemaker_createTransformJobCmd.Flags().String("tags", "", "(Optional) An array of key-value pairs.")
	sagemaker_createTransformJobCmd.Flags().String("transform-input", "", "Describes the input source and the way the transform job consumes it.")
	sagemaker_createTransformJobCmd.Flags().String("transform-job-name", "", "The name of the transform job.")
	sagemaker_createTransformJobCmd.Flags().String("transform-output", "", "Describes the results of the transform job.")
	sagemaker_createTransformJobCmd.Flags().String("transform-resources", "", "Describes the resources, including ML instance types and ML instance count, to use for the transform job.")
	sagemaker_createTransformJobCmd.MarkFlagRequired("model-name")
	sagemaker_createTransformJobCmd.MarkFlagRequired("transform-input")
	sagemaker_createTransformJobCmd.MarkFlagRequired("transform-job-name")
	sagemaker_createTransformJobCmd.MarkFlagRequired("transform-output")
	sagemaker_createTransformJobCmd.MarkFlagRequired("transform-resources")
	sagemakerCmd.AddCommand(sagemaker_createTransformJobCmd)
}
