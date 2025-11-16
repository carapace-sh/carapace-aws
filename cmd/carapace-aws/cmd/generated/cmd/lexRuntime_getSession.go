package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexRuntime_getSessionCmd = &cobra.Command{
	Use:   "get-session",
	Short: "Returns session information for a specified bot, alias, and user ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexRuntime_getSessionCmd).Standalone()

	lexRuntime_getSessionCmd.Flags().String("bot-alias", "", "The alias in use for the bot that contains the session data.")
	lexRuntime_getSessionCmd.Flags().String("bot-name", "", "The name of the bot that contains the session data.")
	lexRuntime_getSessionCmd.Flags().String("checkpoint-label-filter", "", "A string used to filter the intents returned in the `recentIntentSummaryView` structure.")
	lexRuntime_getSessionCmd.Flags().String("user-id", "", "The ID of the client application user.")
	lexRuntime_getSessionCmd.MarkFlagRequired("bot-alias")
	lexRuntime_getSessionCmd.MarkFlagRequired("bot-name")
	lexRuntime_getSessionCmd.MarkFlagRequired("user-id")
	lexRuntimeCmd.AddCommand(lexRuntime_getSessionCmd)
}
