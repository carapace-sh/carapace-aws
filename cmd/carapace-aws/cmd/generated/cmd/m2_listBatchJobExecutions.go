package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var m2_listBatchJobExecutionsCmd = &cobra.Command{
	Use:   "list-batch-job-executions",
	Short: "Lists historical, current, and scheduled batch job executions for a specific application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(m2_listBatchJobExecutionsCmd).Standalone()

	m2_listBatchJobExecutionsCmd.Flags().String("application-id", "", "The unique identifier of the application.")
	m2_listBatchJobExecutionsCmd.Flags().String("execution-ids", "", "The unique identifier of each batch job execution.")
	m2_listBatchJobExecutionsCmd.Flags().String("job-name", "", "The name of each batch job execution.")
	m2_listBatchJobExecutionsCmd.Flags().String("max-results", "", "The maximum number of batch job executions to return.")
	m2_listBatchJobExecutionsCmd.Flags().String("next-token", "", "A pagination token to control the number of batch job executions displayed in the list.")
	m2_listBatchJobExecutionsCmd.Flags().String("started-after", "", "The time after which the batch job executions started.")
	m2_listBatchJobExecutionsCmd.Flags().String("started-before", "", "The time before the batch job executions started.")
	m2_listBatchJobExecutionsCmd.Flags().String("status", "", "The status of the batch job executions.")
	m2_listBatchJobExecutionsCmd.MarkFlagRequired("application-id")
	m2Cmd.AddCommand(m2_listBatchJobExecutionsCmd)
}
