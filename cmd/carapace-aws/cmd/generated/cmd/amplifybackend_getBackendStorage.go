package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifybackend_getBackendStorageCmd = &cobra.Command{
	Use:   "get-backend-storage",
	Short: "Gets details for a backend storage resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifybackend_getBackendStorageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplifybackend_getBackendStorageCmd).Standalone()

		amplifybackend_getBackendStorageCmd.Flags().String("app-id", "", "The app ID.")
		amplifybackend_getBackendStorageCmd.Flags().String("backend-environment-name", "", "The name of the backend environment.")
		amplifybackend_getBackendStorageCmd.Flags().String("resource-name", "", "The name of the storage resource.")
		amplifybackend_getBackendStorageCmd.MarkFlagRequired("app-id")
		amplifybackend_getBackendStorageCmd.MarkFlagRequired("backend-environment-name")
		amplifybackend_getBackendStorageCmd.MarkFlagRequired("resource-name")
	})
	amplifybackendCmd.AddCommand(amplifybackend_getBackendStorageCmd)
}
