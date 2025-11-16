package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listJobsCmd = &cobra.Command{
	Use:   "list-jobs",
	Short: "Lists jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_listJobsCmd).Standalone()

		iot_listJobsCmd.Flags().String("max-results", "", "The maximum number of results to return per request.")
		iot_listJobsCmd.Flags().String("namespace-id", "", "The namespace used to indicate that a job is a customer-managed job.")
		iot_listJobsCmd.Flags().String("next-token", "", "The token to retrieve the next set of results.")
		iot_listJobsCmd.Flags().String("status", "", "An optional filter that lets you search for jobs that have the specified status.")
		iot_listJobsCmd.Flags().String("target-selection", "", "Specifies whether the job will continue to run (CONTINUOUS), or will be complete after all those things specified as targets have completed the job (SNAPSHOT).")
		iot_listJobsCmd.Flags().String("thing-group-id", "", "A filter that limits the returned jobs to those for the specified group.")
		iot_listJobsCmd.Flags().String("thing-group-name", "", "A filter that limits the returned jobs to those for the specified group.")
	})
	iotCmd.AddCommand(iot_listJobsCmd)
}
