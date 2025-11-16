package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var synthetics_updateCanaryCmd = &cobra.Command{
	Use:   "update-canary",
	Short: "Updates the configuration of a canary that has already been created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(synthetics_updateCanaryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(synthetics_updateCanaryCmd).Standalone()

		synthetics_updateCanaryCmd.Flags().String("artifact-config", "", "A structure that contains the configuration for canary artifacts, including the encryption-at-rest settings for artifacts that the canary uploads to Amazon S3.")
		synthetics_updateCanaryCmd.Flags().String("artifact-s3-location", "", "The location in Amazon S3 where Synthetics stores artifacts from the test runs of this canary.")
		synthetics_updateCanaryCmd.Flags().String("browser-configs", "", "A structure that specifies the browser type to use for a canary run.")
		synthetics_updateCanaryCmd.Flags().String("code", "", "A structure that includes the entry point from which the canary should start running your script.")
		synthetics_updateCanaryCmd.Flags().String("dry-run-id", "", "Update the existing canary using the updated configurations from the DryRun associated with the DryRunId.")
		synthetics_updateCanaryCmd.Flags().String("execution-role-arn", "", "The ARN of the IAM role to be used to run the canary.")
		synthetics_updateCanaryCmd.Flags().String("failure-retention-period-in-days", "", "The number of days to retain data about failed runs of this canary.")
		synthetics_updateCanaryCmd.Flags().String("name", "", "The name of the canary that you want to update.")
		synthetics_updateCanaryCmd.Flags().String("provisioned-resource-cleanup", "", "Specifies whether to also delete the Lambda functions and layers used by this canary when the canary is deleted.")
		synthetics_updateCanaryCmd.Flags().String("run-config", "", "A structure that contains the timeout value that is used for each individual run of the canary.")
		synthetics_updateCanaryCmd.Flags().String("runtime-version", "", "Specifies the runtime version to use for the canary.")
		synthetics_updateCanaryCmd.Flags().String("schedule", "", "A structure that contains information about how often the canary is to run, and when these runs are to stop.")
		synthetics_updateCanaryCmd.Flags().String("success-retention-period-in-days", "", "The number of days to retain data about successful runs of this canary.")
		synthetics_updateCanaryCmd.Flags().String("visual-reference", "", "Defines the screenshots to use as the baseline for comparisons during visual monitoring comparisons during future runs of this canary.")
		synthetics_updateCanaryCmd.Flags().String("visual-references", "", "A list of visual reference configurations for the canary, one for each browser type that the canary is configured to run on.")
		synthetics_updateCanaryCmd.Flags().String("vpc-config", "", "If this canary is to test an endpoint in a VPC, this structure contains information about the subnet and security groups of the VPC endpoint.")
		synthetics_updateCanaryCmd.MarkFlagRequired("name")
	})
	syntheticsCmd.AddCommand(synthetics_updateCanaryCmd)
}
