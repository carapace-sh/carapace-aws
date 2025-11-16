package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_stopJobRunCmd = &cobra.Command{
	Use:   "stop-job-run",
	Short: "Stops a particular run of a job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_stopJobRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(databrew_stopJobRunCmd).Standalone()

		databrew_stopJobRunCmd.Flags().String("name", "", "The name of the job to be stopped.")
		databrew_stopJobRunCmd.Flags().String("run-id", "", "The ID of the job run to be stopped.")
		databrew_stopJobRunCmd.MarkFlagRequired("name")
		databrew_stopJobRunCmd.MarkFlagRequired("run-id")
	})
	databrewCmd.AddCommand(databrew_stopJobRunCmd)
}
