package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_createQueueEnvironmentCmd = &cobra.Command{
	Use:   "create-queue-environment",
	Short: "Creates an environment for a queue that defines how jobs in the queue run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_createQueueEnvironmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_createQueueEnvironmentCmd).Standalone()

		deadline_createQueueEnvironmentCmd.Flags().String("client-token", "", "The unique token which the server uses to recognize retries of the same request.")
		deadline_createQueueEnvironmentCmd.Flags().String("farm-id", "", "The farm ID of the farm to connect to the environment.")
		deadline_createQueueEnvironmentCmd.Flags().String("priority", "", "Sets the priority of the environments in the queue from 0 to 10,000, where 0 is the highest priority (activated first and deactivated last).")
		deadline_createQueueEnvironmentCmd.Flags().String("queue-id", "", "The queue ID to connect the queue and environment.")
		deadline_createQueueEnvironmentCmd.Flags().String("template", "", "The environment template to use in the queue.")
		deadline_createQueueEnvironmentCmd.Flags().String("template-type", "", "The template's file type, `JSON` or `YAML`.")
		deadline_createQueueEnvironmentCmd.MarkFlagRequired("farm-id")
		deadline_createQueueEnvironmentCmd.MarkFlagRequired("priority")
		deadline_createQueueEnvironmentCmd.MarkFlagRequired("queue-id")
		deadline_createQueueEnvironmentCmd.MarkFlagRequired("template")
		deadline_createQueueEnvironmentCmd.MarkFlagRequired("template-type")
	})
	deadlineCmd.AddCommand(deadline_createQueueEnvironmentCmd)
}
