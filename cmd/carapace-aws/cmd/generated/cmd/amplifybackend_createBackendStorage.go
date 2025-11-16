package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifybackend_createBackendStorageCmd = &cobra.Command{
	Use:   "create-backend-storage",
	Short: "Creates a backend storage resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifybackend_createBackendStorageCmd).Standalone()

	amplifybackend_createBackendStorageCmd.Flags().String("app-id", "", "The app ID.")
	amplifybackend_createBackendStorageCmd.Flags().String("backend-environment-name", "", "The name of the backend environment.")
	amplifybackend_createBackendStorageCmd.Flags().String("resource-config", "", "The resource configuration for creating backend storage.")
	amplifybackend_createBackendStorageCmd.Flags().String("resource-name", "", "The name of the storage resource.")
	amplifybackend_createBackendStorageCmd.MarkFlagRequired("app-id")
	amplifybackend_createBackendStorageCmd.MarkFlagRequired("backend-environment-name")
	amplifybackend_createBackendStorageCmd.MarkFlagRequired("resource-config")
	amplifybackend_createBackendStorageCmd.MarkFlagRequired("resource-name")
	amplifybackendCmd.AddCommand(amplifybackend_createBackendStorageCmd)
}
