package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifybackend_createTokenCmd = &cobra.Command{
	Use:   "create-token",
	Short: "Generates a one-time challenge code to authenticate a user into your Amplify Admin UI.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifybackend_createTokenCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplifybackend_createTokenCmd).Standalone()

		amplifybackend_createTokenCmd.Flags().String("app-id", "", "The app ID.")
		amplifybackend_createTokenCmd.MarkFlagRequired("app-id")
	})
	amplifybackendCmd.AddCommand(amplifybackend_createTokenCmd)
}
