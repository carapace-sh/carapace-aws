package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrServerless_getJobRunCmd = &cobra.Command{
	Use:   "get-job-run",
	Short: "Displays detailed information about a job run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrServerless_getJobRunCmd).Standalone()

	emrServerless_getJobRunCmd.Flags().String("application-id", "", "The ID of the application on which the job run is submitted.")
	emrServerless_getJobRunCmd.Flags().String("attempt", "", "An optimal parameter that indicates the amount of attempts for the job.")
	emrServerless_getJobRunCmd.Flags().String("job-run-id", "", "The ID of the job run.")
	emrServerless_getJobRunCmd.MarkFlagRequired("application-id")
	emrServerless_getJobRunCmd.MarkFlagRequired("job-run-id")
	emrServerlessCmd.AddCommand(emrServerless_getJobRunCmd)
}
