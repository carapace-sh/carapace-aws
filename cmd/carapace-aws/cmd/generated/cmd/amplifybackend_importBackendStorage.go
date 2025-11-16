package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifybackend_importBackendStorageCmd = &cobra.Command{
	Use:   "import-backend-storage",
	Short: "Imports an existing backend storage resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifybackend_importBackendStorageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplifybackend_importBackendStorageCmd).Standalone()

		amplifybackend_importBackendStorageCmd.Flags().String("app-id", "", "The app ID.")
		amplifybackend_importBackendStorageCmd.Flags().String("backend-environment-name", "", "The name of the backend environment.")
		amplifybackend_importBackendStorageCmd.Flags().String("bucket-name", "", "The name of the S3 bucket.")
		amplifybackend_importBackendStorageCmd.Flags().String("service-name", "", "The name of the storage service.")
		amplifybackend_importBackendStorageCmd.MarkFlagRequired("app-id")
		amplifybackend_importBackendStorageCmd.MarkFlagRequired("backend-environment-name")
		amplifybackend_importBackendStorageCmd.MarkFlagRequired("service-name")
	})
	amplifybackendCmd.AddCommand(amplifybackend_importBackendStorageCmd)
}
