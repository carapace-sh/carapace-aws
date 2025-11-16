package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var synthetics_createCanaryCmd = &cobra.Command{
	Use:   "create-canary",
	Short: "Creates a canary.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(synthetics_createCanaryCmd).Standalone()

	synthetics_createCanaryCmd.Flags().String("artifact-config", "", "A structure that contains the configuration for canary artifacts, including the encryption-at-rest settings for artifacts that the canary uploads to Amazon S3.")
	synthetics_createCanaryCmd.Flags().String("artifact-s3-location", "", "The location in Amazon S3 where Synthetics stores artifacts from the test runs of this canary.")
	synthetics_createCanaryCmd.Flags().String("browser-configs", "", "CloudWatch Synthetics now supports multibrowser canaries for `syn-nodejs-puppeteer-11.0` and `syn-nodejs-playwright-3.0` runtimes.")
	synthetics_createCanaryCmd.Flags().String("code", "", "A structure that includes the entry point from which the canary should start running your script.")
	synthetics_createCanaryCmd.Flags().String("execution-role-arn", "", "The ARN of the IAM role to be used to run the canary.")
	synthetics_createCanaryCmd.Flags().String("failure-retention-period-in-days", "", "The number of days to retain data about failed runs of this canary.")
	synthetics_createCanaryCmd.Flags().String("name", "", "The name for this canary.")
	synthetics_createCanaryCmd.Flags().String("provisioned-resource-cleanup", "", "Specifies whether to also delete the Lambda functions and layers used by this canary when the canary is deleted.")
	synthetics_createCanaryCmd.Flags().String("resources-to-replicate-tags", "", "To have the tags that you apply to this canary also be applied to the Lambda function that the canary uses, specify this parameter with the value `lambda-function`.")
	synthetics_createCanaryCmd.Flags().String("run-config", "", "A structure that contains the configuration for individual canary runs, such as timeout value and environment variables.")
	synthetics_createCanaryCmd.Flags().String("runtime-version", "", "Specifies the runtime version to use for the canary.")
	synthetics_createCanaryCmd.Flags().String("schedule", "", "A structure that contains information about how often the canary is to run and when these test runs are to stop.")
	synthetics_createCanaryCmd.Flags().String("success-retention-period-in-days", "", "The number of days to retain data about successful runs of this canary.")
	synthetics_createCanaryCmd.Flags().String("tags", "", "A list of key-value pairs to associate with the canary.")
	synthetics_createCanaryCmd.Flags().String("vpc-config", "", "If this canary is to test an endpoint in a VPC, this structure contains information about the subnet and security groups of the VPC endpoint.")
	synthetics_createCanaryCmd.MarkFlagRequired("artifact-s3-location")
	synthetics_createCanaryCmd.MarkFlagRequired("code")
	synthetics_createCanaryCmd.MarkFlagRequired("execution-role-arn")
	synthetics_createCanaryCmd.MarkFlagRequired("name")
	synthetics_createCanaryCmd.MarkFlagRequired("runtime-version")
	synthetics_createCanaryCmd.MarkFlagRequired("schedule")
	syntheticsCmd.AddCommand(synthetics_createCanaryCmd)
}
