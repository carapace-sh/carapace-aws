package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifybackend_removeAllBackendsCmd = &cobra.Command{
	Use:   "remove-all-backends",
	Short: "Removes all backend environments from your Amplify project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifybackend_removeAllBackendsCmd).Standalone()

	amplifybackend_removeAllBackendsCmd.Flags().String("app-id", "", "The app ID.")
	amplifybackend_removeAllBackendsCmd.Flags().String("clean-amplify-app", "", "Cleans up the Amplify Console app if this value is set to true.")
	amplifybackend_removeAllBackendsCmd.MarkFlagRequired("app-id")
	amplifybackendCmd.AddCommand(amplifybackend_removeAllBackendsCmd)
}
