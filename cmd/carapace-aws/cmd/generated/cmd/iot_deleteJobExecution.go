package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_deleteJobExecutionCmd = &cobra.Command{
	Use:   "delete-job-execution",
	Short: "Deletes a job execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_deleteJobExecutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_deleteJobExecutionCmd).Standalone()

		iot_deleteJobExecutionCmd.Flags().String("execution-number", "", "The ID of the job execution to be deleted.")
		iot_deleteJobExecutionCmd.Flags().String("force", "", "(Optional) When true, you can delete a job execution which is \"IN\\_PROGRESS\".")
		iot_deleteJobExecutionCmd.Flags().String("job-id", "", "The ID of the job whose execution on a particular device will be deleted.")
		iot_deleteJobExecutionCmd.Flags().String("namespace-id", "", "The namespace used to indicate that a job is a customer-managed job.")
		iot_deleteJobExecutionCmd.Flags().String("thing-name", "", "The name of the thing whose job execution will be deleted.")
		iot_deleteJobExecutionCmd.MarkFlagRequired("execution-number")
		iot_deleteJobExecutionCmd.MarkFlagRequired("job-id")
		iot_deleteJobExecutionCmd.MarkFlagRequired("thing-name")
	})
	iotCmd.AddCommand(iot_deleteJobExecutionCmd)
}
