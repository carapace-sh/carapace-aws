package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifybackend_createBackendConfigCmd = &cobra.Command{
	Use:   "create-backend-config",
	Short: "Creates a config object for a backend.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifybackend_createBackendConfigCmd).Standalone()

	amplifybackend_createBackendConfigCmd.Flags().String("app-id", "", "The app ID.")
	amplifybackend_createBackendConfigCmd.Flags().String("backend-manager-app-id", "", "The app ID for the backend manager.")
	amplifybackend_createBackendConfigCmd.MarkFlagRequired("app-id")
	amplifybackendCmd.AddCommand(amplifybackend_createBackendConfigCmd)
}
