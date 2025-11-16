package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrServerless_startJobRunCmd = &cobra.Command{
	Use:   "start-job-run",
	Short: "Starts a job run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrServerless_startJobRunCmd).Standalone()

	emrServerless_startJobRunCmd.Flags().String("application-id", "", "The ID of the application on which to run the job.")
	emrServerless_startJobRunCmd.Flags().String("client-token", "", "The client idempotency token of the job run to start.")
	emrServerless_startJobRunCmd.Flags().String("configuration-overrides", "", "The configuration overrides for the job run.")
	emrServerless_startJobRunCmd.Flags().String("execution-iam-policy", "", "You can pass an optional IAM policy.")
	emrServerless_startJobRunCmd.Flags().String("execution-role-arn", "", "The execution role ARN for the job run.")
	emrServerless_startJobRunCmd.Flags().String("execution-timeout-minutes", "", "The maximum duration for the job run to run.")
	emrServerless_startJobRunCmd.Flags().String("job-driver", "", "The job driver for the job run.")
	emrServerless_startJobRunCmd.Flags().String("mode", "", "The mode of the job run when it starts.")
	emrServerless_startJobRunCmd.Flags().String("name", "", "The optional job run name.")
	emrServerless_startJobRunCmd.Flags().String("retry-policy", "", "The retry policy when job run starts.")
	emrServerless_startJobRunCmd.Flags().String("tags", "", "The tags assigned to the job run.")
	emrServerless_startJobRunCmd.MarkFlagRequired("application-id")
	emrServerless_startJobRunCmd.MarkFlagRequired("client-token")
	emrServerless_startJobRunCmd.MarkFlagRequired("execution-role-arn")
	emrServerlessCmd.AddCommand(emrServerless_startJobRunCmd)
}
