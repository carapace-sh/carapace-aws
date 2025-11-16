package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateContactRoutingDataCmd = &cobra.Command{
	Use:   "update-contact-routing-data",
	Short: "Updates routing priority and age on the contact (**QueuePriority** and **QueueTimeAdjustmentInSeconds**).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateContactRoutingDataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_updateContactRoutingDataCmd).Standalone()

		connect_updateContactRoutingDataCmd.Flags().String("contact-id", "", "The identifier of the contact in this instance of Amazon Connect.")
		connect_updateContactRoutingDataCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_updateContactRoutingDataCmd.Flags().String("queue-priority", "", "Priority of the contact in the queue.")
		connect_updateContactRoutingDataCmd.Flags().String("queue-time-adjustment-seconds", "", "The number of seconds to add or subtract from the contact's routing age.")
		connect_updateContactRoutingDataCmd.Flags().String("routing-criteria", "", "Updates the routing criteria on the contact.")
		connect_updateContactRoutingDataCmd.MarkFlagRequired("contact-id")
		connect_updateContactRoutingDataCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_updateContactRoutingDataCmd)
}
