package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifybackend_deleteBackendApiCmd = &cobra.Command{
	Use:   "delete-backend-api",
	Short: "Deletes an existing backend API resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifybackend_deleteBackendApiCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplifybackend_deleteBackendApiCmd).Standalone()

		amplifybackend_deleteBackendApiCmd.Flags().String("app-id", "", "The app ID.")
		amplifybackend_deleteBackendApiCmd.Flags().String("backend-environment-name", "", "The name of the backend environment.")
		amplifybackend_deleteBackendApiCmd.Flags().String("resource-config", "", "Defines the resource configuration for the data model in your Amplify project.")
		amplifybackend_deleteBackendApiCmd.Flags().String("resource-name", "", "The name of this resource.")
		amplifybackend_deleteBackendApiCmd.MarkFlagRequired("app-id")
		amplifybackend_deleteBackendApiCmd.MarkFlagRequired("backend-environment-name")
		amplifybackend_deleteBackendApiCmd.MarkFlagRequired("resource-name")
	})
	amplifybackendCmd.AddCommand(amplifybackend_deleteBackendApiCmd)
}
