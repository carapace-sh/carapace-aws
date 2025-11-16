package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securitylake_createSubscriberNotificationCmd = &cobra.Command{
	Use:   "create-subscriber-notification",
	Short: "Notifies the subscriber when new data is written to the data lake for the sources that the subscriber consumes in Security Lake.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securitylake_createSubscriberNotificationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securitylake_createSubscriberNotificationCmd).Standalone()

		securitylake_createSubscriberNotificationCmd.Flags().String("configuration", "", "Specify the configuration using which you want to create the subscriber notification.")
		securitylake_createSubscriberNotificationCmd.Flags().String("subscriber-id", "", "The subscriber ID for the notification subscription.")
		securitylake_createSubscriberNotificationCmd.MarkFlagRequired("configuration")
		securitylake_createSubscriberNotificationCmd.MarkFlagRequired("subscriber-id")
	})
	securitylakeCmd.AddCommand(securitylake_createSubscriberNotificationCmd)
}
