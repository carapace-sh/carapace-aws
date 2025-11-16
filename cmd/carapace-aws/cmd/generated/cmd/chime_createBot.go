package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_createBotCmd = &cobra.Command{
	Use:   "create-bot",
	Short: "Creates a bot for an Amazon Chime Enterprise account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_createBotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chime_createBotCmd).Standalone()

		chime_createBotCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
		chime_createBotCmd.Flags().String("display-name", "", "The bot display name.")
		chime_createBotCmd.Flags().String("domain", "", "The domain of the Amazon Chime Enterprise account.")
		chime_createBotCmd.MarkFlagRequired("account-id")
		chime_createBotCmd.MarkFlagRequired("display-name")
	})
	chimeCmd.AddCommand(chime_createBotCmd)
}
