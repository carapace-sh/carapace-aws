package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateQueueHoursOfOperationCmd = &cobra.Command{
	Use:   "update-queue-hours-of-operation",
	Short: "Updates the hours of operation for the specified queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateQueueHoursOfOperationCmd).Standalone()

	connect_updateQueueHoursOfOperationCmd.Flags().String("hours-of-operation-id", "", "The identifier for the hours of operation.")
	connect_updateQueueHoursOfOperationCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_updateQueueHoursOfOperationCmd.Flags().String("queue-id", "", "The identifier for the queue.")
	connect_updateQueueHoursOfOperationCmd.MarkFlagRequired("hours-of-operation-id")
	connect_updateQueueHoursOfOperationCmd.MarkFlagRequired("instance-id")
	connect_updateQueueHoursOfOperationCmd.MarkFlagRequired("queue-id")
	connectCmd.AddCommand(connect_updateQueueHoursOfOperationCmd)
}
