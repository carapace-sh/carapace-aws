package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_getEventsConfigurationCmd = &cobra.Command{
	Use:   "get-events-configuration",
	Short: "Gets details for an events configuration that allows a bot to receive outgoing events, such as an HTTPS endpoint or Lambda function ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_getEventsConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chime_getEventsConfigurationCmd).Standalone()

		chime_getEventsConfigurationCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
		chime_getEventsConfigurationCmd.Flags().String("bot-id", "", "The bot ID.")
		chime_getEventsConfigurationCmd.MarkFlagRequired("account-id")
		chime_getEventsConfigurationCmd.MarkFlagRequired("bot-id")
	})
	chimeCmd.AddCommand(chime_getEventsConfigurationCmd)
}
