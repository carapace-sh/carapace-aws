package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_updateFunctionCodeCmd = &cobra.Command{
	Use:   "update-function-code",
	Short: "Updates a Lambda function's code.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_updateFunctionCodeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lambda_updateFunctionCodeCmd).Standalone()

		lambda_updateFunctionCodeCmd.Flags().String("architectures", "", "The instruction set architecture that the function supports.")
		lambda_updateFunctionCodeCmd.Flags().Bool("dry-run", false, "Set to true to validate the request parameters and access permissions without modifying the function code.")
		lambda_updateFunctionCodeCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function.")
		lambda_updateFunctionCodeCmd.Flags().String("image-uri", "", "URI of a container image in the Amazon ECR registry.")
		lambda_updateFunctionCodeCmd.Flags().Bool("no-dry-run", false, "Set to true to validate the request parameters and access permissions without modifying the function code.")
		lambda_updateFunctionCodeCmd.Flags().Bool("no-publish", false, "Set to true to publish a new version of the function after updating the code.")
		lambda_updateFunctionCodeCmd.Flags().Bool("publish", false, "Set to true to publish a new version of the function after updating the code.")
		lambda_updateFunctionCodeCmd.Flags().String("revision-id", "", "Update the function only if the revision ID matches the ID that's specified.")
		lambda_updateFunctionCodeCmd.Flags().String("s3-bucket", "", "An Amazon S3 bucket in the same Amazon Web Services Region as your function.")
		lambda_updateFunctionCodeCmd.Flags().String("s3-key", "", "The Amazon S3 key of the deployment package.")
		lambda_updateFunctionCodeCmd.Flags().String("s3-object-version", "", "For versioned objects, the version of the deployment package object to use.")
		lambda_updateFunctionCodeCmd.Flags().String("source-kmskey-arn", "", "The ARN of the Key Management Service (KMS) customer managed key that's used to encrypt your function's .zip deployment package.")
		lambda_updateFunctionCodeCmd.Flags().String("zip-file", "", "The base64-encoded contents of the deployment package.")
		lambda_updateFunctionCodeCmd.MarkFlagRequired("function-name")
		lambda_updateFunctionCodeCmd.Flag("no-dry-run").Hidden = true
		lambda_updateFunctionCodeCmd.Flag("no-publish").Hidden = true
	})
	lambdaCmd.AddCommand(lambda_updateFunctionCodeCmd)
}
