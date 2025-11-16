package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotJobsData_startNextPendingJobExecutionCmd = &cobra.Command{
	Use:   "start-next-pending-job-execution",
	Short: "Gets and starts the next pending (status IN\\_PROGRESS or QUEUED) job execution for a thing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotJobsData_startNextPendingJobExecutionCmd).Standalone()

	iotJobsData_startNextPendingJobExecutionCmd.Flags().String("status-details", "", "A collection of name/value pairs that describe the status of the job execution.")
	iotJobsData_startNextPendingJobExecutionCmd.Flags().String("step-timeout-in-minutes", "", "Specifies the amount of time this device has to finish execution of this job.")
	iotJobsData_startNextPendingJobExecutionCmd.Flags().String("thing-name", "", "The name of the thing associated with the device.")
	iotJobsData_startNextPendingJobExecutionCmd.MarkFlagRequired("thing-name")
	iotJobsDataCmd.AddCommand(iotJobsData_startNextPendingJobExecutionCmd)
}
