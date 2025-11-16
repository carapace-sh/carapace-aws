package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifybackend_updateBackendStorageCmd = &cobra.Command{
	Use:   "update-backend-storage",
	Short: "Updates an existing backend storage resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifybackend_updateBackendStorageCmd).Standalone()

	amplifybackend_updateBackendStorageCmd.Flags().String("app-id", "", "The app ID.")
	amplifybackend_updateBackendStorageCmd.Flags().String("backend-environment-name", "", "The name of the backend environment.")
	amplifybackend_updateBackendStorageCmd.Flags().String("resource-config", "", "The resource configuration for updating backend storage.")
	amplifybackend_updateBackendStorageCmd.Flags().String("resource-name", "", "The name of the storage resource.")
	amplifybackend_updateBackendStorageCmd.MarkFlagRequired("app-id")
	amplifybackend_updateBackendStorageCmd.MarkFlagRequired("backend-environment-name")
	amplifybackend_updateBackendStorageCmd.MarkFlagRequired("resource-config")
	amplifybackend_updateBackendStorageCmd.MarkFlagRequired("resource-name")
	amplifybackendCmd.AddCommand(amplifybackend_updateBackendStorageCmd)
}
