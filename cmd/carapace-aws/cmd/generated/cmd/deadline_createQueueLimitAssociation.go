package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_createQueueLimitAssociationCmd = &cobra.Command{
	Use:   "create-queue-limit-association",
	Short: "Associates a limit with a particular queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_createQueueLimitAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_createQueueLimitAssociationCmd).Standalone()

		deadline_createQueueLimitAssociationCmd.Flags().String("farm-id", "", "The unique identifier of the farm that contains the queue and limit to associate.")
		deadline_createQueueLimitAssociationCmd.Flags().String("limit-id", "", "The unique identifier of the limit to associate with the queue.")
		deadline_createQueueLimitAssociationCmd.Flags().String("queue-id", "", "The unique identifier of the queue to associate with the limit.")
		deadline_createQueueLimitAssociationCmd.MarkFlagRequired("farm-id")
		deadline_createQueueLimitAssociationCmd.MarkFlagRequired("limit-id")
		deadline_createQueueLimitAssociationCmd.MarkFlagRequired("queue-id")
	})
	deadlineCmd.AddCommand(deadline_createQueueLimitAssociationCmd)
}
