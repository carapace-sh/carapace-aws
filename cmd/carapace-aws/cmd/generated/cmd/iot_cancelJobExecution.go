package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_cancelJobExecutionCmd = &cobra.Command{
	Use:   "cancel-job-execution",
	Short: "Cancels the execution of a job for a given thing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_cancelJobExecutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_cancelJobExecutionCmd).Standalone()

		iot_cancelJobExecutionCmd.Flags().String("expected-version", "", "(Optional) The expected current version of the job execution.")
		iot_cancelJobExecutionCmd.Flags().String("force", "", "(Optional) If `true` the job execution will be canceled if it has status IN\\_PROGRESS or QUEUED, otherwise the job execution will be canceled only if it has status QUEUED.")
		iot_cancelJobExecutionCmd.Flags().String("job-id", "", "The ID of the job to be canceled.")
		iot_cancelJobExecutionCmd.Flags().String("status-details", "", "A collection of name/value pairs that describe the status of the job execution.")
		iot_cancelJobExecutionCmd.Flags().String("thing-name", "", "The name of the thing whose execution of the job will be canceled.")
		iot_cancelJobExecutionCmd.MarkFlagRequired("job-id")
		iot_cancelJobExecutionCmd.MarkFlagRequired("thing-name")
	})
	iotCmd.AddCommand(iot_cancelJobExecutionCmd)
}
