package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifybackend_deleteBackendCmd = &cobra.Command{
	Use:   "delete-backend",
	Short: "Removes an existing environment from your Amplify project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifybackend_deleteBackendCmd).Standalone()

	amplifybackend_deleteBackendCmd.Flags().String("app-id", "", "The app ID.")
	amplifybackend_deleteBackendCmd.Flags().String("backend-environment-name", "", "The name of the backend environment.")
	amplifybackend_deleteBackendCmd.MarkFlagRequired("app-id")
	amplifybackend_deleteBackendCmd.MarkFlagRequired("backend-environment-name")
	amplifybackendCmd.AddCommand(amplifybackend_deleteBackendCmd)
}
