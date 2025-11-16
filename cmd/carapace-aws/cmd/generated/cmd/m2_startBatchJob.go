package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var m2_startBatchJobCmd = &cobra.Command{
	Use:   "start-batch-job",
	Short: "Starts a batch job and returns the unique identifier of this execution of the batch job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(m2_startBatchJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(m2_startBatchJobCmd).Standalone()

		m2_startBatchJobCmd.Flags().String("application-id", "", "The unique identifier of the application associated with this batch job.")
		m2_startBatchJobCmd.Flags().String("auth-secrets-manager-arn", "", "The Amazon Web Services Secrets Manager containing user's credentials for authentication and authorization for Start Batch Job execution operation.")
		m2_startBatchJobCmd.Flags().String("batch-job-identifier", "", "The unique identifier of the batch job.")
		m2_startBatchJobCmd.Flags().String("job-params", "", "The collection of batch job parameters.")
		m2_startBatchJobCmd.MarkFlagRequired("application-id")
		m2_startBatchJobCmd.MarkFlagRequired("batch-job-identifier")
	})
	m2Cmd.AddCommand(m2_startBatchJobCmd)
}
