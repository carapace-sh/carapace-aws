package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotJobsData_getPendingJobExecutionsCmd = &cobra.Command{
	Use:   "get-pending-job-executions",
	Short: "Gets the list of all jobs for a thing that are not in a terminal status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotJobsData_getPendingJobExecutionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotJobsData_getPendingJobExecutionsCmd).Standalone()

		iotJobsData_getPendingJobExecutionsCmd.Flags().String("thing-name", "", "The name of the thing that is executing the job.")
		iotJobsData_getPendingJobExecutionsCmd.MarkFlagRequired("thing-name")
	})
	iotJobsDataCmd.AddCommand(iotJobsData_getPendingJobExecutionsCmd)
}
