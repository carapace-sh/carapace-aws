package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifybackend_createBackendCmd = &cobra.Command{
	Use:   "create-backend",
	Short: "This operation creates a backend for an Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifybackend_createBackendCmd).Standalone()

	amplifybackend_createBackendCmd.Flags().String("app-id", "", "The app ID.")
	amplifybackend_createBackendCmd.Flags().String("app-name", "", "The name of the app.")
	amplifybackend_createBackendCmd.Flags().String("backend-environment-name", "", "The name of the backend environment.")
	amplifybackend_createBackendCmd.Flags().String("resource-config", "", "The resource configuration for creating a backend.")
	amplifybackend_createBackendCmd.Flags().String("resource-name", "", "The name of the resource.")
	amplifybackend_createBackendCmd.MarkFlagRequired("app-id")
	amplifybackend_createBackendCmd.MarkFlagRequired("app-name")
	amplifybackend_createBackendCmd.MarkFlagRequired("backend-environment-name")
	amplifybackendCmd.AddCommand(amplifybackend_createBackendCmd)
}
