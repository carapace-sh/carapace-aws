package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var m2_cancelBatchJobExecutionCmd = &cobra.Command{
	Use:   "cancel-batch-job-execution",
	Short: "Cancels the running of a specific batch job execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(m2_cancelBatchJobExecutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(m2_cancelBatchJobExecutionCmd).Standalone()

		m2_cancelBatchJobExecutionCmd.Flags().String("application-id", "", "The unique identifier of the application.")
		m2_cancelBatchJobExecutionCmd.Flags().String("auth-secrets-manager-arn", "", "The Amazon Web Services Secrets Manager containing user's credentials for authentication and authorization for Cancel Batch Job Execution operation.")
		m2_cancelBatchJobExecutionCmd.Flags().String("execution-id", "", "The unique identifier of the batch job execution.")
		m2_cancelBatchJobExecutionCmd.MarkFlagRequired("application-id")
		m2_cancelBatchJobExecutionCmd.MarkFlagRequired("execution-id")
	})
	m2Cmd.AddCommand(m2_cancelBatchJobExecutionCmd)
}
