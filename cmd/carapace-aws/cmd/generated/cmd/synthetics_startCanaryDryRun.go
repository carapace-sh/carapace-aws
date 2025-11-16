package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var synthetics_startCanaryDryRunCmd = &cobra.Command{
	Use:   "start-canary-dry-run",
	Short: "Use this operation to start a dry run for a canary that has already been created",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(synthetics_startCanaryDryRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(synthetics_startCanaryDryRunCmd).Standalone()

		synthetics_startCanaryDryRunCmd.Flags().String("artifact-config", "", "")
		synthetics_startCanaryDryRunCmd.Flags().String("artifact-s3-location", "", "The location in Amazon S3 where Synthetics stores artifacts from the test runs of this canary.")
		synthetics_startCanaryDryRunCmd.Flags().String("browser-configs", "", "A structure that specifies the browser type to use for a canary run.")
		synthetics_startCanaryDryRunCmd.Flags().String("code", "", "")
		synthetics_startCanaryDryRunCmd.Flags().String("execution-role-arn", "", "The ARN of the IAM role to be used to run the canary.")
		synthetics_startCanaryDryRunCmd.Flags().String("failure-retention-period-in-days", "", "The number of days to retain data about failed runs of this canary.")
		synthetics_startCanaryDryRunCmd.Flags().String("name", "", "The name of the canary that you want to dry run.")
		synthetics_startCanaryDryRunCmd.Flags().String("provisioned-resource-cleanup", "", "Specifies whether to also delete the Lambda functions and layers used by this canary when the canary is deleted.")
		synthetics_startCanaryDryRunCmd.Flags().String("run-config", "", "")
		synthetics_startCanaryDryRunCmd.Flags().String("runtime-version", "", "Specifies the runtime version to use for the canary.")
		synthetics_startCanaryDryRunCmd.Flags().String("success-retention-period-in-days", "", "The number of days to retain data about successful runs of this canary.")
		synthetics_startCanaryDryRunCmd.Flags().String("visual-reference", "", "")
		synthetics_startCanaryDryRunCmd.Flags().String("visual-references", "", "A list of visual reference configurations for the canary, one for each browser type that the canary is configured to run on.")
		synthetics_startCanaryDryRunCmd.Flags().String("vpc-config", "", "")
		synthetics_startCanaryDryRunCmd.MarkFlagRequired("name")
	})
	syntheticsCmd.AddCommand(synthetics_startCanaryDryRunCmd)
}
