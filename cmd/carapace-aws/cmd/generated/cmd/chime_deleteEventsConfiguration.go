package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_deleteEventsConfigurationCmd = &cobra.Command{
	Use:   "delete-events-configuration",
	Short: "Deletes the events configuration that allows a bot to receive outgoing events.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_deleteEventsConfigurationCmd).Standalone()

	chime_deleteEventsConfigurationCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
	chime_deleteEventsConfigurationCmd.Flags().String("bot-id", "", "The bot ID.")
	chime_deleteEventsConfigurationCmd.MarkFlagRequired("account-id")
	chime_deleteEventsConfigurationCmd.MarkFlagRequired("bot-id")
	chimeCmd.AddCommand(chime_deleteEventsConfigurationCmd)
}
