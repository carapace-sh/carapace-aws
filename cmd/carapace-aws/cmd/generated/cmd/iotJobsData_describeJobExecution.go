package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotJobsData_describeJobExecutionCmd = &cobra.Command{
	Use:   "describe-job-execution",
	Short: "Gets details of a job execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotJobsData_describeJobExecutionCmd).Standalone()

	iotJobsData_describeJobExecutionCmd.Flags().String("execution-number", "", "Optional.")
	iotJobsData_describeJobExecutionCmd.Flags().String("include-job-document", "", "Optional.")
	iotJobsData_describeJobExecutionCmd.Flags().String("job-id", "", "The unique identifier assigned to this job when it was created.")
	iotJobsData_describeJobExecutionCmd.Flags().String("thing-name", "", "The thing name associated with the device the job execution is running on.")
	iotJobsData_describeJobExecutionCmd.MarkFlagRequired("job-id")
	iotJobsData_describeJobExecutionCmd.MarkFlagRequired("thing-name")
	iotJobsDataCmd.AddCommand(iotJobsData_describeJobExecutionCmd)
}
