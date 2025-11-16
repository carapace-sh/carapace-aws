package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_updateFunctionConfigurationCmd = &cobra.Command{
	Use:   "update-function-configuration",
	Short: "Modify the version-specific settings of a Lambda function.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_updateFunctionConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lambda_updateFunctionConfigurationCmd).Standalone()

		lambda_updateFunctionConfigurationCmd.Flags().String("dead-letter-config", "", "A dead-letter queue configuration that specifies the queue or topic where Lambda sends asynchronous events when they fail processing.")
		lambda_updateFunctionConfigurationCmd.Flags().String("description", "", "A description of the function.")
		lambda_updateFunctionConfigurationCmd.Flags().String("environment", "", "Environment variables that are accessible from function code during execution.")
		lambda_updateFunctionConfigurationCmd.Flags().String("ephemeral-storage", "", "The size of the function's `/tmp` directory in MB.")
		lambda_updateFunctionConfigurationCmd.Flags().String("file-system-configs", "", "Connection settings for an Amazon EFS file system.")
		lambda_updateFunctionConfigurationCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function.")
		lambda_updateFunctionConfigurationCmd.Flags().String("handler", "", "The name of the method within your code that Lambda calls to run your function.")
		lambda_updateFunctionConfigurationCmd.Flags().String("image-config", "", "[Container image configuration values](https://docs.aws.amazon.com/lambda/latest/dg/images-create.html#images-parms) that override the values in the container image Docker file.")
		lambda_updateFunctionConfigurationCmd.Flags().String("kmskey-arn", "", "The ARN of the Key Management Service (KMS) customer managed key that's used to encrypt the following resources:")
		lambda_updateFunctionConfigurationCmd.Flags().String("layers", "", "A list of [function layers](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html) to add to the function's execution environment.")
		lambda_updateFunctionConfigurationCmd.Flags().String("logging-config", "", "The function's Amazon CloudWatch Logs configuration settings.")
		lambda_updateFunctionConfigurationCmd.Flags().String("memory-size", "", "The amount of [memory available to the function](https://docs.aws.amazon.com/lambda/latest/dg/configuration-function-common.html#configuration-memory-console) at runtime.")
		lambda_updateFunctionConfigurationCmd.Flags().String("revision-id", "", "Update the function only if the revision ID matches the ID that's specified.")
		lambda_updateFunctionConfigurationCmd.Flags().String("role", "", "The Amazon Resource Name (ARN) of the function's execution role.")
		lambda_updateFunctionConfigurationCmd.Flags().String("runtime", "", "The identifier of the function's [runtime](https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html). Runtime is required if the deployment package is a .zip file archive.")
		lambda_updateFunctionConfigurationCmd.Flags().String("snap-start", "", "The function's [SnapStart](https://docs.aws.amazon.com/lambda/latest/dg/snapstart.html) setting.")
		lambda_updateFunctionConfigurationCmd.Flags().String("timeout", "", "The amount of time (in seconds) that Lambda allows a function to run before stopping it.")
		lambda_updateFunctionConfigurationCmd.Flags().String("tracing-config", "", "Set `Mode` to `Active` to sample and trace a subset of incoming requests with [X-Ray](https://docs.aws.amazon.com/lambda/latest/dg/services-xray.html).")
		lambda_updateFunctionConfigurationCmd.Flags().String("vpc-config", "", "For network connectivity to Amazon Web Services resources in a VPC, specify a list of security groups and subnets in the VPC.")
		lambda_updateFunctionConfigurationCmd.MarkFlagRequired("function-name")
	})
	lambdaCmd.AddCommand(lambda_updateFunctionConfigurationCmd)
}
