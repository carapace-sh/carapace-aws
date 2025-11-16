package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifybackend_getBackendCmd = &cobra.Command{
	Use:   "get-backend",
	Short: "Provides project-level details for your Amplify UI project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifybackend_getBackendCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplifybackend_getBackendCmd).Standalone()

		amplifybackend_getBackendCmd.Flags().String("app-id", "", "The app ID.")
		amplifybackend_getBackendCmd.Flags().String("backend-environment-name", "", "The name of the backend environment.")
		amplifybackend_getBackendCmd.MarkFlagRequired("app-id")
	})
	amplifybackendCmd.AddCommand(amplifybackend_getBackendCmd)
}
