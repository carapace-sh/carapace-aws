package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_createFunctionCmd = &cobra.Command{
	Use:   "create-function",
	Short: "Creates a Lambda function.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_createFunctionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lambda_createFunctionCmd).Standalone()

		lambda_createFunctionCmd.Flags().String("architectures", "", "The instruction set architecture that the function supports.")
		lambda_createFunctionCmd.Flags().String("code", "", "The code for the function.")
		lambda_createFunctionCmd.Flags().String("code-signing-config-arn", "", "To enable code signing for this function, specify the ARN of a code-signing configuration.")
		lambda_createFunctionCmd.Flags().String("dead-letter-config", "", "A dead-letter queue configuration that specifies the queue or topic where Lambda sends asynchronous events when they fail processing.")
		lambda_createFunctionCmd.Flags().String("description", "", "A description of the function.")
		lambda_createFunctionCmd.Flags().String("environment", "", "Environment variables that are accessible from function code during execution.")
		lambda_createFunctionCmd.Flags().String("ephemeral-storage", "", "The size of the function's `/tmp` directory in MB.")
		lambda_createFunctionCmd.Flags().String("file-system-configs", "", "Connection settings for an Amazon EFS file system.")
		lambda_createFunctionCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function.")
		lambda_createFunctionCmd.Flags().String("handler", "", "The name of the method within your code that Lambda calls to run your function.")
		lambda_createFunctionCmd.Flags().String("image-config", "", "Container image [configuration values](https://docs.aws.amazon.com/lambda/latest/dg/images-create.html#images-parms) that override the values in the container image Dockerfile.")
		lambda_createFunctionCmd.Flags().String("kmskey-arn", "", "The ARN of the Key Management Service (KMS) customer managed key that's used to encrypt the following resources:")
		lambda_createFunctionCmd.Flags().String("layers", "", "A list of [function layers](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html) to add to the function's execution environment.")
		lambda_createFunctionCmd.Flags().String("logging-config", "", "The function's Amazon CloudWatch Logs configuration settings.")
		lambda_createFunctionCmd.Flags().String("memory-size", "", "The amount of [memory available to the function](https://docs.aws.amazon.com/lambda/latest/dg/configuration-function-common.html#configuration-memory-console) at runtime.")
		lambda_createFunctionCmd.Flags().Bool("no-publish", false, "Set to true to publish the first version of the function during creation.")
		lambda_createFunctionCmd.Flags().String("package-type", "", "The type of deployment package.")
		lambda_createFunctionCmd.Flags().Bool("publish", false, "Set to true to publish the first version of the function during creation.")
		lambda_createFunctionCmd.Flags().String("role", "", "The Amazon Resource Name (ARN) of the function's execution role.")
		lambda_createFunctionCmd.Flags().String("runtime", "", "The identifier of the function's [runtime](https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html). Runtime is required if the deployment package is a .zip file archive.")
		lambda_createFunctionCmd.Flags().String("snap-start", "", "The function's [SnapStart](https://docs.aws.amazon.com/lambda/latest/dg/snapstart.html) setting.")
		lambda_createFunctionCmd.Flags().String("tags", "", "A list of [tags](https://docs.aws.amazon.com/lambda/latest/dg/tagging.html) to apply to the function.")
		lambda_createFunctionCmd.Flags().String("timeout", "", "The amount of time (in seconds) that Lambda allows a function to run before stopping it.")
		lambda_createFunctionCmd.Flags().String("tracing-config", "", "Set `Mode` to `Active` to sample and trace a subset of incoming requests with [X-Ray](https://docs.aws.amazon.com/lambda/latest/dg/services-xray.html).")
		lambda_createFunctionCmd.Flags().String("vpc-config", "", "For network connectivity to Amazon Web Services resources in a VPC, specify a list of security groups and subnets in the VPC.")
		lambda_createFunctionCmd.MarkFlagRequired("code")
		lambda_createFunctionCmd.MarkFlagRequired("function-name")
		lambda_createFunctionCmd.Flag("no-publish").Hidden = true
		lambda_createFunctionCmd.MarkFlagRequired("role")
	})
	lambdaCmd.AddCommand(lambda_createFunctionCmd)
}
