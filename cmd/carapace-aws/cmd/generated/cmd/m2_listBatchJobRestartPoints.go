package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var m2_listBatchJobRestartPointsCmd = &cobra.Command{
	Use:   "list-batch-job-restart-points",
	Short: "Lists all the job steps for a JCL file to restart a batch job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(m2_listBatchJobRestartPointsCmd).Standalone()

	m2_listBatchJobRestartPointsCmd.Flags().String("application-id", "", "The unique identifier of the application.")
	m2_listBatchJobRestartPointsCmd.Flags().String("auth-secrets-manager-arn", "", "The Amazon Web Services Secrets Manager containing user's credentials for authentication and authorization for List Batch Job Restart Points operation.")
	m2_listBatchJobRestartPointsCmd.Flags().String("execution-id", "", "The unique identifier of the batch job execution.")
	m2_listBatchJobRestartPointsCmd.MarkFlagRequired("application-id")
	m2_listBatchJobRestartPointsCmd.MarkFlagRequired("execution-id")
	m2Cmd.AddCommand(m2_listBatchJobRestartPointsCmd)
}
