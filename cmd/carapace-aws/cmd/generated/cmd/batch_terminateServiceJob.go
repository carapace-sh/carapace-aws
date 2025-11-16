package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_terminateServiceJobCmd = &cobra.Command{
	Use:   "terminate-service-job",
	Short: "Terminates a service job in a job queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_terminateServiceJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(batch_terminateServiceJobCmd).Standalone()

		batch_terminateServiceJobCmd.Flags().String("job-id", "", "The service job ID of the service job to terminate.")
		batch_terminateServiceJobCmd.Flags().String("reason", "", "A message to attach to the service job that explains the reason for canceling it.")
		batch_terminateServiceJobCmd.MarkFlagRequired("job-id")
		batch_terminateServiceJobCmd.MarkFlagRequired("reason")
	})
	batchCmd.AddCommand(batch_terminateServiceJobCmd)
}
