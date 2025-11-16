package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_startJobRunCmd = &cobra.Command{
	Use:   "start-job-run",
	Short: "Starts a job run using a job definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_startJobRunCmd).Standalone()

	glue_startJobRunCmd.Flags().String("allocated-capacity", "", "This field is deprecated.")
	glue_startJobRunCmd.Flags().String("arguments", "", "The job arguments associated with this run.")
	glue_startJobRunCmd.Flags().String("execution-class", "", "Indicates whether the job is run with a standard or flexible execution class.")
	glue_startJobRunCmd.Flags().String("execution-role-session-policy", "", "This inline session policy to the StartJobRun API allows you to dynamically restrict the permissions of the specified execution role for the scope of the job, without requiring the creation of additional IAM roles.")
	glue_startJobRunCmd.Flags().String("job-name", "", "The name of the job definition to use.")
	glue_startJobRunCmd.Flags().String("job-run-id", "", "The ID of a previous `JobRun` to retry.")
	glue_startJobRunCmd.Flags().String("job-run-queuing-enabled", "", "Specifies whether job run queuing is enabled for the job run.")
	glue_startJobRunCmd.Flags().String("max-capacity", "", "For Glue version 1.0 or earlier jobs, using the standard worker type, the number of Glue data processing units (DPUs) that can be allocated when this job runs.")
	glue_startJobRunCmd.Flags().String("notification-property", "", "Specifies configuration properties of a job run notification.")
	glue_startJobRunCmd.Flags().String("number-of-workers", "", "The number of workers of a defined `workerType` that are allocated when a job runs.")
	glue_startJobRunCmd.Flags().String("security-configuration", "", "The name of the `SecurityConfiguration` structure to be used with this job run.")
	glue_startJobRunCmd.Flags().String("timeout", "", "The `JobRun` timeout in minutes.")
	glue_startJobRunCmd.Flags().String("worker-type", "", "The type of predefined worker that is allocated when a job runs.")
	glue_startJobRunCmd.MarkFlagRequired("job-name")
	glueCmd.AddCommand(glue_startJobRunCmd)
}
