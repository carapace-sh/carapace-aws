package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifybackend_deleteTokenCmd = &cobra.Command{
	Use:   "delete-token",
	Short: "Deletes the challenge token based on the given appId and sessionId.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifybackend_deleteTokenCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplifybackend_deleteTokenCmd).Standalone()

		amplifybackend_deleteTokenCmd.Flags().String("app-id", "", "The app ID.")
		amplifybackend_deleteTokenCmd.Flags().String("session-id", "", "The session ID.")
		amplifybackend_deleteTokenCmd.MarkFlagRequired("app-id")
		amplifybackend_deleteTokenCmd.MarkFlagRequired("session-id")
	})
	amplifybackendCmd.AddCommand(amplifybackend_deleteTokenCmd)
}
