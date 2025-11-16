package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_putEventsConfigurationCmd = &cobra.Command{
	Use:   "put-events-configuration",
	Short: "Creates an events configuration that allows a bot to receive outgoing events sent by Amazon Chime.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_putEventsConfigurationCmd).Standalone()

	chime_putEventsConfigurationCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
	chime_putEventsConfigurationCmd.Flags().String("bot-id", "", "The bot ID.")
	chime_putEventsConfigurationCmd.Flags().String("lambda-function-arn", "", "Lambda function ARN that allows the bot to receive outgoing events.")
	chime_putEventsConfigurationCmd.Flags().String("outbound-events-httpsendpoint", "", "HTTPS endpoint that allows the bot to receive outgoing events.")
	chime_putEventsConfigurationCmd.MarkFlagRequired("account-id")
	chime_putEventsConfigurationCmd.MarkFlagRequired("bot-id")
	chimeCmd.AddCommand(chime_putEventsConfigurationCmd)
}
