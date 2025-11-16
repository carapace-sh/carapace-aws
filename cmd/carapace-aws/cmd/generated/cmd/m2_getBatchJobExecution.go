package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var m2_getBatchJobExecutionCmd = &cobra.Command{
	Use:   "get-batch-job-execution",
	Short: "Gets the details of a specific batch job execution for a specific application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(m2_getBatchJobExecutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(m2_getBatchJobExecutionCmd).Standalone()

		m2_getBatchJobExecutionCmd.Flags().String("application-id", "", "The identifier of the application.")
		m2_getBatchJobExecutionCmd.Flags().String("execution-id", "", "The unique identifier of the batch job execution.")
		m2_getBatchJobExecutionCmd.MarkFlagRequired("application-id")
		m2_getBatchJobExecutionCmd.MarkFlagRequired("execution-id")
	})
	m2Cmd.AddCommand(m2_getBatchJobExecutionCmd)
}
