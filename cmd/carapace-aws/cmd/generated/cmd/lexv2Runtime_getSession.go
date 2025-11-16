package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Runtime_getSessionCmd = &cobra.Command{
	Use:   "get-session",
	Short: "Returns session information for a specified bot, alias, and user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Runtime_getSessionCmd).Standalone()

	lexv2Runtime_getSessionCmd.Flags().String("bot-alias-id", "", "The alias identifier in use for the bot that contains the session data.")
	lexv2Runtime_getSessionCmd.Flags().String("bot-id", "", "The identifier of the bot that contains the session data.")
	lexv2Runtime_getSessionCmd.Flags().String("locale-id", "", "The locale where the session is in use.")
	lexv2Runtime_getSessionCmd.Flags().String("session-id", "", "The identifier of the session to return.")
	lexv2Runtime_getSessionCmd.MarkFlagRequired("bot-alias-id")
	lexv2Runtime_getSessionCmd.MarkFlagRequired("bot-id")
	lexv2Runtime_getSessionCmd.MarkFlagRequired("locale-id")
	lexv2Runtime_getSessionCmd.MarkFlagRequired("session-id")
	lexv2RuntimeCmd.AddCommand(lexv2Runtime_getSessionCmd)
}
