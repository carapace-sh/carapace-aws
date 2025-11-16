package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_getBotCmd = &cobra.Command{
	Use:   "get-bot",
	Short: "Retrieves details for the specified bot, such as bot email address, bot type, status, and display name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_getBotCmd).Standalone()

	chime_getBotCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
	chime_getBotCmd.Flags().String("bot-id", "", "The bot ID.")
	chime_getBotCmd.MarkFlagRequired("account-id")
	chime_getBotCmd.MarkFlagRequired("bot-id")
	chimeCmd.AddCommand(chime_getBotCmd)
}
