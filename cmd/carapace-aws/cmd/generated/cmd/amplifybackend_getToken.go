package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifybackend_getTokenCmd = &cobra.Command{
	Use:   "get-token",
	Short: "Gets the challenge token based on the given appId and sessionId.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifybackend_getTokenCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplifybackend_getTokenCmd).Standalone()

		amplifybackend_getTokenCmd.Flags().String("app-id", "", "The app ID.")
		amplifybackend_getTokenCmd.Flags().String("session-id", "", "The session ID.")
		amplifybackend_getTokenCmd.MarkFlagRequired("app-id")
		amplifybackend_getTokenCmd.MarkFlagRequired("session-id")
	})
	amplifybackendCmd.AddCommand(amplifybackend_getTokenCmd)
}
