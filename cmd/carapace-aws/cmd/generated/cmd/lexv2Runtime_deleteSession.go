package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Runtime_deleteSessionCmd = &cobra.Command{
	Use:   "delete-session",
	Short: "Removes session information for a specified bot, alias, and user ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Runtime_deleteSessionCmd).Standalone()

	lexv2Runtime_deleteSessionCmd.Flags().String("bot-alias-id", "", "The alias identifier in use for the bot that contains the session data.")
	lexv2Runtime_deleteSessionCmd.Flags().String("bot-id", "", "The identifier of the bot that contains the session data.")
	lexv2Runtime_deleteSessionCmd.Flags().String("locale-id", "", "The locale where the session is in use.")
	lexv2Runtime_deleteSessionCmd.Flags().String("session-id", "", "The identifier of the session to delete.")
	lexv2Runtime_deleteSessionCmd.MarkFlagRequired("bot-alias-id")
	lexv2Runtime_deleteSessionCmd.MarkFlagRequired("bot-id")
	lexv2Runtime_deleteSessionCmd.MarkFlagRequired("locale-id")
	lexv2Runtime_deleteSessionCmd.MarkFlagRequired("session-id")
	lexv2RuntimeCmd.AddCommand(lexv2Runtime_deleteSessionCmd)
}
