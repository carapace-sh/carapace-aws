package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifybackend_deleteBackendStorageCmd = &cobra.Command{
	Use:   "delete-backend-storage",
	Short: "Removes the specified backend storage resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifybackend_deleteBackendStorageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplifybackend_deleteBackendStorageCmd).Standalone()

		amplifybackend_deleteBackendStorageCmd.Flags().String("app-id", "", "The app ID.")
		amplifybackend_deleteBackendStorageCmd.Flags().String("backend-environment-name", "", "The name of the backend environment.")
		amplifybackend_deleteBackendStorageCmd.Flags().String("resource-name", "", "The name of the storage resource.")
		amplifybackend_deleteBackendStorageCmd.Flags().String("service-name", "", "The name of the storage service.")
		amplifybackend_deleteBackendStorageCmd.MarkFlagRequired("app-id")
		amplifybackend_deleteBackendStorageCmd.MarkFlagRequired("backend-environment-name")
		amplifybackend_deleteBackendStorageCmd.MarkFlagRequired("resource-name")
		amplifybackend_deleteBackendStorageCmd.MarkFlagRequired("service-name")
	})
	amplifybackendCmd.AddCommand(amplifybackend_deleteBackendStorageCmd)
}
