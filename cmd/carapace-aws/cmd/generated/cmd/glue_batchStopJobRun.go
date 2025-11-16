package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_batchStopJobRunCmd = &cobra.Command{
	Use:   "batch-stop-job-run",
	Short: "Stops one or more job runs for a specified job definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_batchStopJobRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_batchStopJobRunCmd).Standalone()

		glue_batchStopJobRunCmd.Flags().String("job-name", "", "The name of the job definition for which to stop job runs.")
		glue_batchStopJobRunCmd.Flags().String("job-run-ids", "", "A list of the `JobRunIds` that should be stopped for that job definition.")
		glue_batchStopJobRunCmd.MarkFlagRequired("job-name")
		glue_batchStopJobRunCmd.MarkFlagRequired("job-run-ids")
	})
	glueCmd.AddCommand(glue_batchStopJobRunCmd)
}
