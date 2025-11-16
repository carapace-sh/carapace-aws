package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrServerless_cancelJobRunCmd = &cobra.Command{
	Use:   "cancel-job-run",
	Short: "Cancels a job run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrServerless_cancelJobRunCmd).Standalone()

	emrServerless_cancelJobRunCmd.Flags().String("application-id", "", "The ID of the application on which the job run will be canceled.")
	emrServerless_cancelJobRunCmd.Flags().String("job-run-id", "", "The ID of the job run to cancel.")
	emrServerless_cancelJobRunCmd.Flags().String("shutdown-grace-period-in-seconds", "", "The duration in seconds to wait before forcefully terminating the job after cancellation is requested.")
	emrServerless_cancelJobRunCmd.MarkFlagRequired("application-id")
	emrServerless_cancelJobRunCmd.MarkFlagRequired("job-run-id")
	emrServerlessCmd.AddCommand(emrServerless_cancelJobRunCmd)
}
