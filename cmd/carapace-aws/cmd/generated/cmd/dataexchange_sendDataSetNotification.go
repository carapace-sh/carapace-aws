package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchange_sendDataSetNotificationCmd = &cobra.Command{
	Use:   "send-data-set-notification",
	Short: "The type of event associated with the data set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchange_sendDataSetNotificationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dataexchange_sendDataSetNotificationCmd).Standalone()

		dataexchange_sendDataSetNotificationCmd.Flags().String("client-token", "", "Idempotency key for the notification, this key allows us to deduplicate notifications that are sent in quick succession erroneously.")
		dataexchange_sendDataSetNotificationCmd.Flags().String("comment", "", "Free-form text field for providers to add information about their notifications.")
		dataexchange_sendDataSetNotificationCmd.Flags().String("data-set-id", "", "Affected data set of the notification.")
		dataexchange_sendDataSetNotificationCmd.Flags().String("details", "", "Extra details specific to this notification type.")
		dataexchange_sendDataSetNotificationCmd.Flags().String("scope", "", "Affected scope of this notification such as the underlying resources affected by the notification event.")
		dataexchange_sendDataSetNotificationCmd.Flags().String("type", "", "The type of the notification.")
		dataexchange_sendDataSetNotificationCmd.MarkFlagRequired("data-set-id")
		dataexchange_sendDataSetNotificationCmd.MarkFlagRequired("type")
	})
	dataexchangeCmd.AddCommand(dataexchange_sendDataSetNotificationCmd)
}
