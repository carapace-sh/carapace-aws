package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_createWorkerCmd = &cobra.Command{
	Use:   "create-worker",
	Short: "Creates a worker.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_createWorkerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_createWorkerCmd).Standalone()

		deadline_createWorkerCmd.Flags().String("client-token", "", "The unique token which the server uses to recognize retries of the same request.")
		deadline_createWorkerCmd.Flags().String("farm-id", "", "The farm ID of the farm to connect to the worker.")
		deadline_createWorkerCmd.Flags().String("fleet-id", "", "The fleet ID to connect to the worker.")
		deadline_createWorkerCmd.Flags().String("host-properties", "", "The IP address and host name of the worker.")
		deadline_createWorkerCmd.Flags().String("tags", "", "Each tag consists of a tag key and a tag value.")
		deadline_createWorkerCmd.MarkFlagRequired("farm-id")
		deadline_createWorkerCmd.MarkFlagRequired("fleet-id")
	})
	deadlineCmd.AddCommand(deadline_createWorkerCmd)
}
