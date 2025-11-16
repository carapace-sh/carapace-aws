package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotJobsData_updateJobExecutionCmd = &cobra.Command{
	Use:   "update-job-execution",
	Short: "Updates the status of a job execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotJobsData_updateJobExecutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotJobsData_updateJobExecutionCmd).Standalone()

		iotJobsData_updateJobExecutionCmd.Flags().String("execution-number", "", "Optional.")
		iotJobsData_updateJobExecutionCmd.Flags().String("expected-version", "", "Optional.")
		iotJobsData_updateJobExecutionCmd.Flags().String("include-job-document", "", "Optional.")
		iotJobsData_updateJobExecutionCmd.Flags().String("include-job-execution-state", "", "Optional.")
		iotJobsData_updateJobExecutionCmd.Flags().String("job-id", "", "The unique identifier assigned to this job when it was created.")
		iotJobsData_updateJobExecutionCmd.Flags().String("status", "", "The new status for the job execution (IN\\_PROGRESS, FAILED, SUCCESS, or REJECTED).")
		iotJobsData_updateJobExecutionCmd.Flags().String("status-details", "", "Optional.")
		iotJobsData_updateJobExecutionCmd.Flags().String("step-timeout-in-minutes", "", "Specifies the amount of time this device has to finish execution of this job.")
		iotJobsData_updateJobExecutionCmd.Flags().String("thing-name", "", "The name of the thing associated with the device.")
		iotJobsData_updateJobExecutionCmd.MarkFlagRequired("job-id")
		iotJobsData_updateJobExecutionCmd.MarkFlagRequired("status")
		iotJobsData_updateJobExecutionCmd.MarkFlagRequired("thing-name")
	})
	iotJobsDataCmd.AddCommand(iotJobsData_updateJobExecutionCmd)
}
