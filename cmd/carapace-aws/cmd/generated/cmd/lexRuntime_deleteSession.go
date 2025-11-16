package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexRuntime_deleteSessionCmd = &cobra.Command{
	Use:   "delete-session",
	Short: "Removes session information for a specified bot, alias, and user ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexRuntime_deleteSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexRuntime_deleteSessionCmd).Standalone()

		lexRuntime_deleteSessionCmd.Flags().String("bot-alias", "", "The alias in use for the bot that contains the session data.")
		lexRuntime_deleteSessionCmd.Flags().String("bot-name", "", "The name of the bot that contains the session data.")
		lexRuntime_deleteSessionCmd.Flags().String("user-id", "", "The identifier of the user associated with the session data.")
		lexRuntime_deleteSessionCmd.MarkFlagRequired("bot-alias")
		lexRuntime_deleteSessionCmd.MarkFlagRequired("bot-name")
		lexRuntime_deleteSessionCmd.MarkFlagRequired("user-id")
	})
	lexRuntimeCmd.AddCommand(lexRuntime_deleteSessionCmd)
}
